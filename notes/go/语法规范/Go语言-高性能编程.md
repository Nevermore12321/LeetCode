# Go 高性能编程注意点

[toc]





## 常用的数据结构的使用




### 1. 字符串拼接注意点

字符串拼接有多种方法：
1. 使用`+`拼接： 
```go
s = "a" + "b"
```
2. 使用`fmt.Sprintf`拼接：
    - 严格来说，`fmt.Sprintf` 是用来格式化字符串的
```go
s = fmt.Sprintf("%s%s", s1, s2)
```

3. 使用`bytes.Buffer`拼接：
```go
buf := new(bytes.Buffer)
buf.WriteString(s1)
buf.WriteString(s2)
s := buf.String()
```
4. 使用`strings.Builder`拼接：
```go
// strings.Builder的0值可以直接使用
var builder strings.Builder

// 预分配内存的方式 Grow
builder.Grow(n * len(str))

// 向builder中写入字符/字符串
builder.Write([]byte("Hello"))
builder.WriteByte(' ')
builder.WriteString("World")

// String() 方法获得拼接的字符串
builder.String() // "Hello World"
```
5. 使用`[]byte`拼接：
```go
buf := make([]byte, 0)
buf = append(buf, str...)
```

**性能比较**

- 使用 `+` 和 `fmt.Sprintf` 的效率是最低的，和其余的方式相比，性能相差约 1000 倍，而且消耗了超过 1000 倍的内存。
- `fmt.Sprintf` 通常是用来格式化字符串的，一般不会用来拼接字符串。
- `strings.Builder`、`bytes.Buffer` 和 `[]byte` 的性能差距不大，而且消耗的内存也十分接近
- 性能最好且消耗内存最小的是 `[]byte` ，这种方式预分配了内存，在字符串拼接的过程中，不需要进行字符串的拷贝，也不需要分配新的内存，因此性能最好，且内存消耗最小。



**使用推荐**
- 推荐使用 `strings.Builder` 来拼接字符串




### 2. 切片的使用


**使用方法**：


1. 使用内置函数 make 进行初始化
    - T 即元素类型
    - len 是切片的长度
    - cap 是切片的容量，容量是可选参数，默认等于长度。
```go
func make([]T, len, cap) []T
```
2. 使用字面量初始化时

注意：

容量是当前切片已经预分配的内存能够容纳的元素个数，如果往切片中不断地增加新的元素。**如果超过了当前切片的容量，就需要分配新的内存，并将当前切片所有的元素拷贝到新的内存块上**。因此为了减少内存的拷贝次数，**容量在比较小的时候，一般是以 2 的倍数扩大的**，例如 2 4 8 16 …，**当达到 2048 时，会采取新的策略，避免申请内存过大，导致浪费**。


切片操作并不复制切片指向的元素，创建一个新的切片会复用原来切片的底层数组，因此切片操作是非常高效的。

因此，如果一个切片从另一个切片中截取，那么这两个切片的底层数组是一样的，而且，无论哪一个切片修改，都会修改底层数组。



**切片操作及性能**

各个操作的原理，可以参考：https://ueokande.github.io/go-slice-tricks/

- `copy(b, a)` : copy 的切片和原切片公用一个底层内存空间
- `a = append(a, b...)`: 在结尾处追加
- `delete`: 切片的底层是数组，因此删除意味着后面的元素需要逐个向前移位。每次删除的复杂度为 O(N)，因此切片不合适大量随机删除的场景，这种场景下适合使用链表。
- `insert`: insert 和 append 类似。即在某个位置添加一个元素后，将该位置后面的元素再 append 回去。复杂度为 O(N)。因此，不适合大量随机插入的场景。
- `push`: 在末尾追加元素，不考虑内存拷贝的情况，复杂度为 O(1)。在头部追加元素，时间和空间复杂度均为 O(N)，不推荐。
- `pop`: 尾部删除元素，复杂度 O(1).头部删除元素，如果使用切片方式，复杂度为 O(1)。但是需要注意的是，底层数组没有发生改变，第 0 个位置的内存仍旧没有释放。如果有大量这样的操作，头部的内存会一直被占用。





### for range 的使用


for range 可以遍历 array/slice、map、channal，但是有些许不同：
1. 遍历 array/slice
    - 在遍历中，可以使用下标修改数组内容
    - 如果在循环中修改切片的长度不会改变本次循环的次数。也就是在循环内，如果增加或删除元素，遍历的次数不会改变
