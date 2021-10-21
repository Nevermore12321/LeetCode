# Go语言-初步

[toc]



## Go语言优势：
- 可直接编译成机器码，不依赖其他库，glibc的版本有一定要求，部署就是扔一个文件上去就完成了。
- 静态类型语言，但是有动态类型语言的感觉，静态类型语言就是可以在编译的时候检查出大部分问题，动态类型语言的感觉意思就是有很多的包可以使用，写起来效率很高。
- 语言层面支持并发，这个是go最大的特色，天生支持并发，可以充分的利用多核，很容易使用并发。
- 内置runtime，支持垃圾回收，这是动态语言的特性之一，虽然目前来说GC（内存垃圾回收机制）不算完美，但足以应付我们现在所能遇到的大多数情况。
- 简单易学，Go语言的作者有C的基因，Go的关键字是25个但是表达能力很强大，几乎支持大多数其他语言的特性，继承、重载、对象等。
- 丰富的标准库，Go目前支持大量的库，特别是网络库。
- 内置强大的工具，Go语言内置了很多工具链，最好的应该是gofmt工具，自动化格式化代码，能够让团队review更简单，代码模式一样。
- 跨平台编译，如果Go代码中不包括cgo，那么就可以做到windows系统编译linux的应用。
- 内嵌C支持，Go中可以直接包含C代码，利用现有丰富的C库。


## Go语言用途
- 服务器编程，以前利用C或C++做的事情，用Go来做很适合，例如处理日志、数据打包、虚拟机处理、文件系统等。
- 分布式系统，数据库代理器等。
- 网络编程：目前应用最广，包括Web应用，API应用，下载应用。
- 内存数据库，例如Google开发的groupcache，couchbase的部分组件。
- 云平台，目前有不少云平台是用Go语言开发。


## 语法初步

### 1. 变量

#### 变量需要注意的几点：
- 变量的声明：
    - 声明格式：`var 变量名 类型`。例：`var a int`
    - 变量一旦声明，必须使用，否则报错。
    - 只是声明，没有初始化的变量，默认值为0。
    - 在同一个代码块（即{}）中，声明的变量名是唯一的。
    - 可以同时声明多个变量，格式：`var 变量1,变量2… 类型`。例：`var a,b int`
    - 变量的赋值，直接变量名后跟=赋值。例：a=10
- 变量初始化和自动推到类型：
    - 初始化： 即在声明的同时赋值，格式：`var 变量 类型 = 值`。例：`var a int = 10`
    - 自动推导类型：通过初始化的值来确定声明的类型。
        - 格式：`变量 := 值`。例：`e := 10`，e就被声明为int类型并且初始化为10.
        - 函数外的每个语句都必须以关键字开始（var， func等等），因此 := 结构不能在函数外使用。
- 变量的多重赋值：
    - 格式：`变量1,变量2=值1,值2`。例： `a, b = 10, 20`
    - 多重赋值的一个典型应用：交换两个变量的值。例：交换a，b的值，`a, b = b, a`
- 匿名变量：
    - 格式:  `变量1，_ = 值1, 值2`。例：`a, _ = 10, 20`
    - 匿名变量的典型应用：匿名函数多配合函数的返回值一起使用，可以剔除不需要的返回值。例：函数f返回三个值，只想要两边的不想要中间的值，可以利用匿名变量：`a, _, b = f()`
- 常量：
    - 变量用var来声明，常量用const来声明，格式：`const 变量 类型 = 值`。例：`const a int = 10`
    - 常量的自动类型推导中，没有冒号，直接=，格式：`const 变量 = 值`。例：`const b = 10`
- 多个变量或常量的定义：
    - 定义多个变量或常量，可以使用：
    ```
    var {
        a int
        b float64
    }
    const {
        a int = 10
        b float64 = 12.2
    }
    ```
    
#### 打印函数

打印函数fmt.Printf和fmt.Println的区别：  
- fmt.Printf     是格式化输出。例：`fmt.Printf(“a is %d”, a)`
- fmt.Println 是字符串拼接后输出。例：`fmt.Println(“a is “, d)`

占位符的说明：
- 普通占位符  

| 占位符 | 说明                           | 举例                    | 输出                        |
| ------ | ------------------------------ | ----------------------- | --------------------------- |
| %v     | 相应值的默认格式。             | `Printf("%v", people)`  | {zhangsan}                  |
| %+v    | 打印结构体时，会添加字段名     | `Printf("%+v", people)` | {Name:zhangsan}             |
| %#v    | 相应值的Go语法表示             | `Printf("#v", people)`  | main.Human{Name:"zhangsan"} |
| %T     | 相应值的类型的Go语法表示       | `Printf("%T", people)`  | main.Human                  |
| %%     | 字面上的百分号，并非值的占位符 | `Printf("%%")`          | %                           |

- 布尔占位符

| 占位符 | 说明          | 举例                 | 输出 |
| ------ | ------------- | -------------------- | ---- |
| %t     | true 或 false | `Printf("%t", true)` | true |


- 整数占位符

| 占位符 | 说明                                       | 举例                   | 输出   |
| ------ | ------------------------------------------ | ---------------------- | ------ |
| %b     | 二进制表示                                 | `Printf("%b", 5)`      | 101    |
| %c     | 相应Unicode码点所表示的字符                | `Printf("%c", 0x4E2D)` | 中     |
| %d     | 十进制表示                                 | `Printf("%d", 0x12)`   | 18     |
| %o     | 八进制表示                                 | `Printf("%d", 10)`     | 12     |
| %q     | 单引号围绕的字符字面值，由Go语法安全地转义 | `Printf("%q", 0x4E2D)` | '中'   |
| %x     | 十六进制表示，字母形式为小写 a-f           | `Printf("%x", 13)`     | d      |
| %X     | 十六进制表示，字母形式为大写 A-F           | `Printf("%x", 13)`     | D      |
| %U     | Unicode格式：U+1234，等同于 "U+%04X"       | `Printf("%U", 0x4E2D)` | U+4E2D |


- 浮点数和复数的组成部分（实部和虚部）

| 占位符 | 说明                                                         | 举例                     | 输出         |
| ------ | ------------------------------------------------------------ | ------------------------ | ------------ |
| %b     | 无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78 |                          |              |
| %e     | 科学计数法，例如 -1234.456e+78                               | `Printf("%e", 10.2)`     | 1.020000e+01 |
| %E     | 科学计数法，例如 -1234.456E+78                               | `Printf("%e", 10.2)`     | 1.020000E+01 |
| %f     | 有小数点而无指数，例如 123.456                               | `Printf("%f", 10.2)`     | 10.200000    |
| %g     | 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出        | `Printf("%g", 10.20)`    | 10.2         |
| %G     | 根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出        | `Printf("%G", 10.20+2i)` | (10.2+2i)    |


- 字符串与字节切片

| 占位符 | 说明                                   | 举例                             | 输出         |
| ------ | -------------------------------------- | -------------------------------- | ------------ |
| %s     | 输出字符串表示（string类型或[]byte)    | `Printf("%s", []byte("Go语言"))` | Go语言       |
| %q     | 双引号围绕的字符串，由Go语法安全地转义 | `Printf("%q", "Go语言")`         | "Go语言"     |
| %x     | 十六进制，小写字母，每字节两个字符     | `Printf("%x", "golang")`         | 676f6c616e67 |
| %X     | 十六进制，大写字母，每字节两个字符     | `Printf("%X", "golang")`         | 676F6C616E67 |


- 指针

| 占位符 | 说明                  | 举例                    | 输出     |
| ------ | --------------------- | ----------------------- | -------- |
| %p     | 十六进制表示，前缀 0x | `Printf("%p", &people)` | 0x4f57f0 |

#### 输入函数

对于变量值得输入，有两个常用函数，fmt,Scanf和fmt.Scan，两者的用法为：

- `fmt.Scanf("%d",&a)` : 必须指定%类型，才能保存到变量中
- `fmt.Scan(&b)` : 不需要指定，直接保存到变量中



#### iota枚举类型
- iota常量自动生成器，每个一行，自动累加1
- 可以只写一个iota，后面的会自动累加
- iota遇到const时，会重置为1，也就是在const（常量区），在常量区内，有iota会累加，下一个const中iota就会失效。

```
	const (
		a1 = iota
		a2 = iota
		a3
		a4
	)
	const (
		b1 = 10
		b2
		b3
	)
	fmt.Println(a1, a2, a3, a4, b1, b2, b3)

//  打印结果
0 1 2 3 10 10 10
```

