# Pypi Server 搭建 pip 私有仓库

安装步骤：

1. 安装 python3
2. 安装 pip3

```bash
yum install python39
```

3. 安装 pypiserver

```bash
pip3 install pypiserver
```

4. 启动 pypiserver

```bash
nohup pypi-server -p 8080 /packages &
```

5. 创建保存 Packages 的目录，并且下载需要的 python 软件包到该目录下
   1. 如果有一些包找不到 ，原因：没有提供符合条件的二进制包
   2. 可以使用非二进制包安装   --no-binary=:all: package_name

```bash
mkdir /packages
pip3 download \
    --only-binary=:all: \ 		# 对于包以及包的依赖包，都不使用二进制
    --platform linux_x86_64 \  	# 指定系统环境
    -d /packages \				# 下载的文件的存储目录
    -r requirement.txt    		# 指定要下载的包
```

6. 