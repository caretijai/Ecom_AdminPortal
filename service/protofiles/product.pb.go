// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: product.proto

package protofiles

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CatgName     string `protobuf:"bytes,1,opt,name=catg_name,json=catgName,proto3" json:"catg_name,omitempty"`
	CatgDesc     string `protobuf:"bytes,2,opt,name=catg_desc,json=catgDesc,proto3" json:"catg_desc,omitempty"`
	CatgImage    string `protobuf:"bytes,3,opt,name=catg_image,json=catgImage,proto3" json:"catg_image,omitempty"`
	SubCatgName  string `protobuf:"bytes,4,opt,name=sub_catg_name,json=subCatgName,proto3" json:"sub_catg_name,omitempty"`
	SubCatgDesc  string `protobuf:"bytes,5,opt,name=sub_catg_desc,json=subCatgDesc,proto3" json:"sub_catg_desc,omitempty"`
	SubCatgImage string `protobuf:"bytes,6,opt,name=sub_catg_image,json=subCatgImage,proto3" json:"sub_catg_image,omitempty"`
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCategoryRequest) GetCatgName() string {
	if x != nil {
		return x.CatgName
	}
	return ""
}

func (x *CreateCategoryRequest) GetCatgDesc() string {
	if x != nil {
		return x.CatgDesc
	}
	return ""
}

func (x *CreateCategoryRequest) GetCatgImage() string {
	if x != nil {
		return x.CatgImage
	}
	return ""
}

func (x *CreateCategoryRequest) GetSubCatgName() string {
	if x != nil {
		return x.SubCatgName
	}
	return ""
}

func (x *CreateCategoryRequest) GetSubCatgDesc() string {
	if x != nil {
		return x.SubCatgDesc
	}
	return ""
}

func (x *CreateCategoryRequest) GetSubCatgImage() string {
	if x != nil {
		return x.SubCatgImage
	}
	return ""
}

type CreateCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateCategoryResponse) Reset() {
	*x = CreateCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryResponse) ProtoMessage() {}

func (x *CreateCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCategoryResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type GetAllCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAllCategoriesRequest) Reset() {
	*x = GetAllCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCategoriesRequest) ProtoMessage() {}

func (x *GetAllCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCategoriesRequest.ProtoReflect.Descriptor instead.
func (*GetAllCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllCategoriesRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAllCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryData []*FinalSubCatg `protobuf:"bytes,1,rep,name=category_data,json=categoryData,proto3" json:"category_data,omitempty"`
}

func (x *GetAllCategoriesResponse) Reset() {
	*x = GetAllCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCategoriesResponse) ProtoMessage() {}

func (x *GetAllCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCategoriesResponse.ProtoReflect.Descriptor instead.
func (*GetAllCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllCategoriesResponse) GetCategoryData() []*FinalSubCatg {
	if x != nil {
		return x.CategoryData
	}
	return nil
}

type FinalSubCatg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryData    *Category      `protobuf:"bytes,1,opt,name=categoryData,proto3" json:"categoryData,omitempty"`
	SubCategoryData []*SubCategory `protobuf:"bytes,2,rep,name=subCategoryData,proto3" json:"subCategoryData,omitempty"`
}

func (x *FinalSubCatg) Reset() {
	*x = FinalSubCatg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalSubCatg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalSubCatg) ProtoMessage() {}

func (x *FinalSubCatg) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalSubCatg.ProtoReflect.Descriptor instead.
func (*FinalSubCatg) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *FinalSubCatg) GetCategoryData() *Category {
	if x != nil {
		return x.CategoryData
	}
	return nil
}

