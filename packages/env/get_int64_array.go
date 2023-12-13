package env

import (
	"barcode/packages/must"
	"strconv"
	"strings"
)

func GetInt64Array(key string) []int64 {
	str := GetString(key)

	strs := strings.Split(str, ",")
	var ints []int64

	for _, item := range strs {
		intValue := must.Panic2(strconv.ParseInt(item, 10, 64))
		ints = append(ints, intValue)
	}

	return ints
}
