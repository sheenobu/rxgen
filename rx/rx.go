// Package rx contains core types wrapped as reactive structures
package rx

// generate some core types

//go:generate rxgen -name Uint8 -type uint8 .
//go:generate rxgen -name Int8 -type int8 .

//go:generate rxgen -name Uint16 -type uint16 .
//go:generate rxgen -name Int16 -type int16 .

//go:generate rxgen -name Uint32 -type uint32 .
//go:generate rxgen -name Int32 -type int32 .

//go:generate rxgen -name Uint64 -type uint64 .
//go:generate rxgen -name Int64 -type int64 .

//go:generate rxgen -name Bool -type bool .

//go:generate rxgen -name String -type string .
