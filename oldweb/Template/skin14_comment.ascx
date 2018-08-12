<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<link rel="stylesheet" href="/Template/skin14.css" type="text/css">
</head>

<body topmargin="0" leftmargin="0">
<table width=950 border="0" align="center">
  <tr > 
    <td background="/images/skins/skin14/grey-banner.gif" width=900 height="160" colspan="2" align="right" valign="top"><p class="header"><%=path%></p></td>
  </tr>
  <tr> 
    <td width=750 align="left" valign="top" bordercolor="#CCCCCC"> 
<DIV class="blog">
<p><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><font size="3" color=black><b><%=title%></b></font></p>
   <div align="justify"><%=content%> </div><br>
   <div id="tags"><%=tags%></div>
   <div align="right"><%=nick%> <img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=adddate%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div>
	<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><hr>
	<%# Container.DataItem("title") %><br>
	<%# Container.DataItem("content") %><br>
	<div style="text-align:right" id="_r<%# Container.DataItem("cid") %>">By:<%#Container.DataItem("author")%> <img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%#Container.DataItem("add_date")%></div><br/></div>
	</ItemTemplate>
	</asp:Repeater>
<center><p>请<%=ph.g("请发表评论")%></p></center>
<p><form action="/savecomment.aspx" method=post>
	<input type=hidden name="action" value="add">
	<input type=hidden name="article_id" value="<%=id%>">
	<blockquote><font size="2">标题：</font><INPUT type="text" name="TitleText" size=40 style="border: 1px solid #AAAAAA;" value="回复：<%=title%>" maxlength=50><br></blockquote>
	<blockquote><font size="2">名字：</font><INPUT type="text" name="AuthorText"  size=40 style="border: 1px solid #AAAAAA;" maxlength=50 value="<%=viewer%>"><br></blockquote>
	<blockquote><font size="2">主页：</font><INPUT type="text" name="HomepageText"  size=40 style="border: 1px solid #AAAAAA;"  value="<%=hp%>" maxlength=255><br></blockquote>
	<blockquote><font size="2">内容：</font><br/><TEXTAREA name="ContentText" rows="5" cols="50"></TEXTAREA></p></blockquote>
	<blockquote><INPUT type="submit" value="提交" class="logbutton"></blockquote></form><br/>
</div></td>
    <td width=200 background="/images/skins/skin14/grey-colback.gif" align="left" valign="top"><div class="blog" style="width: 200">
	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("导航")%></p><a href="/"><%=ph.g("博客风")%></a><BR>
      	<a href="/<%=uid%>/"><%=path%>首页</a><BR>
  	<a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR></p>
      	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("网志介绍")%></p>
      	<Acme:userpic runat="server" ID="userpic1" deco="<img src='" enddeco="'>"/>
		<p><%=uinfo%></p>
      	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("个人链接")%></p>
	<asp:Repeater id="links" runat="server">
		<ItemTemplate>
			<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
		</ItemTemplate>
	</asp:Repeater></p>
      	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("网志分类")%></p>
	<asp:Repeater id="dlCate" runat="server">
		<ItemTemplate><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
		</ItemTemplate>
	</asp:Repeater>
    	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("网志存档")%></p>
	<asp:Repeater id="archive" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
		</ItemTemplate>
	</asp:Repeater>
	<p class=head><img border="0" src="/images/skins/skin14/grey-comment.gif" width="15" height="15"><%=ph.g("累计浏览")%>：<%=count%><br><br><br><br>
	<Acme:Message runat="server" ID="Message1"/></p></div></td>
  </tr>
</table><br><br><br>
</body>
<script type="text/javascript">tagging();</script>
</html>