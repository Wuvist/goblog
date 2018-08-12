# goblog

这是博客风 www.blogwind.com 使用的后端程序。

博客风原本是我在2003年使用ASP.Net 2.0开发的网站，几乎是我在学生时代写的第一个网站，当年的源码基本保留在[oldweb](https://github.com/Wuvist/goblog/)路径下，仅供备份、参考用户。

这个项目的则是我今年重新使用go实现的后端程序，已演示如何使用go开发一个博客网站，暂时只实现了各种浏览功能。

项目中使用的技术有：

* web框架：[echo](https://github.com/labstack/echo)
* ORM：[sqlboiler](https://github.com/volatiletech/sqlboiler)
* 视图引擎：[gorazor](https://github.com/sipin/gorazor)
* 静态资源内嵌：[esc](https://github.com/mjibson/esc)

值得一提的是，`sqlboiler`、`gorazor`、`esc`均是基于代码生成的工具，项目中并不会直接使用到它们的代码，而仅是使用它们生成的go代码。
