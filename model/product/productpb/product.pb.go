// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: product/productpb/product.proto

package productpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	skupb "github.com/xidongc/mongo_ebenchmark/model/sku/skupb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Category int32

const (
	Category_Electronics Category = 0
	Category_Food        Category = 1
	Category_Sports      Category = 2
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "Electronics",
		1: "Food",
		2: "Sports",
	}
	Category_value = map[string]int32{
		"Electronics": 0,
		"Food":        1,
		"Sports":      2,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_product_productpb_product_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_product_productpb_product_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_product_productpb_product_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_productpb_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_product_productpb_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_product_productpb_product_proto_rawDescGZIP(), []int{0}
}

type NewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Active      bool              `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	Attributes  []string          `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Description string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Images      []string          `protobuf:"bytes,5,rep,name=images,proto3" json:"images,omitempty"`
	Metadata    map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Shippable   bool              `protobuf:"varint,7,opt,name=shippable,proto3" json:"shippable,omitempty"`
	Url         string            `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *NewRequest) Reset() {
	*x = NewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_productpb_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRequest) ProtoMessage() {}

func (x *NewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_productpb_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRequest.ProtoReflect.Descriptor instead.
func (*NewRequest) Descriptor() ([]byte, []int) {
	return file_product_productpb_product_proto_rawDescGZIP(), []int{1}
}

func (x *NewRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewRequest) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *NewRequest) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *NewRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewRequest) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *NewRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *NewRequest) GetShippable() bool {
	if x != nil {
		return x.Shippable
	}
	return false
}

func (x *NewRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_productpb_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_productpb_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_product_productpb_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_productpb_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_productpb_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_product_productpb_product_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Active      bool              `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Attributes  []string          `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Description string            `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Images      []string          `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty"`
	Metadata    map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Shippable   bool              `protobuf:"varint,8,opt,name=shippable,proto3" json:"shippable,omitempty"`
	Url         string            `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_productpb_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_productpb_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_product_productpb_product_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *UpdateRequest) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *UpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateRequest) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *UpdateRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateRequest) GetShippable() bool {
	if x != nil {
		return x.Shippable
	}
	return false
}

func (x *UpdateRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Active      bool              `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Attributes  []string          `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Description string            `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Images      []string          `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty"`
	Metadata    map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Shippable   bool              `protobuf:"varint,8,opt,name=shippable,proto3" json:"shippable,omitempty"`
	Url         string            `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`
	Skus        []*skupb.Sku      `protobuf:"bytes,10,rep,name=skus,proto3" json:"skus,omitempty"`
	Type        Category          `protobuf:"varint,11,opt,name=type,proto3,enum=productpb.Category" json:"type,omitempty"`
	Created     int64             `protobuf:"varint,998,opt,name=created,proto3" json:"created,omitempty"`
	Updated     int64             `protobuf:"varint,999,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_productpb_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_productpb_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_productpb_product_proto_rawDescGZIP(), []int{5}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Product) GetAttributes() []string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Product) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Product) GetShippable() bool {
	if x != nil {
		return x.Shippable
	}
	return false
}

func (x *Product) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Product) GetSkus() []*skupb.Sku {
	if x != nil {
		return x.Skus
	}
	return nil
}

func (x *Product) GetType() Category {
	if x != nil {
		return x.Type
	}
	return Category_Electronics
}

func (x *Product) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Product) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_product_productpb_product_proto protoreflect.FileDescriptor

var file_product_productpb_product_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x1a, 0x13, 0x73, 0x6b,
	0x75, 0x2f, 0x73, 0x6b, 0x75, 0x70, 0x62, 0x2f, 0x73, 0x6b, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xc0, 0x02, 0x0a, 0x0a, 0x4e, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3f, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1c, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd6, 0x02, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x42, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xc9, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x04, 0x73, 0x6b, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x73, 0x6b, 0x75, 0x70, 0x62, 0x2e, 0x53, 0x6b, 0x75, 0x52, 0x04, 0x73,
	0x6b, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0xe6, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0xe7, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a,
	0x31, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x6f, 0x6f, 0x64, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x10, 0x02, 0x32, 0xb6, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x03, 0x4e, 0x65, 0x77, 0x12, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22,
	0x08, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x45, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x13,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x08, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0x4b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x1a, 0x08, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x49, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x08,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x64, 0x6f, 0x6e, 0x67,
	0x63, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x5f, 0x65, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_productpb_product_proto_rawDescOnce sync.Once
	file_product_productpb_product_proto_rawDescData = file_product_productpb_product_proto_rawDesc
)

func file_product_productpb_product_proto_rawDescGZIP() []byte {
	file_product_productpb_product_proto_rawDescOnce.Do(func() {
		file_product_productpb_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_productpb_product_proto_rawDescData)
	})
	return file_product_productpb_product_proto_rawDescData
}

var file_product_productpb_product_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_product_productpb_product_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_product_productpb_product_proto_goTypes = []interface{}{
	(Category)(0),         // 0: productpb.Category
	(*Empty)(nil),         // 1: productpb.Empty
	(*NewRequest)(nil),    // 2: productpb.NewRequest
	(*GetRequest)(nil),    // 3: productpb.GetRequest
	(*DeleteRequest)(nil), // 4: productpb.DeleteRequest
	(*UpdateRequest)(nil), // 5: productpb.UpdateRequest
	(*Product)(nil),       // 6: productpb.Product
	nil,                   // 7: productpb.NewRequest.MetadataEntry
	nil,                   // 8: productpb.UpdateRequest.MetadataEntry
	nil,                   // 9: productpb.Product.MetadataEntry
	(*skupb.Sku)(nil),     // 10: skupb.Sku
}
var file_product_productpb_product_proto_depIdxs = []int32{
	7,  // 0: productpb.NewRequest.metadata:type_name -> productpb.NewRequest.MetadataEntry
	8,  // 1: productpb.UpdateRequest.metadata:type_name -> productpb.UpdateRequest.MetadataEntry
	9,  // 2: productpb.Product.metadata:type_name -> productpb.Product.MetadataEntry
	10, // 3: productpb.Product.skus:type_name -> skupb.Sku
	0,  // 4: productpb.Product.type:type_name -> productpb.Category
	2,  // 5: productpb.ProductService.New:input_type -> productpb.NewRequest
	3,  // 6: productpb.ProductService.Get:input_type -> productpb.GetRequest
	5,  // 7: productpb.ProductService.Update:input_type -> productpb.UpdateRequest
	4,  // 8: productpb.ProductService.Delete:input_type -> productpb.DeleteRequest
	6,  // 9: productpb.ProductService.New:output_type -> productpb.Product
	6,  // 10: productpb.ProductService.Get:output_type -> productpb.Product
	6,  // 11: productpb.ProductService.Update:output_type -> productpb.Product
	1,  // 12: productpb.ProductService.Delete:output_type -> productpb.Empty
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_product_productpb_product_proto_init() }
func file_product_productpb_product_proto_init() {
	if File_product_productpb_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_productpb_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_product_productpb_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRequest); i {
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
		file_product_productpb_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_product_productpb_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_product_productpb_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_product_productpb_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_product_productpb_product_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_productpb_product_proto_goTypes,
		DependencyIndexes: file_product_productpb_product_proto_depIdxs,
		EnumInfos:         file_product_productpb_product_proto_enumTypes,
		MessageInfos:      file_product_productpb_product_proto_msgTypes,
	}.Build()
	File_product_productpb_product_proto = out.File
	file_product_productpb_product_proto_rawDesc = nil
	file_product_productpb_product_proto_goTypes = nil
	file_product_productpb_product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*Product, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Product, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Product, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/productpb.ProductService/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/productpb.ProductService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/productpb.ProductService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/productpb.ProductService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	New(context.Context, *NewRequest) (*Product, error)
	Get(context.Context, *GetRequest) (*Product, error)
	Update(context.Context, *UpdateRequest) (*Product, error)
	Delete(context.Context, *DeleteRequest) (*Empty, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) New(context.Context, *NewRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (*UnimplementedProductServiceServer) Get(context.Context, *GetRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedProductServiceServer) Update(context.Context, *UpdateRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedProductServiceServer) Delete(context.Context, *DeleteRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ProductService/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).New(ctx, req.(*NewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ProductService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ProductService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productpb.ProductService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "productpb.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _ProductService_New_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProductService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProductService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/productpb/product.proto",
}
