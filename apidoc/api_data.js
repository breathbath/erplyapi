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
    "filename": "/home/apidoc/source/auth/frontMiddleware.go",
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
    "filename": "/home/apidoc/source/auth/frontMiddleware.go",
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
    "type": "get",
    "url": "/reports/:code/visits-by-day.:format?from=:from&to=:to",
    "title": "Reports visits by day",
    "name": "Visits_by_day",
    "group": "Reports",
    "description": "<p>Get visits by day</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "from",
            "description": "<p>Defines from date e.g. 2020-04-26T00:00:00, if empty now date -7 days from now is used</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "code",
            "description": "<p>Erply ID has to be provided</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "allowedValues": [
              "\"json\"",
              "\"html\""
            ],
            "optional": false,
            "field": "format",
            "defaultValue": "json",
            "description": "<p>Defines the output format</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "to",
            "description": "<p>Defines to date e.g. 2020-04-27T00:00:00, if empty now date is used</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Json",
        "content": "/reports/visits-by-day.json?from=2020-01-01T00:00&to=2020-01-30T23:59",
        "type": "String"
      },
      {
        "title": "Html",
        "content": "/reports/visits-by-day.html?from=2020-01-01T00:00&to=2020-01-30T23:59",
        "type": "String"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response {json}",
          "content": "HTTP/1.1 200 OK\n{\n\t\"data\": [\n\t\t{\n\t\t\t\"key\": \"27-04-2020\",\n\t\t\t\"value\": 6\n\t\t},\n\t\t{\n\t\t\t\"key\": \"28-04-2020\",\n\t\t\t\"value\": 3\n\t\t}\n\t]\n}",
          "type": "json"
        }
      ]
    },
    "permission": [
      {
        "name": "user"
      }
    ],
    "version": "0.0.0",
    "filename": "/home/apidoc/source/reports/handlers.go",
    "groupTitle": "Reports",
    "header": {
      "fields": {
        "Header": [
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
          "title": "Authorization Header",
          "content": "Authorization: \"Bearer eyJhbGciOi.JSUzUxMiIsIn.R5cCI6IkpXVCJ9\"",
          "type": "String"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Unauthorized(401)",
          "content": "HTTP/1.1 401 Unauthorized\n{\n    \"code\": 401,\n    \"message\": \"cookie token is empty\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/reports/:code/visits-by-hour.:format?from=:from&to=:to",
    "title": "Reports visits by hour",
    "name": "Visits_by_hour",
    "group": "Reports",
    "description": "<p>Get visits by hour</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "from",
            "description": "<p>Defines from date e.g. 2020-04-26T00:00:00, if empty now date -24 hours from now is used</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "code",
            "description": "<p>Erply ID has to be provided</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "allowedValues": [
              "\"json\"",
              "\"html\""
            ],
            "optional": false,
            "field": "format",
            "defaultValue": "json",
            "description": "<p>Defines the output format</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "to",
            "description": "<p>Defines to date e.g. 2020-04-27T00:00:00, if empty now date is used</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Json",
        "content": "/reports/visits-by-hour.json?from=2020-01-01T00:00&to=2020-01-01T23:00",
        "type": "String"
      },
      {
        "title": "Html",
        "content": "/reports/visits-by-hour.html?from=2020-01-01T00:00&to=2020-01-01T23:00",
        "type": "String"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response {json}",
          "content": "HTTP/1.1 200 OK\n{\n\t\"data\": [\n\t\t{\n\t\t\t\"key\": \"27-04-2020 09:00\",\n\t\t\t\"value\": 6\n\t\t},\n\t\t{\n\t\t\t\"key\": \"28-04-2020 06:00\",\n\t\t\t\"value\": 3\n\t\t}\n\t]\n}",
          "type": "json"
        }
      ]
    },
    "permission": [
      {
        "name": "user"
      }
    ],
    "version": "0.0.0",
    "filename": "/home/apidoc/source/reports/handlers.go",
    "groupTitle": "Reports",
    "header": {
      "fields": {
        "Header": [
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
          "title": "Authorization Header",
          "content": "Authorization: \"Bearer eyJhbGciOi.JSUzUxMiIsIn.R5cCI6IkpXVCJ9\"",
          "type": "String"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Unauthorized(401)",
          "content": "HTTP/1.1 401 Unauthorized\n{\n    \"code\": 401,\n    \"message\": \"cookie token is empty\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/reports/:code/visits-by-location.:format?from=:from&to=:to",
    "title": "Reports visits by location",
    "name": "Visits_by_location",
    "group": "Reports",
    "description": "<p>Get visits by location</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "from",
            "description": "<p>Defines from date e.g. 2020-04-26T00:00:00, if empty now date -1 day from now is used</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "code",
            "description": "<p>Erply ID has to be provided</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "allowedValues": [
              "\"json\"",
              "\"html\""
            ],
            "optional": false,
            "field": "format",
            "defaultValue": "json",
            "description": "<p>Defines the output format</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "to",
            "description": "<p>Defines to date e.g. 2020-04-27T00:00:00, if empty now date is used</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Json",
        "content": "/reports/visits-by-location.json?from=2020-01-01T00:00&to=2020-01-01T23:00",
        "type": "String"
      },
      {
        "title": "Html",
        "content": "/reports/visits-by-location.html?from=2020-01-01T00:00&to=2020-01-01T23:00",
        "type": "String"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response {json}",
          "content": "HTTP/1.1 200 OK\n{\n\t\"data\": [\n\t\t{\n\t\t\t\"key\": \"Rome\",\n\t\t\t\"value\": 6\n\t\t},\n\t\t{\n\t\t\t\"key\": \"Berlin\",\n\t\t\t\"value\": 3\n\t\t}\n\t]\n}",
          "type": "json"
        }
      ]
    },
    "permission": [
      {
        "name": "user"
      }
    ],
    "version": "0.0.0",
    "filename": "/home/apidoc/source/reports/handlers.go",
    "groupTitle": "Reports",
    "header": {
      "fields": {
        "Header": [
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
          "title": "Authorization Header",
          "content": "Authorization: \"Bearer eyJhbGciOi.JSUzUxMiIsIn.R5cCI6IkpXVCJ9\"",
          "type": "String"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Unauthorized(401)",
          "content": "HTTP/1.1 401 Unauthorized\n{\n    \"code\": 401,\n    \"message\": \"cookie token is empty\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/reports/:code/visits-by-month.:format?from=:from&to=:to",
    "title": "Reports visits by month",
    "name": "Visits_by_month",
    "group": "Reports",
    "description": "<p>Get visits by month</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "from",
            "description": "<p>Defines from date e.g. 2020-04-26T00:00:00, if empty now date -1 month from now is used</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "code",
            "description": "<p>Erply ID has to be provided</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "allowedValues": [
              "\"json\"",
              "\"html\""
            ],
            "optional": false,
            "field": "format",
            "defaultValue": "json",
            "description": "<p>Defines the output format</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "to",
            "description": "<p>Defines to date e.g. 2020-04-27T00:00:00, if empty now date is used</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Json",
        "content": "/reports/visits-by-month.json?from=2020-01-01T00:00&to=2020-01-30T23:59",
        "type": "String"
      },
      {
        "title": "Html",
        "content": "/reports/visits-by-month.html?from=2020-01-01T00:00&to=2020-01-30T23:59",
        "type": "String"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response {json}",
          "content": "HTTP/1.1 200 OK\n{\n\t\"data\": [\n\t\t{\n\t\t\t\"key\": \"04-2020\",\n\t\t\t\"value\": 6\n\t\t},\n\t\t{\n\t\t\t\"key\": \"04-2020\",\n\t\t\t\"value\": 3\n\t\t}\n\t]\n}",
          "type": "json"
        }
      ]
    },
    "permission": [
      {
        "name": "user"
      }
    ],
    "version": "0.0.0",
    "filename": "/home/apidoc/source/reports/handlers.go",
    "groupTitle": "Reports",
    "header": {
      "fields": {
        "Header": [
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
          "title": "Authorization Header",
          "content": "Authorization: \"Bearer eyJhbGciOi.JSUzUxMiIsIn.R5cCI6IkpXVCJ9\"",
          "type": "String"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Unauthorized(401)",
          "content": "HTTP/1.1 401 Unauthorized\n{\n    \"code\": 401,\n    \"message\": \"cookie token is empty\"\n}",
          "type": "json"
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
