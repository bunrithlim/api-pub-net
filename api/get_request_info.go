// api-pub-net/api
//
// This package holds our API handlers which we use to service REST API
// requests about the HTTP request.

package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/bunrithlim/api-pub-net/models"
	"net/http"
	"net"
)

// GetRequestInfo returns info about the HTTP request. Includes, user agent,
// public facing IP address (IPv4 OR IPv6), and referrer URL.
//
// By default, it will return the info in plain text, but can also return
// data in both JSON and JSONP if requested to.
func GetRequestInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// Get the desired request info
	ip := GetUserPublicIP(r)
	ua := r.UserAgent()
	rf := r.Referer()

	_all := fmt.Sprintf("%s\n%s\n%s", ua, rf, ip)

	// If the user specifies a 'format' querystring, we'll try to return the
	// user's IP address in the specified format.
	if format, ok := r.Form["format"]; ok && len(format) > 0 {
		jsonStr, _ := json.Marshal(models.RequestInfo{ip, ua, rf})

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

	// If no 'format' querystring was specified, we'll default to returning the
	// IP in plain text.
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, _all)
}

func GetRequestHeaders(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	_all := fmt.Sprintf("%s\n%s\n%s", r.Method, r.URL, r.Proto)

	_request := models.RequestHeaders{
		Method : r.Method,
		URL : fmt.Sprintf("%s", r.URL),
		Protocol : r.Proto,
		Headers : make([]models.RequestHeader, 0),
	}

	_all = fmt.Sprintf("%s\nHost = %q", _all, r.Host)
	_all = fmt.Sprintf("%s\nRemoteAddr= %q", _all, r.RemoteAddr)

	_request.Host = r.Host
	_request.RemoteIpAddress = r.RemoteAddr

	_headers := make([]models.RequestHeader, 0)

	//Iterate over all header fields
	for k, v := range r.Header {
		_all = fmt.Sprintf("%s\nHeader field %q, Value %q", _all, k, v)
		_hd := models.RequestHeader{
			Header : string(k),
			Value : fmt.Sprintf("%q", v),
		}
		_headers = append(_headers, _hd)
	}

	_request.Headers = _headers;

	// If the user specifies a 'format' querystring, we'll try to return the
	// user's IP address in the specified format.
	if format, ok := r.Form["format"]; ok && len(format) > 0 {
		jsonStr, _ := json.Marshal(_request)

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

	// If no 'format' querystring was specified, we'll default to returning the
	// IP in plain text.
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, _all)
}