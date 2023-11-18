package replacement

import (
	"fmt"
)

func debugPanic(msg string, args ...interface{}) {
	panic(sprintf(msg, args...))
}

func sprintf(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}
