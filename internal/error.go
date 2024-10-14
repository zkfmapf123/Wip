package internal

import (
	"log"

	"github.com/fatih/color"
)

func PanicError(err error) {
	if err != nil {
		log.Fatalln(color.RedString(err.Error()))
	}
}

func WrapMustFunc[T any](fn func() (T, error)) T {
	v, err := fn()
	if err != nil {
		PanicError(err)
	}

	return v
}
