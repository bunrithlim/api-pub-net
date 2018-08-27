// api-pub-net/api
//
// This package holds our API handlers which we use to service REST API
// requests about UTC times.

package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/bunrithlim/api-pub-net/models"
	"net/http"
	"time"
)

func GetTimeUTCNano(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// Get the UTC time in nanoseconds
	tm := fmt.Sprintf("%d", time.Now().UnixNano())

	FormatTimeUTCResponse(w, r, tm);
}

func GetTimeUTCMilli(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// Get the UTC time in milliseconds
	tm := fmt.Sprintf("%d", time.Now().UnixNano() / int64(time.Millisecond))

	FormatTimeUTCResponse(w, r, tm);
}

func GetTimeUTC(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// Get the UTC time in seconds
	tm := fmt.Sprintf("%d", time.Now().Unix())

	FormatTimeUTCResponse(w, r, tm);
}

func FormatTimeUTCResponse(w http.ResponseWriter, r *http.Request, tm string) {

	// If the user specifies a 'format' querystring, we'll try to return the
	// response in the specified format.

	if format, ok := r.Form["format"]; ok && len(format) > 0 {
		jsonStr, _ := json.Marshal(models.TimeUTC{tm})

		switch format[0] {
		case "json":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(jsonStr))
			return
		case "jsonp":
			// If the user specifies a 'callback' parameter, we'll use that as
			// the name of our JSONP callback.
			callback := "callback"
			if val, ok := r.Form["callback"]; ok && len(val) > 0 {
				callback = val[0]
			}

			w.Header().Set("Content-Type", "application/javascript")
			fmt.Fprintf(w, callback+"("+string(jsonStr)+");")
			return
		}
	}

	// If no 'format' querystring was specified, we'll default to returning
	// results in plain text.
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, tm)
}
