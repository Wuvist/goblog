<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tDefault.ascx.vb" Inherits="Blogwind.tDefault" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3c.org/TR/1999/REC-html401-19991224/loose.dtd">
<HTML lang=gb2312 xmlns="http://www.w3.org/1999/xhtml">
<head>
<script type="text/javascript" src="/icon.js"></script>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<TITLE><%=path%></TITLE>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<META http-equiv=Content-Language content=gb2312>
<META content=all name=robots>
<!-- 调用样式表 -->
<LINK media=all href="/Template/skin15.css" type=text/css rel=stylesheet>

<META content="MSHTML 6.00.3790.186" name=GENERATOR></HEAD>
<BODY id=w3cn><!--页面顶部-->
<DIV id=top>
<DIV id=popimg></DIV>
<DIV id=logo></DIV><!-- 主菜单 -->
<DIV id=blogname><%=path%></DIV>
</DIV><!--布局中间区域第一行-->
<DIV id=middle1>
<DIV id=leftmenu>
<UL>
  <LI><%=ph.g("导航")%></LI>
  </UL>
  <span class="textmenu"><a href="/"><%=ph.g("博客风")%></a><BR>
      	<a href="/<%=uid%>/"><%=path%>首页</a><BR>
  	<a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a>
  </span> 
<UL>
  <LI><%=ph.g("网志介绍")%></LI></UL>
  <span class="textmenu"><Acme:userpic runat="server" ID="userpic1"/><%=uinfo%></span> 
<UL>
  <LI><%=ph.g("网志分类")%></LI>
  </UL>
  <span class="textmenu"><asp:Repeater id="dlCate" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
		</ItemTemplate>
	</asp:Repeater></span>
<UL>
  <LI><%=ph.g("网志存档")%></LI>
  </UL>
  <span class="textmenu"><asp:Repeater id="archive" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
		</ItemTemplate>
	</asp:Repeater></span> 
<UL>
  <LI><%=ph.g("个人链接")%></LI>
  </UL>
  <span class="textmenu"><asp:Repeater id="links" runat="server">
		<ItemTemplate>
			<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
		</ItemTemplate>
	</asp:Repeater></span> 
<UL>
  <LI><%=ph.g("累计浏览")%>：<%=count%></LI>
</UL><span class="textmenu"><Acme:Message runat="server" ID="Message1"/></span></DIV>
<div id="note">
<asp:Repeater id="dlArticles" runat="server">
		<ItemTemplate>				
			<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><span class="title12"><%# Container.DataItem("title") %></span></a>
			<div class="text"><%#Container.DataItem("content")%>
			<div class="PostFooter"><%=nick%> @ <%# Container.DataItem("add_date") %><a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"> 评论:<%#Container.DataItem("comment")%></a></div><hr></div>
		</ItemTemplate>
	</asp:Repeater><uc1:userpaging runat="server" ID="paging"/>
</DIV><!--布局中间区域第二行-->
<DIV id=boxtitle></DIV>
<!--布局中间区域第三行-->
<!--底部 -->
<DIV id=bottom>
<DIV id=btmenu>Powered By ：BlogWind Blog System 1.1.0    Copyright ?2003 - 2004 <A href="http://www.blogwind.com">BlogWind.com</a></DIV></DIV>
<DIV id=copyright>
</DIV>
</BODY></HTML>