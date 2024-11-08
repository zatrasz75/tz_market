// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Михаил Токмачев",
            "url": "https://t.me/Zatrasz",
            "email": "zatrasz@ya.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/en/building": {
            "get": {
                "description": "Принимает обязательные city, year_built, а так же не обязательные floors.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Building"
                ],
                "summary": "Возвращает список строений, с возможностью фильтрации по городу, году и кол-ву этажей(параметры не обязательные)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "по названию города",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "по году сдачи",
                        "name": "year_built",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "кол-ву этажей",
                        "name": "floors",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список домов",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Building"
                            }
                        }
                    },
                    "400": {
                        "description": "обязательные поля отсутствуют",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при получении данных",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Принимает обязательные поля name , city , year_built , а так же не обязательные floors.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Building"
                ],
                "summary": "Создайте новой записи",
                "parameters": [
                    {
                        "description": "Данные здания",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Building"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "запись успешно добавлена",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "обязательные поля отсутствуют",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "название уже существует",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при добавлении данных",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Building": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "floors": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "year_built": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger API",
	Description:      "ТЗ market_backend.\nЗадание (Golang + PostgreSQL + Gin)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
