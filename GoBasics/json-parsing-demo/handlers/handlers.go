package handlers

import (
	"net/http"
	"encoding/json"
	"log"
)

func GetHandler(w http.ResponseWriter, r *http.Request){
	stud := make([]Student,0)
	stud = append(stud,Student{"Mahesh","123",25,Marksheet{50,60,40,70}})
	stud = append(stud,Student{"Shardul","234",25,Marksheet{70,40,60,50}})
	stud = append(stud,Student{"Mayur","345",25,Marksheet{80,70,60,60}})
	stud = append(stud,Student{"Sagar","456",25,Marksheet{80,50,70,50}})
	response := JSONData{stud}
	SendJSONResponse(w,200,response)
	return
}

func SendJSONResponse(w http.ResponseWriter, http_status int, responseMetadata interface{}){
	jsonResponse,err := json.Marshal(responseMetadata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error while marshalling the data...sending blank json")
		return
	}
	log.Println("successfully marshalled the data")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http_status)
	//json.NewEncoder(w).Encode(jsonResponse)
	w.Write(jsonResponse)
}