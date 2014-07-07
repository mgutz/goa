package filter

import (
	"github.com/mgutz/goa"
	"strings"
)

// ReplacePath replaces the leading part of a path in all assets.
//
//      ReplacePath("views/", "dist/views")
//
// This should be used before the Write() filter.
func ReplacePath(from string, to string) func(*goa.Asset) error {
	return func(asset *goa.Asset) error {
		s := asset.WritePath
		if !strings.HasPrefix(s, from) {
			return nil
		}
		s = to + s[len(from):]
		asset.WritePath = s
		return nil
	}
}
