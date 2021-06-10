package main

import (
	"encoding/json"
	"io"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func main(u User) {
	w := io.Writer
	u.FirstName = "Caige"
	u.LastName = "Kelly"

	enc := json.NewEncoder(w)
	enc.Encode(data)
}
