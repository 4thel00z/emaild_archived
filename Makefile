build-emaild:
	go build -o build/emaild main/emaild.go

run: build-emaild
	build/emaild --accountsPath config/accounts.json

clean:
	rm -rf build/*

	
