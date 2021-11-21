# KVM设备高级管理

[toc]

## 1. 半虚拟化驱动 - virtio

### virtio 概述

- KVM 是必须使用硬件虚拟化辅助技术（如Intel VT-x、AMD-V）的 Hypervisor，在 CPU 运行效率方面有硬件支持，其效率是比较高的；
- 在有 Intel EPT 特性支持的平台上，内存虚拟化的效率也较高；
- 有 Intel VT-d 的支持，其 I/O 虚拟化的效率也很高。

在 KVM 中，可以在客户机中使用**半虚拟化驱动**（Paravirtualized Drivers，PV Drivers）<u>来提高客户机的性能（特别是I/O性能）</u>。目前，**KVM 中实现半虚拟化驱动的方式是采用 virtio  这个 Linux 上的设备驱动标准框架**。

在 virtio 还没有出现的时候 ，使用的是利用 QEMU 软件虚拟化来模拟 I/O 操作。下面就来看一下 QEMU 软件虚拟化与 virtio 半虚拟化的区别。

### QEMU 软件虚拟化与 virtio 半虚拟化驱动的原理

1. QEMU 模拟 I/O 设备的基本原理和优缺点

QEMU以纯软件方式模拟现实世界中的I/O设备的基本过程如下图：

![QEMU纯软件模拟IO](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/QEMU%E7%BA%AF%E8%BD%AF%E4%BB%B6%E6%A8%A1%E6%8B%9FIO.PNG)

过程为：

- 在使用 QEMU 模拟 I/O 的情况下，当客户机中的设备驱动程序（Device Driver）发起 I/O 操作请求时，KVM 模块（Module）中的 I/O 操作捕获代码会拦截这次 I/O 请求，然后在经过处理后将本次 I/O 请求的信息存放到 I/O 共享页（sharing page），并通知用户空间的 QEMU 程序。
- QEMU 模拟程序获得 I/O 操作的具体信息之后，交由硬件模拟代码（Emulation Code）来模拟出本次的 I/O 操作，完成之后，将结果放回到  I/O  共享页，并通知 KVM 模块中的 I/O 操作捕获代码。
- 最后，由 KVM 模块中的捕获代码读取 I/O 共享页中的操作结果，并把结果返回客户机中。
- 当然，在这个操作过程中，客户机作为一个 QEMU 进程在等待 I/O  时也可能被阻塞。另外，当客户机通过 DMA（Direct Memory Access）访问大块 I/O 时，QEMU 模拟程序将不会把操作结果放到 I/O  共享页中，而是通过内存映射的方式将结果直接写到客户机的内存中去，然后通过 KVM 模块告诉客户机 DMA 操作已经完成。

QEMU 纯软件模拟 I/O 设备的优缺点：

- 优点：
    - 可以通过软件模拟出各种各样的硬件设备，包括一些不常用的或很老很经典的设备
    - 不用修改客户机操作系统，就可以使模拟设备在客户机中正常工作
- 缺点：
    - 每次I/O操作的路径比较长，有较多的VMEntry、VMExit发生，需要多次上下文切换（context switch），也需要多次数据复制，所以它的性能较差。



2. virtio 半虚拟化的基本原理和优缺点

virtio 是一个在 Hypervisor 之上的抽象 API 接口，<u>让客户机知道自己运行在虚拟化环境中</u>，进而根据 virtio 标准与 Hypervisor 协作，从而在客户机中达到更好的性能（特别是 I/O 性能）。

在 QEMU/KVM 中，virtio 的基本结构如下图：

![KVM中的virtio结构](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/KVM%E4%B8%AD%E7%9A%84virtio%E7%BB%93%E6%9E%84.PNG)

**前端驱动**（frondend，如virtio-blk、virtio-net等）**是在客户机中存在的驱动程序模块**，而**后端处理程序（backend）是在QEMU中实现的**。在**前后端驱动之间，还定义了两层来支持客户机与QEMU之间的通信**。

- “virtio” 这一层是**虚拟队列接口**，它<u>在概念上将前端驱动程序附加到后端处理程序</u>。一个前端驱动程序可以使用0个或多个队列，具体数量取决于需求。
    - 例如，virtio-net 网络驱动程序使用两个虚拟队列（一个用于接收，另一个用于发送），而 virtio-blk 块驱动程序仅使用一个虚拟队列。
    - <u>虚拟队列实际上被实现为跨越客户机操作系统和 Hypervisor 的衔接点，但该衔接点可以通过任意方式实现，前提是客户机操作系统和 virtio 后端程序都遵循一定的标准，以相互匹配的方式实现它</u>
