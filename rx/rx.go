// Package rx contains core types wrapped as reactive structures
package rx

// generate some core types

//go:generate rxgen -name Uint8 -type uint8 rx.go
//go:generate rxgen -name Int8 -type int8 rx.go

//go:generate rxgen -name Uint16 -type uint16 rx.go
//go:generate rxgen -name Int16 -type int16 rx.go

//go:generate rxgen -name Uint32 -type uint32 rx.go
//go:generate rxgen -name Int32 -type int32 rx.go

//go:generate rxgen -name Uint64 -type uint64 rx.go
//go:generate rxgen -name Int64 -type int64 rx.go

//go:generate rxgen -name Bool -type bool rx.go

//go:generate rxgen -name String -type string rx.go
