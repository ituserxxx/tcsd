package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"go_api/app/dao"
	"go_api/app/in_out"
	"go_api/app/model"
	"sync"
)

var Post *post

type post struct {
}

func (p *post) Create(req *in_out.CreatePostReq) error {


	if tx, err := g.DB().Begin(); err == nil {
		u, err := dao.TcUser.OmitEmpty().Where("id=?", req.UID).FindOne()
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		var uInfo model.TcUser
		if err := u.Struct(&uInfo); err != nil {
			return err
		}
		_, err = dao.TcPost.TX(tx).OmitEmpty().Insert(&model.TcPost{
			Uid:         req.UID,
			Content:     req.Context,
			IsAnonymous: uInfo.IsAnonymous,
			IsSquare:    uInfo.IsSquare,
			CreateTime:  gtime.Now(),
		})
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		_, err = dao.TcUser.TX(tx).OmitEmpty().Where("id=?", req.UID).Increment("post_sum", 1)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		_ = tx.Commit()
	}
	return nil
}

func (p *post) PostList(req *in_out.PostListReq) (in_out.PostListResp, error) {
	var data in_out.PostListResp
	listRecord, err := dao.TcPost.OmitEmpty().Page(req.Page, 10).OrderDesc("id").FindAll()
	if err != nil {
		return data, err
	}
	var list []*model.TcPost
	if err := listRecord.Structs(&list); err != nil {
		return data, err
	}

	var postList []in_out.PostListItemResp
	for _, pRecord := range list {
		var ll in_out.PostListItemResp
		var pMasterInfo in_out.PostListRespUserInfo
		var pInfo in_out.PostListRespPostInfo

		pInfo.Id = pRecord.Id
		pInfo.Content = pRecord.Content
		pInfo.LoveSum = pRecord.LoveSum

		pML, err := dao.TcUser.OmitEmpty().Where("id=?", pRecord.Uid).FindOne()
		if err != nil {
			return data, err
		}
		var pMastInfo model.TcUser
		if err := pML.Struct(&pMastInfo); err != nil {
			return data, err
		}
		pMasterInfo.Id = pRecord.Uid
		pMasterInfo.AvatarUrl = pMastInfo.AvatarUrl
		pMasterInfo.NickName = pMastInfo.Nickname

		isDianzan, err := dao.TcPostLove.
			OmitEmpty().
			Where("user_id =? and post_id = ?", req.UID, pRecord.Id).
			FindOne()
		if err != nil {
			return data, err
		}
		if !isDianzan.IsEmpty() {
			pInfo.IsGiveLove = true
		}
		postLoveUserL, err := dao.TcPostLove.
			OmitEmpty().
			Where("post_id=?", pRecord.Id).
			OrderDesc("id").
			Limit(5).
			FindAll()
		if err != nil {
			return data, err
		}
		var postLoveUserList []*model.TcPostLove
		if err := postLoveUserL.Structs(&postLoveUserList); err != nil {
			return data, err
		}
		var sImg []string
		for _, lov := range postLoveUserList {
			uOne, err := dao.TcUser.OmitEmpty().Where("id=?", lov.UserId).FindOne()
			if err != nil {
				return data, err
			}
			var uM model.TcUser
			if err := uOne.Struct(&uM); err != nil {
				return data, err
			}
			//此处应该随机图片
			sImg = append(sImg, "https://thirdwx.qlogo.cn/mmopen/vi_32/vlBrF1UadszIEGzzePcicuDg2Vic0VdokPiaq7QaSPR5PicafOJLG0BLW8icJOmpt0ey792I4vhRKJzPNOU7sbJFayg/132")
		}
		ll.PostInfo.LoveUserImgList = sImg
		ll.MasterInfo = pMasterInfo
		ll.PostInfo = pInfo
		postList = append(postList, ll)
	}
	data.List = postList
	return data, nil
}

func (p *post) GiveLove(req *in_out.GiveLoveReq) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var loc sync.RWMutex
		loc.RLock()
		defer func() {
			loc.RUnlock()
			wg.Done()
		}()
		if tx,err := g.DB().Begin();err == nil{
			l, err := dao.TcPostLove.OmitEmpty().TX(tx).Where("user_id =? and post_id = ?", req.UID, req.PostId).FindOne()
			if err != nil {
				_ = tx.Rollback()
				return
			}
			if l.IsEmpty() {
				_, _ = dao.TcPostLove.OmitEmpty().TX(tx).Data(model.TcPostLove{
					PostId:     req.PostId,
					UserId:     req.UID,
					CreateTime: gtime.Now(),
				}).Insert()
				_,_ = dao.TcPost.OmitEmpty().TX(tx).Where("id = ?",req.PostId).Increment("love_sum",1)

			} else {
				_, _ = dao.TcPostLove.OmitEmpty().TX(tx).Where("user_id =? and post_id = ?", req.UID, req.PostId).Delete()
				_,_ = dao.TcPost.OmitEmpty().TX(tx).Where("id = ?",req.PostId).Decrement("love_sum",1)
			}
			_ = tx.Commit()
		}
	}()

	wg.Wait()
}
