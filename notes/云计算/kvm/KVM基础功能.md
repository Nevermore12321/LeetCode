# KVM 核心基础功能

[toc]

## 1. CPU 配置

在 QEMU/KVM 中，QEMU 提供对 CPU 的模拟，展现给客户机一定的 CPU 数目和 CPU 的特性。在 KVM 打开的情况下，**客户机中 CPU 指令的执行由硬件处理器的虚拟化功能（如Intel VT-x和AMD AMD-V）来辅助执行，具有非常高的执行效率**

### vCPU 概念

**QEMU/KVM 为客户机提供一套完整的硬件系统环境，在客户机看来，其所拥有的 CPU 即是 vCPU（virtual CPU）。**

其实，<u>在 KVM 环境中，每个客户机都是一个标准的 Linux 进程（QEMU进程），而每一个 vCPU 在宿主机中是 QEMU 进程派生的一个普通线程。</u>

**vCPU 在 3 种执行模式下的不同分工如下**：

1. **用户模式（User Mode）**
    - 主要处理 I/O 的模拟和管理，由 QEMU 的代码实现。
2. **内核模式（Kernel Mode）**
    - 主要处理特别需要高性能和安全相关的指令，如处理客户模式到内核模式的转换，处理客户模式下的 I/O 指令或其他特权指令引起的退出（VM-Exit），处理影子内存管理（shadow MMU），后改为扩展页表（EPT Extended Page Tables）。
3. **客户模式（Guest Mode）**
    - 主要执行 Guest 中的大部分指令，I/O 和一些特权指令除外（它们会引起 VM-Exit，被 Hypervisor 截获并模拟）。

在 KVM 环境中，整个系统的基本分层架构如下图：

![kvm分层模型](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/kvm%E5%88%86%E5%B1%82%E6%A8%A1%E5%9E%8B.PNG)

注意：

1. **在系统的底层 CPU 硬件中需要有硬件辅助虚拟化技术的支持（Intel VT或AMD-V等）**
2. 宿主机就运行在硬件之上，**KVM 的内核部分是作为可动态加载内核模块运行在宿主机中的**
    1. 其中一个模块是<u>与硬件平台无关的实现虚拟化核心基础架构的 kvm 模块</u>
    2. 另一个是<u>硬件平台相关的 kvm_intel（或kvm_amd）模块</u>。
3.  **KVM 中的一个客户机是作为一个用户空间进程（qemu）运行的**，它和其他普通的用户空间进程（如gnome、kde、firefox、chrome等）一样**由内核来调度**，**使其运行在物理 CPU 上，不过它由 kvm 模块的控制，可以在前面介绍的3种执行模式下运行**。
    1. 多个客户机就是宿主机中的多个QEMU进程，而一个客户机的多个vCPU就是一个QEMU进程中的多个线程。
    2. 和普通操作系统一样，在客户机系统中，同样分别运行着客户机的内核和客户机的用户空间应用程序。

### SMP 的支持

**SMP（Symmetric Multi-Processor，对称多处理器）系统**：是一种多处理器的电脑硬件架构，在对称多处理架构下，每个处理器的地位都是平等的，对资源的使用权限相同。

在 SMP 系统中，多个程序（进程）可以**真正做到并行执行**，**而且单个进程的多个线程也可以得到并行执行**，这极大地提高了计算机系统并行处理能力和整体性能。随着多核技术、超线程（Hyper-Threading）技术的出现，SMP系统使用多处理器、多核、超线程等技术中的一个或多个

下面的Bash脚本（cpu-info.sh）可以根据 /proc/cpuinfo 文件来检查当前系统中的 CPU 数量、多核及超线程的使用情况：

