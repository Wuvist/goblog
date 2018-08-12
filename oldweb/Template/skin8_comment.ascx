<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>  
	<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<link rel=stylesheet href="/Template/skin8.css" type=text/css>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body bgcolor="#000000" text="#000000">
<table border="0"
background="/images/skins/skin8/mettest3.jpg"  bgcolor="ffffff"
class="borrrder" cellpadding="5" cellspacing="0" width="750" height="100%" align="center">
<tr><td colspan="2" valign="bottom" align="left" height="160" 
background="/images/skins/skin8/abuseshawdows.jpg">
</td></tr> 
<td colspan="2"  bgcolor="666666"height="20"> 
<div align="right" >
<font color="999999">
<i><%=path%></i>
</font>
</div>
</td>
<tr valign="top">

<td  width="550" align="justify" valign="top" height="100%" >
<div class="main">
	<div class="datetimeheader">
<%=title%>
</div><p><%=content%> </p>
<div id="tags"><%=tags%></div>
<div class="datetimefooter">
posted at 
<%=adddate%> by  
<%=nick%> <br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div><br><br>

	<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a>
	   <div class="post">
	   <div id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("author")%> 在 <%#Container.DataItem("add_date")%> 说:</div>
		<%# Container.DataItem("content") %><br>
		</div>
		<br/></div>
	</ItemTemplate>
	</asp:Repeater>
	 <p><%=ph.g("请发表评论")%></p>
	<p><form action="/savecomment.aspx" method=post>
		<input type=hidden name="article_id" value="<%=id%>">
	　 <font size="2">标题：</font><INPUT type="text" name="TitleText" size=40 class="input" value="回复：<%=title%>" maxlength=50><br>
	　 <font size="2">名字：</font><INPUT type="text" name="AuthorText"  size=40 class="input" maxlength=50 value="<%=viewer%>"><br>
	　 <font size="2">主页：</font><INPUT type="text" name="HomepageText"  size=40 class="input" value="<%=hp%>" maxlength=255><br>
	　 <font size="2">内容：</font><br/>　 <TEXTAREA name="ContentText" rows="5" cols="50"></TEXTAREA></p>
    　　　<INPUT type="submit" value="提交" style="border: 1px solid #AAAAAA;background-color: #999999;"></form><br/
</td>

<td  width="150"  height="100%" > 
<div class="side" align="left">
<div class="sidetop"><br>
<%=ph.g("导航")%>
</div>
<a href="/"><%=ph.g("博客风")%></a><BR>
<a href="/<%=uid%>/"><%=path%>首页</a><BR>
<a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR>

<br><br><br><div class="sidetop">
<%=ph.g("网志介绍")%>
</div>
<p><Acme:userpic runat="server" ID="userpic1" deco="<div align='center'><img src='" enddeco="'></div>"/><%=uinfo%></p>

<br><br><br><div class="sidetop">
<%=ph.g("网志分类")%>
</div>
<asp:Repeater id="dlCate" runat="server">
	<ItemTemplate>
		<br><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)
	</ItemTemplate>
</asp:Repeater>



<br><br><br><div class="sidetop">
<%=ph.g("网志存档")%>
</div>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate>
		<br><a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a>
	</ItemTemplate>
</asp:Repeater>


<br><br><br><div class="sidetop">
<%=ph.g("个人链接")%>
</div>
<asp:Repeater id="links" runat="server">
	<ItemTemplate>
		<br><a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a>
	</ItemTemplate>
</asp:Repeater>

<br><br><br><div class="sidetop">
<%=ph.g("累计浏览")%>：<%=count%><br></div>
<Acme:Message runat="server" ID="Message1"/>
</div>
</td>


</tr>
<tr><td colspan="2"  valign="bottom" align="right" class="small" height="10" bgcolor="666666">
| maystar designs |</td></tr></table>
</body><script type="text/javascript">tagging();</script>
</html>