package pqextended

import (
	"database/sql"
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/joho/godotenv"
)

func TestInt8Array(t *testing.T) {
	godotenv.Load()

	postgres, err := sql.Open("postgres", os.Getenv("PQ_EXTENDED_TEST_POSTGRES_URI"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = postgres.Exec(`CREATE TABLE IF NOT EXISTS int8_slices (
		id integer,
		int8_slice smallint[]
	)`)
	if err != nil {
		t.Fatal(err)
	}

	input := []int8{0, -1, 127, -128}

	_, err = postgres.Exec(`INSERT INTO int8_slices (id, int8_slice) VALUES ($1, $2) ON CONFLICT DO NOTHING`, 1, Array(input))
	if err != nil {
		t.Fatal(err)
	}

	var output []int8

	err = postgres.QueryRow(`SELECT int8_slice FROM int8_slices WHERE id=$1`, 1).Scan(Array(&output))
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(input, output) {
		t.Fatalf("Input and output did not match. Expected %v, got %v.", input, output)
	}

	fmt.Printf("Input []int8: %v\n", input)
	fmt.Printf("Output []int8: %v\n", output)
}
