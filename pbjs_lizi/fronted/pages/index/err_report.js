module.exports ={
    "nested": {
        "err": {
            "nested": {
                "ErrReportRequest": {
                    "fields": {
                        "errInfo": {
                            "type": "err.ErrInfo",
                            "id": 1
                        },
                        "accessToken": {
                            "type": "string",
                            "id": 2
                        }
                    }
                },
                "ErrReportResponse": {
                    "fields": {
                        "errCode": {
                            "type": "int64",
                            "id": 1
                        },
                        "errMessage": {
                            "type": "string",
                            "id": 2
                        },
                        "serverInfo": {
                            "type": "ServerInfo",
                            "id": 3
                        }
                    }
                },
                "ServerInfo": {
                    "fields": {
                        "version": {
                            "type": "string",
                            "id": 1
                        }
                    }
                },
                "ErrInfo": {
                    "fields": {
                        "appId": {
                            "type": "string",
                            "id": 1
                        },
                        "pageUrl": {
                            "type": "string",
                            "id": 2
                        },
                        "errInfo": {
                            "type": "string",
                            "id": 3
                        },
                        "errTime": {
                            "type": "int64",
                            "id": 4
                        }
                    }
                }
            }
        }
    }
}