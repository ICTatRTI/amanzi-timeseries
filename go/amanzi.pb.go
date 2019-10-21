// Code generated by protoc-gen-go. DO NOT EDIT.
// source: amanzi.proto

package rti_amanzi_protobuf

import (
	fmt "fmt"
	_struct "github.com/golang/proto/ptypes/struct"
	timestamp "github.com/golang/proto/ptypes/timestamp"
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

// `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
//  The JSON representation for `NullValue` is JSON `null`.
type NullValue int32

const (
	// Null value.
	NullValue_NULL_VALUE NullValue = 0
)

var NullValue_name = map[int32]string{
	0: "NULL_VALUE",
}

var NullValue_value = map[string]int32{
	"NULL_VALUE": 0,
}

func (x NullValue) String() string {
	return proto.EnumName(NullValue_name, int32(x))
}

func (NullValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{0}
}

type NumberValue struct {
	// Types that are valid to be assigned to Kind:
	//	*NumberValue_DoubleValue
	//	*NumberValue_FloatValue
	//	*NumberValue_Int64Value
	//	*NumberValue_Int32Value
	//	*NumberValue_Uint64Value
	//	*NumberValue_Unit32Value
	//	*NumberValue_NullValue
	Kind                 isNumberValue_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NumberValue) Reset()         { *m = NumberValue{} }
func (m *NumberValue) String() string { return proto.CompactTextString(m) }
func (*NumberValue) ProtoMessage()    {}
func (*NumberValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{0}
}

func (m *NumberValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberValue.Unmarshal(m, b)
}
func (m *NumberValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberValue.Marshal(b, m, deterministic)
}
func (m *NumberValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberValue.Merge(m, src)
}
func (m *NumberValue) XXX_Size() int {
	return xxx_messageInfo_NumberValue.Size(m)
}
func (m *NumberValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberValue.DiscardUnknown(m)
}

var xxx_messageInfo_NumberValue proto.InternalMessageInfo

type isNumberValue_Kind interface {
	isNumberValue_Kind()
}

type NumberValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,1,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type NumberValue_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,2,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type NumberValue_Int64Value struct {
	Int64Value int64 `protobuf:"varint,3,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type NumberValue_Int32Value struct {
	Int32Value int32 `protobuf:"varint,4,opt,name=int32_value,json=int32Value,proto3,oneof"`
}

type NumberValue_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,5,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

type NumberValue_Unit32Value struct {
	Unit32Value uint32 `protobuf:"varint,6,opt,name=unit32_value,json=unit32Value,proto3,oneof"`
}

type NumberValue_NullValue struct {
	NullValue NullValue `protobuf:"varint,7,opt,name=null_value,json=nullValue,proto3,enum=rti.amanzi.protobuf.NullValue,oneof"`
}

func (*NumberValue_DoubleValue) isNumberValue_Kind() {}

func (*NumberValue_FloatValue) isNumberValue_Kind() {}

func (*NumberValue_Int64Value) isNumberValue_Kind() {}

func (*NumberValue_Int32Value) isNumberValue_Kind() {}

func (*NumberValue_Uint64Value) isNumberValue_Kind() {}

func (*NumberValue_Unit32Value) isNumberValue_Kind() {}

func (*NumberValue_NullValue) isNumberValue_Kind() {}

func (m *NumberValue) GetKind() isNumberValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *NumberValue) GetDoubleValue() float64 {
	if x, ok := m.GetKind().(*NumberValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *NumberValue) GetFloatValue() float32 {
	if x, ok := m.GetKind().(*NumberValue_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *NumberValue) GetInt64Value() int64 {
	if x, ok := m.GetKind().(*NumberValue_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *NumberValue) GetInt32Value() int32 {
	if x, ok := m.GetKind().(*NumberValue_Int32Value); ok {
		return x.Int32Value
	}
	return 0
}

func (m *NumberValue) GetUint64Value() uint64 {
	if x, ok := m.GetKind().(*NumberValue_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *NumberValue) GetUnit32Value() uint32 {
	if x, ok := m.GetKind().(*NumberValue_Unit32Value); ok {
		return x.Unit32Value
	}
	return 0
}

func (m *NumberValue) GetNullValue() NullValue {
	if x, ok := m.GetKind().(*NumberValue_NullValue); ok {
		return x.NullValue
	}
	return NullValue_NULL_VALUE
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NumberValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NumberValue_DoubleValue)(nil),
		(*NumberValue_FloatValue)(nil),
		(*NumberValue_Int64Value)(nil),
		(*NumberValue_Int32Value)(nil),
		(*NumberValue_Uint64Value)(nil),
		(*NumberValue_Unit32Value)(nil),
		(*NumberValue_NullValue)(nil),
	}
}

// A Source object holds spatial and origin meta data associated with the timeseries
type Source struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{1}
}

func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (m *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(m, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Source) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Elevation struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	Units                string   `protobuf:"bytes,2,opt,name=units,proto3" json:"units,omitempty"`
	Datum                string   `protobuf:"bytes,3,opt,name=datum,proto3" json:"datum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Elevation) Reset()         { *m = Elevation{} }
func (m *Elevation) String() string { return proto.CompactTextString(m) }
func (*Elevation) ProtoMessage()    {}
func (*Elevation) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{2}
}

func (m *Elevation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Elevation.Unmarshal(m, b)
}
func (m *Elevation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Elevation.Marshal(b, m, deterministic)
}
func (m *Elevation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Elevation.Merge(m, src)
}
func (m *Elevation) XXX_Size() int {
	return xxx_messageInfo_Elevation.Size(m)
}
func (m *Elevation) XXX_DiscardUnknown() {
	xxx_messageInfo_Elevation.DiscardUnknown(m)
}

var xxx_messageInfo_Elevation proto.InternalMessageInfo

func (m *Elevation) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Elevation) GetUnits() string {
	if m != nil {
		return m.Units
	}
	return ""
}