#### 变量类型
| 类型          | 名称     | 长度 | 零值  | 说明                                       |
| ------------- | -------- | ---- | ----- | ------------------------------------------ |
| bool          | 布尔类型 | 1    | false | 其值为true或false，不可以用数字0,1代表真假 |
| byte          | 字节类型 | 1    | 0     | uint8的别名                                |
| rune          | 字符类型 | 4    | 0     | 专用于存储unicode编码，等价于uint32        |
| int、uint     | 整型     | 4、8 | 0     | 32位或64位                                 |
| int8、uint8   | 整型     | 1    | 0     | -128~127、0~255                            |
| int16、uint16 | 整型     | 2    | 0     | -32768~32767、0~65535                      |
| int32、uint32 | 整型     | 4    | 0     | -21亿~21亿，0~42亿                         |
| int64、uint64 | 整型     | 8    | 0     | 无                                         |
| float32       | 浮点型   | 4    | 0.0   | 小数位精确到7位                            |
| float64       | 浮点型   | 8    | 0.0   | 小数位精确到15位                           |
| complex64     | 复数类型 | 8    | 无    | 无                                         |
| complex128    | 复数类型 | 16   | 无    | 无                                         |
| uintptr       | 整型     | 4或8 | 无    | 足以存储指针的uint32或uint64的整数         |
| string        | 字符串   | 0    | “”    | utf-8的字符串                              |


注意：  
- 字符是单引号，字符串是双引号，字符串有一个常用的内置函数len()，计算字符串的长度。字符串隐藏了结尾符‘\0’。
- 字符串可以用下标索引，即str[0]。
- 复数类型有实部和虚部，取一个复数的实部和虚部的函数为：real()和imag()。


#### 类型转换：

表达式 T(v) 将值 v 转换为类型 T.

例：`var j float64 = float64(i)`

注意：bool类型与int整型不能相互转换。



#### 类型别名：

- 格式： `type char byte` : 
- byte类型重命名为char，以后就可以使用char来定义变量`var s1 char = 'a'`


### 2. 流程控制语句
Go语言支持三种基本的程序运行结构：顺序结构，选择结构，循环结构。

**选择结构**  
- 选择结构： if语句  
    - 表达式外无需小括号 () ，而大括号 {} 则是必须的。
    - if 语句可以在条件表达式前执行一个简单的语句。该语句声明的变量作用域仅在 if 之内。例：`if a==10 {…}`
    - if支持一个初始化语句，初始化语句和判断语句以分号分割。例：`if a:=10; a==10 {…}`
    - if else的使用：`if 判断语句 {…} else if 判断语句 {…} else {…}`
    - 在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。
- 选择结构： switch语句
    - switch语句的语法规则
    ```
    switch var1 {
        case val1:
            ...
        case val2:
            ...
        default:
            ...
    }
    ```
    - 每个switch语句的case下，都会默认存在有一个break语句，来跳出switch。
    - switch中的case下： **fallthrough** 语句，代表不跳出switch语句，并且后面的case无条件执行，直到遇到没有fallthrough的case后，会跳出switch。
    - switch同样也支持一个初始化语句，初始化语句与变量本身以分号分割。
    - case的值可以有多个，switch后可以没有变量。没有条件的 switch 同 switch true 一样。这种形式能将一长串 if-then-else 写得更加清晰。
    - Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
    
    

**循环结构**   
- 循环结构： for语句
    - for语句的规则：`for 初始化条件; 判断条件; 后置条件变化 {…}`
    - 初始化语句和后置语句是可选的。
    - for 是 Go 中的 “while”，可以类似while的使用：for sum<100 { … }
    - for 后面什么条件也不写，就是无限循环。例：for { … }
- 循环结构： range语句
    - range会默认返回两个值，一个是元素的位置，一个是元素本身，意思就是如果range str，一个是i，一个是str[i]。
    - 注意：range的第二个返回值默认是丢弃的，意思就是可以接收一个变量，也可以接收两个变量。
    - 可以将下标或值赋予 _ 来忽略它。若你只需要索引，去掉 value 的部分即可。
    - 例子：
    ```
    // 定义一个字符串变量
    var str string = "guoshaohe"
    
    // 遍历str，使用len
    for i := 0; i < len(str); i++ {
      fmt.Printf("str[%d] : %c\t", i, str[i])
    }
    
    //  遍历str，使用range
    //  注意range前的变量，如果是新的变量需要 := 来初始化
    //   接收range返回的两个参数
    for i, data := range str {
      fmt.Printf("str[%d] : %c\t", i, data)
    }
    
    //   接收range返回的第一个参数，第二个默认丢弃
    fmt.Println("")
    for i := range str {
      fmt.Printf("str[%d] : %c\t", i, str[i])
    }
    ```
    
    
### 3. 函数

#### 函数定义
- 函数的一般定义格式为：
```
//  函数的定义格式：
func FuncNmae (参数列表) (返回类型列表) {
    //函数体
    return v1，v2    //可以返回多个值，但与返回类型列表对应
}
```

- 函数定义放在main之前之后都可以正常使用。
- 无参数无返回值的函数定义： `func FuncName() { 函数体 }`
- 有参数无返回值的函数定义： `func FuncName(变量 变量类型) { 函数体 }`
    - 注意：当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
    - 例：`func FuncNmae(变量1，变量2 变量类型) { 函数体 }`          其中，变量1和变量2是相同类型的。
- 以上是固定参数，即参数列表是固定的个数
- 不定参数表示参数个数可以变化。定义：`func FuncName(变量1 变量类型，变量列表名 …不定参数类型) { 函数体 }`。例：`func myfucn(a int, list …string) { 函数体 }`
    - 注意：不定参数只能放在形参的最后。固定参数一定要传参数
    - 不定参数可传可不传。
    - 而将所有的不定参数传给另一个函数中，可以使用myfunc(arg …) 。例：
    ```
    func myfunc(a int, list ...string) {
    	fmt.Println("a =  ", a)
    	for i := range list {
    		fmt.Printf("list[%d] = %s\n", i, list[i])
    	}
    }
    //可以将所有的不定参数列表直接全部传给另一个函数的做不定参数
    func myfunc1(args ...string) {
    	myfunc(222, args ...)
        myfunc(333, args[:2] ...)
    }
    ```
- 函数的返回值.
    - 只有一个返回值：`func FuncName(参数列表）返回值类型 { 函数体 + return }`
    - 有多个返回值： `func FuncName(参数列表）(返回值类型列表) { 函数体 + return }`
    - 注意：go推荐写法，给返回值起一个名字，然后直接return，例：
    ```
    //   在返回值列表中列出名字和类型，直接return，则返回全部返回值
    func myfunc4(a, b int) (way string, result int) {
    	result = a + b
    	way = "add"
    	return
    }
    
    //   接收两个返回值
    func main() {
      	res1, res2 := myfunc4(1,2)
    	fmt.Println(res1, "'s result is :", res2)
    }
    ```

#### 函数类型

函数也是一种数据类型，可以通过type给函数类型起一个名字，然后通过函数类型类定义一个变量，通过赋值后进行调用。
```
//   定义两个函数，分别是加法和减法
func Add(a, b int) (res int) {
	res = a + b
	return 
}
func Minus(a, b int) (res int) {
	res = a - b 
	return res
}

//    type 给函数类型起一个名字
type FuncType func(int, int) int

//    就可以通过FuncType来定义变量了
func main() {
//定义一个变量为Functype类型
    var ftest FuncType
//赋值
    ftest = Add
//调用
    r := ftest(9, 3)
	fmt.Println("result is : ", r)
	ftest = Minus
	r = ftest(9, 3)
	fmt.Println("result is : ", r)
}
```

#### 回调函数

函数有一个参数是函数类型，这个函数就是回调函数。

注意： 回调函数可以实现多态，即调用同一个接口，可以有不同的实现。还有一点好处就是可以先定义回调函数，而具体的功能可以以后再做添加。

例：
```
func Add(a, b int) (res int) {
	res = a + b
	return 
}
func Minus(a, b int) (res int) {
	res = a - b 
	return res
}

type FuncType func(int, int) int

//   回调函数
func Calc(a, b int, ftest FuncType) (result int) {
	fmt.Println("In Calc:")
	result = ftest(a, b)
	return 
}

func main() {
  	result := Calc(1,3, Add)
	fmt.Println("Calc's result is : ", result)
	result = Calc(1,3, Minus)
	fmt.Println("Calc's result is : ", result)
}
```

