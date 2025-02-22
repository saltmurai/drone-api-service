import asyncio
import asyncpg
from mavsdk import System
import websockets
import json
import requests
import aioredis
import re


DB_HOST="localhost"
DB_USER="saltmurai"
DB_PASSWORD="saltmurai"
DB_NAME="gcs"

def extract_float_from_string(input_string):
    match = re.search(r"\d+\.\d+", input_string)
    if match:
        return float(match.group())
    else:
        return None


def convert_to_coordinates(x, y):
    lat = x * 1e-7
    long = y * 1e-7
    if lat > 90 or lat < -90 or long > 180 or long < -180 or lat == 0 or long == 0:
        return None
    return long, lat


# Connect to Redis database
async def connect_to_redis():
    print("Connecting to the Redis database...")
    redis = await aioredis.from_url("redis://localhost:6379")
    print("Connected to the Redis database")
    return redis


# Connect to PostgreSQL database
async def connect_to_database():
    print("Connecting to the database...")
    conn = await asyncpg.connect(
        user=DB_USER,
        password=DB_NAME,
        database=DB_NAME,
        host=DB_HOST,
        port="5432",
    )
    print("Connected to the database")
    return conn


# Query the drone system address from the database
async def get_drone_system_addresses(conn):
    print("Querying drone system addresses from the database...")
    query = "SELECT id, address FROM drones WHERE status = false"
    addresses = [record["address"] for record in await conn.fetch(query)]
    ids = [record["id"] for record in await conn.fetch(query)]
    print(ids)
    print("Drone system addresses retrieved from the database:", addresses)
    return addresses, ids


# Set the status of the drone to true in the database
async def update_drone_status(conn, address, port):
    print("Updating the drone status in the database...")
    query = "UPDATE drones SET status = true, port = $2  WHERE address = $1"
    await conn.execute(query, address, port)
    print("Drone status updated in the database")


# Set the status of the drone to false in the database
async def reset_drone_status(conn, address):
    print("Resetting the drone status in the database...")
    query = "UPDATE drones SET status = false WHERE address = $1"
    await conn.execute(query, address)
    print("Drone status reset in the database")


async def connect_to_drone(address, port):
    print("Connecting to the drone " + address)
    # create a new instance of the drone
    drone = System(port=port)
    try:
        await asyncio.wait_for(drone.connect(system_address=address), timeout=3)
        print("Connected to the drone")
        return drone
    except TimeoutError:
        raise TimeoutError("Connection to the drone timed out")
    except Exception as e:
        raise Exception("Connection to the drone failed")


async def get_position(drone, websocket, drone_id, r):
    data = {
        "drone_id": drone_id,
    }
    async for position in drone.telemetry.position():
        data["latitude_deg"] = position.latitude_deg
        data["longitude_deg"] = position.longitude_deg
        data["absolute_altitude_m"] = position.absolute_altitude_m
        data["relative_altitude_m"] = position.relative_altitude_m
        try:
            # await websocket.send(json.dumps(data))
            await r.set(f"{str(drone_id)}-postion", json.dumps(data))
        except websockets.exceptions.ConnectionClosed:
            break
        except Exception as e:
            print(e)


async def get_heading(drone, websocket, id, r):
    data = {
        "drone_id": id,
    }
    async for heading_deg in drone.telemetry.heading():
        data["heading_deg"] = extract_float_from_string(str(heading_deg))
        try:
            await r.set(f"{str(id)}-heading_deg", json.dumps(data))
        except websockets.exceptions.ConnectionClosed:
            break
        except Exception as e:
            print(e)


async def get_status(drone, websocket, id, r):
    async for status_text in drone.telemetry.StatusText():
        # data["status_text"] = extract_float_from_string(str(heading_deg))
        data = {
            "drone_id": id,
            "type": "status_text",
        }
        try:
            data["data"] = str(status_text)
            await websocket.send(json.dumps(data))
            # await r.set(f"{str(id)}-heading_deg", json.dumps(data))
        except websockets.exceptions.ConnectionClosed:
            break
        except Exception as e:
            print(e)


async def get_battery(drone, websocket, id, r):
    data = {
        "drone_id": id,
    }
    async for battery in drone.telemetry.battery():
        data["voltage_v"] = battery.voltage_v
        data["remaining_percent"] = battery.remaining_percent
        try:
            # await websocket.send(json.dumps(data))
            await r.set(f"{str(id)}-battery", json.dumps(data))
        except websockets.exceptions.ConnectionClosed:
            break
        except Exception as e:
            print(e)


async def get_flight_mode(drone, websocket, id, r):
    data = {
        "drone_id": id,
    }
    async for flight_mode in drone.telemetry.flight_mode():
        data["flight_mode"] = f"${flight_mode}"
        try:
            # await websocket.send(json.dumps(data))
            await r.set(f"{str(id)}-flight_mode", json.dumps(data))
        except websockets.exceptions.ConnectionClosed:
            break
        except Exception as e:
            print(e)
        await asyncio.sleep(1)


