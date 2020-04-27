//index.js
//获取应用实例
const app = getApp()
const BaseUrl = "http://localhost:8080"
console.log(BaseUrl)
let protoBuff = require('../../protobufjs/protobuf') // 引入 protoBuff 的库
let messageConfig = require('./message')
let messageRoot = protoBuff.Root.fromJSON(messageConfig)
console.log(messageRoot)
// proto 定义的类型， 这里的类型没有 package 
let AMiniResponse = messageRoot.lookupType('AMiniResponse')
console.log(AMiniResponse)
let AMiniPostRequest = messageRoot.lookupType('AMiniPostRequest')
console.log(AMiniPostRequest)
let AMiniPostResponse = messageRoot.lookupType('AMiniPostResponse')
console.log(AMiniPostResponse)

function Uint8ArrayToString(fileData) {
  let dataString = "";
  for (let i = 0; i < fileData.length; i++) {
    dataString += String.fromCharCode(fileData[i]);
  }
  return dataString
}

function StringToUint8Array(str) {
  let arr = [];
  for (let i = 0, j = str.length; i < j; ++i) {
    arr.push(str.charCodeAt(i));
  }
  let tmpUint8Array = new Uint8Array(arr);
  return tmpUint8Array
}
Page({
  data: {
    aMiniPostResponse:{},
    aMiniPostRequest:{},
    aMiniResponse:{}
  },

  onLoad: function () {},
  getQeury(evt) {
    //console.log(evt)
    wx.request({
      url: BaseUrl + '/query',
      method: 'GET',
      success: (res) => {
        console.log(res);
        let buffer = StringToUint8Array(res.data)
        console.log(buffer)
        let rb = AMiniResponse.decode(buffer)
        console.log(rb)
      }
    })
  },
  postQuery(evt) {
    //console.log(evt);
    wx.request({
      url: BaseUrl + '/query',
      method: 'POST',
      success: (res) => {
        console.log(res)
        let data = StringToUint8Array(res.data)
        let a = AMiniPostResponse.decode(data)
        console.log(a)
        this.setData({
          aMiniPostResponse:a,
        })
        // console.log(data)
      }
    })
  }

})