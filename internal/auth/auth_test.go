package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "no auth header",
			headers: http.Header{},
			//want:    "",
			want:    "123",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "malformed auth header",
			headers: http.Header{"Authorization": []string{"Bearer"}},
			//want:    "",
			want:    "123",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name:    "valid auth header",
			headers: http.Header{"Authorization": []string{"ApiKey 123"}},
			want:    "123",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetAPIKey(tt.headers)
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
