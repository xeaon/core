// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bid.proto

/*
Package sonm is a generated protocol buffer package.

It is generated from these files:
	bid.proto
	capabilities.proto
	hub.proto
	insonmnia.proto
	marketplace.proto
	miner.proto

It has these top-level messages:
	Geo
	Resources
	Slot
	Order
	Capabilities
	CPUDevice
	RAMDevice
	GPUDevice
	ListRequest
	ListReply
	HubInfoRequest
	TaskRequirements
	HubStartTaskRequest
	HubStartTaskReply
	HubStatusMapRequest
	HubStatusRequest
	HubStatusReply
	DealRequest
	GetMinerPropertiesRequest
	GetMinerPropertiesReply
	SetMinerPropertiesRequest
	GetSlotsRequest
	GetSlotsReply
	AddSlotRequest
	RemoveSlotRequest
	PingRequest
	PingReply
	CPUUsage
	MemoryUsage
	NetworkUsage
	ResourceUsage
	InfoReply
	StopTaskRequest
	StopTaskReply
	TaskStatusRequest
	TaskStatusReply
	StatusMapReply
	ContainerRestartPolicy
	TaskLogsRequest
	TaskLogsChunk
	EmptyReply
	DiscoverHubRequest
	TaskResourceRequirements
	Empty
	Timestamp
	GetOrdersRequest
	GetOrdersReply
	GetOrderRequest
	MinerInfoRequest
	MinerHandshakeRequest
	MinerHandshakeReply
	MinerStartRequest
	MinerStartReply
	TaskInfo
	MinerStatusMapRequest
*/
package sonm

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

type OrderType int32

const (
	OrderType_ANY OrderType = 0
	OrderType_BID OrderType = 1
	OrderType_ASK OrderType = 2
)

var OrderType_name = map[int32]string{
	0: "ANY",
	1: "BID",
	2: "ASK",
}
var OrderType_value = map[string]int32{
	"ANY": 0,
	"BID": 1,
	"ASK": 2,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}