#### 闭包和匿名函数

- 所谓的闭包，就是一个函数“捕获”了和它在同一作用域的其他常量和变量，这就意味着，当闭包被调用的时候，不管在程序什么位置调用，闭包可以使用这些常量和变量。
    - 闭包不关心这些常量和变量是否超出作用域，只要闭包在还在使用,这些变量就还会存在。
    - 注意：在Go语言中，所有的匿名函数（Go规范中称为函数字面量）都是闭包，匿名函数是指不需要定义函数名的一种函数的实现方式。
- 匿名函数：没有函数名的函数定义。
    - 例如：
    ```
    func main() {
        // 匿名函数通常用自动推导类型来赋值，
        f := func (a, b int) (max, min int){
    		if a > b {
    			max = a
    			min = b
    		} else {
    			max = b
    			min = a
    		}
    		return
    	}
        // 匿名函数调用可以直接在定义时在结尾大括号}后直接（）执行
    	x, y := f(12, 3)
    	fmt.Println("max is ", x, "min is ", y)
    
    }
    ```
    - 在匿名函数内，可以使用外部的变量和常量。并且只要闭包在还在使用，这些变量就还会存在。例：
    ```
    //  定义一个函数，并且返回一个匿名函数
    func test() func() int {
    	var x int
    	return func() int {
    		x++
    		return x * x
    	}
    }
    
    func main() {
    	f := test()
    //  多次调用，可以发现，每次调用，x的值会递增，因为匿名函数在使用变量x的期间，变量一直会存在
    	fmt.Println(f())
    	fmt.Println(f())
    	fmt.Println(f())
    	fmt.Println(f())
    }
    ```

#### 延迟调用defer关键字

- 关键字defer，用于延迟一个函数或方法（或者当前所创建的匿名函数）的执行。defer 语句会将函数推迟到外层函数返回之后执行。推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
    - defer语句经常被用于处理成对的操作，如打开、关闭，连接、断开链接，加锁、释放锁等。
    - 通过defer机制，不论函数多复杂，都能保证在任何执行路径下，释放资源的defer应该跟在请求资源的语句后面。
    - 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
    defer语句与匿名函数的结合使用，例：
```

```

#### 获取命令行参数：

使用os.Args可以获取到命令行的参数。例：list := os.Args     其中，list相当于一个字符串数组，里面存放了所有的命令行参数。


### 4. 复合类型

复合类型有：指针（pointer）、数组（array）、切片（slice）、字典（map）、结构体（struct）。

#### 指针类型

- 定义：`var 变量名 *类型`，一级指针，`var 变量名 **类型`，二级指针…
- Go语言指针与其他语言的不同之处在于：空指针为nil，没有NULL。并且不支持->操作，直接用.操作。
- 同C/C++类似，如果要想在函数中修改实参的值，那么要在函数定义时，将形参设置为该变量的指针，如果修改一级指针，那么传到函数中的就需要是二级指针。


**new函数**
- `new(T)`，将创建一个T类型的匿名变量，分配一块T类型大小的内存空间，并且返回内存空间的地址
- new函数的使用无需担心内存的生命周期或怎样将其删除，Go会自动处理这些问题。


#### 数组

- 定义： `var 数组名 [n]类型`。  例：`var id [10]int`
- 操作数组：
    - 可以使用len()来求数组的长度，然后用for遍历。
    - 还可以直接用 `range 数组名`  来操作数组，同样返回两个值，序号与序号所对应的数组元素。
- 数组做函数参数，如果要想修改数组的内容，需要将数组的地址传入，函数接收数组指针。
- 数组的初始化：
    - `var 数组名 [n]类型 = [n]类型{ 值1，值2，…. }` 。例：`var a [5]int = [5]int{1, 2, 3}`
        - 注意：没有初始化的为0
    - 自动推到类型 ：`数组名 := [n]类型{值1，值2，…. }`   。例：`a := [5]int{1, 2, 3}`
    - 指定元素初始化 ：`数组名 := [n]类型{下标1:值1，下标2:值2，…. }`  。例：`a := [5]int{3:1, 5:2}`
- 多维数组，同C语言，多个[]就是多维数组。


#### 切片
- 切片不是数组或 数组指针，而是**通过内部指针和相关属性引用数组片段**，以实现变长的方案。
- 切片并不是真正意义的动态数组，而是一个**引用类型**，slice总是指向一个底层array数组，slice的声明也可以像array一样，只是不需要长度。
- 规则：`slice := array[low:high:cap]`。
    - 其中slice取array数组的[low，high)位，并且容量为cap。
- 定义1： `slice :=[]int{}`，这样会定义一个切片，当向其中append元素时，len和cap都会增加。
- 定义2（利用make函数），make函数：`make(切片类型，长度，容量)`，例：`slice := make([]int, 5, 10)`
    - 例：`slice1 := []int{}` ,然后追加元素：`slice1 = append(slice1, 11)`
- 切片如果超过容量，通常以两倍扩容。
- 切片 s 的长度和容量可通过表达式 `len(s)` 和 `cap(s)` 来获取。
- 切片中copy函数的使用：`copy(destslice, srcslice)` ，就是将srcslice一一对应拷贝到destslice。
- 切片做函数参数，由于切片实际上是数组的引用，因此直接传入切片本身，就可以修改。



#### map映射
- map映射将键映射到值。
- map的定义1：`var mymap map[keyType]valueType` 。例：`var mymap map[int]string`
- map的定义2：可以利用make函数定义：`mymap := make(map[keyType]valueType,len)`，可以指定长度，也可以不指定长度。超过长度会自动扩容。
- 注意： map是无序的，无法决定map的返回顺序，每次打印的结果可能不同。
- map中的键是唯一的。
- 如果给map赋值时，键不存在，会自动添加并扩容。
- map的遍历，使用range，返回key，value的值，例：`for key, value := range mymap{ 代码块 }`
- 判断一个键是否存在，可以通过：`value, ok := mymap[key]`，其中，第一个value表示key对应的value值，第二个ok表示key是否存在的bool值。
- map的删除操作：`delete(map变量, map中要删除的键)`，例：`delete(mymap, 1)`
- map做函数参数： 直接将map变量名传入函数即可，在函数中可以对map进行更改操作。



#### 结构体

- 一个结构体（struct）就是一个字段的集合。定义一个结构体类型：
```
    type Vertex struct {
        X int
        Y int
    }
```
- 结构体初始化:
    - `var a Vertex = Vertex{ 结构体中的每个元素都必须进行初始化 }`
    - 结构体指定部分初始化（自动推到类型） ： `a := Vertex{ X : 1 }`, 没有赋值的默认为0
- 结构体指针： `var p *Vertex = &Vertex{ 结构体中的每个元素都必须进行初始化 }`，例如：`p := &Vertex{ X : 1 }`
- 通过new函数来申请一个结构体的内存空间： `p := new(Vertex)`
- 无论是结构体变量，还是结构体指针变量，都可以直接用（.）操作符来操作成员。
- 结构体做函数参数：
    - 值传递：直接将结构体变量名传入函数，函数内无法修改结构体内容。
    -地址传递： 将结构体的地址传入函数，在函数内才可以更改结构体的内容。
- 注意：
    - 如果想让别的包程序调用该包的函数、结构体类型、结构体成员，那么该函数名、结构体类型名、结构体成员变量名、首字母必须大写。如果首字母小写，只能在该包中使用。


## 工作区

- 工作区：
    - Go代码必须放在工作区中，工作区其实就是一个对应于特定工程的目录，它包含三个目录：分别是src目录、pkg目录和bin目录。
        - src目录：用于以代码包的形式组织并保存Go源码文件。（比如.go 、.c 、.h 、.s文件）。
        - pkg目录：用于存放经由go install命令构建安装后的代码包（包含Go库源码文件）的“.a”归档文件。
        - bin目录：与pkg目录类似，在通过go install命令完成安装后，保存由Go命令源文件生成的可执行文件。
    - 目录src用于包含所有的源代码，是Go命令规则的一个强制的规则，而pkg和bin目录则无需手动创建，如果必要Go命令行工具在构建过程中会自动创建这些目录。
    - 注意：只有当环境变量GOPATH中只包含一个工作区目录时，go install命令才会把命令源码安装到工作区的bin目录下。若环境变量GOPATH中有多个工作区的路径时，执行go install会失败，此时们必须设置环境变量GOBIN。
