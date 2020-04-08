package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var g = CreateRandomGrid(10, 10, 0.3)

func main() {
	http.HandleFunc("/next", getNextStepHttp)
	http.HandleFunc("/init", initHttp)
	_ = http.ListenAndServe(":8090", nil)
}

func getNextStepHttp(writer http.ResponseWriter, request *http.Request) {
	g = g.NextStep()
	flat := flatten(g)
	bytes, _ := json.Marshal(flat)
	writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
	_, _ = fmt.Fprint(writer, string(bytes))
}

func initHttp(writer http.ResponseWriter, request *http.Request) {

	g = CreateRandomGrid(10, 10, 0.3)
	flat := flatten(g)
	bytes, _ := json.Marshal(flat)
	writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
	_, _ = fmt.Fprint(writer, string(bytes))
}
