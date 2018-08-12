<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcate.ascx.vb" Inherits="Blogwind.tcate" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<title><%=path%></title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<script type="text/javascript" src="/icon.js"></script>
<link rel="stylesheet" href="/Template/skin1.css" type="text/css"> 
</head>

<body style="FILTER: progid:DXImageTransform.Microsoft.Gradient(endColorstr='#ffffff', startColorstr='#9999CC', gradientType='1'); WIDTH: 100%; HEIGHT: 100%">
<div id="Layer1" style="position:absolute; width:400px; height:100%; z-index:1; top: 0; left: 350;" class="entry"><img src="/images/skins/bluesky/bluesky.jpg" width="400" height="299"><br>
<div class="main">
<asp:Repeater id="dlArticles" runat="server">
	<ItemTemplate>
			<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml" class="titles"><%# Container.DataItem("title") %></a><br><br>
		   <div class="com"><%=nick%>　@<%# Container.DataItem("add_date") %>
		 <a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"> <%=ph.g("评论")%>:<%#Container.DataItem("comment")%></a></div>
	</ItemTemplate>
</asp:Repeater>
</div>
</div>
<div id="Layer2" style="position:absolute; width:200px; height:115px; z-index:2; left: 125; top: 0;" class="navbar">
  <p class="nav"><%=ph.g("导航")%></p>
  <p>
  
  <a href="/"><%=ph.g("博客风")%></a><BR>
  <a href="/<%=uid%>/"><%=path%><%=ph.g("首页")%></a><BR>
  <a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a></p>

  <p class="nav"><%=ph.g("网志介绍")%></p>
  <Acme:userpic runat="server" ID="userpic1" deco="<div align='center'><img src='" enddeco="'></div>"/>
  <p><%=uinfo%></p>
  <p class="nav"><%=ph.g("网志分类")%></p>
  <p><asp:Repeater id="dlCate" runat="server">
						<ItemTemplate>
							<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
						</ItemTemplate>
					</asp:Repeater>
  <p class="nav"><%=ph.g("网志存档")%></p>
  <p><asp:Repeater id="archive" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %> - <%# Container.DataItem("month") %></a><br>
		</ItemTemplate>
	</asp:Repeater></P>
  <p class="nav"><%=ph.g("个人链接")%></p>
  <p><asp:Repeater id="links" runat="server">
						<ItemTemplate>
							<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
						</ItemTemplate>
					</asp:Repeater></p>  
  <p class="nav"><%=ph.g("累计浏览")%>: <%=count%></p>
<Acme:Message runat="server" ID="Message1"/>
</div>
</body>
</html>