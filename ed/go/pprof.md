pprof
-

````bash
go tool pprof -h

# test
goapp test ./ -v
goapp test ./ -cover
goapp test ./ -race
# cpu
goapp test ./ -v -cpuprofile cpu.out
go tool pprof -pdf -output report.cpu.pdf cpu.out
# mem
goapp test ./ -v -memprofile mem.out
go tool pprof -pdf -output report.mem.pdf mem.out
#
goapp test ./ -v -mutexprofile mtx.out
go tool pprof -pdf -output report.mtx.pdf mtx.out

# web
# 1️⃣ add pprof endpoint to web server AND run next commands:
go tool pprof http://localhost:8080/debug/pprof/heap
````
