// Code generated by protoc-gen-go.
// source: bloat_report.proto
// DO NOT EDIT!

/*
Package pganalyze_collector is a generated protocol buffer package.

It is generated from these files:
	bloat_report.proto
	buffercache_report.proto
	compact_activity_snapshot.proto
	compact_log_snapshot.proto
	compact_snapshot.proto
	compact_system_snapshot.proto
	full_snapshot.proto
	report.proto
	sequence_report.proto
	shared.proto
	vacuum_report.proto

It has these top-level messages:
	BloatReportData
	RelationBloatStatistic
	IndexBloatStatistic
	BuffercacheReportData
	BuffercacheEntry
	CompactActivitySnapshot
	Backend
	CompactLogSnapshot
	LogFileReference
	LogLineInformation
	QuerySample
	CompactSnapshot
	CompactSystemSnapshot
	FullSnapshot
	CollectorStatistic
	RoleInformation
	DatabaseInformation
	Setting
	Replication
	StandbyReference
	StandbyInformation
	StandbyStatistic
	TablespaceReference
	TablespaceInformation
	QueryStatistic
	HistoricQueryStatistics
	RelationInformation
	RelationStatistic
	RelationEvent
	IndexInformation
	IndexStatistic
	FunctionInformation
	FunctionStatistic
	Report
	SequenceReportData
	SequenceReference
	SequenceInformation
	SerialColumnInformation
	NullString
	NullTimestamp
	PostgresVersion
	RoleReference
	DatabaseReference
	RelationReference
	IndexReference
	FunctionReference
	QueryReference
	QueryInformation
	System
	SystemInformation
	SystemInformationSelfHosted
	SystemInformationAmazonRDS
	SchedulerStatistic
	MemoryStatistic
	CPUInformation
	CPUReference
	CPUStatistic
	NetworkReference
	NetworkStatistic
	DiskReference
	DiskInformation
	DiskStatistic
	DiskPartitionReference
	DiskPartitionInformation
	DiskPartitionStatistic
	VacuumReportData
	VacuumStatistic
*/
package pganalyze_collector

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BloatLookupMethod int32

const (
	BloatLookupMethod_ESTIMATE_FAST BloatLookupMethod = 0
	BloatLookupMethod_ESTIMATE_SLOW BloatLookupMethod = 1
	BloatLookupMethod_FULL_SCAN     BloatLookupMethod = 2
)

var BloatLookupMethod_name = map[int32]string{
	0: "ESTIMATE_FAST",
	1: "ESTIMATE_SLOW",
	2: "FULL_SCAN",
}
var BloatLookupMethod_value = map[string]int32{
	"ESTIMATE_FAST": 0,
	"ESTIMATE_SLOW": 1,
	"FULL_SCAN":     2,
}

func (x BloatLookupMethod) String() string {
	return proto.EnumName(BloatLookupMethod_name, int32(x))
}
func (BloatLookupMethod) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BloatReportData struct {
	DatabaseReferences      []*DatabaseReference      `protobuf:"bytes,10,rep,name=database_references,json=databaseReferences" json:"database_references,omitempty"`
	RelationReferences      []*RelationReference      `protobuf:"bytes,11,rep,name=relation_references,json=relationReferences" json:"relation_references,omitempty"`
	IndexReferences         []*IndexReference         `protobuf:"bytes,12,rep,name=index_references,json=indexReferences" json:"index_references,omitempty"`
	RelationBloatStatistics []*RelationBloatStatistic `protobuf:"bytes,20,rep,name=relation_bloat_statistics,json=relationBloatStatistics" json:"relation_bloat_statistics,omitempty"`
	IndexBloatStatistics    []*IndexBloatStatistic    `protobuf:"bytes,21,rep,name=index_bloat_statistics,json=indexBloatStatistics" json:"index_bloat_statistics,omitempty"`
}

func (m *BloatReportData) Reset()                    { *m = BloatReportData{} }
func (m *BloatReportData) String() string            { return proto.CompactTextString(m) }
func (*BloatReportData) ProtoMessage()               {}
func (*BloatReportData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BloatReportData) GetDatabaseReferences() []*DatabaseReference {
	if m != nil {
		return m.DatabaseReferences
	}
	return nil
}

func (m *BloatReportData) GetRelationReferences() []*RelationReference {
	if m != nil {
		return m.RelationReferences
	}
	return nil
}

func (m *BloatReportData) GetIndexReferences() []*IndexReference {
	if m != nil {
		return m.IndexReferences
	}
	return nil
}

func (m *BloatReportData) GetRelationBloatStatistics() []*RelationBloatStatistic {
	if m != nil {
		return m.RelationBloatStatistics
	}
	return nil
}

