package utils

import (
	"go/build"
	"path/filepath"
)

// GoPath returns the current GOPATH. If the system
// has multiple GOPATHs then the first is used.
func GoPath() string {
	gpDefault := build.Default.GOPATH
	gps := filepath.SplitList(gpDefault)
	return gps[0]
}
