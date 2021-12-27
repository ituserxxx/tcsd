package route

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go_api/app/api"
)

func init() {
	s := g.Server()
	s.BindHandler("/xxx", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！123")
		r.Exit()
	})
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.ALL("/user/login", api.User.Login)
		group.ALL("/user/update", api.User.UpdateUserInfo)
		group.ALL("/user/update/set", api.User.UpdateUserSet)
		group.ALL("/post/create", api.Post.CreatePost)
		group.ALL("/post/list", api.Post.PostList)
		group.ALL("/post/give_love", api.Post.GiveLove)
	})

}
