# OpenStack Horizon Dashboard 部署到 Httpd

[toc]

## httpd 配置文件

- `/etc/httpd/conf/httpd.conf`：主配置文件；
- `/etc/httpd/conf.modules.d/*.conf`：模块配置文件；
- `/etc/httpd/conf.d/*.conf`：辅助配置文件；
- `/var/log/httpd/access.log`：访问日志；
- `/var/log/httpd/error_log`：错误日志；
- `/var/www/html/`：用户的 html 项目代码。


如果要将各个不同的项目部署在 httpd 中，可以在 conf.d 目录中添加不同的项目配置。

安装 httpd: `yum install httpd`

## Django 项目部署

### WSGI 协议

Django 项目是 WSGI 协议的 Web 框架。

*****WSGI**：全称是Web Server Gateway Interface，WSGI 是一种规范，用来描述 web server 如何与 web application 通信的规范。
- 是 web 底层跟 application 解耦的协议
- web服务器使用WSGI协议来调用application称为WSGI server
- 常见的web server（如Nginx，Apache）无法与web application（如Flask，django、tornado）直接通信，需要WSGI server作为桥梁。


WSGI协议主要包括server和application两部分：
- WSGI server负责从客户端接收请求，将request转发给application，将application返回的response返回给客户端；
- WSGI application接收由server转发的request，处理请求，并将处理结果返回给server。
	- application中可以包括多个栈式的中间件(middlewares)，这些中间件需要同时实现server与application，因此可以在WSGI服务器与WSGI应用之间起调节作用
	- 对服务器来说，中间件扮演应用程序(执行程序)，对应用程序来说，中间件扮演服务器(WSGI服务器)。

![wsgi基本流程](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/OpenStack/wsgi%E5%9F%BA%E6%9C%AC%E6%B5%81%E7%A8%8B.png)


### WSGI 协议原理

WSGI的工作原理分为服务器层和应用程序层：
- 服务器层：
	- 将来自socket的数据包解析为http，调用application，给application提供环境信息environ，这个environ包含wsgi自身的信息（host，post，进程模式等），还有client的header和body信息。
	- 同时还给application提供一个start_response的回调函数，这个回调函数主要在应用程序层进行响应信息处理。
- 应用程序层：
	- 在WSGI提供的start_response，生成header，body和status后将这些信息socket send返回给客户端。
	
WSGI 原理流程图如下：
![WSGI原理图](https://raw.githubusercontent.com/Nevermore12321/LeetCode/blog/%E4%BA%91%E8%AE%A1%E7%AE%97/OpenStack/wsgi%E5%8E%9F%E7%90%86.png)

可以看到，wsgi server已经完成了底层http的解析和数据转发等一系列网络底层的实现，开发者可以更加专注于开发web application。


## Django 打包

### 1. Django 配置
配置 settings 模块

当 WSGI 服务器加载应用时，Django 需要导入配置模块——完整定义应用的地方。
- Django 利用 DJANGO_SETTINGS_MODULE 环境变量来定位合适的配置模块。它必须包含到配置模块的路径。
- 开发环境和生产环境可以配置不同值；这都取决于你是如何组织配置的。
- 若未设置该变量， wsgi.py 默认将其设置为 mysite.settings， mysite 即工程名字。这就是 runserver 默认的发现默认配置行为。

例如，在 Horizon 的 openstack_dashboard/wsgi.py 文件中：
```python
import os
import sys

from django.core.wsgi import get_wsgi_application

# Add this file path to sys.path in order to import settings
sys.path.insert(0, os.path.normpath(os.path.join(
    os.path.dirname(os.path.realpath(__file__)), '..')))
os.environ['DJANGO_SETTINGS_MODULE'] = 'openstack_dashboard.settings'
sys.stdout = sys.stderr

application = get_wsgi_application()
```
这里配置了 openstack_dashboard 这个 app

### 2. 执行打包命令
在 Django 的根目录下，执行打包以及压缩的命令
```bash
# 打包
python manage.py collectstatic -l
# 压缩
python manage.py compress -f

```

打包后，会发现在根目录下，多了 static 目录。


## httpd 配置

### 1. 将需要部署的项目拷贝到指定目录

httpd 的默认配置在 `/etc/httpd/conf/httpd.conf` 中，配置了 
```
DocumentRoot "/var/www/html"
```
也就是指定了 httpd 的默认项目根目录

这里我们用 horizon 项目为例，需要将 horizon 项目拷贝到 `/var/www/html` 目录下：
```bash
[root@172-1-0-1 html]# pwd
/var/www/html
[root@172-1-0-1 html]# ls
horizon  vue-demo
[root@172-1-0-1 html]# ls horizon/
babel-django.cfg    geckodriver.log   manage.py            package.json         README.rst        setup.cfg              test-shim.js
babel-djangojs.cfg  HACKING.rst       MANIFEST.in          package-lock.json    releasenotes      setup.py               tools
bindep.txt          horizon           node_modules         playbooks            reno.yaml         static                 tox.ini
CONTRIBUTING.rst    horizon.egg-info  openstack_auth       plugin-registry.csv  requirements.txt  test_reports
doc                 LICENSE           openstack_dashboard  __pycache__          roles             test-requirements.txt

```

### 2. 配置 httpd 

1. 配置 httpd 支持 wsgi 协议

首先，使用 pip 安装 mod_wsgi 包
```bash
pip3 install mod_wsgi
```

然后，查看 python 的路径，以及 mod_wsgi 的位置：
```
[root@172-1-0-1 html]# mod_wsgi-express module-config
LoadModule wsgi_module "/usr/local/lib64/python3.6/site-packages/mod_wsgi/server/mod_wsgi-py36.cpython-36m-x86_64-linux-gnu.so"
WSGIPythonHome "/usr"
```

最后，在 `/etc/httpd/conf/httpd.conf` 基本配置文件中，配置 httpd 支持 wsgi，其实就是使用 uwsgi 监听 django 的项目，而static静态文件直接去 static 目录中拿
```
LoadModule wsgi_module "/usr/local/lib64/python3.6/site-packages/mod_wsgi/server/mod_wsgi-py36.cpython-36m-x86_64-linux-gnu.so"
WSGIPythonHome "/usr"                                                          
Include conf.modules.d/*.conf  
```

2. 部署 horizon 到 httpd 

在`/etc/httpd/conf.d/` 目录中，新建文件，django.conf，添加内容如下 ：
```
listen 34568                                                                                                                            
<VirtualHost *:34568>                                                                                                                   
                                                                                                                                        
        WSGIDaemonProcess dashboard                                                                                                     
        WSGIProcessGroup dashboard                                                                                                      
                                                                                                                                        
        WSGIApplicationGroup %{GLOBAL}                                                                                                  
                                                                                                                                        
        WSGIScriptAlias / /var/www/html/horizon/openstack_dashboard/wsgi/django.wsgi                                                    
        Alias /static /var/www/html/horizon/static
        <Directory /var/www/html/horizon/openstack_dashboard/wsgi/django.wsgi>
                Options All
                AllowOverride All
                Require all granted
        </Directory>

        <Directory /var/www/html/horizon/static>
                Options All
                AllowOverride All
                Require all granted
        </Directory>

</VirtualHost>

```
3. 重启 httpd 服务
```
systemctl restart httpd
```


