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
	mainPageHTML += "<H1>Welcome to the MongoDB Miner!</H1>"
	mainPageHTML += "<H2>Possible endpoints:</H2>"
	mainPageHTML += "<ul>"
	mainPageHTML += "<li>/ : This main page</li>"
	mainPageHTML += "<li>/api-status : Tell if the server is up and running <br>"
	mainPageHTML += "<button onClick=\"routeToApiStatus()\">Check the API Status</button>"
	mainPageHTML += "</li>"
	mainPageHTML += "<li>/mine/{ip_addr}: Search the database at an ip address you provide <br>"
	mainPageHTML += "<button onClick=\"routeToMine()\">Query the local MongoDB (at ip_addr 127.0.0.1)</button>"
	mainPageHTML += "</li></ul>"
	mainPageHTML += "<script>"
	mainPageHTML += "function routeToApiStatus() {"
	mainPageHTML += "window.location.href = '/api-status'; }"
	mainPageHTML += "function routeToMine() {"
	mainPageHTML += "window.location.href = '/mine/127.0.0.1'; }"
	mainPageHTML += "</script>"
	mainPageHTML += "</body></html>"
	fmt.Fprintf(w, mainPageHTML)
}

func ApiStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "status" : "API is up and running }"`))
}

func Mine(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)

	params := mux.Vars(r)
	ip_addr := params["ip_addr"]

	mineResults := Main(ip_addr)

	for _, resultLine := range mineResults {
		log.Printf(resultLine)
	}

	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)
	mainPageHTML := "<html><body>"
	mainPageHTML += "<H1>The results for mining ip addr " + ip_addr + ":</H1>"
	for _, resultLine := range mineResults {
		mainPageHTML += "<p>" + resultLine + "</p>"
	}
	mainPageHTML += "</body></html>"
	fmt.Fprintf(w, mainPageHTML)
}
