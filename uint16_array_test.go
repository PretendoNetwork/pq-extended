package pqextended

import (
	"database/sql"
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/joho/godotenv"
)

func TestUInt16Array(t *testing.T) {
	godotenv.Load()

	postgres, err := sql.Open("postgres", os.Getenv("PQ_EXTENDED_TEST_POSTGRES_URI"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = postgres.Exec(`CREATE TABLE IF NOT EXISTS uint16_slices (
		id integer,
		uint16_slice integer[]
	)`)
	if err != nil {
		t.Fatal(err)
	}

	input := []uint16{0, 1, 2, 65535}

	_, err = postgres.Exec(`INSERT INTO uint16_slices (id, uint16_slice) VALUES ($1, $2) ON CONFLICT DO NOTHING`, 1, Array(input))
	if err != nil {
		t.Fatal(err)
	}

	var output []uint16

	err = postgres.QueryRow(`SELECT uint16_slice FROM uint16_slices WHERE id=$1`, 1).Scan(Array(&output))
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(input, output) {
		t.Fatalf("Input and output did not match. Expected %v, got %v.", input, output)
	}

	fmt.Printf("Input []uint16: %v\n", input)
	fmt.Printf("Output []uint16: %v\n", output)
}