func (m *Elevation) GetDatum() string {
	if m != nil {
		return m.Datum
	}
	return ""
}

type SourceLocation struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string     `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Elevation            *Elevation `protobuf:"bytes,3,opt,name=elevation,proto3" json:"elevation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SourceLocation) Reset()         { *m = SourceLocation{} }
func (m *SourceLocation) String() string { return proto.CompactTextString(m) }
func (*SourceLocation) ProtoMessage()    {}
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{3}
}

func (m *SourceLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceLocation.Unmarshal(m, b)
}
func (m *SourceLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceLocation.Marshal(b, m, deterministic)
}
func (m *SourceLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceLocation.Merge(m, src)
}
func (m *SourceLocation) XXX_Size() int {
	return xxx_messageInfo_SourceLocation.Size(m)
}
func (m *SourceLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceLocation.DiscardUnknown(m)
}

var xxx_messageInfo_SourceLocation proto.InternalMessageInfo

func (m *SourceLocation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SourceLocation) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SourceLocation) GetElevation() *Elevation {
	if m != nil {
		return m.Elevation
	}
	return nil
}

type Parameter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Units                string   `protobuf:"bytes,3,opt,name=units,proto3" json:"units,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{4}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Parameter) GetUnits() string {
	if m != nil {
		return m.Units
	}
	return ""
}

type TimeInfo struct {
	ReferenceTime        *timestamp.Timestamp `protobuf:"bytes,1,opt,name=referenceTime,proto3" json:"referenceTime,omitempty"`
	Start                *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Interval             string               `protobuf:"bytes,4,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeInfo) Reset()         { *m = TimeInfo{} }
func (m *TimeInfo) String() string { return proto.CompactTextString(m) }
func (*TimeInfo) ProtoMessage()    {}
func (*TimeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{5}
}

