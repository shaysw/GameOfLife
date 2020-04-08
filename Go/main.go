package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const gridHeight = 20
const gridWidth = 20
const threshold = 0.3

var g = CreateRandomGrid(gridHeight, gridWidth, threshold)

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

	g = CreateRandomGrid(gridHeight, gridWidth, threshold)
	flat := flatten(g)
	bytes, _ := json.Marshal(flat)
	writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
	_, _ = fmt.Fprint(writer, string(bytes))
}
