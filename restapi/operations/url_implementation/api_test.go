package url_implementation

import (
	"github.com/paologallinaharbur/shorturl/models"
	"github.com/paologallinaharbur/shorturl/restapi/operations/url"
	"github.com/paologallinaharbur/shorturl/storage"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestApis(t *testing.T) {
	//it accept an interface, but for sake of simplicity we pass the actual obj
	st := storage.NewStorageDB("test")
	defer os.Remove("test")

	testURL := "test"
	//we create the url
	resp := CreateURLHandler(url.CreateURLParams{
		InfoURL: &models.URL{URL: &testURL},
	}, st)

	a := resp.(*url.CreateURLCreated)
	assert.NotNil(t, *a.Payload.Shorturl)

	//We get the url
	resp2 := GetURLHandler(url.GetURLParams{
		ShortURL: *a.Payload.Shorturl,
	}, st)

	b := resp2.(*url.GetURLOK)
	assert.Equal(t, *b.Payload.URL, "test")

	//We delete the url
	resp3 := DeleteURLHandler(url.DeleteURLParams{
		ShortURL: *a.Payload.Shorturl,
	}, st)

	assert.Nil(t, resp3)

	//We get the url again failing
	resp4 := GetURLHandler(url.GetURLParams{
		ShortURL: *a.Payload.Shorturl,
	}, st)

	_, ok := resp4.(*url.GetURLInternalServerError)
	assert.Equal(t, true, ok)
}
