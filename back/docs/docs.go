// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-03-12 22:02:56.411483541 +0100 CET m=+0.081028564

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/navigation/dir": {
            "get": {
                "description": "Gets the navigation route for a start and end room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "NavigationController"
                ],
                "summary": "Get Navigation Route",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the start room name",
                        "name": "startroom",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "the end room name",
                        "name": "endroom",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/navigation.Coordinate"
                            }
                        }
                    }
                }
            }
        },
        "/roomlist/": {
            "get": {
                "description": "Gets all available rooms",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RoomController"
                ],
                "summary": "Get Room List",
                "operationId": "get-room-list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Room"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Door": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "pathNode": {
                    "type": "object",
                    "$ref": "#/definitions/navigation.PathNode"
                },
                "section": {
                    "type": "object",
                    "$ref": "#/definitions/models.Section"
                }
            }
        },
        "models.Room": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "color": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "doors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Door"
                    }
                },
                "floor": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "pathNode": {
                    "type": "object",
                    "$ref": "#/definitions/navigation.PathNode"
                },
                "sections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Section"
                    }
                }
            }
        },
        "models.Section": {
            "type": "object",
            "properties": {
                "end": {
                    "type": "object",
                    "$ref": "#/definitions/navigation.Coordinate"
                },
                "id": {
                    "type": "integer"
                },
                "start": {
                    "type": "object",
                    "$ref": "#/definitions/navigation.Coordinate"
                }
            }
        },
        "navigation.Coordinate": {
            "type": "object",
            "properties": {
                "x": {
                    "type": "integer"
                },
                "y": {
                    "type": "integer"
                },
                "z": {
                    "type": "integer"
                }
            }
        },
        "navigation.PathNode": {
            "type": "object",
            "properties": {
                "connectedNodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/navigation.PathNode"
                    }
                },
                "coordinate": {
                    "type": "object",
                    "$ref": "#/definitions/navigation.Coordinate"
                },
                "group": {
                    "type": "object",
                    "$ref": "#/definitions/navigation.PathNodeGroup"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "navigation.PathNodeGroup": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/navigation.PathNode"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
