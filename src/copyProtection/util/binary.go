package util

import (
	"fmt"
	"strconv"
)

func ToHex(i int, prefix bool) string {
	i64 := int64(i)

	if prefix {
		return "\\x" + strconv.FormatInt(i64, 16)
	}
	return strconv.FormatInt(i64, 16)
}

func ToHexWithDigit(i, digit int, prefix bool) string {
	h := ToHex(i, prefix)
	for i := 1; i < digit; i++ {
		h += ToHex(0, prefix)
	}
	return h
}

func StrToHexBinStr(s string) string {
	b := []byte(s)
	var bs string
	for _, v := range b {
		bs += fmt.Sprintf("%s", ToHex(int(v), true))
	}
	return bs
}
