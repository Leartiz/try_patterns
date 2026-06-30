package main

// classifyHTTPAccidental - accidental: duplicated range checks and redundant branches.
func classifyHTTPAccidental(code int) string {
	if code < 100 {
		return "unknown"
	}
	if code < 200 {
		if code >= 100 {
			return "informational"
		}
		return "unknown"
	}
	if code < 300 {
		if code >= 200 {
			return "success"
		}
		return "unknown"
	}
	if code < 400 {
		if code >= 300 {
			return "redirection"
		}
		return "unknown"
	}
	if code < 500 {
		if code >= 400 {
			return "client_error"
		}
		return "unknown"
	}
	if code < 600 {
		if code >= 500 {
			return "server_error"
		}
		return "unknown"
	}
	if code >= 600 {
		return "unknown"
	}
	return "unknown"
}
