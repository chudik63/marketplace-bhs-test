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
        "/marketplace/assets": {
            "post": {
                "description": "Create a new asset",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "assets"
                ],
                "summary": "Create an asset",
                "parameters": [
                    {
                        "description": "Asset data",
                        "name": "asset",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/marketplace-bhs-test_internal_entity.Asset"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/marketplace-bhs-test_internal_entity.Asset"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/marketplace/assets/{id}": {
            "delete": {
                "description": "Delete an asset by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "assets"
                ],
                "summary": "Delete an asset",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Asset ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Purchase an asset by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "assets"
                ],
                "summary": "Buy an asset",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Asset ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/sign-in": {
            "post": {
                "description": "Authenticate a user and return access and refresh tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Sign in a user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/marketplace-bhs-test_internal_service.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
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
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/sign-out": {
            "post": {
                "description": "Clear the user's access and refresh tokens",
                "tags": [
                    "users"
                ],
                "summary": "Sign out a user",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "description": "Register a new user with the provided credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Sign up a new user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/marketplace-bhs-test_internal_service.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/users/{id}/balance/{count}": {
            "patch": {
                "description": "Update the user's balance by a specified count",
                "tags": [
                    "users"
                ],
                "summary": "Update user's balance",
                "parameters": [
                    {
                        "type": "number",
                        "description": "Balance change amount",
                        "name": "count",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "marketplace-bhs-test_internal_entity.Asset": {
            "description": "Information about an asset in the marketplace",
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "description": "Detailed description of the asset\nRequired: true",
                    "type": "string",
                    "example": "Digital artwork"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "Unique name for the asset\nRequired: true",
                    "type": "string",
                    "example": "Artwork"
                },
                "price": {
                    "description": "Price of the asset in the marketplace\nRequired: true",
                    "type": "number",
                    "example": 100.5
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_id": {
                    "description": "ID of the user who owns the asset\nRequired: true",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "marketplace-bhs-test_internal_service.SignUpInput": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Marketplace-BHS-test",
	Description:      "test tast for Marketplace-Hive project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
