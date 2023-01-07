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
        "contact": {
            "name": "Aulia Wiguna",
            "url": "https://github.com/auliawiguna/",
            "email": "wigunaahmadaulia@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1": {
            "get": {
                "description": "Show greeting",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Root"
                ],
                "summary": "Say hi",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/notes": {
            "get": {
                "description": "Show notes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Show notes",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Default 10",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Default 10",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "ID asc",
                            "ID desc",
                            "title asc",
                            "title desc"
                        ],
                        "type": "string",
                        "description": "Sorting",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "description": "Create new note",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Create new note",
                "parameters": [
                    {
                        "description": "title",
                        "name": "notesRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.NoteCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/notes/{id}": {
            "get": {
                "description": "Show detail note",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Show detail note",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "put": {
                "description": "Update existing note",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Update existing note",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "ID",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "description": "title",
                        "name": "notesRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.NoteCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete existing note",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notes"
                ],
                "summary": "Delete existing note",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structs.NoteCreate": {
            "type": "object",
            "properties": {
                "subtitle": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Goshaka Golang API Starter",
	Description:      "This is a API boilerplate using Golang",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
