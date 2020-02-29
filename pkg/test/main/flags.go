package main

import "time"

// Flags struct defined here
type Flags struct {
	Debug bool

	Port         uint
	GatewayPort  uint
	UpstreamPort uint
	BasePort     uint
	AlsPort      uint

	Delay    time.Duration
	Requests int
	Updates  int

	Mode          string
	Clusters      int
	HTTPListeners int
	TCPListeners  int
	Runtimes      int
	TLS           bool

	NodeID string

	V2 bool
	V3 bool
}
