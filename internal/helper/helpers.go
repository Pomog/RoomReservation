package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"udemyCourse1/internal/config"
)

var app *config.AppConfig

// NewHelpers sets the config for the helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func IsAutenticated (r *http.Request) bool {
	exist := app.Session.Exists(r.Context(), "user_id")
	return exist
}
