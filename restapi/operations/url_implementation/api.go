package url_implementation

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/paologallinaharbur/shorturl/models"
	"github.com/paologallinaharbur/shorturl/restapi/operations/url"
	"github.com/paologallinaharbur/shorturl/storage"
	"log"
	"math/rand"
	"time"
)

//CreateURLHandler handles /api/url requests
func CreateURLHandler(createURLParams url.CreateURLParams, db storage.Storage) middleware.Responder {

	short := "short-" + randStringRunes(10)
	if createURLParams.InfoURL.URL == nil {
		return url.NewDeleteURLBadRequest()
	}

	err := db.Write(short, *createURLParams.InfoURL.URL)
	if err != nil {
		eMess := err.Error()
		return url.NewCreateURLInternalServerError().WithPayload(&models.Error{Message: &eMess})
	}
	return url.NewCreateURLCreated().WithPayload(&models.Shorturl{Shorturl: &short})
}

//DeleteURLHandler handles DELETE /api/url/{shortURL} requests
func DeleteURLHandler(deleteURLParams url.DeleteURLParams, db storage.Storage) middleware.Responder {

	err := db.Delete(deleteURLParams.ShortURL)
	if err != nil {
		log.Println(err)
		return url.NewDeleteURLInternalServerError()
	}
	return nil
}

//GetURLHandler handles GET /api/url/{shortURL} requests
func GetURLHandler(getURLParams url.GetURLParams, db storage.Storage) middleware.Responder {
	s, err := db.Read(getURLParams.ShortURL)
	if err != nil {
		log.Println(err)
		return url.NewGetURLInternalServerError()
	}
	return url.NewGetURLOK().WithPayload(&models.URL{
		URL: &s,
	})
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