- GOPATH设置
    - 为了构建工程，需要把所需工程的根目录添加到环境变量GOPATH中，否则，即使处于同一工作目录下，代码之间也无法通过绝对代码包路径完成调用。
    - 当有多个工作区时，需要将这些工作区都添加到GOPATH环境变量中。



## 包 Package 

- 每个 Go 程序都是由包构成的。导入包时有几种方法：
    - `import “包名” ` : 在使用时，需要带包名来调用方法，例：fmt.Println
    - `import . “包名”` : 在使用时，无需使用包名，直接调用方法，例：Println
    - `import ( “包1”  “包2” )` : 导入多个包时使用。
    - `import 别名 “包名” ` : 给导入的包起别名，例：import io “fmt”
    - `import _ “包名”` : 忽略包名，目的是为了调用包的init函数
- main包的含义：
    - Go语言的编译程序都必须有一个名叫main的包，编译程序会试图把main包编译为二进制可执行程序，一个可执行程序有且只有一个main包。
    - main包中，一定需要包含main函数，main函数是程序的入口，没有main函数，程序无法开始执行。
    - 会使用声明main包代码所在目录的目录名作为二进制可执行程序的文件名。
- 对于多文件编程（源文件在同级目录下），需要注意：
    - 分文件编程（有多个源文件），所有源文件必须放在src目录
    - 设置GOPATH环境变量，必须包括src的根目录
    - 在同一个目录下，包名必须一样
    - 可以使用go env查看相关的环境变量
    - 同一个目录下，调用别的文件的函数，直接调用即可，无需包名引用
- go install 命令：
    - 如果GOPATH中有多个工作区，那么go install前需要设置GOBIN环境变量
    - 将GOBIN环境变量设置为：当前工作区的根目录/bin
    - 然后执行go install 命令，会将工程自动编译并生成bin 和 pkg目录。
- 对于多文件编程（源文件不在同级目录下），需要注意：
    - 在src目录下，有新的目录calc，其中src目录的文件想要调用src/calc目录下的函数
    - 不同的目录，包名不一样
    - 调用不同包里的函数，首先要import的导入，然后调用格式：包名.函数名
    - 调用别的包的函数，这个函数名字如果首字母是小写，无法调用，首字母必须大写，才能调用。
- init函数：
    - 导入一个包时，该包的init函数会首先执行。
    - 每个包可以包含任意多个init函数，这些init函数会在main函数执行前执行。
    - `import _  “包名”`      意思就是调用该包的init函数，但并不导入这些包的方法。
    




## 面向对象编程

尽管Go语言没有封装、继承、多态这些概念，但通过其他方式可以实现这些功能。
- 封装： 通过方法实现。
- 继承： 通过匿名字段实现。
- 多态： 通过接口实现。

### 匿名组合

**匿名字段**： 
- 定义结构体时，结构体内部的字段名与类型是一一对应的
- 但是Go语言支持只提供类型，而不写字段名的方式，也就是匿名字段，也成为嵌入字段。
- 对匿名字段继承来的成员的调用，直接用（.）操作符。
- 如果通过匿名字段继承来的成员与本结构体中定义的字段同名，就近原则，先在本作用域内查找成员，找不到再去外层作用域。
- 也可以使用基础类型的匿名字段，例如，在结构体中直接写：int，不加变量名。与结构体匿名字段操作类似。
```
type Person struct {
	name	string
	sex 	byte
	age 	int
}

type Student struct {
	//匿名字段，只有类型，没有名字，继承了Person的所有成员。
	Person
	id 		int
	addr	string
}

func main() {
	//注意，有匿名字段的结构体，要注意对内部结构体的初始化
	//顺序初始化，全部初始化
	var a Student = Student{Person{"mike", 'm', 19}, 1, "bj"}
	fmt.Printf("struct a: %+v\n", a)

	//自动推导类型，可以指定部分字段进行初始化。
	b := Student{id:0}
	fmt.Printf("struct b: %+v\n", b)
}
```

### 方法

**方法**   
- 在面向对象编程中，一个对象其实也就是一个简单的值或者一个变量，在这个对象中包含一些函数，这种带有接受者的函数，我们称之为方法。本质上，一个方法则是一个与特殊类型关联的函数。
- 在Go语言中，可以给任意自定义类型（包括内置类型，但不包括指针类型）添加相应的方法。
- 普通函数、匿名函数、和方法，都是函数，只是匿名函数没有函数名，而方法是普通函数绑定在某一个类型上称之为方法。
- 方法总是绑定在方法实例上，并隐式的将实例作为第一参数，方法的定义为：
    `func  ( receiver   ReceiverType )   funcName  (  参数列表  )   (  返回值  )    {  方法代码  }`

例：
```
//  重命名一个类型
type long int

//  给tmp 实例添加一个方法，也就是每个long类型都有这个方法
//  tmp 作为接受者，同时也作为函数的一个参数传入函数
func (tmp long) add(other long) long {
	return tmp + other
}

func main() {
	var a long
	var b long
	a = 10
	b = 20
    //   方法的调用，直接用（.）
	r := a.add(b)
	fmt.Println("result is : ", r)
}
```

注意： 
- 不能给指针类型添加方法：
    - 这样定义是错误的，receiver本身不能是指针
    ```
    type  pointer *int
    func   (tmp  pointer)  test() { 函数体 }      
    ```
    - 这样定义时ok的
    ```
    type   long  int
    func   (tmp *long)  test() { 函数体}     
    ```
- 这意味着对于某类型 T，接收者的类型可以用 *T 的文法
- 同理，T类型可以是结构体，**对结构体添加的方法，不能修改结构体本身，而对  \*结构体  （结构体指针）添加的方法，是可以修改结构体的内容的**。
- 使用指针接收者的原因有二：首先，方法能够修改其接收者指向的值。其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
- 而以指针为接收者的方法被调用时，接收者既能为值又能为指针。
Go语言不支持方法的重载。
- 无论是定义的\*T还是T的方法，在调用时，无论实例是值还是指针都可以调用T和*T的方法。
```
func  (tmp *T) func1 ( 参数 ) (返回值) { 函数体 }
func  (tmp T) func2( 参数 ) (返回值) { 函数体 }
在实例：var a T 和 var p *T中，都可以调这两种方法。
```


### 接口

- 在Go语言中，接口是一个自定义类型，接口类型描述了一系列方法的集合。
- 接口类型是一种抽象的类型，它不会暴露出它所代表的对象的内部值的结构和这个对象的基础操作的集合，它只会展示出自己的方法，因此**接口类型不能实例化**。
- 接口类型的定义与结构体比较类似，结构体内定义的是变量，而接口中是函数的声明。
- 接口类型的定义与实现
    - 定义： `type 接口名 interface { 方法声明 }`


下面看一个具体的例子：  
```
//  第一个结构体Student
type Student struct {
	id int
	name string
}
//   第一个结构体的方法sayhi
func (tmp *Student) sayhi() {
	fmt.Printf("Student : %d %s sayhi !\n", tmp.id, tmp.name)
}
//  第二个结构体 Teacher
type Teacher struct {
	addr string
	group string
}
//  第二个结构体的方法sayhi
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher : %s %s sayhi!\n", tmp.addr, tmp.group)
}
//  第三个自定义类型Mystring
type Mystring string
//  自定义类型的方法 sayhi
func (tmp *Mystring) sayhi() {
	fmt.Printf("Mystring: %s sayhi!\n", *tmp)
}

//  接口的定义
type Myinterface interface {
	sayhi()
}

func main() {
    //  接口的使用，不能实例化的意思是不能初始化值
	var i Myinterface
    //   注意  接口函数的接受者类型要正确
	var s *Student = &Student{1, "mike"}
	i = s
	i.sayhi()

	var t *Teacher = &Teacher{"Nanjing", "IT"}
	i = t
	i.sayhi()

	var mys Mystring = Mystring("this is string")
	i = &mys
	i .sayhi()
	}
```

