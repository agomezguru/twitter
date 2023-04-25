package routers

import (
	"encoding/json"
	"net/http"

	"github.com/agomezguru/twitter/db"
	"github.com/agomezguru/twitter/models"
)

/* User register routine */
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User 
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Data retrieved error: " + err.Error(), 400)
		return
	}

	// Start validations
	if len(t.Email) == 0 {
		http.Error(w, "User email is required: " + err.Error(), 400)
		return
	}

	if len(t.Password) < 8 {
		http.Error(w, "Password is to short, use 8 characters or more.", 400)
		return
	}

	_, finded, _ := db.UserExist(t.Email)
	if finded == true {
		http.Error(w, "User previously registered with this email." , 400)
		return
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "Error ocurred when try to register a new user:" + err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Insert register in DB failed.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