```go
words := []string{"Go", "语言", "高性能", "编程"}
for i, s := range words {
    words = append(words, "test")
    fmt.Println(i, s)
}
```
2. 遍历 map
    - 在迭代过程中，如果创建新的键值对，那么新增键值对，可能被迭代，也可能不会被迭代。
    - 和切片不同的是，迭代过程中，删除还未迭代到的键值对，则该键值对不会被迭代。
```go
m := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
}
for k, v := range m {
    delete(m, "two")
    m["four"] = 4
    fmt.Printf("%v: %v\n", k, v)
}
```
3. 遍历 channel
    - 发送给信道(channel) 的值可以使用 for 循环迭代，直到信道被关闭。
    - **如果是 nil 信道，循环将永远阻塞。**
```go
for n := range ch {
    fmt.Println(n)
}
```

**for 和 for range 的性能比较**

- range 在迭代过程中返回的是迭代值的拷贝，如果每次迭代的元素的内存占用很低，那么 for 和 range 的性能几乎是一样，例如 []int。
- 但是如果迭代的元素内存占用较高，例如一个包含很多属性的 struct 结构体，那么 for 的性能将显著地高于 range，有时候甚至会有上千倍的性能差异。对于这种场景，建议使用 for，如果使用 range，建议只迭代下标，通过下标访问迭代值，这种使用方式和 for 就没有区别了。
- 如果想使用 range 同时迭代下标和值，则需要将切片/数组的元素改为指针，才能不影响性能。




### 反射机制

反射机制是指在运行时态能够调用对象的方法和属性，主要用途就是使给定的程序能够动态地适应不同的运行情况


reflect 包里定义了一个接口和一个结构体，即 reflect.Type 和 reflect.Value，它们提供很多函数来获取存储在接口里的类型信息。
- `reflect.Type`:  主要提供关于类型相关的信息
- `reflect.Value`: 可以获取甚至改变类型的值



reflect 包中提供了两个基础的关于反射的函数来获取上述的接口和结构体：
- `reflect.TypeOf` 能获取`reflect.Type`类型信息；
- `reflect.ValueOf` 能获取数据的运行时表示`reflect.Value`；




#### 1. reflect.TypeOf

reflect.TypeOf 的定义规则为：
```go
func TypeOf(i interface{}) Type 
```

TypeOf 函数用来提取一个接口中值的类型信息。由于它的输入参数是一个空的 interface{}，调用此函数时，实参会先被转化为 interface{} 类型。这样，实参的类型信息、方法集、值信息都存储到 interface{} 变量里了。

返回值 Type 实际上是一个接口，定义了很多方法，用来获取类型相关的各种信息：
```go
type Type interface {
    // 所有的类型都可以调用下面这些函数

    // 此类型的变量对齐后所占用的字节数
    Align() int

    // 如果是 struct 的字段，对齐后占用的字节数
    FieldAlign() int

    // 返回类型方法集里的第 `i` (传入的参数)个方法
    Method(int) Method

    // 通过名称获取方法
    MethodByName(string) (Method, bool)

    // 获取类型方法集里导出的方法个数
    NumMethod() int

    // 类型名称
    Name() string

    // 返回类型所在的路径，如：encoding/base64
    PkgPath() string

    // 返回类型的大小，和 unsafe.Sizeof 功能类似
    Size() uintptr

    // 返回类型的字符串表示形式
    String() string

    // 返回类型的类型值
    Kind() Kind

    // 类型是否实现了接口 u
    Implements(u Type) bool

    // 是否可以赋值给 u
    AssignableTo(u Type) bool

    // 是否可以类型转换成 u
    ConvertibleTo(u Type) bool

    // 类型是否可以比较
    Comparable() bool

    // 下面这些函数只有特定类型可以调用
    // 如：Key, Elem 两个方法就只能是 Map 类型才能调用

    // 类型所占据的位数
    Bits() int

    // 返回通道的方向，只能是 chan 类型调用
    ChanDir() ChanDir

    // 返回类型是否是可变参数，只能是 func 类型调用
    // 比如 t 是类型 func(x int, y ... float64)
    // 那么 t.IsVariadic() == true
    IsVariadic() bool

    // 返回内部子元素类型，只能由类型 Array, Chan, Map, Ptr, or Slice 调用
    Elem() Type

    // 返回结构体类型的第 i 个字段，只能是结构体类型调用
    // 如果 i 超过了总字段数，就会 panic
    Field(i int) StructField

    // 返回嵌套的结构体的字段
    FieldByIndex(index []int) StructField

    // 通过字段名称获取字段
    FieldByName(name string) (StructField, bool)

    // FieldByNameFunc returns the struct field with a name
    // 返回名称符合 func 函数的字段
    FieldByNameFunc(match func(string) bool) (StructField, bool)

    // 获取函数类型的第 i 个参数的类型
    In(i int) Type

    // 返回 map 的 key 类型，只能由类型 map 调用
    Key() Type

    // 返回 Array 的长度，只能由类型 Array 调用
    Len() int

    // 返回类型字段的数量，只能由类型 Struct 调用
    NumField() int

    // 返回函数类型的输入参数个数
    NumIn() int

    // 返回函数类型的返回值个数
    NumOut() int

    // 返回函数类型的第 i 个值的类型
    Out(i int) Type

    // 返回类型结构体的相同部分
    common() *rtype

    // 返回类型结构体的不同部分
    uncommon() *uncommonType
}
```



