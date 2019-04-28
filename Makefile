hello:
	curl localhost:8080/api/v1/resource/my-resource
add-type:
	curl -H "Content-Type:application/raw-json" localhost:8080/api/v1/type/namespaces/my-company.com/kinds/artifact/versions/v1 -d @customized_schema.json | jq .

build:
	bazel build --action_env=PATH --define=grpc_no_ares=true //...

update-build-target:
	bazel run //:gazelle

run:
	bazel run --action_env=PATH --define=grpc_no_ares=true //server -- --logtostderr
