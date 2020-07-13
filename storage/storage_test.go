package storage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//For the sake of semplicity we put everything in only one test, but it would be better to break them down
func TestStorage(t *testing.T) {
	st := NewStorageDB("test")
	defer os.Remove("test")

	//NotFound
	s, err := st.Read("testShortURL")
	assert.Error(t, err)
	assert.Equal(t, s, "")

	//creation
	err = st.Write("testShortURL", "testURL")
	assert.NoError(t, err)

	//doublecreation fails
	err = st.Write("testShortURL", "testURL")
	assert.Error(t, err)

	//read works
	s, err = st.Read("testShortURL")
	assert.NoError(t, err)
	assert.Equal(t, s, "testURL")

	//delete
	err = st.Delete("testShortURL")
	assert.NoError(t, err)

	//read fails after deletions
	s, err = st.Read("testShortURL")
	assert.Error(t, err)
	assert.Equal(t, s, "")

	//write works after deletions
	err = st.Write("testShortURL", "testURL")
	assert.NoError(t, err)

	//delete
	err = st.Delete("testShortURL")
	assert.NoError(t, err)

	//delete again generate no error
	err = st.Delete("testShortURL")
	assert.NoError(t, err)

}
