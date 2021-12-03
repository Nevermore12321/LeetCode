# yum repo 搭建合集

[toc]

在 CentOS 系统下关于使用国内 Yum 源，本地 Yum 源，私有 Yum 源的整体解决方案，搞定“软件依赖关系问题”！

## 工具清单

- **rpm**：用于安装、卸载和管理 rpm 包的命令工具。CentOS 默认安装。
- **dnf/yum**：用于从 yum 源下载 rpm 包、自动处理包依赖关系，并且能够一次安装所有所依赖 rpm 包的命令工具。CentOS8 默认可以使用 dnf/yum 命令，CentOS7 及以下版本默认只能使用 yum 命令。
- **reposync**：用于同步互联网 yum 源的 rpm 包到本地磁盘中的命令工具。通过 yum 源下载 yum-utils 安装。、
- **createrepo**：用于扫描本地磁盘中的 rpm 包并生成元数据文件，建立本地 yum 源的命令工具。建成本地 yum 源后，服务器可以通过 file 协议使用本地 yum 源。
- **modifyrepo：**用于导入 yum 源模块文件的命令工具。
- **wget**：用于从网络上自动下载文件的工具。
- **python3**：通过 python3 http.server 工具，将本地 yum 源发布到局域网中，局域网中其他服务器可通过 http 协议使用已发布的私有 yum 源。CentOS8 默认安装，如未安装可通过 yum 源下载安装。

## 使用国内 yum 源

在一个连接互联网的服务器上安装完 Centos 系统后，从国外的 yum 源下载文件很慢，将发行方提供的国外 yum 源改为第三方提供的国内 yum 源。

**步骤：**

1. **备份现有的 yum 源配置文件(root)。**

```bash
[root@localhost ~]# cp -r /etc/yum.repos.d /etc/yum.repos.d.backup
```

2.  **备份后，将原先 yum 配置文件删除掉，根据安装的 Centos 版本，下载对应的 阿里源 yum 配置文件**

```bash
[root@localhost ~]# cd /etc/yum.repos.d
[root@localhost yum.repos.d]# rm -rf *
# CentOS 5
[root@localhost yum.repos.d]# wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-5.repo
# CentOS 6
[root@localhost yum.repos.d]# wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-6.repo
# CentOS 7
[root@localhost yum.repos.d]# wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
# CentOS 8
[root@localhost yum.repos.d]# wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-8.repo
```

3. **刷新 yum 源缓存。使新的 yum 源配置生效**

```bash
[root@localhost ~]# dnf clean all
Repository extras is listed more than once in the configuration
33 文件已删除
[root@localhost ~]# dnf makecache
Repository extras is listed more than once in the configuration
CentOS-8 - Base - mirrors.aliyun.com                                        2.9 MB/s | 3.5 MB     00:01
CentOS-8 - Extras - mirrors.aliyun.com                                       19 kB/s |  10 kB     00:00
CentOS-8 - AppStream - mirrors.aliyun.com                                   4.7 MB/s | 8.1 MB     00:01
CentOS Linux 8 - AppStream                                                  4.0 MB/s | 8.1 MB     00:02
CentOS Linux 8  BaseOS                                                     722 kB/s | 3.5 MB     00:04
元数据缓存已建立。
```

4. **查询启用的 yum 源清单**

```bash
[root@localhost ~]# dnf repolist -v
Repo-id            : AppStream
Repo-name          : CentOS-8 - AppStream - mirrors.aliyun.com
Repo-revision      : 8.5.2111
Repo-distro-tags      : [cpe:/o:centos:centos:8]:  , 8, C, O, S, e, n, t
Repo-updated       : 2021年11月18日 星期四 05时10分03秒
Repo-pkgs          : 6,353
Repo-available-pkgs: 5,469
Repo-size          : 9.8 G
Repo-baseurl       : https://mirrors.aliyun.com/centos/8/AppStream/x86_64/os/,
                   : http://mirrors.aliyuncs.com/centos/8/AppStream/x86_64/os/,
                   : http://mirrors.cloud.aliyuncs.com/centos/8/AppStream/x86_64/os/
Repo-expire        : 172,800 秒 （最近 2021年11月18日 星期四 14时08分50秒）
Repo-filename      : /etc/yum.repos.d/CentOS-Base.repo

Repo-id            : appstream
Repo-name          : CentOS Linux 8 - AppStream
Repo-revision      : 8.5.2111
Repo-distro-tags      : [cpe:/o:centos:centos:8]:  , 8, C, O, S, e, n, t
Repo-updated       : 2021年11月18日 星期四 05时10分03秒
Repo-pkgs          : 6,353
Repo-available-pkgs: 5,469
Repo-size          : 9.8 G
Repo-mirrors       : http://mirrorlist.centos.org/?release=8&arch=x86_64&repo=AppStream&infra=stock
Repo-baseurl       : http://ftp.sjtu.edu.cn/centos/8.5.2111/AppStream/x86_64/os/ (9 more)
Repo-expire        : 172,800 秒 （最近 2021年11月18日 星期四 14时08分53秒）
Repo-filename      : /etc/yum.repos.d/CentOS-Linux-AppStream.repo
...
```