func (OrderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NetworkType int32

const (
	NetworkType_NO_NETWORK NetworkType = 0
	NetworkType_OUTBOUND   NetworkType = 1
	NetworkType_INCOMING   NetworkType = 2
)

var NetworkType_name = map[int32]string{
	0: "NO_NETWORK",
	1: "OUTBOUND",
	2: "INCOMING",
}
var NetworkType_value = map[string]int32{
	"NO_NETWORK": 0,
	"OUTBOUND":   1,
	"INCOMING":   2,
}

func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}
func (NetworkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Geo represent GeoIP results for node
type Geo struct {
	Country string  `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
	City    string  `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	Lat     float32 `protobuf:"fixed32,3,opt,name=lat" json:"lat,omitempty"`
	Lon     float32 `protobuf:"fixed32,4,opt,name=lon" json:"lon,omitempty"`
}

func (m *Geo) Reset()                    { *m = Geo{} }
func (m *Geo) String() string            { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()               {}
func (*Geo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Geo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Geo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Geo) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Geo) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Resources struct {
	// CPU core count
	CpuCores uint64 `protobuf:"varint,1,opt,name=cpuCores" json:"cpuCores,omitempty"`
	// RAM, in bytes
	RamBytes uint64 `protobuf:"varint,2,opt,name=ramBytes" json:"ramBytes,omitempty"`
	// GPU devices count
	GpuCount uint64 `protobuf:"varint,3,opt,name=gpuCount" json:"gpuCount,omitempty"`
	// todo: discuss
	// storage volume, in Megabytes
	Storage uint64 `protobuf:"varint,4,opt,name=storage" json:"storage,omitempty"`
	// Inbound network traffic (the higher value), in bytes
	NetTrafficIn uint64 `protobuf:"varint,5,opt,name=netTrafficIn" json:"netTrafficIn,omitempty"`
	// Outbound network traffic (the higher value), in bytes
	NetTrafficOut uint64 `protobuf:"varint,6,opt,name=netTrafficOut" json:"netTrafficOut,omitempty"`
	// Allowed network connections
	NetworkType NetworkType `protobuf:"varint,7,opt,name=networkType,enum=sonm.NetworkType" json:"networkType,omitempty"`
	// Other properties/benchmarks. The higher means better
	Props map[string]string `protobuf:"bytes,8,rep,name=props" json:"props,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Resources) Reset()                    { *m = Resources{} }
func (m *Resources) String() string            { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()               {}
func (*Resources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Resources) GetCpuCores() uint64 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *Resources) GetRamBytes() uint64 {
	if m != nil {
		return m.RamBytes
	}
	return 0
}

func (m *Resources) GetGpuCount() uint64 {
	if m != nil {
		return m.GpuCount
	}
	return 0
}

func (m *Resources) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *Resources) GetNetTrafficIn() uint64 {
	if m != nil {
		return m.NetTrafficIn
	}
	return 0
}

func (m *Resources) GetNetTrafficOut() uint64 {
	if m != nil {
		return m.NetTrafficOut
	}
	return 0
}

func (m *Resources) GetNetworkType() NetworkType {
	if m != nil {
		return m.NetworkType
	}
	return NetworkType_NO_NETWORK
}

func (m *Resources) GetProps() map[string]string {
	if m != nil {
		return m.Props
	}
	return nil
}

type Slot struct {
	// Virtual computer start of life (hour grained).
	StartTime *Timestamp `protobuf:"bytes,1,opt,name=startTime" json:"startTime,omitempty"`
	// Virtual computer end of life (hour grained).
	EndTime *Timestamp `protobuf:"bytes,2,opt,name=endTime" json:"endTime,omitempty"`
	// Buyer’s rating. Got from Buyer’s profile for BID orders rating_supplier.
	BuyerRating int64 `protobuf:"varint,3,opt,name=buyerRating" json:"buyerRating,omitempty"`
	// Supplier’s rating. Got from Supplier’s profile for ASK orders.
	SupplierRating int64 `protobuf:"varint,4,opt,name=supplierRating" json:"supplierRating,omitempty"`
	// Geo represent Worker's position
	Geo *Geo `protobuf:"bytes,5,opt,name=geo" json:"geo,omitempty"`
	// Hardware resources requirements
	Resources *Resources `protobuf:"bytes,6,opt,name=resources" json:"resources,omitempty"`
}

func (m *Slot) Reset()                    { *m = Slot{} }
func (m *Slot) String() string            { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()               {}
func (*Slot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Slot) GetStartTime() *Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Slot) GetEndTime() *Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Slot) GetBuyerRating() int64 {
	if m != nil {
		return m.BuyerRating
	}
	return 0
}

func (m *Slot) GetSupplierRating() int64 {
	if m != nil {
		return m.SupplierRating
	}
	return 0
}

func (m *Slot) GetGeo() *Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (m *Slot) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

type Order struct {
	// Order ID, UUIDv4
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Buyer's EtherumID
	ByuerID string `protobuf:"bytes,2,opt,name=byuerID" json:"byuerID,omitempty"`
	// Supplier's is EtherumID
	SupplierID string `protobuf:"bytes,3,opt,name=supplierID" json:"supplierID,omitempty"`
	// Order price
	Price int64 `protobuf:"varint,4,opt,name=price" json:"price,omitempty"`
	// Order type (Bid or Ask)
	OrderType OrderType `protobuf:"varint,5,opt,name=orderType,enum=sonm.OrderType" json:"orderType,omitempty"`
	// Slot describe resource requiements
	Slot *Slot `protobuf:"bytes,6,opt,name=slot" json:"slot,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetByuerID() string {
	if m != nil {
		return m.ByuerID
	}
	return ""
}

func (m *Order) GetSupplierID() string {
	if m != nil {
		return m.SupplierID
	}
	return ""
}

func (m *Order) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_ANY
}

func (m *Order) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

