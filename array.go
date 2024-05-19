package pqextended

import (
	"database/sql"
	"database/sql/driver"

	"github.com/lib/pq"
)

// Array returns the optimal driver.Valuer and sql.Scanner for an array or
// slice of any dimension.
//
// Wraps around pq.Array, but adds support for more array types.
func Array(a interface{}) interface {
	driver.Valuer
	sql.Scanner
} {
	switch a := a.(type) {
	case []int8:
		return (*Int8Array)(&a)
	case *[]int8:
		return (*Int8Array)(a)
	case []int16:
		return (*Int16Array)(&a)
	case *[]int16:
		return (*Int16Array)(a)

	case []uint8:
		return (*UInt8Array)(&a)
	case *[]uint8:
		return (*UInt8Array)(a)
	case []uint16:
		return (*UInt16Array)(&a)
	case *[]uint16:
		return (*UInt16Array)(a)
	case []uint32:
		return (*UInt32Array)(&a)
	case *[]uint32:
		return (*UInt32Array)(a)
	case []uint64:
		return (*UInt64Array)(&a)
	case *[]uint64:
		return (*UInt64Array)(a)
	}

	return pq.Array(a)
}
