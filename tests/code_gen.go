// Generated by github.com/davyxu/protoplus
// DO NOT EDIT!
package tests

import (
	"github.com/davyxu/protoplus/proto"
	"unsafe"
)

var (
	_ *proto.Buffer
	_ unsafe.Pointer
)

type MyEnum int32

const (
	MyEnum_Zero MyEnum = 0
	MyEnum_One  MyEnum = 1
	MyEnum_Two  MyEnum = 2
)

var (
	MyEnumMapperValueByName = map[string]int32{
		"Zero": 0,
		"One":  1,
		"Two":  2,
	}

	MyEnumMapperNameByValue = map[int32]string{
		0: "Zero",
		1: "One",
		2: "Two",
	}
)

func (self MyEnum) String() string {
	return MyEnumMapperNameByValue[int32(self)]
}

type MyTypeMini struct {
	Str  string
	Bool bool
}

func (self *MyTypeMini) String() string { return proto.CompactTextString(self) }

func (self *MyTypeMini) Size() (ret int) {

	ret += proto.SizeString(1, self.Str)

	ret += proto.SizeBool(2, self.Bool)

	return
}

func (self *MyTypeMini) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 1, self.Str)

	proto.MarshalBool(buffer, 2, self.Bool)

	return nil
}

func (self *MyTypeMini) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 1:
		return proto.UnmarshalString(buffer, wt, &self.Str)
	case 2:
		return proto.UnmarshalBool(buffer, wt, &self.Bool)

	}

	return proto.ErrUnknownField
}

type MySubType struct {
	Bool         bool
	Int32        int32
	UInt32       uint32
	Int64        int64
	UInt64       uint64
	Float32      float32
	Float64      float64
	Str          string
	BytesSlice   []byte
	BoolSlice    []bool
	Int32Slice   []int32
	UInt32Slice  []uint32
	Int64Slice   []int64
	UInt64Slice  []uint64
	Float32Slice []float32
	Float64Slice []float64
	StrSlice     []string
	Enum         MyEnum
	EnumSlice    []MyEnum
}

func (self *MySubType) String() string { return proto.CompactTextString(self) }

func (self *MySubType) Size() (ret int) {

	ret += proto.SizeBool(1, self.Bool)

	ret += proto.SizeInt32(2, self.Int32)

	ret += proto.SizeUInt32(3, self.UInt32)

	ret += proto.SizeInt64(4, self.Int64)

	ret += proto.SizeUInt64(5, self.UInt64)

	ret += proto.SizeFloat32(6, self.Float32)

	ret += proto.SizeFloat64(7, self.Float64)

	ret += proto.SizeString(8, self.Str)

	ret += proto.SizeBytes(9, self.BytesSlice)

	ret += proto.SizeBoolSlice(10, self.BoolSlice)

	ret += proto.SizeInt32Slice(11, self.Int32Slice)

	ret += proto.SizeUInt32Slice(12, self.UInt32Slice)

	ret += proto.SizeInt64Slice(13, self.Int64Slice)

	ret += proto.SizeUInt64Slice(14, self.UInt64Slice)

	ret += proto.SizeFloat32Slice(15, self.Float32Slice)

	ret += proto.SizeFloat64Slice(16, self.Float64Slice)

	ret += proto.SizeStringSlice(17, self.StrSlice)

	ret += proto.SizeInt32(18, int32(self.Enum))

	ret += proto.SizeInt32Slice(19, *(*[]int32)(unsafe.Pointer(&self.EnumSlice)))

	return
}

func (self *MySubType) Marshal(buffer *proto.Buffer) error {

	proto.MarshalBool(buffer, 1, self.Bool)

	proto.MarshalInt32(buffer, 2, self.Int32)

	proto.MarshalUInt32(buffer, 3, self.UInt32)

	proto.MarshalInt64(buffer, 4, self.Int64)

	proto.MarshalUInt64(buffer, 5, self.UInt64)

	proto.MarshalFloat32(buffer, 6, self.Float32)

	proto.MarshalFloat64(buffer, 7, self.Float64)

	proto.MarshalString(buffer, 8, self.Str)

	proto.MarshalBytes(buffer, 9, self.BytesSlice)

	proto.MarshalBoolSlice(buffer, 10, self.BoolSlice)

	proto.MarshalInt32Slice(buffer, 11, self.Int32Slice)

	proto.MarshalUInt32Slice(buffer, 12, self.UInt32Slice)

	proto.MarshalInt64Slice(buffer, 13, self.Int64Slice)

	proto.MarshalUInt64Slice(buffer, 14, self.UInt64Slice)

	proto.MarshalFloat32Slice(buffer, 15, self.Float32Slice)

	proto.MarshalFloat64Slice(buffer, 16, self.Float64Slice)

	proto.MarshalStringSlice(buffer, 17, self.StrSlice)

	proto.MarshalInt32(buffer, 18, int32(self.Enum))

	proto.MarshalInt32Slice(buffer, 19, *(*[]int32)(unsafe.Pointer(&self.EnumSlice)))

	return nil
}

