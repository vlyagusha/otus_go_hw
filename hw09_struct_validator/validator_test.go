package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          200,
			expectedErr: nil,
		},
		{
			in:          "HTTP 200 OK",
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "ff829198-8db4-11ec-b909-0242ac120002",
				Name:   "John",
				Age:    30,
				Email:  "mail@example.com",
				Role:   "admin",
				Phones: []string{"79101234567"},
				meta:   nil,
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "ff829198-8db4-11ec-b909",
				Name:   "John",
				Age:    17,
				Email:  "mailexample.com",
				Role:   "admin2",
				Phones: []string{"79101234567"},
				meta:   nil,
			},
			expectedErr: ValidationErrors{ValidationError{
				Field: "ID",
				Err:   nil,
			}},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			// require.True(t, errors.Is(err, tt.expectedErr))
			fmt.Println(err)
		})
	}
}