func (m *BloatReportData) GetIndexBloatStatistics() []*IndexBloatStatistic {
	if m != nil {
		return m.IndexBloatStatistics
	}
	return nil
}

type RelationBloatStatistic struct {
	RelationIdx       int32             `protobuf:"varint,1,opt,name=relation_idx,json=relationIdx" json:"relation_idx,omitempty"`
	BloatLookupMethod BloatLookupMethod `protobuf:"varint,2,opt,name=bloat_lookup_method,json=bloatLookupMethod,enum=pganalyze.collector.BloatLookupMethod" json:"bloat_lookup_method,omitempty"`
	TotalBytes        int64             `protobuf:"varint,3,opt,name=total_bytes,json=totalBytes" json:"total_bytes,omitempty"`
	BloatBytes        int64             `protobuf:"varint,4,opt,name=bloat_bytes,json=bloatBytes" json:"bloat_bytes,omitempty"`
	LiveTupleBytes    int64             `protobuf:"varint,5,opt,name=live_tuple_bytes,json=liveTupleBytes" json:"live_tuple_bytes,omitempty"`
	LiveTupleCount    int64             `protobuf:"varint,6,opt,name=live_tuple_count,json=liveTupleCount" json:"live_tuple_count,omitempty"`
	DeadTupleBytes    int64             `protobuf:"varint,7,opt,name=dead_tuple_bytes,json=deadTupleBytes" json:"dead_tuple_bytes,omitempty"`
	DeadTupleCount    int64             `protobuf:"varint,8,opt,name=dead_tuple_count,json=deadTupleCount" json:"dead_tuple_count,omitempty"`
}

func (m *RelationBloatStatistic) Reset()                    { *m = RelationBloatStatistic{} }
func (m *RelationBloatStatistic) String() string            { return proto.CompactTextString(m) }
func (*RelationBloatStatistic) ProtoMessage()               {}
func (*RelationBloatStatistic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RelationBloatStatistic) GetRelationIdx() int32 {
	if m != nil {
		return m.RelationIdx
	}
	return 0
}

func (m *RelationBloatStatistic) GetBloatLookupMethod() BloatLookupMethod {
	if m != nil {
		return m.BloatLookupMethod
	}
	return BloatLookupMethod_ESTIMATE_FAST
}

func (m *RelationBloatStatistic) GetTotalBytes() int64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

func (m *RelationBloatStatistic) GetBloatBytes() int64 {
	if m != nil {
		return m.BloatBytes
	}
	return 0
}

func (m *RelationBloatStatistic) GetLiveTupleBytes() int64 {
	if m != nil {
		return m.LiveTupleBytes
	}
	return 0
}

func (m *RelationBloatStatistic) GetLiveTupleCount() int64 {
	if m != nil {
		return m.LiveTupleCount
	}
	return 0
}

func (m *RelationBloatStatistic) GetDeadTupleBytes() int64 {
	if m != nil {
		return m.DeadTupleBytes
	}
	return 0
}

func (m *RelationBloatStatistic) GetDeadTupleCount() int64 {
	if m != nil {
		return m.DeadTupleCount
	}
	return 0
}

type IndexBloatStatistic struct {
	IndexIdx          int32             `protobuf:"varint,1,opt,name=index_idx,json=indexIdx" json:"index_idx,omitempty"`
	BloatLookupMethod BloatLookupMethod `protobuf:"varint,2,opt,name=bloat_lookup_method,json=bloatLookupMethod,enum=pganalyze.collector.BloatLookupMethod" json:"bloat_lookup_method,omitempty"`
	TotalBytes        int64             `protobuf:"varint,3,opt,name=total_bytes,json=totalBytes" json:"total_bytes,omitempty"`
	BloatBytes        int64             `protobuf:"varint,4,opt,name=bloat_bytes,json=bloatBytes" json:"bloat_bytes,omitempty"`
	InternalPages     int64             `protobuf:"varint,5,opt,name=internal_pages,json=internalPages" json:"internal_pages,omitempty"`
	LeafPages         int64             `protobuf:"varint,6,opt,name=leaf_pages,json=leafPages" json:"leaf_pages,omitempty"`
	EmptyPages        int64             `protobuf:"varint,7,opt,name=empty_pages,json=emptyPages" json:"empty_pages,omitempty"`
	DeletedPages      int64             `protobuf:"varint,8,opt,name=deleted_pages,json=deletedPages" json:"deleted_pages,omitempty"`
	AvgLeafDensity    float64           `protobuf:"fixed64,9,opt,name=avg_leaf_density,json=avgLeafDensity" json:"avg_leaf_density,omitempty"`
}

