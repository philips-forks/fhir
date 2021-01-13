//    Copyright 2020 Google Inc.
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        https://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: proto/google/fhir/proto/r5/core/resources/flag.proto

package flag_go_proto

import (
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/google/fhir/go/proto/google/fhir/proto/annotations_go_proto"
	codes_go_proto "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/codes_go_proto"
	datatypes_go_proto "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/datatypes_go_proto"
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

// Auto-generated from StructureDefinition for Flag, last updated
// 2019-12-31T21:03:40.621+11:00. Key information to flag to healthcare
// providers. See http://hl7.org/fhir/StructureDefinition/Flag
type Flag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Logical id of this artifact
	Id *datatypes_go_proto.Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Metadata about the resource
	Meta *datatypes_go_proto.Meta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *datatypes_go_proto.Uri `protobuf:"bytes,3,opt,name=implicit_rules,json=implicitRules,proto3" json:"implicit_rules,omitempty"`
	// Language of the resource content
	Language *datatypes_go_proto.Code `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *datatypes_go_proto.Narrative `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []*any.Any `protobuf:"bytes,6,rep,name=contained,proto3" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []*datatypes_go_proto.Extension `protobuf:"bytes,8,rep,name=extension,proto3" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []*datatypes_go_proto.Extension `protobuf:"bytes,9,rep,name=modifier_extension,json=modifierExtension,proto3" json:"modifier_extension,omitempty"`
	// Business identifier
	Identifier []*datatypes_go_proto.Identifier `protobuf:"bytes,10,rep,name=identifier,proto3" json:"identifier,omitempty"`
	Status     *Flag_StatusCode                 `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	// Clinical, administrative, etc.
	Category []*datatypes_go_proto.CodeableConcept `protobuf:"bytes,12,rep,name=category,proto3" json:"category,omitempty"`
	// Coded or textual message to display to user
	Code *datatypes_go_proto.CodeableConcept `protobuf:"bytes,13,opt,name=code,proto3" json:"code,omitempty"`
	// Who/What is flag about?
	Subject *datatypes_go_proto.Reference `protobuf:"bytes,14,opt,name=subject,proto3" json:"subject,omitempty"`
	// Time period when flag is active
	Period *datatypes_go_proto.Period `protobuf:"bytes,15,opt,name=period,proto3" json:"period,omitempty"`
	// Alert relevant during encounter
	Encounter *datatypes_go_proto.Reference `protobuf:"bytes,16,opt,name=encounter,proto3" json:"encounter,omitempty"`
	// Flag creator
	Author *datatypes_go_proto.Reference `protobuf:"bytes,17,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *Flag) Reset() {
	*x = Flag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r5_core_resources_flag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flag) ProtoMessage() {}

func (x *Flag) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r5_core_resources_flag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flag.ProtoReflect.Descriptor instead.
func (*Flag) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDescGZIP(), []int{0}
}

func (x *Flag) GetId() *datatypes_go_proto.Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Flag) GetMeta() *datatypes_go_proto.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Flag) GetImplicitRules() *datatypes_go_proto.Uri {
	if x != nil {
		return x.ImplicitRules
	}
	return nil
}

func (x *Flag) GetLanguage() *datatypes_go_proto.Code {
	if x != nil {
		return x.Language
	}
	return nil
}

func (x *Flag) GetText() *datatypes_go_proto.Narrative {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *Flag) GetContained() []*any.Any {
	if x != nil {
		return x.Contained
	}
	return nil
}

func (x *Flag) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

func (x *Flag) GetModifierExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.ModifierExtension
	}
	return nil
}

func (x *Flag) GetIdentifier() []*datatypes_go_proto.Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *Flag) GetStatus() *Flag_StatusCode {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Flag) GetCategory() []*datatypes_go_proto.CodeableConcept {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Flag) GetCode() *datatypes_go_proto.CodeableConcept {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *Flag) GetSubject() *datatypes_go_proto.Reference {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *Flag) GetPeriod() *datatypes_go_proto.Period {
	if x != nil {
		return x.Period
	}
	return nil
}

func (x *Flag) GetEncounter() *datatypes_go_proto.Reference {
	if x != nil {
		return x.Encounter
	}
	return nil
}

func (x *Flag) GetAuthor() *datatypes_go_proto.Reference {
	if x != nil {
		return x.Author
	}
	return nil
}

