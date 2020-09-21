package strs

import (
	"strconv"
)

func StrIsEmpty(str string) bool {
	return str == ""
}

func StrNotEmpty(str string) bool {
	return str != ""
}

func StrWithFallback(str, fallback string) string {
	if StrNotEmpty(str) {
		return str
	}

	return fallback
}

func StrToInt64WithFallback(str string, fallback int64) int64 {
	output, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return fallback
	}
	return output
}

func StrToInt64WithDefaultZero(str string) int64 {
	return StrToInt64WithFallback(str, 0)
}

func StrToUint64WithFallback(str string, fallback uint64) uint64 {
	output, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return fallback
	}
	return output
}

func StrToUint64WithDefaultZero(str string) uint64 {
	return StrToUint64WithFallback(str, 0)
}

func StrToIntWithFallback(str string, fallback int) int {
	output, err := strconv.Atoi(str)
	if err != nil {
		return fallback
	}
	return output
}

func StrToIntWithDefaultZero(str string) int {
	return StrToIntWithFallback(str, 0)
}

func StrToFloat32WithFallback(str string, fallback float32) float32 {
	output, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return fallback
	}
	return float32(output)
}

func StrToFloat32WithDefaultZero(str string) float32 {
	return StrToFloat32WithFallback(str, 0)
}

func StrToFloat64WithFallback(str string, fallback float64) float64 {
	output, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fallback
	}
	return output
}

func StrToFloat64WithDefaultZero(str string) int {
	return StrToIntWithFallback(str, 0)
}
