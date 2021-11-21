# KVM内存管理高级技巧

[toc]

## 1. 大页

### 大页的介绍

**x86（包括x86-32和x86-64）架构的CPU默认使用4KB大小的内存页面**，但是**它们也支持较大的内存页，如x86-64系统就支持2MB及1GB大小的大页（Huge Page）。**

Linux 2.6及以上的内核都支持Huge Page。如果在系统中使用了Huge Page，

- 则<u>内存页的数量会减少，从而需要更少的页表（Page Table），节约了页表所占用的内存数量，并且所需的地址转换也减少了，TLB缓存失效的次数就减少了，从而提高了内存访问的性能</u>。
- 另外，<u>由于地址转换所需的信息一般保存在CPU的缓存中，Huge Page的使用让地址转换信息减少，从而减少了CPU缓存的使用，减轻了CPU缓存的压力，让CPU缓存能更多地用于应用程序的数据缓存，也能够在整体上提升系统的性能</u>。

### KVM虚拟化对大页的利用

操作系统之上的应用程序（包括QEMU创建的客户机）要利用上大页，有下面3种途径之一：

- mmap系统调用使用MAP_HUGETLB flag创建一段内存映射。
- shmget系统调用使用SHM_HUGETLB flag创建一段共享内存区。
- 使用hugetlbfs创建一个文件挂载点，这个挂载目录之内的文件就都是使用大页的了

在KVM虚拟化环境中，就是使用第3种方法：**创建hugetlbfs的挂载点**: 

- 通过“-mem-path FILE”选项，让客户机使用宿主机的huge page挂载点作为它内存的backend。
- 另外，还有一个参数“-mem-prealloc”是让宿主机在启动客户机时就全部分配好客户机的内存，而不是在客户机实际用到更多内存时才按需分配。-mem-prealloc必须在有“-mem-path”参数时才能使用。

## 2. 透明大页

使用大页可以提高系统内存的使用效率和性能。**不过大页有如下几个缺点**：

- 大页必须在使用前就准备好。
- 应用程序代码必须显式地使用大页（一般是调用mmap、shmget系统调用，或者使用libhugetlbfs库对它们的封装）。
- 大页必须常驻物理内存中，不能交换到交换分区中。
- 需要超级用户权限来挂载hugetlbfs文件系统，尽管挂载之后可以指定挂载点的uid、gid、mode等使用权限供普通用户使用。
- 如果预留了大页内存但没实际使用，就会造成物理内存的浪费。

<u>透明大页（Transparent Hugepage）既发挥了大页的一些优点，又能避免了上述缺点。透明大页（THP）是Linux内核的一个特性。</u>

- 透明大页，如它的名称描述的一样，对所有应用程序都是透明的（transparent），应用程序不需要任何修改即可享受透明大页带来的好处。在使用透明大页时，普通的使用hugetlbfs大页依然可以正常使用，而在没有普通的大页可供使用时，才使用透明大页。
- 透明大页是可交换的（swapable），当需要交换到交换空间时，透明大页被打碎为常规的4KB大小的内存页。在使用透明大页时，
    - 如果因为内存碎片导致大页内存分配失败，这时系统可以优雅地使用常规的4KB页替换，而且不会发生任何错误、故障或用户态的通知。
    - 而当系统内存较为充裕、有很多的大页可用时，常规的页分配的物理内存可以通过khugepaged内核线程自动迁往透明大页内存。
- 内核线程khugepaged的作用是，扫描正在运行的进程，然后试图将使用的常规内存页转换到使用大页。



## 3. KSM

在现代操作系统中，**共享内存**被很普遍地应用。如在Linux系统中，当使用fork函数创建一个进程时，子进程与其父进程共享全部的内存，而当子进程或父进程试图修改它们的共享内存区域之时，内核会分配一块新的内存区域，并试图将修改的共享内存区域复制到新的内存区域上，然后让进程去修改复制的内存。这就是著名的**“写时复制”（copy-on-write，COW）**技术。

### KSM基本原理

**KSM是“Kernel SamePage Merging”的缩写**，中文可称为“**内核同页合并**”。

- KSM允许内核在两个或多个进程（包括虚拟客户机）之间共享完全相同的内存页。
- <u>KSM让内核扫描检查正在运行中的程序并比较它们的内存，如果发现它们有完全相同的内存区域或内存页，就将多个相同的内存合并为一个单一的内存页，并将其标识为“写时复制”。</u>
- 之后，<u>如果有进程试图去修改被标识为“写时复制”的合并内存页，就为该进程复制出一个新的内存页供其使用。</u>

**在QEMU/KVM中，一个虚拟客户机就是一个QEMU进程，所以使用KSM也可以实现多个客户机之间的相同内存合并。**

如果在同一宿主机上的多个客户机运行的是相同的操作系统或应用程序，则客户机之间的相同内存页的数量就可能比较大，这种情况下KSM的作用就更加显著。

