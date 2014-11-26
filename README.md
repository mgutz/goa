# goa

[godoc](https://godoc.org/github.com/mgutz/goa)

    import "github.com/mgutz/goa"

Package goa passes file assets through a pipeline, in which each asset is
processed by a one or more filters.

Goa can be used in any project but it synergizes well with Godo.

    p.Task("add-copyright", function() {
        pi := goa.NewPipeline()
        pi.Pipe(
            Load("./**/*.go"),
            AddHeader("Copyright 2014 Mario Gutierrez\n"),
            Write(),
        ).Run()
    })

## Usage

```go
var Verbose = false
```

#### type Asset

```go
type Asset struct {
	bytes.Buffer
	Info *godo.FileAsset
	// WritePath is the write destination of the asset.
	WritePath string
	Pipeline  *Pipeline
}
```

Asset is any file which can be loaded and processed by a filter.

#### func (*Asset) ChangeExt

```go
func (asset *Asset) ChangeExt(newExt string)
```
ChangeExt changes the extension of asset.WritePath. ChangExt is used by filters
which transpile source. For example, a filter for Markdown would use
ChangeExt(".html") to write the asset as an HTML file.

#### func (*Asset) Dump

```go
func (asset *Asset) Dump() string
```
Dump returns a console friendly representation of asset. Note, String() returns
the string value of Buffer.

#### func (*Asset) Ext

```go
func (asset *Asset) Ext() string
```
Ext returns the extension of asset.WritePath.

#### func (*Asset) IsText

```go
func (asset *Asset) IsText() bool
```
IsText return true if it thinks this asset is text based, meaning it can be
manipulated with string functions.

#### func (*Asset) MimeType

```go
func (asset *Asset) MimeType() string
```
MimeType returns an educated guess of the content type of asset.

#### func (*Asset) Rewrite

```go
func (asset *Asset) Rewrite(bytes []byte)
```
Rewrite sets the buffer to bytes.

#### func (*Asset) RewriteString

```go
func (asset *Asset) RewriteString(s string)
```
RewriteString sets the buffer to a string value.

#### type Pipeline

```go
type Pipeline struct {
	Assets  []*Asset
	Filters []interface{}
}
```

Pipeline is a asset flow through which each asset is processed by one or more
filters. For text files this could be something as simple as adding a header or
minification. Some filters process assets in batches combining them, for example
concatenating JavaScript or CSS.

#### func  Pipe

```go
func Pipe(filters ...interface{}) (*Pipeline, error)
```
Pipe creates a pipeline with filters and runs it.

#### func (*Pipeline) AddAsset

```go
func (pipeline *Pipeline) AddAsset(asset *Asset)
```
AddAsset adds an asset

#### func (*Pipeline) Pipe

```go
func (pipeline *Pipeline) Pipe(filters ...interface{}) *Pipeline
```
Pipe adds one or more filters to the pipeline. Pipe may be called more than
once.

Filters are simple function. Options are handle through closures. The supported
handlers are

1. Single asset handler. Use this to transorm each asset individually.

    AddHeader filter is an example.

      // signature
      func(*goa.Asset) error

2. Multi asset handler. Does not modify the number of elements. See

    Write filter is an example.

      //  signature
      func(assets []*goa.Asset) error

3. Pipeline handler. Use this to have unbridled control. Load filter

    is an example.

      // signature
      func(*Pipeline) error

#### func (*Pipeline) Run

```go
func (pipeline *Pipeline) Run()
```
Run runs assets through the pipeline.

#### func (*Pipeline) Truncate

```go
func (pipeline *Pipeline) Truncate()
```
Truncate removes all assets, resetting Assets to empty slice.
