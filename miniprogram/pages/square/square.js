import Toast from '../../miniprogram_npm/@vant/weapp/toast/toast';
const api = require("../../api/api.js");
const app = getApp()
Page({
  data: {
    post_list: [
      // {
      //   master_info: {
      //     avatar_url: "https://thirdwx.qlogo.cn/mmopen/vi_32/vlBrF1UadszIEGzzePcicuDg2Vic0VdokPiaq7QaSPR5PicafOJLG0BLW8icJOmpt0ey792I4vhRKJzPNOU7sbJFayg/132",
      //     nick_name: "巨石糖果山",
      //     id: 1,
      //   },
      //   post_info: {
      //     img_url: [
      //       "https://thirdwx.qlogo.cn/mmopen/vi_32/vlBrF1UadszIEGzzePcicuDg2Vic0VdokPiaq7QaSPR5PicafOJLG0BLW8icJOmpt0ey792I4vhRKJzPNOU7sbJFayg/132",
      //       "https://thirdwx.qlogo.cn/mmopen/vi_32/vlBrF1UadszIEGzzePcicuDg2Vic0VdokPiaq7QaSPR5PicafOJLG0BLW8icJOmpt0ey792I4vhRKJzPNOU7sbJFayg/132",
      //       "https://thirdwx.qlogo.cn/mmopen/vi_32/vlBrF1UadszIEGzzePcicuDg2Vic0VdokPiaq7QaSPR5PicafOJLG0BLW8icJOmpt0ey792I4vhRKJzPNOU7sbJFayg/132"
      //     ],
      //     content: "内容",
      //     id: 1,
      //     is_give_love:true,
      //      love_sum:1,
      //   }
      // },
    ],
    page: 1,
    show_share: false,
    share_options: [
      { name: '微信好友', icon: 'wechat', openType: 'share' },
    ],
    
  },
  onLoad() {
    this.getPostList(this.data.page)
  },
  onPullDownRefresh: function () {
    this.getPostList(this.data.page)
  },
  onReachBottom: function () {
    this.setData({
      page: ++this.data.page
    })
    this.getPostList(this.data.page)
  },

  isPostMaster: function (poMasterId) {
    console.log("判断是否是主人", poMasterId, app.globalData.uid)
    return poMasterId == app.globalData.uid
  },
  getPostList: function (page) {
    api.postList({
      page: page,
      uid: app.globalData.uid
    }).then(res => {
      if (res.code == 1 && res.data.list != null) {
        var arr1 = this.data.post_list
        this.setData({
          post_list: arr1.concat(res.data.list),
          page: this.data.page
        })
      } else {
        Toast.fail('加载失败了~');
      }
    })
  },
  giveLove: function (e) {
    var postId = e.currentTarget.dataset.post_id
    console.log(postId, app.globalData.uid)
    api.giveLove({
      uid: app.globalData.uid,
      post_id: postId
    }).then(res => {
      if (res.code == 1) {
        Toast.success('没想到你也觉得很赞~ye');
      } else {
        Toast.fail('等会再试吧~');
      }
    })
  },
  sharePost: function (e) {
    var postId = e.currentTarget.dataset.post_id
    // console.log(postId,app.globalData.uid)
    this.setData({
      show_share: true
    })
  },
  close_share: function () {
    this.setData({
      show_share: false
    })
  }
})
