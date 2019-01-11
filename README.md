
# gexp example program

```
# launch go program
go run main.go
# install protobuf
go get github.com/golang/protobuf/proto
# install protoc-gen-go
github.com/golang/protobuf/protoc-gen-go
# build the protobuf
protoc --go_out=. *.proto

# set protoc-gen-go to shell path
export PATH=$PATH:$GOPATH/bin

# test the program by running
go run main.go vehicle.pb.go

# file stats
stat -f%z vehicle.xml
stat -f%z vehicle.json
```


### References: 
https://developers.google.com/protocol-buffers/docs/gotutorial
https://developers.google.com/protocol-buffers/docs/reference/go-generated
http://blog.ralch.com/tutorial/golang-proto-buffer/

