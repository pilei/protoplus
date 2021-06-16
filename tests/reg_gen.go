// Generated by github.com/davyxu/protoplus
package tests

import (
	"github.com/davyxu/cellnet/codec"
	"github.com/davyxu/cellnet/meta"
	"reflect"
)

var (
	_ cellmeta.Meta
	_ cellcodec.CodecRecycler
	_ reflect.Kind
)

func init() {

	cellmeta.Register(&cellmeta.Meta{
		FullName: "tests.MyTypeMini",
		ID:       0,
		New:      func() interface{} { return &MyTypeMini{} },
		Type:     reflect.TypeOf((*MyTypeMini)(nil)).Elem(),
		Codec:    cellcodec.MustGetByName("protoplus"),
	})
	cellmeta.Register(&cellmeta.Meta{
		FullName: "tests.MySubType",
		ID:       0,
		New:      func() interface{} { return &MySubType{} },
		Type:     reflect.TypeOf((*MySubType)(nil)).Elem(),
		Codec:    cellcodec.MustGetByName("protoplus"),
	})
	cellmeta.Register(&cellmeta.Meta{
		FullName: "tests.MyType",
		ID:       0,
		New:      func() interface{} { return &MyType{} },
		Type:     reflect.TypeOf((*MyType)(nil)).Elem(),
		Codec:    cellcodec.MustGetByName("protoplus"),
	})
	cellmeta.Register(&cellmeta.Meta{
		FullName: "tests.LoginREQ",
		ID:       17076,
		New:      func() interface{} { return &LoginREQ{} },
		Type:     reflect.TypeOf((*LoginREQ)(nil)).Elem(),
		Codec:    cellcodec.MustGetByName("protoplus"),
	})
	cellmeta.Register(&cellmeta.Meta{
		FullName: "tests.LoginACK",
		ID:       44443,
		New:      func() interface{} { return &LoginACK{} },
		Type:     reflect.TypeOf((*LoginACK)(nil)).Elem(),
		Codec:    cellcodec.MustGetByName("protoplus"),
	})
}
