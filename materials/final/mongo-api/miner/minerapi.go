package miner

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)
	mainPageHTML := "<html><body>"
	mainPageHTML += "<H1>Welcome to the Database Miner Interface!</H1>"
	mainPageHTML += "<H2>Possible endpoints:</H2>"
	mainPageHTML += "<ul>"
	mainPageHTML += "<li>/ : This main page</li>"
	mainPageHTML += "<li>/api-status : Tell if the server is up and running <br>"
	mainPageHTML += "<button onClick=\"routeToApiStatus()\">Check the API Status</button>"
	mainPageHTML += "</li>"
	mainPageHTML += "<li>/mongo-mine/{ip_addr}: Search a Mongo database at an ip address you provide <br>"
	mainPageHTML += "<button onClick=\"routeToMongoMine()\">Query the local MongoDB (at ip_addr 127.0.0.1)</button>"
	mainPageHTML += "</li></ul>"
	mainPageHTML += "<script>"
	mainPageHTML += "function routeToApiStatus() {"
	mainPageHTML += "window.location.href = '/api-status'; }"
	mainPageHTML += "function routeToMongoMine() {"
	mainPageHTML += "window.location.href = '/mongo-mine/127.0.0.1'; }"
	mainPageHTML += "</script>"
	mainPageHTML += "</body></html>"
	fmt.Fprintf(w, mainPageHTML)
}

func ApiStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	mainPageHTML := "<html><body>"
	mainPageHTML += "<button onClick=\"routeToHome()\">Return Home</button>"
	mainPageHTML += "<script>"
	mainPageHTML += "function routeToHome() {"
	mainPageHTML += "window.location.href = '/'; }"
	mainPageHTML += "</script>"
	mainPageHTML += "<p>{ \"status\" : \"API is up and running\" }</p>"
	mainPageHTML += "</body></html>"

	fmt.Fprintf(w, mainPageHTML)
}

func MongoMine(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)

	params := mux.Vars(r)
	ip_addr := params["ip_addr"]

	mineResults := MongoMain(ip_addr)

	for _, resultLine := range mineResults {
		log.Printf(resultLine)
	}

	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)
	mainPageHTML := "<html><body>"
	mainPageHTML += "<button onClick=\"routeToHome()\">Return Home</button>"
	mainPageHTML += "<script>"
	mainPageHTML += "function routeToHome() {"
	mainPageHTML += "window.location.href = '/'; }"
	mainPageHTML += "</script>"

	mainPageHTML += "<H1>The results for mining the MongoDB at ip address " + ip_addr + ":</H1>"
	for _, resultLine := range mineResults {
		mainPageHTML += "<p>" + resultLine + "</p>"
	}

	mainPageHTML += "</body></html>"
	fmt.Fprintf(w, mainPageHTML)
}
