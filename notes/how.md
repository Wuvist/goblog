# 如何永久免费的运行一个网站？

除却域名的年费，我的博客站点是运行在云服务器上，如果没有意外，维护的费用应该是零。

## 云主机

云服务器我使用的是[Google Cloud](https://cloud.google.com/free/)，谷歌云应该是目前唯一一个承诺提供`永远免费 Always Free`选项的云服务器提供商，其它云服务商应该只提供12个月免费使用选项给**新用户**。

当然，谷歌云提供的`永远免费`选项仅能是使用一个性能超弱的实例：

* 处于美国的机房
* 单个共享的CPU核心
* 600M的内存
* 30G的存储

性能可能还比不上一个树莓派，但也已经足够让我用来跑一个网站。

## 环境

操作系统我为方便使用了：Ubuntu 18.04 LTS

### 数据库

数据库因故使用 MySql 8.0，但这直接在`apt-get install mysql-server`之后报错，看`/var/log/mysql/error.log`，显示：

```
2018-08-12T17:55:07.393408Z 0 [System] [MY-013170] [Server] /usr/sbin/mysqld (mysqld 8.0.12) initializing of server has completed
2018-08-12T17:55:11.130992Z 0 [System] [MY-010116] [Server] /usr/sbin/mysqld (mysqld 8.0.12) starting as process 3112
2018-08-12T17:55:11.782754Z 0 [ERROR] [MY-012681] [InnoDB] InnoDB: mmap(137428992 bytes) failed; errno 12
2018-08-12T17:55:11.782817Z 1 [ERROR] [MY-012956] [InnoDB] InnoDB: Cannot allocate memory for the buffer pool
2018-08-12T17:55:11.782832Z 1 [ERROR] [MY-012930] [InnoDB] InnoDB: Plugin initialization aborted with error Generic error.
2018-08-12T17:55:11.782886Z 1 [ERROR] [MY-010334] [Server] Failed to initialize DD Storage Engine
2018-08-12T17:55:11.784009Z 0 [ERROR] [MY-010020] [Server] Data Dictionary initialization failed.
2018-08-12T17:55:11.784055Z 0 [ERROR] [MY-010119] [Server] Aborting
2018-08-12T17:55:11.789921Z 0 [System] [MY-010910] [Server] /usr/sbin/mysqld: Shutdown complete (mysqld 8.0.12)  MySQL Community Server - GPL.
2018-08-12T17:56:06.388162Z 0 [System] [MY-010116] [Server] /usr/sbin/mysqld (mysqld 8.0.12) starting as process 3276
2018-08-12T17:56:07.253414Z 0 [ERROR] [MY-012681] [InnoDB] InnoDB: mmap(137428992 bytes) failed; errno 12
2018-08-12T17:56:07.253518Z 1 [ERROR] [MY-012956] [InnoDB] InnoDB: Cannot allocate memory for the buffer pool
2018-08-12T17:56:07.253535Z 1 [ERROR] [MY-012930] [InnoDB] InnoDB: Plugin initialization aborted with error Generic error.
2018-08-12T17:56:07.253564Z 1 [ERROR] [MY-010334] [Server] Failed to initialize DD Storage Engine
2018-08-12T17:56:07.255109Z 0 [ERROR] [MY-010020] [Server] Data Dictionary initialization failed.
2018-08-12T17:56:07.255993Z 0 [ERROR] [MY-010119] [Server] Aborting
2018-08-12T17:56:07.320087Z 0 [System] [MY-010910] [Server] /usr/sbin/mysqld: Shutdown complete (mysqld 8.0.12)  MySQL Community Server - GPL.
```

直接内存不足，mysqld无法启动。

只好修改`/etc/mysql/mysql.conf.d/mysqld.cnf`文件，在结尾增加：

    performance_schema = off

一行，禁用`performance_schema`节省内存占用，mysql才得以正常启动。

### nginx

当然，也还需要nginx，直接 `apt-get install nginx`即可。

## 应用

整个网站应用我是使用go编写，并且使用[gorazor](https://github.com/sipin/gorazor)、[esc](https://github.com/mjibson/esc)等工具将用到的模板、静态资源文件等到打包进可执行文件中。

整个程序运行时占用不到20M内存，还有百余兆空余内存：

```bash

$ cat /proc/16537/status 

Name:   goblog                                                    
Umask:  0002                                                      
State:  S (sleeping)                                              
Tgid:   16537                                                     
Ngid:   0                                                         
Pid:    16537                                                     
PPid:   16403                                                     
TracerPid:      0                                                 
Uid:    1001    1001    1001    1001                              
Gid:    1002    1002    1002    1002                              
FDSize: 256                                                       
Groups: 4 20 24 25 29 30 44 46 108 114 1000 1001 1002             
NStgid: 16537                                                     
NSpid:  16537                                                     
NSpgid: 16537                                                     
NSsid:  16403                                                     
VmPeak:    16820 kB                                               
VmSize:    16820 kB                                               
VmLck:         0 kB                                               
VmPin:         0 kB                                               
VmHWM:     13356 kB                                               
VmRSS:      9564 kB                                               
RssAnon:            5720 kB                                       
RssFile:            3844 kB                                       
RssShmem:              0 kB                                       
VmData:     8920 kB                                               
VmStk:       132 kB                                               
VmExe:      4320 kB                                               
VmLib:         8 kB                                               
VmPTE:        80 kB                                               
VmSwap:        0 kB                 

$ free -h
              total        used        free      shared  buff/cache   available
Mem:           581M        352M         66M        900K        162M        135M
Swap:            0B          0B          0B
```

## 流量

谷歌云服务的免费实例运行时不收钱，但实例跑网站产生的流量可能是需要钱的，比方说，从中国、澳洲产生的访问流量则完全不免费。

怎么办？

在网站前面直接再套一个[cloudflare](https://www.cloudflare.com/plans/)的CDN，cloudflare有提供免费的CDN供个人站点使用！

cloudflare溯源去谷歌云的话，显然也不可能走中国、澳洲的流量。

## 总结

这样，我们就可以安安静静的近乎永久免费的跑一个网站了。
