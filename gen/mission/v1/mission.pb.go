// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: mission/v1/mission.proto

package missionv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Termination int32

const (
	Termination_TERMINATION_AUTO Termination = 0
	Termination_TERMINATION_STD  Termination = 1
)

// Enum value maps for Termination.
var (
	Termination_name = map[int32]string{
		0: "TERMINATION_AUTO",
		1: "TERMINATION_STD",
	}
	Termination_value = map[string]int32{
		"TERMINATION_AUTO": 0,
		"TERMINATION_STD":  1,
	}
)

func (x Termination) Enum() *Termination {
	p := new(Termination)
	*p = x
	return p
}

func (x Termination) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Termination) Descriptor() protoreflect.EnumDescriptor {
	return file_mission_v1_mission_proto_enumTypes[0].Descriptor()
}

func (Termination) Type() protoreflect.EnumType {
	return &file_mission_v1_mission_proto_enumTypes[0]
}

func (x Termination) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Termination.Descriptor instead.
func (Termination) EnumDescriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{0}
}

type Action int32

const (
	Action_ACTION_TAKEOFF   Action = 0
	Action_ACTION_DISARM    Action = 1
	Action_ACTION_SELFCHECK Action = 2
	Action_ACTION_RELEASE   Action = 3
	Action_ACTION_RTLHOME   Action = 4
	Action_ACTION_HOLD      Action = 5
	Action_ACTION_AUTOLAND  Action = 6
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "ACTION_TAKEOFF",
		1: "ACTION_DISARM",
		2: "ACTION_SELFCHECK",
		3: "ACTION_RELEASE",
		4: "ACTION_RTLHOME",
		5: "ACTION_HOLD",
		6: "ACTION_AUTOLAND",
	}
	Action_value = map[string]int32{
		"ACTION_TAKEOFF":   0,
		"ACTION_DISARM":    1,
		"ACTION_SELFCHECK": 2,
		"ACTION_RELEASE":   3,
		"ACTION_RTLHOME":   4,
		"ACTION_HOLD":      5,
		"ACTION_AUTOLAND":  6,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_mission_v1_mission_proto_enumTypes[1].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_mission_v1_mission_proto_enumTypes[1]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{1}
}

type Planner int32

const (
	Planner_PLANNER_EGO      Planner = 0
	Planner_PLANNER_FAST     Planner = 1
	Planner_PLANNER_MARKER   Planner = 3
	Planner_PLANNER_SAFELAND Planner = 4
)

// Enum value maps for Planner.
var (
	Planner_name = map[int32]string{
		0: "PLANNER_EGO",
		1: "PLANNER_FAST",
		3: "PLANNER_MARKER",
		4: "PLANNER_SAFELAND",
	}
	Planner_value = map[string]int32{
		"PLANNER_EGO":      0,
		"PLANNER_FAST":     1,
		"PLANNER_MARKER":   3,
		"PLANNER_SAFELAND": 4,
	}
)

func (x Planner) Enum() *Planner {
	p := new(Planner)
	*p = x
	return p
}

func (x Planner) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Planner) Descriptor() protoreflect.EnumDescriptor {
	return file_mission_v1_mission_proto_enumTypes[2].Descriptor()
}

func (Planner) Type() protoreflect.EnumType {
	return &file_mission_v1_mission_proto_enumTypes[2]
}

func (x Planner) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Planner.Descriptor instead.
func (Planner) EnumDescriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{2}
}

type Controller int32

const (
	Controller_CONTROLLER_PX4_VELO_FB Controller = 0
	Controller_CONTROLLER_A_FB        Controller = 1
	Controller_CONTROLLER_A_FW        Controller = 2
	Controller_CONTROLLER_A_ADRJ      Controller = 3
)

// Enum value maps for Controller.
var (
	Controller_name = map[int32]string{
		0: "CONTROLLER_PX4_VELO_FB",
		1: "CONTROLLER_A_FB",
		2: "CONTROLLER_A_FW",
		3: "CONTROLLER_A_ADRJ",
	}
	Controller_value = map[string]int32{
		"CONTROLLER_PX4_VELO_FB": 0,
		"CONTROLLER_A_FB":        1,
		"CONTROLLER_A_FW":        2,
		"CONTROLLER_A_ADRJ":      3,
	}
)

func (x Controller) Enum() *Controller {
	p := new(Controller)
	*p = x
	return p
}

func (x Controller) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Controller) Descriptor() protoreflect.EnumDescriptor {
	return file_mission_v1_mission_proto_enumTypes[3].Descriptor()
}

func (Controller) Type() protoreflect.EnumType {
	return &file_mission_v1_mission_proto_enumTypes[3]
}

func (x Controller) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Controller.Descriptor instead.
func (Controller) EnumDescriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{3}
}

type InitInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define a repeated field of integers for the peripheral devices.
	Peripheral []int32 `protobuf:"varint,1,rep,packed,name=peripheral,proto3" json:"peripheral,omitempty"`
	// Define a string field for the controller.
	Controller Controller `protobuf:"varint,2,opt,name=controller,proto3,enum=mission.v1.Controller" json:"controller,omitempty"`
	// Define a string field for the standard.
	Terminate Termination `protobuf:"varint,3,opt,name=terminate,proto3,enum=mission.v1.Termination" json:"terminate,omitempty"`
}

func (x *InitInstruction) Reset() {
	*x = InitInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mission_v1_mission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitInstruction) ProtoMessage() {}

func (x *InitInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_mission_v1_mission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitInstruction.ProtoReflect.Descriptor instead.
func (*InitInstruction) Descriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{0}
}

func (x *InitInstruction) GetPeripheral() []int32 {
	if x != nil {
		return x.Peripheral
	}
	return nil
}

func (x *InitInstruction) GetController() Controller {
	if x != nil {
		return x.Controller
	}
	return Controller_CONTROLLER_PX4_VELO_FB
}

func (x *InitInstruction) GetTerminate() Termination {
	if x != nil {
		return x.Terminate
	}
	return Termination_TERMINATION_AUTO
}

// Define the message for the TRAVEL instruction.
type TravelInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define a float field for the distance.
	Planner    Planner   `protobuf:"varint,1,opt,name=planner,proto3,enum=mission.v1.Planner" json:"planner,omitempty"`
	Waypoint   []float64 `protobuf:"fixed64,2,rep,packed,name=waypoint,proto3" json:"waypoint,omitempty"`
	Constraint []float64 `protobuf:"fixed64,3,rep,packed,name=constraint,proto3" json:"constraint,omitempty"`
	// Define a string field for the standard.
	Terminate Termination `protobuf:"varint,4,opt,name=terminate,proto3,enum=mission.v1.Termination" json:"terminate,omitempty"`
}

func (x *TravelInstruction) Reset() {
	*x = TravelInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mission_v1_mission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelInstruction) ProtoMessage() {}

func (x *TravelInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_mission_v1_mission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelInstruction.ProtoReflect.Descriptor instead.
func (*TravelInstruction) Descriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{1}
}

func (x *TravelInstruction) GetPlanner() Planner {
	if x != nil {
		return x.Planner
	}
	return Planner_PLANNER_EGO
}

func (x *TravelInstruction) GetWaypoint() []float64 {
	if x != nil {
		return x.Waypoint
	}
	return nil
}

func (x *TravelInstruction) GetConstraint() []float64 {
	if x != nil {
		return x.Constraint
	}
	return nil
}

func (x *TravelInstruction) GetTerminate() Termination {
	if x != nil {
		return x.Terminate
	}
	return Termination_TERMINATION_AUTO
}