func (self *MySubType) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 1:
		return proto.UnmarshalBool(buffer, wt, &self.Bool)
	case 2:
		return proto.UnmarshalInt32(buffer, wt, &self.Int32)
	case 3:
		return proto.UnmarshalUInt32(buffer, wt, &self.UInt32)
	case 4:
		return proto.UnmarshalInt64(buffer, wt, &self.Int64)
	case 5:
		return proto.UnmarshalUInt64(buffer, wt, &self.UInt64)
	case 6:
		return proto.UnmarshalFloat32(buffer, wt, &self.Float32)
	case 7:
		return proto.UnmarshalFloat64(buffer, wt, &self.Float64)
	case 8:
		return proto.UnmarshalString(buffer, wt, &self.Str)
	case 9:
		return proto.UnmarshalBytes(buffer, wt, &self.BytesSlice)
	case 10:
		return proto.UnmarshalBoolSlice(buffer, wt, &self.BoolSlice)
	case 11:
		return proto.UnmarshalInt32Slice(buffer, wt, &self.Int32Slice)
	case 12:
		return proto.UnmarshalUInt32Slice(buffer, wt, &self.UInt32Slice)
	case 13:
		return proto.UnmarshalInt64Slice(buffer, wt, &self.Int64Slice)
	case 14:
		return proto.UnmarshalUInt64Slice(buffer, wt, &self.UInt64Slice)
	case 15:
		return proto.UnmarshalFloat32Slice(buffer, wt, &self.Float32Slice)
	case 16:
		return proto.UnmarshalFloat64Slice(buffer, wt, &self.Float64Slice)
	case 17:
		return proto.UnmarshalStringSlice(buffer, wt, &self.StrSlice)
	case 18:
		return proto.UnmarshalInt32(buffer, wt, (*int32)(&self.Enum))
	case 19:
		return proto.UnmarshalInt32Slice(buffer, wt, (*[]int32)(unsafe.Pointer(&self.EnumSlice)))

	}

	return proto.ErrUnknownField
}

type MyType struct {
	Bool         bool
	Int32        int32
	UInt32       uint32
	Int64        int64
	UInt64       uint64
	Float32      float32
	Float64      float64
	Str          string
	Struct       MySubType
	BytesSlice   []byte
	BoolSlice    []bool
	Int32Slice   []int32
	UInt32Slice  []uint32
	Int64Slice   []int64
	UInt64Slice  []uint64
	Float32Slice []float32
	Float64Slice []float64
	StrSlice     []string
	StructSlice  []MySubType
	Enum         MyEnum
	EnumSlice    []MyEnum
}

func (self *MyType) String() string { return proto.CompactTextString(self) }

func (self *MyType) Size() (ret int) {

	ret += proto.SizeBool(1, self.Bool)

	ret += proto.SizeInt32(2, self.Int32)

	ret += proto.SizeUInt32(3, self.UInt32)

	ret += proto.SizeInt64(4, self.Int64)

	ret += proto.SizeUInt64(5, self.UInt64)

	ret += proto.SizeFloat32(6, self.Float32)

	ret += proto.SizeFloat64(7, self.Float64)

	ret += proto.SizeString(8, self.Str)

	ret += proto.SizeStruct(9, &self.Struct)

	ret += proto.SizeBytes(10, self.BytesSlice)

	ret += proto.SizeBoolSlice(11, self.BoolSlice)

	ret += proto.SizeInt32Slice(12, self.Int32Slice)

	ret += proto.SizeUInt32Slice(13, self.UInt32Slice)

	ret += proto.SizeInt64Slice(14, self.Int64Slice)

	ret += proto.SizeUInt64Slice(15, self.UInt64Slice)

	ret += proto.SizeFloat32Slice(16, self.Float32Slice)

	ret += proto.SizeFloat64Slice(17, self.Float64Slice)

	ret += proto.SizeStringSlice(18, self.StrSlice)

	if len(self.StructSlice) > 0 {
		for _, elm := range self.StructSlice {
			ret += proto.SizeStruct(19, &elm)
		}
	}

	ret += proto.SizeInt32(20, int32(self.Enum))

	ret += proto.SizeInt32Slice(21, *(*[]int32)(unsafe.Pointer(&self.EnumSlice)))

	return
}

