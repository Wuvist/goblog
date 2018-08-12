<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcate.ascx.vb" Inherits="Blogwind.tcate" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html>
<head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title><%=path%></title>
<link rel="stylesheet" href="/Template/skin12.css" type="text/css"> 
<script type="text/javascript" src="/icon.js"></script>
</head>

<body topmargin="0" leftmargin="0" link="#808000" vlink="#808000" alink="#808000">
<table border="0" width="780" cellspacing="0" cellpadding="0" align=left>
<tr><td height=153 colspan=2 background="/images/skins/skin12/mao.jpg"><p class="header"><b><%=path%><br>
</td></tr>
<tr><td valign="top" width="580">
<DIV class="blog">
	  <asp:Repeater id="dlArticles" runat="server">
		<ItemTemplate>				
				<table width=100%><tr><td><a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><p><font size="3" color=black><b><%# Container.DataItem("title") %></b></font></a>
				<td align="right"><img border="0" src="/images/skins/skin12/mao1.jpg" width="84" height="56"></table><div class="PostFooter">
			   - - - - - - - - - - - - - - - - - - - - - - - <font size="2"><%=nick%> <img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%# Container.DataItem("add_date") %> 
			<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml">评论:<%#Container.DataItem("comment")%></a></div><br>
		</ItemTemplate>
		<AlternatingItemTemplate>				
				<table width=100%><tr><td><a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><p><font size="3" color=black><b><%# Container.DataItem("title") %></b></font></a>
				<td align="right"><img border="0" src="/images/skins/skin12/mao2.jpg" width="77" height="50"></table><div class="PostFooter">
			   - - - - - - - - - - - - - - - - - - - - - - - <font size="2"><%=nick%> <img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%# Container.DataItem("add_date") %> 
			<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml">评论:<%#Container.DataItem("comment")%></a></div><br>
		</AlternatingItemTemplate>
	</asp:Repeater>
</div>
</td><td valign="top" align=right width="200">
	      <div class="blog" style="width: 200">
	  <p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("导航")%></p><a href="/"><%=ph.g("博客风")%></a><BR>
      <a href="/<%=uid%>/"><%=path%>首页</a><BR>
  <a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR></p>
      <p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("网志介绍")%></p>
	  <Acme:userpic runat="server" ID="userpic1"/>
      <p><%=uinfo%></p>
      <p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("个人链接")%></p><asp:Repeater id="links" runat="server">
						<ItemTemplate>
							<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
						</ItemTemplate>
					</asp:Repeater></p>
      <p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("网志分类")%></p>
	  <asp:Repeater id="dlCate" runat="server">
						<ItemTemplate><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
						</ItemTemplate>
					</asp:Repeater>
      <p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("网志存档")%></p>
<asp:Repeater id="archive" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
		</ItemTemplate>
	</asp:Repeater>
	<p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("累计浏览")%>：<%=count%><br>
	<Acme:Message runat="server" ID="Message1"/></p></div></td>
</tr></table>
</body></html>