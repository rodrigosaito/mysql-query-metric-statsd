build: clean
	CGO_ENABLED=0 go build -a -installsuffix cgo

clean:
	go clean

docker: build
	docker build -t rodrigosaito/mysql-query-metric-statsd .
