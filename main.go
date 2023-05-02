package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/go-chi/chi/v5"

	missionv1 "github.com/saltmurai/drone-api-service/gen/mission/v1"
	"github.com/saltmurai/drone-api-service/gen/mission/v1/missionv1connect"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type MissionServer struct {
	missionv1connect.UnimplementedMissionServiceHandler
}

func (s *MissionServer) SendMission(
	ctx context.Context,
	req *connect.Request[missionv1.SendMissionRequest],
) (*connect.Response[missionv1.SendMissionResult], error) {
	id := req.Msg.GetId()
	init := req.Msg.GetInitInstructions()
	travel := req.Msg.GetTravelInstructions()
	fmt.Println(init.ProtoReflect())
	action := req.Msg.GetActionInstructions()
	fmt.Printf("Got %s %+v %v %v\n", id, init, travel, action)
	return connect.NewResponse(&missionv1.SendMissionResult{
		Success:      true,
		ErrorMessage: "none",
	}), nil
}

func main() {
	log, _ := zap.NewProduction()
	defer log.Sync()
	sugar := log.Sugar()
	sugar.Infof("Starting server on 8080")
	missioner := &MissionServer{}
	mux := http.NewServeMux()
	r := chi.NewRouter()
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("testing"))
	})
	path, handler := missionv1connect.NewMissionServiceHandler(missioner)
	mux.Handle(path, handler)
	mux.Handle("/", r)
	err := http.ListenAndServe(":3002", h2c.NewHandler(mux, &http2.Server{}))
	if err != nil {
		sugar.Error(err)
	}
}
