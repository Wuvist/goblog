<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcate.ascx.vb" Inherits="Blogwind.tcate" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<title><%=path%></title>
<script type="text/javascript" src="/icon.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<link rel="stylesheet" href="/Template/skin14.css" type="text/css">
</head>

<body topmargin="0" leftmargin="0">
<table width=950 border="0" align="center">
  <tr > 
    <td background="/images/skins/skin14/grey-banner.gif" width=900 height="160" colspan="2" align="right" valign="top"><p class="header"><%=path%></p></td>
  </tr>
  <tr> 
    <td width=750 align="left" valign="top" bordercolor="#CCCCCC"><DIV class="blog">
	<asp:Repeater id="dlArticles" runat="server">
		<ItemTemplate>				
			<table width=100%><tr><td><a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><font size="3" color=#333333><b><%# Container.DataItem("title") %></b></font></a>
			<td align="right"></table><br><div class="PostFooter">
<img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=nick%><%# Container.DataItem("add_date") %> 
			<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml">评论:<%#Container.DataItem("comment")%></a></div>
			<hr><br>
		</ItemTemplate>
	</asp:Repeater>
</div></td>
	<td width=200 background="/images/skins/skin14/grey-colback.gif" align="left" valign="top"><div class="blog" style="width: 200">
	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("导航")%></p><a href="/"><%=ph.g("博客风")%></a><BR>
      	<a href="/<%=uid%>/"><%=path%>首页</a><BR>
  	<a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR></p>
      	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("网志介绍")%></p>
		<Acme:userpic runat="server" ID="userpic1" deco="<img src='" enddeco="'>"/>
      	<p><%=uinfo%></p>
      	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("个人链接")%></p>
	<asp:Repeater id="links" runat="server">
		<ItemTemplate>
			<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
		</ItemTemplate>
	</asp:Repeater></p>
      	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("网志分类")%></p>
	 <asp:Repeater id="dlCate" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
		</ItemTemplate>
	</asp:Repeater>
    	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("网志存档")%></p>
	<asp:Repeater id="archive" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
		</ItemTemplate>
	</asp:Repeater>
	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("累计浏览")%>：<%=count%><br><br><br><br>
	<Acme:Message runat="server" ID="Message1"/></p></div></td>
  </tr>
</table><br><br><br>
</body>
</html>