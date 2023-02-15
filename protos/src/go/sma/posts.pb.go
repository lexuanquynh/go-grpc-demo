// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: posts.proto

package sma

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PostStatus int32

const (
	PostStatus_POST_STATUS_UNSPECIFIED PostStatus = 0
	PostStatus_POST_STATUS_DRAFT       PostStatus = 1
	PostStatus_POST_STATUS_SCHEDULED   PostStatus = 2
	PostStatus_POST_STATUS_ARCHIVED    PostStatus = 3
	PostStatus_POST_STATUS_ACTIVE      PostStatus = 4
)

// Enum value maps for PostStatus.
var (
	PostStatus_name = map[int32]string{
		0: "POST_STATUS_UNSPECIFIED",
		1: "POST_STATUS_DRAFT",
		2: "POST_STATUS_SCHEDULED",
		3: "POST_STATUS_ARCHIVED",
		4: "POST_STATUS_ACTIVE",
	}
	PostStatus_value = map[string]int32{
		"POST_STATUS_UNSPECIFIED": 0,
		"POST_STATUS_DRAFT":       1,
		"POST_STATUS_SCHEDULED":   2,
		"POST_STATUS_ARCHIVED":    3,
		"POST_STATUS_ACTIVE":      4,
	}
)

func (x PostStatus) Enum() *PostStatus {
	p := new(PostStatus)
	*p = x
	return p
}

func (x PostStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PostStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_posts_proto_enumTypes[0].Descriptor()
}

func (PostStatus) Type() protoreflect.EnumType {
	return &file_posts_proto_enumTypes[0]
}

func (x PostStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PostStatus.Descriptor instead.
func (PostStatus) EnumDescriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{0}
}

type Filters int32

const (
	Filters_FILTER_UNESPECIFIED Filters = 0
	Filters_FILTER_DRAFT        Filters = 1
	Filters_FILTER_SCHEDULED    Filters = 2
	Filters_FILTER_ARCHIVED     Filters = 3
	Filters_FILTER_ACTIVE       Filters = 4
)

// Enum value maps for Filters.
var (
	Filters_name = map[int32]string{
		0: "FILTER_UNESPECIFIED",
		1: "FILTER_DRAFT",
		2: "FILTER_SCHEDULED",
		3: "FILTER_ARCHIVED",
		4: "FILTER_ACTIVE",
	}
	Filters_value = map[string]int32{
		"FILTER_UNESPECIFIED": 0,
		"FILTER_DRAFT":        1,
		"FILTER_SCHEDULED":    2,
		"FILTER_ARCHIVED":     3,
		"FILTER_ACTIVE":       4,
	}
)

func (x Filters) Enum() *Filters {
	p := new(Filters)
	*p = x
	return p
}

func (x Filters) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Filters) Descriptor() protoreflect.EnumDescriptor {
	return file_posts_proto_enumTypes[1].Descriptor()
}

func (Filters) Type() protoreflect.EnumType {
	return &file_posts_proto_enumTypes[1]
}

func (x Filters) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Filters.Descriptor instead.
func (Filters) EnumDescriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{1}
}

// CreatePostReq
//
// The request message to create a post
type CreatePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The title for the Post
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// The description for the Post
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The id of the user who created the Post
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The collection of the Posts's media files
	MediaIds []string `protobuf:"bytes,4,rep,name=media_ids,json=mediaIds,proto3" json:"media_ids,omitempty"`
	// The schedule date for the post
	ScheduledAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=scheduled_at,json=scheduledAt,proto3" json:"scheduled_at,omitempty"`
}

func (x *CreatePostReq) Reset() {
	*x = CreatePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostReq) ProtoMessage() {}

func (x *CreatePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostReq.ProtoReflect.Descriptor instead.
func (*CreatePostReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePostReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreatePostReq) GetMediaIds() []string {
	if x != nil {
		return x.MediaIds
	}
	return nil
}

func (x *CreatePostReq) GetScheduledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledAt
	}
	return nil
}

// UpdatePostReq
//
// The request message to update a post
type UpdatePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Field mask for a post
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The id for the post
	PostId string `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	// The updatable fields that are allowed to be updated
	Post *UpdatePost `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *UpdatePostReq) Reset() {
	*x = UpdatePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostReq) ProtoMessage() {}

func (x *UpdatePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostReq.ProtoReflect.Descriptor instead.
func (*UpdatePostReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePostReq) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdatePostReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *UpdatePostReq) GetPost() *UpdatePost {
	if x != nil {
		return x.Post
	}
	return nil
}

