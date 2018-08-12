<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tDefault.ascx.vb" Inherits="Blogwind.tDefault" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<script type="text/javascript" src="/icon.js"></script>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<title><%=path%></title>
<link rel="stylesheet" href="/Template/skin2.css" type="text/css"> 
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body bgcolor=black> 
<div style="position:absolute; width:304; height:200; z-index:2; left:47; top:20 ;border-left: 1px solid #660000;border-top: 1px solid #660000;border-right: 1px solid #660000;border-bottom: 1px solid #660000; " >
<img src="/images/skins/life/match.jpg" width="333" height="511" >
</div>
<div style="position:absolute; width:883; height:25; z-index:1; left:47; top:535 ;border-left: 1px solid #660000;border-top: 1px solid #660000;border-right: 1px solid #660000;border-bottom: 1px  solid #660000; background-color: black;" >
<font color="#999999">
<asp:Repeater id="links" runat="server">
	<ItemTemplate>
	&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
		<a href = "<%# Container.DataItem("URL") %>" target="blank"><%# Container.DataItem("link") %></a>
	</ItemTemplate>
</asp:Repeater>
</font>
</div>
<div style="position:absolute; width:970; height:100; z-index:1; left:0; top:0 ;border-left: 1px none #999999;border-top: 1px none #999999;border-right: 1px none #999999;border-bottom: 1px none #999999; background-color: black;" >
</div>
<div class="maindiv">

<div class="headerdiv">
<br>
<center>
<font color="orange"><%=path%></font>

</center>
</div>

<div class="sidediv">
<font color="orange"><%=ph.g("导航")%><br><br></font>
  
  <a href="/"><%=ph.g("博客风")%></a><BR>
  <a href="/<%=uid%>/"><%=path%><%=ph.g("首页")%></a><BR>
  <a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR>
<hr color="660000">
<font color="orange"><%=ph.g("网志介绍")%></font><br><br><Acme:userpic runat="server" ID="userpic1" deco="<div align='center'><img src='" enddeco="'></div>"/><%=uinfo%><br><hr color="660000">
<font color="orange"><%=ph.g("网志分类")%></font>
<br><br>
<asp:Repeater id="dlCate" runat="server">
	<ItemTemplate>
		<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
	</ItemTemplate>
</asp:Repeater>
<hr color="660000">
<font color="orange"><%=ph.g("网志存档")%></font>
<br><br>
<asp:Repeater id="archive" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
		</ItemTemplate>
	</asp:Repeater>
<hr color="660000">
<font color="orange"><%=ph.g("累计浏览")%>：<%=count%></font>
<hr color="660000">
<Acme:Message runat="server" ID="Message1"/>
</div>

<div class="blogdiv">
<asp:Repeater id="dlArticles" runat="server">
	<ItemTemplate>
	<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><font color="orange" size="3"><b><%# Container.DataItem("title") %></b></font></a>
	<p><%#Container.DataItem("content")%></p>
	<p><font color="orange">>></font><font color="#999999">　<%=nick%>　发表于：  <%# Container.DataItem("add_date") %> <a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml">评论： <font color="orange"><%#Container.DataItem("comment")%></a></font>. </font></p>

<hr color="660000">
	</ItemTemplate>
</asp:Repeater><uc1:userpaging runat="server" ID="paging"/>
</div>

</div>
</body>
</html>