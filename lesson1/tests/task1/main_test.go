package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		num  float64
		want float64
	}{
		{
			name: "positive",
			num:  5,
			want: 5,
		},
		{
			name: "negative",
			num:  -5,
			want: 5,
		},
		{
			name: "negative decimal",
			num:  -2.00001,
			want: 2.00001,
		},
		{
			name: "positive decimal",
			num:  2.00001,
			want: 2.00001,
		},
		{
			name: "negative small number",
			num:  -0.00000001,
			want: 0.00000001,
		},
		{
			name: "positive small number",
			num:  0.00000001,
			want: 0.00000001,
		},
		{
			name: "zero",
			num:  0,
			want: 0,
		},
		{
			name: "large number",
			num:  999999999999999,
			want: 999999999999999,
		}, {
			name: "negative large number",
			num:  -999999999999999,
			want: 999999999999999,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Abs(test.num))
		})
	}
}
