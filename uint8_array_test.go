package pqextended

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestUInt8Array(t *testing.T) {
	godotenv.Load()

	postgres, err := sql.Open("postgres", os.Getenv("PQ_EXTENDED_TEST_POSTGRES_URI"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = postgres.Exec(`CREATE TABLE IF NOT EXISTS uint8_slices (
		id integer,
		uint8_slice smallint[]
	)`)
	if err != nil {
		t.Fatal(err)
	}

	input := []uint8{0, 1, 2, 255}

	_, err = postgres.Exec(`INSERT INTO uint8_slices (id, uint8_slice) VALUES ($1, $2) ON CONFLICT DO NOTHING`, 1, Array(input))
	if err != nil {
		t.Fatal(err)
	}

	var output []uint8

	err = postgres.QueryRow(`SELECT uint8_slice FROM uint8_slices WHERE id=$1`, 1).Scan(Array(&output))
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(input, output) {
		t.Fatalf("Input and output did not match. Expected %v, got %v.", input, output)
	}

	fmt.Printf("Input []uint8: %v\n", input)
	fmt.Printf("Output []uint8: %v\n", output)
}

func TestEmptyArray(t *testing.T) {
	godotenv.Load()

	postgres, err := sql.Open("postgres", os.Getenv("PQ_EXTENDED_TEST_POSTGRES_URI"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = postgres.Exec(`CREATE TABLE IF NOT EXISTS uint8_slices (
		id integer,
		uint8_slice smallint[]
	)`)
	if err != nil {
		t.Fatal(err)
	}

	input := []uint8{}

	_, err = postgres.Exec(`INSERT INTO uint8_slices (id, uint8_slice) VALUES ($1, $2) ON CONFLICT DO NOTHING`, 2, Array(input))
	if err != nil {
		t.Fatal(err)
	}

	var output []uint8

	err = postgres.QueryRow(`SELECT uint8_slice FROM uint8_slices WHERE id=$1`, 2).Scan(Array(&output))
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(input, output) {
		t.Fatalf("Input and output did not match. Expected %v, got %v.", input, output)
	}

	fmt.Printf("Input []uint8: %v\n", input)
	fmt.Printf("Output []uint8: %v\n", output)
}
