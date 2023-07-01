package tiktoken

// NOTE: There should be NO space between the comments and the `import "C"` line.

/*
#cgo LDFLAGS: ${SRCDIR}/libtiktoken_ffi.a -ldl -lstdc++
#include <stdlib.h>
#include "tiktoken.h"
*/
import "C"

// NOTE: There should be NO space between the comments and the `import "C"` line.
import (
	"io"
	"unsafe"
)

type CoreBPE struct {
	bpe unsafe.Pointer
}

var _ io.Closer = (*CoreBPE)(nil)

func FromModel(model string) (*CoreBPE, error) {
	cModel := C.CString(model)
	bpe := C.get_bpe_from_model(cModel)
	return &CoreBPE{bpe: bpe}, nil
}

func (b *CoreBPE) Close() error {
	C.free_bpe(b.bpe)
	b.bpe = nil
	return nil
}