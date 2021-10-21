# testing package 使用

[toc]




## testing 使用格式

测试用例有四种形式：
- `TestXxxx(t *testing.T)` : // 基本测试用例
- `BenchmarkXxxx(b *testing.B)` : // 压力测试的测试用例
- `Example_Xxx()` :  // 测试控制台输出的例子
- `TestMain(m *testing.M)` : // 测试 Main 函数

注意：
- Xxx 可以是任何字母数字字符串，但是第一个字母不能是小写字母。
- 在这些函数中，使用 Error、Fail 或相关方法来发出失败信号。
- 要编写一个新的测试套件，需要创建一个名称以 _test.go 结尾的文件，该文件包含 TestXxx 函数

    



## testing 常用变量

使用 go test 命令可以使用的常用的选项：

- `-test.short` : 一个快速测试的标记，在测试用例中可以使用 testing.Short() 来绕开一些测试
- `-test.outputdir` : 输出目录
- `-test.coverprofile` : 测试覆盖率参数，指定输出文件
- `-test.run` : 指定正则来运行某个 / 某些测试用例
- `-test.memprofile` : 内存分析参数，指定输出文件
- `-test.memprofilerate` : 内存分析参数，内存分析的抽样率
- `-test.cpuprofile` : cpu 分析输出参数，为空则不做 cpu 分析
- `-test.blockprofile` : 阻塞事件的分析参数，指定输出文件
- `-test.blockprofilerate` : 阻塞事件的分析参数，指定抽样频率
- `-test.timeout` : 超时时间
- `-test.cpu` : 指定 cpu 数量
- `-test.parallel` : 指定运行测试用例的并行数



## testing 常用的结构体

testing 模块中的常用的结构体：

- `B` : 压力测试
- `BenchmarkResult` : 压力测试结果
- `Cover` : 代码覆盖率相关结构体
- `CoverBlock` : 代码覆盖率相关结构体
- `InternalBenchmark` : 内部使用的结构体
- `InternalExample` : 内部使用的结构体
- `InternalTest` : 内部使用的结构体
- `M` : main 测试使用的结构体
- `PB` : Parallel benchmarks 并行测试使用的结构体
- `T` : 普通测试用例
- `TB` : 测试用例的接口


## testing 的通用方法


testing.T 结构内部是继承自 common 结构，common 结构提供集中方法，是我们经常会用到的：

1. 当我们遇到一个断言错误的时候，我们就会判断这个测试用例失败，就会使用到：
    - `Fail()` : case 失败，测试用例继续
    - `FailedNow()` : case 失败，测试用例中断
2. 当我们遇到一个断言错误，只希望跳过这个错误，但是不希望标示测试用例失败，会使用到：
    - `SkipNow()` : case 跳过，测试用例不继续
3. 当我们只希望在一个地方打印出信息，我们会用到 :
    - `Log()` : 输出信息
    - `Logf()` : 输出有 format 的信息
4. 当我们希望跳过这个用例，并且打印出信息 :
    - `Skip()` : Log + SkipNow
    - `Skipf()` : Logf + SkipNow
5. 当我们希望断言失败的时候，测试用例失败，打印出必要的信息，但是测试用例继续：
    - `Error()` : Log + Fail
    - `Errorf()` : Logf + Fail
6. 当我们希望断言失败的时候，测试用例失败，打印出必要的信息，测试用例中断：
    - `Fatal()` : Log + FailNow
    - `Fatalf()` : Logf + FailNow



## testing 单元测试


### 单元测试简单例子

1. 需要测试的函数代码为：
```
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
```

2. 编写单元测试代码：
```
func TestFib(t *testing.T) {
	var (
		in = 7
		expect = 13
	)

	actual := Fib(in)
	if actual != expect {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expect)
	}
}
```

3. go test 执行测试，查看结果
```
C:\install-tools\goland\golang-project\grpc-test\benchmark>go test .
ok      benchmark       0.922s

```


### 单元测试 Table-Driven (输入输出表格)

使用结构体数据，多数据进行测试。
```
func TestFib(t *testing.T) {
	//  struct 数据
	var fibTest = []struct{
		in int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}
	for _, item := range fibTest {
		actual := Fib(item.in)
		if actual != item.expected {
			t.Errorf("Fib(%d) = %d; expected %d", item.in, actual, item.expected)
		}
	}
}
```


### Parallel 测试

**Parallel 方法表示当前测试只会与其他带有 Parallel 方法的测试并行进行测试。**

