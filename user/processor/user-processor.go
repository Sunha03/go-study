package user

import (
	"errors"
	model "go-study/user/model"
)

// GetUserList : user list 조회
func GetUserList() ([]*model.User, error) {
	userList := model.Users

	return userList, nil
}

// GetUser : id를 이용하여 user 조회
func GetUserById(id string) (*model.User, error) {
	userList := model.Users
	for _, u := range userList {
		if u.Id.String() == id {
			return u, nil
		}
	}

	return nil, errors.New("cannot find user")
}

// CreateUser : user 생성
func CreateUser(user *model.User) (id string, err error) {
	userList := model.Users

	userList = append(userList, user)
	model.Users = userList

	return user.Id.String(), nil
}

// UpdateUser : user 수정
func UpdateUser(id string, user *model.User) (string, error) {
	userList := model.Users
	for _, u := range userList {
		if u.Id.String() == id {
			u.Name = user.Name
			u.Email = user.Email
			u.PhoneNumber = user.PhoneNumber
			u.Address = user.Address
		}
	}

	return id, nil
}

// DeleteUser : user 삭제
func DeleteUser(id string) (string, error) {
	newList := []*model.User{}
	userList := model.Users
	for i, u := range userList {
		if u.Id.String() == id {
			newList = append(userList[:i], userList[i+1:]...)
		}
	}

	model.Users = newList

	return id, nil
}
