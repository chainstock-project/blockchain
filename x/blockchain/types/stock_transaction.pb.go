// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blockchain/stock_transaction.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HoldingStock struct {
	Code          string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Count         int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	PurchasAmount int32  `protobuf:"varint,3,opt,name=purchas_amount,json=purchasAmount,proto3" json:"purchas_amount,omitempty"`
}

func (m *HoldingStock) Reset()         { *m = HoldingStock{} }
func (m *HoldingStock) String() string { return proto.CompactTextString(m) }
func (*HoldingStock) ProtoMessage()    {}
func (*HoldingStock) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c7f749cf77905bb, []int{0}
}
func (m *HoldingStock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HoldingStock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HoldingStock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HoldingStock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HoldingStock.Merge(m, src)
}
func (m *HoldingStock) XXX_Size() int {
	return m.Size()
}
func (m *HoldingStock) XXX_DiscardUnknown() {
	xxx_messageInfo_HoldingStock.DiscardUnknown(m)
}

var xxx_messageInfo_HoldingStock proto.InternalMessageInfo

func (m *HoldingStock) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *HoldingStock) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *HoldingStock) GetPurchasAmount() int32 {
	if m != nil {
		return m.PurchasAmount
	}
	return 0
}

type StockTransaction struct {
	Creator       string          `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	HoldingStocks []*HoldingStock `protobuf:"bytes,2,rep,name=holding_stocks,json=holdingStocks,proto3" json:"holding_stocks,omitempty"`
}

func (m *StockTransaction) Reset()         { *m = StockTransaction{} }
func (m *StockTransaction) String() string { return proto.CompactTextString(m) }
func (*StockTransaction) ProtoMessage()    {}
func (*StockTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c7f749cf77905bb, []int{1}
}
func (m *StockTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StockTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StockTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StockTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockTransaction.Merge(m, src)
}
func (m *StockTransaction) XXX_Size() int {
	return m.Size()
}
func (m *StockTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_StockTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_StockTransaction proto.InternalMessageInfo

func (m *StockTransaction) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *StockTransaction) GetHoldingStocks() []*HoldingStock {
	if m != nil {
		return m.HoldingStocks
	}
	return nil
}

type StockRecord struct {
	Code       string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Count      int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Amount     int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Date       string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	RecordType string `protobuf:"bytes,5,opt,name=record_type,json=recordType,proto3" json:"record_type,omitempty"`
}

func (m *StockRecord) Reset()         { *m = StockRecord{} }
func (m *StockRecord) String() string { return proto.CompactTextString(m) }
func (*StockRecord) ProtoMessage()    {}
func (*StockRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c7f749cf77905bb, []int{2}
}
func (m *StockRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StockRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StockRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StockRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockRecord.Merge(m, src)
}
func (m *StockRecord) XXX_Size() int {
	return m.Size()
}
func (m *StockRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_StockRecord.DiscardUnknown(m)
}

var xxx_messageInfo_StockRecord proto.InternalMessageInfo

func (m *StockRecord) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *StockRecord) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *StockRecord) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *StockRecord) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *StockRecord) GetRecordType() string {
	if m != nil {
		return m.RecordType
	}
	return ""
}

type StockTransactionRecord struct {
	Creator      string         `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	StockRecords []*StockRecord `protobuf:"bytes,2,rep,name=stock_records,json=stockRecords,proto3" json:"stock_records,omitempty"`
}

func (m *StockTransactionRecord) Reset()         { *m = StockTransactionRecord{} }
func (m *StockTransactionRecord) String() string { return proto.CompactTextString(m) }
func (*StockTransactionRecord) ProtoMessage()    {}
func (*StockTransactionRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c7f749cf77905bb, []int{3}
}
func (m *StockTransactionRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StockTransactionRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StockTransactionRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StockTransactionRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockTransactionRecord.Merge(m, src)
}
func (m *StockTransactionRecord) XXX_Size() int {
	return m.Size()
}
func (m *StockTransactionRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_StockTransactionRecord.DiscardUnknown(m)
}

var xxx_messageInfo_StockTransactionRecord proto.InternalMessageInfo

func (m *StockTransactionRecord) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *StockTransactionRecord) GetStockRecords() []*StockRecord {
	if m != nil {
		return m.StockRecords
	}
	return nil
}

