package main

// classifyHTTPEssential - essential: many branches match the domain (HTTP classes).
func classifyHTTPEssential(code int) string {
	switch {
	case code >= 100 && code < 200:
		return "informational"
	case code >= 200 && code < 300:
		return "success"
	case code >= 300 && code < 400:
		return "redirection"
	case code >= 400 && code < 500:
		return "client_error"
	case code >= 500 && code < 600:
		return "server_error"
	default:
		return "unknown"
	}
}
