package leonardo

import (
	"net/url"
)

// Helper function to escape URL path segments
func urlPathEscape(s string) string {
	return url.PathEscape(s)
}
