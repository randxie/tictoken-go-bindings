package tiktoken

import (
	_ "embed"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromModel(t *testing.T) {
	_, err := FromModel("code-davinci-002")
	require.Nil(t, err)

	_, err = FromModel("not-exist")
	require.Equal(t, err, errors.New("failed to load tokenizer for model: not-exist"))
}

func TestEncodeDecodeCorrectly(t *testing.T) {
	bpe, err := FromModel("gpt-3.5-turbo-16k")
	require.Nil(t, err, "Encoding  init should not be nil")

	// Encode
	actualTokens := bpe.Encode("hello world!你好，世界！")
	expectedTokens := []uint32{15339, 1917, 0, 57668, 53901, 3922, 3574, 244, 98220, 6447}
	require.Equal(t, actualTokens, expectedTokens, "Encoding should be equal")

	// Decode
	actualText := bpe.Decode(actualTokens)
	require.Equal(t, actualText, "hello world!你好，世界！", "Decoded result should be equal")
}
