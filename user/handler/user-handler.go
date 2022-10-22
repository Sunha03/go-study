package user

import (
	"encoding/json"
	"net/http"

	model "go-study/user/model"
	processor "go-study/user/processor"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetUserList(w http.ResponseWriter, r *http.Request) {
	userList, err := processor.GetUserList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&userList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	userId := v["userId"]

	user, err := processor.GetUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *model.User

	// request body parsing
	d := json.NewDecoder(r.Body)
	err := d.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// validation
	if user.Id == uuid.Nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := processor.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"id": res})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *model.User

	// request body parsing
	v := mux.Vars(r)
	userId := v["userId"]

	d := json.NewDecoder(r.Body)
	err := d.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// validation
	if userId == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := processor.UpdateUser(userId, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"id": res})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// request body parsing
	v := mux.Vars(r)
	userId := v["userId"]

	// validation
	if userId == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := processor.DeleteUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"id": res})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
