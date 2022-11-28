// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: validatorstate/validator_state.proto

package validatorstate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetMinimumHeightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *GetMinimumHeightResponse) Reset() {
	*x = GetMinimumHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validatorstate_validator_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMinimumHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMinimumHeightResponse) ProtoMessage() {}

func (x *GetMinimumHeightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_validatorstate_validator_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMinimumHeightResponse.ProtoReflect.Descriptor instead.
func (*GetMinimumHeightResponse) Descriptor() ([]byte, []int) {
	return file_validatorstate_validator_state_proto_rawDescGZIP(), []int{0}
}

func (x *GetMinimumHeightResponse) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type GetCurrentHeightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *GetCurrentHeightResponse) Reset() {
	*x = GetCurrentHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validatorstate_validator_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentHeightResponse) ProtoMessage() {}

func (x *GetCurrentHeightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_validatorstate_validator_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentHeightResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentHeightResponse) Descriptor() ([]byte, []int) {
	return file_validatorstate_validator_state_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrentHeightResponse) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type GetValidatorSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height   uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	SubnetId []byte `protobuf:"bytes,2,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
}

func (x *GetValidatorSetRequest) Reset() {
	*x = GetValidatorSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validatorstate_validator_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValidatorSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValidatorSetRequest) ProtoMessage() {}

func (x *GetValidatorSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_validatorstate_validator_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValidatorSetRequest.ProtoReflect.Descriptor instead.
func (*GetValidatorSetRequest) Descriptor() ([]byte, []int) {
	return file_validatorstate_validator_state_proto_rawDescGZIP(), []int{2}
}

func (x *GetValidatorSetRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *GetValidatorSetRequest) GetSubnetId() []byte {
	if x != nil {
		return x.SubnetId
	}
	return nil
}

type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId []byte `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Weight uint64 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validatorstate_validator_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_validatorstate_validator_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_validatorstate_validator_state_proto_rawDescGZIP(), []int{3}
}

func (x *Validator) GetNodeId() []byte {
	if x != nil {
		return x.NodeId
	}
	return nil
}

func (x *Validator) GetWeight() uint64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type GetValidatorSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validators []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (x *GetValidatorSetResponse) Reset() {
	*x = GetValidatorSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validatorstate_validator_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValidatorSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValidatorSetResponse) ProtoMessage() {}

func (x *GetValidatorSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_validatorstate_validator_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValidatorSetResponse.ProtoReflect.Descriptor instead.
func (*GetValidatorSetResponse) Descriptor() ([]byte, []int) {
	return file_validatorstate_validator_state_proto_rawDescGZIP(), []int{4}
}

func (x *GetValidatorSetResponse) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

var File_validatorstate_validator_state_proto protoreflect.FileDescriptor

var file_validatorstate_validator_state_proto_rawDesc = []byte{
	0x0a, 0x24, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75,
	0x6d, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x32, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x4d, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x09, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x54, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x32, 0xa0,
	0x02, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x54, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74,
	0x12, 0x26, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x76, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x68, 0x65, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_validatorstate_validator_state_proto_rawDescOnce sync.Once
	file_validatorstate_validator_state_proto_rawDescData = file_validatorstate_validator_state_proto_rawDesc
)

func file_validatorstate_validator_state_proto_rawDescGZIP() []byte {
	file_validatorstate_validator_state_proto_rawDescOnce.Do(func() {
		file_validatorstate_validator_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_validatorstate_validator_state_proto_rawDescData)
	})
	return file_validatorstate_validator_state_proto_rawDescData
}

var file_validatorstate_validator_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_validatorstate_validator_state_proto_goTypes = []interface{}{
	(*GetMinimumHeightResponse)(nil), // 0: validatorstate.GetMinimumHeightResponse
	(*GetCurrentHeightResponse)(nil), // 1: validatorstate.GetCurrentHeightResponse
	(*GetValidatorSetRequest)(nil),   // 2: validatorstate.GetValidatorSetRequest
	(*Validator)(nil),                // 3: validatorstate.Validator
	(*GetValidatorSetResponse)(nil),  // 4: validatorstate.GetValidatorSetResponse
	(*emptypb.Empty)(nil),            // 5: google.protobuf.Empty
}
var file_validatorstate_validator_state_proto_depIdxs = []int32{
	3, // 0: validatorstate.GetValidatorSetResponse.validators:type_name -> validatorstate.Validator
	5, // 1: validatorstate.ValidatorState.GetMinimumHeight:input_type -> google.protobuf.Empty
	5, // 2: validatorstate.ValidatorState.GetCurrentHeight:input_type -> google.protobuf.Empty
	2, // 3: validatorstate.ValidatorState.GetValidatorSet:input_type -> validatorstate.GetValidatorSetRequest
	0, // 4: validatorstate.ValidatorState.GetMinimumHeight:output_type -> validatorstate.GetMinimumHeightResponse
	1, // 5: validatorstate.ValidatorState.GetCurrentHeight:output_type -> validatorstate.GetCurrentHeightResponse
	4, // 6: validatorstate.ValidatorState.GetValidatorSet:output_type -> validatorstate.GetValidatorSetResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_validatorstate_validator_state_proto_init() }
func file_validatorstate_validator_state_proto_init() {
	if File_validatorstate_validator_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_validatorstate_validator_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMinimumHeightResponse); i {
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
		file_validatorstate_validator_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentHeightResponse); i {
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
		file_validatorstate_validator_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValidatorSetRequest); i {
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
		file_validatorstate_validator_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
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
		file_validatorstate_validator_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValidatorSetResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_validatorstate_validator_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_validatorstate_validator_state_proto_goTypes,
		DependencyIndexes: file_validatorstate_validator_state_proto_depIdxs,
		MessageInfos:      file_validatorstate_validator_state_proto_msgTypes,
	}.Build()
	File_validatorstate_validator_state_proto = out.File
	file_validatorstate_validator_state_proto_rawDesc = nil
	file_validatorstate_validator_state_proto_goTypes = nil
	file_validatorstate_validator_state_proto_depIdxs = nil
}
