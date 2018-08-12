<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html>
<head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link> <TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<link rel=stylesheet href="/Template/skin10.css" type=text/css> 
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body background="/images/skins/skin10/momentbg.jpg"  leftmargin="0" bottommargin="0" topmargin="0" rightmargin="0">

<div align="left">

<table border="0" width="760" cellpadding="0" cellspacing="0">
<tr><td><img src="/images/skins/skin10/moment1.jpg" width="760" height="230"></tr>
<tr><td><img src="/images/skins/skin10/moment2.jpg" width="760" height="180"></tr>
<tr><td><img src="/images/skins/skin10/moment3.jpg" width="760" height="169"></tr>
</table>

<div id="me">
<table border="0" width="180" cellpadding="0" cellspacing="0">
<tr>
<td valign="top" class=tab>
<img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("网志介绍")%></b>
<br><Acme:userpic runat="server" ID="userpic1"/><%=uinfo%>
<p>
<img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("导航")%></b>
<br>&nbsp;<a href="/"><%=ph.g("博客风")%></a>
<br>&nbsp;<a href="/<%=uid%>/"><%=path%>首页</a>
<br>&nbsp;<a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a>
<p><img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("网志分类")%></b>
<asp:Repeater id="dlCate" runat="server">
<ItemTemplate>
<br>&nbsp;<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)
</ItemTemplate>
</asp:Repeater>
<p><img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("网志存档")%></b>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate>
<br>&nbsp;<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a>
	</ItemTemplate>
</asp:Repeater>
<p>
<img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("个人链接")%></b>
<asp:Repeater id="links" runat="server">
	<ItemTemplate>
	<br>&nbsp;<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a>
	</ItemTemplate>
	</asp:Repeater>
<p>
<img src="/images/skins/skin9/balleticon.jpg" border=0 width="19" height="22">
<b><%=ph.g("累计浏览")%>：<%=count%></b>
<br>
<Acme:Message runat="server" ID="Message1"/>
</td>
</tr>
</table>
</div>



<div id="blog">
<table border="0" width="400" cellpadding="0" cellspacing="0">
<tr>
<td valign="top"><align=justify>
<P>
<div class="DateHeader"><%=title%></div><br><br>
<div class="msg"><%=content%></div><br><br>
		<div id="tags"><%=tags%></div>
<span class="PostFooter">- post by <%=nick%>　@<%=adddate%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></span>
<br><br>		<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a>
		   <div class="post"><div id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("author")%> 在 <%#Container.DataItem("add_date")%> 说:</div><br>
		<%# Container.DataItem("content") %><br>
		</div><hr></div>
	</ItemTemplate>
	</asp:Repeater>
    <p><%=ph.g("请发表评论")%></p>
	<p><form action="/savecomment.aspx" method=post>
		<input type=hidden name="article_id" value="<%=id%>">
<INPUT type="hidden" name="TitleText" value="回复：<%=title%>">
	　 <font size="2">名字：</font><INPUT type="text" name="AuthorText"  size=40 class="input" maxlength=50 value="<%=viewer%>"><br>
	　 <font size="2">主页：</font><INPUT type="text" name="HomepageText"  size=40 class="input" value="<%=hp%>" maxlength=255><br>
	　 <font size="2">内容：</font><br>　 <TEXTAREA name="ContentText" rows="5" cols="40"></TEXTAREA></p><br>
    　　　<INPUT type="submit" value="提交" style="border: 1px solid #AAAAAA;background-color: #A0A0A0;"></form><br/>

   </div></FONT>
<p>
</td>
</tr>
</table>
</div>
</body>
<script type="text/javascript">tagging();</script>
</html>