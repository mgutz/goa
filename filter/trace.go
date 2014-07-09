package filter

import (
	"github.com/mgutz/goa"
	"github.com/mgutz/gosu/util"
)

// Trace traces an asset, printing key properties of asset to the console.
func Trace() func(*goa.Asset) error {
	return func(asset *goa.Asset) error {
		util.Debug("filter", asset.Dump())
		return nil
	}
}
