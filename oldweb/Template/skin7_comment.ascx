<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
   <TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<link rel=stylesheet href="/Template/skin7.css" type=text/css> 
</HEAD>
<BODY BGCOLOR="#FFFFFF">
<CENTER><TABLE BORDER=0 CELLSPACING=0 CELLPADDING=0 WIDTH="100%" HEIGHT=53 BACKGROUND="/images/skins/bluemac/apfelblue_lines.gif">
   <TR>
      <TD VALIGN=top>
         <table width=200 height="33" border="0" cellpadding="0" cellspacing="0" align=center>
		<tr>
			<td width="8" background="/images/skins/bluemac/title_left.gif"></td>
			<td background="/images/skins/bluemac/title.gif" class=intro valign=bottom width="182"><div align="center"> <%=path%> </div></td>
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
<table width="100%" height="23" border="0" cellpadding="0" cellspacing="0">
  <tr>
    <td width="15" background="/images/skins/bluemac/bar_left.gif">&nbsp;</td>
    <td background="/images/skins/bluemac/bar.gif" class="title">　<font color=black><%=title%></font></td>
    <td width="15" background="/images/skins/bluemac/bar_right.gif">&nbsp;</td>
  </tr>
</table><br><%=content%>
<hr>
<div id="tags"><%=tags%></div>
<div class="dayname"><%=nick%>　@<%=adddate%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div><br>
	<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a>
	   	<div id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("author")%> 在 <%#Container.DataItem("add_date")%> 说:</div>
	<%# Container.DataItem("content") %>
	<hr></div><br>
</ItemTemplate>
</asp:Repeater>
    <p><%=ph.g("请发表评论")%></p>
	<p><form action="/savecomment.aspx" method=post>
		<input type=hidden name="action" value="add">
		<input type=hidden name="article_id" value="<%=id%>">
<INPUT type="hidden" name="TitleText" value="回复：<%=title%>">
	　 <font size="2">名字：</font><INPUT type="text" name="AuthorText"  size=40 style="border: 1px solid #AAAAAA;" maxlength=50 value="<%=viewer%>"><br>
	　 <font size="2">主页：</font><INPUT type="text" name="HomepageText"  size=40 style="border: 1px solid #AAAAAA;" value="<%=hp%>" maxlength=255><br>
	　 <font size="2">内容：</font><br/>　 <TEXTAREA name="ContentText" rows="5" cols="45"></TEXTAREA></p>
    　　　<INPUT type="submit" value="提交" style="border: 1px solid #AAAAAA;background-color: #CCD4E0;"></form><br/>
</td></tr></table>
      </TD>
      <TD VALIGN=top WIDTH=10>
      </TD>
   </TR>
</TABLE>
</BODY>
<script type="text/javascript">tagging();</script>
</html>
