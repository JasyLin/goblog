package session

import (
	"github.com/gorilla/sessions"
	"jasy/goblog/pkg/config"
	"jasy/goblog/pkg/logger"
	"net/http"
)

// Store gorilla sessions 的存储库
var Store = sessions.NewCookieStore([]byte(config.GetString("app.key")))

var Session *sessions.Session

var Request *http.Request

var Response http.ResponseWriter

func StartSession(w http.ResponseWriter, r *http.Request) {
	var err error

	Session, err = Store.Get(r, config.GetString("session.session_name"))
	logger.LogError(err)

	Response = w
	Request = r
}

func Put(key string, value interface{}) {
	Session.Values[key] = value
	Save()
}

func Get(key string) interface{} {
	return Session.Values[key]
}

func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

func Flush() {
	Session.Options.MaxAge = -1
	Save()
}

func Save() {
	err := Session.Save(Request, Response)
	logger.LogError(err)
}