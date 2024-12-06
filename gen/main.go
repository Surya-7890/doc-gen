package gen

import (
	"log"
)

type Gen struct {
	Log      *log.Logger
	DataChan chan struct{} // change to appropriate data type
}

// waits for route data
// calls constructJsonFile function when DataChan is closed
func WaitForData() {}
