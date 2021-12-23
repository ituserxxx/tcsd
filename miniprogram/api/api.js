
import {request,getUrl}  from '../utils/util';

const http = getUrl()
// 登录
const userLogin = data => {
    // data.str=""
    return request({data:data,url:http+'/user/login'})
};
// 发布吐槽
const postCreate = data => {
    return request({data:data,url:http+'/post/create'})
};

module.exports = {
    userLogin,
    postCreate
}