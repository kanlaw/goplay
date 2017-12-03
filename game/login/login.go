package login

import (
	"encoding/json"

	"goplay/data"
	"goplay/glog"
	"goplay/pb"
	"utils"
)

//登录验证
func LoginCheck(ctos *pb.CLogin) (stoc *pb.SLogin) {
	stoc = new(pb.SLogin)
	var phone string = ctos.GetPhone()
	var passwd string = ctos.GetPassword()
	if !utils.PhoneRegexp(phone) {
		stoc.Error = pb.PhoneNumberError
	}
	if len(passwd) != 32 {
		stoc.Error = pb.PwdFormatError
	}
	return
}

//登录
func Login(arg *pb.RoleLogin, user *data.User) (stoc *pb.SLogin) {
	stoc = new(pb.RoleLogined)
	var phone string = ctos.GetPhone()
	var passwd string = ctos.GetPassword()
	if !user.VerifyPwd(passwd) {
		//数据库中查找
		user.Phone = phone
		if !user.VerifyPwdByPhone(passwd) {
			stoc.Error = pb.UsernameOrPwdError
		}
	}
	if user.Userid == "" {
		stoc.Error = pb.LoginError
	}
	if stoc.Error != pb.OK {
		return
	}
	result, err := json.Marshal(user)
	if err != nil {
		glog.Errorf("user Marshal err %v", err)
		stoc.Error = pb.LoginError
		return
	}
	stoc.Data = string(result)
	return
}
