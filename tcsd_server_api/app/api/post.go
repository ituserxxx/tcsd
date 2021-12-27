package api

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"go_api/app/in_out"
	"go_api/app/service"
	"go_api/constant"
	"go_api/utils/response"
)

// Post postApi
var Post *postApi

type postApi struct{}

func (p *postApi) CreatePost(r *ghttp.Request) {

	var req *in_out.CreatePostReq
	if err := r.Parse(&req); err != nil {
		response.Err(r, constant.ParamsError, err.Error())
	}
	if st := r.Session.GetInt64(gmd5.Encrypt(req));st > 0 {
		if gtime.Timestamp() - st < 300{
			response.Err(r, constant.ParamsError, "大哥，慢点~~")
		}
	}
	err := service.Post.Create(req)
	if err != nil {
		response.Err(r, -1, err.Error())
	}
	response.Succ(r, "ok")
}

func (p *postApi) PostList(r *ghttp.Request) {
	var req *in_out.PostListReq
	if err := r.Parse(&req); err != nil {
		response.Err(r, constant.ParamsError, err.Error())
	}
	l,err := service.Post.PostList(req)
	if err != nil {
		response.Err(r, constant.LogicError, err.Error())
	}
	response.Succ(r, l)
}

func (p *postApi)GiveLove(r *ghttp.Request) {
	var req *in_out.GiveLoveReq
	if err := r.Parse(&req); err != nil {
		response.Err(r, constant.ParamsError, err.Error())
	}
	service.Post.GiveLove(req)
	response.Succ(r, "")
}