```shell
#!/bin/bash
#filename: cpu-info.sh
#this script only works in a Linux system which has one or more identical physical CPU(s).
echo -n "logical CPU number in total: "
#逻辑CPU个数
cat /proc/cpuinfo | grep "processor" | wc -l

#有些系统没有多核也没有打开超线程，就直接退出脚本
cat /proc/cpuinfo | grep -qi "core id"
if [ $? -ne 0 ]; then
	echo "Warning. No multi-core or hyper-threading is enabled."
	exit 0;
fi

echo -n "physical CPU number in total: "
#物理CPU个数
cat /proc/cpuinfo | grep "physical id" | sort | uniq | wc -l

echo -n "core number in a physical CPU: "
#每个物理CPU上core的个数(未计入超线程)
core_per_phy_cpu=$(cat /proc/cpuinfo | grep "core id" | sort | uniq | wc -l)
echo $core_per_phy_cpu

echo -n "logical CPU number in a physical CPU: "
#每个物理CPU中逻辑CPU(可能是core、threads或both)的个数
logical_cpu_per_phy_cpu=$(cat /proc/cpuinfo | grep "siblings" | sort | uniq | awk -F: '{print $2}')
echo $logical_cpu_per_phy_cpu

#是否打开超线程，以及每个core上的超线程数目
#如果在同一个物理CPU上的两个逻辑CPU具有相同的”core id”，那么超线程是打开的
#此处根据前面计算的core_per_phy_cpu和logical_core_per_phy_cpu的比较来查看超线程
if [ $logical_cpu_per_phy_cpu -gt $core_per_phy_cpu ]; then
	echo "Hyper threading is enabled. Each core has $(expr $logical_cpu_per_phy_cpu / $core_per_phy_cpu ) threads."
elif [ $logical_cpu_per_phy_cpu -eq $core_per_phy_cpu ]; then
	echo "Hyper threading is NOT enabled."
else
	echo "Error. There's something wrong."
fi
```

**QEMU 在给客户机模拟 CPU 时，也可以提供对 SMP 架构的模拟，让客户机运行在 SMP 系统中，充分利用物理硬件的 SMP 并行处理优势。**

由于每个 vCPU 在宿主机中都是一个<u>线程</u>，并且宿主机 Linux 系统是支持多任务处理的，因此可<u>以通过两种操作来实现客户机的 SMP</u>：

- 一是将不同的 vCPU 的进程交换执行（分时调度，即使物理硬件非 SMP，也可以为客户机模拟出 SM P系统环境）
- 二是将在物理 SMP 硬件系统上同时执行多个 vCPU 的进程。

在 qemu 命令行中，“-smp” 参数即是配置客户机的 SMP 系统：

```bash
-smp [cpus=]n[,maxcpus=cpus][,cores=cores][,threads=threads][,sockets=sockets]
```

其中：

- n 用于设置客户机中使用的逻辑 CPU 数量（默认值是1）。
- maxcpus 用于设置客户机中最大可能被使用的 CPU 数量，包括启动时处于下线（offline）状态的CPU数量（可用于热插拔 hot-plug 加入 CPU，但不能超过 maxcpus 这个上限）。
- cores 用于设置每个 CPU 的 core 数量（默认值是1）。
- threads 用于设置每个 core 上的线程数（默认值是1）。
- sockets 用于设置客户机中看到的总的 CPU socket 数量。

由于 vCPU 其实是 QEMU 进程中的一个线程，因此分配多个 vCPU 时，就会起多个线程。

### CPU 过载使用

**KVM 允许客户机过载使用（over-commit）物理资源，即允许为客户机分配的 CPU 和内存数量多于物理上实际存在的资源。**

**CPU 的过载使用是让一个或多个客户机使用 vCPU 的总数量超过实际拥有的物理 CPU数量**。QEMU 会启动更多的线程来为客户机提供服务，这些线程也被 Linux 内核调度运行在物理CPU硬件上。

关于 CPU 的过载使用，**推荐的做法**是<u>对多个单CPU的客户机使用over-commit</u>，比如，在拥有 4 个逻辑 CPU 的宿主机中，同时运行多于 4 个（如 8 个、16 个）客户机，其中每个客户机都分配一个 vCPU。这时，如果每个宿主机的负载不是很大，宿主机Linux对每个客户机的调度是非常有效的，这样的过载使用并不会带来客户机的性能损失。（也就是每个客户机分配不超过宿主机 CPU 数量的 vCPU，而所有客户机的 vCPU 总和大于宿主机逻辑 CPU 的个数）。

关于 CPU 的过载使用，最~~**不推荐的做法**~~是让<u>某一个客户机的 vCPU 数量超过物理系统上存在的CPU数量</u>。比如，在拥有 4 个逻辑 CPU 的宿主机中，同时运行一个或多个客户机，其中每个客户机的 vCPU 数量多于 4 个（如 16 个）

KVM 允许 CPU 的过载使用，但是并不推荐在实际的生产环境（特别是负载较重的环境）中过载使用 CPU。在生产环境中过载使用 CPU，有必要在部署前进行严格的性能和稳定性测试。