func init() {
	proto.RegisterType((*Geo)(nil), "sonm.Geo")
	proto.RegisterType((*Resources)(nil), "sonm.Resources")
	proto.RegisterType((*Slot)(nil), "sonm.Slot")
	proto.RegisterType((*Order)(nil), "sonm.Order")
	proto.RegisterEnum("sonm.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("sonm.NetworkType", NetworkType_name, NetworkType_value)
}

func init() { proto.RegisterFile("bid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0x8d, 0x7e, 0x1c, 0x5b, 0x57, 0xf9, 0x1c, 0x7d, 0x97, 0x2e, 0x84, 0x0b, 0xc1, 0x98, 0x12,
	0xdc, 0x40, 0x4d, 0x71, 0x36, 0x69, 0x77, 0x4d, 0x1c, 0x8c, 0x08, 0x95, 0xca, 0xc4, 0x21, 0x74,
	0x55, 0x64, 0x79, 0x62, 0x44, 0xec, 0x19, 0x31, 0x1a, 0xb5, 0xe8, 0xc5, 0xfa, 0x44, 0x7d, 0x86,
	0xae, 0xcb, 0x8c, 0xfe, 0x9c, 0xd2, 0x95, 0xe6, 0x9c, 0x7b, 0x86, 0x73, 0xe7, 0xdc, 0x2b, 0x70,
	0xd6, 0xe9, 0x66, 0x96, 0x09, 0x2e, 0x39, 0xda, 0x39, 0x67, 0xfb, 0xd1, 0x69, 0xca, 0xd4, 0x97,
	0xa5, 0x71, 0x45, 0x4f, 0x1e, 0xc1, 0x5a, 0x52, 0x8e, 0x3e, 0xf4, 0x13, 0x5e, 0x30, 0x29, 0x4a,
	0xdf, 0x18, 0x1b, 0x53, 0x87, 0x34, 0x10, 0x11, 0xec, 0x24, 0x95, 0xa5, 0x6f, 0x6a, 0x5a, 0x9f,
	0xd1, 0x03, 0x6b, 0x17, 0x4b, 0xdf, 0x1a, 0x1b, 0x53, 0x93, 0xa8, 0xa3, 0x66, 0x38, 0xf3, 0xed,
	0x9a, 0xe1, 0x6c, 0xf2, 0xcb, 0x04, 0x87, 0xd0, 0x9c, 0x17, 0x22, 0xa1, 0x39, 0x8e, 0x60, 0x90,
	0x64, 0xc5, 0x0d, 0x17, 0x34, 0xd7, 0x06, 0x36, 0x69, 0xb1, 0xaa, 0x89, 0x78, 0x7f, 0x5d, 0x4a,
	0x9a, 0x6b, 0x17, 0x9b, 0xb4, 0x58, 0xd5, 0xb6, 0x4a, 0x57, 0xb0, 0xca, 0xce, 0x26, 0x2d, 0x56,
	0x3d, 0xe7, 0x92, 0x8b, 0x78, 0x4b, 0xb5, 0xaf, 0x4d, 0x1a, 0x88, 0x13, 0x38, 0x61, 0x54, 0xae,
	0x44, 0xfc, 0xf4, 0x94, 0x26, 0x01, 0xf3, 0x7b, 0xba, 0xfc, 0x82, 0xc3, 0x37, 0xf0, 0x5f, 0x87,
	0xa3, 0x42, 0xfa, 0xc7, 0x5a, 0xf4, 0x92, 0xc4, 0x4b, 0x70, 0x19, 0x95, 0x3f, 0xb8, 0x78, 0x5e,
	0x95, 0x19, 0xf5, 0xfb, 0x63, 0x63, 0x3a, 0x9c, 0xff, 0x3f, 0x53, 0x19, 0xce, 0xc2, 0xae, 0x40,
	0x0e, 0x55, 0xf8, 0x1e, 0x7a, 0x99, 0xe0, 0x59, 0xee, 0x0f, 0xc6, 0xd6, 0xd4, 0x9d, 0x8f, 0x2a,
	0x79, 0x1b, 0xc6, 0xec, 0x8b, 0x2a, 0xde, 0xaa, 0x74, 0x49, 0x25, 0x1c, 0x5d, 0x01, 0x74, 0xa4,
	0x0a, 0xf3, 0x99, 0x36, 0x83, 0x50, 0x47, 0x7c, 0x05, 0xbd, 0xef, 0xf1, 0xae, 0xa0, 0xf5, 0x14,
	0x2a, 0xf0, 0xd1, 0xbc, 0x32, 0x26, 0xbf, 0x0d, 0xb0, 0xef, 0x77, 0x5c, 0xe2, 0x3b, 0x70, 0x72,
	0x19, 0x0b, 0xb9, 0x4a, 0xf7, 0x54, 0x5f, 0x75, 0xe7, 0xa7, 0x95, 0xb1, 0x62, 0x72, 0x19, 0xef,
	0x33, 0xd2, 0x29, 0xf0, 0x2d, 0xf4, 0x29, 0xdb, 0x68, 0xb1, 0xf9, 0x6f, 0x71, 0x53, 0xc7, 0x31,
	0xb8, 0xeb, 0xa2, 0xa4, 0x82, 0xc4, 0x32, 0x65, 0x5b, 0x3d, 0x06, 0x8b, 0x1c, 0x52, 0x78, 0x0e,
	0xc3, 0xbc, 0xc8, 0xb2, 0x5d, 0xda, 0x8a, 0x6c, 0x2d, 0xfa, 0x8b, 0xc5, 0xd7, 0x60, 0x6d, 0x29,
	0xd7, 0xe3, 0x70, 0xe7, 0x4e, 0x65, 0xb8, 0xa4, 0x9c, 0x28, 0x56, 0x3d, 0x40, 0x34, 0x11, 0xe9,
	0x61, 0xb4, 0x3d, 0xb5, 0xc9, 0x91, 0x4e, 0x31, 0xf9, 0x69, 0x40, 0x2f, 0x12, 0x1b, 0x2a, 0x70,
	0x08, 0x66, 0xba, 0xa9, 0xd3, 0x32, 0xd3, 0x8d, 0xda, 0x8b, 0x75, 0x59, 0x50, 0x11, 0x2c, 0xea,
	0xb8, 0x1a, 0x88, 0x67, 0x00, 0x4d, 0x47, 0xc1, 0x42, 0x3f, 0xc4, 0x21, 0x07, 0x8c, 0x8a, 0x39,
	0x13, 0x69, 0x42, 0xeb, 0xf6, 0x2b, 0xa0, 0x1a, 0xe3, 0xca, 0x48, 0x6f, 0x40, 0x4f, 0x6f, 0x40,
	0xdd, 0x58, 0xd4, 0xd0, 0xa4, 0x53, 0xe0, 0x19, 0xd8, 0xf9, 0x8e, 0xcb, 0xfa, 0x09, 0x50, 0x29,
	0xd5, 0x88, 0x88, 0xe6, 0x2f, 0xce, 0xc1, 0x69, 0xef, 0x61, 0x1f, 0xac, 0x4f, 0xe1, 0x57, 0xef,
	0x48, 0x1d, 0xae, 0x83, 0x85, 0x67, 0x68, 0xe6, 0xfe, 0xce, 0x33, 0x2f, 0x3e, 0x80, 0x7b, 0xb0,
	0x61, 0x38, 0x04, 0x08, 0xa3, 0x6f, 0xe1, 0xed, 0xea, 0x31, 0x22, 0x77, 0xde, 0x11, 0x9e, 0xc0,
	0x20, 0x7a, 0x58, 0x5d, 0x47, 0x0f, 0xa1, 0xba, 0x75, 0x02, 0x83, 0x20, 0xbc, 0x89, 0x3e, 0x07,
	0xe1, 0xd2, 0x33, 0xd7, 0xc7, 0xfa, 0xdf, 0xbe, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xdd,
	0xa0, 0xcf, 0xff, 0x03, 0x00, 0x00,
}
