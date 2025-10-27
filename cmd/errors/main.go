package main

import (
	"errors"
	"fmt"
)

type Fatal struct {
	error
}

func (f *Fatal) Unwrap() error {
	return f.error
}

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}

type DetailedError struct {
	Code    int
	Message string
}

func (e DetailedError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func main() {

	err := errors.New("something went wrong")
	fmt.Println(err) // Output: something went wrong

	err = MyError{Message: "something went wrong"}
	fmt.Println(err) // Output: something went wrong

	err = DetailedError{Code: 404, Message: "Page not found"}
	fmt.Println(err) // Output: Error 404: Page not found

	err = &Fatal{error: errors.New("something went wrong")}

}