### CPU 模型

每一种虚拟机管理程序（Virtual Machine Monitor，简称 VMM 或 Hypervisor）都会定义自己的策略，让客户机看起来有一个默认的 CPU 类型。**也就是 VMM 通过策略，定义客户机中 vCPU 的类型**。

有的 Hypervisor 会简单地将宿主机中 CPU 的类型和特性直接传递给客户机使用，而 QEMU/KVM 在默认情况下会向客户机提供一个名为 qemu64 或 qemu32 的基本 CPU 模型。

QEMU/KVM 的这种策略会带来一些好处：

- 如可以对 CPU 特性提供一些高级的过滤功能
- 还可以将物理平台根据提供的基本 CPU 模型进行分组，从而使客户机在同一组硬件平台上的动态迁移更加平滑和安全

通过如下的命令行可以查看当前的 QEMU 支持的所有 CPU 模型：

```bash
[root@kvm-host ~]# qemu-system-x86_64 -cpu ?
x86 qemu64 QEMU Virtual CPU version 2.5+
x86 phenom AMD Phenom(tm) 9550 Quad-Core Processor
x86 core2duo Intel(R) Core(TM)2 Duo CPU T7700 @ 2.40GHz
x86 kvm64 Common KVM processor
x86 qemu32 QEMU Virtual CPU version 2.5+
x86 kvm32 Common 32-bit KVM processor
x86 coreduo Genuine Intel(R) CPU T2600 @ 2.16GHz
x86 486
x86 pentium
x86 pentium2
x86 pentium3
x86 athlon QEMU Virtual CPU version 2.5+
x86 n270 Intel(R) Atom(TM) CPU N270 @ 1.60GHz
x86 Conroe Intel Celeron_4x0 (Conroe/Merom Class Core 2)
x86 Penryn Intel Core 2 Duo P9xxx (Penryn Class Core 2)
x86 Nehalem Intel Core i7 9xx (Nehalem Class Core i7)
x86 Westmere Westmere E56xx/L56xx/X56xx (Nehalem-C)
x86 SandyBridge Intel Xeon E312xx (Sandy Bridge)
x86 IvyBridge Intel Xeon E3-12xx v2 (Ivy Bridge)
x86 Haswell-noTSX Intel Core Processor (Haswell, no TSX)
x86 Haswell Intel Core Processor (Haswell)
x86 Broadwell-noTSX Intel Core Processor (Broadwell, no TSX)
x86 Broadwell Intel Core Processor (Broadwell)
x86 Skylake-Client Intel Core Processor (Skylake)
x86 Opteron_G1 AMD Opteron 240 (Gen 1 Class Opteron)
x86 Opteron_G2 AMD Opteron 22xx (Gen 2 Class Opteron)
x86 Opteron_G3 AMD Opteron 23xx (Gen 3 Class Opteron)
x86 Opteron_G4 AMD Opteron 62xx class CPU
x86 Opteron_G5 AMD Opteron 63xx class CPU
x86 host KVM processor with all supported host features (only available in KVM mode)
...
```

**在 qemu 命令行中，可以用 “-cpu cpu_model” 来指定在客户机中的CPU模型。**

### 进程的处理器亲和性和 vCPU 的绑定

通常在 SMP 系统中，Linux 内核的**进程调度器**根据自有的调度策略将系统中的一个进程调度到某个 CPU 上执行。

例如：一个进程在前一个执行时间是在 cpuM（M为系统中的某 CPU 的 ID）上运行，而在后一个执行时间是在 cpuN（N 为系统中另一 CPU 的 ID）上运行。因为 Linux 对进程执行的调度采用时间片法则（即用完自己的时间片即被暂停执行），而在默认情况下，一个普通进程或线程的处理器亲和性体现在所有可用的 CPU 上，进程或线程有可能在这些 CPU 之中的任何一个（包括超线程）上执行

**进程的处理器亲和性（Processor Affinity）**：即 CPU 的绑定设置，<u>是指将进程绑定到特定的一个或多个 CPU 上去执行，而不允许将进程调度到其他的 CPU 上</u>。Linux 内核对进程的调度算法也是遵守进程的处理器亲和性设置的。

