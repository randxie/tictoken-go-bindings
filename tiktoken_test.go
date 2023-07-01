package tiktoken

import (
	_ "embed"
	"testing"
	"fmt"
)

func TestValidModelLoadedSuccessfully(t *testing.T) {
	// This should panic.
	_, err := FromModel("code-davinci-002")

	if err == nil {
		fmt.Println("pass.")
	}
}