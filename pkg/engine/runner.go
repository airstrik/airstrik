package engine

import (
	"fmt"
	"github.com/itsparser/airstrike/pkg/core/model/engine"
	"github.com/itsparser/airstrike/pkg/core/model/system"
)

func ProcessSuit(suit engine.Suit) {
	fmt.Print("Runner is Initialized for running Test Suit Id : "+suit.SuitId+" Version ", suit.Version)
	testSuit := &TestSuit{TestSuitId: suit.SuitId, Version: suit.Version, Status: "PENNING", suit: &suit}
	testSuit.start()
}

type TestSuit struct {
	system.Model
	TestSuitId string
	Version    int64
	Status     string
	report     TestReport
	suit       *engine.Suit
}

type TestReport struct {
	Total   int
	Success int
	Failed  int
	Running int
	status  string // Running, Hold , Stop,
}

func (ts *TestSuit) start() {
	ts.updateReport()
	ts.run()

}

func (ts *TestSuit) updateReport() {

}

func (ts *TestSuit) run() {
	browser := NewBrowser("")
	if ts.suit.StartUrl != "" {
		browser.Navigate(ts.suit.StartUrl)
	}

}
