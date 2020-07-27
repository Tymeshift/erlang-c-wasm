package main

import (
	"syscall/js"

	erlangc "github.com/Tymeshift/erlang-c-go"
)

// CalculateFTEWasm - wrapper for erlangc calculations for WASM
func CalculateFTEWasm(this js.Value, p []js.Value) interface{} {
	var fteParams []erlangc.FteParams
	for i := 0; i < len(p); i += 7 {
		params := erlangc.FteParams{
			Volume:             float64(p[i].Float()),
			IntervalLength:     int64(p[i+1].Int()),
			MaxOccupancy:       float64(p[i+2].Float()),
			Shrinkage:          float64(p[i+3].Float()),
			Aht:                int64(p[i+4].Int()),
			TargetServiceLevel: float64(p[i+5].Float()),
			TargetTime:         int64(p[i+6].Int()),
		}
		fteParams = append(fteParams, params)
	}
	return erlangc.CalculateFte(fteParams)
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("CalculateFTEWasm", js.FuncOf(CalculateFTEWasm))

	<-c
}
