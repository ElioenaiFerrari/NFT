.PHONY: dev


dev: 
	npx nodemon -e go --signal SIGTERM --exec 'go' run .