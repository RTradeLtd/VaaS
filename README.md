# VaaS

VaaS (Vanity As A Service) is an experimental SaaS product, offering secure Vanity Address generation leveraging RTrade's extensive computing power.

To build the protobuf messages run `protoc --go_out=. *.proto`

Requires ETCDv3 running to coordinate workers

Everything unless specified otherwise (encrypt.go) is licensed under Apache-2

## Running

1) Start etcd instance
2) Start actor `./VaaS distributor 127.0.0.1:0`
3) Start worker `./VaaS distributor 127.0.0.1:0`
4) Make get API call too `http://{{api_url}}:6767/api/v1/ethereum/generate/distributed`