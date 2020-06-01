// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "null",
        "contact": {
            "name": "null"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/file/{id}": {
            "get": {
                "description": "获取文件",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "文件处理"
                ],
                "summary": "获取某个文件",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文件ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hello": {
            "get": {
                "description": "向你说Hello",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试"
                ],
                "summary": "测试SayHello",
                "parameters": [
                    {
                        "type": "string",
                        "description": "人名",
                        "name": "who",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"hello Razeen\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"who are you\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/index": {
            "post": {
                "description": "null",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "index 测试"
                ],
                "summary": "test index",
                "parameters": [
                    {
                        "description": "查询信息",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Request"
                        }
                    },
                    {
                        "type": "string",
                        "description": "查询的id",
                        "name": "query_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Response"
                        }
                    }
                }
            }
        },
        "/json": {
            "post": {
                "description": "获取JSON的示例",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "JSON"
                ],
                "summary": "获取JSON的示例",
                "parameters": [
                    {
                        "description": "需要上传的JSON",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.JSONParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回",
                        "schema": {
                            "$ref": "#/definitions/main.JSONParams"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "description": "文件列表",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件处理"
                ],
                "summary": "查看文件列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Files"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "登入",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登陆"
                ],
                "summary": "登陆",
                "parameters": [
                    {
                        "type": "string",
                        "default": "admin",
                        "description": "用户名",
                        "name": "user",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"login success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"user or password error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "description": "上传文件",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件处理"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "file",
                        "description": "文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.File"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.File": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "len": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "main.Files": {
            "type": "object",
            "properties": {
                "files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.File"
                    }
                },
                "len": {
                    "type": "integer"
                }
            }
        },
        "main.JSONParams": {
            "type": "object",
            "properties": {
                "array": {
                    "description": "这是一个字符串数组",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "int": {
                    "description": "这是一个数字",
                    "type": "integer"
                },
                "str": {
                    "description": "这是一个字符串",
                    "type": "string"
                },
                "struct": {
                    "description": "这是一个结构",
                    "type": "object",
                    "properties": {
                        "field": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "main.Request": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "main.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "description of your project's api",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}