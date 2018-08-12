<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcate.ascx.vb" Inherits="Blogwind.tcate" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>  
	<title><%=path%></title>
<link rel=stylesheet href="/Template/skin8.css" type=text/css>
<script type="text/javascript" src="/icon.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body bgcolor="#000000" text="#000000">
<table border="0"
background="/images/skins/skin8/mettest3.jpg"  bgcolor="ffffff"
class="borrrder" cellpadding="5" cellspacing="0" width="750" height="100%" align="center">
<tr><td colspan="2" valign="bottom" align="left" height="160" 
background="/images/skins/skin8/abuseshawdows.jpg">
</td></tr> 
<td colspan="2"  bgcolor="666666"height="20"> 
<div align="right" >
<font color="999999">
<i><%=path%></i>
</font>
</div>
</td>
<tr valign="top">

<td  width="550" align="justify" valign="top" height="100%" >
<asp:Repeater id="dlArticles" runat="server">
	<ItemTemplate>
	<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><div class="datetimeheader">
<%# Container.DataItem("title") %></a>
</div><a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><div class="datetimefooter">
posted at 

<%# Container.DataItem("add_date") %> by  
<%=nick%> 评论:<%#Container.DataItem("comment")%></div></a><br><br>
	</ItemTemplate>
</asp:Repeater>
</td>

<td  width="150"  height="100%" > 
<div class="side" align="left">
<div class="sidetop"><br>
<%=ph.g("导航")%>
</div>
<a href="/"><%=ph.g("博客风")%></a><BR>
<a href="/<%=uid%>/"><%=path%>首页</a><BR>
<a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR>

<br><br><br><div class="sidetop">
<%=ph.g("网志介绍")%>
</div>
<p><Acme:userpic runat="server" ID="userpic1" deco="<div align='center'><img src='" enddeco="'></div>"/><%=uinfo%></p>

<br><br><br><div class="sidetop">
<%=ph.g("网志分类")%>
</div>
<asp:Repeater id="dlCate" runat="server">
	<ItemTemplate>
		<br><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)
	</ItemTemplate>
</asp:Repeater>



<br><br><br><div class="sidetop">
<%=ph.g("网志存档")%>
</div>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate>
		<br><a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a>
	</ItemTemplate>
</asp:Repeater>
</div>

<br><br><br><div class="sidetop">
<%=ph.g("个人链接")%>
</div>
<asp:Repeater id="links" runat="server">
	<ItemTemplate>
		<br><a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a>
	</ItemTemplate>
</asp:Repeater>

<br><br><br><div class="sidetop">
<%=ph.g("累计浏览")%>：<%=count%><br></div>
<Acme:Message runat="server" ID="Message1"/>
</div>
</td>


</tr>
<tr><td colspan="2"  valign="bottom" align="right" class="small" height="10" bgcolor="666666">
| maystar designs |</td></tr></table>
</body></html>