### 算法与数据结构

- #### 快排，堆排，最小堆，最大堆实现

    ```
    func partition(nums []int, lo, hi int) int {
        start, end, key := lo, hi, nums[lo]
        for start < end {
            for start <= hi && nums[start] <= key {
                start++
            }
            for end >= lo && nums[end] > key {
                end--
            }
            if start < end {
                nums[start], nums[end] = nums[end], nums[start]
            }
        }
        nums[lo], nums[end] = nums[end], key
        return end
    }

    func quickSort(nums []int) {
        _quickSort(nums, 0, len(nums)-1)
    }

    // 递归实现快排
    func _quickSort(nums []int, lo, hi int) {
        if hi <= lo {
            return
        }
        j := partition(nums, lo, hi)
        _quickSort(nums, lo, j-1) // nums[lo,,,j-1]
        _quickSort(nums, j+1, hi) // nums[j+1,,,hi]

    }

    ```
	- 红黑树的特征
	- 手撕代码：判断BST

### 计算机与操作系统
- #### 进程和线程的区别
    ```
    - 进程是资源分配的最小单位，线程是程序执行的最小单位（资源调度的最小单位）
    - 进程有自己的独立地址空间，每启动一个进程，系统就会为它分配地址空间，建立数据表来维护代码段、堆栈段和数据段，这种操作非常昂贵。而线程是共享进程中的数据的，使用相同的地址空间，因此CPU切换一个线程的花费远比进程要小很多，同时创建一个线程的开销也比进程要小很多。
    - 线程之间的通信更方便，同一进程下的线程共享全局变量、静态变量等数据，而进程之间的通信需要以通信的方式（IPC)进行。不过如何处理好同步与互斥是编写多线程程序的难点。
    - 多进程程序更健壮，多线程程序只要有一个线程死掉，整个进程也死掉了，而一个进程死掉并不会对另外一个进程造成影响，因为进程有自己独立的地址空间。
    ```
- #### 线程状态
    ```
    new、ready 、running、block、dead
    ```

- #### 进程与线程的资源
    ```
    - 线程共享：进程代码段、进程的公有数据(利用这些共享的数据，线程很容易的实现相互之间的通讯)、进程打开的文件描述符、信号的处理器、进程的当前目录和进程用户ID与进程组ID。
    - 线程独有：栈（保存其运行状态和局部自动变量）、程序计数器
    ```

- #### 进程与线程的同步
    ```
    - 进程：无名管道、有名管道、信号、共享内存、消息队列、信号量
    - 线程：互斥量、读写锁、自旋锁、线程信号、条件变量
    ```

###　数据库
- #### MySQL锁的种类 和 读写控制。
- #### 事务特性 ，隔离级别
    ```
    - 原子性、一致性、隔离性、持久性
    ```
	- 脏读，不可重复读，幻读，解释
	- 索引有几种，聚簇和非聚簇，回表操作
	- mysql存储引擎，区别
	- redis持久化方式，原理，效率，区别
	- Redis内存淘汰机制

### 网络
- #### GET 和 POST 的区别
    ```
    - GET在浏览器回退时是无害的，而POST会再次提交请求。
    - GET产生的URL地址可以被Bookmark，而POST不可以。
    - GET请求会被浏览器主动cache，而POST不会，除非手动设置。
    - GET请求只能进行url编码，而POST支持多种编码方式。
    - GET请求参数会被完整保留在浏览器历史记录里，而POST中的参数不会被保留。
    - GET请求在URL中传送的参数是有长度限制的，而POST么有。
    - 对参数的数据类型，GET只接受ASCII字符，而POST没有限制。
    - GET比POST更不安全，因为参数直接暴露在URL上，所以不能用来传递敏感信息。
    - GET参数通过URL传递，POST放在Request body中。
    ```

	- TCP丢包处理
	- 三次握手
	- https请求过程
	- http和tcp的区别
	- HTTP头部属性
- #### http和https的区别
```
 - https协议需要到CA（Certificate Authority，证书颁发机构）申请证书，一般免费证书较少，因而需要一定费用。
 - http是超文本传输协议，信息是明文传输，https则是具有安全性的ssl加密传输协议。
 - http和https使用的是完全不同的连接方式，用的端口也不一样，前者是80，后者是443。
 - http的连接很简单，是无状态的。Https协议是由SSL+Http协议构建的可进行加密传输、身份认证的网络协议，比http协议安全。(无状态的意思是其数据包的发送、传输和接收都是相互独立的。无连接的意思是指通信双方都不长久的维持对方的任何信息。)
```
Python
	- python 装饰器， 深浅拷贝， list和set 查询x，谁复杂度低，why。
	- python dict 实现原理（哈希表),, 利用dict查看1亿个手机号的众数，内存装不下dict

## golang
	- Golang协程与线程的区别
	- gc
	- go的值类型和引用类型
- ### product-sonsumer
    ```

    ```

docker
	- docker 进程如何相互隔离
