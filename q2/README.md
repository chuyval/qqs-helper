# Why does proto marshal then unmarshal fail when dependency repository has a vendor directory?

See [project](https://github.com/chuyval/qqs/tree/master/q2) this repository is related to.

Contains `main.go` file that when run, it successfully Unmarshals a protobuf struct
before and after installing this repos dependencies.

Same `main.go` file reports an error when it has been run in the other
repository when this repository contains a vendor directory.

[Stackoverflow](https://stackoverflow.com/questions/52241759/why-does-proto-marshal-then-unmarshal-fail-when-dependency-repository-has-a-vend)


## Build proto files
```
protoc -I. --go_out=$GOPATH/src ./*.proto
```

## Installing dependencies
```
glide install
```