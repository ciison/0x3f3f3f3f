//index.js
//获取应用实例
const app = getApp()
const BaseUrl = "http://localhost:8080"
console.log(BaseUrl)

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
  data: {},

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
        let data = StringToUint8Array(res.data);
        console.log(data)
      }
    })
  }

})