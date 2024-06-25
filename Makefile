up:
	docker-compose up -d

down:
	docker-compose down

mock:
	mockgen -source=./repository/interface.go -destination=./repository/mock.go -package=repository
	mockgen -source=./usecase/interface.go -destination=./usecase/mock.go -package=usecase
	mockgen -source=./transaction/interface.go -destination=./transaction/mock.go -package=transaction
	mockgen -source=./query/interface.go -destination=./query/mock.go -package=query
	
gotest:
	docker-compose exec app-test go test -v ./... -count=1