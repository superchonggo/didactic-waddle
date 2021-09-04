package helpers

import (
	"log"
)

func HandleError(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

func HandleFatalError(err error, errorMessage string) error {
	if err != nil {
		log.Fatal(errorMessage)
		return err
	}

	return nil
}
