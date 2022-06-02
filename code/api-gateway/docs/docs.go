// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "get": {
                "description": "微服务模块 Order 中提供的获取订单列表服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order 服务"
                ],
                "summary": "获取订单列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户 id",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "偏移量",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "限制数量",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"message\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "微服务模块 Order 中提供的创建订单服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order 服务"
                ],
                "summary": "创建订单",
                "parameters": [
                    {
                        "description": "订单",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.OrderCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"message\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "微服务模块 User 中提供的用户登录服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User 服务"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"message\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/orders": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "User 服务中提供的用户订单列表服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User 服务"
                ],
                "summary": "用户订单列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "User 服务中提供的用户创建订单服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User 服务"
                ],
                "summary": "用户创建订单",
                "parameters": [
                    {
                        "description": "订单",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.OrderCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "微服务模块 User 中提供的用户注册服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User 服务"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "注册",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"message\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.OrderCreateReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "service.OrderCreateRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "@inject_tag: json:\"name\" form:\"name\" uri:\"name\"",
                    "type": "string"
                },
                "user_id": {
                    "description": "@inject_tag: json:\"user_id\" form:\"user_id\" uri:\"user_id\"",
                    "type": "integer"
                }
            }
        },
        "service.UserLoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "@inject_tag: json:\"password\" form:\"password\" uri:\"password\"",
                    "type": "string"
                },
                "username": {
                    "description": "@inject_tag: json:\"username\" form:\"username\" uri:\"username\"",
                    "type": "string"
                }
            }
        },
        "service.UserRegisterRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "@inject_tag: json:\"password\" form:\"password\" uri:\"password\"",
                    "type": "string"
                },
                "password_confirm": {
                    "description": "@inject_tag: json:\"password_confirm\" form:\"password_confirm\" uri:\"password_confirm\"",
                    "type": "string"
                },
                "username": {
                    "description": "@inject_tag: json:\"username\" form:\"username\" uri:\"username\"",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger API-Gateway API 入口",
	Description:      "网关 API 入口.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
