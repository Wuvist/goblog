<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcate.ascx.vb" Inherits="Blogwind.tcate" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<title><%=path%></title>
<link rel="stylesheet" href="/Template/skin30.css" type="text/css"> 
<script type="text/javascript" src="/icon.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
<DIV class=pagelayout>
<DIV id=header><SPAN><A id=Header1_HeaderTitle href="#"><%=path%></A><BR><%=uinfo%></SPAN></DIV>
<DIV id=menu><UL class=list><DIV class=spacer>&nbsp;</DIV><Acme:userpic runat="server" ID="userpic1"/></UL>
<H1><%=ph.g("导航")%></H1>
<UL class=list>
<LI class=listitem><a href="/"><%=ph.g("博客风")%></a>
<LI class=listitem><a href="/<%=uid%>/" class=listitem id=MyLinks1_PersonalHome title="访问 <%=nick%> 的首页"><%=path%>首页</a>
<LI class=listitem><A class=listitem id=MyLinks1_ContactLink accessKey=9 href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%>作者</A></LI></UL>
<H1><%=ph.g("网志分类")%></H1>
<UL class=list>
<asp:Repeater id="dlCate" runat="server">
<ItemTemplate>
<LI class=listitem><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %>(<%# Container.DataItem("files") %>)</a></ItemTemplate>
</asp:Repeater></LI></UL>
<H1><%=ph.g("网志存档")%></H1>
<UL class=list>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate><LI class=listitem><a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a></ItemTemplate>
</asp:Repeater></LI></UL>
<H1><%=ph.g("个人链接")%></H1>
<UL class=list>
<asp:Repeater id="links" runat="server">
	<ItemTemplate><LI class=listitem><a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a></ItemTemplate>
</asp:Repeater></LI></UL>
<H1><%=ph.g("累计浏览")%>: <%=count%></H1>
<UL class=list><Acme:Message runat="server" ID="Message1"/></UL>
<DIV class=spacer>&nbsp;</DIV></DIV>
<DIV id=main>

<asp:Repeater id="dlArticles" runat="server">
	<ItemTemplate><DIV class=post>
<DIV class=postTitle><a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml" class="title"><%# Container.DataItem("title") %></a></DIV>
<DIV class=postFoot><%=nick%> 发表于 <%# Container.DataItem("add_date") %> <a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"> 评论:<%#Container.DataItem("comment")%></a></DIV>
</DIV><DIV class=spacer>&nbsp;</DIV>
	</ItemTemplate>
</asp:Repeater>
<DIV class=spacer>&nbsp;</DIV></DIV>
<P id=footer>Powered by: <BR><A id=Footer2_Hyperlink2 href="http://www.blogwind.com" name=Hyperlink1>Blogwind</a></P></DIV>

</html>