- virtio-ring 实现了**环形缓冲区（ring buffer）**，<u>用于保存前端驱动和后端处理程序执行的信息</u>。
    - 该环形缓冲区可以一次性保存前端驱动的多次 I/O 请求，并且**交由后端驱动去批量处理**，最后<u>实际调用宿主机中设备驱动实现物理上的 I/O 操作，这样做就可以根据约定实现批量处理而不是客户机中每次 I/O 请求都需要处理一次，从而提高客户机与 Hypervisor 信息交换的效率</u>。

virtio 的优缺点：

- 优点：
    - 可以获得很好的 I/O 性能，其性能几乎可以达到与 native（即非虚拟化环境中的原生系统）差不多的 I/O 性能
    - 如果宿主机内核和客户机都支持virtio，一般推荐使用 virtio
- 缺点：
    - 要求客户机必须安装特定的 virtio 驱动使其知道是运行在虚拟化环境中，并且按照 virtio 的规定格式进行数据传输



### virtio_balloon 介绍

1. ballooning 简介

通常来说，要改变客户机占用的宿主机内存，要先关闭客户机，修改启动时的内存配置，然后重启客户机才能实现。而<u>内存的ballooning（气球）技术可以在客户机运行时动态地调整它所占用的宿主机内存资源，而不需要关闭客户机。</u>

**ballooning 技术形象地在客户机占用的内存中引入气球（balloon）的概念**（如下图）。

![内存balloon概念](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/%E5%86%85%E5%AD%98balloon%E6%A6%82%E5%BF%B5.PNG)

- <u>气球中的内存是可以供宿主机使用的（但不能被客户机访问或使用）</u>，所以，当宿主机内存紧张，空余内存不多时，可以请求客户机回收利用已分配给客户机的部分内存，客户机就会释放其空闲的内存。
- 此时若客户机空闲内存不足，可能还会回收部分使用中的内存，可能会将部分内存换出到客户机的交换分区（swap）中，从而使内存“气球”充气膨胀，进而使宿主机回收气球中的内存用于其他进程（或其他客户机）。
- 反之，当客户机中内存不足时，也可以让客户机的内存气球压缩，释放出内存气球中的部分内存，让客户机有更多的内存可用。

总之，内存 balloon 其实就是一块备用内存，宿主机内存不够的时候，压缩客户机的内存，让 balloon 充气，供宿主机使用。当客户机内存不够的时候，可以使用 balloon 中的内存，使其放气压缩。



2. KVM 中 ballooning 的原理及优劣势

KVM 中的 ballooning 过程：

- Hypervisor（即 KVM）发送请求到客户机操作系统，让其归还一定数量的内存给 Hypervisor。
- 客户机操作系统中的 virtio_balloon 驱动接收到 Hypervisor 的请求。
- virtio_balloon 驱动使客户机的内存气球膨胀，气球中的内存就不能被客户机访问。如果此时客户机中内存剩余量不多（如某应用程序绑定/申请了大量的内存），并且不能让内存气球膨胀到足够大以满足 Hypervisor 的请求，那么 virtio_balloon 驱动也会尽可能多地提供内存使气球膨胀，尽量去满足 Hypervisor 所请求的内存数量（即使不一定能完全满足）。
- 客户机操作系统归还气球中的内存给 Hypervisor。
- Hypervisor 可以将从气球中得来的内存分配到任何需要的地方。
- 即使从气球中得到的内存没有处于使用中，Hypervisor 也可以将内存返还给客户机。这个过程为：Hypervisor 发送请求到客户机的 virtio_balloon 驱动；这个请求使客户机操作系统压缩内存气球；在气球中的内存被释放出来，重新由客户机访问和使用。

KVM 中 ballooning 的优缺点：