5. **使用国内阿里云 yum 源进行系统更新**

```bash
[root@localhost ~]# dnf update -y
```

## 逻辑隔离局域网搭建私有yum源

![逻辑隔离局域网拓扑图](https://github.com/Nevermore12321/LeetCode/blob/blog/%E5%B7%A5%E4%BD%9C/repos/%E9%80%BB%E8%BE%91%E9%9A%94%E7%A6%BB%E5%B1%80%E5%9F%9F%E7%BD%91%E6%8B%93%E6%89%91%E5%9B%BE.jpg?raw=true)

在局域网中，只有"私有 yum 源服务器"可以通过防火墙访问互联网的 yum 源，其他服务器不能访问互联网，使用"私有 yum 源服务器"。

**"私有 yum 源服务器"先同步互联网的 yum 源到本地，建立本地 yum 源，再为局域网中其他服务器提供私有 yum 源**。

步骤：

1. **使用国内 yum 源可以大幅提高下载速度。详见 使用国内 yum 源 章节。**
2. **创建本地 yum 源主目录（也就是存放下载包的目录）。本地 yum 源主目录所在磁盘的可用空间建议 ≥ 40G**

```bash
[root@localhost ~]# mkdir -p /repos/centos-8-x86_64
```

3. **安装 reposync、createrepo 工具。**

```bash
[root@localhost ~]# dnf install yum-utils createrepo
```

4. **查询启用的 yum 源清单，获取仓库 id。**

```bash
[root@localhost ~]# dnf repolist
仓库 id                               仓库名称
AppStream                             CentOS-8 - AppStream - mirrors.aliyun.com
base                                 CentOS-8 - Base - mirrors.aliyun.com
extras                                CentOS-8 - Extras - mirrors.aliyun.com
```

5. **使用 reposync 命令同步指定 yum 源标识的 rpm 包到本地 yum 源主目录**
   1. 若下载过程中发生异常，可重新执行该命令，对于已经下载的rpm包会自动跳过。
   2. 参数 "-n" 表示只下载最新版本得 rpm 包，"-p" 表示存储目录，"--repoid" 表示要下载的仓库 id。

```bash
[root@localhost ~]# reposync --repoid=AppStream --repoid=base --repoid=extras -p /repos/centos-8-x86_64
```

6. **使用 wget 命令下载阿里云 yum 源的校验文件到本地 yum 源主目录**
   1. 这里的校验文件，就是 repo 文件中 gpgcheck=1 时，需要校验 gpgkey 
   2. 阿里源的校验文件在：http://mirrors.aliyun.com/centos/RPM-GPG-KEY-CentOS-Official

```bash
[root@localhost ~]# wget -O /repos/centos-8-x86_64/RPM-GPG-KEY-CentOS-Official http://mirrors.aliyun.com/centos/RPM-GPG-KEY-CentOS-Official
```

7. **下载完后，使用 createrepo 命令扫描本地 yum 源目录中的 rpm 包并生成元数据文件，建立本地 yum 源**
   1. reposync 命令只是将所有的 rpm 包同步到本地
   2. createrepo 会生成元数据，就是 repo 的 xml 索引

```bash
[root@localhost ~]# createrepo -p /repos/centos-8-x86_64/AppStream

[root@localhost ~]# createrepo -p /repos/centos-8-x86_64/base

[root@localhost ~]# createrepo -p /repos/centos-8-x86_64/extras
```

8. **使用 wget 命令下载阿里云 yum 源的 modules 文件。有的 yum 源需要引用 modules 文件才能正常使用，比如：AppStream。没有 modules 文件的 yum 源可以略过此步骤，如：BaseOS 和 extras**

   - 首先确定 yum 源是否需要 modules 文件。在浏览器中访问阿里云 yum 源的元数据目录（repodata）中的"repomd.xml"文件。文件地址为：http://mirrors.aliyun.com/centos/8/AppStream/x86_64/os/repodata/repomd.xml
   - 直接 ctrl + f 搜索 modules
   - 如果有下面这段，说明需要 modules 文件：

   ```bash
   <data type="modules">
       <checksum type="sha256">7d90930238ed179faa4f32c968cece05d35d264d24c36fec2961b098d04d7c2b</checksum>
       <open-checksum type="sha256">c9b249adc568b221cd5858624313dc91c147b23f95f70c1ae16b35a485e2e4ae</open-checksum>
       <location href="repodata/7d90930238ed179faa4f32c968cece05d35d264d24c36fec2961b098d04d7c2b-modules.yaml.xz"/>
       <timestamp>1637183389</timestamp>
       <size>61344</size>
       <open-size>535971</open-size>
   </data>
   ```

   - 在 yum 源的元数据目录（repodata）中找到`<data type="modules">`节点下的`<location />`节点中"href"属性的 "*.yaml.xz" 文件并下载到本地 yum 源目录，并重命名为 "modules.yaml.xz"（文件名不能更改）：

   ```bash
   [root@localhost ~]# wget -O /repos/centos-8-x86_64/AppStream/modules.yaml.xz http://mirrors.aliyun.com/centos/8/AppStream/x86_64/os/repodata/7d90930238ed179faa4f32c968cece05d35d264d24c36fec2961b098d04d7c2b-modules.yaml.xz
   
   ```

   - 将下载到本地yum源目录下的"modules.yaml.gz"文件解压缩：

   ```bash
   [root@localhost ~]# cd /repos/centos-8-x86_64/AppStream
   [root@localhost AppStream]# xz -d modules.yaml.xz
   ```

9. **使用 modifyrepo 命令导入 modules 文件到本地 yum 源的元数据目录（repodata）。**

```bash
[root@localhost ~]# modifyrepo /repos/centos-8-x86_64/AppStream/modules.yaml /repos/centos-8-x86_64/AppStream/repodata
```

10. **设置定时更新程序。编写 Shell 脚本，配置系统定时器定时执行更新或创建本地 yum 源。**

    - 在 /repos 目录下创建 `update-repos.sh` 文件，文件内容如下：

    ```shell
    #!/bin/bash
    
    uid=$(id -u)
    
    if [ $uid != "0" ]; then
    
      echo "请使用root权限运行"
    
      exit 0
    
    fi
    
    echo CentOS8 YUM源-同步开始
    
    if [ ! -d "/repos/centos-8-x86_64" ];then
    
      mkdir -p /repos/centos-8-x86_64
    
      wget -O /repos/centos-8-x86_64/RPM-GPG-KEY-CentOS-Official https://mirrors.aliyun.com/centos/RPM-GPG-KEY-CentOS-Official
    
    fi
    
    reposync --repoid=AppStream --repoid=BaseOS --repoid=extras -p /repos/centos-8-x86_64
    
    echo CentOS8 YUM源-同步结束
    
    read -p "是否更新本地YUM源数据(y/N)？" input
    
    if [ $input = "y" -o $input = "Y" ]; then
    
       echo
    
    else
    
      exit 0
    
    fi
    
    echo CentOS8 本地YUM源元数据-更新开始
    
    if [ -d "/repos/centos-8-x86_64/AppStream/repodata" ];then
    
      rm -rf /repos/centos-8-x86_64/AppStream/repodata
    
    fi
    
    if [ -d "/repos/centos-8-x86_64/AppStream/.repodata" ];then
    
      rm -rf /repos/centos-8-x86_64/AppStream/.repodata
    
    fi
    
    if [ -d "/repos/centos-8-x86_64/BaseOS/repodata" ];then
    
      rm -rf /repos/centos-8-x86_64/BaseOS/repodata
    
    fi
    
    if [ -d "/repos/centos-8-x86_64/BaseOS/.repodata" ];then
    
      rm -rf /repos/centos-8-x86_64/BaseOS/.repodata
    
    fi
    
    if [ -d "/repos/centos-8-x86_64/extras/repodata" ];then
    
      rm -rf /repos/centos-8-x86_64/extras/repodata
    
    fi
    
    if [ -d "/repos/centos-8-x86_64/extras/.repodata" ];then
    
      rm -rf /repos/centos-8-x86_64/extras/.repodata
    
    fi
    
    createrepo -p /repos/centos-8-x86_64/AppStream
    
    createrepo -p /repos/centos-8-x86_64/BaseOS
    
    createrepo -p /repos/centos-8-x86_64/extras
    
    echo CentOS8 本地YUM源元数据-更新结束
    
    
    
    echo CentOS8 本地YUM源模块-配置开始
    
    if [ -f "/repos/centos-8-x86_64/AppStream/modules.yaml" ];then
    
      modifyrepo /repos/centos-8-x86_64/AppStream/modules.yaml /repos/centos-8-x86_64/AppStream/repodata
    
    fi
    
    echo CentOS8 本地YUM源模块-配置结束
    
    
    
    echo CentOS8 本地YUM源更新完成.时间：$(date +"%Y-%m-%d %T")
    
    echo CentOS8 本地YUM源更新完成.时间：$(date +"%Y-%m-%d %T") >> update-repos.log
    ```

    - 编写完成后为文件增加可执行权限

    ```bash
    [root@localhost ~]# chmod 755 /repos/update-repos.sh
    ```

    - 设置定时执行程序，设置内容为每天凌晨 00:00 开始执行

    ```bash
    [root@localhost ~]# crontab -u root -e
    0 0 * * * /repos/update-repos.sh
    ```

11. **将当前目录切换为"/repos"，使用 python3 http.server 发布私有yum源。**

```bash
[root@localhost ~]# nohup python3 -m http.server 80 &
```

12. **打开防火墙端口**

```bash
[root@localhost ~]# firewall-cmd --zone=public --add-port=80/tcp --permanent
[root@localhost ~]# firewall-cmd --reload
```

13. **在浏览器中查看已发布的yum源。**
    1. 可以在浏览器中访问 http://localhost 查看 yum 源的发布



## 物理隔离局域网搭建私有 yum 源和使用本地 yum 源

![物理隔离局域网拓扑图](![物理隔离局域网拓扑图.jpg](https://github.com/Nevermore12321/LeetCode/blob/blog/%E5%B7%A5%E4%BD%9C/repos/%E7%89%A9%E7%90%86%E9%9A%94%E7%A6%BB%E5%B1%80%E5%9F%9F%E7%BD%91%E6%8B%93%E6%89%91%E5%9B%BE.jpg?raw=true))

局域网的"私有YUM源服务器"及其他服务器都不能访问互联网。"代理YUM源服务器"可以访问互联网，并与局域网物理隔离。

"代理YUM源服务器"先同步互联网的yum源到本地挂载的可移动存储上，建立本地yum源，再将存有本地yum源的可移动存储拆卸，安装到局域网的"私有YUM源服务器"。

"私有YUM源服务器"通过挂载可移动存储到本地，建立本地yum源，再为局域网中其他服务器提供私有yum源。

步骤:

1. **使用国内 yum 源可以大幅提高下载速度。详见 使用国内 yum 源 章节。**

2. **搭建本地 yum 源。详见 逻辑隔离局域网搭建私有 yum 源 章节中的第 3-10 步。**

3. **存有本地 yum 源制作成 ISO 镜像文件后上传到"私有 yum 源服务器"的虚拟化平台中。**

   1. CentOS8 中制作 ISO 文件的命令是（注意要加 -J ，否则文件名会被截断）:

   ```bash
   [root@localhost ~]# genisoimage -J -o /repos.iso /repos/centos-8-x86_64
   [root@localhost ~]# genisoimage -joliet-long -o /repos.iso /repos/centos-8-x86_64
   ```

4. 将生成的 repos.iso 文件，拷贝到 **私有 yum 源服务器** ，然后在 私有 yum 源服务器 配置使用该 iso 文件的 yum 源

   - **挂载 iso 文件到指定的目录**

     ```bash
     [root@localhost ~]# mkdir repos
     [root@localhost ~]# mount /repos.iso /repos
     ```

   - **备份现有的 yum 源配置文件。yum 源配置文件是"/etc/yum.repos.d"目录下的"\*.repo"文件，备份该目录，以便于操作失败时恢复。**

     ```bash
     [root@localhost ~]# cp -r /etc/yum.repos.d /etc/yum.repos.d.bak
     ```

   - **修改"/etc/yum.repos.d"目录下的"CentOS-AppStream.repo"、"CentOS-Base.repo"和"CentOS-Extras.repo"文件。改为本地 yum 源。**三个文件分别为：

     ```bash
     [AppStream]
     
     name=CentOS-$releasever- AppStream
     
     # /repos目录是"私有YUM源服务器"的可移动磁盘挂载目录。
     
     baseurl=file:///repos/centos-$releasever-$basearch/AppStream
     
     gpgcheck=1
     
     enabled=1
     
     gpgkey=file:///repos/centos-$releasever-$basearch/RPM-GPG-KEY-CentOS-Official
     ```

     ```bash
     [BaseOS]
     
     name=CentOS-$releasever- Base
     
     # /repos目录是"私有YUM源服务器"的可移动磁盘挂载目录。
     
     baseurl=file:///repos/centos-$releasever-$basearch/BaseOS
     
     gpgcheck=1
     
     enabled=1
     
     gpgkey=file:///repos/centos-$releasever-$basearch/RPM-GPG-KEY-CentOS-Official
     ```

     ```bash
     [extras]
     
     name=CentOS-$releasever- Extras
     
     # /repos目录是"私有YUM源服务器"的可移动磁盘挂载目录。
     
     baseurl=file:///repos/centos-$releasever-$basearch/extras
     
     gpgcheck=1
     
     enabled=1
     
     gpgkey=file:///repos/centos-$releasever-$basearch/RPM-GPG-KEY-CentOS-Official
     ```

   - **刷新 yum 源缓存。使新的 yum 源配置生效。**

     ```bash
     [root@localhost ~]# dnf clean all
     [root@localhost ~]# dnf makecache
     ```

     