func init() {
	proto.RegisterType((*HoldingStock)(nil), "chainstockproject.blockchain.blockchain.HoldingStock")
	proto.RegisterType((*StockTransaction)(nil), "chainstockproject.blockchain.blockchain.StockTransaction")
	proto.RegisterType((*StockRecord)(nil), "chainstockproject.blockchain.blockchain.StockRecord")
	proto.RegisterType((*StockTransactionRecord)(nil), "chainstockproject.blockchain.blockchain.StockTransactionRecord")
}

func init() {
	proto.RegisterFile("blockchain/stock_transaction.proto", fileDescriptor_5c7f749cf77905bb)
}

var fileDescriptor_5c7f749cf77905bb = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4b, 0x4b, 0xf3, 0x40,
	0x14, 0xed, 0xf4, 0xf5, 0xf1, 0xdd, 0x3e, 0x90, 0x41, 0x4a, 0x56, 0xb1, 0x04, 0xc4, 0x6e, 0x4c,
	0xc1, 0xc7, 0xca, 0x95, 0xae, 0x5c, 0xc7, 0xba, 0x50, 0x84, 0x30, 0x9d, 0x0c, 0x4d, 0x6c, 0x9b,
	0x09, 0x33, 0x53, 0xb0, 0x7b, 0x57, 0x82, 0xe0, 0xcf, 0x72, 0xd9, 0xa5, 0x4b, 0x69, 0xff, 0x88,
	0x64, 0x26, 0x25, 0xa3, 0x20, 0xd4, 0xdd, 0x3d, 0xf7, 0x24, 0xe7, 0xde, 0x73, 0xe6, 0x82, 0x37,
	0x9e, 0x71, 0x3a, 0xa5, 0x31, 0x49, 0xd2, 0xa1, 0x54, 0x9c, 0x4e, 0x43, 0x25, 0x48, 0x2a, 0x09,
	0x55, 0x09, 0x4f, 0xfd, 0x4c, 0x70, 0xc5, 0xf1, 0x91, 0xa6, 0x35, 0x9b, 0x09, 0xfe, 0xc8, 0xa8,
	0xf2, 0xcb, 0xbf, 0xac, 0xd2, 0x0b, 0xa1, 0x7d, 0xcd, 0x67, 0x51, 0x92, 0x4e, 0x6e, 0xf2, 0x8f,
	0x31, 0x86, 0x3a, 0xe5, 0x11, 0x73, 0x50, 0x1f, 0x0d, 0xfe, 0x07, 0xba, 0xc6, 0xfb, 0xd0, 0xa0,
	0x7c, 0x91, 0x2a, 0xa7, 0xda, 0x47, 0x83, 0x46, 0x60, 0x00, 0x3e, 0x84, 0x6e, 0xb6, 0x10, 0x34,
	0x26, 0x32, 0x24, 0x73, 0x4d, 0xd7, 0x34, 0xdd, 0x29, 0xba, 0x97, 0xba, 0xe9, 0xbd, 0x20, 0xd8,
	0xd3, 0xd2, 0xa3, 0x72, 0x49, 0xec, 0xc0, 0x3f, 0x2a, 0x18, 0x51, 0x5c, 0x14, 0x83, 0xb6, 0x10,
	0x3f, 0x40, 0x37, 0x36, 0xfb, 0x84, 0x7a, 0x7b, 0xe9, 0x54, 0xfb, 0xb5, 0x41, 0xeb, 0xe4, 0xdc,
	0xdf, 0xd1, 0x91, 0x6f, 0xdb, 0x09, 0x3a, 0xb1, 0x85, 0xa4, 0xf7, 0x8c, 0xa0, 0x65, 0x08, 0x46,
	0xb9, 0x88, 0xfe, 0xe0, 0xb6, 0x07, 0xcd, 0x6f, 0x2e, 0x0b, 0x94, 0x2b, 0x44, 0x44, 0x31, 0xa7,
	0x6e, 0x14, 0xf2, 0x1a, 0x1f, 0x40, 0x4b, 0x68, 0xfd, 0x50, 0x2d, 0x33, 0xe6, 0x34, 0x34, 0x05,
	0xa6, 0x35, 0x5a, 0x66, 0xcc, 0x7b, 0x45, 0xd0, 0xfb, 0x99, 0x49, 0xb1, 0xd1, 0xef, 0xc9, 0xdc,
	0x41, 0xc7, 0xbc, 0xb6, 0x11, 0xda, 0x06, 0x73, 0xb6, 0x73, 0x30, 0x96, 0xf1, 0xa0, 0x2d, 0x4b,
	0x20, 0xaf, 0x6e, 0xdf, 0xd7, 0x2e, 0x5a, 0xad, 0x5d, 0xf4, 0xb9, 0x76, 0xd1, 0xdb, 0xc6, 0xad,
	0xac, 0x36, 0x6e, 0xe5, 0x63, 0xe3, 0x56, 0xee, 0x2f, 0x26, 0x89, 0x8a, 0x17, 0x63, 0x9f, 0xf2,
	0xf9, 0xb0, 0x9c, 0x73, 0x5c, 0x0c, 0x1a, 0x5a, 0x97, 0xf8, 0x64, 0x83, 0xdc, 0xb7, 0x1c, 0x37,
	0xf5, 0x2d, 0x9e, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x07, 0x49, 0xee, 0x97, 0xb1, 0x02, 0x00,
	0x00,
}

