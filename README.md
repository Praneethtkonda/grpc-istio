# go_signc
go_signc is an updated / newly written version of signc where we eliminate a couple of cons from the previous version.
The protocol is updated to grpc now which gives us support for scalability, security and lot more features.

### Start the grpc python server
```bash
$ pip install -r requirement.txt

$ python -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. *.proto
```

### Run the go grpc client to send the signing request (developer)
#### Setting up the protobuf compiler environment
```bash
$ export PATH=$PATH:$(go env GOPATH)/bin
$ export GOPATH=$HOME/go
$ go mod download
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ protoc --proto_path=. --go_out=../pb --go-grpc_out=../pb --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto
```
#### Running the go client
```bash
$ go run client.go
```

### Building the go client binary
```bash
$ go build -o signc ./
$ ./client --help
```

### K8s setup

This POC mainly uses Istio as the service mesh concept for L7 load balancing.

#### Istio installation
```bash
$ minikube start --cpus 6 --memory 8192
# Download Latest version of Istio
$ curl -L https://istio.io/downloadIstio | sh -
$ cd istio-1.17.1
$ ./istioctl install --set profile=demo -y
$ minikube kubectl -- get pods -n istio-system
# For side car proxy injection
$ kubectl label namespace default istio-injection=enabled
$ kubectl get ns default --show-labels
$ kubectl get pods -n istio-system
```