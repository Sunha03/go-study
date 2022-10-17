package data

import (
	userpb "go-study/grpcserver/pb"
)

// UserData : user 데이터 셋팅
var UserData = []*userpb.UserMessage{
	{
		UserId:      "1",
		Name:        "Henry",
		PhoneNumber: "01012341234",
		Age:         20,
	},
	{
		UserId:      "2",
		Name:        "Micheal",
		PhoneNumber: "01056785678",
		Age:         21,
	},
	{
		UserId:      "3",
		Name:        "Jessie",
		PhoneNumber: "01090129012",
		Age:         22,
	},
	{
		UserId:      "4",
		Name:        "Max",
		PhoneNumber: "01045129012",
		Age:         23,
	},
	{
		UserId:      "5",
		Name:        "Tony",
		PhoneNumber: "01012345678",
		Age:         24,
	},
}
