# app版本比较

## 使用及单测
****
```go
ret := CompareVersion("8.1", "8.1")
assert.Equal(t, 0, ret)
ret = CompareVersion("8.1", "8.01")
assert.Equal(t, 0, ret)
ret = CompareVersion("8.21", "8.01")
assert.Equal(t, 1, ret)
ret = CompareVersion("8.101", "8.201")
assert.Equal(t, -1, ret)
ret = CompareVersion("8.1.2", "8.01")
assert.Equal(t, 1, ret)
ret = CompareVersion("7.0.09.000", "7.0.09")
assert.Equal(t, 0, ret)
ret = CompareVersion("7.0.08.9999", "7.0.09.9999")
assert.Equal(t, -1, ret)
ret = CompareVersion("7.0.9.9999", "7.0.09.999")
assert.Equal(t, 1, ret)
ret = CompareVersion("9.01", "9.0")
assert.Equal(t, 1, ret)
ret = CompareVersion("7.0.9.9999", "6.6.0.0000")
assert.Equal(t, 1, ret)
ret = CompareVersion("", "9")
assert.Equal(t, -1, ret)
```
**CompareVersionWithCache**
```go
ret := CompareVersionWithCache("8.1", "8.1")
assert.Equal(t, 0, ret)
ret = CompareVersionWithCache("8.1", "8.01")
assert.Equal(t, 0, ret)
ret = CompareVersionWithCache("8.21", "8.01")
assert.Equal(t, 1, ret)
ret = CompareVersionWithCache("8.101", "8.201")
assert.Equal(t, -1, ret)
ret = CompareVersionWithCache("8.1.2", "8.01")
assert.Equal(t, 1, ret)
ret = CompareVersionWithCache("7.0.09.000", "7.0.09")
assert.Equal(t, 0, ret)
ret = CompareVersionWithCache("7.0.08.9999", "7.0.09.9999")
assert.Equal(t, -1, ret)
ret = CompareVersionWithCache("7.0.9.9999", "7.0.09.999")
assert.Equal(t, 1, ret)
ret = CompareVersionWithCache("9.01", "9.0")
assert.Equal(t, 1, ret)
ret = CompareVersionWithCache("7.0.9.9999", "6.6.0.0000")
assert.Equal(t, 1, ret)
ret = CompareVersionWithCache("", "9")
assert.Equal(t, -1, ret)
```

## Benchmark
```
pkg: github.com/Isites/ares/strs
cpu: Intel(R) Core(TM) i7-7567U CPU @ 3.50GHz
BenchmarkCompareVersion-4            	 6377474	       193.9 ns/op	       8 B/op	       2 allocs/op
BenchmarkCompareVersionWithCache-4   	 8369961	       151.2 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/Isites/ares/strs	3.009s
```