// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Stas Bukovskyi",
            "email": "stas.bukovskyi@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/images/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Upload image and get its id and url",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Upload Image",
                "responses": {
                    "200": {
                        "description": "uploaded image",
                        "schema": {
                            "$ref": "#/definitions/entity.Image"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/images/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete an image by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "images"
                ],
                "summary": "Delete Image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "image id",
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
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/wishlist-items/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new wishlist item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wishlist-items"
                ],
                "summary": "Create a new wishlist item",
                "parameters": [
                    {
                        "description": "create wishlist item request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AddWishlistItemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "wishlist item",
                        "schema": {
                            "$ref": "#/definitions/entity.WishlistItem"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/wishlist-items/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get wishlist item by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wishlist-items"
                ],
                "summary": "Get wishlist item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wishlist item id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "wishlist item",
                        "schema": {
                            "$ref": "#/definitions/entity.WishlistItem"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create existing wishlist item by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wishlist-items"
                ],
                "summary": "Update a wishlist item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wishlist item id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update wishlist item request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateWishlistItemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "wishlist item",
                        "schema": {
                            "$ref": "#/definitions/entity.WishlistItem"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete wishlist item by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wishlist-items"
                ],
                "summary": "Delete wishlist item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wishlist item id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "wishlist item",
                        "schema": {
                            "$ref": "#/definitions/entity.WishlistItem"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/wishlists/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all wishlists of this user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wishlists"
                ],
                "summary": "List all wishlists",
                "responses": {
                    "200": {
                        "description": "wishlists",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Wishlist"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new, empty wishlist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wishlists"
                ],
                "summary": "Create a new wishlist",
                "parameters": [
                    {
                        "description": "wishlist request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.WishlistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "created wishlist",
                        "schema": {
                            "$ref": "#/definitions/entity.Wishlist"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/wishlists/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get wishlist by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wishlists"
                ],
                "summary": "Get wishlist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wishlist id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "wishlist",
                        "schema": {
                            "$ref": "#/definitions/entity.Wishlist"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a wishlist by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wishlists"
                ],
                "summary": "Update a wishlist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wishlist id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "wishlist request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.WishlistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "updated wishlist",
                        "schema": {
                            "$ref": "#/definitions/entity.Wishlist"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a wishlist by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wishlists"
                ],
                "summary": "Delete a wishlist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wishlist id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "deleted wishlist",
                        "schema": {
                            "$ref": "#/definitions/entity.Wishlist"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/wishlists/{id}/items": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get wishlist items by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wishlists"
                ],
                "summary": "Get wishlist items",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wishlist id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "wishlist items",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.WishlistItem"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Create user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "body request to sign in",
                        "name": "q",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "access token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Create user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "body request to sign up",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httperrs.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Image": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "image_name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "wishlist_item_id": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                }
            }
        },
        "entity.Wishlist": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.WishlistItem": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Image"
                    }
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "wishlist_id": {
                    "type": "string"
                }
            }
        },
        "httperrs.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/httperrs.ServiceError"
                }
            }
        },
        "httperrs.ServiceError": {
            "type": "object",
            "properties": {
                "kind": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "param": {
                    "type": "string"
                }
            }
        },
        "request.AddWishlistItemRequest": {
            "type": "object",
            "required": [
                "url",
                "wishlist_id"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "image_ids": {
                    "type": "array",
                    "maxItems": 5,
                    "items": {
                        "type": "string"
                    }
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "wishlist_id": {
                    "type": "string"
                }
            }
        },
        "request.SignInRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 4
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "request.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 4
                },
                "name": {
                    "type": "string",
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "request.UpdateWishlistItemRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "image_ids": {
                    "type": "array",
                    "maxItems": 5,
                    "items": {
                        "type": "string"
                    }
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "request.WishlistRequest": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "minLength": 3
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go Wishlist API",
	Description:      "This is a simple API to create and manage wishlists with any sort of products or content",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