- 优点：
    - 因为 ballooning 能够被控制和监控，所以能够潜在地节约大量的内存。它不同于内存页共享技术（KSM是内核自发完成的，不可控），客户机系统的内存只有在通过命令行调整 balloon 时才会随之改变，所以能够监控系统内存并验证 ballooning 引起的变化。
    - ballooning 对内存的调节很灵活，既可以精细地请求少量内存，也可以粗犷地请求大量的内存。
    - Hypervisor 使用 ballooning 让客户机归还部分内存，从而缓解其内存压力。而且从气球中回收的内存也不要求一定要被分配给另外某个进程（或另外的客户机）
- 缺点：
    - ballooning 需要客户机操作系统加载 virtio_balloon 驱动，然而并非每个客户机系统都有该驱动（如Windows需要自己安装该驱动）。
    - 如果有大量内存需要从客户机系统中回收，那么 ballooning 可能会降低客户机操作系统运行的性能。
        - 一方面，内存的减少可能会让客户机中作为磁盘数据缓存的内存被放到气球中，从而使客户机中的磁盘 I/O 访问增加；
        - 另一方面，如果处理机制不够好，也可能让客户机中正在运行的进程由于内存不足而执行失败。
    - 目前没有比较方便的、自动化的机制来管理 ballooning，一般都采用在 QEMU monitor 中执行 balloon 命令来实现 ballooning。没有对客户机的有效监控，没有自动化的 ballooning 机制，这可能会不便于在生产环境中实现大规模自动化部署。
    - 内存的动态增加或减少可能会使内存被过度碎片化，从而降低内存使用时的性能。另外，内存的变化会影响客户机内核对内存使用的优化，比如，内核起初根据目前状态对内存的分配采取了某个策略，而后由于balloon的原因突然使可用内存减少了很多，这时起初的内存策略就可能不是太优化了。



### virtio_net 介绍



在选择 KVM 中的网络设备时，一般来说应优先选择半虚拟化的网络设备，而不是纯软件模拟的设备。

**使用 virtio_net 半虚拟化驱动可以提高网络吞吐量（thoughput）和降低网络延迟（latency），从而让客户机中网络达到几乎和非虚拟化系统中使用原生网卡的网络差不多的性能。**

virtio_net 其实就是在客户集中的虚拟网卡驱动，流程与 virtio_ballooning 类似：

- 客户机发送请求到 virtio_net 网卡驱动
- virtio_net 通过 virtio 与 virtio_ring 将发送的消息传递给 Hypervisor 
- Hypervisor  处理请求后，返回结果给 客户机

### virtio_blk 介绍

virtio_blk 驱动使用 virtio API 为客户机提供了一个高效访问块设备 I/O 的方法。在 QEMU/KVM 中对块设备使用 virtio，需要在两方面进行配置：

- 客户机中的前端驱动模块 virtio_blk 
- 宿主机中的QEMU**提供后端处理程序**

注意：使用 virtio_blk 驱动的磁盘显示为 “/dev/vda”，这不同于 IDE 硬盘的 “/dev/hda” 或 SATA 硬盘的 “/dev/sda” 这样的显示标识。

而 “/dev/vd*” 这样的磁盘设备名称可能会导致从前分配在磁盘上的swap分区失效，因为有些客户机系统中记录文件系统信息的 “/etc/fstab” 文件中有类似如下的对swap分区的写法。

```shell
/dev/sda2 swap swap defaults 0 0
/dev/hda2 swap swap defaults 0 0
```

原因就是“/dev/vda2”这样的磁盘分区名称未被正确识别。解决这个问题的方法就很简单了，只需要修改它为如下形式并保存到“/etc/fstab”文件，然后重启客户机系统即可。

```shell
/dev/vda2 swap swap defaults 0 0
```

### 内核态的 vhost-net 后端以及网卡多队列

前面提到 virtio 在宿主机中的后端处理程序（backend）一般是由用户空间的 QEMU 提供的，然而，**如果对于网络 I/O 请求的后端处理能够在内核空间来完成，则效率会更高，会提高网络吞吐量和减少网络延迟**。在比较新的**内核中有一个叫作 “vhost-net” 的驱动模块，它作为一个内核级别的后端处理程序，将 virtio-net 的后端处理任务放到内核空间中执行，从而提高效率**。

vhost 技术对 virtio-net 进行了优化，在内核中加入了 vhost-net.ko 模块，使得对网络数据可以在内核态得到处理。

