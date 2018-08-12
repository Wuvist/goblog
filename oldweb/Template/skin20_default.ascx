<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tDefault.ascx.vb" Inherits="Blogwind.tDefault" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<TITLE><%=path%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<link rel="stylesheet" type="text/css" href="/Template/skin20.css">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</HEAD>
<BODY>
<table width="100%" cellpadding="0">
	<tr>
		<td width="250" valign="top">
			<object classid="clsid:d27cdb6e-ae6d-11cf-96b8-444553540000" codebase="http://download.macromedia.com/pub/shockwave/cabs/flash/swflash.cab#version=7,0,0,0" width="200" height="200" id="Untitled-1" align="middle">
			<param name="allowScriptAccess" value="sameDomain" />
			<param name="movie" value="/images/skins/skin20/moon1.swf" />
			<param name="quality" value="high" />
			<param name="bgcolor" value="#1b1b1b" />
			<embed src="/images/skins/skin20/moon1.swf" quality="high" bgcolor="#1b1b1b" width="200" height="200" name="Untitled-1" align="middle" allowScriptAccess="sameDomain" type="application/x-shockwave-flash" pluginspage="http://www.macromedia.com/go/getflashplayer" />
			</object>


			
			
			<table>
				<tr><td><img src="/images/skins/skin20/star.jpg" align="middle"></td><td><%=ph.g("导航")%></td></tr>
				<tr><td></td><td><a class="s" href="http://www.blogwind.com"><%=ph.g("博客风")%></a></td></tr>
				<tr><td></td><td><a class="s" href="/<%=uid%>/"><%=path%>首页</a></td></tr>
				<tr><td></td><td> <a class="s" href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%><%=nick%></a></td></tr>
				<tr><td><img src="/images/skins/skin20/star.jpg" align="center"></td><td><%=ph.g("网志介绍")%></td></tr>
				<tr><td></td><td><Acme:userpic runat="server" ID="userpic1" deco="<img src='" enddeco="'>"/><%=uinfo%></td></tr>
				<tr><td><img src="/images/skins/skin20/star.jpg" align="center"></td><td><%=ph.g("网志分类")%></td></tr>
				<asp:Repeater id="dlCate" runat="server">
					<ItemTemplate>
					<tr><td></td><td><a class="s" href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)</td></tr>
					</ItemTemplate>
				</asp:Repeater>
				<tr><td><img src="/images/skins/skin20/star.jpg" align="center"></td><td><%=ph.g("个人链接")%></td></tr>
				<asp:Repeater id="links" runat="server">
					<ItemTemplate>
					<tr><td></td><td><a class="s" href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a></td></tr>
					</ItemTemplate>
				</asp:Repeater>
				<tr><td><img src="/images/skins/skin20/star.jpg" align="center"></td><td><%=ph.g("网志存档")%></td></tr>
				<asp:Repeater id="archive" runat="server">
					<ItemTemplate><tr><td></td><td><a class="s" href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a></td></tr>
					</ItemTemplate>
				</asp:Repeater>
				<tr><td valign="top"><img src="/images/skins/skin20/star.jpg" align="center"></td><td><%=ph.g("累计浏览")%>：<%=count%></td></tr>
				<tr><td></td><td><Acme:Message runat="server" ID="Message1"/></td></tr>
			</table>
		</td>
		<td valign="top">
		
<asp:Repeater id="dlArticles" runat="server">
	<ItemTemplate>
		<table width="100%" cellpadding="0" cellspacing="5">
				<tr><td valign="center">
			<object classid="clsid:d27cdb6e-ae6d-11cf-96b8-444553540000" codebase="http://download.macromedia.com/pub/shockwave/cabs/flash/swflash.cab#version=7,0,0,0" width="30" height="30" id="Untitled-1" align="middle">
			<param name="allowScriptAccess" value="sameDomain" />
			<param name="movie" value="/images/skins/skin20/star.swf" />
			<param name="quality" value="high" />
			<param name="bgcolor" value="#1b1b1b" />
			<embed src="/images/skins/skin20/star.swf" quality="high" bgcolor="#1b1b1b" width="30" height="30" name="Untitled-1" align="middle" allowScriptAccess="sameDomain" type="application/x-shockwave-flash" pluginspage="http://www.macromedia.com/go/getflashplayer" />
			</object>
			<b><a class="s" href="/<%# uid & "/" & Container.DataItem("index")%>.shtml" class="title"><%# Container.DataItem("title") %></a></b>
				</td></tr>
				<tr><td><%#Container.DataItem("content")%></td></tr>
				<tr><td><div style="text-align:right">
					<%=nick%> 发表于 <%# Container.DataItem("add_date") %> | <a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml" class="s">评论 [<%#Container.DataItem("comment")%>]</a><br><br>
				</div>
			<br><br><center><img src="/images/skins/skin20/bar.gif" align="center"></center><br><br><br>

				</td></tr>
			</table>
			</ItemTemplate>
</asp:Repeater><uc1:userpaging runat="server" ID="paging"/>

		</td>
			</table>	

		</td>
	</tr>
</table>
</BODY>
</HTML>