#### 2. reflect.ValueOf

ValueOf 函数。返回值 reflect.Value 表示 interface{} 里存储的实际变量，它能提供实际变量的各种信息。相关的方法常常是需要结合类型信息和值信息。


ValueOf 函数的定义为：
```go
func ValueOf(i interface{}) Value
```



Value 结构体定义了很多方法，通过这些方法可以直接操作 Value 字段 ptr 所指向的实际数据：
```go
// 设置切片的 len 字段，如果类型不是切片，就会panic
 func (v Value) SetLen(n int)

 // 设置切片的 cap 字段
 func (v Value) SetCap(n int)

 // 设置字典的 kv
 func (v Value) SetMapIndex(key, val Value)

 // 返回切片、字符串、数组的索引 i 处的值
 func (v Value) Index(i int) Value

 // 根据名称获取结构体的内部字段值
 func (v Value) FieldByName(name string) Value

 // ……
Value 字段还有很多其他的方法。例如：

// 用来获取 int 类型的值
func (v Value) Int() int64

// 用来获取结构体字段（成员）数量
func (v Value) NumField() int

// 尝试向通道发送数据（不会阻塞）
func (v Value) TrySend(x reflect.Value) bool

// 通过参数列表 in 调用 v 值所代表的函数（或方法
func (v Value) Call(in []Value) (r []Value) 

// 调用变参长度可变的函数
func (v Value) CallSlice(in []Value) []Value
```
还有很多，没有列举完全。



#### 3. 反射的三大定律

反射三大定律：
1. 从 interface{} 变量可以反射出反射对象；
    - 反射是一种检测存储在 interface 中的类型和值机制。这可以通过 TypeOf 函数和 ValueOf 函数得到。
2. 从反射对象可以获取 interface{} 变量；
    - 将 ValueOf 的返回值通过 Interface() 函数反向转变成 interface 变量。
3. 要修改反射对象，其值必须可设置；
    - 反射变量可设置的本质是它存储了原变量本身，这样对反射变量的操作，就会反映到原变量本身；



### 提高反射的性能


#### 避免使用反射

避免使用反射
使用反射赋值，效率非常低下，如果有替代方案，尽可能避免使用反射，特别是会被反复调用的热点代码。例如 RPC 协议中，需要对结构体进行序列化和反序列化，这个时候避免使用 Go 语言自带的 json 的 Marshal 和 Unmarshal 方法，因为标准库中的 json 序列化和反序列化是利用反射实现的。可选的替代方案有 easyjson，在大部分场景下，相比标准库，有 5 倍左右的性能提升。

#### 缓存

FieldByName 相比于 Field 有一个数量级的性能劣化。那在实际的应用中，就要避免直接调用 FieldByName。我们可以利用字典将 Name 和 Index 的映射缓存起来。避免每次反复查找，耗费大量的时间。




### 空结构体 struct{} 的使用


#### 1. 空结构体不占用空间

可以使用 unsafe.Sizeof 计算出一个数据类型实例需要占用的字节数。


```go
fmt.Println(unsafe.Sizeof(struct {}{}))
```

打印结果是0，因此 空结构体不占用空间。




#### 2. 空结构体的作用

因为空结构体不占据内存空间，因此被广泛作为各种场景下的占位符使用。一是节省资源，二是空结构体本身就具备很强的语义，即这里不需要任何值，仅作为占位符。


