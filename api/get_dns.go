// api-pub-net/api
//
// This package holds our API handlers which we use to service REST API
// requests about DNS entries.
//
// Set the paramter"target=<HOST>" to define the host/domain to lookup.

package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/bunrithlim/api-pub-net/models"
	"net/http"
	"net"
	//"io/ioutil"
	"log"
    "net/http/httputil"
)

const(
	URL_DNS_DEFAULT_TARGET = "google.com"
)

func GetDnsIP(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	_host := URL_DNS_DEFAULT_TARGET

	if target, ok := r.Form["target"]; ok && len(target) > 0 {
        _host = target[0];
	}

	output := ""
    log.Print("Host: ", _host)
	res, err := http.Get(_host)
	if err != nil {
		log.Print(err)
        output = fmt.Sprint(err)
	} else {

	if res.StatusCode == 200 {
		//header, herr := ioutil.ReadAll(res.Body)
        output = fmt.Sprintf("%s\n%s", output, res.Header)
		//robots, err := ioutil.ReadAll(res.Body)
        //output = fmt.Sprintf("%s\n\n%s", output, robots)

	dump, derr := httputil.DumpResponse(res, true)
	if derr != nil {
		log.Print(derr)
        output = fmt.Sprint(derr)
	}
	output = fmt.Sprintf("%s\n\n%s", output, dump)    
        
        
		res.Body.Close()
		if err != nil {
			log.Print(err)
            output = fmt.Sprint(err)
		}
	}
    }
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, string(output))
}

func GetDnsCname(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

func GetDnsHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

func GetDnsNs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

func GetDnsMx(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

func GetDnsTxt(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

func GetDnsAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	_dns := new(models.DNSAll)
	_host := URL_DNS_DEFAULT_TARGET

	if target, ok := r.Form["target"]; ok && len(target) > 0 {
        _host = target[0];
	}

	cname, _ := net.LookupCNAME(_host)
	_dns.Cname = cname

	addrs, _ := net.LookupHost(_host)
	_dns.Host_addrs = addrs

	mx, _ := net.LookupMX(_host)
	_dns.Mx = mx

	ns, _ := net.LookupNS(_host)
	_dns.Ns = ns

	netip, _ := net.LookupIP(_host)
	_dns.IP = netip

	txts, _ := net.LookupTXT(_host)
	_dns.Txts = txts

	// If the user specifies a 'format' querystring, we'll try to return the
	// user's IP address in the specified format.
	if format, ok := r.Form["format"]; ok && len(format) > 0 {
		jsonStr, _ := json.Marshal(_dns)

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
}
