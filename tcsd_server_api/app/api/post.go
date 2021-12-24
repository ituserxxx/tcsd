package api

import (
	"github.com/gogf/gf/net/ghttp"
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
	err := service.Post.Create(req)
	if err != nil {
		response.Err(r, -1, err.Error())
	}
	response.Succ(r, "ok")
}

func (p *postApi) HomePostList(r *ghttp.Request) {
	//var req *in_out.HomePostReq
	//if err := r.Parse(&req); err != nil {
	//	utils.Err(r, -1, err.Error())
	//}
	//if req.Page == 0 {
	//	req.Page = 1
	//}
	//listRecord, err := dao.TcPost.OmitEmpty().Where(model.TcPost{
	//	IsAnonymous: 1,
	//}).Page(req.Page, 10).FindAll()
	//if err != nil {
	//	utils.Err(r, -1, err.Error())
	//}
	//var list []*model.TcPost
	//if err := listRecord.Structs(&list); err != nil && err != sql.ErrNoRows {
	//	utils.Err(r, -1, err.Error())
	//}
	//
	//utils.Succ(r, list)
}
