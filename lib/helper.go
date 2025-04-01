package lib

import (
	"encoding/json"
	"os"
	"strings"
)

func If[T any](cond bool, TrueValue, FalseValue T) T {
	if cond {
		return TrueValue
	}
	return FalseValue
}

func ReadJSONFile[T any](filepath string) (T, error) {
	var result T

	// read file as []byte
	data, err := os.ReadFile(filepath)
	if err != nil {
		return result, err
	}

	// decode JSON to format type
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func CutSlices[T any](values []T, start, end int) []T {
	// validate start and end position
	size := len(values)
	start = If(start < 0, 0, start)
	end = If(end > size, size, end)

	if start > end {
		return []T{}
	}

	return values[start:end]
}

func OffsetStartEnd(page, limit int) (start int, end int) {
	start = (page - 1) * limit
	end = (page) * limit
	return
}

func Trim(value string) string {
	return strings.TrimSpace(value)
}

func Ptr[T any](value T) *T {
	return &value
}
