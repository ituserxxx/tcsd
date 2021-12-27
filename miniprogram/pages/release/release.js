import Toast from '../../miniprogram_npm/@vant/weapp/toast/toast';
const api = require("../../api/api.js");
const app = getApp()
Page({
  data: {
    placeholder: "请开始你的表演~~~",
    content: "",
    IsAnonymous: true,
    IsSquare: true,
    show_switch: false,
    isAuth: true,

  },
  onLoad() {
  },
  bindFormSubmit: function (e) {
    // console.log(e)
    // console.log(app.globalData, e.detail.value)
    // console.log(e.detail)
    this.setData({
      content: e.detail.value.textarea
    })
    // if (!app.globalData.isAuth){
    //     this.upUserInfo()
    // }else 
    if (app.globalData.postaSum == 0) {
      this.setData({
        show_switch: true,
      })
    } else if (this.isNull(this.ltrim(e.detail.value.textarea))) {
      Toast("还是随便敲几个字符吧~~")
      return
    } else {
      var isRe = app.preventActive(()=>{
        this.pushPost()
      })
      if (isRe){
        Toast(isRe)
      }
    }
  },
  think_some: function (e) {
    return
  },
  pushPost: function () {
    var data = {
      uid: app.globalData.uid,
      context: this.data.content,
    }
    api.postCreate(data).then(res => {
      if (res.code == 1) {
        Toast.success("发布成功~");
        app.globalData.postaSum++
      } else {
        Toast.fail(res.msg);
      }
    })
  },
  clear_from: function (e) {
    this.setData({
      content: ""
    })
  },
  selct_IsAnonymous: function (e) {
    this.setData({
      IsAnonymous: e.detail
    })
  },
  selct_IsSquare: function (e) {
    this.setData({
      IsSquare: e.detail
    })
  },
  clonse_switch: function (e) {
    this.setData({
      show_switch: false
    })
    var data = {
      uid: app.globalData.uid,
      is_anonymous: this.data.IsAnonymous ? 1 : 2,
      is_square: this.data.IsSquare ? 1 : 2,
    }
    api.upUserPostSet(data)
    this.pushPost()
  },
  upUserInfo: function () {
    wx.getUserProfile({
      desc: '授权登录',
      success: (res) => {
        if (res.userInfo) {
          var data = {
            uid: app.globalData.uid,
            nick_name: res.userInfo.nickName,
            gender: res.userInfo.gender,
            avatar_url: res.userInfo.avatarUrl,
            encrypted_data: res.encryptedData,
            iv: res.iv,
          }
          api.userUpdateInfo(data).then(res2 => {
            if (res2.code != 1) {
              console.log("userUpdateInfo fail")
              return
            } else {
              app.globalData.isAuth = true
              app.globalData.nickName = res.userInfo.nickName
              app.globalData.avatarUrl = res.userInfo.avatarUrl
              this.setData({
                isAuth: true
              })
            }
          })
        }
      }
    })
  },
  ltrim: function (s) {
    return s.replace(/\s+/g, '');
  },
  isNull: function (obj) {
    if (obj === null) return true;
    if (typeof obj === 'undefined') {
      return true;
    }
    if (typeof obj === 'string') {
      if (obj === "") {
        return true;
      }
      var reg = new RegExp("^([ ]+)|([　]+)$");
      return reg.test(obj);
    }
    return false;
  },
})