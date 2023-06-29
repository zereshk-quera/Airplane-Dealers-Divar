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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/ads": {
            "get": {
                "description": "Retrieves all ads from database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ads"
                ],
                "summary": "Get ads.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Ad"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid parameter id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Could not retrieve ads",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ads/{id}": {
            "get": {
                "description": "Retrieves an ad based on the provided ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ads"
                ],
                "summary": "Get ad by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Ad ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Ad"
                        }
                    },
                    "400": {
                        "description": "Invalid parameter id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Could not retrieve ads",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Ad": {
            "type": "object",
            "properties": {
                "airplaneModel": {
                    "type": "string"
                },
                "category": {
                    "$ref": "#/definitions/models.Category"
                },
                "categoryID": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "expertCheck": {
                    "type": "boolean"
                },
                "flyTime": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "planeAge": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "repairCheck": {
                    "type": "boolean"
                },
                "status": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Bookmarks": {
            "type": "object",
            "properties": {
                "ads": {
                    "$ref": "#/definitions/models.Ad"
                },
                "adsID": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "bookmarks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Bookmarks"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "integer"
                },
                "username": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Airplane-Divar",
	Description:      "Quera Airplane-Divar server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