空结构体可以用作：
1. 可以实现 Set，使用map来实现集合Set，由于 集合 只有 key 没有 value，因此可以将value置为空结构体来节省空间
```go
type Set map[string]struct{}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}
...
```
2. 不发送数据的信道(channel)
    - 有时候使用 channel 不需要发送任何的数据，只用来通知子协程(goroutine)执行任务，或只用来控制协程并发度。
```go
func worker(ch chan struct{}) {
	<- ch
	fmt.Println("worker start")
	close(ch)
}

func main() {
	fmt.Println("main start")
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}
}
```
3. 仅包含方法的结构体
    - 在部分场景下，结构体只包含方法，不包含任何的字段。
```go
type Door struct{}

func (d Door) Open() {
	fmt.Println("Open the door")
}

func (d Door) Close() {
	fmt.Println("Close the door")
}
```


### 内存对齐


操作系统并非一个字节一个字节访问内存，而是按2, 4, 8这样的字长来访问。因此，当 CPU 从存储器读数据到寄存器，或者从寄存器写数据到存储器，IO 的数据长度通常是字长。如 32 位系统访问粒度是 4 字节（bytes），64 位系统的是 8 字节。


CPU 访问内存时，并不是逐个字节访问，而是以字长（word size）为单位访问。比如 32 位的 CPU ，字长为 4 字节，那么 CPU 访问内存的单位也是 4 字节。

这么设计的目的，是减少 CPU 访问内存的次数，加大 CPU 访问内存的吞吐量。比如同样读取 8 个字节的数据，一次读取 4 个字节那么只需要读取 2 次。


内存对齐，也就是说，A 类型占 3 个字节，B 类型也占 3 个字节，如果不对齐，那么 AB 是连续存储的，一共占6字节，操作系统在访问内存时，假设以4字节长度读取，那么读取B类型时，就需要读取两次，先读第一个4字节，然后都下一个4字节。





**unsafe.Alignof**

unsafe 标准库提供了 Alignof 方法，可以返回一个类型的对齐值，也可以叫做对齐系数或者对齐倍数

效果：
- 对于任意类型的变量 x ，unsafe.Alignof(x) 至少为 1。
- 对于 struct 结构体类型的变量 x，计算 x 每一个字段 f 的 unsafe.Alignof(x.f)，unsafe.Alignof(x) 等于其中的最大值。也就是等于结构体中长度最大的值
- 对于 array 数组类型的变量 x，unsafe.Alignof(x) 等于构成数组的元素类型的对齐倍数。




**合理布局减少内存占用**

下面不同顺序有不同的结果：
```go
type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	c int32
	b int16
}

func main() {
	fmt.Println(unsafe.Sizeof(demo1{})) // 8
	fmt.Println(unsafe.Sizeof(demo2{})) // 12
}
```

首先，各个类型的占用内存的大小为：
- int8: 8位,就是一个字节
- int16: 2个字节
- int32: 4个字节
- int64:8个字节

（如果以32位系统为例，cpu读取以4字节为单位）
解析：

- demo1 中，a 占1个字节，b 占2字节，不超过4字节，因此a和b可以在同一个 4 字节中，c 占4字节，就在下一个4字节中，一共占8字节
- demo2 中，a 占1个字节，b 占4字节，已经超过4字节，因此，b在下一个4字节中，c 占2字节，在第三个4字节中，因此一共 4 + 4 + 4 = 12



**空 struct{} 的对齐**

- 空 struct{} 大小为 0，作为其他 struct 的字段时，一般不需要内存对齐。
- 但是有一种情况除外：即当 struct{} 作为结构体最后一个字段时，需要内存对齐。因为如果有指针指向该字段, 返回的地址将在结构体之外，如果此指针一直存活不释放对应的内存，就会有内存泄露的问题（该内存不因结构体释放而释放）。

```go
type demo3 struct {
	c int32
	a struct{}
}

type demo4 struct {
	a struct{}
	c int32
}

func main() {
	fmt.Println(unsafe.Sizeof(demo3{})) // 8
	fmt.Println(unsafe.Sizeof(demo4{})) // 4
}
```

可以看到，demo4{} 的大小为 4 字节，与字段 c 占据空间一致，而 demo3{} 的大小为 8 字节，即额外填充了 4 字节的空间。




## 并发编程


### 读写锁和互斥锁的区别


#### 1. 互斥锁(sync.Mutex)

