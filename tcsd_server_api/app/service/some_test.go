package service

import (
	"fmt"
	"go_api/app/in_out"
	"testing"
)

func TestName(t *testing.T) {
	l,err := Post.PostList(&in_out.PostListReq{
		UID:  3,
		Page: 1,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}else{
		fmt.Println(l.List[0])
	}

}
