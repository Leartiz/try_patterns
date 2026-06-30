package main

type Session struct {
	Valid         bool
	MsgType       string // request, answer, error
	PeerKnown     bool
	Routed        bool
	HasRouteTable bool
}

type RouteResult string

const (
	RouteOK          RouteResult = "ok"
	RouteInvalid     RouteResult = "invalid"
	RouteUnknownPeer RouteResult = "unknown_peer"
	RouteSkip        RouteResult = "skip"
	RouteForward     RouteResult = "forward"
)