func (m *TimeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeInfo.Unmarshal(m, b)
}
func (m *TimeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeInfo.Marshal(b, m, deterministic)
}
func (m *TimeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeInfo.Merge(m, src)
}
func (m *TimeInfo) XXX_Size() int {
	return xxx_messageInfo_TimeInfo.Size(m)
}
func (m *TimeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TimeInfo proto.InternalMessageInfo

func (m *TimeInfo) GetReferenceTime() *timestamp.Timestamp {
	if m != nil {
		return m.ReferenceTime
	}
	return nil
}

func (m *TimeInfo) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *TimeInfo) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *TimeInfo) GetInterval() string {
	if m != nil {
		return m.Interval
	}
	return ""
}

type Origin struct {
	ModifiedOn           *timestamp.Timestamp `protobuf:"bytes,1,opt,name=modifiedOn,proto3" json:"modifiedOn,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ReferenceIds         []string             `protobuf:"bytes,3,rep,name=referenceIds,proto3" json:"referenceIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Origin) Reset()         { *m = Origin{} }
func (m *Origin) String() string { return proto.CompactTextString(m) }
func (*Origin) ProtoMessage()    {}
func (*Origin) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{6}
}

func (m *Origin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Origin.Unmarshal(m, b)
}
func (m *Origin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Origin.Marshal(b, m, deterministic)
}
func (m *Origin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Origin.Merge(m, src)
}
func (m *Origin) XXX_Size() int {
	return xxx_messageInfo_Origin.Size(m)
}
func (m *Origin) XXX_DiscardUnknown() {
	xxx_messageInfo_Origin.DiscardUnknown(m)
}

var xxx_messageInfo_Origin proto.InternalMessageInfo

func (m *Origin) GetModifiedOn() *timestamp.Timestamp {
	if m != nil {
		return m.ModifiedOn
	}
	return nil
}

func (m *Origin) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Origin) GetReferenceIds() []string {
	if m != nil {
		return m.ReferenceIds
	}
	return nil
}

