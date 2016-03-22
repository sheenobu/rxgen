// package rx contains core types wrapped as reactive structures
package rx

// generate some core types

//go:generate rxgen -name RxUint8 -type uint8 .
//go:generate rxgen -name RxInt8 -type int8 .

//go:generate rxgen -name RxUint16 -type uint16 .
//go:generate rxgen -name RxInt16 -type int16 .

//go:generate rxgen -name RxUint32 -type uint32 .
//go:generate rxgen -name RxInt32 -type int32 .

//go:generate rxgen -name RxUint64 -type uint64 .
//go:generate rxgen -name RxInt64 -type int64 .

//go:generate rxgen -name RxBool -type bool .

//go:generate rxgen -name RxString -type string .
