<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcate.ascx.vb" Inherits="Blogwind.tcate" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<HTML>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<TITLE><%=path%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<link rel="stylesheet" type="text/css" href="/Template/skin17.css">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</HEAD>
<BODY>
<table width="100%" cellpadding="0">
	<tr>
		<td width="20%" valign="top">
			<table>
				<tr><td><img src="/images/skins/skin17/k.gif" align="middle"></td><td>die Menuefuehrung</td></tr>
				<tr><td></td><td><a href="/">Blogwind Home</a></td></tr>
				<tr><td></td><td><a href="/<%=uid%>/"><%=path%>Home</a></td></tr>
				<tr><td></td><td> <a href="/email.aspx?user=<%=uid%>">Kontakt</a></td></tr>
				<tr><td><img src="/images/skins/skin17/k.gif" align="center"></td><td><%=ph.g("网志介绍")%></td></tr>
				<tr><td></td><td><Acme:userpic runat="server" ID="userpic1"/><%=uinfo%></td></tr>
				<tr><td><img src="/images/skins/skin17/k.gif" align="center"></td><td>die Kategorien</td></tr>
				<asp:Repeater id="dlCate" runat="server">
					<ItemTemplate>
					<tr><td></td><td><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)</td></tr>
					</ItemTemplate>
				</asp:Repeater>
				<tr><td><img src="/images/skins/skin17/k.gif" align="center"></td><td><%=ph.g("个人链接")%></td></tr>
				<asp:Repeater id="links" runat="server">
					<ItemTemplate>
					<tr><td></td><td><a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a></td></tr>
					</ItemTemplate>
				</asp:Repeater>
				<tr><td><img src="/images/skins/skin17/k.gif" align="center"></td><td><%=ph.g("网志存档")%></td></tr>
				<asp:Repeater id="archive" runat="server">
					<ItemTemplate><tr><td></td><td><a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %> - <%# Container.DataItem("month") %></a></td></tr>
					</ItemTemplate>
				</asp:Repeater>
				<tr><td valign="top"><img src="/images/skins/skin17/k.gif" align="center"></td><td>die Besucherzaehlung <br><%=count%></td></tr>
				<tr><td></td><td><Acme:Message runat="server" ID="Message1"/></td></tr>
			</table>
		</td>
		<td valign="top">	<table width="100%" cellpadding="0" cellspacing="5"><asp:Repeater id="dlArticles" runat="server">	
	<ItemTemplate><tr><td colspan=3 valign="top"><img src="/images/skins/skin17/open.gif"</td></tr>
				<tr><td width="15"></td><td><a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml" class="title"><%# Container.DataItem("title") %></a></td><td><div style="text-align:right"><%# Container.DataItem("add_date") %> von <%=nick%> <a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"> Kommentieren ( <%#Container.DataItem("comment")%> ) </a></div></td></tr>
				<tr><td colspan=3 valign="top" align="right"><img src="/images/skins/skin17/close.gif"</td></tr>
			</ItemTemplate>
</asp:Repeater>
			</table>	
		</td>
	</tr>
	<tr>
		<td colspan="2"></td>
	</tr>
</table>
</BODY>
</HTML>