
# goexp example program

 
Download and install Golang https://golang.org/dl/
```console

# install protobuf
go get github.com/golang/protobuf/proto

# install protoc-gen-go
github.com/golang/protobuf/protoc-gen-go

# set protoc-gen-go to shell path
export PATH=$PATH:$GOPATH/bin

# build the protobuf
protoc --go_out=. *.proto

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

