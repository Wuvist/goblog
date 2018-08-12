<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html>
<head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<link rel="stylesheet" href="/Template/skin11.css" type="text/css"> 
</head>

<body vlink="#2D6382" alink="#3A81AB" link="#3A81AB">
<div id="Layer1" style="position:absolute; width:100%; height:100%; z-index:1; top: 0; background:#000000">
<div align="center">
  <center>

<table border="0" width="90%" cellspacing="0" cellpadding="0" height="100%">
  <tr>
    <td width="167" valign="top" class=borderrr><img border="0" src="/images/skins/skin11/lmimage.jpg" width="200" height="455"></td>
    <td height="100%" valign="top" class=borderrr>


<div class="blog" style="width: 100%; height=100%">
	<div class="date"><p><font size="3"><b><%=title%></p></b></font></div>
   <div align="justify"><%=content%> </div><br>
   <div id="tags"><%=tags%></div>
   <div align="right"><%=nick%> @<%=adddate%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div>
	<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><hr>
		<%# Container.DataItem("title") %><br>
		<%# Container.DataItem("content") %><br>
		<div style="text-align:right" id="_r<%# Container.DataItem("cid") %>">By:<%#Container.DataItem("author")%> @<%#Container.DataItem("add_date")%></div>
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
    　　　<INPUT type="submit" value="提交" class="logbutton"></form><br/><br/>
   </div>
      </div>
      </td>

    <td width="200" valign="top" class=borderrr align=center>
	      <div class="blog" style="width: 200;height: 100%">
	  <p class=head><%=ph.g("导航")%></p>
	  <p><a href="/"><%=ph.g("博客风")%></a><BR>
      <a href="/<%=uid%>/"><%=path%>首页</a><BR>
  <a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR>
  <A href="#"><IMG alt="RSS频道" src="../images/xml.gif" border="0" width="36" height="14"></A></p>
      <p class=head><%=ph.g("网志介绍")%></p>
	  <Acme:userpic runat="server" ID="userpic1" deco="<div align='center'><img src='" enddeco="'></div>"/>
      <p><%=uinfo%></p>
      <p class=head><%=ph.g("个人链接")%></p>
      <p><asp:Repeater id="links" runat="server">
						<ItemTemplate>
							<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
						</ItemTemplate>
					</asp:Repeater></p>
      <p class=head><%=ph.g("网志分类")%></p>
	  <asp:Repeater id="dlCate" runat="server">
						<ItemTemplate><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
						</ItemTemplate>
					</asp:Repeater>
      <p class=head><%=ph.g("网志存档")%></p>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate>
		<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
	</ItemTemplate>
</asp:Repeater>
	<p class=head><%=ph.g("累计浏览")%>：<%=count%></p>
	<Acme:Message runat="server" ID="Message1"/>

      </div>
</td>
  </tr>
</table>

</div>
</div>
</body><script type="text/javascript">tagging();</script>
</html>
