package app

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		fmt.Sprintln(err.Key, err.Message)
	}

	return
}
