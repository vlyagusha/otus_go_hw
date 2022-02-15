package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
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

	SuccessResponses struct {
		Codes []int `validate:"in:200,201,202,203,204,205,206,207,208,226"`
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
				Phones: []string{"9101234567", "-"},
				meta:   nil,
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   ErrStrLen,
				},
				ValidationError{
					Field: "Age",
					Err:   ErrIntMin,
				},
				ValidationError{
					Field: "Email",
					Err:   ErrStrRegexp,
				},
				ValidationError{
					Field: "Role",
					Err:   ErrStrIn,
				},
				ValidationError{
					Field: "Phones",
					Err:   ErrStrLen,
				},
				ValidationError{
					Field: "Phones",
					Err:   ErrStrLen,
				},
			},
		},
		{
			in:          App{Version: "1.0.0"},
			expectedErr: nil,
		},
		{
			in:          App{Version: "v1.0.0"},
			expectedErr: ValidationErrors{ValidationError{Field: "Version", Err: ErrStrLen}},
		},
		{
			in: Token{
				Header:    nil,
				Payload:   nil,
				Signature: nil,
			},
			expectedErr: nil,
		},
		{
			in: Token{
				Header:    []byte{'1', '2', '3'},
				Payload:   []byte{'a', 'b', 'c'},
				Signature: []byte{'!', '@', '#'},
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 200,
				Body: "OK",
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 201,
				Body: "Created",
			},
			expectedErr: ValidationErrors{ValidationError{Field: "Code", Err: ErrIntIn}},
		},
		{
			in:          SuccessResponses{Codes: []int{200, 201}},
			expectedErr: nil,
		},
		{
			in: SuccessResponses{Codes: []int{400, 404}},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Codes",
					Err:   ErrIntIn,
				},
				ValidationError{
					Field: "Codes",
					Err:   ErrIntIn,
				},
			},
		},
		{
			in: SuccessResponses{Codes: []int{203, 503}},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Codes",
					Err:   ErrIntIn,
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			require.True(t, reflect.DeepEqual(err, tt.expectedErr))
		})
	}
}
