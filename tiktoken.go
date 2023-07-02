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

type USize interface {
    uint32 | uint64
}

type CoreBPE struct {
	bpe unsafe.Pointer
}

var _ io.Closer = (*CoreBPE)(nil)

func FromModel(model string) (*CoreBPE, error) {
	cModel := C.CString(model)
	defer C.free(unsafe.Pointer(cModel))

	bpe := C.get_bpe_from_model(cModel)
	return &CoreBPE{bpe: bpe}, nil
}

func (b *CoreBPE) Close() error {
	C.free_bpe(b.bpe)
	b.bpe = nil
	return nil
}

func (b *CoreBPE) Encode(prompt string) []uint32 {
	cPrompt := C.CString(prompt)
	defer C.free(unsafe.Pointer(cPrompt))

	var len C.uint
	res := C.encode(b.bpe, cPrompt, &len)
	if len > 0 {
		// can't dealloc nil
		defer C.free(unsafe.Pointer(res))
	}
	slice := unsafe.Slice(res, len)
	tokenIDs := make([]uint32, len)
	for i, v := range slice {
		tokenIDs[i] = uint32(v)
	}

	return tokenIDs
}