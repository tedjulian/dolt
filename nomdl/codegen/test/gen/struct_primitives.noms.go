// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_primitives_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("StructPrimitives",
			[]types.Field{
				types.Field{"uint64", types.MakePrimitiveTypeRef(types.UInt64Kind), false},
				types.Field{"uint32", types.MakePrimitiveTypeRef(types.UInt32Kind), false},
				types.Field{"uint16", types.MakePrimitiveTypeRef(types.UInt16Kind), false},
				types.Field{"uint8", types.MakePrimitiveTypeRef(types.UInt8Kind), false},
				types.Field{"int64", types.MakePrimitiveTypeRef(types.Int64Kind), false},
				types.Field{"int32", types.MakePrimitiveTypeRef(types.Int32Kind), false},
				types.Field{"int16", types.MakePrimitiveTypeRef(types.Int16Kind), false},
				types.Field{"int8", types.MakePrimitiveTypeRef(types.Int8Kind), false},
				types.Field{"float64", types.MakePrimitiveTypeRef(types.Float64Kind), false},
				types.Field{"float32", types.MakePrimitiveTypeRef(types.Float32Kind), false},
				types.Field{"bool", types.MakePrimitiveTypeRef(types.BoolKind), false},
				types.Field{"string", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"blob", types.MakePrimitiveTypeRef(types.BlobKind), false},
				types.Field{"value", types.MakePrimitiveTypeRef(types.ValueKind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_primitives_CachedRef = types.RegisterPackage(&p)
}

// StructPrimitives

type StructPrimitives struct {
	_uint64  uint64
	_uint32  uint32
	_uint16  uint16
	_uint8   uint8
	_int64   int64
	_int32   int32
	_int16   int16
	_int8    int8
	_float64 float64
	_float32 float32
	_bool    bool
	_string  string
	_blob    types.Blob
	_value   types.Value

	ref *ref.Ref
}

func NewStructPrimitives() StructPrimitives {
	return StructPrimitives{
		_uint64:  uint64(0),
		_uint32:  uint32(0),
		_uint16:  uint16(0),
		_uint8:   uint8(0),
		_int64:   int64(0),
		_int32:   int32(0),
		_int16:   int16(0),
		_int8:    int8(0),
		_float64: float64(0),
		_float32: float32(0),
		_bool:    false,
		_string:  "",
		_blob:    types.NewEmptyBlob(),
		_value:   types.Bool(false),

		ref: &ref.Ref{},
	}
}

type StructPrimitivesDef struct {
	Uint64  uint64
	Uint32  uint32
	Uint16  uint16
	Uint8   uint8
	Int64   int64
	Int32   int32
	Int16   int16
	Int8    int8
	Float64 float64
	Float32 float32
	Bool    bool
	String  string
	Blob    types.Blob
	Value   types.Value
}

func (def StructPrimitivesDef) New() StructPrimitives {
	return StructPrimitives{
		_uint64:  def.Uint64,
		_uint32:  def.Uint32,
		_uint16:  def.Uint16,
		_uint8:   def.Uint8,
		_int64:   def.Int64,
		_int32:   def.Int32,
		_int16:   def.Int16,
		_int8:    def.Int8,
		_float64: def.Float64,
		_float32: def.Float32,
		_bool:    def.Bool,
		_string:  def.String,
		_blob:    def.Blob,
		_value:   def.Value,
		ref:      &ref.Ref{},
	}
}

func (s StructPrimitives) Def() (d StructPrimitivesDef) {
	d.Uint64 = s._uint64
	d.Uint32 = s._uint32
	d.Uint16 = s._uint16
	d.Uint8 = s._uint8
	d.Int64 = s._int64
	d.Int32 = s._int32
	d.Int16 = s._int16
	d.Int8 = s._int8
	d.Float64 = s._float64
	d.Float32 = s._float32
	d.Bool = s._bool
	d.String = s._string
	d.Blob = s._blob
	d.Value = s._value
	return
}

var __typeRefForStructPrimitives types.TypeRef
var __typeDefForStructPrimitives types.TypeRef

func (m StructPrimitives) TypeRef() types.TypeRef {
	return __typeRefForStructPrimitives
}

func init() {
	__typeRefForStructPrimitives = types.MakeTypeRef(__genPackageInFile_struct_primitives_CachedRef, 0)
	__typeDefForStructPrimitives = types.MakeStructTypeRef("StructPrimitives",
		[]types.Field{
			types.Field{"uint64", types.MakePrimitiveTypeRef(types.UInt64Kind), false},
			types.Field{"uint32", types.MakePrimitiveTypeRef(types.UInt32Kind), false},
			types.Field{"uint16", types.MakePrimitiveTypeRef(types.UInt16Kind), false},
			types.Field{"uint8", types.MakePrimitiveTypeRef(types.UInt8Kind), false},
			types.Field{"int64", types.MakePrimitiveTypeRef(types.Int64Kind), false},
			types.Field{"int32", types.MakePrimitiveTypeRef(types.Int32Kind), false},
			types.Field{"int16", types.MakePrimitiveTypeRef(types.Int16Kind), false},
			types.Field{"int8", types.MakePrimitiveTypeRef(types.Int8Kind), false},
			types.Field{"float64", types.MakePrimitiveTypeRef(types.Float64Kind), false},
			types.Field{"float32", types.MakePrimitiveTypeRef(types.Float32Kind), false},
			types.Field{"bool", types.MakePrimitiveTypeRef(types.BoolKind), false},
			types.Field{"string", types.MakePrimitiveTypeRef(types.StringKind), false},
			types.Field{"blob", types.MakePrimitiveTypeRef(types.BlobKind), false},
			types.Field{"value", types.MakePrimitiveTypeRef(types.ValueKind), false},
		},
		types.Choices{},
	)
	types.RegisterStructBuilder(__typeRefForStructPrimitives, builderForStructPrimitives)
}

func (s StructPrimitives) InternalImplementation() types.Struct {
	// TODO: Remove this
	m := map[string]types.Value{
		"uint64":  types.UInt64(s._uint64),
		"uint32":  types.UInt32(s._uint32),
		"uint16":  types.UInt16(s._uint16),
		"uint8":   types.UInt8(s._uint8),
		"int64":   types.Int64(s._int64),
		"int32":   types.Int32(s._int32),
		"int16":   types.Int16(s._int16),
		"int8":    types.Int8(s._int8),
		"float64": types.Float64(s._float64),
		"float32": types.Float32(s._float32),
		"bool":    types.Bool(s._bool),
		"string":  types.NewString(s._string),
		"blob":    s._blob,
		"value":   s._value,
	}
	return types.NewStruct(__typeRefForStructPrimitives, __typeDefForStructPrimitives, m)
}

func builderForStructPrimitives() chan types.Value {
	c := make(chan types.Value)
	s := StructPrimitives{ref: &ref.Ref{}}
	go func() {
		s._uint64 = uint64((<-c).(types.UInt64))
		s._uint32 = uint32((<-c).(types.UInt32))
		s._uint16 = uint16((<-c).(types.UInt16))
		s._uint8 = uint8((<-c).(types.UInt8))
		s._int64 = int64((<-c).(types.Int64))
		s._int32 = int32((<-c).(types.Int32))
		s._int16 = int16((<-c).(types.Int16))
		s._int8 = int8((<-c).(types.Int8))
		s._float64 = float64((<-c).(types.Float64))
		s._float32 = float32((<-c).(types.Float32))
		s._bool = bool((<-c).(types.Bool))
		s._string = (<-c).(types.String).String()
		s._blob = (<-c).(types.Blob)
		s._value = (<-c)

		c <- s
	}()
	return c
}

func (s StructPrimitives) Equals(other types.Value) bool {
	return other != nil && __typeRefForStructPrimitives.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s StructPrimitives) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s StructPrimitives) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForStructPrimitives.Chunks()...)
	chunks = append(chunks, s._blob.Chunks()...)
	chunks = append(chunks, s._value.Chunks()...)
	return
}

