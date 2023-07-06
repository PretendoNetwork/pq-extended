package pqextended

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// UInt8Array implements the pq Scan and Value interface for uint8 slice type
type UInt8Array []uint8

// Scan implements the sql.Scanner interface
func (a *UInt8Array) Scan(src interface{}) error {
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

func (a *UInt8Array) scanBytes(src []byte) error {
	str := string(src)
	str = strings.TrimSuffix(str, "}")
	str = strings.TrimPrefix(str, "{")
	strSlice := strings.Split(str, ",")

	b := make([]uint8, 0, len(strSlice))

	for i := 0; i < len(strSlice); i++ {
		number, err := strconv.Atoi(strSlice[i])

		if err != nil {
			return err
		}

		b = append(b, uint8(number))
	}

	*a = b

	return nil
}

// Value implements the driver.Valuer interface
func (a UInt8Array) Value() (driver.Value, error) {
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
