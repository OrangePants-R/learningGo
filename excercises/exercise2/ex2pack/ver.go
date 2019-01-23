package ex2pack

import "runtime"

// Version returns the current version of golang
func Version() string {
	return runtime.Version()
}
