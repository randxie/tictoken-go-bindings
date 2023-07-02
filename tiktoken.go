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
	"errors"
	"io"
	"unsafe"
)

type CoreBPE struct {
	bpe unsafe.Pointer
}

// Equivalent to "impl ... for" in Rust.
var _ io.Closer = (*CoreBPE)(nil)

func FromModel(model string) (*CoreBPE, error) {
	cModel := C.CString(model)
	defer C.free(unsafe.Pointer(cModel))

	bpe := C.get_bpe_from_model(cModel)
	if bpe == nil {
		return nil, errors.New("failed to load tokenizer for model: " + model)
	}
	return &CoreBPE{bpe: bpe}, nil
}

func (e *CoreBPE) Close() error {
	C.free_bpe(e.bpe)
	e.bpe = nil
	return nil
}

func (e *CoreBPE) Encode(prompt string) []uint32 {
	cPrompt := C.CString(prompt)
	defer C.free(unsafe.Pointer(cPrompt))

	var len C.uint
	res := C.encode(e.bpe, cPrompt, &len)
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

func (e *CoreBPE) Decode(tokenIDs []uint32) string {
	if len(tokenIDs) == 0 {
		return ""
	}
	len := C.uint(len(tokenIDs))
	output := C.decode(e.bpe, (*C.uint)(unsafe.Pointer(&tokenIDs[0])), len)
	defer C.free(unsafe.Pointer(output))
	return C.GoString(output)
}
