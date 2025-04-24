//go:build (!cgo || nocgo) && darwin

package fitz

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/ffi"
)

const (
	libname = "libmupdf.dylib"
)

// loadLibrary loads the so and panics on error.
func loadLibrary() uintptr {
	handle, err := purego.Dlopen(libname, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(fmt.Errorf("cannot load library: %w", err))
	}

	return uintptr(handle)
}

func newBundle(name string, rType *ffi.Type, aTypes ...*ffi.Type) *bundle {
	b := new(bundle)
	var err error

	if b.sym, err = purego.Dlsym(libmupdf, name); err != nil {
		panic(err)
	}

	nArgs := uint32(len(aTypes))

	if status := ffi.PrepCif(&b.cif, ffi.DefaultAbi, nArgs, rType, aTypes...); status != ffi.OK {
		panic(status)
	}

	return b
}
