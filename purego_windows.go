//go:build (!cgo || nocgo) && windows

package fitz

import (
	"fmt"
	"syscall"

	"github.com/jupiterrider/ffi"
	"golang.org/x/sys/windows"
)

const (
	libname = "libmupdf.dll"
)

// loadLibrary loads the dll and panics on error.
func loadLibrary() uintptr {
	handle, err := syscall.LoadLibrary(libname)
	if err != nil {
		panic(fmt.Errorf("cannot load library %s: %w", libname, err))
	}

	return uintptr(handle)
}

func newBundle(name string, rType *ffi.Type, aTypes ...*ffi.Type) *bundle {
	b := new(bundle)
	var err error

	var h windows.Handle
	if h, err = windows.GetStdHandle(uint32(libmupdf)); err != nil {
		panic(err)
	}
	if b.sym, err = windows.GetProcAddress(h, name); err != nil {
		panic(err)
	}

	nArgs := uint32(len(aTypes))

	if status := ffi.PrepCif(&b.cif, ffi.DefaultAbi, nArgs, rType, aTypes...); status != ffi.OK {
		panic(status)
	}

	return b
}
