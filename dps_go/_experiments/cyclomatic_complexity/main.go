package main

import "fmt"

func main() {
	fmt.Println("=== Part 1: HTTP status (essential vs accidental) ===")
	codes := []int{102, 200, 301, 404, 503, 999}
	fmt.Println("code | essential | accidental")
	for _, code := range codes {
		e := classifyHTTPEssential(code)
		a := classifyHTTPAccidental(code)
		fmt.Printf("%4d | %-11s | %s\n", code, e, a)
	}

	fmt.Println()
	fmt.Println("=== Part 2: routing handle (before vs after) ===")
	sessions := []Session{
		{Valid: false, MsgType: "request", PeerKnown: true},
		{Valid: true, MsgType: "request", PeerKnown: false},
		{Valid: true, MsgType: "request", PeerKnown: true, Routed: true, HasRouteTable: true},
		{Valid: true, MsgType: "request", PeerKnown: true, HasRouteTable: false},
		{Valid: true, MsgType: "request", PeerKnown: true, HasRouteTable: true},
		{Valid: true, MsgType: "answer", PeerKnown: true},
		{Valid: true, MsgType: "error", PeerKnown: true},
		{Valid: true, MsgType: "unknown", PeerKnown: true},
	}
	fmt.Println("valid type    peer | before     | after")
	for _, s := range sessions {
		b := handleRoutingBefore(s)
		a := handleRoutingAfter(s)
		fmt.Printf("%5v %7s %5v | %-10s | %s\n", s.Valid, s.MsgType, s.PeerKnown, b, a)
	}
}
