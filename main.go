// This little function is expecting to receive data via stdin. If there is any
// data it exits with a error code of 1. If stdin is empty, then it exists with
// an error code of 0 (ie. no error).
// This is useful when using CI and things like gofmt.
//
// Example:
// If there is any content in the gofmt command (ie. there is a formatting
// error), then stdinerror returns a error response (exit(1)).
//     gofmt -l folder | ./stdinerror
//

package main

import "os"
import "log"
import "bufio"

func main() {
	s := bufio.NewScanner(os.Stdin)
	err := false

	for s.Scan() {
		if len(s.Text()) > 0 {
			err = true
			log.Println("ERROR ", s.Text())
		}

	}

	if err {
		os.Exit(1)
	}
	os.Exit(0)

}
