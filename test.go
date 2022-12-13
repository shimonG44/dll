package main

import (
	"C"

	"github.com/tawesoft/golib/dialog"
)

func create_popup() bool {
	dialog.Alert("Message")
	return true
}

func main() {

}