func (self *MyType) Marshal(buffer *proto.Buffer) error {

	proto.MarshalBool(buffer, 1, self.Bool)

	proto.MarshalInt32(buffer, 2, self.Int32)

	proto.MarshalUInt32(buffer, 3, self.UInt32)

	proto.MarshalInt64(buffer, 4, self.Int64)

	proto.MarshalUInt64(buffer, 5, self.UInt64)

	proto.MarshalFloat32(buffer, 6, self.Float32)

	proto.MarshalFloat64(buffer, 7, self.Float64)

	proto.MarshalString(buffer, 8, self.Str)

	proto.MarshalStruct(buffer, 9, &self.Struct)

	proto.MarshalBytes(buffer, 10, self.BytesSlice)

	proto.MarshalBoolSlice(buffer, 11, self.BoolSlice)

	proto.MarshalInt32Slice(buffer, 12, self.Int32Slice)

	proto.MarshalUInt32Slice(buffer, 13, self.UInt32Slice)

	proto.MarshalInt64Slice(buffer, 14, self.Int64Slice)

	proto.MarshalUInt64Slice(buffer, 15, self.UInt64Slice)

	proto.MarshalFloat32Slice(buffer, 16, self.Float32Slice)

	proto.MarshalFloat64Slice(buffer, 17, self.Float64Slice)

	proto.MarshalStringSlice(buffer, 18, self.StrSlice)

	for _, elm := range self.StructSlice {
		proto.MarshalStruct(buffer, 19, &elm)
	}

	proto.MarshalInt32(buffer, 20, int32(self.Enum))

	proto.MarshalInt32Slice(buffer, 21, *(*[]int32)(unsafe.Pointer(&self.EnumSlice)))

	return nil
}

func (self *MyType) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 1:
		return proto.UnmarshalBool(buffer, wt, &self.Bool)
	case 2:
		return proto.UnmarshalInt32(buffer, wt, &self.Int32)
	case 3:
		return proto.UnmarshalUInt32(buffer, wt, &self.UInt32)
	case 4:
		return proto.UnmarshalInt64(buffer, wt, &self.Int64)
	case 5:
		return proto.UnmarshalUInt64(buffer, wt, &self.UInt64)
	case 6:
		return proto.UnmarshalFloat32(buffer, wt, &self.Float32)
	case 7:
		return proto.UnmarshalFloat64(buffer, wt, &self.Float64)
	case 8:
		return proto.UnmarshalString(buffer, wt, &self.Str)
	case 9:
		return proto.UnmarshalStruct(buffer, wt, &self.Struct)
	case 10:
		return proto.UnmarshalBytes(buffer, wt, &self.BytesSlice)
	case 11:
		return proto.UnmarshalBoolSlice(buffer, wt, &self.BoolSlice)
	case 12:
		return proto.UnmarshalInt32Slice(buffer, wt, &self.Int32Slice)
	case 13:
		return proto.UnmarshalUInt32Slice(buffer, wt, &self.UInt32Slice)
	case 14:
		return proto.UnmarshalInt64Slice(buffer, wt, &self.Int64Slice)
	case 15:
		return proto.UnmarshalUInt64Slice(buffer, wt, &self.UInt64Slice)
	case 16:
		return proto.UnmarshalFloat32Slice(buffer, wt, &self.Float32Slice)
	case 17:
		return proto.UnmarshalFloat64Slice(buffer, wt, &self.Float64Slice)
	case 18:
		return proto.UnmarshalStringSlice(buffer, wt, &self.StrSlice)
	case 19:
		var elm MySubType
		if err := proto.UnmarshalStruct(buffer, wt, &elm); err != nil {
			return err
		} else {
			self.StructSlice = append(self.StructSlice, elm)
			return nil
		}
	case 20:
		return proto.UnmarshalInt32(buffer, wt, (*int32)(&self.Enum))
	case 21:
		return proto.UnmarshalInt32Slice(buffer, wt, (*[]int32)(unsafe.Pointer(&self.EnumSlice)))

	}

	return proto.ErrUnknownField
}

type LoginREQ struct {
}

func (self *LoginREQ) String() string { return proto.CompactTextString(self) }

func (self *LoginREQ) Size() (ret int) {

	return
}

func (self *LoginREQ) Marshal(buffer *proto.Buffer) error {

	return nil
}

func (self *LoginREQ) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {

	}

	return proto.ErrUnknownField
}

type LoginACK struct {
}

func (self *LoginACK) String() string { return proto.CompactTextString(self) }

func (self *LoginACK) Size() (ret int) {

	return
}

func (self *LoginACK) Marshal(buffer *proto.Buffer) error {

	return nil
}

func (self *LoginACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {

	}

	return proto.ErrUnknownField
}

func init() {

}
