package filter

import (
	"github.com/mgutz/goa"
	"os"
	"path/filepath"
)

// Write writes all assets to the file system.
func Write() func(assets []*goa.Asset) error {
	return func(assets []*goa.Asset) error {
		// cache directories to avoid Mkdirall
		madeDirs := map[string]bool{}
		for _, asset := range assets {
			writePath := asset.WritePath
			dir := filepath.Dir(writePath)
			if !madeDirs[dir] {
				err := os.MkdirAll(dir, 0755)
				if err != nil {
					return err
				}
				madeDirs[dir] = true
			}
			f, err := os.Create(writePath)
			if err != nil {
				return err
			}
			defer f.Close()
			asset.Buffer.WriteTo(f)
		}
		return nil
	}
}
