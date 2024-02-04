package mock

import (
	"fmt"
	"testing"
	"time"
)

func TestGenRecordNwithdrawalRecords(t *testing.T) {
	n := 100
	appliedAt := time.Date(2024, 2, 4, 0, 0, 0, 0, time.Local)
	got := GenRecordNwithdrawalRecords(n, appliedAt, "_1")

	fmt.Println(got)
}

func TestGenAndImportNwithdrawalRecords(t *testing.T) {
	GenAndImportNwithdrawalRecords()
}
