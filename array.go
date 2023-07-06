package pqextended

import (
	"database/sql"
	"database/sql/driver"

	"github.com/lib/pq"
)

func Array(a interface{}) interface {
	driver.Valuer
	sql.Scanner
} {
	switch a := a.(type) {
	case []uint8:
		return (*UInt8Array)(&a)
	case *[]uint8:
		return (*UInt8Array)(a)
	}

	return pq.Array(a)
}
