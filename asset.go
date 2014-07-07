package goa

import (
	"bytes"
	"path"

	"github.com/mgutz/gosu"
)

// Asset is any file which can be loaded and processed by a filter.
type Asset struct {
	bytes.Buffer
	Info *gosu.FileAsset
	// WritePath is the write destination of the asset.
	WritePath string
	Pipeline  *Pipeline
}

// ChangeExt changes the extesion of asset.WritePath. ChangExt is used
// by filters which transpile source. For example, a filter for Markdown
// would use ChangeExt(".html") to write the asset as an HTML file.
func (asset *Asset) ChangeExt(newExt string) {
	extension := path.Ext(asset.WritePath)
	base := asset.WritePath[0:len(asset.WritePath)-len(extension)] + newExt
	asset.WritePath = path.Join(path.Dir(asset.WritePath), base)
}

// Ext returns the extension of asset.WritePath.
func (asset *Asset) Ext() string {
	return path.Ext(asset.WritePath)
}