**接口实现多态**   
```
//   同样还是三个结构体的定义，以及方法的定义
type Student struct {
	name string
	id int
}

func (tmp *Student) sayhi() {
fmt.Printf("Student : %s  %d  say hi!\n", tmp.name, tmp.id)
}

type Teacher struct {
	group string
	addr string
}

func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher : %s  %s  say hi!\n", tmp.group, tmp.addr)
}

type Mystring string

func (tmp *Mystring) sayhi() {
	fmt.Printf("Mystring : %s  say hi!\n", *tmp)
}

type Myinterface interface {
	sayhi()
}

//   定义了一个普通函数，以接口类型为参数，实现多态
func whosayhi(tmp Myinterface) {
	tmp.sayhi()
}

func main() {
	var s *Student = &Student{"mike", 1}
	var t *Teacher = &Teacher{"go", "Nanjing"}
	var mys Mystring = Mystring("this is a test")
	
    //   实现了一个接口whosayhi，多种形式。
	whosayhi(s)
	whosayhi(t)
	whosayhi(&mys)
	
	sli := make([]Myinterface, 3)
	sli[0] = s
	sli[1] = t
	sli[2] = &mys
	for _, i := range sli {
		i.sayhi()
	}
}

```

- 接口继承：
    - 同样接口类型也可以通过匿名字段来继承。
    - 父接口称之为子集，而子接口称之为超集，超集包含父接口，所以超集可以接收子集的类型作为接受者。
- 空接口：
    - 空接口不包括任何的方法，所有的类型都实现了空接口
    - 因此，空接口可以存储任意类型的数值，有点类似与C语言中的void *类型。
    - 空接口相当于一个万能类型，保存任意类型的值，例：
    ```
        var i interface{} = 1
        i = "abc"
    ```
    - 当函数可以接收任意的对象实例时，可以将函数参数声明为`interface{}`，一个典型例子就是fmt库中的print系列函数
    ```
    func xxx(arg …interface{}) {}
    ```
- 既然有了空接口可以传入任何类型，那么就存在函数对不同类型的处理，这就需要类型判断，也就是类型断言。使用： 
```
var i interface{} = 10
value, ok := i.(int)
// 或者
switch value := i.(type) {
    case int: 
    ...
}
```


## 异常处理

### error接口 

Go语言引入一个关于错误处理的标准模式，即error接口，它是Go语言内建的接口类型。

- error 接口的定义为：
    - `type  error  interface{  Error()  string   }`              
    - Error()  返回值为 string 
- 说明error其实也是内置的变量类型，本质是一个接口。可以通过var来定义：`var  err1   error`


### 产生error类型的几个方法

- 第一种方法是利用Go语言的标准库error，提供了New方法，产生一个error类型。
- 第二种方法是利用fmt库中的Errorf函数，来产生error类型。

例：
```
err1 := errors.New(“this is error1 ！”)
fmt.Println(“error1: “, err1)
err2 := fmt.Errorf(“this is error2 !”)
fmt.Println(“error2: “, err2)
```

### 异常处理的应用

```
func Mydiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("除数为0")
	} else {
		result = a / b
	}
	return
}

func main() {
	result, err := Mydiv(10,2)
	if err == nil {
		fmt.Println("result is : ", result)
	} else {
		fmt.Println("err is : ", err)
	}
}
```

### panic函数

当遇到不可恢复的错误时，如数组越界、空指针引用等，这些错误在运行时会引起panic异常.

在一般情况下，不应通过调用panic函数来报告普通的错误，而应该利用该函数报告致命的错误。当panic函数发生时，程序会终止运行。


### recover函数

- Go语言提供了一种专用于拦截运行时panic的内建函数：**recover函数**。
- recover函数可以将程序从当前程序运行时的panic状态中恢复并重新获得流程控制权。
- 注意： 
    - recover只有在defer调用的函数中有效。
    - 如果调用了内置函数recover，并且定义该defer语句的函数内发生了panic异常，recover会使程序从panic中恢复，并返回panic value
    - 而导致panic异常的该函数不会继续运行，但可以正常返回。
    - 在未发生panic时调用recover，recover会返回nil。例：

```
func testa() {
	fmt.Println("test a func")
}

//  testb 函数数组越界后，recover捕获panic异常，并打印错误信息，而后面执行的testc()函数正常执行
func testb(x int) {
	defer func() {
		// 设置recover函数，并且打印错误信息
		if err := recover(); err != nil {
			fmt.Println("err: ", err)
		}
	}()        //  注意要加()  表示在函数testb执行完后执行该函数
	var a [10]int
	a[x] = 111   //设置数组越界
	fmt.Println("test b func")
}

func testc() {
	fmt.Println("test c func")
}

func main() {
	testa()
	testb(20)
	testc()
}

#打印结果为：
test a func
err:  runtime error: index out of range
test c func
```


## 文本文件处理

### 字符串处理

在strings标准库中，有一些常用字符串操作函数：
- `Contains(s, substr string)` ：字符串s中包含substr，返回bool值
- `Join(a []string, sep string)` ：字符串连接，把slice a通过sep连接成一个字符串返回。
- `Index (s, sep string)` ：在字符串s中查找sep子字符串所在的位置，返回位置值，不存在返回-1。
- `Repeat(s string, count int)` ：重复s字符串count次，最后返回重复的字符串。
- `Replace(s, old, new string, n int)` ：在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换。
- `Split(s, sep string)` ：把s字符串按照 sep 分割，返回slice。
- `Trim(s string, cutset string)` ：在s字符串的头部和尾部去除cutset指定的字符串。
- `Fields(s string)` ：去除s字符串的空格符，并且按照空格分割字符串，最后返回slice


### 字符串的转换

字符串转换的函数在strconv标准库中，下面三个函数是常用的字符串转换函数。

- Append系列函数，将整数、bool值等类型转换为字符串后，在添加到现有的字节数组中。
```
    slice := make([]byte, 10)
    // 将1234以十进制方式转换为字符串后追加到slice中，并返回。
    slice = strconv.AppendInt(slice, 1234, 10)
    slice = strconv.AppendBool(slice, false)
    slice = strconv.AppendQuote(slice, “quote”)
    slice = strconv.AppendQuoteRune(slice, ‘人’)
```
- Format系列函数，把其他类型转换为字符串。
```
    a := strconv.FormatInt(1234, 10)
    b := strconv.FormatBool(false)
    c := strconv.FormatUint(1234, 10) 
    d := strconv.Itoa(1023)
```
- 字符串转其他类型，使用strconv.Parse类型。


### 正则表达式

Go语言通过regexp标准包为正则表达式提供官方支持。

例如： 
```
buf := “abc dad asd rwg ;ml ,fs”

//  MustCompile函数用来解析正则表达式，如果成功，返回一个Regexp指针类型解析器
//  函数内使用的是反引号 波浪号下面的那个，表示原生字符串
regexp.MustCompile(`a.d`)
//   然后可以使用find系列函数来
result := reg1.FindAllStringSubmatch(buf, -1)
```

### Json格式的编解码

Json格式的编解码是通过encoding/json的标准库实现的。

1. 通过结构体生成json字符串 -- Marshal函数将结构体转换成json字符串，并且返回一个字符切片和error
```
type Student struct {
	Company string
	Skill []string
	Name string
	Id int
	IsChina bool
}

func main() {
	s := Student{"oracle",[]string{"C++", "Go", "Python"}, "ggg",3,true}
	
	//   Marshal函数将结构体转换成json字符串，并且返回一个字符切片和error
  	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
}

#打印结果：
buf =  {"Company":"oracle","Skill":["C++","Go","Python"],"Name":"ggg","Id":3,"IsChina":true}
```
注意： 
还可以对输出的Json字符串做一个格式化的输出：
直接将上面的Marshal函数替换为MarshalIndent，即：`buf, err := json.MarshalIndent(s, “”, ” “)`


2. tag 的使用
   可以选择的控制字段有三种：
    - -：不要解析这个字段
    - omitempty：当字段为空（默认值）时，不要解析这个字段。比如 false、0、nil、长度为 0 的 array，map，slice，string
    - FieldName：当解析 json 的时候，使用这个名字

例如： 
```
type Student struct {
	Company string      `json:"-"`          //忽略此字段
	Skill []string 		`json:"skill"`        //  二次编码，别名skill
	Name string
	Id int				`json:",string"`     //   将int转化为string类型
	IsChina bool		`json:"local,omitempty"`          //   如果为空就不解析此字段。
}
```

3. 解析Json格式的字符串到结构体  --  使用Unmarshal函数来解析json字符串
```
	jsonBuf := `
{
 "Company": "oracle",
 "skill": [
  "C++",
  "Go",
  "Python"
 ],
 "Name": "ggg",
 "Id": 3,
 "IsChina": true
}
`
	var s Student
	//  使用Unmarshal函数来解析json字符串
	err := json.Unmarshal([]byte(jsonBuf), &s)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("s = ", s)

```


