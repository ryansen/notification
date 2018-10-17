package proto

const (
	Swagger = `  {
  "swagger": "2.0",
  "info": {
    "title": "package notification;",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/hello": {
      "post": {
        "summary": "Sends a greeting",
        "operationId": "SayHello",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/nf_protoHelloReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/nf_protoHelloRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/nf/CreateNfWAppFilter": {
      "post": {
        "summary": "create notification with App Filter",
        "operationId": "CreateNfWAppFilter",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/nf_protoCreateNfResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/nf_protoCreateNfWAppFilterRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/nf/CreateNfWUserFilter": {
      "post": {
        "summary": "create notification with User Filter",
        "operationId": "CreateNfWUserFilter",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/nf_protoCreateNfResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/nf_protoCreateNfWUserFilterRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/nf/CreateNfWaddrs": {
      "post": {
        "summary": "create notification with addrs(email addrs, phone numbers)",
        "operationId": "CreateNfWaddrs",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/nf_protoCreateNfResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/nf_protoCreateNfWaddrsRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/nf/DescribeNfs": {
      "post": {
        "summary": "describe Notification with filter",
        "operationId": "DescribeNfs",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/nf_protoDescribeNfsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/nf_protoDescribeNfsRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/nf/DescribeUserNfs": {
      "post": {
        "summary": "describe User Notification with filter",
        "operationId": "DescribeUserNfs",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/nf_protoDescribeNfsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/nf_protoDescribeNfsRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    }
  },
  "definitions": {
    "nf_protoCreateNfResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "title": "google.protobuf.StringValue nf_post_id = 1;"
        }
      }
    },
    "nf_protoCreateNfWAppFilterRequest": {
      "type": "object",
      "properties": {
        "nf_post_type": {
          "type": "string"
        },
        "notify_type": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "short_content": {
          "type": "string"
        },
        "expired_days": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "app_id": {
          "type": "string"
        },
        "app_versions": {
          "type": "string"
        },
        "cluster_status": {
          "type": "string"
        }
      }
    },
    "nf_protoCreateNfWUserFilterRequest": {
      "type": "object",
      "properties": {
        "nf_post_type": {
          "type": "string"
        },
        "notify_type": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "short_content": {
          "type": "string"
        },
        "expired_days": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "user_filter_type": {
          "type": "string"
        },
        "user_filter_condition": {
          "type": "string"
        },
        "userid_list": {
          "type": "string"
        }
      }
    },
    "nf_protoCreateNfWaddrsRequest": {
      "type": "object",
      "properties": {
        "nf_post_type": {
          "type": "string"
        },
        "notify_type": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "short_content": {
          "type": "string"
        },
        "expired_days": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "addrs_str": {
          "type": "string"
        }
      }
    },
    "nf_protoDescribeNfsRequest": {
      "type": "object",
      "properties": {
        "nf_post_type": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "notify_type": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "title": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "userids_str": {
          "type": "string"
        },
        "status": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "limit": {
          "type": "integer",
          "format": "int64"
        },
        "offset": {
          "type": "integer",
          "format": "int64"
        },
        "sort_key": {
          "type": "string"
        }
      }
    },
    "nf_protoDescribeNfsResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "nf_protoHelloReply": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      },
      "title": "The response message containing the greetings"
    },
    "nf_protoHelloRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      },
      "description": "The request message containing the user's name."
    }
  }
}

`
)