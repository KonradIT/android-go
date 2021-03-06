// Code generated by protoc-gen-go. DO NOT EDIT.
// source: update_metadata.proto

package chromeos_update_engine

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type InstallOperation_Type int32

const (
	InstallOperation_REPLACE    InstallOperation_Type = 0
	InstallOperation_REPLACE_BZ InstallOperation_Type = 1
	InstallOperation_MOVE       InstallOperation_Type = 2
	InstallOperation_BSDIFF     InstallOperation_Type = 3
	// On minor version 2 or newer, these operations are supported:
	InstallOperation_SOURCE_COPY   InstallOperation_Type = 4
	InstallOperation_SOURCE_BSDIFF InstallOperation_Type = 5
	// On minor version 3 or newer and on major version 2 or newer, these
	// operations are supported:
	InstallOperation_REPLACE_XZ InstallOperation_Type = 8
	// On minor version 4 or newer, these operations are supported:
	InstallOperation_ZERO          InstallOperation_Type = 6
	InstallOperation_DISCARD       InstallOperation_Type = 7
	InstallOperation_BROTLI_BSDIFF InstallOperation_Type = 10
	// On minor version 5 or newer, these operations are supported:
	InstallOperation_PUFFDIFF InstallOperation_Type = 9
)

var InstallOperation_Type_name = map[int32]string{
	0:  "REPLACE",
	1:  "REPLACE_BZ",
	2:  "MOVE",
	3:  "BSDIFF",
	4:  "SOURCE_COPY",
	5:  "SOURCE_BSDIFF",
	8:  "REPLACE_XZ",
	6:  "ZERO",
	7:  "DISCARD",
	10: "BROTLI_BSDIFF",
	9:  "PUFFDIFF",
}

var InstallOperation_Type_value = map[string]int32{
	"REPLACE":       0,
	"REPLACE_BZ":    1,
	"MOVE":          2,
	"BSDIFF":        3,
	"SOURCE_COPY":   4,
	"SOURCE_BSDIFF": 5,
	"REPLACE_XZ":    8,
	"ZERO":          6,
	"DISCARD":       7,
	"BROTLI_BSDIFF": 10,
	"PUFFDIFF":      9,
}

func (x InstallOperation_Type) Enum() *InstallOperation_Type {
	p := new(InstallOperation_Type)
	*p = x
	return p
}

func (x InstallOperation_Type) String() string {
	return proto.EnumName(InstallOperation_Type_name, int32(x))
}

func (x *InstallOperation_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InstallOperation_Type_value, data, "InstallOperation_Type")
	if err != nil {
		return err
	}
	*x = InstallOperation_Type(value)
	return nil
}

func (InstallOperation_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa6d72a1ce634b79, []int{4, 0}
}

// Data is packed into blocks on disk, always starting from the beginning
// of the block. If a file's data is too large for one block, it overflows
// into another block, which may or may not be the following block on the
// physical partition. An ordered list of extents is another
// representation of an ordered list of blocks. For example, a file stored
// in blocks 9, 10, 11, 2, 18, 12 (in that order) would be stored in
// extents { {9, 3}, {2, 1}, {18, 1}, {12, 1} } (in that order).
// In general, files are stored sequentially on disk, so it's more efficient
// to use extents to encode the block lists (this is effectively
// run-length encoding).
// A sentinel value (kuint64max) as the start block denotes a sparse-hole
// in a file whose block-length is specified by num_blocks.
// Signatures: Updates may be signed by the OS vendor. The client verifies
// an update's signature by hashing the entire download. The section of the
// download that contains the signature is at the end of the file, so when
// signing a file, only the part up to the signature part is signed.
// Then, the client looks inside the download's Signatures message for a
// Signature message that it knows how to handle. Generally, a client will
// only know how to handle one type of signature, but an update may contain
// many signatures to support many different types of client. Then client
// selects a Signature message and uses that, along with a known public key,
// to verify the download. The public key is expected to be part of the
// client.
type Extent struct {
	StartBlock           *uint64  `protobuf:"varint,1,opt,name=start_block,json=startBlock" json:"start_block,omitempty"`
	NumBlocks            *uint64  `protobuf:"varint,2,opt,name=num_blocks,json=numBlocks" json:"num_blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Extent) Reset()         { *m = Extent{} }
func (m *Extent) String() string { return proto.CompactTextString(m) }
func (*Extent) ProtoMessage()    {}
func (*Extent) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6d72a1ce634b79, []int{0}
}

func (m *Extent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extent.Unmarshal(m, b)
}
func (m *Extent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extent.Marshal(b, m, deterministic)
}
func (m *Extent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extent.Merge(m, src)
}
func (m *Extent) XXX_Size() int {
	return xxx_messageInfo_Extent.Size(m)
}
func (m *Extent) XXX_DiscardUnknown() {
	xxx_messageInfo_Extent.DiscardUnknown(m)
}

var xxx_messageInfo_Extent proto.InternalMessageInfo

func (m *Extent) GetStartBlock() uint64 {
	if m != nil && m.StartBlock != nil {
		return *m.StartBlock
	}
	return 0
}

func (m *Extent) GetNumBlocks() uint64 {
	if m != nil && m.NumBlocks != nil {
		return *m.NumBlocks
	}
	return 0
}

type Signatures struct {
	Signatures           []*Signatures_Signature `protobuf:"bytes,1,rep,name=signatures" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Signatures) Reset()         { *m = Signatures{} }
