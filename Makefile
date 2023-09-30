test:
	go clean -testcache && \
	go test -cover ./...

run:
	docker-compose rm --stop && docker-compose build && docker-compose up --force-recreate