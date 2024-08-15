// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Kiselyov Vladimir"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/image/upload": {
            "post": {
                "description": "Uploads the image to the task",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "image"
                ],
                "summary": "Upload image",
                "parameters": [
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "task_id",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Uploaded image",
                        "name": "file",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/image.ResponseUploadImage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "423": {
                        "description": "Locked",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/ping": {
            "post": {
                "description": "Returns \"pong\" if the server is healthy",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCommon"
                        }
                    }
                }
            }
        },
        "/api/task": {
            "get": {
                "description": "Returns a task with statistics and a list of images and their data on recognized faces",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Get task",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/task.RequestGetTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task.ResponseGetTask"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete task by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Delete task",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/task.RequestDeleteTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task.ResponseDeleteTask"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "423": {
                        "description": "Locked",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/task/add": {
            "post": {
                "description": "Creates a task with the passed name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Add task",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/task.RequestAddTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task.ResponseAddTask"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/task/process": {
            "post": {
                "description": "Processes images from the task and calculates statistics for the task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Process task",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/task.RequestProcessTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task.ResponseProcessTask"
                        }
                    },
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/task.ResponseProcessTask"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "image.ResponseUploadImage": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "models.ResponseCommon": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "models.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "task.FaceData": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "bounding_box": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "sex": {
                    "$ref": "#/definitions/task.Sex"
                }
            }
        },
        "task.ImageData": {
            "type": "object",
            "properties": {
                "faces": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/task.FaceData"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "task.RequestAddTask": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "task.RequestDeleteTask": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "task.RequestGetTask": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "task.RequestProcessTask": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "task.ResponseAddTask": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "task.ResponseDeleteTask": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "task.ResponseGetTask": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/task.ImageData"
                    }
                },
                "statistic": {
                    "$ref": "#/definitions/task.Statistic"
                },
                "status": {
                    "$ref": "#/definitions/task.Status"
                }
            }
        },
        "task.ResponseProcessTask": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "task.Sex": {
            "type": "string",
            "enum": [
                "male"
            ],
            "x-enum-varnames": [
                "Male"
            ]
        },
        "task.Statistic": {
            "type": "object",
            "properties": {
                "avg_female_age": {
                    "type": "integer"
                },
                "avg_male_age": {
                    "type": "integer"
                },
                "face_count": {
                    "type": "integer"
                },
                "female_count": {
                    "type": "integer"
                },
                "male_count": {
                    "type": "integer"
                }
            }
        },
        "task.Status": {
            "type": "string",
            "enum": [
                "Being formed"
            ],
            "x-enum-varnames": [
                "Forming"
            ]
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "TevianTestTask API documentation",
	Description:      "This is the server for the Tevian test task.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
