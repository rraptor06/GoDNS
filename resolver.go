package resolver

func ResolveDomain(domain string) string {
	if domain == "example.com" {
		return "93.184.216.34" // IP Example
	}
	return "0.0.0.0"
}
