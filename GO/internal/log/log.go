package log

import "fmt"

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
