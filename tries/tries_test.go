package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Trie(t *testing.T) {

	var trie Trie = Constructor()

	assert.Equal(t, false, trie.Search(""))

	trie.Insert("apple")
	assert.Equal(t, true, trie.Search("apple"))
	assert.Equal(t, false, trie.Search("app"))
	assert.Equal(t, true, trie.StartsWith("apple"))

	trie.Insert("dog")
	trie.Insert("app")
	assert.Equal(t, true, trie.Search("app"))
}
