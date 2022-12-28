package utils

import (
	"github.com/itsparser/airstrike/pkg/core/utils/helper"
	"log"
)

func ErrorHandling(err error) {
	log.Print("Error Have been raised from the following Method " + helper.MyCaller())
	log.Fatal(err)
}