1. 需要进行并行测试的函数代码：
```
var (
	//  同时读写 data，需要加锁
	data = make(map[string]string)
	//  锁
	locker sync.RWMutex
)

//  写
func WriteToMap(k, v string) {
	locker.Lock()
	defer locker.Unlock()
	data[k] = v
}

// 读
func ReadFromMap(k string) string {
	locker.RLocker()
	defer locker.RUnlock()
	return data[k]
}
```
2. 编写 Parallel 测试 代码：
```

var testPairs = []struct{
	k string
	v string
}{
	{"Golang", " gsh1"},
	{"C++", "gsh2"},
	{"Python", " gsh3"},
	{"Java", "gsh4"},
	{"JavaScript", " gsh5"},
	{"Rust", "gsh6"},
	{"Scala", " gsh7"},
	{"C", "gsh8"},
	{"PHP", " gsh9"},
	{"TypeScript", "gsh10"},
	{"Shell", " gsh11"},
	{"CSharp", "gsh12"},
}

// 注意 TestWriteToMap 需要在 TestReadFromMap 之前
func TestWriteToMap(t *testing.T) {
	t.Parallel()
	for _, pair := range testPairs {
		WriteToMap(pair.k, pair.v)
	}
}

//  Parallel 测试会并行的运行 具有 t.Parallel() 函数的 函数代码
func TestReadFromMap(t *testing.T) {
	t.Parallel()
	for _, pair := range testPairs {
		actual := ReadFromMap(pair.k)
		if actual != pair.v {
			t.Errorf("the value of key(%s) is %s, expected: %s", pair.k, actual, pair.v)
		}
	}
}
```


## Benchmark 基准测试

在 _test.go 结尾的测试文件中，如下形式的函数：
```
func BenchmarkXxx(*testing.B)
```
被认为是基准测试，通过 go test 命令，加上 -bench 标志来执行。多个基准测试按照顺序运行。



### Benchmark 简单示例

1. 编写需要被测试的函数代码：
```
func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}
```
2. 编写 benchmark 测试代码：
```
//  benchmark 需要传入 testing.B
func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}
```
3. 使用 go test 执行测试
    - go test 命令默认不运行 benchmark 用例的，如果我们想运行 benchmark 用例，则需要加上 -bench 参数。
    - `-bench` 参数支持传入一个正则表达式，匹配到的用例才会得到执行
```
C:\install-tools\goland-2020-3-2\golang-project\grpc-test\go_learn>go test -bench .
goos: windows
goarch: amd64
pkg: go_learn
cpu: Intel(R) Core(TM) i7-4800MQ CPU @ 2.70GHz
BenchmarkFib-8               202           5850107 ns/op
PASS
ok      go_learn        1.843s

```


### benchmark 工作原理

benchmark 用例的参数 b *testing.B，有个属性 `b.N` 表示这个用例需要运行的次数。`b.N` 对于每个用例都是不一样的。

`b.N` 从 1 开始，如果该用例能够在 1s 内完成，b.N 的值便会增加，再次执行。b.N 的值大概以 1, 2, 3, 5, 10, 20, 30, 50, 100 这样的序列递增，越到后面，增加得越快。

可以看到刚刚的输出：
```
BenchmarkFib-8               202           5850107 ns/op
```

注意：
- `BenchmarkFib-8` : -8 即 GOMAXPROCS, 默认等于 CPU 核数。可以通过 -cpu 参数改变 GOMAXPROCS，-cpu 支持传入一个列表作为参数，例如: `go test -bench='Fib$' -cpu=2,4 .`



### benchmark 计时方法


**有三个方法用于计时：**
- `StartTimer`：开始对测试进行计时。该方法会在基准测试开始时自动被调用，我们也可以在调用 StopTimer 之后恢复计时；
- `StopTimer`：停止对测试进行计时。当你需要执行一些复杂的初始化操作，并且你不想对这些操作进行测量时，就可以使用这个方法来暂时地停止计时；
- `ResetTimer`：对已经逝去的基准测试时间以及内存分配计数器进行清零。对于正在运行中的计时器，这个方法不会产生任何效果。本节开头有使用示例。


实例：
```
func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		nums := generateWithCap(10000)
		b.StartTimer()
		bubbleSort(nums)
	}
}
```


### benchmark 并行执行

通过 RunParallel 方法能够并行地执行给定的基准测试。RunParallel会创建出多个 goroutine，并将 b.N 分配给这些 goroutine 执行，其中 goroutine 数量的默认值为 GOMAXPROCS。

RunParallel 通常会与 -cpu 标志一同使用。


### benchmark 内存统计

内存统计有两种方式：
1. 在执行 `go test -bench="正则" .` 时，可以添加 -benchmem 参数，**-benchmem 参数可以度量内存分配的次数。**
2. 在编写测试代码时，可以使用 `ReportAllocs()` 方法用于打开当前基准测试的内存统计功能, 与 go test 使用 -benchmem 标志类似, 但 `ReportAllocs()` 只影响那些调用了该函数的基准测试。
```
func BenchmarkTmplExucte(b *testing.B) {
    b.ReportAllocs()
    ...
}
```


### benchmark 提升准确度


