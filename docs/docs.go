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
        "/auth/check-user": {
            "post": {
                "description": "Проверяет, зарегистрирован ли пользователь с указанным email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Проверка пользователя",
                "parameters": [
                    {
                        "description": "Запрос для проверки пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CheckUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.CheckUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.CheckUserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.CheckUserResponse"
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
        "/auth/sign-in": {
            "post": {
                "description": "Выполняет вход пользователя в систему с использованием email и пароля",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Вход пользователя",
                "parameters": [
                    {
                        "description": "Запрос для входа пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.SignInResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/controllers.SignInResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Регистрирует нового пользователя с использованием email и пароля",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Регистрация пользователя",
                "parameters": [
                    {
                        "description": "Запрос для регистрации пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.SignUpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.SignUpResponse"
                        }
                    }
                }
            }
        },
        "/catalog": {
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
            },
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
        "/catalog/get": {
            "post": {
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
                        "description": "Запрос для получения новостей из каталога",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.GetCatalogRequest"
                        }
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
            }
        },
        "/profile": {
            "get": {
                "description": "Получает информацию о профиле пользователя по email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profiles"
                ],
                "summary": "Получение информации о профиле пользователя",
                "parameters": [
                    {
                        "description": "Запрос для получения информации о профиле пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ProfileInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.ProfileInfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ProfileInfoResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ProfileInfoResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Создает профиль пользователя с указанными данными",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profiles"
                ],
                "summary": "Создание профиля пользователя",
                "parameters": [
                    {
                        "description": "Запрос для создания профиля пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateProfileResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateProfileResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет профиль пользователя по идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profiles"
                ],
                "summary": "Удаление профиля пользователя",
                "parameters": [
                    {
                        "description": "Запрос для удаления профиля пользователя",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.DeleteProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.DeleteProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.DeleteProfileResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.DeleteProfileResponse"
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
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controllers.CheckUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "controllers.CheckUserResponse": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "type": "string"
                },
                "isRegistered": {
                    "type": "boolean"
                }
            }
        },
        "controllers.CreateCatalogRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
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
        "controllers.CreateProfileRequest": {
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
                "profileImage": {
                    "type": "string"
                },
                "sex": {
                    "$ref": "#/definitions/models.SexType"
                }
            }
        },
        "controllers.CreateProfileResponse": {
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
        "controllers.DeleteProfileRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "controllers.DeleteProfileResponse": {
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
        "controllers.GetCatalogRequest": {
            "type": "object",
            "properties": {
                "category": {
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
        "controllers.ProfileInfoRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "controllers.ProfileInfoResponse": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "type": "string"
                },
                "profileInfo": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Profile"
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
        "controllers.SignInRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.SignInResponse": {
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
        "controllers.SignUpRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.SignUpResponse": {
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
        "models.CatalogItem": {
            "type": "object",
            "properties": {
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
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
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
        "models.Profile": {
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
	Host:             "b683-2a03-d000-6583-1ba0-48f1-c600-58b8-2a33.ngrok-free.app",
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
