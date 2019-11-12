all:
	go build -o scp -mod=vendor
fmt:
	for file in `find -name "*.go" `; do gofmt -l -w $file; done
