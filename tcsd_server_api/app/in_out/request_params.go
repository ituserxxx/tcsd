package in_out

// LoginReq 登录
type LoginReq struct {
	Code string `json:"code" v:"required"`
}



type UpdateUserInfo struct {
	UID           int    `json:"uid" v:"required"`
	NickName      string `json:"nick_name"`
	Gender        int    `json:"gender"`
	AvatarUrl     string `json:"avatar_url"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}
type UpdateUserSetReq struct {
	UID           int    `json:"uid" v:"required"`
	IsAnonymous int    `json:"is_anonymous" d:"1"`
	IsSquare    int    `json:"is_square" d:"1"`
}


// 创建Post
type CreatePostReq struct {
	UID         int    `json:"uid" v:"required"`
	Context     string `json:"context" v:"required"`
}

type HomePostReq struct {
	UID  int `json:"uid" v:"required"`
	Page int `json:"page"`
}