func (m *IndexBloatStatistic) Reset()                    { *m = IndexBloatStatistic{} }
func (m *IndexBloatStatistic) String() string            { return proto.CompactTextString(m) }
func (*IndexBloatStatistic) ProtoMessage()               {}
func (*IndexBloatStatistic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IndexBloatStatistic) GetIndexIdx() int32 {
	if m != nil {
		return m.IndexIdx
	}
	return 0
}

func (m *IndexBloatStatistic) GetBloatLookupMethod() BloatLookupMethod {
	if m != nil {
		return m.BloatLookupMethod
	}
	return BloatLookupMethod_ESTIMATE_FAST
}

func (m *IndexBloatStatistic) GetTotalBytes() int64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

func (m *IndexBloatStatistic) GetBloatBytes() int64 {
	if m != nil {
		return m.BloatBytes
	}
	return 0
}

func (m *IndexBloatStatistic) GetInternalPages() int64 {
	if m != nil {
		return m.InternalPages
	}
	return 0
}

func (m *IndexBloatStatistic) GetLeafPages() int64 {
	if m != nil {
		return m.LeafPages
	}
	return 0
}

func (m *IndexBloatStatistic) GetEmptyPages() int64 {
	if m != nil {
		return m.EmptyPages
	}
	return 0
}

func (m *IndexBloatStatistic) GetDeletedPages() int64 {
	if m != nil {
		return m.DeletedPages
	}
	return 0
}

func (m *IndexBloatStatistic) GetAvgLeafDensity() float64 {
	if m != nil {
		return m.AvgLeafDensity
	}
	return 0
}

func init() {
	proto.RegisterType((*BloatReportData)(nil), "pganalyze.collector.BloatReportData")
	proto.RegisterType((*RelationBloatStatistic)(nil), "pganalyze.collector.RelationBloatStatistic")
	proto.RegisterType((*IndexBloatStatistic)(nil), "pganalyze.collector.IndexBloatStatistic")
	proto.RegisterEnum("pganalyze.collector.BloatLookupMethod", BloatLookupMethod_name, BloatLookupMethod_value)
}

