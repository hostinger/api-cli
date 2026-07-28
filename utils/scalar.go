package utils

import (
	"strconv"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func StringPtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func IntPtrOrNil(i int) *int {
	if i == -1 {
		return nil
	}
	return &i
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func StringArrayToIntArray(strArr []string) []int {
	intArr := make([]int, 0, len(strArr))

	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intArr = append(intArr, num)
	}

	return intArr
}

func StringToTime(dateStr string) time.Time {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}

	return t
}

func StringToDate(dateStr string) openapi_types.Date {
	t, err := time.Parse(openapi_types.DateFormat, dateStr)
	if err != nil {
		panic(err)
	}

	return openapi_types.Date{Time: t}
}

// StringToEmail keeps format:email query params converting through utils like
// every other formatted kind, so generated commands need no extra import.
func StringToEmail(email string) openapi_types.Email {
	return openapi_types.Email(email)
}
