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
[Link to collection](https://www.getpostman.com/collections/6708aa075ed682aac4d3)
