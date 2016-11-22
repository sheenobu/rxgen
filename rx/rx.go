// +build ignore

// Package rx contains core types wrapped as reactive structures
package rx

// generate some core types

// +gen rx:"Builtin[uint8]"
type Uint8_ uint8

// +gen rx:"Builtin[int8]"
type Int8_ int8

// +gen rx:"Builtin[uint8]"
type Uint16_ uint16

// +gen rx:"Builtin[int16]"
type Int16_ int16

// +gen rx:"Builtin[uint16]"
type Uint32_ uint32

// +gen rx:"Builtin[int32]"
type Int32_ int32

// +gen rx:"Builtin[uint64]"
type Uint64_ uint64

// +gen rx:"Builtin[int64]"
type Int64_ int64

// +gen rx:"Builtin[bool]"
type Bool_ bool

// +gen rx:"Builtin[string]"
type String_ string
