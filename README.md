# REST API server with two endpoints

# Run

make run

# First end point
POST /in-memory	&ensp;Stores key-value pair </br>
```
curl -X POST -d '{"key":"foo", "value":"bar"}' http://localhost:3000/in-memory
```
GET /in-memory?key= &ensp;Retrieves the value by key
```
curl http://localhost:3000/in-memory?key=foo
```
# Second end point
GET /mongo Retrieves records from the database based on a query
```
curl -X GET -d '{"startDate":"2016-01-01", "endDate":"2022-01-31", "minCount":10, "maxCount":10000}' http://localhost:3000/mongo
```
