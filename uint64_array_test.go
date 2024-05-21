package pqextended

import (
	"database/sql"
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/joho/godotenv"
)

func TestUInt64Array(t *testing.T) {
	godotenv.Load()

	postgres, err := sql.Open("postgres", os.Getenv("PQ_EXTENDED_TEST_POSTGRES_URI"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = postgres.Exec(`CREATE TABLE IF NOT EXISTS uint64_slices (
		id integer,
		uint64_slice numeric(20)[]
	)`)
	if err != nil {
		t.Fatal(err)
	}

	input := []uint64{0, 1, 2, 18446744073709551615}

	_, err = postgres.Exec(`INSERT INTO uint64_slices (id, uint64_slice) VALUES ($1, $2) ON CONFLICT DO NOTHING`, 1, Array(input))
	if err != nil {
		t.Fatal(err)
	}

	var output []uint64

	err = postgres.QueryRow(`SELECT uint64_slice FROM uint64_slices WHERE id=$1`, 1).Scan(Array(&output))
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(input, output) {
		t.Fatalf("Input and output did not match. Expected %v, got %v.", input, output)
	}

	fmt.Printf("Input []uint64: %v\n", input)
	fmt.Printf("Output []uint64: %v\n", output)
}
