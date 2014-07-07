package filter

import (
	"io/ioutil"

	"github.com/mgutz/goa"
	"github.com/mgutz/gosu"
)

// Load loads all the files from glob patterns and creates the initial
// asset array for a pipeline.
func Load(patterns ...string) func(*goa.Pipeline) error {
	fileAssets, _, _ := gosu.Glob(patterns)
	return func(pipeline *goa.Pipeline) error {
		for _, info := range fileAssets {
			data, err := ioutil.ReadFile(info.Path)
			if err != nil {
				return err
			}
			asset := &goa.Asset{Info: info}
			asset.Write(data)
			asset.WritePath = info.Path
			pipeline.AddAsset(asset)
		}
		return nil
	}
}