// UpdatePost
//
// The message that keeps track of the permitted attributes
// to be updated for a post
type UpdatePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The title for the post
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// The description for the post
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The collection of the Posts's media files
	MediaIds []string `protobuf:"bytes,3,rep,name=media_ids,json=mediaIds,proto3" json:"media_ids,omitempty"`
	// The schedule date for the post
	ScheduledAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=scheduled_at,json=scheduledAt,proto3" json:"scheduled_at,omitempty"`
}

func (x *UpdatePost) Reset() {
	*x = UpdatePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePost) ProtoMessage() {}

func (x *UpdatePost) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePost.ProtoReflect.Descriptor instead.
func (*UpdatePost) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdatePost) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdatePost) GetMediaIds() []string {
	if x != nil {
		return x.MediaIds
	}
	return nil
}

func (x *UpdatePost) GetScheduledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledAt
	}
	return nil
}

// PostIdReq
//
// The request message to show a post
type PostIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id for the post
	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *PostIdReq) Reset() {
	*x = PostIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostIdReq) ProtoMessage() {}

func (x *PostIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostIdReq.ProtoReflect.Descriptor instead.
func (*PostIdReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{3}
}

func (x *PostIdReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

// ListPostsReq
//
// The request message to list the collection of posts for a given user
type ListPostsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id for the user
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Search query string
	S string `protobuf:"bytes,2,opt,name=s,proto3" json:"s,omitempty"`
	// Requested page
	Page string `protobuf:"bytes,3,opt,name=page,proto3" json:"page,omitempty"`
	// Requested number of items per page
	// Default: 20
	// Max: 50
	PerPage string `protobuf:"bytes,4,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	// Filtered collection by status
	Filter Filters `protobuf:"varint,5,opt,name=filter,proto3,enum=sma.Filters" json:"filter,omitempty"`
}

func (x *ListPostsReq) Reset() {
	*x = ListPostsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsReq) ProtoMessage() {}

func (x *ListPostsReq) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsReq.ProtoReflect.Descriptor instead.
func (*ListPostsReq) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{4}
}

func (x *ListPostsReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListPostsReq) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *ListPostsReq) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *ListPostsReq) GetPerPage() string {
	if x != nil {
		return x.PerPage
	}
	return ""
}

func (x *ListPostsReq) GetFilter() Filters {
	if x != nil {
		return x.Filter
	}
	return Filters_FILTER_UNESPECIFIED
}

// ListPostsResp
//
// The response message to list the collection of posts for a given user
type ListPostsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection of posts
	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	// Page info
	PageInfo *PageInfo `protobuf:"bytes,2,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
}

func (x *ListPostsResp) Reset() {
	*x = ListPostsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsResp) ProtoMessage() {}

func (x *ListPostsResp) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsResp.ProtoReflect.Descriptor instead.
func (*ListPostsResp) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{5}
}

func (x *ListPostsResp) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *ListPostsResp) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current page
	Page uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	// The page size
	PageSize uint64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The total items
	TotalItems uint64 `protobuf:"varint,3,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	// The total pages
	TotalPages uint64 `protobuf:"varint,4,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{6}
}

func (x *PageInfo) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageInfo) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageInfo) GetTotalItems() uint64 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *PageInfo) GetTotalPages() uint64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

// Post
//
// The message that represents a post
type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id for the Post
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The title for the Post
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// The description for the Post
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The is of the user who created the Post
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The current status of the Post
	Status PostStatus `protobuf:"varint,5,opt,name=status,proto3,enum=sma.PostStatus" json:"status,omitempty"`
	// The date when the Post was created
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The date when the Post was updated
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{7}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Post) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Post) GetStatus() PostStatus {
	if x != nil {
		return x.Status
	}
	return PostStatus_POST_STATUS_UNSPECIFIED
}

func (x *Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Post) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// EmptyResp
//
// An empty response
type EmptyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResp) Reset() {
	*x = EmptyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResp) ProtoMessage() {}

func (x *EmptyResp) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResp.ProtoReflect.Descriptor instead.
func (*EmptyResp) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{8}
}

var File_posts_proto protoreflect.FileDescriptor

