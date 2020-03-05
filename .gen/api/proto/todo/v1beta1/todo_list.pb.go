// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo/v1beta1/todo_list.proto

package todov1beta1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateTodoRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRequest) Reset()         { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()    {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd164c9ed9f675e, []int{0}
}

func (m *CreateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRequest.Unmarshal(m, b)
}
func (m *CreateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRequest.Marshal(b, m, deterministic)
}
func (m *CreateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRequest.Merge(m, src)
}
func (m *CreateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRequest.Size(m)
}
func (m *CreateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRequest proto.InternalMessageInfo

func (m *CreateTodoRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CreateTodoResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoResponse) Reset()         { *m = CreateTodoResponse{} }
func (m *CreateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoResponse) ProtoMessage()    {}
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd164c9ed9f675e, []int{1}
}

func (m *CreateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoResponse.Unmarshal(m, b)
}
func (m *CreateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoResponse.Merge(m, src)
}
func (m *CreateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoResponse.Size(m)
}
func (m *CreateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoResponse proto.InternalMessageInfo

func (m *CreateTodoResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListTodosRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodosRequest) Reset()         { *m = ListTodosRequest{} }
func (m *ListTodosRequest) String() string { return proto.CompactTextString(m) }
func (*ListTodosRequest) ProtoMessage()    {}
func (*ListTodosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd164c9ed9f675e, []int{2}
}

func (m *ListTodosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodosRequest.Unmarshal(m, b)
}
func (m *ListTodosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodosRequest.Marshal(b, m, deterministic)
}
func (m *ListTodosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodosRequest.Merge(m, src)
}
func (m *ListTodosRequest) XXX_Size() int {
	return xxx_messageInfo_ListTodosRequest.Size(m)
}
func (m *ListTodosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodosRequest proto.InternalMessageInfo

type ListTodosResponse struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodosResponse) Reset()         { *m = ListTodosResponse{} }
func (m *ListTodosResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodosResponse) ProtoMessage()    {}
func (*ListTodosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd164c9ed9f675e, []int{3}
}

func (m *ListTodosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodosResponse.Unmarshal(m, b)
}
func (m *ListTodosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodosResponse.Marshal(b, m, deterministic)
}
func (m *ListTodosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodosResponse.Merge(m, src)
}
func (m *ListTodosResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodosResponse.Size(m)
}
func (m *ListTodosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodosResponse proto.InternalMessageInfo

func (m *ListTodosResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type MarkAsDoneRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkAsDoneRequest) Reset()         { *m = MarkAsDoneRequest{} }
func (m *MarkAsDoneRequest) String() string { return proto.CompactTextString(m) }
func (*MarkAsDoneRequest) ProtoMessage()    {}
func (*MarkAsDoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd164c9ed9f675e, []int{4}
}

func (m *MarkAsDoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkAsDoneRequest.Unmarshal(m, b)
}
func (m *MarkAsDoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkAsDoneRequest.Marshal(b, m, deterministic)
}
func (m *MarkAsDoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkAsDoneRequest.Merge(m, src)
}
func (m *MarkAsDoneRequest) XXX_Size() int {
	return xxx_messageInfo_MarkAsDoneRequest.Size(m)
}
func (m *MarkAsDoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkAsDoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MarkAsDoneRequest proto.InternalMessageInfo

func (m *MarkAsDoneRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MarkAsDoneResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkAsDoneResponse) Reset()         { *m = MarkAsDoneResponse{} }
func (m *MarkAsDoneResponse) String() string { return proto.CompactTextString(m) }
func (*MarkAsDoneResponse) ProtoMessage()    {}
func (*MarkAsDoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fd164c9ed9f675e, []int{5}
}

func (m *MarkAsDoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkAsDoneResponse.Unmarshal(m, b)
}
func (m *MarkAsDoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkAsDoneResponse.Marshal(b, m, deterministic)
}
func (m *MarkAsDoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkAsDoneResponse.Merge(m, src)
}
func (m *MarkAsDoneResponse) XXX_Size() int {
	return xxx_messageInfo_MarkAsDoneResponse.Size(m)
}
func (m *MarkAsDoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkAsDoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MarkAsDoneResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateTodoRequest)(nil), "todo.v1beta1.CreateTodoRequest")
	proto.RegisterType((*CreateTodoResponse)(nil), "todo.v1beta1.CreateTodoResponse")
	proto.RegisterType((*ListTodosRequest)(nil), "todo.v1beta1.ListTodosRequest")
	proto.RegisterType((*ListTodosResponse)(nil), "todo.v1beta1.ListTodosResponse")
	proto.RegisterType((*MarkAsDoneRequest)(nil), "todo.v1beta1.MarkAsDoneRequest")
	proto.RegisterType((*MarkAsDoneResponse)(nil), "todo.v1beta1.MarkAsDoneResponse")
}

func init() {
	proto.RegisterFile("todo/v1beta1/todo_list.proto", fileDescriptor_6fd164c9ed9f675e)
}

var fileDescriptor_6fd164c9ed9f675e = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x49, 0xfa, 0xbe, 0x62, 0xa7, 0x55, 0x9a, 0x41, 0xb0, 0x04, 0xb1, 0x61, 0x15, 0xcc,
	0x69, 0x4b, 0xea, 0xd9, 0x83, 0xd1, 0x63, 0xa5, 0xa5, 0x04, 0x29, 0x22, 0x48, 0x6a, 0xf6, 0x10,
	0xd4, 0x6e, 0xcd, 0xae, 0xe2, 0xe7, 0xf1, 0xe8, 0x47, 0xf1, 0x23, 0x79, 0x92, 0x49, 0xb6, 0x36,
	0x7f, 0xe8, 0x2d, 0x33, 0xf3, 0xcb, 0x3c, 0xcf, 0x33, 0x09, 0x1c, 0x69, 0x99, 0xc8, 0xe1, 0x7b,
	0xb0, 0x10, 0x3a, 0x0e, 0x86, 0x54, 0x3c, 0x3c, 0xa7, 0x4a, 0xf3, 0x55, 0x26, 0xb5, 0xc4, 0x2e,
	0x35, 0xb8, 0x99, 0xba, 0x87, 0x0d, 0xb6, 0xc0, 0xd8, 0x19, 0x38, 0x57, 0x99, 0x88, 0xb5, 0x88,
	0x64, 0x22, 0x67, 0xe2, 0xf5, 0x4d, 0x28, 0x8d, 0x08, 0xff, 0xb4, 0xf8, 0xd0, 0x7d, 0xcb, 0xb3,
	0xfc, 0xf6, 0x2c, 0x7f, 0x66, 0xa7, 0x80, 0x65, 0x50, 0xad, 0xe4, 0x52, 0x09, 0xdc, 0x07, 0x3b,
	0x4d, 0x0c, 0x67, 0xa7, 0x09, 0x43, 0xe8, 0x8d, 0x53, 0xa5, 0x89, 0x51, 0x66, 0x1b, 0xbb, 0x00,
	0xa7, 0xd4, 0x33, 0x2f, 0xfa, 0xf0, 0x9f, 0x5c, 0xa8, 0xbe, 0xe5, 0xb5, 0xfc, 0xce, 0x08, 0x79,
	0xd9, 0x2e, 0xcf, 0x35, 0x0a, 0x80, 0x9d, 0x80, 0x73, 0x13, 0x67, 0x4f, 0x97, 0xea, 0x5a, 0x2e,
	0xc5, 0xda, 0x61, 0x5d, 0xf7, 0x00, 0xb0, 0x0c, 0x15, 0x22, 0xa3, 0x1f, 0x0b, 0x76, 0x69, 0x15,
	0xc9, 0xe3, 0x04, 0x60, 0x13, 0x00, 0x07, 0x55, 0xc1, 0xc6, 0x0d, 0x5c, 0x6f, 0x3b, 0x60, 0x22,
	0x8c, 0xa1, 0xfd, 0x97, 0x0b, 0x8f, 0xab, 0x78, 0xfd, 0x08, 0xee, 0x60, 0xeb, 0xdc, 0x6c, 0x9b,
	0x00, 0x6c, 0x12, 0xd4, 0xed, 0x35, 0x0e, 0x50, 0xb7, 0xd7, 0x0c, 0x1f, 0x46, 0xd0, 0x7b, 0x94,
	0x2f, 0x15, 0x2c, 0xdc, 0x5b, 0x5f, 0x63, 0x4a, 0x1f, 0x7f, 0x6a, 0xdd, 0x75, 0x68, 0x6c, 0xa6,
	0x9f, 0x76, 0x2b, 0x9a, 0xcf, 0xbf, 0xec, 0x2e, 0x41, 0xfc, 0x36, 0x08, 0xa9, 0xf9, 0x5d, 0x94,
	0xf7, 0xa6, 0x5c, 0xec, 0xe4, 0xbf, 0xcd, 0xf9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x64,
	0x31, 0x31, 0x7d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoListClient is the client API for TodoList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoListClient interface {
	// CreateTodo adds a new todo to the todo list.
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	// ListTodos returns the list of todos.
	ListTodos(ctx context.Context, in *ListTodosRequest, opts ...grpc.CallOption) (*ListTodosResponse, error)
	// MarkAsDone marks a todo as done.
	MarkAsDone(ctx context.Context, in *MarkAsDoneRequest, opts ...grpc.CallOption) (*MarkAsDoneResponse, error)
}

type todoListClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoListClient(cc grpc.ClientConnInterface) TodoListClient {
	return &todoListClient{cc}
}

func (c *todoListClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.v1beta1.TodoList/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) ListTodos(ctx context.Context, in *ListTodosRequest, opts ...grpc.CallOption) (*ListTodosResponse, error) {
	out := new(ListTodosResponse)
	err := c.cc.Invoke(ctx, "/todo.v1beta1.TodoList/ListTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) MarkAsDone(ctx context.Context, in *MarkAsDoneRequest, opts ...grpc.CallOption) (*MarkAsDoneResponse, error) {
	out := new(MarkAsDoneResponse)
	err := c.cc.Invoke(ctx, "/todo.v1beta1.TodoList/MarkAsDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoListServer is the server API for TodoList service.
type TodoListServer interface {
	// CreateTodo adds a new todo to the todo list.
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	// ListTodos returns the list of todos.
	ListTodos(context.Context, *ListTodosRequest) (*ListTodosResponse, error)
	// MarkAsDone marks a todo as done.
	MarkAsDone(context.Context, *MarkAsDoneRequest) (*MarkAsDoneResponse, error)
}

// UnimplementedTodoListServer can be embedded to have forward compatible implementations.
type UnimplementedTodoListServer struct {
}

func (*UnimplementedTodoListServer) CreateTodo(ctx context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (*UnimplementedTodoListServer) ListTodos(ctx context.Context, req *ListTodosRequest) (*ListTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodos not implemented")
}
func (*UnimplementedTodoListServer) MarkAsDone(ctx context.Context, req *MarkAsDoneRequest) (*MarkAsDoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsDone not implemented")
}

func RegisterTodoListServer(s *grpc.Server, srv TodoListServer) {
	s.RegisterService(&_TodoList_serviceDesc, srv)
}

func _TodoList_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1beta1.TodoList/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_ListTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).ListTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1beta1.TodoList/ListTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).ListTodos(ctx, req.(*ListTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_MarkAsDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAsDoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).MarkAsDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1beta1.TodoList/MarkAsDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).MarkAsDone(ctx, req.(*MarkAsDoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.v1beta1.TodoList",
	HandlerType: (*TodoListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoList_CreateTodo_Handler,
		},
		{
			MethodName: "ListTodos",
			Handler:    _TodoList_ListTodos_Handler,
		},
		{
			MethodName: "MarkAsDone",
			Handler:    _TodoList_MarkAsDone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo/v1beta1/todo_list.proto",
}