func init() { proto.RegisterFile("bloat_report.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x94, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0xdf, 0xac, 0xef, 0xc6, 0x7a, 0xfa, 0x67, 0x99, 0x3b, 0x46, 0x00, 0xa1, 0x95, 0x4d,
	0xa0, 0x08, 0xa4, 0x5e, 0xc0, 0x27, 0xd8, 0xd6, 0x4d, 0x54, 0xea, 0x06, 0x4a, 0x0b, 0xbb, 0x23,
	0x72, 0xeb, 0xd3, 0x2e, 0xc2, 0x8b, 0xa3, 0xd8, 0xad, 0x5a, 0x2e, 0xf9, 0x9c, 0x7c, 0x15, 0x24,
	0x64, 0x3b, 0xed, 0xda, 0x34, 0xdb, 0x35, 0x97, 0xfd, 0xf9, 0xe7, 0xe7, 0x9c, 0x3e, 0x96, 0x02,
	0x64, 0xc0, 0x05, 0x55, 0x61, 0x8a, 0x89, 0x48, 0x55, 0x2b, 0x49, 0x85, 0x12, 0xa4, 0x91, 0x8c,
	0x69, 0x4c, 0xf9, 0xfc, 0x27, 0xb6, 0x86, 0x82, 0x73, 0x1c, 0x2a, 0x91, 0xbe, 0xa8, 0xca, 0x5b,
	0x9a, 0x22, 0xb3, 0xca, 0xf1, 0xef, 0x12, 0xec, 0x9d, 0xe9, 0x9b, 0x81, 0xb9, 0xd8, 0xa6, 0x8a,
	0x92, 0x1b, 0x68, 0x30, 0xaa, 0xe8, 0x80, 0x4a, 0x0c, 0x53, 0x1c, 0x61, 0x8a, 0xf1, 0x10, 0xa5,
	0x07, 0xcd, 0x92, 0x5f, 0xf9, 0xf0, 0xb6, 0x55, 0x10, 0xda, 0x6a, 0x67, 0x7e, 0xb0, 0xd0, 0x03,
	0xc2, 0xf2, 0x48, 0xea, 0xe0, 0x14, 0x39, 0x55, 0x91, 0x88, 0x57, 0x83, 0x2b, 0x8f, 0x04, 0x07,
	0x99, 0xbf, 0x12, 0x9c, 0xe6, 0x91, 0x24, 0xd7, 0xe0, 0x46, 0x31, 0xc3, 0xd9, 0x6a, 0x6a, 0xd5,
	0xa4, 0x9e, 0x14, 0xa6, 0x76, 0xb4, 0x7c, 0x1f, 0xb9, 0x17, 0xad, 0xfd, 0x96, 0x64, 0x0c, 0xcf,
	0x97, 0x8b, 0xda, 0x5e, 0xa5, 0xa2, 0x2a, 0x92, 0x2a, 0x1a, 0x4a, 0xef, 0xc0, 0x04, 0xbf, 0x7f,
	0x74, 0x5d, 0x53, 0x69, 0x6f, 0x71, 0x27, 0x78, 0x96, 0x16, 0x72, 0x49, 0xbe, 0xc3, 0xa1, 0x5d,
	0x7c, 0x63, 0xca, 0x53, 0x33, 0xc5, 0x7f, 0x78, 0xfd, 0xdc, 0x88, 0x83, 0x68, 0x13, 0xca, 0xe3,
	0x3f, 0x5b, 0x70, 0x58, 0xbc, 0x13, 0x79, 0x0d, 0xd5, 0xe5, 0x7f, 0x8c, 0xd8, 0xcc, 0x73, 0x9a,
	0x8e, 0xbf, 0x1d, 0x54, 0x16, 0xac, 0xc3, 0x66, 0xe4, 0x1b, 0x34, 0xec, 0x5e, 0x5c, 0x88, 0x1f,
	0x93, 0x24, 0xbc, 0x43, 0x75, 0x2b, 0x98, 0xb7, 0xd5, 0x74, 0xfc, 0xfa, 0x03, 0xef, 0x65, 0x86,
	0x74, 0x8d, 0x7e, 0x65, 0xec, 0x60, 0x7f, 0x90, 0x47, 0xe4, 0x08, 0x2a, 0x4a, 0x28, 0xca, 0xc3,
	0xc1, 0x5c, 0xa1, 0xf4, 0x4a, 0x4d, 0xc7, 0x2f, 0x05, 0x60, 0xd0, 0x99, 0x26, 0x5a, 0xb0, 0x83,
	0xad, 0xf0, 0xbf, 0x15, 0x0c, 0xb2, 0x82, 0x0f, 0x2e, 0x8f, 0xa6, 0x18, 0xaa, 0x49, 0xc2, 0x31,
	0xb3, 0xb6, 0x8d, 0x55, 0xd7, 0xbc, 0xaf, 0x71, 0x91, 0x39, 0x14, 0x93, 0x58, 0x79, 0x3b, 0x39,
	0xf3, 0x5c, 0x53, 0x6d, 0x32, 0xa4, 0x6c, 0x2d, 0xf3, 0x89, 0x35, 0x35, 0x5f, 0xcf, 0x5c, 0x31,
	0x6d, 0xe6, 0x6e, 0xce, 0x34, 0x99, 0xc7, 0xbf, 0x4a, 0xd0, 0x28, 0x78, 0x2d, 0xf2, 0x12, 0xca,
	0xf6, 0xdd, 0xef, 0x9b, 0xdf, 0x35, 0xe0, 0xdf, 0xae, 0xfd, 0x0d, 0xd4, 0xa3, 0x58, 0x61, 0x1a,
	0x53, 0x1e, 0x26, 0x74, 0xbc, 0x2c, 0xbd, 0xb6, 0xa0, 0x5f, 0x34, 0x24, 0xaf, 0x00, 0x38, 0xd2,
	0x51, 0xa6, 0xd8, 0xb6, 0xcb, 0x9a, 0xd8, 0xe3, 0x23, 0xa8, 0xe0, 0x5d, 0xa2, 0xe6, 0xd9, 0xb9,
	0xed, 0x18, 0x0c, 0xb2, 0xc2, 0x09, 0xd4, 0x18, 0x72, 0x54, 0xc8, 0x32, 0xc5, 0x96, 0x5b, 0xcd,
	0xa0, 0x95, 0x7c, 0x70, 0xe9, 0x74, 0x1c, 0x9a, 0x41, 0x0c, 0x63, 0x19, 0xa9, 0xb9, 0x57, 0x6e,
	0x3a, 0xbe, 0x13, 0xd4, 0xe9, 0x74, 0xdc, 0x45, 0x3a, 0x6a, 0x5b, 0xfa, 0xee, 0x13, 0xec, 0x6f,
	0xf4, 0x43, 0xf6, 0xa1, 0x76, 0xd1, 0xeb, 0x77, 0xae, 0x4e, 0xfb, 0x17, 0xe1, 0xe5, 0x69, 0xaf,
	0xef, 0xfe, 0xb7, 0x86, 0x7a, 0xdd, 0xcf, 0x37, 0xae, 0x43, 0x6a, 0x50, 0xbe, 0xfc, 0xda, 0xed,
	0x86, 0xbd, 0xf3, 0xd3, 0x6b, 0x77, 0x6b, 0xb0, 0x63, 0x3e, 0x9a, 0x1f, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xb4, 0x55, 0xa0, 0x63, 0x6d, 0x05, 0x00, 0x00,
}
