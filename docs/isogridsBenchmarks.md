~~~
>pwd
~/go/src/github.com/taironas/tinygraphs/controllers/isogrids
~~~

~~~
> go test -v -run=^$ -bench=.
PASS
BenchmarkIsogrids-8	    3000	    407587 ns/op
ok  	github.com/taironas/tinygraphs/controllers/isogrids	1.279s
~~~

~~~
> go test -v -run=^$ -bench=^BenchmarkIsogrids$ -benchtime=2s -cpuprofile=prof.cpu
PASS
BenchmarkIsogrids-8	   10000	    417154 ns/op
ok  	github.com/taironas/tinygraphs/controllers/isogrids	4.228s
~~~

~~~
> go tool pprof isogrids.test prof.cpu
Entering interactive mode (type "help" for commands)
(pprof) top
3.82s of 4.17s total (91.61%)
Dropped 70 nodes (cum <= 0.02s)
Showing top 10 nodes out of 81 (cum >= 0.34s)
      flat  flat%   sum%        cum   cum%
     3.41s 81.77% 81.77%      3.41s 81.77%  runtime.mach_semaphore_wait
     0.09s  2.16% 83.93%      0.09s  2.16%  runtime.mach_semaphore_signal
     0.07s  1.68% 85.61%      0.19s  4.56%  runtime.mallocgc
     0.05s  1.20% 86.81%      0.13s  3.12%  runtime.rawstring
     0.05s  1.20% 88.01%      0.05s  1.20%  runtime.usleep
     0.04s  0.96% 88.97%      0.17s  4.08%  fmt.(*pp).doPrintf
     0.03s  0.72% 89.69%      0.05s  1.20%  fmt.(*fmt).integer
     0.03s  0.72% 90.41%      0.03s  0.72%  runtime.memmove
     0.03s  0.72% 91.13%      0.07s  1.68%  runtime.typedmemmove
     0.02s  0.48% 91.61%      0.34s  8.15%  github.com/ajstarks/svgo.(*SVG).pp
(pprof) top --cum
3.41s of 4.17s total (81.77%)
Dropped 70 nodes (cum <= 0.02s)
Showing top 10 nodes out of 81 (cum >= 3.25s)
      flat  flat%   sum%        cum   cum%
         0     0%     0%      3.52s 84.41%  runtime.systemstack
         0     0%     0%      3.50s 83.93%  runtime.schedule
         0     0%     0%      3.46s 82.97%  runtime.stopm
         0     0%     0%      3.43s 82.25%  runtime.semasleep.func1
         0     0%     0%      3.43s 82.25%  runtime.semasleep1
     3.41s 81.77% 81.77%      3.41s 81.77%  runtime.mach_semaphore_wait
         0     0% 81.77%      3.41s 81.77%  runtime.notesleep
         0     0% 81.77%      3.41s 81.77%  runtime.semasleep
         0     0% 81.77%      3.26s 78.18%  runtime.findrunnable
         0     0% 81.77%      3.25s 77.94%  runtime.mcall
~~~

~~~
> go test -v -run=^$ -bench=^BenchmarkIsogrids$ -benchtime=2s -memprofile=prof.mem
PASS
BenchmarkIsogrids-8	   10000	    402797 ns/op
ok  	github.com/taironas/tinygraphs/controllers/isogrids	4.091s
~~~

~~~
> go tool pprof isogrids.test prof.mem
Entering interactive mode (type "help" for commands)
(pprof) top
1069.45kB of 1069.45kB total (  100%)
Dropped 72 nodes (cum <= 5.35kB)
Showing top 10 nodes out of 12 (cum >= 512.19kB)
      flat  flat%   sum%        cum   cum%
  557.26kB 52.11% 52.11%   557.26kB 52.11%  html.init
  512.19kB 47.89%   100%   512.19kB 47.89%  runtime.malg
         0     0%   100%   557.26kB 52.11%  github.com/taironas/tinygraphs/controllers/isogrids.init
         0     0%   100%   557.26kB 52.11%  github.com/taironas/tinygraphs/write.init
         0     0%   100%   557.26kB 52.11%  html/template.init
         0     0%   100%   557.26kB 52.11%  main.init
         0     0%   100%   557.26kB 52.11%  runtime.goexit
         0     0%   100%   557.26kB 52.11%  runtime.main
         0     0%   100%   512.19kB 47.89%  runtime.mcommoninit
         0     0%   100%   512.19kB 47.89%  runtime.mpreinit
(pprof) top --cum
1069.45kB of 1069.45kB total (  100%)
Dropped 72 nodes (cum <= 5.35kB)
Showing top 10 nodes out of 12 (cum >= 512.19kB)
      flat  flat%   sum%        cum   cum%
         0     0%     0%   557.26kB 52.11%  github.com/taironas/tinygraphs/controllers/isogrids.init
         0     0%     0%   557.26kB 52.11%  github.com/taironas/tinygraphs/write.init
  557.26kB 52.11% 52.11%   557.26kB 52.11%  html.init
         0     0% 52.11%   557.26kB 52.11%  html/template.init
         0     0% 52.11%   557.26kB 52.11%  main.init
         0     0% 52.11%   557.26kB 52.11%  runtime.goexit
         0     0% 52.11%   557.26kB 52.11%  runtime.main
  512.19kB 47.89%   100%   512.19kB 47.89%  runtime.malg
         0     0%   100%   512.19kB 47.89%  runtime.mcommoninit
         0     0%   100%   512.19kB 47.89%  runtime.mpreinit
~~~
