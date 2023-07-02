package tiktoken

import (
	_ "embed"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestValidModelLoadedSuccessfully(t *testing.T) {
	// This should panic.
	_, err := FromModel("code-davinci-002")
	require.Nil(t, err)
}

func TestEncodeCorrectly(t *testing.T) {
	bpe, err := FromModel("gpt-3.5-turbo-16k")
	require.Nil(t, err, "Encoding  init should not be nil")
	actualTokens := bpe.Encode("hello world!你好，世界！")
	// these tokens are converted from the original python code
	expectedTokens := []uint32{15339, 1917, 0, 57668, 53901, 3922, 3574, 244, 98220, 6447}
	require.Equal(t, actualTokens, expectedTokens, "Encoding should be equal")
}