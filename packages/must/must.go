package must

import "log"

func Panic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func Panic2[T any](val T, err error) T {
	Panic(err)

	return val
}

func Panic3[T, V any](val1 T, val2 V, err error) (T, V) {
	Panic(err)

	return val1, val2
}
