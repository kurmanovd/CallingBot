build:
	go build -o bin/voip-backend -mod vendor cmd/api/main.go

run:
	cd cmd/api; ./rundev.sh

clean:
	rm -rf bin/voip-backend