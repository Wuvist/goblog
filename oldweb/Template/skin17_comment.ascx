<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<HTML>
<head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
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
		<td valign="top">	<table width="100%" cellpadding="0" cellspacing="5">
				<tr><td>
					<img src="/images/skins/skin17/open.gif"><br>
					<center><%=title%><br>
					<hr width="30%">
				</td></tr>
				<tr><td><%=content%></td></tr>
				<tr><td><div id="tags"><%=tags%></div><div style="text-align:right">
					<%=adddate%> von <%=nick%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a><br><br>
					<img src="/images/skins/skin17/close.gif" align="center"><br>
				</div></td></tr>
			</table>

		<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a>
			<div class="comments">
				<p><%# Container.DataItem("title") %></p>
				<p><%# Container.DataItem("content") %></p>
				<div style="text-align:right" id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("add_date")%> von <%#Container.DataItem("author")%></div>
				<hr>
			</div></div>
		</ItemTemplate>
	</asp:Repeater>
			<div class="addcomment">
				<form action="/savecomment.aspx" method=post>
				<input type=hidden name="action" value="add">
				<input type=hidden name="article_id" value="<%=id%>">
				<table width="100%">
				<tr><td colspan="2">Kommentar hinterlassen:</td></tr>
				<tr><td>Betreff:</td><td><input type="text" name="TitleText" size=50 value="回复：<%=title%>" maxlength=50></td></tr>
				<tr><td>Von:</td><td><input type="text" name="AuthorText"  size=50 maxlength=50 value="<%=viewer%>"></td></tr>
				<tr><td>Homepage:</td><td><input type="text" name="HomepageText" size=50  value="<%=hp%>" maxlength=255></td></tr>
				<tr><td valign="top">Nachricht:</td><td><textarea cols="50" rows="5" name="ContentText"></textarea></td></tr>
				<tr><td></td><td><input type="submit" value="Kommentar abschicken"></td></tr>
				</table>
				</form>
			</div>
		</td>
			</table>	
		</td>
	</tr>
	<tr>
		<td colspan="2"></td>
	</tr>
</table>
</BODY>
<script type="text/javascript">tagging();</script>
</html>