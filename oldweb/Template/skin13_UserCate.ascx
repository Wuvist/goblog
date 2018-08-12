<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcate.ascx.vb" Inherits="Blogwind.tcate" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html>
<head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title><%=path%></title>
<link rel="stylesheet" href="/Template/skin13.css" type="text/css"> 
<script type="text/javascript" src="/icon.js"></script>
</head>

<BODY>

<table width= "100%" cellpadding="0" cellspacing="0">

<tr>
<td width="18%" valign="top">
	<table width="100%">
	<tr><td class="cells"><center>
      <div class="blog" style="width: 200">
	  <p class=head><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"><%=ph.g("导航")%><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"></p><a href="/"><%=ph.g("博客风")%></a><BR>
      <a href="/<%=uid%>/"><%=path%>首页</a><BR>
  <a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR></p>
      <p class=head><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"><%=ph.g("网志介绍")%><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"></p>
	  <Acme:userpic runat="server" ID="userpic1"/>
      <p><%=uinfo%></p>
      <p class=head><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"><%=ph.g("个人链接")%><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"></p><asp:Repeater id="links" runat="server">
						<ItemTemplate>
							<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
						</ItemTemplate>
					</asp:Repeater></p>
      <p class=head><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"><%=ph.g("网志分类")%><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"></p>
	  <asp:Repeater id="dlCate" runat="server">
						<ItemTemplate><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
						</ItemTemplate>
					</asp:Repeater>
      <p class=head><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"><%=ph.g("网志存档")%><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"></p>
<asp:Repeater id="archive" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
		</ItemTemplate>
	</asp:Repeater>
	<p class=head><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"><%=ph.g("累计浏览")%>：<%=count%><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"><br>
	<Acme:Message runat="server" ID="Message1"/></p></div>

	</td></tr></table>

</td><td width="2%"></td>
<td width="80%">
	<table width="100%" cellpadding="10"><tr><td class="cells">
	<font><b><big><%=path%></b></font color></font size>
	</td></tr>
	<asp:Repeater id="dlArticles" runat="server">
		<ItemTemplate>
			<tr><td>&nbsp;</td></tr><tr><td class="cells">
			<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><b><%# Container.DataItem("title") %></b></a><br>
				<span class="PostFooter">- - posted by <%=nick%><img border="0" src="/images/skins/skin13/f1.gif" width="24" height="18"><%# Container.DataItem("add_date") %>
				 <a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml">评论:<%#Container.DataItem("comment")%></a> </span>
			</td></tr>
	
		</ItemTemplate>
	</asp:Repeater>
	</table>
</td>
</tr>
</table>
<script language=JavaScript src="/Template/skin13_mouse.js"></script>
</body>
</html>