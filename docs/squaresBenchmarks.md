Squares Benchmark
=================

~~~
> pwd
~/go/src/github.com/taironas/tinygraphs/controllers/squares
> go test -v -run=^$ -bench=.
PASS
BenchmarkSquares-8	     200	   6899024 ns/op
ok  	github.com/taironas/tinygraphs/controllers/squares	2.097s
~~~

~~~
>➜go test -v -run=^$ -bench=^BenchmarkSquares$ -benchtime=2s -cpuprofile=prof.cpu
PASS
BenchmarkSquares-8	     500	   6758214 ns/op
ok  	github.com/taironas/tinygraphs/controllers/squares	4.096s
~~~

~~~
> go tool pprof squares.test prof.cpu
Entering interactive mode (type "help" for commands)
(pprof) top
3510ms of 3930ms total (89.31%)
Dropped 42 nodes (cum <= 19.65ms)
Showing top 10 nodes out of 76 (cum >= 70ms)
      flat  flat%   sum%        cum   cum%
    2870ms 73.03% 73.03%     2870ms 73.03%  runtime.mach_semaphore_wait
     110ms  2.80% 75.83%      110ms  2.80%  runtime.mach_semaphore_signal
     100ms  2.54% 78.37%      100ms  2.54%  image/jpeg.fdct
      80ms  2.04% 80.41%       80ms  2.04%  runtime.memmove
      70ms  1.78% 82.19%       70ms  1.78%  runtime.mach_semaphore_timedwait
      60ms  1.53% 83.72%      530ms 13.49%  github.com/taironas/tinygraphs/draw/squares.Image
      60ms  1.53% 85.24%      160ms  4.07%  image.(*RGBA).Set
      60ms  1.53% 86.77%       90ms  2.29%  runtime.scanobject
      50ms  1.27% 88.04%      150ms  3.82%  image/jpeg.(*encoder).writeBlock
      50ms  1.27% 89.31%       70ms  1.78%  image/jpeg.rgbaToYCbCr
(pprof) top --cum
2.88s of 3.93s total (73.28%)
Dropped 42 nodes (cum <= 0.02s)
Showing top 10 nodes out of 76 (cum >= 2.63s)
      flat  flat%   sum%        cum   cum%
     0.01s  0.25%  0.25%      3.04s 77.35%  runtime.systemstack
         0     0%  0.25%      2.94s 74.81%  runtime.semasleep.func1
         0     0%  0.25%      2.94s 74.81%  runtime.semasleep1
         0     0%  0.25%      2.90s 73.79%  runtime.schedule
         0     0%  0.25%      2.88s 73.28%  runtime.stopm
     2.87s 73.03% 73.28%      2.87s 73.03%  runtime.mach_semaphore_wait
         0     0% 73.28%      2.87s 73.03%  runtime.notesleep
         0     0% 73.28%      2.87s 73.03%  runtime.semasleep
         0     0% 73.28%      2.67s 67.94%  runtime.findrunnable
         0     0% 73.28%      2.63s 66.92%  runtime.mcall
(pprof) 
~~~

~~~
> go test -v -run=^$ -bench=^BenchmarkSquares$ -benchtime=2s -memprofile=prof.mem
PASS
BenchmarkSquares-8	     500	   6829991 ns/op
ok  	github.com/taironas/tinygraphs/controllers/squares	4.135s
~~~

~~~
go tool pprof --alloc_space squares.test prof.mem
Entering interactive mode (type "help" for commands)
(pprof) top
723.66MB of 729.17MB total (99.24%)
Dropped 25 nodes (cum <= 3.65MB)
      flat  flat%   sum%        cum   cum%
  565.51MB 77.56% 77.56%   565.51MB 77.56%  github.com/taironas/tinygraphs/draw/squares.Image
  158.15MB 21.69% 99.24%   727.17MB 99.73%  github.com/taironas/tinygraphs/controllers/squares.Square
         0     0% 99.24%   728.67MB 99.93%  github.com/taironas/route.(*Router).ServeHTTP
         0     0% 99.24%   727.17MB 99.73%  github.com/taironas/route.clearHandler.func1
         0     0% 99.24%   728.67MB 99.93%  github.com/taironas/tinygraphs/controllers/squares.BenchmarkSquares
         0     0% 99.24%   727.17MB 99.73%  net/http.HandlerFunc.ServeHTTP
         0     0% 99.24%   728.67MB 99.93%  runtime.goexit
         0     0% 99.24%   728.67MB 99.93%  testing.(*B).launch
         0     0% 99.24%   728.67MB 99.93%  testing.(*B).runN
(pprof) top --cum
723.66MB of 729.17MB total (99.24%)
Dropped 25 nodes (cum <= 3.65MB)
      flat  flat%   sum%        cum   cum%
         0     0%     0%   728.67MB 99.93%  github.com/taironas/route.(*Router).ServeHTTP
         0     0%     0%   728.67MB 99.93%  github.com/taironas/tinygraphs/controllers/squares.BenchmarkSquares
         0     0%     0%   728.67MB 99.93%  runtime.goexit
         0     0%     0%   728.67MB 99.93%  testing.(*B).launch
         0     0%     0%   728.67MB 99.93%  testing.(*B).runN
         0     0%     0%   727.17MB 99.73%  github.com/taironas/route.clearHandler.func1
  158.15MB 21.69% 21.69%   727.17MB 99.73%  github.com/taironas/tinygraphs/controllers/squares.Square
         0     0% 21.69%   727.17MB 99.73%  net/http.HandlerFunc.ServeHTTP
  565.51MB 77.56% 99.24%   565.51MB 77.56%  github.com/taironas/tinygraphs/draw/squares.Image
(pprof) 
~~~
