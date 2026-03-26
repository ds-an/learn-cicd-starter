package auth_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name string
		input http.Header 
		want string
    wantErr bool
	}{
		{
			name: "Happy",
			input: http.Header{
				"Authorization": []string{"ApiKey sdlfksfansldfj"},
			},
			want: "sdlfksfansldfj",
			wantErr: false,
		},
		{
			name: "Empty header",
			input: http.Header{},
			want: "",
			wantErr: true,
		},
	}
    for i, tc := range tests {
        got, gotErr := auth.GetAPIKey(tc.input)
        if (gotErr != nil) != tc.wantErr {
            t.Fatalf("test %d, \"%s\": expected error: %v, got: %v", i+1, tc.name, tc.wantErr, gotErr)
        }
        if !reflect.DeepEqual(tc.want, got) {
            t.Fatalf("test %d, \"%s\": expected: %v, got: %v", i+1, tc.name, tc.want, got)
        }
    }
}
