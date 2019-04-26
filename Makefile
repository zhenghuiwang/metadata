hello:
	curl localhost:8080/api/v1/resource/my-resource | jq .
add-type:
	curl -H "Content-Type:application/raw-json" localhost:8080/api/v1/type/namespaces/my-company.com/kinds/artifact/versions/v1 -d @customized_schema.json | jq .
