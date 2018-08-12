<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<link rel="stylesheet" href="/Template/skin5.css" type="text/css">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body bgcolor="ffffff" marginwidth="0" marginheight="0">

<table width="100%" border="0" cellspacing="0" cellpadding="0">
<tr bgcolor="#ffffff">
<td class="name"  width="100%" height="60"><%=path%>
</td>
</tr>
</table>
<table width="100%" border="0" cellspacing="0" cellpadding="0">
<tr bgcolor="#990000">
<td class="description" align="right" width="100%" height="30"><%=uinfo%></td>
</tr>
</table>

<table width="100%" cellpadding="0" cellspacing="0"><tr><td valign="top">
<div class="blog">
<p class="title"><%=title%></p>
	   <%=content%>
	   <div id="tags"><%=tags%></div>
	   <div class="com"><%=nick%>　@<%=adddate%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div><br><br>
		<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onMouseOver="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a>
	   <div class="post">	<div id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("author")%> 在 <%#Container.DataItem("add_date")%> 说:</div>
		<br>
	<%# Container.DataItem("content") %><hr />
	</div></div>
</ItemTemplate>
</asp:Repeater>
</div>
<p><%=ph.g("请发表评论")%></p>
	<p><form action="/savecomment.aspx" method=post>
		<input type=hidden name="action" value="add">
		<input type=hidden name="article_id" value="<%=id%>">
	　 <INPUT type="hidden" name="TitleText" value="回复：<%=title%>" maxlength=50><br>
	　 <font size="2">名字：</font><INPUT type="text" name="AuthorText"  size=40 style="border: 1px solid #AAAAAA;" maxlength=50 value="<%=viewer%>"><br>
	　 <font size="2">主页：</font><INPUT type="text" name="HomepageText"  size=40 style="border: 1px solid #AAAAAA;" value="<%=hp%>" maxlength=255><br>
	　 <font size="2">内容：</font><br/>　 <TEXTAREA name="ContentText" rows="5" cols="50"></TEXTAREA></p>
    　　　<INPUT type="submit" value="提交" style="border: 1px solid #AAAAAA;background-color: #CCD4E0;"></form><br/><br/>
</div>
</TD>
<TD bgColor="#e0e0b1" height="500" width="150" valign="top"><div class="side">
<div class="sidetop">
 <%=ph.g("导航")%>
 </div> 
  <a href="/"><%=ph.g("博客风")%></a><BR>
  <a href="/<%=uid%>/"><%=path%>首页</a><BR>
  <a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR>
<br>

<div class="sidetop"><br>
个人档案
</div>
<Acme:userpic runat="server" ID="userpic1" deco="<div align='center'><br><img src='" enddeco="'></div>"/>
<%=uinfo%>
<br>
<br><br><br><div class="sidetop">
<%=ph.g("网志分类")%>
</div>
<asp:Repeater id="dlCate" runat="server">
<ItemTemplate>
	<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
</ItemTemplate>
</asp:Repeater>
<br><br><br><div class="sidetop">
<%=ph.g("网志存档")%>
</div>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate>
		<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
	</ItemTemplate>
</asp:Repeater>

<br><br><br><div class="sidetop">
<%=ph.g("个人链接")%>
</div>
<asp:Repeater id="links" runat="server">
	<ItemTemplate>
		<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
	</ItemTemplate>
</asp:Repeater>
<p class="nav"><%=ph.g("累计浏览")%>: <%=count%></p>
<br>
<Acme:Message runat="server" ID="Message1"/>
<br>
</div>
</TD>
</TR>
</TABLE>
</body>
<script type="text/javascript">tagging();</script>
</html>
