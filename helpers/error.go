package helpers

import (
	"fmt"
	"os"
)

func CmdErrorHandler(err error) {
	fmt.Println(err)
	os.Exit(1)
}