在KVM环境下使用KSM，还允许KVM请求哪些相同的内存页是可以被共享而合并的，所以KSM只会识别并合并那些不会干扰客户机运行且不会影响宿主机或客户机运行的安全内存页。可见，在KVM虚拟化环境中，KSM能够提高内存的速度和使用效率。

- **在KSM的帮助下，相同的内存页被合并了，减少了客户机的内存使用量**。
    - 一方面，内存中的内容更容易被保存到CPU的缓存中
    - 另一方面，有更多的内存可用于缓存一些磁盘中的数据。
    - 因此，不管是内存的缓存命中率（CPU缓存命中率），还是磁盘数据的缓存命中率（在内存中命中磁盘数据缓存的命中率）都会提高，从而提高了KVM客户机中操作系统或应用程序的运行速度。
- **KSM是内存过载使用的一种较好的方式**。
    - KSM通过减少每个客户机实际占用的内存数量，可以让多个客户机分配的内存数量之和大于物理上的内存数量。
    - 对于使用相同内存量的客户机而言，在物理内存量不变的情况下，可以在一个宿主机中创建更多的客户机，提高了虚拟化客户机部署的密度，提高了物理资源的利用效率

KSM最初就是为KVM虚拟化中的使用而开发的，不过它对非虚拟化的系统依然非常有用。在KSM的帮助下，有人在物理内存为16GB的机器上，用KVM成功运行了多达52个1GB内存的Windows XP客户机。

**KSM必须有一个或多个进程去检测和找出哪些内存页是完全相同可以用于合并的，而且需要找到那些不会经常更新的内存页，这样的页才是最适合于合并的。因此，KSM让内存使用量降低了，但是CPU使用率会有一定程度的升高，也可能会带来隐蔽的性能问题，需要在实际使用环境中进行适当配置KSM的使用，以便达到较好的平衡。**

使用推荐：

KSM对内存合并而节省内存的数量<u>与客户机操作系统类型及其上运行的应用程序有关</u>

- 如果宿主机上的客户机操作系统相同且其上运行的应用程序也类似，节省内存的效果就会很显著，甚至节省超过50%的内存都有可能的。
- 如果客户机操作系统不同，且运行的应用程序也大不相同，KSM节省内存效率就不高，可能连5%都不到

### ksmd守护进程

内核的KSM守护进程是ksmd，配置和监控ksmd的文件在“/sys/kernel/mm/ksm/”目录下：

```bash
[root@bogon ~]# ls -l /sys/kernel/mm/ksm/
total 0
-r--r--r--. 1 root root 4096 Nov 21 21:34 full_scans
-rw-r--r--. 1 root root 4096 Nov 21 21:34 max_page_sharing
-rw-r--r--. 1 root root 4096 Nov 21 21:34 merge_across_nodes
-r--r--r--. 1 root root 4096 Nov 21 21:34 pages_shared
-r--r--r--. 1 root root 4096 Nov 21 21:34 pages_sharing
-rw-r--r--. 1 root root 4096 Nov 21 21:34 pages_to_scan
-r--r--r--. 1 root root 4096 Nov 21 21:34 pages_unshared
-r--r--r--. 1 root root 4096 Nov 21 21:34 pages_volatile
-rw-r--r--. 1 root root 4096 Nov 21 22:20 run
-rw-r--r--. 1 root root 4096 Nov 21 21:34 sleep_millisecs
-r--r--r--. 1 root root 4096 Nov 21 21:34 stable_node_chains
-rw-r--r--. 1 root root 4096 Nov 21 21:34 stable_node_chains_prune_millisecs
-r--r--r--. 1 root root 4096 Nov 21 21:34 stable_node_dups
```

下面简单介绍各个文件的作用:

- **full_scans**：记录已经对所有可合并的内存区域扫描过的次数。
- **merge_across_nodes**：在NUMA（见7.4节）架构的平台上，是否允许跨节点（node）合并内存页。
- **pages_shared**：记录正在使用中的共享内存页的数量。
- **pages_sharing**：记录有多少数量的内存页正在使用被合并的共享页，不包括合并的内存页本身。这就是实际节省的内存页数量。
- **pages_unshared**：记录了守护进程去检查并试图合并，却发现了因没有重复内容而不能被合并的内存页数量。
- **pages_volatile**：记录了因为其内容很容易变化而不被合并的内存页。
- **pages_to_scan**：在ksmd进程休眠之前扫描的内存页数量。
- **sleep_millisecs：ksmd**进程休眠的时间（单位：毫秒），ksmd的两次运行之间的间隔。
- **run**：控制ksmd进程是否运行的参数，默认值为0，要激活KSM必须要设置其值为1（除非内核关闭了sysfs的功能）。设置为0，表示停止运行ksmd但保持它已经合并的内存页；设置为1，表示马上运行ksmd进程；设置为2表示停止运行ksmd，并且分离已经合并的所有内存页，但是保持已经注册为可合并的内存区域给下一次运行使用

说明：

- pages_sharing的值越大，说明KSM节省的内存越多，KSM效果越好。
- pages_sharing除以pages_shared得到的值越大，说明相同内存页重复的次数越多，KSM效率就越高。
- pages_unshared除以pages_sharing得到的值越大，说明ksmd扫描不能合并的内存页越多，KSM的效率越低