- 好处：可以减少进程在多个 CPU 之间交换运行带来的缓存命中失效（cache missing），从该进程运行的角度来看，可能带来一定程度上的性能提升。
- 坏处：破坏了原有 SMP 系统中各个 CPU 的负载均衡（load balance），这可能会导致整个系统的进程调度变得低效。

每个 vCPU 都是宿主机中一个普通的 QEMU 线程，可以使用 taskset 工具对其设置处理器亲和性，使其绑定到某一个或几个固定的 CPU 上去调度。

【实际案例】：提供一个有两个逻辑 CPU 计算能力的一个客户机，要求 CPU 资源独立被占用，不受宿主机中其他客户机的负载水平的影响。实现步骤为：

1. 启动宿主机时隔离出两个逻辑CPU专门供一个客户机使用。也就是通过修改 linux 内核参数，使得某些 CPU 可以隔离，也就是设置了隔离的 CPU ，普通的默认进程不会被调度到被隔离的 CPU 上。具体的实现步骤：
    1. 查看机器的逻辑 CPU 的个数。（注意：物理 CPU 就是实际的 CPU 槽数；每个 CPU 有多个核；如果没有开启超线程，那么逻辑 CPU 的个数 = 物理 CPU 的个数 * 每个 CPU 的核数；如果开启了超线程，那么 逻辑 CPU 的个数 >= 物理 CPU 的个数 * 每个 CPU 的核数
    2. 修改 `/etc/tuned/realtime-variables.conf` 文件，添加 `isolcpus=参数`，表示设置哪几个 CPU 被隔离出来，不受 Linux 内核调度。
    3. 激活配置文件，并且需要重启机器。下面就是用到的一些命令（Centos）：

```bash
1. 查看逻辑 CPU 个数：
lscpu
2. 修改隔离 CPU 的配置：
 vim realtime-variables.conf 添加 isolated_cores=0-3,5,7 
3. 激活隔离 CPU
tuned-adm profile realtime
4. 解除隔离 CPU
tuned-adm profile balanced
5. 每次添加隔离 CPU 或者解除隔离 CPU 都需要重启才会生效。
6. 查看隔离 CPU 是否生效：
cat /proc/cmdline
```

2. 在设置了隔离 CPU 后，就可以启动虚拟客户机，来使得客户机的 vCPU 绑定到宿主机的这两个隔离的 CPU 上。过程如下：
    1. 启动一个客户机 ： `qemu-system-x86_64 -enable-kvm -smp 2 -m 4G rhel7.img -daemonize`
    2. 查看代表 vCPU 的 QEMU 线程 `ps -eLo ruser,pid,ppid,lwp,psr,args | grep qemu | grep -v grep`
    3. 绑定代表整个客户机的 QEMU 进程号，使其运行在 cpu4 上 : `taskset -pc 4 8645`
    4. 后面继续绑定另一个 QEMU 进程，使其运行在特定的 CPU 上。



注意：上述过程用到了两个重要的命令：

- `taskset-pc cpulist pid` - 将进程号为 pid 的进程，绑定到一系列 cpulist 上
- `ps -eLo ruser,pid,ppid,lwp,psr,args` - ps 命令显示当前系统的进程信息的状态
    - “-e” ： 参数用于显示所有的进程
    - “-L”：参数用于将线程（light-weight process，LWP）也显示出来
    - “-o”：参数表示以用户自定义的格式输出
        - “psr”：表示当前分配给进程运行的处理器编号
        - “lwp”：列表示线程的ID
        - “ruser”：表示运行进程的用户
        - “pid”：表示进程的ID
        - “ppid”：表示父进程的ID
        - “args”：表示运行的命令及其参数）



实际操作：

1. 隔离出两个 CPU 专门供一个客户机使用。

    1. 修改隔离 CPU 的配置
    2. 激活隔离 CPU
    3. 重启，重启后，检查是否设置成功

    ```bash
    ```

    

2. 启动客户机，并且绑定到这两个隔离出来的 CPU

    1. 需要创建一个镜像文件，qemu-img还支持创建其他格式的 image 文件，比如 qcow2，甚至是其他虚拟机用到的文件格式，比如 VMware 的vmdk、vdi、vhd 等
    2. 准备好 ISO 镜像
    3. 启动客户机，并且绑定隔离出来 CPU

    ```bash
    ```

    

【总结】：

