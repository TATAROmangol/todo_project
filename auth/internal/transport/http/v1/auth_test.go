 package v1

import (
	"net/http"
	"testing"
)

func TestAuthHandler_Register(t *testing.T) {
	type fields struct {
		as AuthService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ah := &AuthHandler{
				as: tt.fields.as,
			}
			ah.Register(tt.args.w, tt.args.r)
		})
	}
}
