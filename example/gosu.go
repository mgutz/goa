package main

import (
	"github.com/mgutz/goa"
	f "github.com/mgutz/goa/filter"
	"github.com/mgutz/gosu"
)

func Project(p *gosu.Project) {
	p.Task("default", []string{"add-headers", "bundle-scripts"})

	p.Task("add-headers", func() {
		pipeline := goa.NewPipeline()
		pipeline.Pipe(
			f.Load("test/**/*.txt"),
			f.AddHeader("COPYRIGHT\n"),
			f.ReplacePath("test", "dist"),
			f.Write(),
		).Run()
	})

	p.Task("bundle-scripts", func() {
		pipeline := goa.NewPipeline()
		pipeline.Pipe(
			f.Load("test/**/*.js"),
			f.Cat(";", "dist/bundle.js"),
			f.AddHeader("COPYRIGHT\n"),
			f.ReplacePath("test", "dist"),
			f.Write(),
		).Run()
	})
}

func main() {
	gosu.Run(Project)
}
