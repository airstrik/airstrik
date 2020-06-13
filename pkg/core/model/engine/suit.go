package engine

import "github.com/itsparser/airstrike/pkg/core/model/system"

type Suit struct {
	system.Model
	Name        string
	Description string
	Profile     interface{}
	StartUrl    string
	Version     int64
	SuitId      string
	Commands    []Command `json:"Commands"`
}

type Command struct {
	Type     string
	Value    string
	Variable string
	Xpath    string
}
