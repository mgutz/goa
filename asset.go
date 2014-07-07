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
	// WritePath is the path to where the asset is written.
	WritePath string
	Pipeline  *Pipeline
}

// Filter processes a single asset.
type Filter interface {
	Filter(*Asset)
}

// Batcher process one or more assets and can modify the collection
// of assets. Some example of batch operations are concatenators,
// combiners, archivers.
type Batcher interface {
	Batch([]*Asset)
}

// ChangeExt changes the extesion of asset.WritePath.
func (asset *Asset) ChangeExt(newExt string) {
	extension := path.Ext(asset.WritePath)
	base := asset.WritePath[0:len(asset.WritePath)-len(extension)] + newExt
	asset.WritePath = path.Join(path.Dir(asset.WritePath), base)
}

// ChangeExt changes the extesion of asset.WritePath.
func (asset *Asset) Ext() string {
	return path.Ext(asset.WritePath)
}
