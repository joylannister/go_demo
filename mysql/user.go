package mysql

import (
	"demo/model"
	"demo/pb"
)

func GetUserInfo(in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	sqlStr := "select id,user_name,nickname from user_db where user_name=?"
	var u model.UserInfo
	err := db.Get(&u, sqlStr, in.Username)
	if err != nil {
		return &pb.GetUserResponse{}, err
	}
	return &pb.GetUserResponse{
		Id:       u.Id,
		Username: u.UserName,
		Nickname: u.NickName,
	}, nil
}