## 文件操作（os标准库）

### 新建文件：
通过两种方法新建文件：
- 根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666，返回的文件对象是可读写的。
    - `func  Create(name string)  (file  *File, err Error)`
- 根据文件描述符创建相应的文件，返回一个文件对象。
    - `func   NewFile(fd uintptr ,  name string)  *File`

### 打开文件
通过两种方法打开文件：
- 读方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。
    - `func  Open(name  string)  (file  *File,  err Error)`
- 打开名称为name的文件，flag是打开的方式，读、写等，perm 是权限。
    - `func  OpenFile(name  string,  flag  int,  perm  uint32)  (file  *File,  err Error)`

### 写文件
- 写入byte类型的信息到文件。  
    - `func   (file  *File) Write(b []byte)  (n int , err Error)`
- 在指定位置写入byte类型的信息到文件。
    - `func   (file  *File) WriteAt(b []byte,  off  int64)  (n int , err Error)`
- 写入string类型的信息到文件。
    - `func   (file  *File）WriteString(s  string)  (n int , err Error)`

### 读文件：
- 读取数据到b中。
    - `func  (file  *File)  Read(b  []byte)  (n int , err Error)`
- 在指定位置off开始读取数据到b中。
    - `func  (file  *File)  ReadAt(b  []byte,  off  int64)  (n int , err Error)`

### 删除文件：
- 删除名为name的文件。
`func  Remove(name  string)  Error`

例如： 
```
package main

import (
	"fmt"
	"io"
	"os"
)

func WriteFile(name string) {
	//  创建文件
	f, err := os.Create(name)
	if err != nil {
		fmt.Println("err = ", err)
	}
	//  关闭文件
	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("No %d, test test test!\n", i)
		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Println("n = ", n)
	}
}

func ReadFile(name string) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("err = ", err)
	}
	//  关闭文件
	defer f.Close()

	buf := make([]byte, 1024*4)
	n, err := f.Read(buf)
	//  文件出错，并且错误不是读到文件末尾时
	if err != nil && err != io.EOF {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf[0:n]))
}

func main() {
	filename := "./test.txt"
	WriteFile(filename)
	ReadFile(filename)
}

```



## 并发编程

### 并发与并行
- 并行（parallel）：指在同一时刻，有多个指令在多个处理器上同时执行。
- 并发（concurrence）：指在同一时刻，只能有一条指令执行，但多个进程指令可以被快速的轮换执行。使得在宏观上，有多个执行同时执行的效果，但在微观上，没有同时执行，只是把时间分割成若干段，使多个进程快速的交替执行。
- 其中，并发表示怎么样能够高效率的利用硬件资源，而并行表示有多少个cpu就可以并行多少个。



### Go语言并发的优势
- Go语言在语言层面上支持并发，并发的内存管理是非常复杂的，而Go语言提供了自动垃圾回收机制；
- Go语言为并发编程而内置的上层API基于CSP（顺序通信进程）模型，这就意味着显式锁都是可以避免的，因为Go语言通过安全的二通道发送和接收数据以实现同步，这大大简化了并发编程的编写。


### goroutine
- goroutine是Go语言并发设计的核心，也就是协程的意思，它比线程更小，十几个goroutine可能底层也就五六个线程，Go语言内部实现了这些goroutine的内存共享，goroutine比thread更易用，更高效，更轻便。
- 创建goroutine非常简单：  
    - 只需要在函数调用语句前添加go关键字即可创建并发执行单元，调度器会自动将其安排在合适的系统现呈上执行。

例如：
```
func NewTask() {
	for {
		fmt.Println("this is NewTask")
		time.Sleep(time.Second)
	}
}

func main() {
    //  开启子协程
	go NewTask()
    //  开始运行后，子协程和main一起争夺cpu资源
	for {
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second)
	}
}
```

**注意**   
- 当主函数main函数退出时，所有的子协程也会退出。
- 当主协程main退出时，有时候会导致子协程未来得及调用，这点需要注意。

### runtime包
1. Gosched
    - runtime.Gosched()用于让出时间片，让出当前goroutine的执行权限，调度器安排其他等待的任务执行，并在下次某个时刻，从该位置恢复执行。例：
    ```
    func main() {
    
    	go func() {
    		for i := 0; i < 2; i++ {
    			fmt.Println("Hello Newtask!")
    		}
    	}()
    
    	for i := 0; i < 2; i++ {
    		// 让出时间片，先让别的协程执行，它执行完，再回来执行此协程
    		//   如果不加runtime.Gosched(), 会导致主协程执行完了，子协程还没来得及执行
    		runtime.Gosched()
    		fmt.Println("Hello main!")
    	}
    }
    ```
2. Goexit
    - 调用runtime.Goexit()将立即终止当前goroutine执行，调度器将保证所有已注册的defer延迟调用被执行。例：
    ```
    func test() {
    	fmt.Println("子协程中调用函数test()开始====》")
    	defer fmt.Println("defer结束test后调用===")
    	//  终止Goexit() 所在的协程，但是保证所有已注册的defer能够执行
    	runtime.Goexit()
    }
    
    func main() {
    	//  创建新的协程
    	go func() {
    		fmt.Println("子协程开始--->")
    		defer fmt.Println("defer子协程结束后执行---")
    		test()
    		fmt.Println("子协程结束--->")
    	}()
    
    	//  设置死循环，保证主协程不退出
    	for {
    	}
    }
    
    #如果没有Goexit() 函数，执行结果正常：
    子协程开始--->
    子协程中调用函数test()开始====》
    defer结束test后调用===
    子协程结束--->
    defer子协程结束后执行---
    
    #如果添加Goexit()函数，执行结果为：（所有注册了的defer都执行了）
    子协程开始--->
    子协程中调用函数test()开始====》
    defer结束test后调用===
    defer子协程结束后执行---
    
    ```
3. GOMAXPROCS
    - 调用runtime.GOMAXPROCS（）用来设置可以 并行计算的CPU核数的最大值，并返回之前的值。



### sync方式同步

**使用方法：**
```
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时操作
	wg.Done()
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download("a.com/" + string(i+'0'))
	}
	wg.Wait()
	fmt.Println("Done!")
}
```

注意：
- `wg.Add(1)`：为 wg 添加一个计数，wg.Done()，减去一个计数。
- `go download()`：启动新的协程并发执行 download 函数。
- `wg.Wait()`：等待所有的协程执行结束。





### 并发的资源竞争问题

竞争问题抛出：
```
//  打印每个字符，并sleep一秒
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
}

func main() {
	//  两个协程同时开始调用Printer函数，一起竞争CPU的资源，所以打印结果是乱的
	go func() {
		Printer("HELLOWORLD")
	}()

	go func() {
		Printer("helloworld")
	}()

	for {

	}
}

#可以看到打印结果是乱的：
hHEelLLlOowWOoRrLlD
```


### channel

**goroutine运行在相同的地址空间，因此访问内存需要同步，goroutine通过通信来共享内存，而不是共享内存来通信。**

**channel类型**是CSP模式的具体实现，用于多个goroutine通讯，其内部实现了同步，以确保并发安全。


####  channel类型
- 和map类似，channel也是一个对应make创建的底层数据结构的引用。
- 当我们复制了一个channel或用于参数传递时，可以修改其底层的数据，类似于其他引用。和其他类型一样，channel的零值也是nil。
- channel类型的定义：
    - `make(chan  Type)` : 等价于`make(chan  Type,  0)`
    - `make(chan  Type,  capacity)`
    - 当capacity=0时，channel时无缓冲，通过阻塞来读写，阻塞就是当channel中无数据时，阻塞等待，一直到有数据为止。
    - 当channel>0时，是非阻塞的，直到写满capacity个元素后才阻塞写入。
- channel通过操作符 `<-` 来接收和发送数据
    - 发送给channel： `channel  <-  value`
    - 从channel 中取： `x := <-  channel` 
    - 可以使用 range 来读取数据： `for num := range ch {}`