type Group struct {
	GroupId              string   `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{7}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *Group) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type TimeSeriesMetaInfo struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string          `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Type                 string          `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Properties           *_struct.Struct `protobuf:"bytes,5,opt,name=properties,proto3" json:"properties,omitempty"`
	Groups               []*Group        `protobuf:"bytes,6,rep,name=groups,proto3" json:"groups,omitempty"`
	Source               *Source         `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	SourceLocation       *SourceLocation `protobuf:"bytes,8,opt,name=sourceLocation,proto3" json:"sourceLocation,omitempty"`
	Parameter            *Parameter      `protobuf:"bytes,9,opt,name=parameter,proto3" json:"parameter,omitempty"`
	TimeInfo             *TimeInfo       `protobuf:"bytes,10,opt,name=timeInfo,proto3" json:"timeInfo,omitempty"`
	Origin               *Origin         `protobuf:"bytes,11,opt,name=origin,proto3" json:"origin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TimeSeriesMetaInfo) Reset()         { *m = TimeSeriesMetaInfo{} }
func (m *TimeSeriesMetaInfo) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesMetaInfo) ProtoMessage()    {}
func (*TimeSeriesMetaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{8}
}

func (m *TimeSeriesMetaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeriesMetaInfo.Unmarshal(m, b)
}
func (m *TimeSeriesMetaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeriesMetaInfo.Marshal(b, m, deterministic)
}
func (m *TimeSeriesMetaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeriesMetaInfo.Merge(m, src)
}
func (m *TimeSeriesMetaInfo) XXX_Size() int {
	return xxx_messageInfo_TimeSeriesMetaInfo.Size(m)
}
func (m *TimeSeriesMetaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeriesMetaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeriesMetaInfo proto.InternalMessageInfo

func (m *TimeSeriesMetaInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TimeSeriesMetaInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimeSeriesMetaInfo) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *TimeSeriesMetaInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TimeSeriesMetaInfo) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *TimeSeriesMetaInfo) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *TimeSeriesMetaInfo) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *TimeSeriesMetaInfo) GetSourceLocation() *SourceLocation {
	if m != nil {
		return m.SourceLocation
	}
	return nil
}

func (m *TimeSeriesMetaInfo) GetParameter() *Parameter {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func (m *TimeSeriesMetaInfo) GetTimeInfo() *TimeInfo {
	if m != nil {
		return m.TimeInfo
	}
	return nil
}

func (m *TimeSeriesMetaInfo) GetOrigin() *Origin {
	if m != nil {
		return m.Origin
	}
	return nil
}

type TimeRecord struct {
	Datetime             *timestamp.Timestamp `protobuf:"bytes,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Value                *NumberValue         `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Qualifiers           []string             `protobuf:"bytes,3,rep,name=qualifiers,proto3" json:"qualifiers,omitempty"`
	Properties           *_struct.Struct      `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeRecord) Reset()         { *m = TimeRecord{} }
func (m *TimeRecord) String() string { return proto.CompactTextString(m) }
func (*TimeRecord) ProtoMessage()    {}
func (*TimeRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{9}
}

func (m *TimeRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRecord.Unmarshal(m, b)
}
func (m *TimeRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRecord.Marshal(b, m, deterministic)
}
func (m *TimeRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRecord.Merge(m, src)
}
func (m *TimeRecord) XXX_Size() int {
	return xxx_messageInfo_TimeRecord.Size(m)
}
func (m *TimeRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRecord.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRecord proto.InternalMessageInfo

func (m *TimeRecord) GetDatetime() *timestamp.Timestamp {
	if m != nil {
		return m.Datetime
	}
	return nil
}

func (m *TimeRecord) GetValue() *NumberValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TimeRecord) GetQualifiers() []string {
	if m != nil {
		return m.Qualifiers
	}
	return nil
}

func (m *TimeRecord) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type TimeSeries struct {
	// a meta object describing the timeseries
	MetaInfo *TimeSeriesMetaInfo `protobuf:"bytes,1,opt,name=metaInfo,proto3" json:"metaInfo,omitempty"`
	// the time , value records of the timeseries
	Data                 []*TimeRecord `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TimeSeries) Reset()         { *m = TimeSeries{} }
func (m *TimeSeries) String() string { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()    {}
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63c41d33ddcceee, []int{10}
}