var file_posts_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73,
	0x6d, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6d, 0x61,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73,
	0x74, 0x22, 0xa0, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x24, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x24, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x73, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x73, 0x6d, 0x61, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73,
	0x6d, 0x61, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x7d, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x86, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x73, 0x6d, 0x61, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x0b, 0x0a,
	0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x2a, 0x8d, 0x01, 0x0a, 0x0a, 0x50,
	0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4f, 0x53,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a,
	0x15, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x43, 0x48,
	0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4f, 0x53, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x2a, 0x72, 0x0a, 0x07, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f,
	0x55, 0x4e, 0x45, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44,
	0x55, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52,
	0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x44, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x32, 0xea,
	0x01, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x73, 0x6d, 0x61, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x73, 0x6d, 0x61,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x73, 0x6d, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x73, 0x6d, 0x61, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e,
	0x2e, 0x73, 0x6d, 0x61, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x09,
	0x2e, 0x73, 0x6d, 0x61, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x73, 0x6d, 0x61, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x73, 0x6d, 0x61, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x6d,
	0x61, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x73, 0x6d,
	0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x73,
	0x72, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_posts_proto_rawDescOnce sync.Once
	file_posts_proto_rawDescData = file_posts_proto_rawDesc
)

func file_posts_proto_rawDescGZIP() []byte {
	file_posts_proto_rawDescOnce.Do(func() {
		file_posts_proto_rawDescData = protoimpl.X.CompressGZIP(file_posts_proto_rawDescData)
	})
	return file_posts_proto_rawDescData
}

var file_posts_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_posts_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_posts_proto_goTypes = []interface{}{
	(PostStatus)(0),               // 0: sma.PostStatus
	(Filters)(0),                  // 1: sma.Filters
	(*CreatePostReq)(nil),         // 2: sma.CreatePostReq
	(*UpdatePostReq)(nil),         // 3: sma.UpdatePostReq
	(*UpdatePost)(nil),            // 4: sma.UpdatePost
	(*PostIdReq)(nil),             // 5: sma.PostIdReq
	(*ListPostsReq)(nil),          // 6: sma.ListPostsReq
	(*ListPostsResp)(nil),         // 7: sma.ListPostsResp
	(*PageInfo)(nil),              // 8: sma.PageInfo
	(*Post)(nil),                  // 9: sma.Post
	(*EmptyResp)(nil),             // 10: sma.EmptyResp
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 12: google.protobuf.FieldMask
}
var file_posts_proto_depIdxs = []int32{
	11, // 0: sma.CreatePostReq.scheduled_at:type_name -> google.protobuf.Timestamp
	12, // 1: sma.UpdatePostReq.update_mask:type_name -> google.protobuf.FieldMask
	4,  // 2: sma.UpdatePostReq.post:type_name -> sma.UpdatePost
	11, // 3: sma.UpdatePost.scheduled_at:type_name -> google.protobuf.Timestamp
	1,  // 4: sma.ListPostsReq.filter:type_name -> sma.Filters
	9,  // 5: sma.ListPostsResp.posts:type_name -> sma.Post
	8,  // 6: sma.ListPostsResp.page_info:type_name -> sma.PageInfo
	0,  // 7: sma.Post.status:type_name -> sma.PostStatus
	11, // 8: sma.Post.created_at:type_name -> google.protobuf.Timestamp
	11, // 9: sma.Post.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 10: sma.Posts.CreatePost:input_type -> sma.CreatePostReq
	3,  // 11: sma.Posts.UpdatePost:input_type -> sma.UpdatePostReq
	5,  // 12: sma.Posts.ShowPost:input_type -> sma.PostIdReq
	6,  // 13: sma.Posts.ListPosts:input_type -> sma.ListPostsReq
	5,  // 14: sma.Posts.DeletePost:input_type -> sma.PostIdReq
	9,  // 15: sma.Posts.CreatePost:output_type -> sma.Post
	9,  // 16: sma.Posts.UpdatePost:output_type -> sma.Post
	9,  // 17: sma.Posts.ShowPost:output_type -> sma.Post
	7,  // 18: sma.Posts.ListPosts:output_type -> sma.ListPostsResp
	10, // 19: sma.Posts.DeletePost:output_type -> sma.EmptyResp
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_posts_proto_init() }
func file_posts_proto_init() {
	if File_posts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_posts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostReq); i {
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
		file_posts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostReq); i {
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
		file_posts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePost); i {
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
		file_posts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostIdReq); i {
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
		file_posts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostsReq); i {
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
		file_posts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostsResp); i {
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
		file_posts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
		file_posts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_posts_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResp); i {
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
			RawDescriptor: file_posts_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_posts_proto_goTypes,
		DependencyIndexes: file_posts_proto_depIdxs,
		EnumInfos:         file_posts_proto_enumTypes,
		MessageInfos:      file_posts_proto_msgTypes,
	}.Build()
	File_posts_proto = out.File
	file_posts_proto_rawDesc = nil
	file_posts_proto_goTypes = nil
	file_posts_proto_depIdxs = nil
}
