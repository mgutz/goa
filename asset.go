package goa

import (
	"bytes"
	"fmt"
	"net/http"
	"path"

	"gopkg.in/godo.v1"
)

// Asset is any file which can be loaded and processed by a filter.
type Asset struct {
	bytes.Buffer
	Info *godo.FileAsset
	// WritePath is the write destination of the asset.
	WritePath string
	Pipeline  *Pipeline
	mimeType  string
}

// ChangeExt changes the extension of asset.WritePath. ChangExt is used
// by filters which transpile source. For example, a filter for Markdown
// would use ChangeExt(".html") to write the asset as an HTML file.
func (asset *Asset) ChangeExt(newExt string) {
	extension := path.Ext(asset.WritePath)
	base := asset.WritePath[0:len(asset.WritePath)-len(extension)] + newExt
	asset.WritePath = path.Join(path.Dir(asset.WritePath), base)
}

// Dump returns a console friendly representation of asset. Note, String()
// returns the string value of Buffer.
func (asset *Asset) Dump() string {
	return fmt.Sprintf("ReadPath: \"%s\" WritePath: \"%s\" MimeType: \"%s\"\n", asset.Info.Path, asset.WritePath, asset.MimeType())
}

// Ext returns the extension of asset.WritePath.
func (asset *Asset) Ext() string {
	return path.Ext(asset.WritePath)
}

// IsText return true if it thinks this asset is text based, meaning it
// can be manipulated with string functions.
func (asset *Asset) IsText() bool {
	mimeType := asset.MimeType()
	// TODO more mimes need to be checked
	switch mimeType {
	default:
		return true
	case "application/octet-stream":
		return false
	}
}

// MimeType returns an educated guess of the content type of asset.
func (asset *Asset) MimeType() string {
	if asset.mimeType == "" {
		// TODO is passing all bytes expensive?
		asset.mimeType = http.DetectContentType(asset.Buffer.Bytes())
	}
	return asset.mimeType
}

// RewriteString sets the buffer to a string value.
func (asset *Asset) RewriteString(s string) {
	asset.Reset()
	asset.WriteString(s)
}

// Rewrite sets the buffer to bytes.
func (asset *Asset) Rewrite(bytes []byte) {
	asset.Reset()
	asset.Write(bytes)
}
