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
            "url": "https://github.com/daniferdinandall",
            "email": "1214050@std.ac.id"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/presensi": {
            "get": {
                "description": "Mengambil semua data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get All Data Presensi.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.JamKerja": {
            "type": "object",
            "properties": {
                "durasi": {
                    "type": "integer"
                },
                "gmt": {
                    "type": "integer"
                },
                "hari": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "jam_keluar": {
                    "type": "string"
                },
                "jam_masuk": {
                    "type": "string"
                },
                "piket_tim": {
                    "type": "string"
                },
                "shift": {
                    "type": "integer"
                }
            }
        },
        "controller.Karyawan": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "hari_kerja": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "jabatan": {
                    "type": "string"
                },
                "jam_kerja": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.JamKerja"
                    }
                },
                "nama": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "controller.Presensi": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "biodata": {
                    "$ref": "#/definitions/controller.Karyawan"
                },
                "checkin": {
                    "description": "Datetime     primitive.DateTime ` + "`" + `bson:\"datetime,omitempty\" json:\"datetime,omitempty\"` + "`" + `",
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "location": {
                    "type": "string"
                },
                "longitude": {
                    "type": "number"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "ws-dani.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "TES SWAG",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
