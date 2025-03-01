package utils

import "runtime"

func GetStackTrace() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, true)
	return string(buf[:n])
}
