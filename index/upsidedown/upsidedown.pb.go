package upsidedown

import (
	fmt "fmt"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var (
	ErrInvalidLengthUpsidedown = fmt.Errorf("proto: negative length found during unmarshaling")
)

type BackIndexTermsEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         *uint32                `protobuf:"varint,1,req,name=field" json:"field,omitempty"`
	Terms         []string               `protobuf:"bytes,2,rep,name=terms" json:"terms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BackIndexTermsEntry) Reset() { _ = "STUB: not implemented"; return }

func (x *BackIndexTermsEntry) String() string { _ = "STUB: not implemented"; return "" }

func (*BackIndexTermsEntry) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *BackIndexTermsEntry) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*BackIndexTermsEntry) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *BackIndexTermsEntry) GetField() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *BackIndexTermsEntry) GetTerms() []string { _ = "STUB: not implemented"; return nil }

func (x *BackIndexTermsEntry) MarshalTo(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (x *BackIndexTermsEntry) Size() (n int) { _ = "STUB: not implemented"; return 0 }

type BackIndexStoreEntry struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Field          *uint32                `protobuf:"varint,1,req,name=field" json:"field,omitempty"`
	ArrayPositions []uint64               `protobuf:"varint,2,rep,name=arrayPositions" json:"arrayPositions,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BackIndexStoreEntry) Reset() { _ = "STUB: not implemented"; return }

func (x *BackIndexStoreEntry) String() string { _ = "STUB: not implemented"; return "" }

func (*BackIndexStoreEntry) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *BackIndexStoreEntry) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*BackIndexStoreEntry) Descriptor() ([]byte, []int) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x *BackIndexStoreEntry) GetField() uint32 { _ = "STUB: not implemented"; return 0 }

func (x *BackIndexStoreEntry) GetArrayPositions() []uint64 { _ = "STUB: not implemented"; return nil }

func (x *BackIndexStoreEntry) MarshalTo(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (x *BackIndexStoreEntry) Size() (n int) { _ = "STUB: not implemented"; return 0 }

type BackIndexRowValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TermsEntries  []*BackIndexTermsEntry `protobuf:"bytes,1,rep,name=termsEntries" json:"termsEntries,omitempty"`
	StoredEntries []*BackIndexStoreEntry `protobuf:"bytes,2,rep,name=storedEntries" json:"storedEntries,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BackIndexRowValue) Reset() { _ = "STUB: not implemented"; return }

func (x *BackIndexRowValue) String() string { _ = "STUB: not implemented"; return "" }

func (*BackIndexRowValue) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *BackIndexRowValue) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*BackIndexRowValue) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *BackIndexRowValue) GetTermsEntries() []*BackIndexTermsEntry {
	_ = "STUB: not implemented"
	return nil
}

func (x *BackIndexRowValue) GetStoredEntries() []*BackIndexStoreEntry {
	_ = "STUB: not implemented"
	return nil
}

func (x *BackIndexRowValue) MarshalTo(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (x *BackIndexRowValue) Size() (n int) { _ = "STUB: not implemented"; return 0 }

func skipUpsidedown(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func sovUpsidedown(x uint64) (n int) { _ = "STUB: not implemented"; return 0 }

func encodeVarintUpsidedown(data []byte, offset int, v uint64) int {
	_ = "STUB: not implemented"
	return 0
}

var File_index_upsidedown_upsidedown_proto protoreflect.FileDescriptor

const file_index_upsidedown_upsidedown_proto_rawDesc = "" +
	"\n" +
	"!index/upsidedown/upsidedown.proto\"A\n" +
	"\x13BackIndexTermsEntry\x12\x14\n" +
	"\x05field\x18\x01 \x02(\rR\x05field\x12\x14\n" +
	"\x05terms\x18\x02 \x03(\tR\x05terms\"S\n" +
	"\x13BackIndexStoreEntry\x12\x14\n" +
	"\x05field\x18\x01 \x02(\rR\x05field\x12&\n" +
	"\x0earrayPositions\x18\x02 \x03(\x04R\x0earrayPositions\"\x89\x01\n" +
	"\x11BackIndexRowValue\x128\n" +
	"\ftermsEntries\x18\x01 \x03(\v2\x14.BackIndexTermsEntryR\ftermsEntries\x12:\n" +
	"\rstoredEntries\x18\x02 \x03(\v2\x14.BackIndexStoreEntryR\rstoredEntries"

var (
	file_index_upsidedown_upsidedown_proto_rawDescOnce sync.Once
	file_index_upsidedown_upsidedown_proto_rawDescData []byte
)

func file_index_upsidedown_upsidedown_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_index_upsidedown_upsidedown_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_index_upsidedown_upsidedown_proto_goTypes = []any{
	(*BackIndexTermsEntry)(nil),
	(*BackIndexStoreEntry)(nil),
	(*BackIndexRowValue)(nil),
}
var file_index_upsidedown_upsidedown_proto_depIdxs = []int32{
	0,
	1,
	2,
	2,
	2,
	2,
	0,
}

func init()                                        { file_index_upsidedown_upsidedown_proto_init() }
func file_index_upsidedown_upsidedown_proto_init() { _ = "STUB: not implemented"; return }
