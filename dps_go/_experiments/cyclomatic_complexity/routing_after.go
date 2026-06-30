package main

// handleRoutingAfter - same rules, lower nesting (lower CC).
func handleRoutingAfter(s Session) RouteResult {
	if !s.Valid { // guard clause: fail fast on invalid session
		return RouteInvalid
	}

	switch s.MsgType { // switch: replace message-type if-chain
	case "request":
		if !s.PeerKnown { // guard clause: request needs known peer
			return RouteUnknownPeer
		}
		return routeRequest(s) // extract function: isolate request flow
	case "answer":
		if !s.PeerKnown { // guard clause: answer needs known peer
			return RouteUnknownPeer
		}
		return RouteOK
	case "error":
		return RouteSkip
	default:
		return RouteSkip
	}
}

func routeRequest(s Session) RouteResult {
	if s.Routed { // guard clause: already routed -> forward
		return RouteForward
	}
	if !s.HasRouteTable { // guard clause: no table -> skip
		return RouteSkip
	}
	return RouteOK
}