func (m *Signatures) String() string { return proto.CompactTextString(m) }
func (*Signatures) ProtoMessage()    {}
func (*Signatures) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6d72a1ce634b79, []int{1}
}

func (m *Signatures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signatures.Unmarshal(m, b)
}
func (m *Signatures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signatures.Marshal(b, m, deterministic)
}
func (m *Signatures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signatures.Merge(m, src)
}
func (m *Signatures) XXX_Size() int {
	return xxx_messageInfo_Signatures.Size(m)
}
func (m *Signatures) XXX_DiscardUnknown() {
	xxx_messageInfo_Signatures.DiscardUnknown(m)
}

var xxx_messageInfo_Signatures proto.InternalMessageInfo

func (m *Signatures) GetSignatures() []*Signatures_Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type Signatures_Signature struct {
	Version              *uint32  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signatures_Signature) Reset()         { *m = Signatures_Signature{} }
func (m *Signatures_Signature) String() string { return proto.CompactTextString(m) }
func (*Signatures_Signature) ProtoMessage()    {}
func (*Signatures_Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6d72a1ce634b79, []int{1, 0}
}

func (m *Signatures_Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signatures_Signature.Unmarshal(m, b)
}
func (m *Signatures_Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signatures_Signature.Marshal(b, m, deterministic)
}
func (m *Signatures_Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signatures_Signature.Merge(m, src)
}
func (m *Signatures_Signature) XXX_Size() int {
	return xxx_messageInfo_Signatures_Signature.Size(m)
}
func (m *Signatures_Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signatures_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signatures_Signature proto.InternalMessageInfo

func (m *Signatures_Signature) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Signatures_Signature) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PartitionInfo struct {
	Size                 *uint64  `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionInfo) Reset()         { *m = PartitionInfo{} }
func (m *PartitionInfo) String() string { return proto.CompactTextString(m) }
func (*PartitionInfo) ProtoMessage()    {}
func (*PartitionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6d72a1ce634b79, []int{2}
}

func (m *PartitionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionInfo.Unmarshal(m, b)
}
func (m *PartitionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionInfo.Marshal(b, m, deterministic)
}
func (m *PartitionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionInfo.Merge(m, src)
}
func (m *PartitionInfo) XXX_Size() int {
	return xxx_messageInfo_PartitionInfo.Size(m)
}
func (m *PartitionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionInfo proto.InternalMessageInfo

func (m *PartitionInfo) GetSize() uint64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *PartitionInfo) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// Describe an image we are based on in a human friendly way.
// Examples:
//   dev-channel, x86-alex, 1.2.3, mp-v3
//   nplusone-channel, x86-alex, 1.2.4, mp-v3, dev-channel, 1.2.3
//
// All fields will be set, if this message is present.
type ImageInfo struct {
	Board   *string `protobuf:"bytes,1,opt,name=board" json:"board,omitempty"`
	Key     *string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Channel *string `protobuf:"bytes,3,opt,name=channel" json:"channel,omitempty"`
	Version *string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	// If these values aren't present, they should be assumed to match
	// the equivalent value above. They are normally only different for
	// special image types such as nplusone images.
	BuildChannel         *string  `protobuf:"bytes,5,opt,name=build_channel,json=buildChannel" json:"build_channel,omitempty"`
	BuildVersion         *string  `protobuf:"bytes,6,opt,name=build_version,json=buildVersion" json:"build_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageInfo) Reset()         { *m = ImageInfo{} }
func (m *ImageInfo) String() string { return proto.CompactTextString(m) }
func (*ImageInfo) ProtoMessage()    {}
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6d72a1ce634b79, []int{3}
}

func (m *ImageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageInfo.Unmarshal(m, b)
}
func (m *ImageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageInfo.Marshal(b, m, deterministic)
}
func (m *ImageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageInfo.Merge(m, src)
}
func (m *ImageInfo) XXX_Size() int {
	return xxx_messageInfo_ImageInfo.Size(m)
}
func (m *ImageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ImageInfo proto.InternalMessageInfo

func (m *ImageInfo) GetBoard() string {
	if m != nil && m.Board != nil {
		return *m.Board
	}
	return ""
}

func (m *ImageInfo) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ImageInfo) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *ImageInfo) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *ImageInfo) GetBuildChannel() string {
	if m != nil && m.BuildChannel != nil {
		return *m.BuildChannel
	}
	return ""
}

func (m *ImageInfo) GetBuildVersion() string {
	if m != nil && m.BuildVersion != nil {
		return *m.BuildVersion
	}
	return ""
}

