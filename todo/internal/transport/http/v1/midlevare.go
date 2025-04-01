package v1

import (
	"context"
	"net/http"
)

func Log(ctx context.Context, h func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		h(w,r)
	}
}