<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html>
<head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<link rel="stylesheet" href="/Template/skin12.css" type="text/css"> 
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
</head>

<body topmargin="0" leftmargin="0" link="#808000" vlink="#808000" alink="#808000">
<table border="0" width="780" cellspacing="0" cellpadding="0" align=left>
<tr><td height=153 colspan=2 background="/images/skins/skin12/mao.jpg"><p class="header"><b><%=path%><br>
</td></tr>
<tr><td valign="top" width="580">
<DIV class="blog">
<p><img border="0" src="/images/skins/skin12/mao2.jpg" width="77" height="50"><font size="3" color=black><b><%=title%></b></font></p>
   <div align="justify"><%=content%> </div><br>
   <div id="tags"><%=tags%></div>
   <div align="right"><%=nick%> <img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=adddate%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div>
	<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><hr>
		<%# Container.DataItem("title") %><br>
		<%# Container.DataItem("content") %><br>
		<div style="text-align:right" id="_r<%# Container.DataItem("cid") %>">By:<%#Container.DataItem("author")%> <img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%#Container.DataItem("add_date")%></div>
		<br/></div>
	</ItemTemplate>
	</asp:Repeater>
						<p><%=ph.g("请发表评论")%></p>
	<p><form action="/savecomment.aspx" method=post>
		<input type=hidden name="action" value="add">
		<input type=hidden name="article_id" value="<%=id%>">
	　 <font size="2">标题：</font><INPUT type="text" name="TitleText" size=40 style="border: 1px solid #AAAAAA;" value="回复：<%=title%>" maxlength=50><br>
	　 <font size="2">名字：</font><INPUT type="text" name="AuthorText"  size=40 style="border: 1px solid #AAAAAA;" maxlength=50 value="<%=viewer%>"><br>
	　 <font size="2">主页：</font><INPUT type="text" name="HomepageText"  size=40 style="border: 1px solid #AAAAAA;"  value="<%=hp%>" maxlength=255><br>
	　 <font size="2">内容：</font><br/>　 <TEXTAREA name="ContentText" rows="5" cols="50"></TEXTAREA></p>
    　　　<INPUT type="submit" value="提交" class="logbutton"></form><br/>
</div>
</td><td valign="top" align=right width="200">
	      <div class="blog" style="width: 200">
	  <p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("导航")%></p><a href="/"><%=ph.g("博客风")%></a><BR>
      <a href="/<%=uid%>/"><%=path%>首页</a><BR>
  <a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR></p>
      <p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("网志介绍")%></p>
	  <Acme:userpic runat="server" ID="userpic1"/>
      <p><%=uinfo%></p>
      <p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("个人链接")%></p><asp:Repeater id="links" runat="server">
						<ItemTemplate>
							<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
						</ItemTemplate>
					</asp:Repeater></p>
      <p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("网志分类")%></p>
	  <asp:Repeater id="dlCate" runat="server">
						<ItemTemplate><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
						</ItemTemplate>
					</asp:Repeater>
      <p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("网志存档")%></p>
<asp:Repeater id="archive" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
		</ItemTemplate>
	</asp:Repeater>
	<p class=head><img border="0" src="/images/skins/skin12/miao.jpg" width="24" height="19"><%=ph.g("累计浏览")%>：<%=count%><br>
	<Acme:Message runat="server" ID="Message1"/></p></div></td>
</tr></table>
</body><script type="text/javascript">tagging();</script>
</html>