// active | inactive | entered-in-error
type Flag_StatusCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     codes_go_proto.FlagStatusCode_Value `protobuf:"varint,1,opt,name=value,proto3,enum=google.fhir.r5.core.FlagStatusCode_Value" json:"value,omitempty"`
	Id        *datatypes_go_proto.String          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Extension []*datatypes_go_proto.Extension     `protobuf:"bytes,3,rep,name=extension,proto3" json:"extension,omitempty"`
}

func (x *Flag_StatusCode) Reset() {
	*x = Flag_StatusCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r5_core_resources_flag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flag_StatusCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flag_StatusCode) ProtoMessage() {}

func (x *Flag_StatusCode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r5_core_resources_flag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flag_StatusCode.ProtoReflect.Descriptor instead.
func (*Flag_StatusCode) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Flag_StatusCode) GetValue() codes_go_proto.FlagStatusCode_Value {
	if x != nil {
		return x.Value
	}
	return codes_go_proto.FlagStatusCode_INVALID_UNINITIALIZED
}

func (x *Flag_StatusCode) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Flag_StatusCode) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

var File_proto_google_fhir_proto_r5_core_resources_flag_proto protoreflect.FileDescriptor

var file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDesc = []byte{
	0x0a, 0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66,
	0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x35, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x6c, 0x61, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66,
	0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x35, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x35, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa2, 0x0c, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x27, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68,
	0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x3f, 0x0a, 0x0e, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55,
	0x72, 0x69, 0x52, 0x0d, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69,
	0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x61, 0x72,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x32, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64,
	0x12, 0x3c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69,
	0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d,
	0x0a, 0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e,
	0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x44,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x42, 0x06, 0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x40, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68,
	0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x42, 0x06, 0xf0, 0xd0, 0x87, 0xeb,
	0x04, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0xd4, 0x01, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x99, 0x01, 0xf0, 0xd0, 0x87,
	0xeb, 0x04, 0x01, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xf2, 0xff,
	0xfc, 0xc2, 0x06, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x0c, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xf2, 0xff, 0xfc, 0xc2, 0x06,
	0x0c, 0x50, 0x72, 0x61, 0x63, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0xf2, 0xff, 0xfc,
	0xc2, 0x06, 0x10, 0x50, 0x72, 0x61, 0x63, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x0e, 0x50, 0x6c, 0x61, 0x6e, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x0a, 0x4d, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x09, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x33, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x4d, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x0f, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x09, 0x45,
	0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x8b, 0x01, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68,
	0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x42, 0x53, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x06, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x0c, 0x50, 0x72, 0x61, 0x63, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x65, 0x72, 0xf2, 0xff, 0xfc, 0xc2, 0x06, 0x10, 0x50, 0x72, 0x61, 0x63, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x1a, 0xa0, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c,
	0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e,
	0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x66, 0xc0, 0x9f,
	0xe3, 0xb6, 0x05, 0x01, 0x8a, 0xf9, 0x83, 0xb2, 0x05, 0x28, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f,
	0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x53, 0x65, 0x74, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x2d, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x9a, 0xb5, 0x8e, 0x93, 0x06, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x68,
	0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x64, 0x65, 0x3a, 0x38, 0xc0, 0x9f, 0xe3, 0xb6, 0x05, 0x03, 0xb2, 0xfe, 0xe4, 0x97,
	0x06, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x46, 0x6c, 0x61, 0x67, 0x4a, 0x04,
	0x08, 0x07, 0x10, 0x08, 0x42, 0x74, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x50,
	0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x35, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x98, 0xc6, 0xb0, 0xb5, 0x07, 0x05, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDescOnce sync.Once
	file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDescData = file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDesc
)

func file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDescGZIP() []byte {
	file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDescOnce.Do(func() {
		file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDescData)
	})
	return file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDescData
}

