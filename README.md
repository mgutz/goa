# goa

Go asset pipeline.

## Example

```go
import (
    "fmt"

    "github.com/mgutz/goa"
    f "github.com/mgutz/goa/filter"
    "github.com/mgutz/gosu"
)

func Project(p *gosu.Project) {
    p.Task("default", []string{"add-headers", "bundle-scripts"})

    p.Task("add-headers", func() {
        goa.Pipe(
            f.Load("test/**/*.txt"),
            f.AddHeader("COPYRIGHT\n"),
            f.ReplacePath("test", "dist"),
            // tap into it
            func(asset *goa.Asset) {
                fmt.Printf("WritePath %s MimeType %s\n", asset.WritePath, asset.MimeType())
            },
            f.Write(),
        )
    })

    p.Task("bundle-scripts", func() {
        goa.Pipe(
            f.Load("test/**/*.js"),
            f.Cat(";", "dist/bundle.js"),
            f.AddHeader("COPYRIGHT\n"),
            f.Write(),
        )
    })
}
```