type InstallOperation struct {
	Type *InstallOperation_Type `protobuf:"varint,1,req,name=type,enum=chromeos_update_engine.InstallOperation_Type" json:"type,omitempty"`
	// The offset into the delta file (after the protobuf)
	// where the data (if any) is stored
	DataOffset *uint32 `protobuf:"varint,2,opt,name=data_offset,json=dataOffset" json:"data_offset,omitempty"`
	// The length of the data in the delta file
	DataLength *uint32 `protobuf:"varint,3,opt,name=data_length,json=dataLength" json:"data_length,omitempty"`
	// Ordered list of extents that are read from (if any) and written to.
	SrcExtents []*Extent `protobuf:"bytes,4,rep,name=src_extents,json=srcExtents" json:"src_extents,omitempty"`
	// Byte length of src, equal to the number of blocks in src_extents *
	// block_size. It is used for BSDIFF and SOURCE_BSDIFF, because we need to
	// pass that external program the number of bytes to read from the blocks we
	// pass it.  This is not used in any other operation.
	SrcLength  *uint64   `protobuf:"varint,5,opt,name=src_length,json=srcLength" json:"src_length,omitempty"`
	DstExtents []*Extent `protobuf:"bytes,6,rep,name=dst_extents,json=dstExtents" json:"dst_extents,omitempty"`
	// Byte length of dst, equal to the number of blocks in dst_extents *
	// block_size. Used for BSDIFF and SOURCE_BSDIFF, but not in any other
	// operation.
	DstLength *uint64 `protobuf:"varint,7,opt,name=dst_length,json=dstLength" json:"dst_length,omitempty"`
	// Optional SHA 256 hash of the blob associated with this operation.
	// This is used as a primary validation for http-based downloads and
	// as a defense-in-depth validation for https-based downloads. If
	// the operation doesn't refer to any blob, this field will have
	// zero bytes.
	DataSha256Hash []byte `protobuf:"bytes,8,opt,name=data_sha256_hash,json=dataSha256Hash" json:"data_sha256_hash,omitempty"`
	// Indicates the SHA 256 hash of the source data referenced in src_extents at
	// the time of applying the operation. If present, the update_engine daemon
	// MUST read and verify the source data before applying the operation.
	SrcSha256Hash        []byte   `protobuf:"bytes,9,opt,name=src_sha256_hash,json=srcSha256Hash" json:"src_sha256_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallOperation) Reset()         { *m = InstallOperation{} }
func (m *InstallOperation) String() string { return proto.CompactTextString(m) }
func (*InstallOperation) ProtoMessage()    {}
func (*InstallOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6d72a1ce634b79, []int{4}
}

func (m *InstallOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallOperation.Unmarshal(m, b)
}
func (m *InstallOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallOperation.Marshal(b, m, deterministic)
}
func (m *InstallOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallOperation.Merge(m, src)
}
func (m *InstallOperation) XXX_Size() int {
	return xxx_messageInfo_InstallOperation.Size(m)
}
func (m *InstallOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallOperation.DiscardUnknown(m)
}

var xxx_messageInfo_InstallOperation proto.InternalMessageInfo

func (m *InstallOperation) GetType() InstallOperation_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return InstallOperation_REPLACE
}

func (m *InstallOperation) GetDataOffset() uint32 {
	if m != nil && m.DataOffset != nil {
		return *m.DataOffset
	}
	return 0
}

func (m *InstallOperation) GetDataLength() uint32 {
	if m != nil && m.DataLength != nil {
		return *m.DataLength
	}
	return 0
}

func (m *InstallOperation) GetSrcExtents() []*Extent {
	if m != nil {
		return m.SrcExtents
	}
	return nil
}

func (m *InstallOperation) GetSrcLength() uint64 {
	if m != nil && m.SrcLength != nil {
		return *m.SrcLength
	}
	return 0
}

func (m *InstallOperation) GetDstExtents() []*Extent {
	if m != nil {
		return m.DstExtents
	}
	return nil
}

func (m *InstallOperation) GetDstLength() uint64 {
	if m != nil && m.DstLength != nil {
		return *m.DstLength
	}
	return 0
}

func (m *InstallOperation) GetDataSha256Hash() []byte {
	if m != nil {
		return m.DataSha256Hash
	}
	return nil
}

func (m *InstallOperation) GetSrcSha256Hash() []byte {
	if m != nil {
		return m.SrcSha256Hash
	}
	return nil
}

// Describes the update to apply to a single partition.
type PartitionUpdate struct {
	// A platform-specific name to identify the partition set being updated. For
	// example, in Chrome OS this could be "ROOT" or "KERNEL".
	PartitionName *string `protobuf:"bytes,1,req,name=partition_name,json=partitionName" json:"partition_name,omitempty"`
	// Whether this partition carries a filesystem with post-install program that
	// must be run to finalize the update process. See also |postinstall_path| and
	// |filesystem_type|.
	RunPostinstall *bool `protobuf:"varint,2,opt,name=run_postinstall,json=runPostinstall" json:"run_postinstall,omitempty"`
	// The path of the executable program to run during the post-install step,
	// relative to the root of this filesystem. If not set, the default "postinst"
	// will be used. This setting is only used when |run_postinstall| is set and
	// true.
	PostinstallPath *string `protobuf:"bytes,3,opt,name=postinstall_path,json=postinstallPath" json:"postinstall_path,omitempty"`
	// The filesystem type as passed to the mount(2) syscall when mounting the new
	// filesystem to run the post-install program. If not set, a fixed list of
	// filesystems will be attempted. This setting is only used if
	// |run_postinstall| is set and true.
	FilesystemType *string `protobuf:"bytes,4,opt,name=filesystem_type,json=filesystemType" json:"filesystem_type,omitempty"`
	// If present, a list of signatures of the new_partition_info.hash signed with
	// different keys. If the update_engine daemon requires vendor-signed images
	// and has its public key installed, one of the signatures should be valid
	// for /postinstall to run.
	NewPartitionSignature []*Signatures_Signature `protobuf:"bytes,5,rep,name=new_partition_signature,json=newPartitionSignature" json:"new_partition_signature,omitempty"`
	OldPartitionInfo      *PartitionInfo          `protobuf:"bytes,6,opt,name=old_partition_info,json=oldPartitionInfo" json:"old_partition_info,omitempty"`
	NewPartitionInfo      *PartitionInfo          `protobuf:"bytes,7,opt,name=new_partition_info,json=newPartitionInfo" json:"new_partition_info,omitempty"`
	// The list of operations to be performed to apply this PartitionUpdate. The
	// associated operation blobs (in operations[i].data_offset, data_length)
	// should be stored contiguously and in the same order.
	Operations []*InstallOperation `protobuf:"bytes,8,rep,name=operations" json:"operations,omitempty"`
	// Whether a failure in the postinstall step for this partition should be
	// ignored.
	PostinstallOptional  *bool    `protobuf:"varint,9,opt,name=postinstall_optional,json=postinstallOptional" json:"postinstall_optional,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionUpdate) Reset()         { *m = PartitionUpdate{} }
func (m *PartitionUpdate) String() string { return proto.CompactTextString(m) }
func (*PartitionUpdate) ProtoMessage()    {}
func (*PartitionUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6d72a1ce634b79, []int{5}
}

func (m *PartitionUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionUpdate.Unmarshal(m, b)
}
func (m *PartitionUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionUpdate.Marshal(b, m, deterministic)
}
func (m *PartitionUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionUpdate.Merge(m, src)
}
func (m *PartitionUpdate) XXX_Size() int {
	return xxx_messageInfo_PartitionUpdate.Size(m)
}
func (m *PartitionUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionUpdate proto.InternalMessageInfo

func (m *PartitionUpdate) GetPartitionName() string {
	if m != nil && m.PartitionName != nil {
		return *m.PartitionName
	}
	return ""
}

func (m *PartitionUpdate) GetRunPostinstall() bool {
	if m != nil && m.RunPostinstall != nil {
		return *m.RunPostinstall
	}
	return false
}

func (m *PartitionUpdate) GetPostinstallPath() string {
	if m != nil && m.PostinstallPath != nil {
		return *m.PostinstallPath
	}
	return ""
}

func (m *PartitionUpdate) GetFilesystemType() string {
	if m != nil && m.FilesystemType != nil {
		return *m.FilesystemType
	}
	return ""
}

func (m *PartitionUpdate) GetNewPartitionSignature() []*Signatures_Signature {
	if m != nil {
		return m.NewPartitionSignature
	}
	return nil
}

func (m *PartitionUpdate) GetOldPartitionInfo() *PartitionInfo {
	if m != nil {
		return m.OldPartitionInfo
	}
	return nil
}

func (m *PartitionUpdate) GetNewPartitionInfo() *PartitionInfo {
	if m != nil {
		return m.NewPartitionInfo
	}
	return nil
}

func (m *PartitionUpdate) GetOperations() []*InstallOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *PartitionUpdate) GetPostinstallOptional() bool {
	if m != nil && m.PostinstallOptional != nil {
		return *m.PostinstallOptional
	}
	return false
}

type DeltaArchiveManifest struct {
	// Only present in major version = 1. List of install operations for the
	// kernel and rootfs partitions. For major version = 2 see the |partitions|
	// field.
	InstallOperations       []*InstallOperation `protobuf:"bytes,1,rep,name=install_operations,json=installOperations" json:"install_operations,omitempty"`
	KernelInstallOperations []*InstallOperation `protobuf:"bytes,2,rep,name=kernel_install_operations,json=kernelInstallOperations" json:"kernel_install_operations,omitempty"`
	// (At time of writing) usually 4096
	BlockSize *uint32 `protobuf:"varint,3,opt,name=block_size,json=blockSize,def=4096" json:"block_size,omitempty"`
	// If signatures are present, the offset into the blobs, generally
	// tacked onto the end of the file, and the length. We use an offset
	// rather than a bool to allow for more flexibility in future file formats.
	// If either is absent, it means signatures aren't supported in this
	// file.
	SignaturesOffset *uint64 `protobuf:"varint,4,opt,name=signatures_offset,json=signaturesOffset" json:"signatures_offset,omitempty"`
	SignaturesSize   *uint64 `protobuf:"varint,5,opt,name=signatures_size,json=signaturesSize" json:"signatures_size,omitempty"`
	// Only present in major version = 1. Partition metadata used to validate the
	// update. For major version = 2 see the |partitions| field.
	OldKernelInfo *PartitionInfo `protobuf:"bytes,6,opt,name=old_kernel_info,json=oldKernelInfo" json:"old_kernel_info,omitempty"`
	NewKernelInfo *PartitionInfo `protobuf:"bytes,7,opt,name=new_kernel_info,json=newKernelInfo" json:"new_kernel_info,omitempty"`
	OldRootfsInfo *PartitionInfo `protobuf:"bytes,8,opt,name=old_rootfs_info,json=oldRootfsInfo" json:"old_rootfs_info,omitempty"`
	NewRootfsInfo *PartitionInfo `protobuf:"bytes,9,opt,name=new_rootfs_info,json=newRootfsInfo" json:"new_rootfs_info,omitempty"`
	// old_image_info will only be present for delta images.
	OldImageInfo *ImageInfo `protobuf:"bytes,10,opt,name=old_image_info,json=oldImageInfo" json:"old_image_info,omitempty"`
	NewImageInfo *ImageInfo `protobuf:"bytes,11,opt,name=new_image_info,json=newImageInfo" json:"new_image_info,omitempty"`
	// The minor version, also referred as "delta version", of the payload.
	MinorVersion *uint32 `protobuf:"varint,12,opt,name=minor_version,json=minorVersion,def=0" json:"minor_version,omitempty"`
	// Only present in major version >= 2. List of partitions that will be
	// updated, in the order they will be updated. This field replaces the
	// |install_operations|, |kernel_install_operations| and the
	// |{old,new}_{kernel,rootfs}_info| fields used in major version = 1. This
	// array can have more than two partitions if needed, and they are identified
	// by the partition name.
	Partitions []*PartitionUpdate `protobuf:"bytes,13,rep,name=partitions" json:"partitions,omitempty"`
	// The maximum timestamp of the OS allowed to apply this payload.
	// Can be used to prevent downgrading the OS.
	MaxTimestamp         *int64   `protobuf:"varint,14,opt,name=max_timestamp,json=maxTimestamp" json:"max_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeltaArchiveManifest) Reset()         { *m = DeltaArchiveManifest{} }
func (m *DeltaArchiveManifest) String() string { return proto.CompactTextString(m) }
func (*DeltaArchiveManifest) ProtoMessage()    {}
func (*DeltaArchiveManifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6d72a1ce634b79, []int{6}
}

func (m *DeltaArchiveManifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeltaArchiveManifest.Unmarshal(m, b)
}
func (m *DeltaArchiveManifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeltaArchiveManifest.Marshal(b, m, deterministic)
}
func (m *DeltaArchiveManifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeltaArchiveManifest.Merge(m, src)
}
func (m *DeltaArchiveManifest) XXX_Size() int {
	return xxx_messageInfo_DeltaArchiveManifest.Size(m)
}
func (m *DeltaArchiveManifest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeltaArchiveManifest.DiscardUnknown(m)
}

var xxx_messageInfo_DeltaArchiveManifest proto.InternalMessageInfo

const Default_DeltaArchiveManifest_BlockSize uint32 = 4096
const Default_DeltaArchiveManifest_MinorVersion uint32 = 0

func (m *DeltaArchiveManifest) GetInstallOperations() []*InstallOperation {
	if m != nil {
		return m.InstallOperations
	}
	return nil
}

func (m *DeltaArchiveManifest) GetKernelInstallOperations() []*InstallOperation {
	if m != nil {
		return m.KernelInstallOperations
	}
	return nil
}

func (m *DeltaArchiveManifest) GetBlockSize() uint32 {
	if m != nil && m.BlockSize != nil {
		return *m.BlockSize
	}
	return Default_DeltaArchiveManifest_BlockSize
}

func (m *DeltaArchiveManifest) GetSignaturesOffset() uint64 {
	if m != nil && m.SignaturesOffset != nil {
		return *m.SignaturesOffset
	}
	return 0
}

func (m *DeltaArchiveManifest) GetSignaturesSize() uint64 {
	if m != nil && m.SignaturesSize != nil {
		return *m.SignaturesSize
	}
	return 0
}

func (m *DeltaArchiveManifest) GetOldKernelInfo() *PartitionInfo {
	if m != nil {
		return m.OldKernelInfo
	}
	return nil
}

func (m *DeltaArchiveManifest) GetNewKernelInfo() *PartitionInfo {
	if m != nil {
		return m.NewKernelInfo
	}
	return nil
}

func (m *DeltaArchiveManifest) GetOldRootfsInfo() *PartitionInfo {
	if m != nil {
		return m.OldRootfsInfo
	}
	return nil
}

func (m *DeltaArchiveManifest) GetNewRootfsInfo() *PartitionInfo {
	if m != nil {
		return m.NewRootfsInfo
	}
	return nil
}

func (m *DeltaArchiveManifest) GetOldImageInfo() *ImageInfo {
	if m != nil {
		return m.OldImageInfo
	}
	return nil
}

func (m *DeltaArchiveManifest) GetNewImageInfo() *ImageInfo {
	if m != nil {
		return m.NewImageInfo
	}
	return nil
}

func (m *DeltaArchiveManifest) GetMinorVersion() uint32 {
	if m != nil && m.MinorVersion != nil {
		return *m.MinorVersion
	}
	return Default_DeltaArchiveManifest_MinorVersion
}

func (m *DeltaArchiveManifest) GetPartitions() []*PartitionUpdate {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *DeltaArchiveManifest) GetMaxTimestamp() int64 {
	if m != nil && m.MaxTimestamp != nil {
		return *m.MaxTimestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("chromeos_update_engine.InstallOperation_Type", InstallOperation_Type_name, InstallOperation_Type_value)
	proto.RegisterType((*Extent)(nil), "chromeos_update_engine.Extent")
	proto.RegisterType((*Signatures)(nil), "chromeos_update_engine.Signatures")
	proto.RegisterType((*Signatures_Signature)(nil), "chromeos_update_engine.Signatures.Signature")
	proto.RegisterType((*PartitionInfo)(nil), "chromeos_update_engine.PartitionInfo")
	proto.RegisterType((*ImageInfo)(nil), "chromeos_update_engine.ImageInfo")
	proto.RegisterType((*InstallOperation)(nil), "chromeos_update_engine.InstallOperation")
	proto.RegisterType((*PartitionUpdate)(nil), "chromeos_update_engine.PartitionUpdate")
	proto.RegisterType((*DeltaArchiveManifest)(nil), "chromeos_update_engine.DeltaArchiveManifest")
}

func init() { proto.RegisterFile("update_metadata.proto", fileDescriptor_fa6d72a1ce634b79) }

var fileDescriptor_fa6d72a1ce634b79 = []byte{
	// 1047 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xef, 0x6e, 0x23, 0xb5,
	0x17, 0xfd, 0xa5, 0x99, 0x36, 0x99, 0x9b, 0x4c, 0x32, 0xeb, 0x5f, 0x97, 0x0d, 0x48, 0x40, 0xc9,
	0x6a, 0xb7, 0x41, 0x40, 0xb5, 0x54, 0xb0, 0x68, 0xf9, 0x82, 0xfa, 0x27, 0x6d, 0x23, 0x5a, 0x12,
	0x39, 0xed, 0x02, 0xfd, 0x32, 0x72, 0x33, 0x4e, 0x33, 0xea, 0x8c, 0x1d, 0x8d, 0x9d, 0x6d, 0xcb,
	0x7b, 0xf0, 0x08, 0x3c, 0x00, 0x12, 0x4f, 0xc0, 0x1b, 0xf0, 0x46, 0xc8, 0x77, 0xfe, 0x76, 0xd9,
	0x8a, 0x86, 0x6f, 0xf6, 0xb9, 0xf7, 0x1e, 0x9f, 0xeb, 0x9c, 0xb9, 0x0e, 0x3c, 0x5e, 0xcc, 0x7d,
	0xa6, 0xb9, 0x17, 0x71, 0xcd, 0x7c, 0xa6, 0xd9, 0xd6, 0x3c, 0x96, 0x5a, 0x92, 0xf7, 0x26, 0xb3,
	0x58, 0x46, 0x5c, 0x2a, 0x2f, 0x8d, 0x73, 0x71, 0x19, 0x08, 0xde, 0x3d, 0x82, 0xb5, 0xfe, 0x8d,
	0xe6, 0x42, 0x93, 0x8f, 0xa1, 0xa1, 0x34, 0x8b, 0xb5, 0x77, 0x11, 0xca, 0xc9, 0x55, 0xa7, 0xb2,
	0x51, 0xe9, 0x59, 0x14, 0x10, 0xda, 0x35, 0x08, 0xf9, 0x10, 0x40, 0x2c, 0xa2, 0x24, 0xac, 0x3a,
	0x2b, 0x18, 0xb7, 0xc5, 0x22, 0xc2, 0xa8, 0xea, 0xfe, 0x5a, 0x01, 0x18, 0x07, 0x97, 0x82, 0xe9,
	0x45, 0xcc, 0x15, 0x39, 0x06, 0x50, 0xf9, 0xae, 0x53, 0xd9, 0xa8, 0xf6, 0x1a, 0xdb, 0x9f, 0x6f,
	0xbd, 0x5b, 0xc5, 0x56, 0x51, 0x57, 0x2c, 0x69, 0xa9, 0xfe, 0x83, 0x57, 0x60, 0xe7, 0x01, 0xd2,
	0x81, 0xda, 0x1b, 0x1e, 0xab, 0x40, 0x0a, 0x54, 0xe9, 0xd0, 0x6c, 0x4b, 0x08, 0x58, 0xa6, 0x67,
	0x14, 0xd7, 0xa4, 0xb8, 0xee, 0x7e, 0x03, 0xce, 0x88, 0xc5, 0x3a, 0xd0, 0x81, 0x14, 0x03, 0x31,
	0x95, 0x26, 0x49, 0x05, 0xbf, 0xf0, 0xb4, 0x43, 0x5c, 0x1b, 0x6c, 0xc6, 0xd4, 0x2c, 0x2b, 0x34,
	0xeb, 0xee, 0xef, 0x15, 0xb0, 0x07, 0x11, 0xbb, 0xe4, 0x58, 0xb5, 0x0e, 0xab, 0x17, 0x92, 0xc5,
	0x3e, 0x96, 0xd9, 0x34, 0xd9, 0x10, 0x17, 0xaa, 0x57, 0xfc, 0x16, 0xcb, 0x6c, 0x6a, 0x96, 0x46,
	0xdc, 0x64, 0xc6, 0x84, 0xe0, 0x61, 0xa7, 0x8a, 0x68, 0xb6, 0x2d, 0xcb, 0xb6, 0x92, 0x48, 0x26,
	0xfb, 0x29, 0x38, 0x17, 0x8b, 0x20, 0xf4, 0xbd, 0xac, 0x72, 0x15, 0xe3, 0x4d, 0x04, 0xf7, 0xd2,
	0xf2, 0x3c, 0x29, 0x23, 0x59, 0x2b, 0x25, 0xbd, 0x4e, 0xb0, 0xee, 0x5f, 0x16, 0xb8, 0x03, 0xa1,
	0x34, 0x0b, 0xc3, 0xe1, 0x9c, 0xc7, 0xcc, 0x34, 0x4d, 0x76, 0xc0, 0xd2, 0xb7, 0x73, 0xd3, 0xf0,
	0x4a, 0xaf, 0xb5, 0xfd, 0xc5, 0x7d, 0x3f, 0xc2, 0xdb, 0x75, 0x5b, 0xa7, 0xb7, 0x73, 0x4e, 0xb1,
	0xd4, 0x98, 0xc3, 0x5c, 0xa6, 0x27, 0xa7, 0x53, 0xc5, 0x35, 0xf6, 0xeb, 0x50, 0x30, 0xd0, 0x10,
	0x91, 0x3c, 0x21, 0xe4, 0xe2, 0x52, 0xcf, 0xb0, 0xf5, 0x34, 0xe1, 0x18, 0x11, 0xf2, 0x1d, 0x34,
	0x54, 0x3c, 0xf1, 0x38, 0x9a, 0x4d, 0x75, 0x2c, 0x34, 0xc4, 0x47, 0xf7, 0x69, 0x49, 0x3c, 0x49,
	0x41, 0xc5, 0x93, 0x64, 0xa9, 0x8c, 0xfd, 0x0c, 0x41, 0x7a, 0xc0, 0x6a, 0x62, 0x3f, 0x15, 0x4f,
	0x0a, 0x7e, 0x5f, 0xe9, 0x9c, 0x7f, 0xed, 0x61, 0xfc, 0xbe, 0xd2, 0x25, 0x7e, 0x43, 0x90, 0xf2,
	0xd7, 0x12, 0x7e, 0x5f, 0xe9, 0x94, 0xbf, 0x07, 0x2e, 0x36, 0xa8, 0x66, 0x6c, 0xfb, 0xeb, 0x97,
	0x1e, 0xba, 0xa5, 0x8e, 0x6e, 0x69, 0x19, 0x7c, 0x8c, 0xf0, 0x11, 0x53, 0x33, 0xf2, 0x1c, 0xda,
	0x46, 0x68, 0x39, 0xd1, 0xc6, 0x44, 0x47, 0xc5, 0x93, 0x22, 0xaf, 0xfb, 0x5b, 0x05, 0x2c, 0x73,
	0xc5, 0xa4, 0x01, 0x35, 0xda, 0x1f, 0x1d, 0xef, 0xec, 0xf5, 0xdd, 0xff, 0x91, 0x16, 0x40, 0xba,
	0xf1, 0x76, 0xcf, 0xdd, 0x0a, 0xa9, 0x83, 0x75, 0x32, 0x7c, 0xdd, 0x77, 0x57, 0x08, 0xc0, 0xda,
	0xee, 0x78, 0x7f, 0x70, 0x70, 0xe0, 0x56, 0x49, 0x1b, 0x1a, 0xe3, 0xe1, 0x19, 0xdd, 0xeb, 0x7b,
	0x7b, 0xc3, 0xd1, 0xcf, 0xae, 0x45, 0x1e, 0x81, 0x93, 0x02, 0x69, 0xce, 0x6a, 0x99, 0xe9, 0xa7,
	0x73, 0xb7, 0x6e, 0x98, 0xce, 0xfb, 0x74, 0xe8, 0xae, 0x99, 0x03, 0xf7, 0x07, 0xe3, 0xbd, 0x1d,
	0xba, 0xef, 0xd6, 0x4c, 0xe5, 0x2e, 0x1d, 0x9e, 0x1e, 0x0f, 0xb2, 0x4a, 0x20, 0x4d, 0xa8, 0x8f,
	0xce, 0x0e, 0x0e, 0x70, 0x67, 0x77, 0xff, 0xb4, 0xa0, 0x9d, 0x7f, 0x41, 0x67, 0x78, 0x8d, 0xe4,
	0x19, 0xb4, 0xe6, 0x19, 0xe4, 0x09, 0x16, 0x25, 0xe6, 0xb2, 0xa9, 0x93, 0xa3, 0x3f, 0xb0, 0x88,
	0x93, 0x4d, 0x68, 0xc7, 0x0b, 0xe1, 0xcd, 0xa5, 0xd2, 0x41, 0xe2, 0x2e, 0xb4, 0x4e, 0x9d, 0xb6,
	0xe2, 0x85, 0x18, 0x15, 0x28, 0xf9, 0x14, 0xdc, 0x52, 0x92, 0x37, 0x67, 0xa9, 0x87, 0x6c, 0xda,
	0x2e, 0xe1, 0x23, 0xa6, 0x67, 0x86, 0x73, 0x1a, 0x84, 0x5c, 0xdd, 0x2a, 0xcd, 0x23, 0x0f, 0x8d,
	0x9d, 0x7c, 0x4e, 0xad, 0x02, 0xc6, 0x6b, 0xf5, 0xe1, 0x89, 0xe0, 0xd7, 0x5e, 0xa1, 0x33, 0x9f,
	0x27, 0x9d, 0xd5, 0xff, 0x30, 0x8e, 0x1e, 0x0b, 0x7e, 0x9d, 0x5f, 0x43, 0x31, 0x8c, 0xc6, 0x40,
	0x64, 0xe8, 0x97, 0x4e, 0x09, 0xc4, 0x54, 0xe2, 0xb7, 0xd9, 0xd8, 0x7e, 0x76, 0xdf, 0x01, 0x77,
	0x06, 0x12, 0x75, 0x65, 0xe8, 0xdf, 0x1d, 0x51, 0x63, 0x20, 0x77, 0xa5, 0x23, 0x69, 0x6d, 0x29,
	0xd2, 0xb2, 0x5c, 0x24, 0x3d, 0x02, 0x90, 0xd9, 0xb7, 0xad, 0x3a, 0x75, 0xbc, 0x82, 0xde, 0x43,
	0x87, 0x01, 0x2d, 0xd5, 0x92, 0x2f, 0x61, 0xbd, 0xfc, 0x6b, 0xc9, 0xb9, 0x81, 0x59, 0x88, 0x36,
	0xaf, 0xd3, 0xff, 0x97, 0x62, 0xc3, 0x34, 0xd4, 0xfd, 0xa3, 0x06, 0xeb, 0xfb, 0x3c, 0xd4, 0x6c,
	0x27, 0x9e, 0xcc, 0x82, 0x37, 0xfc, 0x84, 0x89, 0x60, 0xca, 0x95, 0x26, 0x3f, 0x02, 0x29, 0x78,
	0x72, 0x75, 0x95, 0x25, 0xd5, 0x3d, 0x0a, 0xde, 0x42, 0x14, 0xf1, 0xe1, 0xfd, 0x2b, 0x1e, 0x0b,
	0x1e, 0x7a, 0xef, 0xe0, 0x5f, 0x59, 0x92, 0xff, 0x49, 0x42, 0x35, 0xf8, 0xc7, 0x29, 0x4f, 0x01,
	0xf0, 0x41, 0xf4, 0xf0, 0x49, 0xc1, 0xb1, 0xf7, 0xad, 0xf5, 0xd5, 0x8b, 0x57, 0x2f, 0xa9, 0x8d,
	0xf8, 0xd8, 0xbc, 0x2e, 0x9f, 0xc1, 0xa3, 0xe2, 0x2d, 0xcb, 0x66, 0xa8, 0x85, 0x13, 0xc6, 0x2d,
	0x02, 0xe9, 0x24, 0xdd, 0x84, 0x76, 0x29, 0x19, 0x69, 0x93, 0x61, 0xd7, 0x2a, 0x60, 0x64, 0x3d,
	0x81, 0xb6, 0x71, 0x5e, 0xde, 0xe4, 0xb2, 0xb6, 0x73, 0x64, 0xe8, 0x7f, 0x9f, 0xb6, 0x35, 0x95,
	0x86, 0xce, 0x78, 0xae, 0x4c, 0xb7, 0x94, 0xe1, 0x1c, 0xc1, 0xaf, 0xef, 0xd2, 0x19, 0x75, 0xb1,
	0x94, 0x7a, 0xaa, 0x12, 0xba, 0xfa, 0xb2, 0xea, 0x28, 0x16, 0x97, 0xd5, 0x95, 0xe9, 0xec, 0x65,
	0xd5, 0x95, 0xe8, 0x0e, 0xa1, 0x65, 0xd4, 0x05, 0xe6, 0x79, 0x4f, 0xd8, 0x00, 0xd9, 0x3e, 0xb9,
	0xd7, 0x11, 0xd9, 0x1f, 0x01, 0xda, 0x94, 0xa1, 0x5f, 0xfc, 0x2d, 0x38, 0x84, 0x96, 0xd1, 0x55,
	0x22, 0x6a, 0x3c, 0x98, 0x48, 0xf0, 0xeb, 0x82, 0xe8, 0x39, 0x38, 0x51, 0x20, 0x64, 0x9c, 0x3f,
	0xef, 0x4d, 0xf4, 0x52, 0xe5, 0x05, 0x6d, 0x22, 0x9e, 0xbe, 0xf0, 0xe4, 0x10, 0x20, 0x1f, 0x0b,
	0xaa, 0xe3, 0xa0, 0x8f, 0x37, 0xff, 0xf5, 0x0e, 0x92, 0xb1, 0x4d, 0x4b, 0xa5, 0xe6, 0xff, 0x44,
	0xc4, 0x6e, 0x3c, 0x1d, 0x44, 0x5c, 0x69, 0x16, 0xcd, 0x3b, 0xad, 0x8d, 0x4a, 0xaf, 0x4a, 0x9b,
	0x11, 0xbb, 0x39, 0xcd, 0xb0, 0xdd, 0x95, 0xa3, 0xea, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3e,
	0x35, 0x0d, 0x12, 0x52, 0x0a, 0x00, 0x00,
}
