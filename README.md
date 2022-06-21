# loggator
Fast, extensible and robust log aggregator platform written in golang.

# Features
1. Send logs to the endpoint `POST /logs/`

# Instructions
1. To migrate the database:
	i) [Install `psql`](https://www.calhoun.io/how-to-install-postgresql-9-6-on-mac-os-x/)  
	ii) Access the database `psql -U postgres`.  
	iii) Create the `loggator` database.  
	iv) `CREATE TABLE loggator;`  
	v) Apply the schema `make localmigrateup`.  

# API documentation
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/6708aa075ed682aac4d3?action=collection%2Fimport)
## Live view
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/1921454-9ce96aa1-be53-4128-a9d0-d90b4645d4be?action=collection%2Ffork&collection-url=entityId%3D1921454-9ce96aa1-be53-4128-a9d0-d90b4645d4be%26entityType%3Dcollection%26workspaceId%3D50bf739e-2b62-474b-b442-3ed1b8ed7fa9)
