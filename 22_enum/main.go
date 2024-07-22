package main

import "fmt"

const (
	StateIdle = iota
	StateConected
	StateError
	StateRetrying
)

var stateName