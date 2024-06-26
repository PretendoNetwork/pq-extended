package pqextended

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// UInt32Array implements the pq Scan and Value interface for uint32 slice type
type UInt32Array []uint32

// Scan implements the sql.Scanner interface
func (a *UInt32Array) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return a.scanBytes(src)
	case string:
		return a.scanBytes([]byte(src))
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("PQ Extended: Cannot convert %T to UInt8Array", src)
}

func (a *UInt32Array) scanBytes(src []byte) error {
	str := string(src)
	str = strings.TrimSuffix(str, "}")
	str = strings.TrimPrefix(str, "{")

	if str == "" {
		*a = make([]uint32, 0)
		return nil
	}

	strSlice := strings.Split(str, ",")

	b := make([]uint32, 0, len(strSlice))

	for i := 0; i < len(strSlice); i++ {
		number, err := strconv.ParseUint(strSlice[i], 10, 32)

		if err != nil {
			return err
		}

		b = append(b, uint32(number))
	}

	*a = b

	return nil
}

// Value implements the driver.Valuer interface
func (a UInt32Array) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		var b strings.Builder

		b.WriteString("{")

		for i := 0; i < n; i++ {
			if i != n-1 {
				b.WriteString(fmt.Sprintf("%d,", a[i]))
			} else {
				b.WriteString(fmt.Sprintf("%d", a[i]))
			}
		}

		b.WriteString("}")

		return b.String(), nil
	}

	return "{}", nil
}
