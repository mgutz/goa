package goa

import (
	"github.com/mgutz/gosu"
)

// Pipeline is a asset flow through which each asset is processed by
// one or more filters. For text
// files this could be something as simple as adding a header or
// minification. Some filters process assets in batches combining them,
// for example concatenating JavaScript or CSS.
type Pipeline struct {
	Assets  []*Asset
	Filters []interface{}
}

// NewPipeline creates a new pipeline with empty assets.
func NewPipeline() *Pipeline {
	return &Pipeline{Assets: []*Asset{}}
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
//
// 3. Pipeline handler. Use this to have unbridled control. Load filter
//    is an example.
//
//      // signature
//      func(*Pipeline) error
//
func (pipeline *Pipeline) Pipe(filters ...interface{}) *Pipeline {
	for _, filter := range filters {
		pipeline.Filters = append(pipeline.Filters, filter)
	}
	return pipeline
}

// Run runs assets through the pipeline.
func (pipeline *Pipeline) Run() {
	for _, filter := range pipeline.Filters {
		switch fn := filter.(type) {
		default:
			gosu.Panicf("pipeline", "Invalid filter type %+v\n", fn)
		// receives a single asset, a filter
		case func(*Asset) error:
			for _, asset := range pipeline.Assets {
				fn(asset)
			}
		// receives all assets read-only
		case func([]*Asset) error:
			fn(pipeline.Assets)
		// receives the pipeline, can add remove assets
		case func(*Pipeline) error:
			fn(pipeline)
		}
		pipeline.Filters = append(pipeline.Filters, filter)
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
