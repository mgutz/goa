package filter

import (
	"strings"

	"github.com/mgutz/goa"
	"gopkg.in/godo.v1/util"
)

// ReplacePath replaces the leading part of a path in all assets.
//
//      ReplacePath("src/", "dist/")
//
// This should be used before the Write() filter.
func ReplacePath(from string, to string) func(*goa.Asset) error {
	return func(asset *goa.Asset) error {
		oldPath := asset.WritePath
		if !strings.HasPrefix(oldPath, from) {
			return nil
		}
		asset.WritePath = to + oldPath[len(from):]
		if Verbose {
			util.Debug("goa", "ReplacePath %s => %s\n", oldPath, asset.WritePath)
		}
		return nil
	}
}
