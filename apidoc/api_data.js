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
    "url": "/customers",
    "title": "Customers create",
    "description": "<p>Creates a new customer</p>",
    "name": "Customers_create",
    "group": "Customers",
    "parameter": {
      "examples": [
        {
          "title": "Body:",
          "content": "{\n\t\"CompanyName\": \"My Personal Inc\",\n\t\"Address\":            \"Elm str\",\n\t\"PostalCode\":         \"100234\",\n\t\"Country\":            \"USA\",\n\t\"FullName\":           \"Big Boss\",\n\t\"RegistryCode\":      \"1234\",\n\t\"VatNumber\":          \"23456\",\n\t\"Email\":              \"no@mail.me\",\n\t\"Phone\":              \"+13434134233134\",\n\t\"BankName\":           \"Best Bank\",\n\t\"BankAccountNumber\":  \"3434937493749813\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response",
          "content": "HTTP/1.1 200 OK\n{\n    \"clientID\": 11,\n    \"customerID\": 11\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Bad request(400)",
          "content": "HTTP/1.1 400 Bad request\n{\n    \"message\": \"ERPLY API: Can not save customer with empty name or registry number status: Error\"\n}",
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
    "filename": "/home/apidoc/source/erply/customers.go",
    "groupTitle": "Customers"
  },
  {
    "type": "get",
    "url": "/customers/:ids",
    "title": "Customers list by ids",
    "description": "<p>Lists customers by comma separated ids list</p>",
    "name": "Customers_list_by_ids",
    "group": "Customers",
    "examples": [
      {
        "title": "With many ids",
        "content": "/customers/1,2,3",
        "type": "String"
      },
      {
        "title": "With one id",
        "content": "/customers/1",
        "type": "String"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response",
          "content": "HTTP/1.1 200 OK\n[\n    {\n        \"id\": 6,\n        \"customerID\": 6,\n        \"type_id\": \"\",\n        \"fullName\": \"Comp INC\",\n        \"companyName\": \"Comp INC\",\n        \"firstName\": \"\",\n        \"lastName\": \"\",\n        \"groupID\": 14,\n        \"EDI\": \"\",\n        \"phone\": \"\",\n        \"eInvoiceEmail\": \"\",\n        \"email\": \"\",\n        \"fax\": \"\",\n        \"code\": \"3333\",\n        \"referenceNumber\": \"\",\n        \"vatNumber\": \"\",\n        \"bankName\": \"\",\n        \"bankAccountNumber\": \"\",\n        \"bankIBAN\": \"\",\n        \"bankSWIFT\": \"\",\n        \"paymentDays\": 0,\n        \"notes\": \"\",\n        \"lastModified\": 1587311074,\n        \"customerType\": \"COMPANY\",\n        \"address\": \"\",\n        \"addresses\": null,\n        \"street\": \"\",\n        \"address2\": \"\",\n        \"city\": \"\",\n        \"postalCode\": \"\",\n        \"country\": \"\",\n        \"state\": \"\",\n        \"contactPersons\": []\n    },\n    {\n        \"id\": 3,\n        \"customerID\": 3,\n        \"type_id\": \"\",\n        \"fullName\": \"mustermann, max\",\n        \"companyName\": \"\",\n        \"firstName\": \"max\",\n        \"lastName\": \"mustermann\",\n        \"groupID\": 14,\n        \"EDI\": \"\",\n        \"phone\": \"\",\n        \"eInvoiceEmail\": \"\",\n        \"email\": \"\",\n        \"fax\": \"\",\n        \"code\": \"\",\n        \"referenceNumber\": \"\",\n        \"vatNumber\": \"\",\n        \"bankName\": \"\",\n        \"bankAccountNumber\": \"\",\n        \"bankIBAN\": \"\",\n        \"bankSWIFT\": \"\",\n        \"paymentDays\": 0,\n        \"notes\": \"\",\n        \"lastModified\": 1587298463,\n        \"customerType\": \"PERSON\",\n        \"address\": \"\",\n        \"addresses\": null,\n        \"street\": \"\",\n        \"address2\": \"\",\n        \"city\": \"\",\n        \"postalCode\": \"\",\n        \"country\": \"\",\n        \"state\": \"\",\n        \"contactPersons\": []\n    }\n]",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Not found(404)",
          "content": "HTTP/1.1 404 Not found\n[]",
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
    "filename": "/home/apidoc/source/erply/customers.go",
    "groupTitle": "Customers"
  }
] });