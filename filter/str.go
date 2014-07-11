package filter

import (
	"github.com/mgutz/goa"
)

// Str passes asset.Buffer string through any `str` filter for processing.
// asset.Buffer is assigned then asigned the result value from filter.
func Str(handler func(string) string) func(*goa.Asset) error {
	return func(asst *goa.Asset) error {
		if asst.IsText() {
			s := asst.String()
			asst.RewriteString(handler(s))
		}
		return nil
	}
}
