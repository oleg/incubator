CONFIG_PATH=${HOME}/.proglog/

.PHONY: init
init:
	mkdir -p ${CONFIG_PATH}

.PHONY: gen-cert
gen-cert:
	cfssl gencert \
			-initca test/ca-csr.json | cfssljson -bare ca

	cfssl gencert \
			-ca=ca.pem \
			-ca-key=ca-key.pem \
			-config=test/ca-config.json \
			-profile=server \
			test/server-csr.json | cfssljson -bare server

	cfssl gencert \
			-ca=ca.pem \
			-ca-key=ca-key.pem \
			-config=test/ca-config.json \
			-profile=client \
			test/client-csr.json | cfssljson -bare client

	mv *.pem *.csr ${CONFIG_PATH}

.PHONY: compile
compile:
	protoc api/v1/*.proto \
                --go_out=. \
                --go_opt=paths=source_relative \
                --go-grpc_out=. \
                --go-grpc_opt=paths=source_relative \
                --proto_path=.

.PHONY: test
test:
	go test ./...

.PHONY: race-test
race-test:
	go test -race ./...

