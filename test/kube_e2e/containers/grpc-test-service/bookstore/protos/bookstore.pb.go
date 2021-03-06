// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/bookstore.proto

/*
Package bookstore is a generated protocol buffer package.

It is generated from these files:
	protos/bookstore.proto

It has these top-level messages:
	Shelf
	Book
	Author
	ListShelvesResponse
	CreateShelfRequest
	GetShelfRequest
	DeleteShelfRequest
	ListBooksRequest
	CreateBookRequest
	GetBookRequest
	UpdateBookRequest
	DeleteBookRequest
	GetAuthorRequest
*/
package bookstore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Author_Gender int32

const (
	Author_UNKNOWN Author_Gender = 0
	Author_MALE    Author_Gender = 1
	Author_FEMALE  Author_Gender = 2
)

var Author_Gender_name = map[int32]string{
	0: "UNKNOWN",
	1: "MALE",
	2: "FEMALE",
}
var Author_Gender_value = map[string]int32{
	"UNKNOWN": 0,
	"MALE":    1,
	"FEMALE":  2,
}

func (x Author_Gender) String() string {
	return proto.EnumName(Author_Gender_name, int32(x))
}
func (Author_Gender) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// A shelf resource.
type Shelf struct {
	// A unique shelf id.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// A theme of the shelf (fiction, poetry, etc).
	Theme string `protobuf:"bytes,2,opt,name=theme" json:"theme,omitempty"`
}

func (m *Shelf) Reset()                    { *m = Shelf{} }
func (m *Shelf) String() string            { return proto.CompactTextString(m) }
func (*Shelf) ProtoMessage()               {}
func (*Shelf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Shelf) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Shelf) GetTheme() string {
	if m != nil {
		return m.Theme
	}
	return ""
}

// A book resource.
type Book struct {
	// A unique book id.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// An author of the book.
	Author string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	// A book title.
	Title string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	// Quotes from the book.
	Quotes []string `protobuf:"bytes,4,rep,name=quotes" json:"quotes,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Book) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetQuotes() []string {
	if m != nil {
		return m.Quotes
	}
	return nil
}

// An author resource.
type Author struct {
	// A unique author id.
	Id        int64         `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Gender    Author_Gender `protobuf:"varint,2,opt,name=gender,enum=bookstore.Author_Gender" json:"gender,omitempty"`
	FirstName string        `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string        `protobuf:"bytes,4,opt,name=last_name,json=lname" json:"last_name,omitempty"`
}

func (m *Author) Reset()                    { *m = Author{} }
func (m *Author) String() string            { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()               {}
func (*Author) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Author) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Author) GetGender() Author_Gender {
	if m != nil {
		return m.Gender
	}
	return Author_UNKNOWN
}

func (m *Author) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Author) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

// Response to ListShelves call.
type ListShelvesResponse struct {
	// Shelves in the bookstore.
	Shelves []*Shelf `protobuf:"bytes,1,rep,name=shelves" json:"shelves,omitempty"`
}

func (m *ListShelvesResponse) Reset()                    { *m = ListShelvesResponse{} }
func (m *ListShelvesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListShelvesResponse) ProtoMessage()               {}
func (*ListShelvesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListShelvesResponse) GetShelves() []*Shelf {
	if m != nil {
		return m.Shelves
	}
	return nil
}

// Request message for CreateShelf method.
type CreateShelfRequest struct {
	// The shelf resource to create.
	Shelf *Shelf `protobuf:"bytes,1,opt,name=shelf" json:"shelf,omitempty"`
}

func (m *CreateShelfRequest) Reset()                    { *m = CreateShelfRequest{} }
func (m *CreateShelfRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateShelfRequest) ProtoMessage()               {}
func (*CreateShelfRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateShelfRequest) GetShelf() *Shelf {
	if m != nil {
		return m.Shelf
	}
	return nil
}

// Request message for GetShelf method.
type GetShelfRequest struct {
	// The ID of the shelf resource to retrieve.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf" json:"shelf,omitempty"`
}