func (s StructPrimitives) Uint64() uint64 {
	return s._uint64
}

func (s StructPrimitives) SetUint64(val uint64) StructPrimitives {
	s._uint64 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Uint32() uint32 {
	return s._uint32
}

func (s StructPrimitives) SetUint32(val uint32) StructPrimitives {
	s._uint32 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Uint16() uint16 {
	return s._uint16
}

func (s StructPrimitives) SetUint16(val uint16) StructPrimitives {
	s._uint16 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Uint8() uint8 {
	return s._uint8
}

func (s StructPrimitives) SetUint8(val uint8) StructPrimitives {
	s._uint8 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Int64() int64 {
	return s._int64
}

func (s StructPrimitives) SetInt64(val int64) StructPrimitives {
	s._int64 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Int32() int32 {
	return s._int32
}

func (s StructPrimitives) SetInt32(val int32) StructPrimitives {
	s._int32 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Int16() int16 {
	return s._int16
}

func (s StructPrimitives) SetInt16(val int16) StructPrimitives {
	s._int16 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Int8() int8 {
	return s._int8
}

func (s StructPrimitives) SetInt8(val int8) StructPrimitives {
	s._int8 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Float64() float64 {
	return s._float64
}

func (s StructPrimitives) SetFloat64(val float64) StructPrimitives {
	s._float64 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Float32() float32 {
	return s._float32
}

func (s StructPrimitives) SetFloat32(val float32) StructPrimitives {
	s._float32 = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Bool() bool {
	return s._bool
}

func (s StructPrimitives) SetBool(val bool) StructPrimitives {
	s._bool = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) String() string {
	return s._string
}

func (s StructPrimitives) SetString(val string) StructPrimitives {
	s._string = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Blob() types.Blob {
	return s._blob
}

func (s StructPrimitives) SetBlob(val types.Blob) StructPrimitives {
	s._blob = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructPrimitives) Value() types.Value {
	return s._value
}

func (s StructPrimitives) SetValue(val types.Value) StructPrimitives {
	s._value = val
	s.ref = &ref.Ref{}
	return s
}
