// Package goa passes file assets through a pipeline, in which each asset
// is processed by a one or more filters.
//
// Goa can be used in any project but it synergizes well with Godo.
//
//      p.Task("add-copyright", function() {
//          pi := goa.NewPipeline()
//          pi.Pipe(
//              Load("./**/*.go"),
//              AddHeader("Copyright 2014 Mario Gutierrez\n"),
//              Write(),
//          ).Run()
//      })
package goa