总的来说，<u>在 KVM 环境中，一般并不推荐手动设置 QEMU 进程的处理器亲和性来绑定 vCPU</u>，但是，在非常了解系统硬件架构的基础上，**根据实际应用的需求，可以将其绑定到特定的 CPU 上，从而提高客户机中的 CPU 执行效率或实现 CPU 资源独享的隔离性**。

**NUMA**（Non-Uniform Memory Access，非一致性内存访问）是一种在多处理系统中的内存设计架构，在多处理器中，CPU 访问系统上各个物理内存的速度可能不一样，一个 CPU 访问其本地内存的速度比访问（同一系统上）其他 CPU 对应的本地内存快一些。

比如一台机器是有 2 个处理器，有 4 个内存块。将1个处理器和两个内存块合起来，称为一个 NUMA node，这样这个机器就会有两个 NUMA node。在物理分布上，NUMA node 的处理器和内存块的物理距离更小，因此访问也更快。



## 2. 内存配置

内存是与 CPU 沟通的一个桥梁，其作用是暂时存放 CPU 中将要执行的指令和数据，所有程序都必须先载入内存中才能够执行。

### 内存设置基本参数

在通过 qemu 命令行启动客户机时设置内存大小的参数如下：

```bash
-m megs #设置客户机的内存为megs MB大小
```

注意：默认的单位为 MB，也支持加上“M”或“G”作为后缀来显式指定使用 MB 或 GB 作为内存分配的单位。如果不设置 -m 参数，QEMU 对客户机分配的内存大小默认值为 128MB

### EPT 和 VPID 简介

**EPT（Extended Page Tables，扩展页表）**，属于 Intel 的第二代硬件虚拟化技术，它是**针对内存管理单元（MMU）的虚拟化扩展。**

- EPT 降低了内存虚拟化的难度（与影子页表相比），也提升了内存虚拟化的性能。
- 从基于 Intel 的 Nehalem 架构的平台开始，EPT 就作为 CPU 的一个特性加入 CPU 硬件中了。（<u>从硬件上支持内存虚拟化</u>）

**地址转换**：

【目的】：在客户机操作系统看来，客户机可用的内存空间也是一个从零地址开始的连续的物理内存空间。

【过程】：Hypervisor（即 KVM）引入了一层新的地址空间，即**客户机物理地址空间（HPA）**，这个地址空间<u>不是真正的硬件上的地址空间，它们之间还有一层映射</u>。所以，在虚拟化环境下，内存使用就需要两层的地址转换，即<u>客户机应用程序可见的客户机虚拟地址（Guest Virtual Address，GVA）到客户机物理地址（Guest Physical Address，GPA）的转换</u>，<u>再从客户机物理地址（GPA）到宿主机物理地址（Host Physical Address，HPA）的转换</u>。其中，前一个转换由客户机操作系统来完成，而后一个转换由 Hypervisor 来负责。

在 EPT 出现之前，使用**影子页表（Shadow Page Tables）**实现（如下图）：

![影子页表的作用](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/%E5%BD%B1%E5%AD%90%E9%A1%B5%E8%A1%A8%E7%9A%84%E4%BD%9C%E7%94%A8.PNG)

注意：

- 从软件上维护了从客户机虚拟地址（GVA）到宿主机物理地址（HPA）之间的映射，<u>每一份客户机操作系统的页表也对应一份影子页表</u>
- <u>Hypervisor 将影子页表载入物理上的内存管理单元（Memory Management Unit，MMU）中进行地址翻译</u>
- 在普通的内存访问时都可实现从 GVA 到 HPA 的直接转换，从而避免了上面前面提到的两次地址转换。

影子页表实现的缺点：

- 实现非常复杂，导致其开发、调试和维护都比较困难
- 影子页表的内存开销也比较大，因为需要为每个客户机进程对应的页表的都维护一个影子页表

为了克服影子页表的缺点，Intel 的 CPU 提供了 **EPT 技术**（AMD 提供的类似技术叫作 NPT，即 Nested Page Tables），**直接在硬件上支持 GVA→GPA→HPA 的两次地址转换，从而降低内存虚拟化实现的复杂度，也进一步提升了内存虚拟化的性能。**（如下图）

![EPT基本原理](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/EPT%E5%9F%BA%E6%9C%AC%E5%8E%9F%E7%90%86.PNG)

注意：

