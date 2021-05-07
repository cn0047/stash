pprof
-

````bash
go tool pprof -h

# memprofile
go build -gcflags='-memprofile mem.out' -o=app main.go
go tool pprof -pdf -output report.mem.pdf ./app mem.out



# test
goapp test ./ -v
goapp test ./ -cover
goapp test ./ -race
# cpu
goapp test ./ -v -cpuprofile cpu.out
go tool pprof -pdf -output report.cpu.pdf cpu.out
# mem
goapp test ./ -v -memprofile mem.out
goapp test ./ -v -memprofile mem.out -test.memprofilerate=1
go tool pprof -pdf -output report.mem.pdf mem.out
#
goapp test ./ -v -mutexprofile mtx.out
go tool pprof -pdf -output report.mtx.pdf mtx.out

# web
# 1️⃣ add pprof endpoint to web server AND run next commands:
go tool pprof http://localhost:8080/debug/pprof/heap
````