**Vhost的目的就是实现一个从 host kernel 到Guest的直通路径，从而 bypass Qemu 进程**。Vhost 协议定义了如何去建立数据路径，但是如何去实现并没有严格规定，这个实现取决于 host 和 Guest 上 ring 的 layout 和数据 buffer 的描述以及数据报的发送和接受机制。



Vhost因此可以在Kernel中实现（vhost-net）,也可以在用户态实现（vhost-user).

之前 virtio-net 的过程是：guest 通过 virtio-net driver 发送命令 -> 通过 ring 虚拟化环形队列传递给 Hypervisor （QEMU-kvm） -> QEMU-kvm 处理命令 ->（内核）

而现在 vhost-net 的过程是：guest 通过 virtio-net driver 发送命令 -> kvm直接和vhost-net.ko通信，然后由vhost-net.ko访问tap**设备**

vhost 的实现原理图：

![vhost-net实现](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/vhost-net%E5%AE%9E%E7%8E%B0.jpg)

归根结底，vhost-net 就是在 Linux 内核中，与 kvm 平级的一个模块，之前 virtio-net 是通过 kvm 转给内核处理，现在直接给内核处理。

为了性能，因为 virtio 本身就是 queue 的定义，为了性能实现了多队列：

![vhost-net的多队列](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/vhost-net%E7%9A%84%E5%A4%9A%E9%98%9F%E5%88%97.PNG)

网卡多队列这个功能对于提升虚拟机网卡处理能力，包括每秒处理报文个数（packets per second，pps）和吞吐量（throughput）都有非常大的帮助。

- <u>当客户机中的 virtio-net 网卡只有一个队列时，那么该网卡的中断就只能集中由一个CPU来处理</u>；
- 如果客户机用作一个高性能的 Web 服务器，其网络较为繁忙、网络压力很大，那么只能用单个 CPU 来处理网卡中断就会成为系统瓶颈。
- 当我们开启网卡多队列时，在宿主机上，我们前面已经看到，会有多个 vhost 线程来处理网络后端，同时在客户机中，virtio-net 多队列网卡也可以将网卡中断打散到多个 CPU 上由它们并行处理，从而避免单个 CPU 处理网卡中断可能带来的瓶颈，从而提高整个系统的网络处理能力。

注意：

使用 vhost-net 作为后端处理驱动可以提高网络的性能。不过，对于一些使用 vhost-net 作为后端的网络负载类型，可能使其性能不升反降。**特别是从宿主机到其客户机之间的UDP流量，如果客户机处理接收数据的速度比宿主机发送数据的速度要慢，这时就容易出现性能下降**。在这种情况下，<u>使用 vhost-net 将会使 UDP socket 的接收缓冲区更快地溢出，从而导致更多的数据包丢失</u>。因此，在这种情况下不使用vhost-net，让传输速度稍微慢一点，反而会提高整体的性能  。

在使用 libvirt 时，默认会优先使用 vhost_net 作为网络后端驱动。



### 使用用户态的 vhost-user 作为后端驱动

vhost-net，是为了减少网络数据交换过程中的多次上下文切换，让 guest 与 host kernel 直接通信，从而提高网络性能。

然而，在大规模使用 KVM 虚拟化的云计算生产环境中，通常都会使用 Open vSwitch 或与其类似的 SDN 方案，以便可以更加灵活地管理网络资源。

- 通常在这种情况下，在宿主机上会运行一个虚拟交换机（vswitch）用户态进程，这时如果使用 vhost 作为后端网络处理程序，那么也会存在宿主机上用户态、内核态的上下文切换。
- **vhost-user 的产生就是为了解决这样的问题，它可以让客户机直接与宿主机上的虚拟交换机进程进行数据交换**。

可以理解为，vhost-net 是实现了 vhost 协议的 Linux 内核态模块，而 vhost-user 是实现了 vhost 协议的用户态模块。如果客户机的网络需要与 OpenVSwitch 的虚拟交换机交互，那么使用 vhost-user 的用户态模块，就会减少内核-用户态的切换造成的损失。

vhost-user 可以理解为在用户态实现了 vhost 的一种协议。**vhost-user 协议实现了在同一个宿主机上两个进程建立共享的虚拟队列（virtqueue）所需要的控制平面**。**控制逻辑的信息交换是通过共享文件描述符的UNIX套接字来实现的**；当然，在**数据平面是通过两个进程间的共享内存来实现的**。



