// ipify-api
//
// This is the main package which starts up and runs our REST API service.
//
// ipify is a simple API service which returns a user's public IP address (it
// supports handling both IPv4 and IPv6 addresses).

package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/bunrithlim/api-pub-net/api"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// main launches our web server which runs indefinitely.
func main() {

	// Setup all routes.  We only service API requests, so this is basic.
	router := httprouter.New()
	router.GET("/",     api.GetHome)
	router.GET("/help", api.GetHelp)

	// IP related features
	router.GET("/ip", api.GetIP)

	// Time related features
	router.GET("/utc",       api.GetTimeUTC)
	router.GET("/utc/milli", api.GetTimeUTCMilli)
	router.GET("/utc/nano",  api.GetTimeUTCNano)

	// Requester related features
	router.GET("/who",         api.GetRequestInfo)
	router.GET("/who/headers", api.GetRequestHeaders)

	// DNS related features
	router.GET("/dns",       api.GetDnsIP)
	router.GET("/dns/all",   api.GetDnsAll)
	router.GET("/dns/cname", api.GetDnsCname)
	router.GET("/dns/host",  api.GetDnsHost)
	router.GET("/dns/ip",    api.GetDnsIP)
	router.GET("/dns/ns",    api.GetDnsNs)
	router.GET("/dns/mx",    api.GetDnsMx)
	router.GET("/dns/txt",   api.GetDnsTxt)


	// Setup 404 / 405 handlers.
	router.NotFound = http.HandlerFunc(api.NotFound)
	router.MethodNotAllowed = http.HandlerFunc(api.MethodNotAllowed)

	// Setup middlewares.  For this we're basically adding:
	//	- Support for CORS to make JSONP work.
	handler := cors.Default().Handler(router)

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Starting HTTP server on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
