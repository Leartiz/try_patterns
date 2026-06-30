package main

// handleRoutingBefore - nested ifs, duplicated checks (high CC).
func handleRoutingBefore(s Session) RouteResult {
	if s.Valid {
		if s.MsgType == "request" {
			if s.PeerKnown {
				if s.Routed {
					return RouteForward
				} else {
					if s.HasRouteTable {
						return RouteOK
					} else {
						return RouteSkip
					}
				}
			} else {
				return RouteUnknownPeer
			}
		} else {
			if s.MsgType == "answer" {
				if s.PeerKnown {
					return RouteOK
				} else {
					return RouteUnknownPeer
				}
			} else {
				if s.MsgType == "error" {
					return RouteSkip
				}
			}
		}
	} else {
		return RouteInvalid
	}
	return RouteSkip
}
