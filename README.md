# PQ Extended
Extends the `pq.Array` method to implement types not officially supported by PostgreSQL. This is a drop-in replacement for `pq.Array`

# Why?
The official Go driver for Postgres doesn't support slices of all types available in Go, such as `int8` and `int16`, and no unsigned ints at all. This results in the need for custom types which implement the `Scan` and `Value` interface to work properly. This module aims to be a drop-in replacement in these situations to add extended type support

# Usage
## pqextended.Array
`pqextended.Array` is a drop-in replacement for `pq.Array`. If a type is unknown, `pq.Array` is called internally. Currently supports:

- `[]uint8`

```go
postgres, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
if err != nil {
	panic(err)
}

_, err = postgres.Exec(`CREATE TABLE IF NOT EXISTS uint8_slices (
	id integer,
	slice integer[]
)`)
if err != nil {
	panic(err)
}

input := []uint8{0, 1, 2, 3}

_, err = postgres.Exec(`INSERT INTO uint8_slices (id, slice) VALUES (1, $1) ON CONFLICT DO NOTHING`, pqextended.Array(input))
if err != nil {
	panic(err)
}

var output []uint8

err = postgres.QueryRow(`SELECT slice FROM uint8_slices WHERE id=1`).Scan(pqextended.Array(&output))
if err != nil {
	panic(err)
}

if !bytes.Equal(input, output) {
	panic(fmt.Sprintf("Input and output did not match. Expected %v, got %v.", input, output))
}

fmt.Printf("Input: %v\n", input)
fmt.Printf("Output: %v\n", output)
```

# Tests
Tests require the use of a Postgres server. Set the `PQ_EXTENDED_TEST_POSTGRES_URI` environment variable to connect to the database used for tests. `.env` files are supported