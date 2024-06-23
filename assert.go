package particles

import (
	"log"
	"log/slog"
)

var asserData map[string]any = map[string]any{}

func runAssert(msg string) {
	for k, v := range asserData {
		slog.Error("context", "key", k, "value", v)
	}
	log.Fatal(msg)
}

func Assert(msg string, expr bool) {
	if !expr {
		runAssert(msg)
	}
}
