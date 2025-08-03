package utils

import "strconv"

func StringToBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return b
}

func StringToInt(str string) int {
	b, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return b
}

func ConvertUintToString(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}
