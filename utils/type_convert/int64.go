package type_convert

import "strconv"

// Int64ToString int64 --> string
func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}
