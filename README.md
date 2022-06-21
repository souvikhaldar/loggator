# loggator
A simple yet fast, extensible and robust log aggregator platform written in golang.

# Features
1. Send logs in JSON to the endpoint and the logs will be carefully stored in database.
2. Search for logs using various filters.
3. Generate charts or graphs or visualize functional performance of the service.
4. Raise alert for anomalous behaviours like too many ERROR logs, too many INFO logs for a particular endpoint signaling possible DDOS, etc.

# Dependency
1. https://github.com/golang-migrate/migrate

# Workflow
1. Send logs to the endpoint 
	The log agent will parse the logs and send to the above endpoint in the following JSON format

	`POST /logs/`
	```
	{
	    "date": "2012-04-23",
	    "time": "23:11:12",
	    "log": "this is a error log",
	    "log_level": "ERROR",
	    "tenant_id": 2,
	    "service_name": "loggator",
	    "package_name": "logs",
	    "file_name": "repository.go"
	}
	```
2. Fetch logs based on filters
	## Request
	GET `/logs?log_level=ERROR&service_name=loggator&package_name=logs&date=2012-04-23&file_name=repository.go&tenant_id=1`
	## Response
	```
	[
	   {
	      "tenant_id":1,
	      "log_id":4,
	      "log":"this is a error log",
	      "created_at":"2022-06-21T12:23:29.985752Z",
	      "date":"2012-04-23T00:00:00Z",
	      "time":"0000-01-01T23:11:12Z",
	      "log_level":"ERROR",
	      "service_name":"loggator",
	      "file_name":"repository.go",
	      "package_name":"logs"
	   },
	   {
	      "tenant_id":1,
	      "log_id":5,
	      "log":"this is a error log",
	      "created_at":"2022-06-21T12:29:22.952193Z",
	      "date":"2012-04-23T00:00:00Z",
	      "time":"0000-01-01T23:11:12Z",
	      "log_level":"ERROR",
	      "service_name":"loggator",
	      "file_name":"repository.go",
	      "package_name":"logs"
	   },
	   {
	      "tenant_id":1,
	      "log_id":6,
	      "log":"this is a error log",
	      "created_at":"2022-06-21T12:31:46.73274Z",
	      "date":"2012-04-23T00:00:00Z",
	      "time":"0000-01-01T23:11:12Z",
	      "log_level":"ERROR",
	      "service_name":"loggator",
	      "file_name":"repository.go",
	      "package_name":"logs"
	   }
	]
	```


# Instructions
1. To migrate the database:  
	i) [Install `psql`](https://www.calhoun.io/how-to-install-postgresql-9-6-on-mac-os-x/)   
	ii) Access the database `psql -U postgres`.   
	iii) Create the `loggator` database `CREATE TABLE loggator;`   
	iv) Apply the schema `make localmigrateup`.   

2. To run the service
	i) Run `go run cmd/server/main.go`
	ii) Run the psql database on port 5432.  

# API documentation
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/1921454-9ce96aa1-be53-4128-a9d0-d90b4645d4be?action=collection%2Ffork&collection-url=entityId%3D1921454-9ce96aa1-be53-4128-a9d0-d90b4645d4be%26entityType%3Dcollection%26workspaceId%3D50bf739e-2b62-474b-b442-3ed1b8ed7fa9)

