
TIME=1s

bench:
	go test -bench=. -benchtime $(TIME)
