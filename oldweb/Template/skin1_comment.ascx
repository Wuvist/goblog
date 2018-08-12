<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<link rel="stylesheet" href="/Template/skin1.css" type="text/css"> 
</head>

<body style="FILTER: progid:DXImageTransform.Microsoft.Gradient(endColorstr='#ffffff', startColorstr='#9999CC', gradientType='1'); WIDTH: 100%; HEIGHT: 100%">
<div id="Layer1" style="position:absolute; width:400px; height:100%; z-index:1; top: 0; left: 350;" class="entry"><img src="/images/skins/bluesky/bluesky.jpg" width="400" height="299"><br>
<div class="main">
<p class="title"><%=title%></p>
<%=content%><br><br>
<div id="tags"><%=tags%></div>
<div class="com"><%=nick%> @ <%=adddate%><br><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div>
		<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a>
		   <div class="post"><div style="text-align:left" id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("author")%> <%=ph.g("写了")%> @ <%#Container.DataItem("add_date")%>：</div><br><div style="word-break:break-all">
		<%# Container.DataItem("content") %></div>		
		</div></div>
	</ItemTemplate>
	</asp:Repeater>
    <p>　<%=ph.g("请发表评论")%></p>
	<p><form action="/savecomment.aspx" method=post>
		<input type=hidden name="article_id" value="<%=id%>">
		<INPUT type="hidden" name="TitleText" value="回复：<%=title%>" maxlength=50>
	　 <font size="2"><%=ph.g("名字")%>：</font><INPUT type="text" name="AuthorText"  size=40 class="input" maxlength=50 value="<%=viewer%>"><br>
	　 <font size="2"><%=ph.g("主页")%>：</font><INPUT type="text" name="HomepageText"  size=40 class="input" value="<%=hp%>" maxlength=255><br>
	　 <font size="2"><%=ph.g("内容")%>：</font><br/>　 <TEXTAREA name="ContentText" rows="5" cols="40"></TEXTAREA></p>
    　　　<INPUT type="submit" value="<%=ph.g("发表")%>" style="border: 1px solid #AAAAAA;background-color: #CCD4E0;"></form><br/><br/>
	</div>
</div>

<div id="Layer2" style="position:absolute; width:200px; height:115px; z-index:2; left: 125; top: 0;" class="navbar">
  <p class="nav"><%=ph.g("导航")%></p>
  <p>
  
  <a href="/"><%=ph.g("博客风")%></a><BR>
  <a href="/<%=uid%>/"><%=path%><%=ph.g("首页")%></a><BR>
  <a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a></p>

  <p class="nav"><%=ph.g("网志介绍")%></p>
  <Acme:userpic runat="server" ID="userpic1" deco="<div align='center'><img src='" enddeco="'></div>"/>
  <p><%=uinfo%></p>
  <p class="nav"><%=ph.g("网志分类")%></p>
  <p><asp:Repeater id="dlCate" runat="server">
						<ItemTemplate>
							<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
						</ItemTemplate>
					</asp:Repeater>
  <p class="nav"><%=ph.g("网志存档")%></p>
  <p><asp:Repeater id="archive" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %> - <%# Container.DataItem("month") %></a><br>
		</ItemTemplate>
	</asp:Repeater></P>
  <p class="nav"><%=ph.g("个人链接")%></p>
  <p><asp:Repeater id="links" runat="server">
						<ItemTemplate>
							<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
						</ItemTemplate>
					</asp:Repeater></p>  
  <p class="nav"><%=ph.g("累计浏览")%>: <%=count%></p>
<Acme:Message runat="server" ID="Message1"/>
</div>
</body>
<script type="text/javascript">tagging();</script>
</html>