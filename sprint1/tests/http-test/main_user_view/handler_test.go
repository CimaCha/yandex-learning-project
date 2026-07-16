package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserViewHandler(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		user        User
	}
	tests := []struct {
		name    string
		request string
		users   map[string]User
		want    want
	}{
		{
			name: "positive",
			users: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				contentType: "application/json",
				statusCode:  200,
				user: User{ID: "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			request: "/users?user_id=id1",
		},
		{
			name: "status: not found",
			users: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  404,
				user:        User{},
			},
			request: "/users?user_id=id2",
		},
		{
			name: "status: bad request",
			users: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  400,
				user:        User{},
			},
			request: "/users?user_id=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			handler := UserViewHandler(tt.users)
			handler(w, request)
			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			if result.StatusCode == http.StatusOK {
				userResult, err := io.ReadAll(result.Body)
				require.NoError(t, err)
				defer result.Body.Close()

				var user User
				err = json.Unmarshal(userResult, &user)
				require.NoError(t, err)

				assert.Equal(t, tt.want.user, user)
			}
		})
	}
}
