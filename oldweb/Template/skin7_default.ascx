<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tDefault.ascx.vb" Inherits="Blogwind.tDefault" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
   <TITLE><%=path%></TITLE>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<link rel=stylesheet href="/Template/skin7.css" type=text/css> 
<script type="text/javascript" src="/icon.js"></script>
</HEAD>
<BODY BGCOLOR="#FFFFFF">
<CENTER><TABLE BORDER=0 CELLSPACING=0 CELLPADDING=0 WIDTH="100%" HEIGHT=53 BACKGROUND="/images/skins/bluemac/apfelblue_lines.gif">
   <TR>
      <TD VALIGN=top>
         <table width=200 height="33" border="0" cellpadding="0" cellspacing="0" align=center>
		<tr>
			<td width="8" background="/images/skins/bluemac/title_left.gif"></td>
			<td background="/images/skins/bluemac/title.gif" class=intro valign=bottom width="182><div align="center"> <%=path%> </div></td>
			<td width="8" background="/images/skins/bluemac/title_right.gif"></td>
		</tr>
		</table>
      </TD>
   </TR>
</TABLE>
<BR>
<TABLE BORDER=0 CELLSPACING=0 CELLPADDING=0 WIDTH=100%>
   <TR><TD WIDTH=20>
      </TD>
      <TD VALIGN=top WIDTH=150>
<table width="150" height="23" border="0" cellpadding="0" cellspacing="0">
  <tr>
    <td width="15" background="/images/skins/bluemac/bar_left.gif">&nbsp;</td>
    <td background="/images/skins/bluemac/bar.gif" class=bartitle>导　航</td>
    <td width="15" background="/images/skins/bluemac/bar_right.gif">&nbsp;</td>
  </tr>
</table>
		<BR>
		<IMG src="/images/skins/bluemac/point.gif" width="12" height="12">&nbsp;<a href="/"><%=ph.g("博客风")%></a> <br>
		<IMG src="/images/skins/bluemac/point.gif" width="12" height="12">&nbsp;<a href="/<%=uid%>/"><%=path%>首页</a> <br>
		<IMG src="/images/skins/bluemac/point.gif" width="12" height="12">&nbsp;<a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR>

         <P><table width="150" height="23" border="0" cellpadding="0" cellspacing="0">
  <tr>
    <td width="15" background="/images/skins/bluemac/bar_left.gif">&nbsp;</td>
    <td background="/images/skins/bluemac/bar.gif" class=bartitle><%=ph.g("网志介绍")%></td>
    <td width="15" background="/images/skins/bluemac/bar_right.gif">&nbsp;</td>
  </tr>
</table></P>
         <p><Acme:userpic runat="server" ID="userpic1"/><%=uinfo%></P>
         <P><table width="150" height="23" border="0" cellpadding="0" cellspacing="0">
  <tr>
    <td width="15" background="/images/skins/bluemac/bar_left.gif">&nbsp;</td>
    <td background="/images/skins/bluemac/bar.gif" class=bartitle><%=ph.g("网志分类")%></td>
    <td width="15" background="/images/skins/bluemac/bar_right.gif">&nbsp;</td>
  </tr>
</table></P>
         <asp:Repeater id="dlCate" runat="server">
					<ItemTemplate>
					<IMG src="/images/skins/bluemac/point.gif" width="12" height="12">&nbsp;<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
					</ItemTemplate>
					</asp:Repeater>
			<P><table width="150" height="23" border="0" cellpadding="0" cellspacing="0">
  <tr>
    <td width="15" background="/images/skins/bluemac/bar_left.gif">&nbsp;</td>
    <td background="/images/skins/bluemac/bar.gif" class=bartitle><%=ph.g("个人链接")%></td>
    <td width="15" background="/images/skins/bluemac/bar_right.gif">&nbsp;</td>
  </tr>
</table></P><asp:Repeater id="links" runat="server">
				<ItemTemplate>
				<IMG 
            src="/images/skins/bluemac/point.gif"" width="12" height="12">&nbsp;<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
				</ItemTemplate>
				</asp:Repeater>
         <P><table width="150" height="23" border="0" cellpadding="0" cellspacing="0">
  <tr>
    <td width="15" background="/images/skins/bluemac/bar_left.gif">&nbsp;</td>
    <td background="/images/skins/bluemac/bar.gif" class=bartitle><%=ph.g("网志存档")%></td>
    <td width="15" background="/images/skins/bluemac/bar_right.gif">&nbsp;</td>
  </tr>
</table></P>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate>
<IMG src="/images/skins/bluemac/point.gif" width="12" height="12">&nbsp;<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
	</ItemTemplate>
</asp:Repeater>
		 <P><table width="150" height="23" border="0" cellpadding="0" cellspacing="0">
  <tr>
    <td width="15" background="/images/skins/bluemac/bar_left.gif">&nbsp;</td>
    <td background="/images/skins/bluemac/bar.gif" class=bartitle><%=ph.g("累计浏览")%>：<%=count%></td>
    <td width="15" background="/images/skins/bluemac/bar_right.gif">&nbsp;</td>
  </tr>
</table></P>
		 <P><Acme:Message runat="server" ID="Message1"/></P>
		</div>
		</TD>
      <TD WIDTH=40>
         <P></P>
      </TD>
      <TD VALIGN=top>
<table width=100% border=0 bordercolor="#396BCE" cellpadding="0"  cellspacing="0">
<tr><td>
<asp:Repeater id="dlArticles" runat="server">
	<ItemTemplate>
<table width="100%" height="23" border="0" cellpadding="0" cellspacing="0">
  <tr>
    <td width="15" background="/images/skins/bluemac/bar_left.gif">&nbsp;</td>
    <td background="/images/skins/bluemac/bar.gif" class="title">　<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><font color=black><%# Container.DataItem("title") %></font></a></td>
    <td width="15" background="/images/skins/bluemac/bar_right.gif">&nbsp;</td>
  </tr>
</table><br><%#Container.DataItem("content")%>
<hr>
<div class="dayname"><%=nick%>　@<%# Container.DataItem("add_date") %> <a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml">评论：<%#Container.DataItem("comment")%></a></div><br>
</ItemTemplate>
</asp:Repeater><uc1:userpaging runat="server" ID="paging"/></td></tr></table>
      </TD>
      <TD VALIGN=top WIDTH=10>
      </TD>
   </TR>
</TABLE>
</BODY>
</HTML>