## 2. 设备直接分配（VT-d）

## VT-d 原理

在QEMU/KVM中，客户机可以使用的设备大致可分为如下3种类型：

1. **Emulated device**：<u>QEMU纯软件模拟的设备</u>，比如-device rt8139等。
2. **virtio device**：实现<u>VIRTIO</u> API的半虚拟化驱动的设备，比如-device virtio-net-pci等。
3. **PCI device assignment**：<u>PCI设备直接分配</u>。

前两种类型，之前已经都介绍过了，现在介绍第三种方式，**PCI设备直接分配（Device Assignment，或PCI pass-through）**，它**允许将宿主机中的物理PCI（或PCI-E）设备直接分配给客户机完全使用**。Intel定义的I/O虚拟化技术规范为 “Intel(R) Virtualization Technology for Directed I/O”（VT-d），而AMD的I/O虚拟化技术规范为“AMD-Vi”（也叫作IOMMU）

KVM虚拟机支持将宿主机中的PCI、PCI-E设备附加到虚拟化的客户机中，从而让客户机以独占方式访问这个PCI（或PCI-E）设备。

- 通过硬件支持的VT-d技术将设备分配给客户机后，在客户机看来，设备是物理上连接在其PCI（或PCI-E）总线上的
- 客户机对该设备的I/O交互操作和实际的物理设备操作完全一样，不需要（或者很少需要）Hypervisor（即KVM）的参与。

