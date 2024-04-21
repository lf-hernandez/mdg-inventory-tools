package utils

import "strconv"

func FormatFloat(value *float64) string {
	if value != nil {
		return strconv.FormatFloat(*value, 'f', 2, 64)
	}
	return ""
}

func FormatInt(value *int) string {
	if value != nil {
		return strconv.Itoa(*value)
	}
	return ""
}
