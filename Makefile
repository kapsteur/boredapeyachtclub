run: create bin/main
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/main: main.go
	go build -o bin/main main.go

clean:
	rm -rf bin