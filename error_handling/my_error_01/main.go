package main

import (
	"errors"
	"fmt"
)

func main() {

	dd, err := mentor(-4)
	fmt.Println(err, dd)
}

var ErrorMentor = &result{name: "xxxx", username: "xxxx", err: errors.New("not negatif number")}

type result struct {
	name     string
	username string
	err      error
}

func (r *result) Error() string {

	return fmt.Sprintf("happend error: %v is %v  %v\n ", r.name, r.username, r.err)

}

func mentor(f int) (int, error) {
	if f < 0 {
		return 0, ErrorMentor
	}

	return f, nil

}

func added(s, d int) (int, error) {
	if s < 0 || d < 0 {
		return 0, ErrorMentor
	}

	return s + d, nil

}
