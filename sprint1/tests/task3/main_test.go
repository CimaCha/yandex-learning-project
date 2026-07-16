package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Member struct {
	person       Person
	relationship Relationship
}

func TestAddNew(t *testing.T) {
	tests := []struct {
		name    string
		members []Member
		want    Family
		wantErr bool
	}{
		{
			name: "positive",
			members: []Member{
				{
					person:       Person{FirstName: "John", LastName: "Doe", Age: 33},
					relationship: "father",
				},
			},
			want: Family{
				Members: map[Relationship]Person{
					"father": {FirstName: "John", LastName: "Doe", Age: 33},
				},
			},
			wantErr: false,
		},
		{
			name: "multiple members",
			members: []Member{
				{
					person:       Person{FirstName: "John", LastName: "Doe", Age: 33},
					relationship: "father",
				},
				{
					person:       Person{FirstName: "Jane", LastName: "Doe", Age: 33},
					relationship: "mother",
				},
			},
			want: Family{
				Members: map[Relationship]Person{
					"father": {FirstName: "John", LastName: "Doe", Age: 33},
					"mother": {FirstName: "Jane", LastName: "Doe", Age: 33},
				},
			},
			wantErr: false,
		},
		{
			name: "add existing member",
			members: []Member{
				{
					person:       Person{FirstName: "John", LastName: "Doe", Age: 33},
					relationship: "father",
				},
				{
					person:       Person{FirstName: "John", LastName: "Doe", Age: 33},
					relationship: "father",
				},
			},
			want: Family{
				Members: map[Relationship]Person{
					"father": {FirstName: "John", LastName: "Doe", Age: 33},
				},
			},
			wantErr: true,
		},
		{
			name: "add person without relationship",
			members: []Member{
				{
					person: Person{FirstName: "John", LastName: "Doe", Age: 33},
				},
			},
			want: Family{
				Members: map[Relationship]Person{
					"": {FirstName: "John", LastName: "Doe", Age: 33},
				},
			},
			wantErr: false,
		},
		{
			name: "add empty person",
			members: []Member{
				{
					person:       Person{},
					relationship: "father",
				},
			},
			want: Family{
				Members: map[Relationship]Person{
					"father": {},
				},
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := Family{}
			for _, v := range test.members {
				err := f.AddNew(v.relationship, v.person)
				if err != nil {
					assert.EqualError(t, err, "relationship already exists")
				}
			}
			assert.Equal(t, f.Members, test.want.Members)
		})
	}
}