- CR3（控制寄存器3）将客户机程序所见的客户机虚拟地址（GVA）转化为客户机物理地址（GPA），然后再通过EPT将客户机物理地址（GPA）转化为宿主机物理地址（HPA）
- 两次地址转换都是由CPU硬件来自动完成的，其转换效率非常高
- EPT只需要维护一张EPT页表，而不需要像“影子页表”那样为每个客户机进程的页表维护一张影子页表，从而也减少了内存的开销



**VPID（Virtual Processor Identifiers，虚拟处理器标识）**是<u>在硬件上对 TLB 资源管理的优化，通过在硬件上为每个 TLB 项增加一个标识，用于不同的虚拟处理器的地址空间，从而能够区分 Hypervisor 和不同处理器的 TLB</u>。

其中，**TLB（translation lookaside buffer，旁路转换缓冲）**是<u>内存管理硬件以提高虚拟地址转换速度的缓存</u>。TLB 是页表（page table）的缓存，保存了一部分页表。

EPT  结合 VPID 的优势：

- 硬件区分了不同的 TLB 项分别属于不同虚拟处理器，因此可以避免每次进行 VM-Entry 和 VM-Exit 时都让 TLB 全部失效，提高了 VM 切换的效率。
- 由于有了这些在 VM 切换后仍然继续存在的 TLB 项，硬件减少了一些不必要的页表访问，减少了内存访问次数，从而提高了 Hypervisor 和客户机的运行速度。
- VPID 也会对客户机的实时迁移（Live Migration）有很好的效率提升，会节省实时迁移的开销，提升实时迁移的速度，降低迁移的延迟（Latency）。



【查看系统是否支持 EPT 和 VPID 功能】：

1. 查看 /proc/cpuinfo 中的 CPU 标志

```bash
[root@bogon ~]# grep -E "ept|vpid" /proc/cpuinfo
flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc aperfmperf eagerfpu pni dtes64 monitor ds_cpl vmx est tm2 ssse3 cx16 xtpr pdcm sse4_1 sse4_2 popcnt lahf_lm ssbd ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid dtherm ida arat spec_ctrl intel_stibp flush_l1d
```

2. 根据 sysfs 文件系统中 kvm_intel 模块

```bash
[root@bogon ~]# cat /sys/module/kvm_intel/parameters/ept
Y
[root@bogon ~]# cat /sys/module/kvm_intel/parameters/vpid
Y
```

【打开或关闭 EPT 和 VPID 特性】：设置 ept 和 vpid 参数的值来打开或关闭 EPT 和 VPID。如果 kvm_intel 模块已经处于加载状态，则需要先卸载这个模块，在重新加载之时加入所需的参数设置。

```bash
[root@kvm-host ~]# modprobe kvm_intel ept=0,vpid=0
[root@kvm-host ~]# rmmod kvm_intel
[root@kvm-host ~]# modprobe kvm_intel ept=1,vpid=1
```

注意：一般不要手动关闭 EPT 和 VPID 功能，否则会导致客户机中内存访问的性能下降。

### 内存过载使用

与 CPU 过载使用类似，**在 KVM 中内存也是允许过载使用（over-commit）的，KVM 能够让分配给客户机的内存总数大于实际可用的物理内存总数。**

**原理**：<u>KVM 中客户机是一个 QEMU 进程</u>，宿主机系统没有特殊对待它而分配特定的内存给 QEMU，只是把它当作一个普通 Linux 进程。**Linux 内核在进程请求更多内存时才分配给它们更多的内存，所以也是在客户机操作系统请求更多内存时，KVM 才向其分配更多的内存**。

有如下 3 种方式来实现内存的过载使用：

1. **内存交换（swapping）**：用交换空间（swap space）来弥补内存的不足。
    1. 用 swapping 方式来让内存过载使用，要求有足够的交换空间（swap space）来满足所有的客户机进程和宿主机中其他进程所需内存
    2. **可用的物理内存空间和交换空间的大小之和应该等于或大于配置给所有客户机的内存总和**，否则，在各个客户机内存使用同时达到较高比率时，可能会有客户机（因内存不足）被强制关闭
2. **气球（ballooning）**：通过 virio_balloon 驱动来实现宿主机 Hypervisor 和客户机之间的协作。
3. **页共享（page sharing）**：通过 KSM（Kernel Samepage Merging）合并多个客户机进程使用的相同内存页。

注意：KVM 允许内存过载使用，但在生产环境中配置内存的过载使用之前，仍然应该根据实际应用进行充分的测试。