func (m *GetShelfRequest) Reset()                    { *m = GetShelfRequest{} }
func (m *GetShelfRequest) String() string            { return proto.CompactTextString(m) }
func (*GetShelfRequest) ProtoMessage()               {}
func (*GetShelfRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetShelfRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

// Request message for DeleteShelf method.
type DeleteShelfRequest struct {
	// The ID of the shelf to delete.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf" json:"shelf,omitempty"`
}

func (m *DeleteShelfRequest) Reset()                    { *m = DeleteShelfRequest{} }
func (m *DeleteShelfRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteShelfRequest) ProtoMessage()               {}
func (*DeleteShelfRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteShelfRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

// Request message for ListBooks method.
type ListBooksRequest struct {
	// ID of the shelf which books to list.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf" json:"shelf,omitempty"`
}

func (m *ListBooksRequest) Reset()                    { *m = ListBooksRequest{} }
func (m *ListBooksRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBooksRequest) ProtoMessage()               {}
func (*ListBooksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListBooksRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

// Request message for CreateBook method.
type CreateBookRequest struct {
	// The ID of the shelf on which to create a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf" json:"shelf,omitempty"`
	// A book resource to create on the shelf.
	Book *Book `protobuf:"bytes,2,opt,name=book" json:"book,omitempty"`
}

func (m *CreateBookRequest) Reset()                    { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()               {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateBookRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

func (m *CreateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

// Request message for GetBook method.
type GetBookRequest struct {
	// The ID of the shelf from which to retrieve a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf" json:"shelf,omitempty"`
	// The ID of the book to retrieve.
	Book int64 `protobuf:"varint,2,opt,name=book" json:"book,omitempty"`
}

func (m *GetBookRequest) Reset()                    { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()               {}
func (*GetBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetBookRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

func (m *GetBookRequest) GetBook() int64 {
	if m != nil {
		return m.Book
	}
	return 0
}

// Request message for UpdateBook method
type UpdateBookRequest struct {
	// The ID of the shelf from which to retrieve a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf" json:"shelf,omitempty"`
	// A book resource to update on the shelf.
	Book *Book `protobuf:"bytes,2,opt,name=book" json:"book,omitempty"`
}

func (m *UpdateBookRequest) Reset()                    { *m = UpdateBookRequest{} }
func (m *UpdateBookRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateBookRequest) ProtoMessage()               {}
func (*UpdateBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateBookRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

func (m *UpdateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

// Request message for DeleteBook method.
type DeleteBookRequest struct {
	// The ID of the shelf from which to delete a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf" json:"shelf,omitempty"`
	// The ID of the book to delete.
	Book int64 `protobuf:"varint,2,opt,name=book" json:"book,omitempty"`
}

func (m *DeleteBookRequest) Reset()                    { *m = DeleteBookRequest{} }
func (m *DeleteBookRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteBookRequest) ProtoMessage()               {}
func (*DeleteBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteBookRequest) GetShelf() int64 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

func (m *DeleteBookRequest) GetBook() int64 {
	if m != nil {
		return m.Book
	}
	return 0
}

// Request message for GetAuthor method.
type GetAuthorRequest struct {
	// The ID of the author resource to retrieve.
	Author int64 `protobuf:"varint,1,opt,name=author" json:"author,omitempty"`
}

func (m *GetAuthorRequest) Reset()                    { *m = GetAuthorRequest{} }
func (m *GetAuthorRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAuthorRequest) ProtoMessage()               {}
func (*GetAuthorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetAuthorRequest) GetAuthor() int64 {
	if m != nil {
		return m.Author
	}
	return 0
}

func init() {
	proto.RegisterType((*Shelf)(nil), "bookstore.Shelf")
	proto.RegisterType((*Book)(nil), "bookstore.Book")
	proto.RegisterType((*Author)(nil), "bookstore.Author")
	proto.RegisterType((*ListShelvesResponse)(nil), "bookstore.ListShelvesResponse")
	proto.RegisterType((*CreateShelfRequest)(nil), "bookstore.CreateShelfRequest")
	proto.RegisterType((*GetShelfRequest)(nil), "bookstore.GetShelfRequest")
	proto.RegisterType((*DeleteShelfRequest)(nil), "bookstore.DeleteShelfRequest")
	proto.RegisterType((*ListBooksRequest)(nil), "bookstore.ListBooksRequest")
	proto.RegisterType((*CreateBookRequest)(nil), "bookstore.CreateBookRequest")
	proto.RegisterType((*GetBookRequest)(nil), "bookstore.GetBookRequest")
	proto.RegisterType((*UpdateBookRequest)(nil), "bookstore.UpdateBookRequest")
	proto.RegisterType((*DeleteBookRequest)(nil), "bookstore.DeleteBookRequest")
	proto.RegisterType((*GetAuthorRequest)(nil), "bookstore.GetAuthorRequest")
	proto.RegisterEnum("bookstore.Author_Gender", Author_Gender_name, Author_Gender_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Bookstore service

type BookstoreClient interface {
	// Returns a list of all shelves in the bookstore.
	ListShelves(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ListShelvesResponse, error)
	// Creates a new shelf in the bookstore.
	CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Creates multiple shelves with one streaming call
	BulkCreateShelf(ctx context.Context, opts ...grpc.CallOption) (Bookstore_BulkCreateShelfClient, error)
	// Returns a specific bookstore shelf.
	GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Deletes a shelf, including all books that are stored on the shelf.
	DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Returns a list of books on a shelf.
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (Bookstore_ListBooksClient, error)
	// Creates a new book.
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Returns a specific book.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Deletes a book from a shelf.
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
	BookstoreOptions(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Returns a specific author.
	GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error)
}

type bookstoreClient struct {
	cc *grpc.ClientConn
}

func NewBookstoreClient(cc *grpc.ClientConn) BookstoreClient {
	return &bookstoreClient{cc}
}

func (c *bookstoreClient) ListShelves(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ListShelvesResponse, error) {
	out := new(ListShelvesResponse)
	err := grpc.Invoke(ctx, "/bookstore.Bookstore/ListShelves", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := grpc.Invoke(ctx, "/bookstore.Bookstore/CreateShelf", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) BulkCreateShelf(ctx context.Context, opts ...grpc.CallOption) (Bookstore_BulkCreateShelfClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bookstore_serviceDesc.Streams[0], c.cc, "/bookstore.Bookstore/BulkCreateShelf", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookstoreBulkCreateShelfClient{stream}
	return x, nil
}

type Bookstore_BulkCreateShelfClient interface {
	Send(*CreateShelfRequest) error
	Recv() (*Shelf, error)
	grpc.ClientStream
}

type bookstoreBulkCreateShelfClient struct {
	grpc.ClientStream
}

func (x *bookstoreBulkCreateShelfClient) Send(m *CreateShelfRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bookstoreBulkCreateShelfClient) Recv() (*Shelf, error) {
	m := new(Shelf)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookstoreClient) GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := grpc.Invoke(ctx, "/bookstore.Bookstore/GetShelf", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/bookstore.Bookstore/DeleteShelf", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (Bookstore_ListBooksClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bookstore_serviceDesc.Streams[1], c.cc, "/bookstore.Bookstore/ListBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookstoreListBooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bookstore_ListBooksClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type bookstoreListBooksClient struct {
	grpc.ClientStream
}

func (x *bookstoreListBooksClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookstoreClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/bookstore.Bookstore/CreateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/bookstore.Bookstore/GetBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/bookstore.Bookstore/DeleteBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/bookstore.Bookstore/UpdateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) BookstoreOptions(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/bookstore.Bookstore/BookstoreOptions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := grpc.Invoke(ctx, "/bookstore.Bookstore/GetAuthor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bookstore service

type BookstoreServer interface {
	// Returns a list of all shelves in the bookstore.
	ListShelves(context.Context, *google_protobuf1.Empty) (*ListShelvesResponse, error)
	// Creates a new shelf in the bookstore.
	CreateShelf(context.Context, *CreateShelfRequest) (*Shelf, error)
	// Creates multiple shelves with one streaming call
	BulkCreateShelf(Bookstore_BulkCreateShelfServer) error
	// Returns a specific bookstore shelf.
	GetShelf(context.Context, *GetShelfRequest) (*Shelf, error)
	// Deletes a shelf, including all books that are stored on the shelf.
	DeleteShelf(context.Context, *DeleteShelfRequest) (*google_protobuf1.Empty, error)
	// Returns a list of books on a shelf.
	ListBooks(*ListBooksRequest, Bookstore_ListBooksServer) error
	// Creates a new book.
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	// Returns a specific book.
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	// Deletes a book from a shelf.
	DeleteBook(context.Context, *DeleteBookRequest) (*google_protobuf1.Empty, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	BookstoreOptions(context.Context, *GetShelfRequest) (*google_protobuf1.Empty, error)
	// Returns a specific author.
	GetAuthor(context.Context, *GetAuthorRequest) (*Author, error)
}

func RegisterBookstoreServer(s *grpc.Server, srv BookstoreServer) {
	s.RegisterService(&_Bookstore_serviceDesc, srv)
}

func _Bookstore_ListShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ListShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/ListShelves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ListShelves(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_CreateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/CreateShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateShelf(ctx, req.(*CreateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_BulkCreateShelf_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BookstoreServer).BulkCreateShelf(&bookstoreBulkCreateShelfServer{stream})
}

type Bookstore_BulkCreateShelfServer interface {
	Send(*Shelf) error
	Recv() (*CreateShelfRequest, error)
	grpc.ServerStream
}

type bookstoreBulkCreateShelfServer struct {
	grpc.ServerStream
}

func (x *bookstoreBulkCreateShelfServer) Send(m *Shelf) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bookstoreBulkCreateShelfServer) Recv() (*CreateShelfRequest, error) {
	m := new(CreateShelfRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Bookstore_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/GetShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetShelf(ctx, req.(*GetShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/DeleteShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteShelf(ctx, req.(*DeleteShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_ListBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBooksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookstoreServer).ListBooks(m, &bookstoreListBooksServer{stream})
}

type Bookstore_ListBooksServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type bookstoreListBooksServer struct {
	grpc.ServerStream
}

func (x *bookstoreListBooksServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

func _Bookstore_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_BookstoreOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).BookstoreOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/BookstoreOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).BookstoreOptions(ctx, req.(*GetShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetAuthor(ctx, req.(*GetAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bookstore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bookstore.Bookstore",
	HandlerType: (*BookstoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListShelves",
			Handler:    _Bookstore_ListShelves_Handler,
		},
		{
			MethodName: "CreateShelf",
			Handler:    _Bookstore_CreateShelf_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _Bookstore_GetShelf_Handler,
		},
		{
			MethodName: "DeleteShelf",
			Handler:    _Bookstore_DeleteShelf_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _Bookstore_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Bookstore_GetBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _Bookstore_DeleteBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _Bookstore_UpdateBook_Handler,
		},
		{
			MethodName: "BookstoreOptions",
			Handler:    _Bookstore_BookstoreOptions_Handler,
		},
		{
			MethodName: "GetAuthor",
			Handler:    _Bookstore_GetAuthor_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkCreateShelf",
			Handler:       _Bookstore_BulkCreateShelf_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ListBooks",
			Handler:       _Bookstore_ListBooks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/bookstore.proto",
}

func init() { proto.RegisterFile("protos/bookstore.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 776 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xda, 0x4a,
	0x10, 0x8e, 0x81, 0xf0, 0x33, 0xe8, 0x04, 0x33, 0xc9, 0x41, 0x8e, 0x13, 0x72, 0xd0, 0xe6, 0x9c,
	0x53, 0x44, 0x55, 0x3b, 0xa2, 0x17, 0x95, 0xa2, 0xf6, 0x22, 0xb4, 0x09, 0xaa, 0x9a, 0x42, 0x45,
	0x9a, 0xe6, 0xa2, 0x95, 0x2a, 0x23, 0x96, 0xc4, 0xc2, 0xb0, 0x04, 0x2f, 0x95, 0xaa, 0x88, 0x9b,
	0xbe, 0x42, 0x5f, 0xa4, 0xef, 0x92, 0x57, 0xe8, 0x83, 0x54, 0xde, 0xb5, 0xc1, 0x60, 0xa0, 0x52,
	0xdb, 0x2b, 0x98, 0xdd, 0x99, 0x6f, 0xbe, 0x9d, 0xfd, 0xbe, 0x35, 0x14, 0x86, 0x23, 0xc6, 0x99,
	0x6b, 0xb6, 0x19, 0xeb, 0xb9, 0x9c, 0x8d, 0xa8, 0x21, 0x16, 0x30, 0x33, 0x5d, 0xd0, 0xf7, 0xaf,
	0x19, 0xbb, 0x76, 0xa8, 0x69, 0x0d, 0x6d, 0xd3, 0x1a, 0x0c, 0x18, 0xb7, 0xb8, 0xcd, 0x06, 0xae,
	0x4c, 0xd4, 0xf7, 0xfc, 0x5d, 0x11, 0xb5, 0xc7, 0x5d, 0x93, 0xf6, 0x87, 0xfc, 0xb3, 0xdc, 0x24,
	0x8f, 0x60, 0xf3, 0xe2, 0x86, 0x3a, 0x5d, 0xdc, 0x82, 0x98, 0xdd, 0xd1, 0x94, 0x92, 0x52, 0x8e,
	0xb7, 0x62, 0x76, 0x07, 0x77, 0x60, 0x93, 0xdf, 0xd0, 0x3e, 0xd5, 0x62, 0x25, 0xa5, 0x9c, 0x69,
	0xc9, 0x80, 0x7c, 0x80, 0x44, 0x8d, 0xb1, 0x5e, 0x24, 0xbb, 0x00, 0x49, 0x6b, 0xcc, 0x6f, 0xd8,
	0xc8, 0x4f, 0xf7, 0x23, 0x81, 0x62, 0x73, 0x87, 0x6a, 0x71, 0x1f, 0xc5, 0x0b, 0xbc, 0xec, 0xdb,
	0x31, 0xe3, 0xd4, 0xd5, 0x12, 0xa5, 0xb8, 0x97, 0x2d, 0x23, 0xf2, 0x4d, 0x81, 0xe4, 0x89, 0x2c,
	0x5c, 0x6c, 0x70, 0x04, 0xc9, 0x6b, 0x3a, 0xe8, 0x50, 0xd9, 0x60, 0xab, 0xaa, 0x19, 0xb3, 0x79,
	0xc8, 0x12, 0xa3, 0x2e, 0xf6, 0x5b, 0x7e, 0x1e, 0x16, 0x01, 0xba, 0xf6, 0xc8, 0xe5, 0x1f, 0x07,
	0x56, 0x3f, 0xe8, 0x9f, 0x11, 0x2b, 0x0d, 0xab, 0x4f, 0x51, 0x83, 0x8c, 0x63, 0x05, 0xbb, 0x09,
	0xc9, 0xce, 0xf1, 0x02, 0xf2, 0x10, 0x92, 0x12, 0x0a, 0xb3, 0x90, 0xba, 0x6c, 0xbc, 0x6a, 0x34,
	0xaf, 0x1a, 0xea, 0x06, 0xa6, 0x21, 0xf1, 0xfa, 0xe4, 0xfc, 0x54, 0x55, 0x10, 0x20, 0x79, 0x76,
	0x2a, 0xfe, 0xc7, 0xc8, 0x09, 0x6c, 0x9f, 0xdb, 0x2e, 0xf7, 0x66, 0xf8, 0x89, 0xba, 0x2d, 0xea,
	0x0e, 0xd9, 0xc0, 0xa5, 0x58, 0x81, 0x94, 0x2b, 0x97, 0x34, 0xa5, 0x14, 0x2f, 0x67, 0xab, 0x6a,
	0x88, 0xaf, 0x18, 0x78, 0x2b, 0x48, 0x20, 0x4f, 0x01, 0x9f, 0x8f, 0xa8, 0xc5, 0xa9, 0x5c, 0xa7,
	0xb7, 0x63, 0xea, 0x72, 0xfc, 0x1f, 0x36, 0xbd, 0x84, 0xae, 0x98, 0xc1, 0xb2, 0x7a, 0xb9, 0x4d,
	0x1e, 0x40, 0xae, 0x4e, 0xf9, 0x5c, 0xe9, 0x4e, 0xb8, 0x34, 0x1e, 0x24, 0x56, 0x00, 0x5f, 0x50,
	0x87, 0x2e, 0xb4, 0x59, 0x9e, 0x5b, 0x06, 0xd5, 0x3b, 0x95, 0x77, 0xd5, 0xee, 0xfa, 0xcc, 0x06,
	0xe4, 0x25, 0x79, 0x2f, 0x77, 0x6d, 0x2a, 0x1e, 0x42, 0xc2, 0x3b, 0x83, 0xb8, 0xc0, 0x6c, 0x35,
	0x17, 0x3a, 0x90, 0xa8, 0x15, 0x9b, 0xe4, 0x18, 0xb6, 0xea, 0x94, 0xff, 0x1c, 0x0c, 0x43, 0x60,
	0x71, 0xbf, 0xb6, 0x01, 0xf9, 0xcb, 0x61, 0xe7, 0xcf, 0x71, 0x79, 0x06, 0x79, 0x39, 0xb1, 0x5f,
	0xa3, 0x53, 0x01, 0xb5, 0x4e, 0xb9, 0x14, 0x67, 0x50, 0x3d, 0xf3, 0x89, 0x2c, 0xf7, 0xa3, 0xea,
	0x7d, 0x1a, 0x32, 0xb5, 0x80, 0x03, 0x5e, 0x41, 0x36, 0x24, 0x2a, 0x2c, 0x18, 0xd2, 0xc1, 0x46,
	0xe0, 0x60, 0xe3, 0xd4, 0x73, 0xb0, 0x7e, 0x10, 0xa2, 0xbd, 0x44, 0x84, 0x44, 0xfd, 0x72, 0xff,
	0xfd, 0x6b, 0x0c, 0x30, 0x6d, 0xfa, 0x52, 0xc3, 0x4b, 0xc8, 0x86, 0xa4, 0x86, 0xc5, 0x10, 0x40,
	0x54, 0x82, 0x7a, 0x44, 0x73, 0xe4, 0x6f, 0x81, 0x98, 0x23, 0x49, 0x81, 0xd8, 0x3d, 0xf6, 0x4f,
	0xdf, 0x81, 0x5c, 0x6d, 0xec, 0xf4, 0x7e, 0x0b, 0x7a, 0x5f, 0x40, 0x17, 0xc8, 0x5f, 0x66, 0x7b,
	0xec, 0xf4, 0x02, 0xc6, 0x7e, 0x87, 0xb2, 0x72, 0xa4, 0x60, 0x0b, 0xd2, 0x81, 0xd2, 0x51, 0x0f,
	0xd5, 0x2f, 0xc8, 0x7f, 0x09, 0xb6, 0x26, 0xb0, 0x11, 0xd5, 0x00, 0xd6, 0xbc, 0x13, 0xb0, 0x13,
	0x3c, 0x83, 0x6c, 0xc8, 0x14, 0x73, 0xac, 0xa3, 0x66, 0xd1, 0x57, 0x5c, 0x04, 0xd9, 0xc0, 0xf7,
	0x90, 0x99, 0x1a, 0x06, 0xf7, 0x16, 0xee, 0x25, 0x6c, 0x23, 0x7d, 0x51, 0x6b, 0xe4, 0x40, 0x90,
	0xd3, 0xb0, 0xb0, 0x48, 0x4e, 0xbe, 0xf8, 0x47, 0x0a, 0x5a, 0x00, 0x33, 0x8f, 0xe1, 0x7e, 0x64,
	0xb2, 0x21, 0x79, 0x46, 0xe1, 0xff, 0x15, 0xf0, 0x07, 0xfa, 0x0a, 0xf8, 0x63, 0xa1, 0x55, 0x7c,
	0x02, 0x29, 0xdf, 0x76, 0xb8, 0x3b, 0x3f, 0xda, 0xb5, 0xe0, 0x1b, 0x68, 0x03, 0xcc, 0x3c, 0x32,
	0xc7, 0x2d, 0x62, 0x9d, 0x95, 0xe3, 0xfb, 0x4f, 0x50, 0xfc, 0xa7, 0x52, 0x5c, 0x4e, 0xd1, 0xbc,
	0xf3, 0x7e, 0x26, 0xd8, 0x03, 0x98, 0xd9, 0x7b, 0xae, 0x55, 0xc4, 0xf5, 0x51, 0xa6, 0x86, 0xe8,
	0x51, 0xae, 0x96, 0xd6, 0xf5, 0x30, 0xec, 0xce, 0xc4, 0x1f, 0x48, 0x0f, 0xd4, 0xa9, 0x1f, 0x9b,
	0x43, 0xf1, 0x39, 0x5d, 0x2b, 0xba, 0x55, 0x67, 0x3b, 0x14, 0x7d, 0x8b, 0xb5, 0x3d, 0x48, 0x35,
	0xdf, 0xbc, 0x7d, 0xd9, 0x6c, 0x5c, 0x2c, 0x51, 0xe1, 0x3b, 0xc8, 0x4c, 0x5f, 0x8a, 0x39, 0xf5,
	0x2c, 0xbe, 0x1f, 0x7a, 0x3e, 0xf2, 0xd9, 0x23, 0xbb, 0xa2, 0xc3, 0x36, 0xe6, 0x4d, 0xf9, 0x96,
	0xb8, 0xe6, 0x9d, 0xfc, 0x33, 0x69, 0x27, 0x05, 0x99, 0xc7, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x8e, 0x5b, 0x05, 0x90, 0x43, 0x08, 0x00, 0x00,
}
