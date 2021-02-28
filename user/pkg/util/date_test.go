package util_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chutified/booking-terminal/user/pkg/util"
)

func TestValidateDate(t *testing.T) {
	tests := []struct {
		name  string
		year  int32
		month int32
		day   int32
		valid bool
	}{
		{
			year:  0,
			month: 1,
			day:   1,
			valid: true,
		},
		{
			year:  9999,
			month: 12,
			day:   31,
			valid: true,
		},
		{
			year:  1234,
			month: 5,
			day:   6,
			valid: true,
		},
		{
			year:  2003,
			month: 4,
			day:   16,
			valid: true,
		},
		{
			year:  2004,
			month: 2,
			day:   27,
			valid: true,
		},
		{
			year:  -1234,
			month: 5,
			day:   6,
			valid: false,
		},
		{
			year:  1234,
			month: -5,
			day:   6,
			valid: false,
		},
		{
			year:  1234,
			month: 5,
			day:   -6,
			valid: false,
		},
		{
			year:  1234,
			month: 0,
			day:   6,
			valid: false,
		},
		{
			year:  1234,
			month: 5,
			day:   0,
			valid: false,
		},
		{
			year:  1234,
			month: 42,
			day:   6,
			valid: false,
		},
		{
			year:  1234,
			month: 5,
			day:   42,
			valid: false,
		},
		{
			year:  1234,
			month: 13,
			day:   6,
			valid: false,
		},
		{
			year:  1234,
			month: 5,
			day:   32,
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d/%d/%d (dd/mm/yyyy)", tt.day, tt.month, tt.year), func(t *testing.T) {
			valid := util.ValidateDate(tt.year, tt.month, tt.day)
			require.Equal(t, tt.valid, valid)
		})
	}
}