func (x *FinalSubCatg) GetSubCategoryData() []*SubCategory {
	if x != nil {
		return x.SubCategoryData
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CatgName  string `protobuf:"bytes,2,opt,name=catg_name,json=catgName,proto3" json:"catg_name,omitempty"`
	CatgDesc  string `protobuf:"bytes,3,opt,name=catg_desc,json=catgDesc,proto3" json:"catg_desc,omitempty"`
	CatgImage string `protobuf:"bytes,4,opt,name=catg_image,json=catgImage,proto3" json:"catg_image,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *Category) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetCatgName() string {
	if x != nil {
		return x.CatgName
	}
	return ""
}

func (x *Category) GetCatgDesc() string {
	if x != nil {
		return x.CatgDesc
	}
	return ""
}

func (x *Category) GetCatgImage() string {
	if x != nil {
		return x.CatgImage
	}
	return ""
}

type SubCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SubCatgName  string `protobuf:"bytes,2,opt,name=sub_catg_name,json=subCatgName,proto3" json:"sub_catg_name,omitempty"`
	SubCatgDesc  string `protobuf:"bytes,3,opt,name=sub_catg_desc,json=subCatgDesc,proto3" json:"sub_catg_desc,omitempty"`
	SubCatgImage string `protobuf:"bytes,4,opt,name=sub_catg_image,json=subCatgImage,proto3" json:"sub_catg_image,omitempty"`
}

func (x *SubCategory) Reset() {
	*x = SubCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCategory) ProtoMessage() {}

func (x *SubCategory) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCategory.ProtoReflect.Descriptor instead.
func (*SubCategory) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{6}
}

func (x *SubCategory) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SubCategory) GetSubCatgName() string {
	if x != nil {
		return x.SubCatgName
	}
	return ""
}

func (x *SubCategory) GetSubCatgDesc() string {
	if x != nil {
		return x.SubCatgDesc
	}
	return ""
}

func (x *SubCategory) GetSubCatgImage() string {
	if x != nil {
		return x.SubCatgImage
	}
	return ""
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x74, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x74, 0x67, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x67, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x61, 0x74, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x73, 0x75, 0x62, 0x5f, 0x63, 0x61, 0x74, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x61, 0x74, 0x67, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x67,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x61, 0x74, 0x67,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75,
	0x62, 0x43, 0x61, 0x74, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x29, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x75, 0x62,
	0x43, 0x61, 0x74, 0x67, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x8b, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x43,
	0x61, 0x74, 0x67, 0x12, 0x38, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a,
	0x0f, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x0f, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x73, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x61, 0x74, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x74,
	0x67, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x67, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x67, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x74, 0x67,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x61, 0x74,
	0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75,
	0x62, 0x43, 0x61, 0x74, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62,
	0x5f, 0x63, 0x61, 0x74, 0x67, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x67, 0x44, 0x65, 0x73, 0x63, 0x12, 0x24, 0x0a,
	0x0e, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x61, 0x74, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x67, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x32, 0x9b, 0x02, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x83, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x7a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_product_proto_goTypes = []interface{}{
	(*CreateCategoryRequest)(nil),    // 0: protofiles.CreateCategoryRequest
	(*CreateCategoryResponse)(nil),   // 1: protofiles.CreateCategoryResponse
	(*GetAllCategoriesRequest)(nil),  // 2: protofiles.GetAllCategoriesRequest
	(*GetAllCategoriesResponse)(nil), // 3: protofiles.GetAllCategoriesResponse
	(*FinalSubCatg)(nil),             // 4: protofiles.FinalSubCatg
	(*Category)(nil),                 // 5: protofiles.Category
	(*SubCategory)(nil),              // 6: protofiles.SubCategory
}
var file_product_proto_depIdxs = []int32{
	4, // 0: protofiles.GetAllCategoriesResponse.category_data:type_name -> protofiles.FinalSubCatg
	5, // 1: protofiles.FinalSubCatg.categoryData:type_name -> protofiles.Category
	6, // 2: protofiles.FinalSubCatg.subCategoryData:type_name -> protofiles.SubCategory
	2, // 3: protofiles.ProductCatalogueService.GetAllCategories:input_type -> protofiles.GetAllCategoriesRequest
	0, // 4: protofiles.ProductCatalogueService.CreateCategory:input_type -> protofiles.CreateCategoryRequest
	3, // 5: protofiles.ProductCatalogueService.GetAllCategories:output_type -> protofiles.GetAllCategoriesResponse
	1, // 6: protofiles.ProductCatalogueService.CreateCategory:output_type -> protofiles.CreateCategoryResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryRequest); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryResponse); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCategoriesRequest); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCategoriesResponse); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalSubCatg); i {
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
		file_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCategory); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
