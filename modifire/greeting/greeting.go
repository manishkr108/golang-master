package greeting

// This function is public (exported)
func Hello() string {
	return "hello"
}

// This function is private (unexported)
func perfect() string {
	return "perfact not export"
}
