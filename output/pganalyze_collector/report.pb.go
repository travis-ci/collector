// Code generated by protoc-gen-go.
// source: report.proto
// DO NOT EDIT!

package pganalyze_collector

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Report struct {
	ReportRunId string                     `protobuf:"bytes,1,opt,name=report_run_id,json=reportRunId" json:"report_run_id,omitempty"`
	ReportType  string                     `protobuf:"bytes,2,opt,name=report_type,json=reportType" json:"report_type,omitempty"`
	CollectedAt *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=collected_at,json=collectedAt" json:"collected_at,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*Report_BloatReportData
	//	*Report_BuffercacheReportData
	Data isReport_Data `protobuf_oneof:"data"`
}

func (m *Report) Reset()                    { *m = Report{} }
func (m *Report) String() string            { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()               {}
func (*Report) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type isReport_Data interface {
	isReport_Data()
}

type Report_BloatReportData struct {
	BloatReportData *BloatReportData `protobuf:"bytes,10,opt,name=bloat_report_data,json=bloatReportData,oneof"`
}
type Report_BuffercacheReportData struct {
	BuffercacheReportData *BuffercacheReportData `protobuf:"bytes,11,opt,name=buffercache_report_data,json=buffercacheReportData,oneof"`
}

func (*Report_BloatReportData) isReport_Data()       {}
func (*Report_BuffercacheReportData) isReport_Data() {}

func (m *Report) GetData() isReport_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Report) GetReportRunId() string {
	if m != nil {
		return m.ReportRunId
	}
	return ""
}

func (m *Report) GetReportType() string {
	if m != nil {
		return m.ReportType
	}
	return ""
}

func (m *Report) GetCollectedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CollectedAt
	}
	return nil
}

func (m *Report) GetBloatReportData() *BloatReportData {
	if x, ok := m.GetData().(*Report_BloatReportData); ok {
		return x.BloatReportData
	}
	return nil
}

func (m *Report) GetBuffercacheReportData() *BuffercacheReportData {
	if x, ok := m.GetData().(*Report_BuffercacheReportData); ok {
		return x.BuffercacheReportData
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Report) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Report_OneofMarshaler, _Report_OneofUnmarshaler, _Report_OneofSizer, []interface{}{
		(*Report_BloatReportData)(nil),
		(*Report_BuffercacheReportData)(nil),
	}
}

func _Report_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Report)
	// data
	switch x := m.Data.(type) {
	case *Report_BloatReportData:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BloatReportData); err != nil {
			return err
		}
	case *Report_BuffercacheReportData:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BuffercacheReportData); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Report.Data has unexpected type %T", x)
	}
	return nil
}

func _Report_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Report)
	switch tag {
	case 10: // data.bloat_report_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BloatReportData)
		err := b.DecodeMessage(msg)
		m.Data = &Report_BloatReportData{msg}
		return true, err
	case 11: // data.buffercache_report_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuffercacheReportData)
		err := b.DecodeMessage(msg)
		m.Data = &Report_BuffercacheReportData{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Report_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Report)
	// data
	switch x := m.Data.(type) {
	case *Report_BloatReportData:
		s := proto.Size(x.BloatReportData)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Report_BuffercacheReportData:
		s := proto.Size(x.BuffercacheReportData)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Report)(nil), "pganalyze.collector.Report")
}

func init() { proto.RegisterFile("report.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4f, 0xfb, 0x30,
	0x14, 0xc4, 0xff, 0xe9, 0x1f, 0x45, 0xc2, 0x29, 0x42, 0x18, 0x21, 0xa2, 0x2c, 0xad, 0x2a, 0x86,
	0x8a, 0xc1, 0x95, 0x60, 0x66, 0xa0, 0x62, 0x80, 0xd5, 0xea, 0x6e, 0x3d, 0x27, 0x2f, 0xa1, 0x92,
	0x1b, 0x5b, 0xe6, 0x65, 0x08, 0x1f, 0x92, 0xcf, 0x84, 0x6a, 0xb7, 0xd0, 0x88, 0x8c, 0xbe, 0xbb,
	0xf7, 0xd3, 0x9d, 0xd9, 0xd4, 0xa3, 0xb3, 0x9e, 0x84, 0xf3, 0x96, 0x2c, 0xbf, 0x76, 0x0d, 0xb4,
	0x60, 0xfa, 0x4f, 0x14, 0xa5, 0x35, 0x06, 0x4b, 0xb2, 0xbe, 0x98, 0x35, 0xd6, 0x36, 0x06, 0x57,
	0x21, 0xa2, 0xbb, 0x7a, 0x45, 0xdb, 0x1d, 0x7e, 0x10, 0xec, 0x5c, 0xbc, 0x2a, 0xb8, 0x36, 0x16,
	0x48, 0x9d, 0x92, 0x8a, 0x5c, 0x77, 0x75, 0x8d, 0xbe, 0x84, 0xf2, 0x1d, 0x07, 0xce, 0xe2, 0x6b,
	0xc2, 0x52, 0x19, 0x04, 0xbe, 0x60, 0x17, 0xd1, 0x52, 0xbe, 0x6b, 0xd5, 0xb6, 0xca, 0x93, 0x79,
	0xb2, 0x3c, 0x97, 0x59, 0x14, 0x65, 0xd7, 0xbe, 0x55, 0x7c, 0xc6, 0x0e, 0x4f, 0x45, 0xbd, 0xc3,
	0x7c, 0x12, 0x12, 0x2c, 0x4a, 0x9b, 0xde, 0x21, 0x7f, 0x62, 0xd3, 0x43, 0x57, 0xac, 0x14, 0x50,
	0xfe, 0x7f, 0x9e, 0x2c, 0xb3, 0x87, 0x42, 0xc4, 0xd6, 0xe2, 0xd8, 0x5a, 0x6c, 0x8e, 0xad, 0x65,
	0xf6, 0x93, 0x7f, 0x26, 0x2e, 0xd9, 0xd5, 0x69, 0x7d, 0x55, 0x01, 0x41, 0xce, 0x02, 0xe3, 0x4e,
	0x8c, 0x7c, 0x87, 0x58, 0xef, 0xd3, 0x71, 0xc0, 0x0b, 0x10, 0xbc, 0xfe, 0x93, 0x97, 0x7a, 0x28,
	0xf1, 0x8a, 0xdd, 0xfe, 0x9d, 0x1f, 0xc9, 0x59, 0x20, 0xdf, 0x8f, 0x93, 0x7f, 0x6f, 0x06, 0xfc,
	0x1b, 0x3d, 0x66, 0xac, 0x53, 0x76, 0xb6, 0x47, 0xea, 0x34, 0x4c, 0x7c, 0xfc, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x09, 0x6e, 0xaf, 0xe4, 0xcb, 0x01, 0x00, 0x00,
}
