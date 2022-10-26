package user

import (
	"encoding/json"
	"net/http"

	model "go-study/user/model"
	processor "go-study/user/processor"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// swagger url : http://localhost:12000/swagger/index.html

// @Summary		GetUserList
// @Description	user 목록 조회
// @Tags		user
// @Accept		json
// @Produce		json
// @Success		200		{array} model.User
// @Failure		400
// @Failure		500
// @Router		/user	[GET]
// GetUserList : user 목록 조회
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

// @Summary		GetUser
// @Description	user 조회 by id
// @Tags		user
// @Accept		json
// @Produce		json
// @Param       userId  	path	string	true	"ex) 7f11671a-eeb5-4d0e-bcb0-f83260a6d375"
// @Success		200			{object} model.User
// @Failure		400
// @Failure		500
// @Router		/user/{userId}	[GET]
// GetUser : user 조회
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

// @Summary		CreateUser
// @Description	user 생성
// @Tags		user
// @Accept		json
// @Produce		json
// @Param       id  		path	string		true	"ex) 7f11671a-eeb5-4d0e-bcb0-f83260a6d375"
// @Param       body  		body	model.User	true	"body params"
// @Success		200			{object} objectid{id=string}
// @Failure		400
// @Failure		500
// @Router		/user		[POST]
// CreateUser : user 생성
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

// @Summary		UpdateUser
// @Description	user 수정
// @Tags		user
// @Accept		json
// @Produce		json
// @Param       id  		path	string		true	"ex) 7f11671a-eeb5-4d0e-bcb0-f83260a6d375"
// @Param       body  		body	model.User	true	"body params"
// @Success		200			{object} objectid{id=string}
// @Failure		400
// @Failure		500
// @Router		/user		[PATCH]
// UpdateUser : user 수정
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

// @Summary		DeleteUser
// @Description	user 삭제
// @Tags		user
// @Accept		json
// @Produce		json
// @Param       id  		path	string	true	"ex) 7f11671a-eeb5-4d0e-bcb0-f83260a6d375"
// @Success		200			{object} objectid{id=string}
// @Failure		400
// @Failure		500
// @Router		/user/{userId}		[DELETE]
// DeleteUser : user 삭제
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
