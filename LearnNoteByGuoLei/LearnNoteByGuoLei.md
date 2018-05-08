# 4月份 go 语言学习笔记

## 前言

4月份，学习了 go 语言，在学习过程中遇到的比较有意思或需要注意的地方在此记录。

> 注：本文作为个人笔记，不包含 go 语言基础只是的详细解释，如果有基础方面的问题直接咨询本人。

## go 语言中的包

#### 声明程序所属包时需要注意以下几点：

* go 语言的包使用 package 关键字声明，并且声明的时候，不包所在的路径，只是一个单纯的包名。

* 使用 go 语言编写的程序都需要放到包中，每个文件夹中只允许存在一个包。

* 在两个不同的文件夹下，即使声明的包名相同，也算是不同的包。

#### 包管理工具

由于 go 语言在编译的时候包之间存在依赖关系，但是 go 不能根据程序逻辑来判定包之间的这些依赖，所以在编译的时候需要人工选择先编译哪个包，后编译哪个包！这就需要开发人员对包之间的依赖有一个详细的了解，在开发过程中，有可能出现新人介入等问题，因此要求所有开发人员对一个比较大的工程所有包都了解这一点是很难做到的，所以为了管理包之间的依赖关系，我们需要一个包管理工具。
go 语言并没有提供官方的包管理工具，所以现在常用的都是第三方包管理工具：

* godep

* vender

* gb

* gpm

* gvp

> 注：时间原因，暂时还没有对这些第三方的包管理工具的使用进行学习，具体内容以后补全。
    
#### main 包

* go 语言在编译的时候，main 包会根据系统不同生成不同后缀的可执行文件， main 之外的包会生成 .a 后缀的文件。

* main 包中必须有一个命名为 main 的主函数，作为程序运行时的主逻辑。

#### 特殊函数

* main 函数，程序主函数。

* init 函数：如果在程序运行之前，需要一些初始化的工作，则需要用到 init 函数：

* 一个包中可以包含多个 init 函数。

* init 函数会在 main 函数之前运行。

* 所有在 main 包中引用的其他包的 init 函数也会在 main 函数之前运行。
            
> go 语言没有一个由官方维护的类似于 npm 或者 maven 的包管理工具，所以，它针对于第三方包的管理比较杂乱，在选用第三方包的时候，应该谨慎！个人感觉 go 语言的基础包相对于其他语言的基础功能要强大，完全可以基于基础包来针对需要的功能自己来做封装。

## go 语言的类型

个人理解的 go 语言的类型分为三大类：基础类型，集合类型，特殊类型
    
#### 基础类型

* 布尔型： 不多解释

* 数字类型： go 语言的数字类型包括整形和浮点类型

 * 对于整形，可以根据长度，按字节分配。go 语言提供了 8位， 16位， 32位 和 64位 整形对应的关键字

 * 对于浮点类型，也可以根据长度，按字节分配。 go 语言提供了 32位， 64位 浮点类型对应的关键字

* 字符串类型：不多解释

* 函数类型：golang 支持面向函数编程

* 接口类型：后面详细解释

* 自定义结构体类型：个人将其当作 java 中的 bean 对象来理解

 在某资料上查到，字符串类型严格意义上来讲也是引用类型，但是在 string 类型值作为函数参数传递的时候也是使用的副本，个人对这部分没有深究，只需知道其使用方式与基础类型相同即可.
        
#### 集合类型

