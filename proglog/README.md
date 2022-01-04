# Proglog 
Simple API that can store and return data from record

# Test 

```
>go run main.go 

>curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzAA"}}'
<{"offset":0}
>curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzEK"}}'
<{"offset":1}
>curl -X POST localhost:8080 -d '{"record": {"value": "TGV0J3MgR28gIzCC"}}'
<{"offset":2}


curl -X GET localhost:8080 -d '{"offset":1}'
{"record":{"value":"TGV0J3MgR28gIzEK","offset":1}}

```