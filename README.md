# protoc-gen-microweb

TBD

## Docs

```
GO111MODULE=off go get -d github.com/grpc-ecosystem/grpc-gateway/...
GO111MODULE=off go get -v github.com/golang/protobuf/protoc-gen-go
GO111MODULE=off go get -v github.com/micro/protoc-gen-micro
GO111MODULE=off go get -v github.com/owncloud/protoc-gen-microweb

protoc \
	-I=/usr/local/include/ \
	-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I=examples/greeter/proto/ \
	--go_out=examples/greeter/proto/ \
	--micro_out=examples/greeter/proto/ \
	--microweb_out=examples/greeter/proto/ \
	greeter.proto
```

## Install

```
GO111MODULE=off go get -v github.com/owncloud/protoc-gen-microweb
```

## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.12.

```bash
go get -d github.com/owncloud/protoc-gen-microweb
cd $GOPATH/src/github.com/owncloud/protoc-gen-microweb

go install
```

## Security

If you find a security issue please contact security@owncloud.com first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

* [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```
Copyright (c) 2021 ownCloud GmbH <https://owncloud.com>
```
