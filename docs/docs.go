// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/frame": {
            "get": {
                "description": "Extract Frame as File based on episode and frame number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/jpeg"
                ],
                "tags": [
                    "extract"
                ],
                "summary": "Extract Frame as File",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Video Name",
                        "name": "video_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1",
                        "description": "Episode",
                        "name": "episode",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Frame Number",
                        "name": "frame",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/gif": {
            "get": {
                "description": "Extract GIF or WebM as File based on episode, start, and end",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/gif",
                    "video/webm"
                ],
                "tags": [
                    "extract"
                ],
                "summary": "Extract GIF or WebM",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Video Name",
                        "name": "video_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Episode",
                        "name": "episode",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Start",
                        "name": "start",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "End",
                        "name": "end",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Format (gif or webm)",
                        "name": "format",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/search": {
            "post": {
                "description": "Search for sentences based on query and other parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Search for sentences",
                "parameters": [
                    {
                        "description": "Search parameters",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SearchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SearchResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.SearchRequest": {
            "type": "object",
            "properties": {
                "episode": {
                    "type": "string"
                },
                "nth_page": {
                    "type": "integer"
                },
                "paged_by": {
                    "type": "integer"
                },
                "query": {
                    "type": "string"
                },
                "video_name": {
                    "enum": [
                        "Ave Mujica",
                        "MyGO"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.VideoNameEnum"
                        }
                    ]
                }
            }
        },
        "models.SearchResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SentenceItem"
                    }
                }
            }
        },
        "models.SentenceItem": {
            "type": "object",
            "properties": {
                "episode": {
                    "type": "string"
                },
                "frame_end": {
                    "type": "integer"
                },
                "frame_start": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "segment_id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "video_name": {
                    "type": "string"
                }
            }
        },
        "models.VideoNameEnum": {
            "type": "string",
            "enum": [
                "Ave Mujica",
                "MyGO"
            ],
            "x-enum-varnames": [
                "AveMujica",
                "MyGO"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "MyGO Backend API",
	Description:      "This is a server for MyGO Sentence Search and Image/GIF Extraction.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