![KVM客户机直接分配PCI-E设备架构](https://github.com/Nevermore12321/LeetCode/blob/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/KVM%E5%AE%A2%E6%88%B7%E6%9C%BA%E7%9B%B4%E6%8E%A5%E5%88%86%E9%85%8DPCI-E%E8%AE%BE%E5%A4%87%E6%9E%B6%E6%9E%84.PNG?raw=true)

VT-d 的优点：

- 设备直接分配让客户机完全占有PCI设备，这样在执行I/O操作时可大量地减少甚至避免了VM-Exit陷入Hypervisor中，极大地提高了I/O性能，几乎可以达到与Native系统中一样的性能。
- 尽管virtio的性能也不错，但VT-d克服了其兼容性不够好和CPU使用率较高的问题

VT-d 的缺点：

- 一台服务器主板上的空间比较有限，允许添加的PCI和PCI-E设备是有限的，如果一台宿主机上有较多数量的客户机，则很难向每台客户机都独立分配VT-d的设备。
- 大量使用VT-d独立分配设备给客户机，导致硬件设备数量增加，这会增加硬件投资成本
- 对于使用VT-d直接分配了设备的客户机，其动态迁移功能将会受限，不过也可以用热插拔或libvirt工具等方式来缓解这个问题

为了克服前两个缺点，可以考虑采用如下两个方案：

- 一是在一台物理宿主机上，仅对I/O（如网络）性能要求较高的少数客户机使用VT-d直接分配设备（如网卡），而对其余的客户机使用纯模拟（emulated）或使用virtio，以达到多个客户机共享同一个设备的目的；
- 二是对于网络I/O的解决方法，可以选择SR-IOV，使一个网卡产生多个独立的虚拟网卡，将每个虚拟网卡分别分配给一个客户机使用

### VFIO简介

VFIO（Virtual Function IO ）最大的改进就是隔离了设备之间的DMA和中断，以及对IOMMU Group的支持，从而有了更好的安全性。IOMMU Group可以认为是对PCI设备的分组，每个group里面的设备被视作IOMMU可以操作的最小整体；换句话说，同一个IOMMU Group里的设备不可以分配给不同的客户机。新的VFIO架构也做到了平台无关，有更好的可移植性

### SR-IOV技术

<u>前面介绍的普通VT-d技术实现了设备直接分配，尽管其性能非常好，但是它的一个物理设备只能分配给一个客户机使用</u>。

**为了实现多个虚拟机能够共享同一个物理设备的资源，并且达到设备直接分配的性能，PCI-SIG 组织发布了SR-IOV（Single Root I/O Virtualization and Sharing）规范**，该规范定义个了一个标准化的机制，**用以原生地支持实现多个共享的设备（不一定是网卡设备）**。不过，<u>目前SR-IOV（单根I/O虚拟化）最广泛的应用还是在以太网卡设备的虚拟化方面</u>

SR-IOV中引入的两个新的功能（function）类型：

- **Physical Function（PF，物理功能）**：拥有包含SR-IOV扩展能力在内的所有完整的PCI-e功能，其中SR-IOV能力使PF可以配置和管理SR-IOV功能。简言之，PF就是一个普通的PCI-e设备（带有SR-IOV功能），可以放在宿主机中配置和管理其他VF，它本身也可以作为一个完整独立的功能使用。
- **Virtual Function（VF，虚拟功能）**：由PF衍生而来的“轻量级”的PCI-e功能，包含数据传送所必需的资源，但是仅谨慎地拥有最小化的配置资源。简言之，VF通过PF的配置之后，可以分配到客户机中作为独立功能使用。

可以理解为：PF 为一个物理网卡，可以抽象成多个虚拟网卡 VF。

<u>SR-IOV为客户机中使用的VF提供了独立的内存空间、中断、DMA流，从而不需要Hypervisor介入数据的传送过程。SR-IOV架构设计的目的是允许一个设备支持多个VF，同时也尽量减小每个VF的硬件成本。</u>

SR-IOV 的总体架构为：

![Intel以太网卡的SR-IOV总体架构](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/Intel%E4%BB%A5%E5%A4%AA%E7%BD%91%E5%8D%A1%E7%9A%84SR-IOV%E6%80%BB%E4%BD%93%E6%9E%B6%E6%9E%84.PNG)

注意：

- 一个具有SR-IOV功能的设备能够被配置为在PCI配置空间（configuration space）中呈现出多个Function（包括一个PF和多个VF），每个VF都有自己独立的配置空间和完整的BAR（Base Address Register，基址寄存器）。
- Hypervisor通过将VF实际的配置空间映射到客户机看到的配置空间的方式，实现将一个或多个VF分配给一个客户机。通过Intel VT-x和VT-d等硬件辅助虚拟化技术提供的内存转换技术，允许直接的DMA传输去往或来自一个客户机，从而绕过了Hypervisor中的软件交换机（software switch）。
- 每个VF在同一个时刻只能被分配到一个客户机中，因为VF需要真正的硬件资源（不同于emulated类型的设备）。在客户机中的VF，表现给客户机操作系统的就是一个完整的普通的设备。

在KVM中，可以将一个或多个VF分配给一个客户机，客户机通过自身的VF驱动程序直接操作设备的VF而不需要Hypervisor（即KVM）的参与，结构如下：

![在KVM中使用SR-IOV功能示意图](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/kvm/%E5%9C%A8KVM%E4%B8%AD%E4%BD%BF%E7%94%A8SR-IOV%E5%8A%9F%E8%83%BD%E7%A4%BA%E6%84%8F%E5%9B%BE.PNG)

为了让SR-IOV工作起来，需要硬件平台支持Intel VT-x和VT-d（或AMD的SVM和IOMMU）硬件辅助虚拟化特性，还需要有支持SR-IOV规范的设备，当然也需要
QEMU/KVM的支持。



使用SR-IOV主要有如下3个优势：

- 真正实现了设备的共享（多个客户机共享一个SR-IOV设备的物理端口）。
- 接近于原生系统的高性能（比纯软件模拟和virtio设备的性能都要好）。
- 相比于VT-d，SR-IOV可以用更少的设备支持更多的客户机，可以提高数据中心的空间利用率。

SR-IOV的不足之处有如下两点：

- 对设备有依赖，只有部分PCI-E设备支持SR-IOV。
- 使用SR-IOV时，不方便动态迁移客户机。

## 3. 热插拔

**热插拔（hot plugging）**即“带电插拔”，指可以在计算机运行时（不关闭电源）插上或拔除硬件。

热插拔最早出现在服务器领域，目的是提高服务器扩展性、灵活性和对灾难的及时恢复能力。

实现热插拔需要有几方面支持：总线电气特性、主板BIOS、操作系统和设备驱动。

- 目前，在服务器硬件中，可实现热插拔的部件主要有SATA硬盘（IDE不支持热插拔）、CPU、内存、风扇、USB、网卡等。
- 在KVM虚拟化环境中，在不关闭客户机的情况下，也可以对客户机的设备进行热插拔。目前，主要支持PCI设备、CPU、内存的热插拔。