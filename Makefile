ROOT_DIR=$(shell pwd)

generate_swagger:
	docker run -it -v ${ROOT_DIR}:${ROOT_DIR} -e SWAGGER_GENERATE_EXTENSION=false --workdir ${ROOT_DIR} quay.io/goswagger/swagger generate spec -o ./docs/swagger.json -m;

generate_openapi:
	curl -X 'POST' \
      'https://converter.swagger.io/api/convert' \
	  -H 'accept: application/yaml' \
	  -H 'Content-Type: application/json' \
      -d '@./docs/swagger.json' > docs/openapi.yaml

generate_both: generate_swagger generate_openapi

docker_build:
	docker build -t referentielijsten:new .

docker_run: docker_delete
	docker run --name referentielijsten -p 8001:8000 -d referentielijsten:new

docker_delete:
	docker rm -f referentielijsten