对于性能测试来说，提升测试准确度的一个重要手段就是增加测试的次数。

提升精准度的方法：

**可以使用 `-benchtime` 和 `-count` 两个参数达到这个目的。**


注意：
- benchmark 的默认时间是 1s，如果可以在1s内完成，那么就增加 B.N 的次数，再次执行
- `-benchtime` 参数修改默认时间，例如指定为 5s
- `-count` 参数可以用来设置 benchmark 的轮数。例如，进行 3 轮 benchmark。



示例：
```
C:\install-tools\goland-2020-3-2\golang-project\grpc-test\go_learn>go test -bench="Fib$" -benchtime=5s -count=5 -benchmem .
goos: windows
goarch: amd64
pkg: go_learn
cpu: Intel(R) Core(TM) i7-4800MQ CPU @ 2.70GHz
BenchmarkFib-8              1046           5616492 ns/op               0 B/op          0 allocs/op
BenchmarkFib-8              1082           5493025 ns/op               0 B/op          0 allocs/op
BenchmarkFib-8              1081           5523761 ns/op               0 B/op          0 allocs/op
BenchmarkFib-8              1082           5520381 ns/op               0 B/op          0 allocs/op
BenchmarkFib-8              1063           5793231 ns/op               0 B/op          0 allocs/op
PASS
ok      go_learn        32.845s

```

### benchmark 结果

根据上面的结果，可以发现:
- 1046 ：表示基准测试的迭代总次数 b.N
- 5616492 ns/op：表示平均每次迭代所消耗的纳秒数
- 0 B/op：表示平均每次迭代内存所分配的字节数
- 0 allocs/op：表示平均每次迭代的内存分配次数





### benchmark 测试不同的输入

例如针对 fib 函数，可以编写多个不同的数据进行测试：
```
func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}
```


benchmark 测试代码可以写为：
```
//  加入参数 i , 该函数微辅助函数
func benchmarkFib(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(n)
	}
}

func BenchmarkFib1(b *testing.B) { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B) { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B) { benchmarkFib(3, b) }
func BenchmarkFib4(b *testing.B) { benchmarkFib(4, b) }
```

执行 go test 结果为：
```
C:\install-tools\goland-2020-3-2\golang-project\grpc-test\go_learn>go test -bench .
goos: windows
goarch: amd64
pkg: go_learn
cpu: Intel(R) Core(TM) i7-4800MQ CPU @ 2.70GHz
BenchmarkFib1-8         567859976                2.109 ns/op
BenchmarkFib2-8         196573300                5.890 ns/op
BenchmarkFib3-8         124905565                9.550 ns/op
BenchmarkFib4-8         100000000               17.57 ns/op
PASS
ok      go_learn        7.233s

```


### benchmark 生成 profile 文件，利用 go tool pprof 分析

testing 支持生成 CPU、memory 和 block 的 profile 文件。
- `-cpuprofile=$FILE`
- `-memprofile=$FILE -memprofilerate=N` 调整记录速率为原来的 1/N。
- `-blockprofile=$FILE`


生成了 profile 文件后，就可以使用 go tool pprof 进行分析了
1. 可以在终端进行分析 ： `go tool pprof /path/to/profile`
2. 可以在网页上查看分析结果：`go tool pprof -http=ip:port /path/to/profile`


注意：
- 只需要在 go test 添加 -cpuprofile 参数即可生成 BenchmarkFib 对应的 CPU profile 文件：
- 生成其他类似



示例：
1. 生成 pprof 文件
```
C:\install-tools\goland-2020-3-2\golang-project\grpc-test\go_learn>go test -bench="Fib$" -cpuprofile=./cpu.pprof .
goos: windows
goarch: amd64
pkg: go_learn
cpu: Intel(R) Core(TM) i7-4800MQ CPU @ 2.70GHz
BenchmarkFib-8               171           6165956 ns/op
PASS
ok      go_learn        2.270s

C:\install-tools\goland-2020-3-2\golang-project\grpc-test\go_learn>dir
 驱动器 C 中的卷是 系统
 卷的序列号是 C05D-7CEA

 C:\install-tools\goland-2020-3-2\golang-project\grpc-test\go_learn 的目录

2021/04/21  09:17    <DIR>          .
2021/04/21  09:17    <DIR>          ..
2021/04/21  09:17             3,885 cpu.pprof
2021/04/21  09:16               482 fib_benchmark_test.go
2021/04/20  19:16                25 go.mod
2021/04/21  09:17         3,336,704 go_learn.test.exe

```
2. 查看分析记过
```
C:\install-tools\goland-2020-3-2\golang-project\grpc-test\go_learn>go tool pprof -http=:1234 ./cpu.pprof
Serving web UI on http://localhost:1234
```

就可以打开浏览器 http://localhost:1234 查看性能分析了。