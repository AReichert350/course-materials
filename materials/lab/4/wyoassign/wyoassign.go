package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type Response struct{
	Assignments []Assignment `json:"assignments"`
}

type Assignment struct {
	Id string `json:"id"`
	Title string `json:"title`
	Description string `json:"desc"`
	Points int `json:"points"`
}

var Assignments []Assignment
const Valkey string = "FooKey"

func InitAssignments(){
	var assignmnet Assignment
	assignmnet.Id = "Mike1A"
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	Assignments = append(Assignments, assignmnet)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")

	var jsonResponse []byte
	var err error

	if len(response.Assignments) == 0 {
		jsonResponse, err = json.Marshal("Yay! Looks like you currently don't have any assignments")
	} else {
		jsonResponse, err = json.Marshal(response)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//TODO 
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, assignment := range Assignments {
		if assignment.Id == params["id"]{
			json.NewEncoder(w).Encode(assignment)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	//TODO : Provide a response if there is no such assignment
	// If reached this point, there is no such assignment
	jsonResponse, _ := json.Marshal("No assignment exists for the requested ID (" + params["id"] + ")")
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonResponse)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
			if assignment.Id == params["id"]{
				Assignments = append(Assignments[:index], Assignments[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	
	var response Response
	response.Assignments = Assignments

	params := mux.Vars(r)

	updatedAssignmentIndex := -1

	// Lookup index in Assignments user passed for assignment they want to update
	r.ParseForm()
	if(r.FormValue("id") != "") {
		for i, assignment := range response.Assignments {
			if assignment.Id == params["id"] {
				updatedAssignmentIndex = i
				break
			}
		}
	}

	// If the assignment index to update was not a pre-existing assignment index,
	// return that the assignment wasn't found
	if updatedAssignmentIndex == -1 {
		w.WriteHeader(http.StatusNotFound)
		jsonResponse, _ := json.Marshal("No assignment exists for the requested ID (" + params["id"] + ")")
		w.Write(jsonResponse)
		return
	}

	// Update the actual Assignment class variable
	Assignments[updatedAssignmentIndex].Title = r.FormValue("title")
	Assignments[updatedAssignmentIndex].Description = r.FormValue("desc")
	Assignments[updatedAssignmentIndex].Points, _ = strconv.Atoi(r.FormValue("points"))
	response.Assignments = Assignments
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignmnet Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id")    !=  "" && 
	   r.FormValue("title") !=  "" &&
	   r.FormValue("desc")  !=  "" &&
	   r.FormValue("points") != "") {
		assignmnet.Id =  r.FormValue("id")
		assignmnet.Title =  r.FormValue("title")
		assignmnet.Description =  r.FormValue("desc")
		assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
		Assignments = append(Assignments, assignmnet)
		jsonResponse, _ := json.Marshal("Created the assignment!")
		w.Write(jsonResponse)
		w.WriteHeader(http.StatusCreated)
		return
	}
	jsonResponse, _ := json.Marshal("Missing value(s) in request body for id, title, desc, or points")
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusNotFound)
}