//index.js
//获取应用实例
const app = getApp()
const BaseUrl = "http://localhost:8080"
console.log(BaseUrl)
let protoBuff = require('../../protobufjs/protobuf') // 引入 protoBuff 的库
let messageConfig = require('./message')
let messageRoot = protoBuff.Root.fromJSON(messageConfig)
console.log(messageRoot)
// proto 定义的类型， 这里的类型没有 package 定义这个 proto 文件时没有写上 package 名称;
let AMiniResponse = messageRoot.lookupType('AMiniResponse')
console.log(AMiniResponse)
let AMiniPostRequest = messageRoot.lookupType('AMiniPostRequest')
console.log(AMiniPostRequest)
let AMiniPostResponse = messageRoot.lookupType('AMiniPostResponse')
console.log(AMiniPostResponse)

// 引入 err_report.js 文件
let errReportConfig = require('./err_report')
let errReportRoot = protoBuff.Root.fromJSON(errReportConfig)

let errReportRequestType = errReportRoot.lookupType('err.ErrReportRequest')
let errReportErrInfoType = errReportRoot.lookupType('err.ErrInfo')
let errReportResponseType = errReportRoot.lookupType('err.ErrReportResponse')
let serverInfoType = errReportRoot.lookupType('err.ServerInfo')

console.log(errReportRequestType)
console.log(errReportErrInfoType)
console.log(errReportResponseType)
console.log(serverInfoType)

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
        aMiniPostResponse: {},
        aMiniPostRequest: {},
        aMiniResponse: {},
        errReportResponse: {}
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
                    aMiniPostResponse: a,
                })
            }
        })
    },
    // 复杂的请求
    postComplex(evt) {
        let pageUrl = '/index/index'
        let appId = '12345'
        let errInfo = {
            pageUrl,
            appId,
            errInfo: 'userClickPage',
        }
        let accessToken = 'jsx'
        let message = {
            errInfo,
            accessToken,
        }
        console.log(message)
        let buffer = errReportRequestType.encode(message).finish()
        let str = Uint8ArrayToString(buffer)
        wx.request({
            url: BaseUrl + '/complex/err',
            method: 'POST',
            header: {
                'content-type': 'text/plaintext'
            },
            data: str,
            success: (res) => {
                console.log(res)
                let data = StringToUint8Array(res.data)
                let errReportResponse = errReportResponseType.decode(data)
                console.log(errReportResponse)
                let errCode = errReportResponse.errCode
                let errMessage = errReportResponse.errMessage
                let serverInfo = errReportResponse.serverInfo
                let msg = {
                    errCode,
                    errMessage,
                    serverInfo,
                }
                this.setData({
                    errReportResponse: msg
                })
            }
        })
    }
})