func (m *TimeSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeries.Unmarshal(m, b)
}
func (m *TimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeries.Marshal(b, m, deterministic)
}
func (m *TimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeries.Merge(m, src)
}
func (m *TimeSeries) XXX_Size() int {
	return xxx_messageInfo_TimeSeries.Size(m)
}
func (m *TimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeries proto.InternalMessageInfo

func (m *TimeSeries) GetMetaInfo() *TimeSeriesMetaInfo {
	if m != nil {
		return m.MetaInfo
	}
	return nil
}

func (m *TimeSeries) GetData() []*TimeRecord {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("rti.amanzi.protobuf.NullValue", NullValue_name, NullValue_value)
	proto.RegisterType((*NumberValue)(nil), "rti.amanzi.protobuf.NumberValue")
	proto.RegisterType((*Source)(nil), "rti.amanzi.protobuf.Source")
	proto.RegisterType((*Elevation)(nil), "rti.amanzi.protobuf.Elevation")
	proto.RegisterType((*SourceLocation)(nil), "rti.amanzi.protobuf.SourceLocation")
	proto.RegisterType((*Parameter)(nil), "rti.amanzi.protobuf.Parameter")
	proto.RegisterType((*TimeInfo)(nil), "rti.amanzi.protobuf.TimeInfo")
	proto.RegisterType((*Origin)(nil), "rti.amanzi.protobuf.Origin")
	proto.RegisterType((*Group)(nil), "rti.amanzi.protobuf.Group")
	proto.RegisterType((*TimeSeriesMetaInfo)(nil), "rti.amanzi.protobuf.TimeSeriesMetaInfo")
	proto.RegisterType((*TimeRecord)(nil), "rti.amanzi.protobuf.TimeRecord")
	proto.RegisterType((*TimeSeries)(nil), "rti.amanzi.protobuf.TimeSeries")
}

func init() { proto.RegisterFile("amanzi.proto", fileDescriptor_e63c41d33ddcceee) }

var fileDescriptor_e63c41d33ddcceee = []byte{
	// 816 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x5e, 0xc7, 0x89, 0x1b, 0x1f, 0x6f, 0x23, 0x34, 0x54, 0xc2, 0xda, 0x42, 0x6b, 0xdc, 0x0b,
	0x2c, 0x84, 0xd2, 0xca, 0x41, 0x8b, 0x40, 0x48, 0x40, 0xab, 0x8a, 0x5d, 0xb1, 0xdd, 0x56, 0xb3,
	0x6d, 0x6f, 0xab, 0x49, 0x3c, 0x89, 0x46, 0xd8, 0x1e, 0x33, 0x1e, 0x47, 0x82, 0x07, 0xe0, 0x82,
	0x87, 0xe1, 0x39, 0x78, 0x01, 0x1e, 0x07, 0x09, 0xcd, 0x8f, 0x1d, 0x67, 0x9b, 0xa4, 0xb9, 0x9b,
	0x73, 0xf2, 0x9d, 0xe3, 0x73, 0xbe, 0xf9, 0xbe, 0x09, 0x9c, 0x92, 0x82, 0x94, 0x7f, 0xb0, 0x69,
	0x25, 0xb8, 0xe4, 0xe8, 0x63, 0x21, 0xd9, 0xb4, 0x9f, 0x99, 0x37, 0xcb, 0xb3, 0x4f, 0x57, 0x9c,
	0xaf, 0x72, 0xfa, 0xb8, 0x4d, 0x3c, 0xae, 0xa5, 0x68, 0x16, 0xd2, 0x00, 0xce, 0x1e, 0xde, 0xfe,
	0x55, 0xb2, 0x82, 0xd6, 0x92, 0x14, 0x95, 0x01, 0xc4, 0x7f, 0x0f, 0x20, 0xb8, 0x6e, 0x8a, 0x39,
	0x15, 0x6f, 0x49, 0xde, 0x50, 0xf4, 0x08, 0x4e, 0x33, 0xde, 0xcc, 0x73, 0xfa, 0x6e, 0xad, 0xe2,
	0xd0, 0x89, 0x9c, 0xc4, 0xb9, 0x38, 0xc1, 0x81, 0xc9, 0x1a, 0xd0, 0xe7, 0x10, 0x2c, 0x73, 0x4e,
	0xa4, 0xc5, 0x0c, 0x22, 0x27, 0x19, 0x5c, 0x9c, 0x60, 0xd0, 0xc9, 0x0e, 0xc2, 0x4a, 0x79, 0xfe,
	0xb5, 0x85, 0xb8, 0x91, 0x93, 0xb8, 0x0a, 0xa2, 0x93, 0x7d, 0xc8, 0x2c, 0xb5, 0x90, 0x61, 0xe4,
	0x24, 0x23, 0x0b, 0x99, 0xa5, 0xdd, 0x34, 0x4d, 0xbf, 0xcd, 0x28, 0x72, 0x92, 0xa1, 0x9a, 0xa6,
	0xe9, 0xf5, 0x51, 0xa0, 0x92, 0x6d, 0x1a, 0x79, 0x91, 0x93, 0xdc, 0xd5, 0x20, 0x9d, 0x35, 0xa0,
	0x1f, 0x00, 0xca, 0x26, 0xcf, 0x2d, 0xe4, 0x4e, 0xe4, 0x24, 0x93, 0xf4, 0xc1, 0x74, 0x07, 0xa1,
	0xd3, 0xeb, 0x26, 0xcf, 0x75, 0xcd, 0xc5, 0x09, 0xf6, 0xcb, 0x36, 0x78, 0xea, 0xc1, 0xf0, 0x57,
	0x56, 0x66, 0xf1, 0x13, 0xf0, 0x6e, 0x78, 0x23, 0x16, 0x14, 0x21, 0x18, 0x96, 0xa4, 0x30, 0x14,
	0xf9, 0x58, 0x9f, 0x55, 0x6e, 0xc1, 0x33, 0x43, 0x89, 0x8f, 0xf5, 0x39, 0x7e, 0x01, 0xfe, 0xf3,
	0x9c, 0xae, 0x89, 0x64, 0xbc, 0x44, 0xf7, 0x60, 0xd4, 0x23, 0x16, 0x9b, 0x40, 0x65, 0xd5, 0xb0,
	0xb5, 0xad, 0x33, 0x81, 0xca, 0x66, 0x44, 0x36, 0x85, 0x66, 0xcf, 0xc7, 0x26, 0x88, 0xd7, 0x30,
	0x31, 0x03, 0x5c, 0xf1, 0x85, 0xe9, 0x79, 0xe4, 0x20, 0xe8, 0x7b, 0xf0, 0x69, 0x3b, 0x88, 0xee,
	0x19, 0xec, 0xa1, 0xa0, 0x1b, 0x17, 0x6f, 0x0a, 0xe2, 0x4b, 0xf0, 0x5f, 0x11, 0x41, 0x0a, 0x2a,
	0xa9, 0x38, 0xfa, 0x93, 0xdd, 0x62, 0x6e, 0x6f, 0xb1, 0xf8, 0x1f, 0x07, 0xc6, 0xaf, 0x59, 0x41,
	0x2f, 0xcb, 0x25, 0x47, 0x3f, 0xc2, 0x5d, 0x41, 0x97, 0x54, 0xd0, 0x72, 0x41, 0x55, 0x52, 0xf7,
	0x0c, 0xd2, 0xb3, 0xa9, 0x91, 0xee, 0x66, 0xaa, 0xd7, 0xad, 0x74, 0xf1, 0x76, 0x01, 0x7a, 0x02,
	0xa3, 0x5a, 0x12, 0x21, 0xf5, 0x97, 0x0f, 0x57, 0x1a, 0x20, 0xfa, 0x0a, 0x5c, 0x5a, 0x66, 0x96,
	0x83, 0x43, 0x78, 0x05, 0x43, 0x67, 0x30, 0x66, 0xa5, 0xa4, 0x62, 0x4d, 0x72, 0xad, 0x52, 0x1f,
	0x77, 0x71, 0xfc, 0x97, 0x03, 0xde, 0x4b, 0xc1, 0x56, 0xac, 0x44, 0xdf, 0x01, 0x14, 0x3c, 0x63,
	0x4b, 0x46, 0xb3, 0x97, 0xe5, 0x11, 0x5b, 0xf4, 0xd0, 0x28, 0x82, 0x20, 0xa3, 0xf5, 0x42, 0xb0,
	0x4a, 0x5f, 0x8e, 0xa1, 0xb0, 0x9f, 0x42, 0x31, 0x9c, 0x76, 0x5b, 0x5f, 0x66, 0x8a, 0x50, 0x37,
	0xf1, 0xf1, 0x56, 0x2e, 0x7e, 0x06, 0xa3, 0x9f, 0x05, 0x6f, 0x2a, 0x14, 0xc2, 0x9d, 0x95, 0x3a,
	0x5c, 0x66, 0xf6, 0x86, 0xda, 0xf0, 0xc3, 0x1f, 0x8a, 0xff, 0x73, 0x01, 0xa9, 0x21, 0x6f, 0xa8,
	0x60, 0xb4, 0x7e, 0x41, 0x25, 0xd1, 0xd7, 0x34, 0x81, 0x01, 0x6b, 0xbb, 0x0d, 0x58, 0xd6, 0x29,
	0x60, 0xb0, 0x43, 0x01, 0x6e, 0x4f, 0x01, 0x08, 0x86, 0xf2, 0xf7, 0x8a, 0x5a, 0xe2, 0xf4, 0x19,
	0x7d, 0x03, 0x50, 0x09, 0x5e, 0x51, 0x21, 0x19, 0xad, 0xb5, 0xa9, 0x83, 0xf4, 0x93, 0xf7, 0x98,
	0xba, 0xd1, 0x0f, 0x19, 0xee, 0x41, 0x51, 0x0a, 0x9e, 0x5e, 0xa4, 0x0e, 0xbd, 0xc8, 0xd5, 0xf4,
	0xee, 0x92, 0xaf, 0xe6, 0x00, 0x5b, 0x24, 0x9a, 0x81, 0x57, 0x6b, 0xbf, 0x68, 0xd7, 0x07, 0xe9,
	0xfd, 0x9d, 0x35, 0xc6, 0x52, 0xd8, 0x42, 0xd1, 0x2f, 0x30, 0xa9, 0xb7, 0x4c, 0x16, 0x8e, 0x75,
	0xf1, 0xa3, 0x03, 0xc5, 0x2d, 0x14, 0xdf, 0x2a, 0x55, 0xbe, 0xab, 0x5a, 0xe7, 0x84, 0xfe, 0x01,
	0xdf, 0x75, 0xfe, 0xc2, 0x9b, 0x02, 0xf4, 0x2d, 0x8c, 0xa5, 0xf5, 0x4a, 0x08, 0xba, 0xf8, 0xb3,
	0x9d, 0xc5, 0xad, 0xa1, 0x70, 0x07, 0x57, 0xab, 0x73, 0xad, 0xcd, 0x30, 0x38, 0xb0, 0xba, 0x91,
	0x2f, 0xb6, 0xd0, 0xf8, 0x5f, 0x07, 0x40, 0xf5, 0xc2, 0x74, 0xc1, 0x45, 0x86, 0xce, 0x61, 0x9c,
	0x11, 0x49, 0xe5, 0x71, 0xce, 0xec, 0xb0, 0xe8, 0xbc, 0x7d, 0xe8, 0x8c, 0x29, 0xa3, 0x3d, 0x6f,
	0x6d, 0xf7, 0xcf, 0xd3, 0x3e, 0x85, 0x0f, 0x00, 0x7e, 0x6b, 0x48, 0xae, 0x8c, 0x21, 0x5a, 0x95,
	0xf7, 0x32, 0xb7, 0xb4, 0x33, 0x3c, 0x5a, 0x3b, 0xf1, 0x9f, 0x76, 0x2f, 0xa3, 0x6b, 0xf4, 0x0c,
	0xc6, 0x85, 0xd5, 0xb6, 0xdd, 0xeb, 0x8b, 0xbd, 0xb4, 0x6e, 0x5b, 0x01, 0x77, 0x85, 0x68, 0x06,
	0xc3, 0x8c, 0x48, 0xa2, 0xc7, 0x0c, 0xd2, 0x87, 0x7b, 0x1b, 0x18, 0x2e, 0xb1, 0x06, 0x7f, 0x79,
	0x1f, 0xfc, 0xee, 0x3f, 0x06, 0x4d, 0x00, 0xae, 0xdf, 0x5c, 0x5d, 0xbd, 0x7b, 0xfb, 0xd3, 0xd5,
	0x9b, 0xe7, 0x1f, 0x9d, 0x3c, 0x0d, 0xe1, 0x1e, 0x17, 0xab, 0xf7, 0x1a, 0xbd, 0x72, 0xe6, 0x9e,
	0x3e, 0xcc, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x02, 0x96, 0x80, 0x06, 0x14, 0x08, 0x00, 0x00,
}