例如： 
```
//    定义一个全局channel，并且capacity为0，表示阻塞等待。
var ch = make(chan int)
func main() {
	//  两个协程同时开始调用Printer函数，一起竞争CPU的资源，所以打印结果是乱的
	go func() {
		Printer("HELLOWORLD")
		//   当子协程1 打印完成后，向管道内写入数据
		ch <- 1
	}()

	go func() {
		//    当子协程2从管道读取到数据时（丢弃）， 才会往下执行，否则阻塞等待
		<- ch
		Printer("helloworld")

	}()

	for {
	}
}	


#打印结果就正常了：
HELLOWORLD
helloworld
```


#### 无缓冲的channe 

**无缓冲的通道 unbuffered channel**是指在接收前没有能力保存任何值得通道。

- 这种类型要求发送goroutine和接收goroutine同时准备好
- 通道会导致先执行发送或接收操作的goroutine阻塞等待，即发送到channel后阻塞等待，等待接收完通道的数据后，才继续运行。
- 任意一个channel的发送或接收操作都无法离开另一个操作单独存在。
- 创建形式，就是使用make创建，无capacity形式的管道，即：`make(chan  Type)` , 等价于`make(chan  Type,  0)`

#### 有缓冲的channel  

**有缓冲的的通道 buffered channel**是一种在被接收前能存储一个或多个值的通道。

- 这种通道并不强制要求goroutine之间必须同时完成发送和接收。
- 只有当通道中没有接收值时，接收才会阻塞。
- 只有当通道中可用缓冲区被塞满后，发送才会阻塞。

#### 关闭channel  

直接使用close( ch )就可以关闭channel了。

注意：  
- channel不像文件那样需要经常关闭，只有当你确实没有任何数据发送或接收了，或者你想显式的结束range循环之类的操作，才去关闭channel。
- 关闭channel后，无法向channel中发送数据，但是可以从channel中接收读取数据。


#### 单向的channel

默认情况下，通道是双向的，即既可以接收数据也可以发送数据。

- 有时我们希望通道是单方向使用的，可以指定通道的方向。
    - 其中，chan <-   表示发送数据，向channel中写数据；
    - <- chan   表示接收数据，从channel中读数据。
    例如： 
```
var  ch1  chan int  （这个就是正常双向的channel）
var  ch2  chan  <- int        （这个是单向的channel，用于发送数据）
var  ch3  <-  chan int        （这个是单向的channel，用于接收数据）
```

生产者消费者模型： 
```
package main

import (
	"fmt"
	"time"
)
// 将双向channel转化为单向，只能写不能读
func producer(in chan <- int)  {
	for i := 0; i < 10; i++ {
		in <- i * i
	}
	close(in)
}
// 将双向channel转化为单向，只能读不能写
func consumer(out <- chan int) {
	for num := range out {
		fmt.Println("num = ", num)
	}
}

func main() {
	//  创建一个双向channel
	ch := make(chan int)

	//   生产者，生产数字，向channel写数据，并且新开一个协程
	//   channel 参数传递，直接传，即引用传递
	go producer(ch)

	//    消费者，购买数字，从channel中读取数字并打印，同样新开一个协程
	go consumer(ch)
	//   main函数sleep1秒，保证子协程都能结束
	time.Sleep(time.Second)
}

```

#### 定时器

有两种定时器，分别是Timer和Ticker。

1. Timer 是一个定时器，代表未来的单一事件，你要告诉Timer你要等待多长时间，它提供一个channel，在将来的该事件为该channel提供一个时间值。
    - Timer的定义： `timer :=  time.NewTimer( 时间值 )`
    -  timer定时器可以停止和重置，可以使用timer.Stop()停止，使用timer.Reset()重置。
    - 注意：Timer定时器定时后，在未来某个时间里，只触发一次。
    - 通过Timer可以实现延时功能：
        - 可以通过<-timer.C，也可以通过<-timer.After(时间值)
        -  `timer :=  time.NewTimer( 时间值 )` ，其中 timer.C 表示将来事件的 channel， 并且到时间后会写入一个时间值。
        - `time.After( 时间 )` 会返回一个 channel 同 timer.C 类似。
        - timer.C 和 time.After 会一直阻塞。直到定时时间到，向channel写入值一个当前时间的事件，阻塞解除，可以从中读取数据。
```
func main() {
  	//   新建一个定时器
	timer := time.NewTimer(time.Second * 3)
	//    在定时时间内timer.C一直处于阻塞，直到定时结束向通道写入一个当前时间的事件
    num := <- timer.C
	fmt.Println("num = ", num)
	fmt.Println("时间到")
	
    //   定时结束前，阻塞，定时结束后，time.After 返回一个通道，并且向通道写入一个当前时间的事件
	num = <- time.After(time.Second * 3)
	fmt.Println("num = ", num)
	fmt.Println("时间到")
}


#结果为：
num =  2018-12-27 16:56:46.850665 +0800 CST m=+3.011172201
时间到
num =  2018-12-27 16:56:49.9658432 +0800 CST m=+6.126350401
时间到
```
2. Ticker是一个定时触发的计时器，它会以一个间隔往channel中发送一个当前时间的事件，而channel的接受者可以以固定时间向channel中读取事件。
    - Ticker的定义：`ticker  :=  time.NewTicker( 时间间隔 )`

```
func main() {
	//   新建一个ticker定时器，每隔两秒写入一个事件
	ticker := time.NewTicker(2 * time.Second)
	i := 0
	for {
		//  阻塞等待，每隔2秒执行后面的代码
		<- ticker.C
		i++
		fmt.Println("i = ", i)
		if i == 10 {
			ticker.Stop()
			break
		}
	}
}
```

#### select

**select的作用**：通过select监听channel上的数据流动。

- select的用法与switch非常类似，有select开始一个选择块，每个条件由case语句来描述。
- select也有比较多的限制。最大的限制就是每个case必须是一个IO操作。
- 在一个select语句中，Go语言会从头至尾评估每一个发送或接收的语句，如果其中的任意一个语句可以继续执行（即不阻塞），并且有多个语句可以执行，那么就从这些可以执行的语句中任意选择一个执行。
- 如果没有一个语句可以执行（即所有通道都阻塞），那么有两种情况：
    1. 如果有default语句，那么就会执行default语句，同时会从select后的语句开始继续执行。
    2. 如果没有default语句，那么select将会阻塞，直到有任意一个case后的条件可以执行。

```
select {
case <- chan1:
    //  如果chan1成功读到数据，则进行该case的处理语句
case chan2 <- 1:
    //  如果chan2成功写入数据，则进行该case的处理语句
default: 
    //  如果上面都没有成功，则进入default处理语句
}
```

使用 select 完成 斐波那契数列的例子：
```
package main

import (
	"fmt"
)

//   ch通道转化为只写，quit通道转化为只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		//  利用select监听数据的流动情况
		select {
		//   如果ch通道可以写
		case ch <- x:
			x, y = y, x+y
		//    如果quit通道可以读
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}
func main() {
	// 新建一个channel，用来数字通信
	ch := make(chan int)
	//  新建一个channel，用来确定程序是否结束
	quit := make(chan bool)

	//  新建协程，作为消费者，从channel中读取斐波那契数列
	go func() {
		for i := 0; i < 10; i++ {
			num := <- ch
			fmt.Println(num)
		}
		//   十次循环后，停止
		quit <- true
	}()

	//  生产者，产生斐波那契数列，并写入channel
	fibonacci(ch, quit)
}

#程序最终打印结果：
1
1
2
3
5
8
13
21
34
55
flag =  true
```


使用 select 完成 超时 的程序：
```
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			//	当ch阻塞后，time.After开始倒计时，时间到后程序超时
			case <-time.After(time.Second * 5):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<- quit
	fmt.Println("程序结束")
}
```


## 单元测试

如果要写单元测试代码，需要将文件名定义为 `_test.go` 结尾的形式。使用golang中的 `testing` package 来实现单元测试的逻辑，例如：

```
// calc_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := add(1, 2); ans != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}
}
```

运行 go test，将自动运行当前 package 下的所有测试用例，如果需要查看详细的信息，可以添加-v参数。



## 网络编程

网络编程主要是了解TCP/IP协议，这里就不再赘述了，Go语言关于Socket编程的标准库为net库，下面是一个socket编程的流程图：

