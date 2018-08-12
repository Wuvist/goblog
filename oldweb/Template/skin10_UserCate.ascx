<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcate.ascx.vb" Inherits="Blogwind.tcate" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html>
<head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link> <TITLE><%=path%></TITLE>
<link rel=stylesheet href="/Template/skin10.css" type=text/css> 
<script type="text/javascript" src="/icon.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body background="/images/skins/skin10/momentbg.jpg"  leftmargin="0" bottommargin="0" topmargin="0" rightmargin="0">

<div align="left">

<table border="0" width="760" cellpadding="0" cellspacing="0">
<tr><td><img src="/images/skins/skin10/moment1.jpg" width="760" height="230"></tr>
<tr><td><img src="/images/skins/skin10/moment2.jpg" width="760" height="180"></tr>
<tr><td><img src="/images/skins/skin10/moment3.jpg" width="760" height="169"></tr>
</table>

<div id="me">
<table border="0" width="180" cellpadding="0" cellspacing="0">
<tr>
<td valign="top" class=tab>
<img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("网志介绍")%></b>
<br><Acme:userpic runat="server" ID="userpic1"/><%=uinfo%>
<p>
<img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("导航")%></b>
<br>&nbsp;<a href="/"><%=ph.g("博客风")%></a>
<br>&nbsp;<a href="/<%=uid%>/"><%=path%>首页</a>
<br>&nbsp;<a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a>
<p><img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("网志分类")%></b>
<asp:Repeater id="dlCate" runat="server">
<ItemTemplate>
<br>&nbsp;<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)
</ItemTemplate>
</asp:Repeater>
<p><img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("网志存档")%></b>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate>
<br>&nbsp;<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
	</ItemTemplate>
</asp:Repeater>
<p>
<img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("个人链接")%></b>
<asp:Repeater id="links" runat="server">
	<ItemTemplate>
	<br>&nbsp;<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a>
	</ItemTemplate>
	</asp:Repeater>
<p>
<img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("累计浏览")%>：<%=count%></b>
<br>
<Acme:Message runat="server" ID="Message1"/>
</td>
</tr>
</table>
</div>



<div id="blog">
<table border="0" width="310" cellpadding="0" cellspacing="0">
<tr>
<td valign="top"><align=justify>
<P>
<asp:Repeater id="dlArticles" runat="server">
	<ItemTemplate>
			<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><div class="DateHeader"><%# Container.DataItem("title") %></div></a>
			<span class="PostFooter">- post by <%=nick%>　@<%# Container.DataItem("add_date") %>
		 <a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml">评论:<%#Container.DataItem("comment")%></a> </span><BR><BR>
	</ItemTemplate>
</asp:Repeater>

   </div>
</p>
</td>
</tr>
</table>
</div>
</body>
</html>