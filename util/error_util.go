//Contains a collection of utility functions
package util

import (
	"fmt"
	"os"
)

//Helper that simply checks if provided error object is nil.
//If provided error object is nil, prints error message and terminates
//program execution.
func CheckError(e error) {
	if e != nil {
		fmt.Println(e.Error())
		//TODO: Implement more graceful way of dealing with errors.
		os.Exit(-1)
	}
}
