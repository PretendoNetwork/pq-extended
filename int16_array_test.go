package pqextended

import (
	"database/sql"
	"fmt"
	"os"
	"slices"
	"testing"

	"github.com/joho/godotenv"
)

func TestInt16Array(t *testing.T) {
	godotenv.Load()

	postgres, err := sql.Open("postgres", os.Getenv("PQ_EXTENDED_TEST_POSTGRES_URI"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = postgres.Exec(`CREATE TABLE IF NOT EXISTS int16_slices (
		id integer,
		int16_slice smallint[]
	)`)
	if err != nil {
		t.Fatal(err)
	}

	input := []int16{0, -1, 32767, -32768}

	_, err = postgres.Exec(`INSERT INTO int16_slices (id, int16_slice) VALUES ($1, $2) ON CONFLICT DO NOTHING`, 1, Array(input))
	if err != nil {
		t.Fatal(err)
	}

	var output []int16

	err = postgres.QueryRow(`SELECT int16_slice FROM int16_slices WHERE id=$1`, 1).Scan(Array(&output))
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(input, output) {
		t.Fatalf("Input and output did not match. Expected %v, got %v.", input, output)
	}

	fmt.Printf("Input []int16: %v\n", input)
	fmt.Printf("Output []int16: %v\n", output)
}
