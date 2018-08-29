// api-pub-net/api
//
// This package holds our API handlers which we use to service information.

package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	URL_HOME   = "https://github.com/bunrithlim/api-pub-net"
	URL_HELP   = "https://github.com/bunrithlim/api-pub-net/blob/master/README.md"
	URL_SOURCE = "https://github.com/bunrithlim/api-pub-net"
)

// GetHome redirects to the github page.
func GetHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, URL_HOME, http.StatusSeeOther)
}

// GetHelp redirects to the github readme page.
func GetHelp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, URL_HELP, http.StatusSeeOther)
}