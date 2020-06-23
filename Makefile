
run-gorm:
	go run ./gorm/...

bench:
	go test -bench=. -benchtime 30s
