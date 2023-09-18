package type_convert

import "strconv"

// StringToInt64IgnoreErr string --> int64 忽略异常
func StringToInt64IgnoreErr(val string) int64 {
	int64Val, _ := strconv.ParseInt(val, 10, 64)
	return int64Val
}

// StringToInt64Ignore string --> int64
func StringToInt64Ignore(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}