# Forward telemetry data to all connected WebSocket clients
async def forward_telemetry_data(websocket, path):
    conn = await connect_to_database()
    r = await connect_to_redis()
    addresses, ids = await get_drone_system_addresses(conn)
    drones = []
    tasks = []
    start_port = 50051
    i = 0
    try:
        i = 0
        for address in addresses:
            try:
                drone = await connect_to_drone(address, start_port + i)
            except Exception as e:
                continue
            drones.append(drone)
            await update_drone_status(conn, address, port=start_port + i)
            i += 1
            # await update_drone_status(conn, address)

            position_task = asyncio.create_task(
                get_position(drone, websocket, ids[addresses.index(address)], r)
            )
            battery_task = asyncio.create_task(
                get_battery(drone, websocket, ids[addresses.index(address)], r)
            )
            flight_mode_task = asyncio.create_task(
                get_flight_mode(drone, websocket, ids[addresses.index(address)], r)
            )
            heading_task = asyncio.create_task(
                get_heading(drone, websocket, ids[addresses.index(address)], r)
            )
            status_task = asyncio.create_task(
                get_status(drone, websocket, ids[addresses.index(address)], r)
            )
            tasks.append(position_task)
            tasks.append(heading_task)
            tasks.append(battery_task)
            tasks.append(flight_mode_task)
            tasks.append(status_task)

        await asyncio.gather(*tasks)
    except asyncio.CancelledError:
        print("Forwarding telemetry data cancelled")
    except Exception as e:
        print(e)
    finally:
        print("Cleaning up...")
        for drone in drones:
            drone._stop_mavsdk_server()
        for address in addresses:
            await reset_drone_status(conn, address)
        for id in ids:
            await r.delete(f"${str(id)}-postion")
            await r.delete(f"${str(id)}-battery")
            await r.delete(f"${str(id)}-flight_mode")


async def receive_client_messages(websocket, path):
    r = await connect_to_redis()
    try:
        async for message in websocket:
            try:
                data = json.loads(message)
                print(data)
                action = data.get("action")
                system_address = data.get("system_address")
                port = data.get("port")
                drone_id = data.get("drone_id")
                if not action or not system_address or not port:
                    print("Invalid message format.")
                    continue

                drone = System(port=port)
                await drone.connect(system_address=system_address)
                # Perform actions based on the received message
                # You can add different if-else conditions here based on different actions
                # For example:
                if action == "return_home":
                    try:
                        await drone.action.return_to_launch()
                    except Exception as e:
                        print(e)
                    # Perform some action on the drone with the given system_address
                elif action == "land":
                    print(data)
                    try:
                        await drone.action.land()
                    except Exception as e:
                        print(e)
                    # Perform some other action on the drone
                elif action == "hold":
                    try:
                        await drone.action.hold()
                    except Exception as e:
                        print(e)
                elif action == "takeoff":
                    print(data)
                    # Perform some other action on the drone
                    try:
                        await drone.action.set_takeoff_altitude(2.0)
                        await drone.action.arm()
                        await drone.action.takeoff()
                    except Exception as e:
                        print(e)
                elif action == "arm":
                    try:
                        await drone.action.arm()
                    except Exception as e:
                        print(e)
                elif action == "disarm":
                    try:
                        await drone.action.disarm()
                    except Exception as e:
                        print(e)
                elif action == "hold":
                    try:
                        await drone.action.hold()
                    except Exception as e:
                        print(e)
                elif action == "kill":
                    try:
                        await drone.action.kill()
                    except Exception as e:
                        print(e)
                elif action == "reboot":
                    try:
                        await drone.action.reboot()
                    except Exception as e:
                        print(e)
                elif action == "start_mission":
                    try:
                        await drone.mission.start_mission()
                    except Exception as e:
                        print(e)
                elif action == "download_mission":
                    mission_items = {
                        "drone_id": drone_id,
                        "type": "mission_items",
                    }
                    try:
                        mission = await drone.mission_raw.download_mission()
                        mission_coordinates = []
                        for mission_item in mission:
                            print(
                                convert_to_coordinates(mission_item.x, mission_item.y)
                            )
                            if (
                                convert_to_coordinates(mission_item.x, mission_item.y)
                                is not None
                            ):
                                mission_coordinates.append(
                                    convert_to_coordinates(
                                        mission_item.x, mission_item.y
                                    )
                                )
                        print(json.dumps(mission_coordinates))
                        mission_items["mission_items"] = json.dumps(mission_coordinates)
                        await websocket.send(json.dumps(mission_items))

                    except Exception as e:
                        print(e)

                else:
                    print("Unknown action received.")
                    print(data)

            except json.JSONDecodeError as e:
                print("Error decoding JSON:", e)
            except Exception as e:
                print("Error processing message:", e)

    except websockets.ConnectionClosed:
        print("Client connection closed.")


# Main program
async def main(websocket, path):
    print("New WebSocket client connected")
    await asyncio.gather(
        forward_telemetry_data(websocket, path),
        receive_client_messages(websocket, path),
    )


# Set up WebSocket server
start_server = websockets.serve(main, "0.0.0.0", 3003)

print("WebSocket server started")
try:
    asyncio.get_event_loop().run_until_complete(start_server)
    asyncio.get_event_loop().run_forever()
except KeyboardInterrupt:
    print("WebSocket server stopped")
    print("Reseting drone status in the database...")
    # called POST api http://localhost:3002/resetDrones
    res = requests.post("http://localhost:3002/resetDrones")
    if res.status_code == 200:
        print("Drone status reset in the database")
