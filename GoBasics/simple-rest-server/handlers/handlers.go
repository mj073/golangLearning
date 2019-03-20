package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	u "GoBasics/simple-rest-server/utils"
	"orc_bitbucket_repo/Godeps/_workspace/src/github.com/gorilla/mux"
	"strings"
)
var stud = make([]u.Student,0)
func init(){
	log.Println("init called")
	stud = append(stud,u.Student{"Mahesh","123",25,u.Marksheet{50,60,40,70}})
	stud = append(stud,u.Student{"Shardul","234",25,u.Marksheet{70,40,60,50}})
	stud = append(stud,u.Student{"Mayur","345",25,u.Marksheet{80,70,60,60}})
	stud = append(stud,u.Student{"Sagar","456",25,u.Marksheet{80,50,70,50}})
}

func GetHandler(w http.ResponseWriter, r *http.Request){
	student := u.Student{}
	urlVars := mux.Vars(r)
	name := urlVars["name"]
	name = strings.TrimPrefix(name,"{")
	name = strings.TrimSuffix(name,"}")
	log.Println("number of students:",len(stud))
	for _,s := range(stud){
		if s.Name == name{
			log.Println("student requested found..")
			student = s
			break
		}
	}
	if student.Id == ""{
		log.Println("Student not found..")
		SendJSONResponse(w,404,nil)
		return
	}
	response := u.JSONData{student}
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