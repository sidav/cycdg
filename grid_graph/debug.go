package graph

func debugPanic(msg string, args ...interface{}) {
	panic(sprintf(msg, args...))
}
