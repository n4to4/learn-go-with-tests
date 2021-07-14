package server

import (
	"fmt"
	"net/http"
)

func PlayerServer(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "20")
}
