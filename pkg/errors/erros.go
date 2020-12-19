package errors

import (
	"errors"
	"log"
)

// VerifyFatal lauches a fatal error if has an error
func VerifyFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// VerifyReport just prints the error if has an error and returns the error
func VerifyReport(err error) error {
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// ThrowError returns a custom error
func ThrowError(err string) error {
	return errors.New(err)
}
