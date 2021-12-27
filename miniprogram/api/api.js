
import {request,getUrl}  from '../utils/util';

const http = getUrl()
// 登录
const userLogin = data => {
    // data.str=""
    return request({data:data,url:http+'/user/login'})
};
const userUpdateInfo = data => {
    // data.str=""
    return request({data:data,url:http+'/user/update'})
};
const upUserPostSet = data => {
    return request({data:data,url:http+'/user/update/set'})
};
// 发布吐槽
const postCreate = data => {
    return request({data:data,url:http+'/post/create'})
};
// 吐槽列表
const postList = data => {
    return request({data:data,url:http+'/post/list'})
};

// 点赞
const giveLove = data => {
    return request({data:data,url:http+'/post/give_love'})
};
module.exports = {
    userLogin,
    postCreate,
    userUpdateInfo,
    upUserPostSet,
    postList,
    giveLove
}