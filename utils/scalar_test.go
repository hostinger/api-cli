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

func TestStringToUUID(t *testing.T) {
	const want = "550e8400-e29b-41d4-a716-446655440000"
	if got := StringToUUID(want); got.String() != want {
		t.Errorf("StringToUUID(%q) = %q, want %q", want, got, want)
	}
}

func TestStringToUUIDPanicsOnGarbage(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("StringToUUID did not panic on an unparseable uuid")
		}
	}()
	StringToUUID("not-a-uuid")
}
