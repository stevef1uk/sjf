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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A simple API to learn how to write OpenAPI Specification",
    "title": "Simple API",
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/v1",
  "paths": {
    "/accounts": {
      "get": {
        "description": "Returns a list containing all accounts.",
        "summary": "Gets some accounts",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "PK",
            "name": "id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "string",
            "description": "PK",
            "name": "name",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A list of records",
            "schema": {
              "type": "array",
              "items": {
                "required": [
                  "id",
                  "name"
                ],
                "properties": {
                  "id": {
                    "type": "integer"
                  },
                  "name": {
                    "type": "string"
                  }
                }
              }
            }
          },
          "400": {
            "description": "Record not found"
          },
          "default": {
            "description": "Sorry something went wrong"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A simple API to learn how to write OpenAPI Specification",
    "title": "Simple API",
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/v1",
  "paths": {
    "/accounts": {
      "get": {
        "description": "Returns a list containing all accounts.",
        "summary": "Gets some accounts",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "PK",
            "name": "id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "string",
            "description": "PK",
            "name": "name",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A list of records",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/getAccountsOKBodyItems"
              }
            }
          },
          "400": {
            "description": "Record not found"
          },
          "default": {
            "description": "Sorry something went wrong"
          }
        }
      }
    }
  },
  "definitions": {
    "getAccountsOKBodyItems": {
      "required": [
        "id",
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    }
  }
}`))
}
