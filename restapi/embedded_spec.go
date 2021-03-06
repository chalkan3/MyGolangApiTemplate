// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json",
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Whatsapp communication API",
    "title": "Whats app integration",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "api"
        ],
        "operationId": "getMessages",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "just a test for igor",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/message"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/sendMessage": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "operationId": "sendMessage",
        "parameters": [
          {
            "description": "create a new message",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/resp"
            }
          }
        }
      }
    },
    "/sync": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "operationId": "sync",
        "parameters": [
          {
            "description": "create a new message",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "description": "OK"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "message": {
      "type": "object",
      "properties": {
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "channels": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "channel": {
                "type": "string"
              },
              "chat": {
                "type": "string"
              },
              "language": {
                "type": "string"
              },
              "token": {
                "type": "object",
                "properties": {
                  "phone1": {
                    "type": "string"
                  },
                  "phone2": {
                    "type": "string"
                  }
                }
              }
            }
          }
        },
        "event_date": {
          "type": "string"
        },
        "event_id": {
          "type": "integer",
          "format": "int64"
        },
        "event_time": {
          "type": "string"
        },
        "host_name": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "idd": {
          "type": "string"
        },
        "processed": {
          "type": "boolean"
        },
        "severity": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "tool": {
          "type": "string"
        }
      }
    },
    "resp": {
      "type": "object",
      "properties": {
        "message_queued": {
          "type": "boolean"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json",
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Whatsapp communication API",
    "title": "Whats app integration",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "api"
        ],
        "operationId": "getMessages",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "just a test for igor",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/message"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/sendMessage": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "operationId": "sendMessage",
        "parameters": [
          {
            "description": "create a new message",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/resp"
            }
          }
        }
      }
    },
    "/sync": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "api"
        ],
        "operationId": "sync",
        "parameters": [
          {
            "description": "create a new message",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "description": "OK"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "MessageChannelsItems0": {
      "type": "object",
      "properties": {
        "channel": {
          "type": "string"
        },
        "chat": {
          "type": "string"
        },
        "language": {
          "type": "string"
        },
        "token": {
          "type": "object",
          "properties": {
            "phone1": {
              "type": "string"
            },
            "phone2": {
              "type": "string"
            }
          }
        }
      }
    },
    "MessageChannelsItems0Token": {
      "type": "object",
      "properties": {
        "phone1": {
          "type": "string"
        },
        "phone2": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "message": {
      "type": "object",
      "properties": {
        "actions": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "channels": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/MessageChannelsItems0"
          }
        },
        "event_date": {
          "type": "string"
        },
        "event_id": {
          "type": "integer",
          "format": "int64"
        },
        "event_time": {
          "type": "string"
        },
        "host_name": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "idd": {
          "type": "string"
        },
        "processed": {
          "type": "boolean"
        },
        "severity": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "tool": {
          "type": "string"
        }
      }
    },
    "resp": {
      "type": "object",
      "properties": {
        "message_queued": {
          "type": "boolean"
        }
      }
    }
  }
}`))
}
