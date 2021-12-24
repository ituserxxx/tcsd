package service

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"go_api/app/dao"
	"go_api/app/in_out"
	"go_api/app/model"
	"go_api/utils/wx"
)

var User *user
type user struct {
}

func (u *user)Login(req *in_out.LoginReq)(in_out.LoginResp,error)  {
	var data in_out.LoginResp
	res, err := wx.Login(req.Code)
	if err != nil {
		return data, gerror.New("LoginReq 1" + err.Error())
	}
	userRecord, err := dao.TcUser.OmitEmpty().Where("openid=?", res.OpenID).FindOne()
	if err != nil {
		return data, err
	}
	//没有用户记录，即新增
	if userRecord == nil {
		result, err := dao.TcUser.OmitEmpty().Save(&model.TcUser{
			Openid:        res.OpenID,
			CreateTime:    gtime.Now(),
			LastLoginTime: gtime.Now(),
		})
		if err != nil {
			return data, err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return data, err
		}
		data.Uid = gconv.Int(id)
		return data,nil
	}
	var userInfo *model.TcUser
	if err := userRecord.Struct(&userInfo); err != nil {
		return data, err
	}
	_, _ = dao.TcUser.OmitEmpty().Where("id=?", userInfo.Id).Update(&model.TcUser{
		LastLoginTime: gtime.Now(),
		SessionKey:    res.SessionKey,
	})
	data.Uid = userInfo.Id
	data.NickName = userInfo.Nickname
	data.AvatarUrl = userInfo.AvatarUrl
	if userInfo.Nickname != ""{
		data.IsAuth = true
	}
	data.PostSum = userInfo.PostSum
	return data,nil
}
func (u *user)UpdateUserInfo(req *in_out.UpdateUserInfo) error {
	go func() {
		l,err := dao.TcUser.OmitEmpty().Where("id=?",req.UID).FindOne()
		if err != nil {
			return
		}
		var u model.TcUser
		if err := l.Struct(&u);err != nil{
			return
		}

		res,err := wx.DecryptMsg(u.SessionKey,req.EncryptedData,req.Iv)
		if err != nil {
			return
		}
		_, err = dao.TcUser.OmitEmpty().Where("id =?", req.UID).
			Data(model.TcUser{
				Id:      req.UID,
				Unionid: res.UnionID,
				Openid:  res.OpenID,
			}).Update()
		if err != nil {
			return
		}
	}()
	_,err := dao.TcUser.OmitEmpty().Where("id =?",req.UID).
		Data(model.TcUser{
		Id:        req.UID,
		Gender:    req.Gender,
		AvatarUrl: req.AvatarUrl,
		Nickname:  req.NickName,
	}).Update()
	if err != nil {
		return err
	}
	return nil
}
func (u *user)UpdateUserSet(req *in_out.UpdateUserSetReq)  {
	_,_ = dao.TcUser.OmitEmpty().Where("id=?",req.UID).Data(model.TcUser{
		IsAnonymous: req.IsAnonymous,
		IsSquare:    req.IsSquare,
	}).Update()
}