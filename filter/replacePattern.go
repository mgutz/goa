package filter

import (
	"regexp"

	"github.com/mgutz/goa"
	"github.com/mgutz/gosu/util"
)

// ReplacePath replaces the leading part of a path in all assets.
//
//      ReplacePath("views/", "dist/views")
//
// This should be used before the Write() filter.
func ReplacePattern(pattern, repl string) func(*goa.Asset) error {
	re := regexp.MustCompile(pattern)
	return func(asset *goa.Asset) error {
		if asset.IsText() {
			s := asset.String()
			if s != "" {
				asset.RewriteString(re.ReplaceAllString(s, repl))
			}
		} else {
			util.Debug("dbg", "Asset is not text, mimeType %s\n", asset.MimeType())
		}
		return nil
	}
}
