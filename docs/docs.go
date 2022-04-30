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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Putu Prima Winangun",
            "email": "primawinangun@gmail.com"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/valid-time/process": {
            "post": {
                "description": "return all available time from the combination of four value of slice of integer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "find all combination of valid time from slice of integer",
                "parameters": [
                    {
                        "description": "string value with space between, ex: 1 2 3 4",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ValidTimeResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/word-count/upload": {
            "post": {
                "description": "return top ten most used words along with how many times they occur in the text.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "upload file via rest api http \u0026 return top ten most used words with the number of used",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file type with words contained",
                        "name": "payload",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.WordCount"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.request": {
            "type": "object",
            "properties": {
                "input": {
                    "type": "string"
                }
            }
        },
        "model.ValidTimeResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.WordCount": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "word": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Ulventech Technical Test",
	Description:      "This is a simple application for ulventech technical test.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}