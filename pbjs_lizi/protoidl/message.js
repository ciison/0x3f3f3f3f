// 导出定义的信息
module.exports = {
    "nested": {
        "AMiniResponse": {
            "fields": {
                "openId": {
                    "type": "string",
                    "id": 1
                },
                "jwt": {
                    "type": "string",
                    "id": 2
                },
                "address": {
                    "type": "string",
                    "id": 3
                },
                "age": {
                    "type": "int64",
                    "id": 4
                }
            }
        },
        "AMiniPostRequest": {
            "fields": {
                "jwt": {
                    "type": "string",
                    "id": 1
                },
                "category": {
                    "type": "string",
                    "id": 2
                }
            }
        },
        "AMiniPostResponse": {
            "fields": {
                "errCode": {
                    "type": "int32",
                    "id": 1
                },
                "errMessage": {
                    "type": "string",
                    "id": 2
                },
                "category": {
                    "rule": "repeated",
                    "type": "string",
                    "id": 3
                }
            }
        }
    }
}