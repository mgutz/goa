package filter

import (
	"io/ioutil"

	"github.com/mgutz/goa"
	"gopkg.in/godo.v1"
)

// Load loads all the files from glob patterns and creates the initial
// asset array for a pipeline. This loads the entire contents of the file, binary
// or text, into a buffer. Consider creating your own loader if dealing
// with large files.
func Load(patterns ...string) func(*goa.Pipeline) error {
	return func(pipeline *goa.Pipeline) error {
		fileAssets, _, err := godo.Glob(patterns)
		if err != nil {
			return err
		}

		for _, info := range fileAssets {
			if !info.IsDir() {
				data, err := ioutil.ReadFile(info.Path)
				if err != nil {
					return err
				}
				asset := &goa.Asset{Info: info}
				asset.Write(data)
				asset.WritePath = info.Path
				pipeline.AddAsset(asset)
			}
		}
		return nil
	}
}
