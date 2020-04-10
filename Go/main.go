package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const gridHeight = 20
const gridWidth = 20

func main() {
	http.HandleFunc("/next", getNextStepHttp)
	http.HandleFunc("/init", initHttp)
	_ = http.ListenAndServe(":8090", nil)
}

var g = &Grid{}

func getNextStepHttp(writer http.ResponseWriter, request *http.Request) {
	g = g.NextStep()
	flat := flatten(g)
	bytes, _ := json.Marshal(flat)
	writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
	_, _ = fmt.Fprint(writer, string(bytes))
}

func initHttp(writer http.ResponseWriter, request *http.Request) {
	var v map[string]int
	json.NewDecoder(request.Body).Decode(&v)
	threshold := float32(v["threshold"]) / 100.0
	liveSlotMinLiveNeighboursToKeepAlive := v["liveSlotMinLiveNeighboursToKeepAlive"]
	liveSlotMaxLiveNeighboursToKeepAlive := v["liveSlotMaxLiveNeighboursToKeepAlive"]
	deadSlotMinLiveNeighboursToBringToLife := v["deadSlotMinLiveNeighboursToBringToLife"]
	deadSlotMaxLiveNeighboursToBringToLife := v["deadSlotMaxLiveNeighboursToBringToLife"]
	g = InitializeGrid(
		gridHeight,
		gridWidth,
		threshold,
		liveSlotMinLiveNeighboursToKeepAlive,
		liveSlotMaxLiveNeighboursToKeepAlive,
		deadSlotMinLiveNeighboursToBringToLife,
		deadSlotMaxLiveNeighboursToBringToLife)
	flat := flatten(g)
	bytes, _ := json.Marshal(flat)
	writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
	_, _ = fmt.Fprint(writer, string(bytes))
}
