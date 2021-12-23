// pages/release/release.js
const app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        userInfo: {},
        inputText:"",
        isAuthInfo: false,
        canIUse: wx.canIUse('button.open-type.getUserInfo'),
        canIUseGetUserProfile: false,
    },
    onLoad() {
        // if (wx.getUserProfile) {
        //   this.setData({
        //     canIUseGetUserProfile: true
        //   })
        // }
        // console.log(this.getUserProfile())

        console.log()
      },
      getUserProfile(e) {
        console.log(111)
        // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认
        // 开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
        wx.getUserProfile({
          desc: '授权登录', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
          success: (res) => {
            console.log(res.userInfo,111)
            this.setData({
              userInfo: res.userInfo,
              isAuthInfo: true
            })
          }
        })
      },

    bindFormSubmit: function(e) {  
        wx.request({
            url: app.data.url+'api/post/create',
            data: {
              uid:app.getUserInfo().uid,
              context:e.detail.value.textarea,
              // is_anonymous 必传  1-开发，2-匿名
              is_anonymous:1
            },
            success:(res)=> {
              if(res.data.code==0){
             
                wx.showToast({
                  title: '发布成功',
                  icon: 'none',    
                  duration: 2000     
                })  
                this.setData({
                  inputText:""
                })
              }else{
                
                wx.showToast({
                  title: res.data.res,
                  icon: 'none',    
                  duration: 2000     
                })  
              }
            }
        })
        
        
       
        // 
        // wx.showModal({
        // title: '提示',
        // content: '微信登录才能发布信息',
        // success: function (res) {

        //     if (res.confirm) { 
        //         wx.login({
        //             success (res) {
        //               if (res.code) {
        //                 //发起网络请求
        //                console.log(res.code)
        //               } else {
        //                 console.log('登录失败！' + res.errMsg)
        //               }
        //             }
        //           })
        //     } else {   
        //     // console.log('点击取消回调')
        //     }
        // }
        // })

    
    },
   
      
})