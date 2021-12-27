const api =  require('./api/api') ;
// app.js
App({
  globalData: {
    uid: 0,
    isAuth:false,
    nickName:"",
    avatarUrl:"",
    postaSum:0,
    PageActive: true
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
  preventActive (fn) {
    const self = this
    if (this.globalData.PageActive) {
      this.globalData.PageActive = false
      if (fn) fn()
      setTimeout(() => {
        self.globalData.PageActive = true
      }, 3000); //设置该时间内重复触发只执行第一次，单位ms，按实际设置
    } else {
      return "慢点，慢点~~"
    }
  }
})
