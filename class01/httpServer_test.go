package main

import (
	"net/http"
	"testing"
)

func testHtteServer(t *testing.T) {
	t.Log("xxx")
}

func Test_healthzHandler(t *testing.T) {
	type args struct {
		writer http.ResponseWriter
		req    *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			healthzHandler(tt.args.writer, tt.args.req)
		})
	}
}
