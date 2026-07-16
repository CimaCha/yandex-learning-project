package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullName(t *testing.T) {
	tests := []struct {
		name string
		user User
		want string
	}{
		{
			name: "positive",
			user: User{
				FirstName: "John",
				LastName:  "Doe",
			},
			want: "John Doe",
		},
		{
			name: "without first name",
			user: User{
				LastName: "Doe",
			},
			want: " Doe",
		},
		{
			name: "without last name",
			user: User{
				FirstName: "John",
			},
			want: "John ",
		},
		{
			name: "empty user",
			user: User{},
			want: " ",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fullName := test.user.FullName()
			assert.Equal(t, test.want, fullName)
		})
	}
}
