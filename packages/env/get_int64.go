package env

import (
	"barcode/packages/must"
	"strconv"
)

func GetInt(key string) int64 {
	str := GetString(key)

	return must.Panic2(strconv.ParseInt(str, 10, 64))
}
