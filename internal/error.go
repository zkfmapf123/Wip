package internal

import (
	"log"

	"github.com/fatih/color"
)

func PanicError(err error) {
	log.Fatalln(color.RedString(err.Error()))
}
