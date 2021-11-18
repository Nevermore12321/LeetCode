# windows + centos8 双系统教程

[toc]

## 1. 制作 U 盘启动

使用 easyboot 制作 U 盘启动，

该 U 盘启动没有启动盘，只是挂在了 U 的文件，通过 U 盘中的 ISO 文件开始系统的安装

也就是下载系统的 ISO 文件拷贝导 easyboot 的文件系统中，这里**需要下载 ubuntu 和 centos 的两个镜像 ISO 文件**。

因为：

windows 的文件系统格式为 nfs，centos8 需要识别本地的 ext4 格式的文件系统中的 ISO 镜像文件，开启安装，否则需要进行网络在线安装，速度非常慢。

也可以制作连个启动，先使用 Ubuntu live 拷贝 iso 文件，然后使用 Centos8 安装系统。

## 2. 磁盘管理

在 windows 下，使用 **磁盘管理** 工具，操作两步：

1. 压缩 C 盘，在 C 盘上右键 -> 压缩卷，可以将 C 盘 压缩成两部分，一个供 windows 使用，一个供 Linux 使用
2. 如果 windows 的其他盘没有用的话，可以删除，搞成一块盘，然后再分成两个，一个供 windows 使用，一个供 Linux 使用

最终的效果如下：



## 3. 修改 Windows BIOS 配置

需要修改两个配置：

1. 修改 Security Boot ，配置为 disable
2. 修改 BIOS 启动顺序为 U 盘启动

修改好后，保存退出，并启动



## 4. 进入 Ubuntu live 修改磁盘格式

修改好 BIOS 后，进入 U 盘启动，启动后，会看到 U 盘中的所有内容。

首先选择 Ubuntu-desktop 的 ISO 镜像，进入导 Ubuntu live 系统，注意是 live 系统，live 系统是启动在 U 盘中的临时系统，通过该 Ubuntu live 将 windows 的磁盘格式，转成 EXT4 .

 主要步骤就是：

1. 进入 U 盘启动，选择 Ubuntu-desktop ISO 镜像
2. 进入系统后，选择 Live 版本，进入临时 Ubuntu 系统
3. 选择左上角的 Activities，进入搜索界面，然后搜索 disks ，进入 disks 
4. 选择想要给 linux 的磁盘，点击删除（如果没有 free 需要删除），在重新添加一块，选择最大，直接一路确定，最终就是 EXT4 格式，将创建好后的磁盘挂载，就是 播放按钮，挂载后，点击打开挂载后的目录
5. 在 disks 中找到 U 盘的盘符，然后查看挂载目录，点击挂载的目录，将 centos8 的 ISO 镜像 copy 到 EXT4 磁盘中。
6. 拷贝完成后，重启机器



## 5. 进入 Centos8 的安装

重启后，同样还是 U 盘启动，进入 easyboot 的 U 盘目录下，选择 centos8 的 ISO 镜像，进入 centos8 的安装向导。

注意：需要选择 ISO 文件的目录，以及磁盘分区时需要选择 除 windows 外的磁盘

主要步骤：

1. 进入 U 盘启动后，选择 Centos8 ISO 镜像
2. 进入 centos8 安装向导后，选择 Centos8 的 ISO 文件，需要选择到 EXT4 这块磁盘中的 Centos8 ISO 文件。
3. 磁盘分区时，需要选择通过 Windows C 盘 压缩后的剩下的盘符做为 Linux 的安装盘，这里不能装错。
4. 这些结束后，就可以下一步，开始安装。



 