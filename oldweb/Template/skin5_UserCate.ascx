<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcate.ascx.vb" Inherits="Blogwind.tcate" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<title><%=path%></title>
<link rel="stylesheet" href="/Template/skin5.css" type="text/css"> 
<script type="text/javascript" src="/icon.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body bgcolor="ffffff" marginwidth="0" marginheight="0">

<table width="100%" border="0" cellspacing="0" cellpadding="0">
<tr bgcolor="#ffffff">
<td class="name"  width="100%" height="60"><%=path%>
</td>
</tr>
</table>

<table width="100%" border="0" cellspacing="0" cellpadding="0">
<tr bgcolor="#990000">
<td class="description" align="right" width="100%" height="30"><%=uinfo%></td>
</tr>
</table>

<table width="100%" cellpadding="0" cellspacing="0"><tr><td valign="top">
<div class="blog">
<asp:Repeater id="dlArticles" runat="server">
	<ItemTemplate><div class="datetimefooter"><a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml" class="title"><%# Container.DataItem("title") %></a>
		   <%=nick%>　@<%# Container.DataItem("add_date") %>
		 <a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml">评论:<%#Container.DataItem("comment")%></a></div><br>
	</ItemTemplate>
</asp:Repeater>
</div>
</div>
</TD>
<TD bgColor="#e0e0b1" height="500" width="150" valign="top"><div class="side">
<div class="sidetop">
 <%=ph.g("导航")%>
 </div> 
  <a href="/"><%=ph.g("博客风")%></a><BR>
  <a href="/<%=uid%>/"><%=path%>首页</a><BR>
  <a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR>
<br>

<div class="sidetop"><br>
个人档案
</div>
<Acme:userpic runat="server" ID="userpic1" deco="<div align='center'><br><img src='" enddeco="'></div>"/>
<%=uinfo%>
<br>
<br><br><br><div class="sidetop">
<%=ph.g("网志分类")%>
</div>
<asp:Repeater id="dlCate" runat="server">
<ItemTemplate>
	<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
</ItemTemplate>
</asp:Repeater>
<br><br><br><div class="sidetop">
<%=ph.g("网志存档")%>
</div>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate>
		<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
	</ItemTemplate>
</asp:Repeater>


<br><br><br><div class="sidetop">
<%=ph.g("个人链接")%>
</div>
<asp:Repeater id="links" runat="server">
	<ItemTemplate>
		<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
	</ItemTemplate>
</asp:Repeater>
<p class="nav"><%=ph.g("累计浏览")%>: <%=count%></p>
<br>
<Acme:Message runat="server" ID="Message1"/>
<br>
</div>
</TD>
</TR>
</TABLE>
</body>
</html>
