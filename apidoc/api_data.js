define({ "api": [
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "/home/apidoc/source/apidoc/main.js",
    "group": "/home/apidoc/source/apidoc/main.js",
    "groupTitle": "/home/apidoc/source/apidoc/main.js",
    "name": ""
  },
  {
    "type": "post",
    "url": "/login",
    "title": "Login",
    "description": "<p>User auth</p>",
    "name": "Login",
    "group": "Auth",
    "parameter": {
      "examples": [
        {
          "title": "Body:",
          "content": "{\n\t\"username\": \"admin\",\n\t\"password\": \"admin\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response",
          "content": "HTTP/1.1 200 OK\n{\n    \"code\": 200,\n    \"expire\": \"2020-04-20T13:30:49+03:00\",\n    \"token\": \"YOUR_TOKEN_RESULT\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Bad request(401)",
          "content": "HTTP/1.1 401 Unauthorized\n{\n    \"code\": 401,\n    \"message\": \"incorrect Username or Password\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "/home/apidoc/source/auth/authMiddleware.go",
    "groupTitle": "Auth",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Content-Type",
            "defaultValue": "application/json",
            "description": "<p>Json content type</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Content-Type",
          "content": "Content-Type:\"application/json\"",
          "type": "String"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/refresh",
    "title": "Refresh token",
    "description": "<p>Refreshes the auth token</p>",
    "name": "Refresh_token",
    "group": "Auth",
    "success": {
      "examples": [
        {
          "title": "Success-Response",
          "content": "HTTP/1.1 200 OK\n{\n    \"code\": 200,\n    \"expire\": \"2020-04-20T13:47:59+03:00\",\n    \"token\": \"YOUR_TOKEN_RESULT\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Bad request(401)",
          "content": "HTTP/1.1 401 Unauthorized\n{\n    \"code\": 401,\n    \"message\": \"cookie token is empty\"\n}",
          "type": "json"
        },
        {
          "title": "Unauthorized(401)",
          "content": "HTTP/1.1 401 Unauthorized\n{\n    \"code\": 401,\n    \"message\": \"cookie token is empty\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "/home/apidoc/source/auth/authMiddleware.go",
    "groupTitle": "Auth",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Content-Type",
            "defaultValue": "application/json",
            "description": "<p>Json content type</p>"
          },
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token value</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Content-Type",
          "content": "Content-Type:\"application/json\"",
          "type": "String"
        },
        {
          "title": "Authorization Header",
          "content": "Authorization: \"Bearer eyJhbGciOi.JSUzUxMiIsIn.R5cCI6IkpXVCJ9\"",
          "type": "String"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/visits",
    "title": "Visit metric create",
    "description": "<p>Adds a new visit metric</p>",
    "name": "Visit_create",
    "group": "Visit",
    "parameter": {
      "examples": [
        {
          "title": "Body:",
          "content": "{\n\t\"location\": \t\"Rome\", #required\n\t\"device_hash\": \t\"djfasdlfjlfjasdlkfas\", #required\n\t\"erply_id\":\t\t\"100234\" #required\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response",
          "content": "HTTP/1.1 200 OK",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Bad request(400)",
          "content": "HTTP/1.1 400 Bad request\n{\n    \"error\": \"Key: 'VisitMetric.Location' Error:Field validation for 'Location' failed on the 'required' tag\"\n}",
          "type": "json"
        },
        {
          "title": "Unauthorized(401)",
          "content": "HTTP/1.1 401 Unauthorized\n{\n    \"code\": 401,\n    \"message\": \"cookie token is empty\"\n}",
          "type": "json"
        }
      ]
    },
    "permission": [
      {
        "name": "registered user"
      }
    ],
    "version": "0.0.0",
    "filename": "/home/apidoc/source/metrics/visits.go",
    "groupTitle": "Visit",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Content-Type",
            "defaultValue": "application/json",
            "description": "<p>Json content type</p>"
          },
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token value</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Content-Type",
          "content": "Content-Type:\"application/json\"",
          "type": "String"
        },
        {
          "title": "Authorization Header",
          "content": "Authorization: \"Bearer eyJhbGciOi.JSUzUxMiIsIn.R5cCI6IkpXVCJ9\"",
          "type": "String"
        }
      ]
    }
  }
] });
