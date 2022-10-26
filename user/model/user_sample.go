package user

import (
	common "go-study/user/common"
)

var Users = []*User{
	{
		Id:          common.StringToUuid("7f11671a-eeb5-4d0e-bcb0-f83260a6d371"),
		Name:        "user1",
		Email:       "user1@gmail.com",
		PhoneNumber: "010-1111-1111",
		Address:     "Seoul",
	},
	{
		Id:          common.StringToUuid("7f11671a-eeb5-4d0e-bcb0-f83260a6d372"),
		Name:        "user2",
		Email:       "user2@gmail.com",
		PhoneNumber: "010-2222-2222",
		Address:     "Busan",
	},
	{
		Id:          common.StringToUuid("7f11671a-eeb5-4d0e-bcb0-f83260a6d373"),
		Name:        "user3",
		Email:       "user3@gmail.com",
		PhoneNumber: "010-3333-3333",
		Address:     "Daegu",
	},
	{
		Id:          common.StringToUuid("7f11671a-eeb5-4d0e-bcb0-f83260a6d374"),
		Name:        "user4",
		Email:       "user4@gmail.com",
		PhoneNumber: "010-4444-4444",
		Address:     "Jeju",
	},
	{
		Id:          common.StringToUuid("7f11671a-eeb5-4d0e-bcb0-f83260a6d375"),
		Name:        "user5",
		Email:       "user5@gmail.com",
		PhoneNumber: "010-5555-5555",
		Address:     "Daejeon",
	},
}
