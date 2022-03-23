// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package errorlogger

import "fmt"

func Example() {
	exampleLog := New()
	exampleLog.Info("Example log info message...")
	fmt.Println("Example import acknowledgement from package errorlogger")
}
