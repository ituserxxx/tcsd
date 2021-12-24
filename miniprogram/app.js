const api =  require('./api/api') ;
// app.js
App({
  globalData: {
    uid: 0,
    isAuth:false,
    nickName:"",
    avatarUrl:"",
    postaSum:0,
  },
  onLaunch() {
    wx.login({
      success: res => {
        api.userLogin({
          code:res.code
        }).then(res=>{
          this.globalData.uid = res.data.uid
          this.globalData.isAuth = res.data.is_auth
          this.globalData.nickName = res.data.nick_name
          this.globalData.avatarUrl = res.data.avatar_url  
          this.globalData.postaSum = res.data.post_sum
        }).catch(e=>{
          console.log(res)
        })
      }
    })
  },
})
