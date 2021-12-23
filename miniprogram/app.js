const api =  require('./api/api') ;
// app.js
App({
  onLaunch() {
    // 登录
    wx.login({
      success: res => {
        api.userLogin({
          code:res.code
        }).then(res=>{
          console.log(res)
          this.globalData.uid = res.data
        }).catch(e=>{
          console.log(res)
        })
      }
    })
  },
  globalData: {
    uid: 0
  },
})