Redhat系列系统（如RHEL 6、RHEL 7）中提供了两个服务ksm和ksmtuned，来动态调节KSM的运行情况。ksmtuned其实是一个实时动态调整ksm行为的后台服务，可以理解它的存在是因为前文我们所讲的KSM本身有利有弊，而有了ksmtuned，方便用户合理有效地使用KSM。

## 2. NUMA

### NUMA 原理

**NUMA（Non-Uniform Memory Access，非统一内存访问架构）**是相对于<u>UMA（Uniform Memory Access）</u>而言的。

早年的计算机架构都是UMA，如下图所示。

![UMA架构](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/UMA%E6%9E%B6%E6%9E%84.PNG)

UMA架构中，所有的CPU处理单元（Processor）均质地通过共享的总线访问内存，<u>所有CPU访问所有内存单元的速度是一样的</u>。**在多处理器的情形下，多个任务会被分派在各个处理器上并发执行，则它们竞争内存资源的情况会非常频繁，从而引起效率的下降。**

NUMA架构，处理器与内存被划分成一个个的节点（node），处理器访问自己节点内的内存会比访问其他节点的内存快。如下图所示：

![NUMA架构](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/NUMA%E6%9E%B6%E6%9E%84.PNG)

### NUMA 相关工具

1. **numastat**

安装：`yum install numactl`

<u>numastat用来查看某个（些）进程或者整个系统的内存消耗在各个NUMA节点的分布情况。</u>

```bash
[root@kvm-host ~]# numastat
				node0	 node1
numa_hit		72050204 55925951
numa_miss 		0		 0
numa_foreign 	0		 0
interleave_hit 	38068	 39139
local_node 		71816493 54027058
other_node 		233711	 1898893
```

说明：

- **numa_hit**表示成功地从该节点分配到的内存页数。
- **numa_miss**表示成功地从该节点分配到的内存页数，但其本意是希望从别的节点分配，失败以后退而求其次从该节点分配。
- **numa_foreign**与numa_miss互为“影子”，每个numa_miss都来自另一个节点的numa_foreign。
- **interleave_hit**，有时候内存请求是没有NUMA节点偏好的，此时会均匀分配自各个节点（interleave），这个数值就是这种情况下从该节点分配出去的内存页面数。
- **local_node**表示分配给运行在同一节点的进程的内存页数。
- **other_node**与上面相反。local_node值加上other_node值就是numa_hit值。以上数值默认都是内存页数，要看具体多少MB，可以通过加上-n参数实现。

numastat还可以只看某些进程，甚至只要名字片段匹配，例如 “numastat qemu”

2. **numad**

安装：`yum install numad`

numad是一个可以自动管理NUMA亲和性（affinity）的工具（同时也是一个后台进程）。它实时监控NUMA拓扑结构（topology）和资源使用，并动态调整。同时它还可以在启动一个程序前，提供NUMA优化建议

与numad功能类似，Kernel的auto NUMA balancing（/proc/sys/kernel/numa_balancing）也是进行动态NUMA资源的调节。numad启动后会覆盖Kernel的auto NUMA balancing功能。

`-p<pid>，-x<pid>，-r<pid>`，分别指定numad针对哪些pid以及不针对哪些pid进行自动的NUMA资源优化

numad自己内部维护一个inclusive list和一个exclusive list；`-p<pid>、-x<pid>`就是分别往这两个list里面添加进程id；`-r<pid>`就是从这两个list里面移除。在numad首次启动时候，可以重复多个-p或者-x；启动后，每次调用numad只能跟一个-p、-x或者-r参数。

3. **numactl**

如果说numad是事后（客户机起来以后）调节NUMA资源分配，那么numactl则是主动地在程序起来时候就指定好它的NUMA节点。

**numactl其实不止它名字表示的那样设置NUMA相关亲和性，它还可以设置共享内存/大页文件系统的内存策略，以及进程的CPU和内存的亲和性**。

numactl 主要参数如下:

- `--hardware`，列出来目前系统中可用的NUMA节点，以及它们之间的距离（distance）。
- `--membind`，确保command执行时候内存都是从指定的节点上分配；如果该节点没有足够内存，返回失败。
- `--cpunodebind`，确保command只在指定node的CPU上面执行。
- `--phycpubind`，确保command只在指定的CPU上执行。
- `--localalloc`，指定内存只从本地节点上分配。
- `--preferred`，指定一个偏好的节点，command执行时内存优先从这个节点分配，不够的话才从别的节点分配。
- `--interleave`，让客户机均匀地占用两个节点的资源

使用推荐：

- 在想要专属地让某个客户机得到优先服务的时候，我们可以把KSM关闭，通过numactl将客户机QEMU进程绑定在某个node或某些CPU上。
- 当我们想要更高的客户机密度，而不考虑特别的服务质量的时候，我们可以通过numactl--all，同时打开KSM，关掉numad。