var file_proto_google_fhir_proto_r5_core_resources_flag_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_google_fhir_proto_r5_core_resources_flag_proto_goTypes = []interface{}{
	(*Flag)(nil),                               // 0: google.fhir.r5.core.Flag
	(*Flag_StatusCode)(nil),                    // 1: google.fhir.r5.core.Flag.StatusCode
	(*datatypes_go_proto.Id)(nil),              // 2: google.fhir.r5.core.Id
	(*datatypes_go_proto.Meta)(nil),            // 3: google.fhir.r5.core.Meta
	(*datatypes_go_proto.Uri)(nil),             // 4: google.fhir.r5.core.Uri
	(*datatypes_go_proto.Code)(nil),            // 5: google.fhir.r5.core.Code
	(*datatypes_go_proto.Narrative)(nil),       // 6: google.fhir.r5.core.Narrative
	(*any.Any)(nil),                            // 7: google.protobuf.Any
	(*datatypes_go_proto.Extension)(nil),       // 8: google.fhir.r5.core.Extension
	(*datatypes_go_proto.Identifier)(nil),      // 9: google.fhir.r5.core.Identifier
	(*datatypes_go_proto.CodeableConcept)(nil), // 10: google.fhir.r5.core.CodeableConcept
	(*datatypes_go_proto.Reference)(nil),       // 11: google.fhir.r5.core.Reference
	(*datatypes_go_proto.Period)(nil),          // 12: google.fhir.r5.core.Period
	(codes_go_proto.FlagStatusCode_Value)(0),   // 13: google.fhir.r5.core.FlagStatusCode.Value
	(*datatypes_go_proto.String)(nil),          // 14: google.fhir.r5.core.String
}
var file_proto_google_fhir_proto_r5_core_resources_flag_proto_depIdxs = []int32{
	2,  // 0: google.fhir.r5.core.Flag.id:type_name -> google.fhir.r5.core.Id
	3,  // 1: google.fhir.r5.core.Flag.meta:type_name -> google.fhir.r5.core.Meta
	4,  // 2: google.fhir.r5.core.Flag.implicit_rules:type_name -> google.fhir.r5.core.Uri
	5,  // 3: google.fhir.r5.core.Flag.language:type_name -> google.fhir.r5.core.Code
	6,  // 4: google.fhir.r5.core.Flag.text:type_name -> google.fhir.r5.core.Narrative
	7,  // 5: google.fhir.r5.core.Flag.contained:type_name -> google.protobuf.Any
	8,  // 6: google.fhir.r5.core.Flag.extension:type_name -> google.fhir.r5.core.Extension
	8,  // 7: google.fhir.r5.core.Flag.modifier_extension:type_name -> google.fhir.r5.core.Extension
	9,  // 8: google.fhir.r5.core.Flag.identifier:type_name -> google.fhir.r5.core.Identifier
	1,  // 9: google.fhir.r5.core.Flag.status:type_name -> google.fhir.r5.core.Flag.StatusCode
	10, // 10: google.fhir.r5.core.Flag.category:type_name -> google.fhir.r5.core.CodeableConcept
	10, // 11: google.fhir.r5.core.Flag.code:type_name -> google.fhir.r5.core.CodeableConcept
	11, // 12: google.fhir.r5.core.Flag.subject:type_name -> google.fhir.r5.core.Reference
	12, // 13: google.fhir.r5.core.Flag.period:type_name -> google.fhir.r5.core.Period
	11, // 14: google.fhir.r5.core.Flag.encounter:type_name -> google.fhir.r5.core.Reference
	11, // 15: google.fhir.r5.core.Flag.author:type_name -> google.fhir.r5.core.Reference
	13, // 16: google.fhir.r5.core.Flag.StatusCode.value:type_name -> google.fhir.r5.core.FlagStatusCode.Value
	14, // 17: google.fhir.r5.core.Flag.StatusCode.id:type_name -> google.fhir.r5.core.String
	8,  // 18: google.fhir.r5.core.Flag.StatusCode.extension:type_name -> google.fhir.r5.core.Extension
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_proto_google_fhir_proto_r5_core_resources_flag_proto_init() }
func file_proto_google_fhir_proto_r5_core_resources_flag_proto_init() {
	if File_proto_google_fhir_proto_r5_core_resources_flag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_google_fhir_proto_r5_core_resources_flag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flag); i {
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
		file_proto_google_fhir_proto_r5_core_resources_flag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flag_StatusCode); i {
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
			RawDescriptor: file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_google_fhir_proto_r5_core_resources_flag_proto_goTypes,
		DependencyIndexes: file_proto_google_fhir_proto_r5_core_resources_flag_proto_depIdxs,
		MessageInfos:      file_proto_google_fhir_proto_r5_core_resources_flag_proto_msgTypes,
	}.Build()
	File_proto_google_fhir_proto_r5_core_resources_flag_proto = out.File
	file_proto_google_fhir_proto_r5_core_resources_flag_proto_rawDesc = nil
	file_proto_google_fhir_proto_r5_core_resources_flag_proto_goTypes = nil
	file_proto_google_fhir_proto_r5_core_resources_flag_proto_depIdxs = nil
}