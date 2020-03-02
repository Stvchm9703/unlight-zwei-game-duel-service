generate_vcred:
	protoc \
		-I proto/ \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/ \
		-I vendor/github.com/gogo/googleapis/ \
		-I vendor/ \
		--go_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		--govalidators_out=gogoimport=true,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		proto/*.proto
	# gvm issue :  move the genrated file to current directory
	mv $(CURDIR)/vendor/*.pb.go $(CURDIR)/proto/
	mv $(CURDIR)/vendor/*.validator.pb.go $(CURDIR)/proto/
	# mv $(CURDIR)/vendor/cred.pb.gw.go $(CURDIR)/proto/
	## Workaround for https://github.com/grpc-ecosystem/grpc-gateway/issues/229.
	# sed -i.bak "s/empty.Empty/types.Empty/g" proto/cred.pb.gw.go && rm proto/cred.pb.gw.go.bak

	# ## Generate static assets for OpenAPI UI
	# statik -m -f -src third_party/OpenAPI/

build:
	go build -o account_server main.go

docker_build:
	go build -o account_server main.go
	docker-compose build 
