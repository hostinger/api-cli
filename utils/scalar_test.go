package utils

import (
	"testing"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func TestStringToDate(t *testing.T) {
	got := StringToDate("2026-03-16")
	if want := "2026-03-16"; got.Format(openapi_types.DateFormat) != want {
		t.Errorf("StringToDate(%q) = %q, want %q", "2026-03-16", got.Format(openapi_types.DateFormat), want)
	}
}

func TestStringToDatePanicsOnGarbage(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("StringToDate did not panic on an unparseable date")
		}
	}()
	StringToDate("16/03/2026")
}