互斥即不可同时运行。即使用了互斥锁的两个代码片段互相排斥，只有其中一个代码片段执行完成后，另一个才能执行。

Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法：
- Lock 加锁
- Unlock 释放锁


可以通过在代码前调用 Lock 方法，在代码后调用 Unlock 方法来保证一段代码的互斥执行，也可以用 defer 语句来保证互斥锁一定会被解锁。在一个 Go 协程调用 Lock 方法获得锁后，其他请求锁的协程都会阻塞在 Lock 方法，直到锁被释放。



#### 2. 读写锁(sync.RWMutex)


读写锁，读写锁分为读锁和写锁，读锁是允许同时执行的，但写锁是互斥的。一般来说，有如下几种情况：
- 读锁之间不互斥，没有写锁的情况下，读锁是无阻塞的，多个协程可以同时获得读锁。
- 写锁之间是互斥的，存在写锁，其他写锁阻塞。
- 写锁与读锁是互斥的，如果存在读锁，写锁阻塞，如果存在写锁，读锁阻塞。


Go 标准库中提供了 sync.RWMutex 互斥锁类型及其四个方法：
- Lock 加写锁
- Unlock 释放写锁
- RLock 加读锁
- RUnlock 释放读锁

读写锁的存在是为了解决读多写少时的性能问题，读场景较多时，读写锁可有效地减少锁阻塞的时间。



#### 3. 饥饿模式

**互斥锁有两种状态：**

- 正常状态
- 饥饿状态。


在正常状态下，所有等待锁的 goroutine 按照 FIFO 顺序等待。唤醒的 goroutine 不会直接拥有锁，而是会和新请求锁的 goroutine 竞争锁的拥有。

新请求锁的 goroutine 具有优势：它正在 CPU 上执行，而且可能有好几个，所以刚刚唤醒的 goroutine 有很大可能在锁竞争中失败。在这种情况下，这个被唤醒的 goroutine 会加入到等待队列的前面。 如果一个等待的 goroutine 超过 1ms 没有获取锁，那么它将会把锁转变为饥饿模式。

在饥饿模式下，锁的所有权将从 unlock 的 goroutine 直接交给交给等待队列中的第一个。新来的 goroutine 将不会尝试去获得锁，即使锁看起来是 unlock 状态, 也不会去尝试自旋操作，而是放在等待队列的尾部。

如果一个等待的 goroutine 获取了锁，并且满足一以下其中的任何一个条件：
- 它是队列中的最后一个；
- 它等待的时候小于1ms。它会将锁的状态转换为正常状态。


正常状态有很好的性能表现，饥饿模式也是非常重要的，因为它能阻止尾部延迟的现象。






### 协程的退出


#### 超时控制，协程退出

**超时控制的逻辑：**
```go
// 模拟 超时 操作，操作完成后，向 channal 中写入完成标志
func DoBadThing(done chan bool) {
	//  操作超时模拟
	time.Sleep(2 * time.Second)
	//  操作完成，传递完成标志
	done <- true
}

//  超时处理 的逻辑， 参数是 想要 对某个函数进行超时监听
func Timeout(f func(chan bool)) error {
	done := make(chan bool)
	//  调用操作
	go f(done)
	select {
	//  如果 操作完成，那么正常逻辑
	case <-done:
		fmt.Println("Exit normally")
		return nil
	case <-time.After(100 * time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func main() {
	err := mytest.Timeout(mytest.DoBadThing)
	if err != nil {
		fmt.Println(err)
	}
}
```

解析：
- 最终的结果是 "timeout"
- 利用 time.After 启动了一个异步的定时器，返回一个 channel，当超过指定的时间后，该 channel 将会接受到信号。
- 启动了子协程执行函数 f，函数执行结束后，将向 channel done 发送结束信号。
- 使用 select 阻塞等待 done 或 time.After 的信息，若超时，则返回错误，若没有超时，则返回 nil。


问题：
- 如果启动多个 Timeout 协程，并且多个操作都超时了，那么协程是否可以正常退出呢？答案是不可以。
- 因为子协程在操作完成后，想要往done channel 中写入时，done 是一个无缓冲区的 channel ，当超时发生时，select 接收到 time.After 的超时信号就返回了，done 没有了接收方(receiver)，而 doBadthing 在执行 1s 后向 done 发送信号，由于没有接收者且无缓存区，发送者(sender)会一直阻塞，导致协程不能退出




**解决超时控制的子协程退出**
- 创建有缓冲区的 channel
- 在DoBadTing中，使用 select 尝试发送
```
func doGoodthing(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}
}
```