![网络变成逻辑图](https://github.com/Nevermore12321/LeetCode/blob/blog/golang/golang%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B.png?raw=true)


**下面是一个并发C/S模型的例子：**

1. **服务器端模型为：**
```
package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	//   处理完成后，遇到exit后关闭连接
	defer conn.Close()

	//  获得连接请求的IP地址
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect successful")

	//  读取客户端发来的数据，并将数据转换成大写
	//   遇到exit断开连接
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		//fmt.Println("buf len = ", len(string(buf[:n])))
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("[%s]: %s ", addr, string(buf[:n]))
		//  通过发送过来的字符串中，多了两个字符“\n\r”
		if string(buf[:n-2]) == "exit" {
			fmt.Println(addr, "exit")
			return
		}
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	//  创建套接字，监听
	listener, err1 := net.Listen("tcp", "127.0.0.1:6443")
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}
	defer listener.Close()
	//   阻塞等待用户连接
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("err2 = ", err2)
			return
		}
		go HandleConn(conn)
	}
}
```

2. **客户端模型为：**
```
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err1 := net.Dial("tcp", "127.0.0.1:6443")
	if err1 != nil {
		fmt.Println("net.Dial err1 = ", err1)
		return
	}
	//  main函数完成后，关闭连接
	defer conn.Close()


	go func() {
		//  从键盘输入内容，给服务器发送数据
		str := make([]byte, 1024)
		for {
			//  读取键盘额输入
			n, err3 := os.Stdin.Read(str[:])
			if err3 != nil {
				fmt.Println("os.Stdin.Read err3 = ", err3)
				return
			}
			_, err4 := conn.Write(str[:n])
			if err4 != nil {
				fmt.Println("os.Stdin.Read err3 = ", err4)
				return
			}
		}
	}()
	//接收服务器回复的数据
	//  切片缓冲
	buf := make([]byte, 1024)
	for {
		//  接收服务器发来的回复数据
		n, err2 := conn.Read(buf)
		if err2 != nil {
			fmt.Println("Read err2 = ", err2)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

```

--------

下面是一个socket编程的传输文件的例子，接收方相当于服务器端，发送方相当于客户端，其过程为：
- 服务器端阻塞等待客户端连接
- 客户端发起连接，连接成功后向服务器端发送传输的文件名
- 服务器端接收到文件名后，向客户端回复一个ok，表示可以接收文件
- 客户端接收到ok指令后，开始发送文件
- 服务器端开始接受文件，完成后断开连接。



具体模型如下：

1. **接收端模型为：**
```
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFile(filename string, conn net.Conn) {
	//  新建文件
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}
	defer f.Close()

	//  接收文件
	buf := make([]byte, 1024 * 4)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完成")
			} else {
				fmt.Println("conn.Read err = ", err)
			}
			return
		}
		f.Write(buf[:n])
	}

}

func main(){
    //   创建套接字，监听
	listener, err := net.Listen("tcp", "127.0.0.1:6543")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	//   阻塞等待客户端连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err1 = ", err1)
		return
	}
	defer conn.Close()
    //   连接建立后，开始接受文件名
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Read err2 = ", err2)
		return
	}
	filename := string(buf[:n])
    //  向客户端回复一个ok，表示可以接收文件
	n, err3 := conn.Write([]byte("ok"))
	if err3 != nil {
		fmt.Println("conn.Write err2 = ", err3)
		return
	}
	// 接收文件
	RecvFile(filename, conn)
}
```

2. **发送端模型为：**
```
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
    //  打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err", err)
		return
	}
	defer f.Close()
    //  发送文件
	buf := make([]byte, 1024 * 4)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err", err)
			}
			return
		}
		conn.Write(buf[:n])
	}
}


func main() {
    //   输入传输的文件，包括路径
	fmt.Println("请输入要传输的文件（包含路径）：")
	var path string
	fmt.Scan(&path)
	//    检查文件，并且通过stat获取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err = ", err)
		return
	}
    //   向接收端发起连接请求
	conn, err1 := net.Dial("tcp", "127.0.0.1:6543")
	if err1 != nil {
		fmt.Println("net.Dial err1 = ", err1)
		return
	}
	defer conn.Close()
	//   连接建立后，发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn.Write err = ", err)
		return
	}
	//    查看接收端是否回复ok
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Write err2 = ", err2)
		return
	}
    //   如果是ok，开始发送文件
	if string(buf[:n]) == "ok" {
		SendFile(path, conn)
	}


}

```



---------------------------

下面是一个并发聊天室服务器端的例子，其中一些功能包括：
- 并发处理用户登录
- 用户发送消息广播给所有已登录的用户
- 用户长时间不操作的超时处理
- 实现who查看所有已登录用户
- 实现rename更改当前用户的名字



具体的并发聊天室服务器模型为：

1. **服务器模型为：**
```
package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	Name string       //  保存用户名
	C chan string     //   用于用户发送数据的管道
	Addr string       //    IP地址
}

//  全局的map，用于保存所有的在线用户
var onlineMap = make(map[string]Client)

//   全局channel，用于转发给所有用户消息
var message = make(chan string)

//   收到消息，广播给每个用户
func Manager() {
	for {
		//   从全局channel  message中读取数据，如果没有数据，阻塞等待
		msg := <- message
		//  如果有数据，遍历map，给所有用户发送消息，即把数据通过用户的channel发送
		for _, cli := range onlineMap {
			cli.C <- msg
		}

	}
}

//   处理用户的连接
func HandConn(conn net.Conn) {
	defer conn.Close()
	//  获取客户端的IP地址
	cliAddr := conn.RemoteAddr().String()
	//   给发来连接请求的用户新建结构体，并保存在map中
	cli := Client{cliAddr, make(chan string), cliAddr}
	onlineMap[cliAddr] = cli
	//   新开协程，专门用于给客户端发送数据
	//   从用户结构体中的channel C中读取从Manager协程中发送的数据，并且通过套接字发送给其他用户
	go WriteMsgToClient(cli, conn)

	//   发送数据的协程开启后，开始广播给所有用户新用户的上线，只要把login信息发送给全局message channel即可
	message <- "[" + cli.Addr + "]" + cli.Name + ":  login"

	// 提示当前用户的IP
	cli.C <- "my IP and Port: "

	//  如果是客户端主动退出
	isQuit := make(chan bool)

	//   检查客户端是否有数据发送过来，如果在一定时间内没有发送，则认为超时，断开连接
	hasData := make(chan bool)

	//  新建一个协程，用来处理用户发送过来的信息 （匿名函数）
	go func() {
		buf := make([]byte, 2048)
		//  循环读取客户端发来的信息
		for {
			n, err := conn.Read(buf)
			//   有两种err的情况，一种是对方断开连接 或者 read 出错
			if n == 0 {
				// 用户主动退出
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}
			msg := string(buf[:n])
			//fmt.Println("len(msg) = ", len(msg))
			//  如果收到 who，查看当前所有在线的用户的IP
			if len(msg) == 5 && msg[:3] == "who" {
				//   遍历map，并且给当前用户发送所有的成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					user := tmp.Addr + " : " + tmp.Name
					conn.Write([]byte(user))
				}
			} else if len(msg) > 7 && msg[:7] == "rename "{
				// 输入rename 修改当前用户的名字，rename+空格+新名字
				//  split按空格把字符串分开，取第二个元素即为新名字
				name := strings.Split(msg, " ")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename OK"))
			}else {
				//  把收到的信息，通过channel  message给到广播消息的协程Manger
				message <- "[" + cli.Addr + "]" + cli.Name + " : " + msg
			}
			//  只要 收到的n不为0 ，表示有数据
			hasData <- true
		}
	}()


	//  为了不让WritemsgToClient协程不退出，因此，当前协程不能退出
	for {
		//使用select检测channel  isquit是否有数据
		select {
		case <-isQuit:
			//  如果isQuit中有true，则执行删除当前用户的信息
			delete(onlineMap, cliAddr)
			//  广播下线通知
			message <- "[" + cli.Addr + "]" + cli.Name + ":  log out"

		//   如果有数据，不做任何处理
		case <-hasData:

		//    如果超过60秒没有收到数据，则超时处理，断开连接
		case <- time.After(time.Second * 60):
			delete(onlineMap, cliAddr)
			message <- "[" + cli.Addr + "]" + cli.Name + ":  Time Out To Leave"
			return
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	//  遍历用户channel中的数据，给其他客户端发送
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}


func main() {
	//  创建套接字，并监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listener.Close()

	//  新开一个协程，用于转发消息，只要消息来了，遍历map，给每个用户发送消息
	go Manager()

	//  主协程，循环阻塞等待用户连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		//  新建协程，处理用于发来的连接请求
		go HandConn(conn)
	}
}
```