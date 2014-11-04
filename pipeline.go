package goa

import (
	"gopkg.in/godo.v1/util"
)

// Verbose indicates whether to log verbosely
var Verbose = false

// Pipeline is a asset flow through which each asset is processed by
// one or more filters. For text
// files this could be something as simple as adding a header or
// minification. Some filters process assets in batches combining them,
// for example concatenating JavaScript or CSS.
type Pipeline struct {
	Assets  []*Asset
	Filters []interface{}
}

// Pipe creates a pipeline with filters and runs it.
func Pipe(filters ...interface{}) (*Pipeline, error) {
	pipeline := &Pipeline{Assets: []*Asset{}}
	pipeline.Pipe(filters...).Run()
	return pipeline, nil
}

// Pipe adds one or more filters to the pipeline. Pipe may be called
// more than once.
//
// Filters are simple function. Options are handle through closures.
// The supported handlers are
//
// 1. Single asset handler. Use this to transorm each asset individually.
//    AddHeader filter is an example.
//
//      // signature
//      func(*goa.Asset) error
//
// 2. Multi asset handler. Does not modify the number of elements. See
//    Write filter is an example.
//
//      //  signature
//      func(assets []*goa.Asset) error
//
// 3. Pipeline handler. Use this to have unbridled control. Load filter
//    is an example.
//
//      // signature
//      func(*Pipeline) error
//
func (pipeline *Pipeline) Pipe(filters ...interface{}) *Pipeline {
	pipeline.Filters = filters
	// for _, filter := range filters {
	// 	pipeline.Filters = append(pipeline.Filters, filter)
	// }
	return pipeline
}

// Run runs assets through the pipeline.
func (pipeline *Pipeline) Run() {
	for i, filter := range pipeline.Filters {
		if i == 1 && len(pipeline.Assets) == 0 {
			util.Info("goa", "There are 0 assets in pipeline. Check your Load filter. %+v\n", pipeline)
		}
		switch fn := filter.(type) {
		default:
			util.Panic("pipeline", "Invalid filter type %+v\n", fn)
		// receives a single asset, a filter
		case func(*Asset) error:
			for _, asset := range pipeline.Assets {
				err := fn(asset)
				if err != nil {
					util.Error("goa", "%+v\n", err)
					break
				}
			}
		// This should only be used inspections like tap, saves from having to return an error
		case func(*Asset):
			for _, asset := range pipeline.Assets {
				fn(asset)
			}
		// receives all assets read-only
		case func([]*Asset) error:
			err := fn(pipeline.Assets)
			if err != nil {
				util.Error("goa", "%+v\n", err)
				break
			}
		// receives the pipeline, can add remove assets
		case func(*Pipeline) error:
			err := fn(pipeline)
			if err != nil {
				util.Error("goa", "%+v\n", err)
				break
			}
		}
	}
}

// AddAsset adds an asset
func (pipeline *Pipeline) AddAsset(asset *Asset) {
	if asset == nil {
		return
	}
	asset.Pipeline = pipeline
	pipeline.Assets = append(pipeline.Assets, asset)
}

// Truncate removes all assets, resetting Assets to empty slice.
func (pipeline *Pipeline) Truncate() {
	pipeline.Assets = pipeline.Assets[:0]
}