注意：  
**goroutine 只能自己退出，而不能被其他 goroutine 强制关闭或杀死。**




### sync.Pool 复用对象

保存和复用临时对象，减少内存分配，降低 GC 压力。

sync.Pool 用于存储那些被分配了但是没有被使用，而未来可能会使用的值。这样就可以不用再次经过内存分配，可直接复用已有对象，减轻 GC 的压力，从而提升系统的性能。

sync.Pool 的大小是可伸缩的，高负载时会动态扩容，存放在池中的对象如果不活跃了会被自动清理。


**使用方法**
```
//  创建 pool 变量
var pool *sync.Pool

type Person struct {
	Name string
	Id int
}


func initPool() {
	//  创建 Person 类型 的 Pool，类似于池
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new pool")
			return new(Person)
		},
	}
}

func main() {
	initPool()
	p := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)

	p.Name = "gsh"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	pool.Put(p)
	p2 := &Person{Name: "zy", Id: 1}
	pool.Put(p2)
	fmt.Println("Pool 里的第一个对象：&{first}，调用 Get: ", pool.Get().(*Person))
	fmt.Println("Pool 里的第二个对象：&{Second}，调用 Get: ", pool.Get().(*Person))
	fmt.Println("Pool 没有对象了，调用 Get: ", pool.Get().(*Person))
}

```

解析：
- `Get()` 用于从对象池中获取对象，因为返回值是 interface{}，因此需要类型转换。
- `Put()` 则是在对象使用完毕后，返回对象池。
- Pool类似一个数据池，先进先出，但是在并发场景下，不能确定进出的顺序，最好的做法是在 Put 前，将对象清空。



### sync.Once

sync.Once 是 Go 标准库提供的使函数只执行一次的实现，常应用于单例模式，例如初始化配置、保持数据库连接等。作用与 init 函数类似，但有区别。
- init 函数是当所在的 package 首次被加载时执行，若迟迟未被使用，则既浪费了内存，又延长了程序加载时间。
- sync.Once 可以在代码的任意位置初始化和调用，因此可以延迟到使用时再执行，并发场景下是线程安全的。


**使用方法**

sync.Once 仅提供了一个方法 Do，参数 f 是对象初始化函数。

```
var (
	once   sync.Once
	config *Config
)

func ReadConfig() *Config {
	once.Do(func() {
		var err error
		config = &Config{Server: os.Getenv("TT_SERVER_URL")}
		config.Port, err = strconv.ParseInt(os.Getenv("TT_PORT"), 10, 0)
		if err != nil {
			config.Port = 8080 // default port
        }
        log.Println("init config")
	})
	return config
}
```



### sync.Cond

sync.Cond 条件变量用来协调想要访问共享资源的那些 goroutine，当共享资源的状态发生变化的时候，它可以用来通知被互斥锁阻塞的 goroutine。



**sync.Cond 的四个方法**
- `func NewCond(l Locker) *Cond` : NewCond 创建 Cond 实例时，需要关联一个锁。
- `func (c *Cond) Broadcast()` : Broadcast 唤醒所有等待条件变量 c 的 goroutine，无需锁保护。
- `func (c *Cond) Signal()` : Signal 只唤醒任意 1 个等待条件变量 c 的 goroutine，无需锁保护。
- `func (c *Cond) Wait()` : 调用 Wait 会自动释放锁 c.L，并挂起调用者所在的 goroutine，因此当前协程会阻塞在 Wait 方法调用的地方。如果其他协程调用了 Signal 或 Broadcast 唤醒了该协程，那么 Wait 方法在结束阻塞时，会重新给 c.L 加锁，并且继续执行 Wait 后面的代码。


实例：
```
package main

import (
	"log"
	"sync"
	"time"
)

//  写完成的标志
var done = false

//  读操作
func read(name string, c *sync.Cond) {
	//  加锁，读取
	c.L.Lock()
	//  如果 done 时false，即写还没完成
	if !done {
		c.Wait()
	}

	log.Println(name, " starts reading")
	c.L.Unlock()
}

//  写操作
func write(name string, c *sync.Cond) {
	log.Println(name, " starts writing")
	time.Sleep(2 * time.Second)
	//  加锁， 写入
	c.L.Lock()
	done = true
	c.L.Unlock()

	log.Println(name, "wakes all")
	//  广播给所有的等待的线程
	c.Broadcast()

}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}

```