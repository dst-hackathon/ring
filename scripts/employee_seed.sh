#!/bin/bash
curl -H "Content-Type: application/vnd.api+json" -d '{ "data": { "type": "employees", "id": "dt00x", "attributes": { "first-name": "Old", "last-name": "Republic" } } }' http://localhost:3000/employees
curl -H "Content-Type: application/vnd.api+json" -d '{ "data": { "type": "employees", "id": "dt77850", "attributes": { "first-name": "Sid", "last-name": "L" } } }' http://localhost:3000/employees
