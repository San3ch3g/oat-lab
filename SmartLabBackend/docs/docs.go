// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/auth/check-code": {
            "post": {
                "description": "Проверяет код, отправленный на указанный email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Проверка кода",
                "parameters": [
                    {
                        "description": "Запрос на проверку кода",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CheckCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.CheckCodeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.CheckCodeResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/controllers.CheckCodeResponse"
                        }
                    }
                }
            }
        },
        "/auth/send-code": {
            "post": {
                "description": "Отправляет код на email и создаёт запись в бд",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Проверка email",
                "parameters": [
                    {
                        "description": "Запрос для отправки кода и создания записи в бд",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SendCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SendCodeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.SendCodeResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.SendCodeResponse"
                        }
                    }
                }
            }
        },
        "/auth/send-code-again": {
            "post": {
                "description": "Повторно отправляет код подтверждения на email пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Повторная отправка кода подтверждения",
                "parameters": [
                    {
                        "description": "Запрос для повторной отправки кода подтверждения",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SendCodeAgainRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SendCodeAgainResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.SendCodeAgainResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.SendCodeAgainResponse"
                        }
                    }
                }
            }
        },
        "/catalog": {
            "delete": {
                "description": "Удаляет новость из каталога по идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Catalog"
                ],
                "summary": "Удаление новости из каталога",
                "parameters": [
                    {
                        "description": "Запрос для удаления новости из каталога",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.DeleteCatalogItemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.DeleteCatalogItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.DeleteCatalogItemResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.DeleteCatalogItemResponse"
                        }
                    }
                }
            }
        },
        "/catalog/": {
            "get": {
                "description": "Получает новости из каталога по категории",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Catalog"
                ],
                "summary": "Получение новостей из каталога",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Категория для получения новостей из каталога",
                        "name": "category",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetCatalogResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetCatalogResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetCatalogResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Создает новую новость в каталоге",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Catalog"
                ],
                "summary": "Создание новости",
                "parameters": [
                    {
                        "description": "Запрос для создания новости",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateCatalogRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateCatalogResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateCatalogResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateCatalogResponse"
                        }
                    }
                }
            }
        },
        "/catalog/by-id": {
            "get": {
                "description": "Получает элемент каталога по его ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Catalog"
                ],
                "summary": "Получить элемент каталога по ID.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID элемента каталога",
                        "name": "itemId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetCatalogItemByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetCatalogItemByIdResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetCatalogItemByIdResponse"
                        }
                    }
                }
            }
        },
        "/med-card": {
            "get": {
                "description": "Получает информацию о мед картах пользователя по email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MedCards"
                ],
                "summary": "Получение информации о мед картах пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email пользователя",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.MedCardsInfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.MedCardsInfoResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.MedCardsInfoResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновляет информацию мед карты",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MedCards"
                ],
                "summary": "Обновление информации мед карты",
                "parameters": [
                    {
                        "description": "Запрос для создания профиля пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateMedCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateMedCardResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateMedCardResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateMedCardResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Создает мед карты пользователя с указанными данными",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MedCards"
                ],
                "summary": "Создание мед карты пользователя",
                "parameters": [
                    {
                        "description": "Запрос для создания мед карты пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateMedCardRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateMedCardResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateMedCardResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateMedCardResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет мед карты пользователя по идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MedCards"
                ],
                "summary": "Удаление мед карты пользователя",
                "parameters": [
                    {
                        "description": "Запрос для удаления мед карты пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.MedCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.MedCardResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.MedCardResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.MedCardResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CheckCodeRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "controllers.CheckCodeResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "controllers.CreateCatalogRequest": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "preparation": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "timeResult": {
                    "type": "string"
                }
            }
        },
        "controllers.CreateCatalogResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controllers.CreateMedCardRequest": {
            "type": "object",
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "middleName": {
                    "type": "string"
                },
                "sex": {
                    "$ref": "#/definitions/models.SexType"
                }
            }
        },
        "controllers.CreateMedCardResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controllers.DeleteCatalogItemRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "controllers.DeleteCatalogItemResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controllers.GetCatalogItemByIdResponse": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "preparation": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "timeRes": {
                    "type": "string"
                }
            }
        },
        "controllers.GetCatalogResponse": {
            "type": "object",
            "properties": {
                "catalogItems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CatalogItem"
                    }
                },
                "errorMessage": {
                    "type": "string"
                }
            }
        },
        "controllers.MedCardRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "controllers.MedCardResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controllers.MedCardsInfoResponse": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "type": "string"
                },
                "profileInfo": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MedCard"
                    }
                }
            }
        },
        "controllers.SendCodeAgainRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "controllers.SendCodeAgainResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controllers.SendCodeRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "controllers.SendCodeResponse": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controllers.UpdateMedCardRequest": {
            "type": "object",
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "medCardId": {
                    "type": "integer"
                },
                "middleName": {
                    "type": "string"
                },
                "profileImageUrl": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                }
            }
        },
        "controllers.UpdateMedCardResponse": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.CatalogItem": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "category": {
                    "$ref": "#/definitions/models.CatalogItemCategory"
                },
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "preparation": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "timeResult": {
                    "type": "string"
                }
            }
        },
        "models.CatalogItemCategory": {
            "type": "string",
            "enum": [
                "all",
                "news",
                "covid",
                "popular",
                "complex"
            ],
            "x-enum-varnames": [
                "Default",
                "News",
                "Covid",
                "Popular",
                "Complex"
            ]
        },
        "models.MedCard": {
            "type": "object",
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "middleName": {
                    "type": "string"
                },
                "profileImageUrl": {
                    "type": "string"
                },
                "sex": {
                    "$ref": "#/definitions/models.SexType"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.SexType": {
            "type": "string",
            "enum": [
                "M",
                "W"
            ],
            "x-enum-varnames": [
                "MaleSex",
                "WomanSex"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "407e-95-189-144-121.ngrok-free.app",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Smart Lab",
	Description:      "API server for Smart Lab",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
