
# Benchmark GORM vs GORP

This benchmark include currently 3 operations:

- Insert
- Update
- Find

## Results

> The outputs below were reordered in order
> to make it easier to read.

```
go test -bench=. -benchtime 30s
goos: linux
goarch: amd64
BenchmarkGorm/Insert-4         	   21552	   1688333 ns/op
BenchmarkGorp/Insert-4         	   25905	   1473178 ns/op
BenchmarkGorm/Update-4         	   19516	   1739969 ns/op
BenchmarkGorp/Update-4         	   22705	   1423827 ns/op
BenchmarkGorm/Read-4           	   87717	    446263 ns/op
BenchmarkGorp/Read-4           	   85447	    485613 ns/op
```

```
go test -bench=. -benchtime 5s
goos: linux
goarch: amd64
pkg: github.com/vingarcia/go-gorm-vs-gorp
BenchmarkGorm/Insert-4         	    3457	   1652128 ns/op
BenchmarkGorp/Insert-4         	    4072	   1527532 ns/op
BenchmarkGorm/Update-4         	    2547	   2324664 ns/op
BenchmarkGorp/Update-4         	    3392	   1866467 ns/op
BenchmarkGorm/Read-4           	   10000	    523259 ns/op
BenchmarkGorp/Read-4           	   13252	    482991 ns/op
```

```
go test -bench=. -benchtime 1s
goos: linux
goarch: amd64
pkg: github.com/vingarcia/go-gorm-vs-gorp
BenchmarkGorm/Insert-4         	     565	   2497875 ns/op
BenchmarkGorp/Insert-4         	     571	   2141689 ns/op
BenchmarkGorm/Update-4         	     531	   2465267 ns/op
BenchmarkGorp/Update-4         	     516	   2144158 ns/op
BenchmarkGorm/Read-4           	    3818	    324860 ns/op
BenchmarkGorp/Read-4           	    3818	    317082 ns/op
```