// Define the message for the ACTION instruction.
type ActionInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define a string field for the action type.
	Action Action `protobuf:"varint,1,opt,name=action,proto3,enum=mission.v1.Action" json:"action,omitempty"`
	// This field is optionals only for release command. If it is not populated, the default value is 0.
	Package []int32 `protobuf:"varint,2,rep,packed,name=package,proto3" json:"package,omitempty"`
	// For takoff, land. Optional. If not populated, the default value is 0.
	Param float64 `protobuf:"fixed64,3,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *ActionInstruction) Reset() {
	*x = ActionInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mission_v1_mission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionInstruction) ProtoMessage() {}

func (x *ActionInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_mission_v1_mission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionInstruction.ProtoReflect.Descriptor instead.
func (*ActionInstruction) Descriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{2}
}

func (x *ActionInstruction) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_TAKEOFF
}

func (x *ActionInstruction) GetPackage() []int32 {
	if x != nil {
		return x.Package
	}
	return nil
}

func (x *ActionInstruction) GetParam() float64 {
	if x != nil {
		return x.Param
	}
	return 0
}

type SequenceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Sequence:
	//
	//	*SequenceItem_InitSequence
	//	*SequenceItem_ActionSequence
	//	*SequenceItem_TravelSequence
	Sequence isSequenceItem_Sequence `protobuf_oneof:"sequence"`
}

func (x *SequenceItem) Reset() {
	*x = SequenceItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mission_v1_mission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SequenceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SequenceItem) ProtoMessage() {}

func (x *SequenceItem) ProtoReflect() protoreflect.Message {
	mi := &file_mission_v1_mission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SequenceItem.ProtoReflect.Descriptor instead.
func (*SequenceItem) Descriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{3}
}

func (m *SequenceItem) GetSequence() isSequenceItem_Sequence {
	if m != nil {
		return m.Sequence
	}
	return nil
}

func (x *SequenceItem) GetInitSequence() *InitInstruction {
	if x, ok := x.GetSequence().(*SequenceItem_InitSequence); ok {
		return x.InitSequence
	}
	return nil
}

func (x *SequenceItem) GetActionSequence() *ActionInstruction {
	if x, ok := x.GetSequence().(*SequenceItem_ActionSequence); ok {
		return x.ActionSequence
	}
	return nil
}

func (x *SequenceItem) GetTravelSequence() *TravelInstruction {
	if x, ok := x.GetSequence().(*SequenceItem_TravelSequence); ok {
		return x.TravelSequence
	}
	return nil
}

type isSequenceItem_Sequence interface {
	isSequenceItem_Sequence()
}

type SequenceItem_InitSequence struct {
	InitSequence *InitInstruction `protobuf:"bytes,1,opt,name=init_sequence,json=initSequence,proto3,oneof"`
}

type SequenceItem_ActionSequence struct {
	ActionSequence *ActionInstruction `protobuf:"bytes,2,opt,name=action_sequence,json=actionSequence,proto3,oneof"`
}

type SequenceItem_TravelSequence struct {
	TravelSequence *TravelInstruction `protobuf:"bytes,3,opt,name=travel_sequence,json=travelSequence,proto3,oneof"`
}

func (*SequenceItem_InitSequence) isSequenceItem_Sequence() {}

func (*SequenceItem_ActionSequence) isSequenceItem_Sequence() {}

func (*SequenceItem_TravelSequence) isSequenceItem_Sequence() {}

type SendMissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define a unique ID for the mission.
	Id            string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SequenceItems []*SequenceItem `protobuf:"bytes,2,rep,name=sequence_items,json=sequenceItems,proto3" json:"sequence_items,omitempty"`
}

func (x *SendMissionRequest) Reset() {
	*x = SendMissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mission_v1_mission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMissionRequest) ProtoMessage() {}

func (x *SendMissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mission_v1_mission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMissionRequest.ProtoReflect.Descriptor instead.
func (*SendMissionRequest) Descriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{4}
}

func (x *SendMissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SendMissionRequest) GetSequenceItems() []*SequenceItem {
	if x != nil {
		return x.SequenceItems
	}
	return nil
}

// For later use. Sending single instruction only
type SingleInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequenceItem *SequenceItem `protobuf:"bytes,1,opt,name=sequence_item,json=sequenceItem,proto3" json:"sequence_item,omitempty"`
}

func (x *SingleInstruction) Reset() {
	*x = SingleInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mission_v1_mission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleInstruction) ProtoMessage() {}

func (x *SingleInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_mission_v1_mission_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleInstruction.ProtoReflect.Descriptor instead.
func (*SingleInstruction) Descriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{5}
}

func (x *SingleInstruction) GetSequenceItem() *SequenceItem {
	if x != nil {
		return x.SequenceItem
	}
	return nil
}

type SendMissionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define a boolean field to indicate success or failure.
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Define a string field for any error messages.
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *SendMissionResult) Reset() {
	*x = SendMissionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mission_v1_mission_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMissionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMissionResult) ProtoMessage() {}

func (x *SendMissionResult) ProtoReflect() protoreflect.Message {
	mi := &file_mission_v1_mission_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMissionResult.ProtoReflect.Descriptor instead.
func (*SendMissionResult) Descriptor() ([]byte, []int) {
	return file_mission_v1_mission_proto_rawDescGZIP(), []int{6}
}

func (x *SendMissionResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SendMissionResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendMissionResult) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_mission_v1_mission_proto protoreflect.FileDescriptor

var file_mission_v1_mission_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xa0, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x69, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x35, 0x0a, 0x09, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x11, 0x54, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2d, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01,
	0x52, 0x08, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x22, 0x6f, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x22, 0xf2, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x42, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x48, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x65, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3f, 0x0a,
	0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x0d, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x52,
	0x0a, 0x11, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0d, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x22, 0x6c, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2a, 0x38, 0x0a, 0x0b, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x10, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x55, 0x54, 0x4f, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x44, 0x10, 0x01, 0x2a, 0x93, 0x01, 0x0a, 0x06, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x41, 0x4b, 0x45, 0x4f, 0x46, 0x46, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x52, 0x4d, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x4c, 0x46, 0x43, 0x48, 0x45, 0x43, 0x4b,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4c,
	0x45, 0x41, 0x53, 0x45, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x52, 0x54, 0x4c, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4c, 0x41, 0x4e, 0x44, 0x10, 0x06,
	0x2a, 0x56, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x50,
	0x4c, 0x41, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x45, 0x47, 0x4f, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x50, 0x4c, 0x41, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x53, 0x54, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x50, 0x4c, 0x41, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x52,
	0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x53, 0x41,
	0x46, 0x45, 0x4c, 0x41, 0x4e, 0x44, 0x10, 0x04, 0x2a, 0x69, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f,
	0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x50, 0x58, 0x34, 0x5f, 0x56, 0x45, 0x4c, 0x4f, 0x5f, 0x46, 0x42,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52,
	0x5f, 0x41, 0x5f, 0x46, 0x42, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x54, 0x52,
	0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x41, 0x5f, 0x46, 0x57, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11,
	0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f, 0x41, 0x5f, 0x41, 0x44, 0x52,
	0x4a, 0x10, 0x03, 0x32, 0x60, 0x0a, 0x0e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0xa8, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x74, 0x6d, 0x75, 0x72, 0x61, 0x69, 0x2f, 0x64,
	0x72, 0x6f, 0x6e, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa,
	0x02, 0x0a, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x4d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0b, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mission_v1_mission_proto_rawDescOnce sync.Once
	file_mission_v1_mission_proto_rawDescData = file_mission_v1_mission_proto_rawDesc
)

func file_mission_v1_mission_proto_rawDescGZIP() []byte {
	file_mission_v1_mission_proto_rawDescOnce.Do(func() {
		file_mission_v1_mission_proto_rawDescData = protoimpl.X.CompressGZIP(file_mission_v1_mission_proto_rawDescData)
	})
	return file_mission_v1_mission_proto_rawDescData
}

var file_mission_v1_mission_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_mission_v1_mission_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mission_v1_mission_proto_goTypes = []interface{}{
	(Termination)(0),           // 0: mission.v1.Termination
	(Action)(0),                // 1: mission.v1.Action
	(Planner)(0),               // 2: mission.v1.Planner
	(Controller)(0),            // 3: mission.v1.Controller
	(*InitInstruction)(nil),    // 4: mission.v1.InitInstruction
	(*TravelInstruction)(nil),  // 5: mission.v1.TravelInstruction
	(*ActionInstruction)(nil),  // 6: mission.v1.ActionInstruction
	(*SequenceItem)(nil),       // 7: mission.v1.SequenceItem
	(*SendMissionRequest)(nil), // 8: mission.v1.SendMissionRequest
	(*SingleInstruction)(nil),  // 9: mission.v1.SingleInstruction
	(*SendMissionResult)(nil),  // 10: mission.v1.SendMissionResult
}
var file_mission_v1_mission_proto_depIdxs = []int32{
	3,  // 0: mission.v1.InitInstruction.controller:type_name -> mission.v1.Controller
	0,  // 1: mission.v1.InitInstruction.terminate:type_name -> mission.v1.Termination
	2,  // 2: mission.v1.TravelInstruction.planner:type_name -> mission.v1.Planner
	0,  // 3: mission.v1.TravelInstruction.terminate:type_name -> mission.v1.Termination
	1,  // 4: mission.v1.ActionInstruction.action:type_name -> mission.v1.Action
	4,  // 5: mission.v1.SequenceItem.init_sequence:type_name -> mission.v1.InitInstruction
	6,  // 6: mission.v1.SequenceItem.action_sequence:type_name -> mission.v1.ActionInstruction
	5,  // 7: mission.v1.SequenceItem.travel_sequence:type_name -> mission.v1.TravelInstruction
	7,  // 8: mission.v1.SendMissionRequest.sequence_items:type_name -> mission.v1.SequenceItem
	7,  // 9: mission.v1.SingleInstruction.sequence_item:type_name -> mission.v1.SequenceItem
	8,  // 10: mission.v1.MissionService.SendMission:input_type -> mission.v1.SendMissionRequest
	10, // 11: mission.v1.MissionService.SendMission:output_type -> mission.v1.SendMissionResult
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_mission_v1_mission_proto_init() }
func file_mission_v1_mission_proto_init() {
	if File_mission_v1_mission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mission_v1_mission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitInstruction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mission_v1_mission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelInstruction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mission_v1_mission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionInstruction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mission_v1_mission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SequenceItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mission_v1_mission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMissionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mission_v1_mission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleInstruction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mission_v1_mission_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMissionResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_mission_v1_mission_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SequenceItem_InitSequence)(nil),
		(*SequenceItem_ActionSequence)(nil),
		(*SequenceItem_TravelSequence)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mission_v1_mission_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mission_v1_mission_proto_goTypes,
		DependencyIndexes: file_mission_v1_mission_proto_depIdxs,
		EnumInfos:         file_mission_v1_mission_proto_enumTypes,
		MessageInfos:      file_mission_v1_mission_proto_msgTypes,
	}.Build()
	File_mission_v1_mission_proto = out.File
	file_mission_v1_mission_proto_rawDesc = nil
	file_mission_v1_mission_proto_goTypes = nil
	file_mission_v1_mission_proto_depIdxs = nil
}
