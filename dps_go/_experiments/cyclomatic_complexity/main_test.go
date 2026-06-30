package main

import "testing"

func Test_classifyHTTP_sameResult(t *testing.T) {
	codes := []int{0, 102, 200, 301, 404, 503, 999}

	for _, code := range codes {
		gotE := classifyHTTPEssential(code)
		gotA := classifyHTTPAccidental(code)
		if gotE != gotA {
			t.Errorf("code %d: essential=%q accidental=%q", code, gotE, gotA)
		}
	}
}

func Test_handleRouting_sameResult(t *testing.T) {
	sessions := []Session{
		{Valid: false, MsgType: "request", PeerKnown: true},
		{Valid: true, MsgType: "request", PeerKnown: false},
		{Valid: true, MsgType: "request", PeerKnown: true, Routed: true, HasRouteTable: true},
		{Valid: true, MsgType: "request", PeerKnown: true, HasRouteTable: false},
		{Valid: true, MsgType: "request", PeerKnown: true, HasRouteTable: true},
		{Valid: true, MsgType: "answer", PeerKnown: true},
		{Valid: true, MsgType: "answer", PeerKnown: false},
		{Valid: true, MsgType: "error", PeerKnown: true},
		{Valid: true, MsgType: "unknown", PeerKnown: true},
	}

	for i, s := range sessions {
		gotB := handleRoutingBefore(s)
		gotA := handleRoutingAfter(s)
		if gotB != gotA {
			t.Errorf("session %d: before=%q after=%q", i, gotB, gotA)
		}
	}
}