func (m *HoldingStock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HoldingStock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HoldingStock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PurchasAmount != 0 {
		i = encodeVarintStockTransaction(dAtA, i, uint64(m.PurchasAmount))
		i--
		dAtA[i] = 0x18
	}
	if m.Count != 0 {
		i = encodeVarintStockTransaction(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintStockTransaction(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StockTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StockTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StockTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HoldingStocks) > 0 {
		for iNdEx := len(m.HoldingStocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HoldingStocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStockTransaction(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintStockTransaction(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StockRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StockRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StockRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RecordType) > 0 {
		i -= len(m.RecordType)
		copy(dAtA[i:], m.RecordType)
		i = encodeVarintStockTransaction(dAtA, i, uint64(len(m.RecordType)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintStockTransaction(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x22
	}
	if m.Amount != 0 {
		i = encodeVarintStockTransaction(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if m.Count != 0 {
		i = encodeVarintStockTransaction(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintStockTransaction(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StockTransactionRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StockTransactionRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StockTransactionRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StockRecords) > 0 {
		for iNdEx := len(m.StockRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StockRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStockTransaction(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintStockTransaction(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStockTransaction(dAtA []byte, offset int, v uint64) int {
	offset -= sovStockTransaction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HoldingStock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovStockTransaction(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovStockTransaction(uint64(m.Count))
	}
	if m.PurchasAmount != 0 {
		n += 1 + sovStockTransaction(uint64(m.PurchasAmount))
	}
	return n
}

func (m *StockTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovStockTransaction(uint64(l))
	}
	if len(m.HoldingStocks) > 0 {
		for _, e := range m.HoldingStocks {
			l = e.Size()
			n += 1 + l + sovStockTransaction(uint64(l))
		}
	}
	return n
}

func (m *StockRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovStockTransaction(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovStockTransaction(uint64(m.Count))
	}
	if m.Amount != 0 {
		n += 1 + sovStockTransaction(uint64(m.Amount))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovStockTransaction(uint64(l))
	}
	l = len(m.RecordType)
	if l > 0 {
		n += 1 + l + sovStockTransaction(uint64(l))
	}
	return n
}

func (m *StockTransactionRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovStockTransaction(uint64(l))
	}
	if len(m.StockRecords) > 0 {
		for _, e := range m.StockRecords {
			l = e.Size()
			n += 1 + l + sovStockTransaction(uint64(l))
		}
	}
	return n
}

func sovStockTransaction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStockTransaction(x uint64) (n int) {
	return sovStockTransaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HoldingStock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStockTransaction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HoldingStock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HoldingStock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStockTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PurchasAmount", wireType)
			}
			m.PurchasAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PurchasAmount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStockTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StockTransaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStockTransaction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StockTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StockTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStockTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HoldingStocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStockTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HoldingStocks = append(m.HoldingStocks, &HoldingStock{})
			if err := m.HoldingStocks[len(m.HoldingStocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStockTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StockRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStockTransaction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StockRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StockRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStockTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStockTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStockTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecordType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStockTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StockTransactionRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStockTransaction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StockTransactionRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StockTransactionRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStockTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StockRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStockTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StockRecords = append(m.StockRecords, &StockRecord{})
			if err := m.StockRecords[len(m.StockRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStockTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStockTransaction
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStockTransaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStockTransaction
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStockTransaction
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStockTransaction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStockTransaction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStockTransaction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStockTransaction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStockTransaction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStockTransaction = fmt.Errorf("proto: unexpected end of group")
)
