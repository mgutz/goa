package filter

import (
	"bytes"
	"github.com/mgutz/goa"
)

// AddHeader prepends header to each asset's buffer.
func AddHeader(header string) func(*goa.Asset) error {
	// TODO need way to determine if a file is binary or text
	return func(asset *goa.Asset) error {
		buffer := bytes.NewBufferString(header)
		buffer.Write(asset.Bytes())
		asset.Buffer = *buffer
		return nil
	}
}
