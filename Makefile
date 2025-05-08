run:
	@docker compose up 

run_test:
	go test internal/services/parcel_test.go