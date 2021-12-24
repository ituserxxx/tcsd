package api

import (
	"github.com/gogf/gf/net/ghttp"
	"go_api/app/in_out"
	"go_api/app/service"
	"go_api/constant"
	"go_api/utils/response"
)

// User userApi 用户管理器
var User *userApi

type userApi struct{}

func (u *userApi) Login(r *ghttp.Request) {
	var req *in_out.LoginReq
	if err := r.Parse(&req); err != nil {
		response.Err(r,constant.ParamsError,err.Error())
	}
	id,err := service.User.Login(req)
	if err != nil {
		response.Err(r, constant.LogicError, err.Error())
	}
	response.Succ(r, id)
}

func (u *userApi) UpdateUserInfo(r *ghttp.Request) {
	var req *in_out.UpdateUserInfo
	if err := r.Parse(&req); err != nil {
		response.Err(r,constant.ParamsError,err.Error())
	}
	err := service.User.UpdateUserInfo(req)
	if err != nil {
		response.Err(r, constant.LogicError, err.Error())
	}
	response.Succ(r, "")
}

func (u *userApi) Xxx(r *ghttp.Request) {
	response.Succ(r, "11111111")
}

func (u *userApi) UpdateUserSet(r *ghttp.Request) {
	var req *in_out.UpdateUserSetReq
	if err := r.Parse(&req); err != nil {
		response.Err(r,constant.ParamsError,err.Error())
	}
	service.User.UpdateUserSet(req)
	response.Succ(r, "")
}