* 数组类型
 数组类型是一个定长，类型固定的值的集合。在使用上和其他语言没有太大的区别，数组在内存单元中的存储方式如下图：![数组内存示意图](https://raw.githubusercontent.com/yunkuhui/learning-go/guolei/LearnNoteByGuoLei/array_struct.jpg)

* 切片类型
 切片类型是基于数组而创建的，唯一的区别是切片的长度是可变的，其内存结构如下图：![切片内存示意图](https://raw.githubusercontent.com/yunkuhui/learning-go/guolei/LearnNoteByGuoLei/slice_struct.jpg)
切片长度是指数组中有值的内存地址的个数，切片的容量是指切片指向的数组起始位置到数组终止位置的长度。

* map类型
 go 语言为了映射的处理效率，其底层使用 hash 来实现的，内存结构如下：![映射内存示意图](https://raw.githubusercontent.com/yunkuhui/learning-go/guolei/LearnNoteByGuoLei/map_struct.jpg)
如上图所示在映射中键值对储存的时候，会先将键值对的键转换为散列值，使用散列值的低位来作为选择桶的散列键。
在上图中也可看出每个桶的内存存储方式，也就是映射最终的内存存储方式：每个映射由两个数组组成，第一个数组为存储散列键高八位的数组，第二个数组会将键和值分开，并存储到连续的内存单元。
当操作映射的时候，会首先根据键值对应的散列值低位来选择桶，然后根据散列键的高位来确定对应键值在内存中的索引，最终取得相应的键值对。

#### 特殊类型

* 指针类型：指针类型的值是变量在内存中的地址，通过指针类型可以直接操作变量本身的值，可以通过这种方式实现变量的“引用传递”。

* channel 类型：通道类型，用于 goroutine 之间的数据传递，具体使用方式在 goroutine 章节会具体介绍。

#### 小例子

 对于集合类型中的数组和切片，下面有一个比较有趣的小例子，这个例子代码量虽然不多，但是可以比较全面的体现 slice 和数组的内存关系，以及 slice 中长度和容量的关系：

	array := []int{10, 20, 30, 40, 50}
	slice := array[1:3]           // slice的值为[20, 30]
	slice[0] = 60                 // 将slice索引为0的值设置为60
	slice[1] = 70                 // 将slice索引为0的值设置为70
	slice = append(slice, 80)     // 给slice切片追加一个80的值
	slice = append(slice, 90)     // 给slice切片追加一个90的值
	fmt.Println(slice)            // 第一次打印slice切片
	fmt.Println(array)            // 第一次打印array数组
	slice = append(slice, 100)    // 给slice切片追加一个100的值
	slice[0] = 110                // 将slice索引为0的值设置为110
	fmt.Println(slice)            // 第二次打印slice切片
	fmt.Println(array)            // 第二次打印array数组


## 方法

#### 方法的概念

有接收者的函数就是方法，方法的接收者分为【值接收者】和【指针接收者】两种，方法的接收者不仅仅声明了方法的所属结构，同时也可以作为一个参数，在方法体中使用。

#### 不同方法接收者的区别

* 值接收者：和方法的参数一样，在调用方法的时候使用的是一个值的副本，当方法体中对接收者进行操作的时候，接收者原值不会改变。

* 指针接收者：类似于引用传递，当在方法体内对接收者进行操作的时候，接收者原值也会跟着改变。

如果一个方法在声明的时候指定了接收者为某个结构的值，则该方法可以同时用于方法的值和指针，反之亦然。例：
 
	// 声明结构体TestStruct1
	type TestStruct1 struct{}
	// 声明方法testFunction1，并且设置方法的接收者为TestStruct1的值
	func (receiver TestStruct1) testFunction1 () {}
	// 声明方法testFunction2，并且设置方法的接收者为TestStruct1的指针
	func (receiver *TestStruct1) testFunction2 () {}
	// 初始化一个TestStruct1值对象
	struct1Value := TestStruct1{}
	// 初始化一个TestStruct1指针对象
	struct1Pointer := TestStruct1{}
	// 正确
	struct1Value.testFunction1()
	// 正确， golang会自动将【struct1Pointer】转换为【*struct1Pointer】
	struct1Pointer.testFunction1()
	// 正确， golang会自动将【struct1Value】转换为【&struct1Value】
	struct1Value.testFunction2()
	// 正确
	struct1Value.testFunction3()

#### 注意：并不是所有的对象都可以获取到相应的内存地址

 例：
 
	// 基于int类型声明一个Integer类型
	type Integer int
		
	func (d * Integer) print() string {
		// 声明一个方法，其接收者为Integer的指针
		return fmt.Sprintf("Integer: %d", *d)
	}
		
	// 编译可以通过，但是运行时报错
	Integer(42).print()

个人理解， golang 获取不到没有句柄的对象的内存地址（如有错误欢迎指正），由于 golang 的编译器检查不出此种错误，所以一定要注意！！！

## 接口

#### 简介

go 语言的接口类似于 java 中的接口，是对应一类对象所具有的特定行为进行描述的类型。
与 java 不同的是， java 需要显示的说明对象具体实现了哪个接口， golang 中只要是对象实现了接口的方法，就说明该对象实现了接口。

#### 内存结构

接口类型的内存结构如下所示：![接口类型内存结构](https://raw.githubusercontent.com/yunkuhui/learning-go/guolei/LearnNoteByGuoLei/interface_struct.jpg)

由上如可以看出：接口类型的 itable 中存放的是实现接口的实例结构体的特征， value 中存放的是实例结构体的对象，当一个实例给接口类型的对象赋值时，会尝试将类型和其方法集套入接口类型的 itable 中，如果符合方法集的限制条件，则将实例的内存地址存入 value 中。
接口方法集的限制条件：当使用实现接口的结构体的值给接口类型赋值时，接口类型的方法集只能放置接收者为值接收者的实现接口方法，当只用实现类的指针给接口类型赋值时，接口类型的方法集可以使用【值接收者】或【指针接收者】实现的接口方法。在开发规则的角度来看的话：当一个结构体实现一个接口的方法的接收者为【指针接收者】，则只有该结构体的指针可以赋值给该接口类型；当一个结构体实现一个接口的方法的接收者为【值接收者】，则该结构体的指针和值都可以赋值给该接口类型。（ go 语言做这种限制的原因是请参照：[方法章节的注意事项](#注意并不是所有的对象都可以获取到相应的内存地址)）
例：

	type TestInterface interface{
		testFunction()
	}
	
	type TestStruct1 struct{
		column string
	}
	
	// TestStruct1使用值接收者实现了TestInterface接口
	func (receiver TestStruct1) testFunction () {
		……
	}
	
	type TestStruct2 struct{
		column string
	}
	
	// TestStruct2使用指针接收者实现了TestInterface接口
	func (receiver *TestStruct2) testFunction () {
		……
	}
	
	// 初始化一个TestInterface类型的对象
	var interface TestInterface
	
	// 初始化一个TestStruct1类型的对象
	var struct1 TestStruct1 = TestStruct1{
		column:"test"
	}
	
	// 初始化一个TestStruct2类型的对象
	var struct2 TestStruct2 = TestStruct1{
		column:"test"
	}
	
	// 正确
	interface = struct1
	
	// 正确
	interface = &struct1
	
	// 编译错误
	interface = struct2
	
	// 正确
	interface = &struct2


实现了接口的结构体对象可以赋值给相应的接口类型，以此来统一处理一类对象。

## goroutine

#### 简介

个人认为 golang 运行效率高的原因是如下几点：

* 编译型系统语言，避免了 JVM 中间层的性能消耗

* goroutine 更简单高效的使用系统资源（此种性能的提升是建立在合理使用 goroutine 和良好编码的基础上）

个人对于 goroutine 的理解是： goroutine 是一组建立在 CPU 线程上的一个逻辑处理器，每个逻辑处理器上可以运行多个 goroutine ，当有多个逻辑处理器时，所有的 goroutine 会平均分布到所有的逻辑处理器。

> golang 可以通过 runtime 包的 NumCPU() 函数查看当前计算机的线程数。
 golang 也可以通过 runtime 包的 GOMAXPROCS() 函数来设置工程允许使用的逻辑处理器个数。

#### 高并发时的竞争

goroutine 之间存在竞争问题，当多个 goroutine 同时操作一个对象时，就会出现竞争，竞争时有可能会导致数据丢失。

> go 工程可以在build时加上 **-race** 参数来检查程序是否存在竞争，当加上 **-race** 时，如果程序存在竞争的情况则编译失败

#### 阻止竞争的方法

* 原子函数
 针对于数字类型， golang 设置了相应的原子函数 **atomic** ，通过原子函数可以实现 goroutine 之间的同步。

* 互斥锁
 golang 在 sync 基础包中提供了 **mutex** ，通过 **mutex** 方法可以锁住某个代码片段，使其在 goroutine 之间同步执行。

* 通道（channel）
 golang 通道的有阻塞的机制，当通道只有输入或者只有输出的时候，会阻塞 goroutine 的运行，因此，可以通过 channel 的阻塞来控制 goroutine 之间的同步问题。

原子函数的使用只能针对于数字类型，有一定的限制，而互斥锁是直接锁定代码片段，不灵活，所以绝大多数都会使用 channel 的阻塞机制来实现 goroutine 的同步，避免竞争的问题。

#### 应用

关于 goroutine 和 channel 的协作的应用（个人想到的应用，不全面，以后想到随时追加）：

* 通过 goroutine 和无缓冲 channel 控制程序的生命周期，例如：一个 web 服务在启动的时候同时启动一个 goroutine ，在 goroutine 中使用 channel 来接收系统的终止信号（ ctrl+c 或者 kill 命令等）。当 channel 没有接收到程序的终止信号时，当前 channel 是阻塞的，导致当前 goroutine 暂时停止运行，只有当接受到终止信号时， channel 停止阻塞，执行 goroutine 后续的服务关闭逻辑。过此方法可以控制一个 web 服务的生命周期

* 通过 goroutine 和有缓冲 channel 实现可循环利用资源的管理池，例如：可以创建多个DB连接放到有缓冲 channel 中；在需要使用DB连接的时候直接在 channel 中取出DB连接对象，使用之后放回 channel 通道中；根据 channel 的阻塞特性，当 channel 中没有可用的DB连接对象时，会阻塞调用DB连接的程序，知道其他 goroutine 使用完 DB 连接并将其返还给 channel 之后，才会重新激活之前阻塞的 goroutine ，这一点很符合资源管理池的特性。

* 通过 goroutine 和无缓冲通道即时处理的特性，创建线程管理池，实现可控线程数的批量处理。

#### select 关键字

通过 select 关键字可以实现类似于 switch 的选择语句，它会实时的接收 channel 中传递的对象，如果 channel 中没有对象传入的时候则阻塞。（有 default 分支的话就会执行 default 分支而不会阻塞）

## 其他小细节（此部分为个人笔记，随时补充）

#### new 和 make 的区别

* new 可以适用所有类型，使用new关键字返回的是对应类型的指针。

* make 只适用于 array，alice map，channel，返回的是对象的值。

> 注意： map 类型在做初始化的时候不能单纯的声明为 nil ，必须使用 make 函数对其分配内存。

#### 结构体的定义方式

* 直接定义：type XXXXXX Type
例：

	type Integer int    // 根据int类型创建一个Integer类型

* 普通定义：type XXXXXX struct{}
例：

	type test1 struct {
		value1 int
		value2 string
	}

* 类型嵌套：type XXXXXX struct {}
例：

	type test1 struct {
		struct1
	}

#### 值初始化方式

在 go 中，通常使用 **var** 关键字做没有赋值的初始化，使用 **:=** 做有赋值的初始化。
例：

	var slice []string
	array := test.NewArray()

这么做的原因是使用 **var** 关键字必须给变量指定类型，而有些类型的作用于只可用于其包内，而在包外是不能使用的；而 **:=** 可以省略值的类型，不管其类型是私有的还是公开的， go 都会自动识别其类型。