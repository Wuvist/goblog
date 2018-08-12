<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>

<html><HEAD><link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<LINK media=all href="/Template/skin26.css" type=text/css rel=Stylesheet>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<BODY>
<DIV class=pagelayout>
<DIV id=header><SPAN><A id=Header1_HeaderTitle href="#"><%=path%></A><BR><%=uinfo%></SPAN></DIV>
<DIV id=menu><UL class=list><Acme:userpic runat="server" ID="userpic1"/></ul>
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
<UL class=list><Acme:Message runat="server" ID="Message1"/></ul>
<DIV class=spacer>&nbsp;</DIV></DIV>
<DIV id=main>
<DIV class=post>
<DIV class=postTitle><%=title%></DIV>
<DIV class=postText><%=content%>
<div id="tags"><%=tags%></div></DIV>
<DIV class=postFoot>
<%=nick%> 发表于<%=adddate%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></DIV></DIV>
<BR>
<BR>
<DIV id=comments>
<H3></H3><asp:Repeater id="dlComment" runat="server">
<ItemTemplate><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><DIV class=post id="_c<%# Container.DataItem("cid") %>" onMouseOver="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)">
<DIV class=postTitle><%#Container.DataItem("author")%> 在 <%#Container.DataItem("add_date")%> 说:<div id="_r<%# Container.DataItem("cid") %>"></div></DIV>
<DIV class=postText><%# Container.DataItem("content") %></DIV></DIV><BR>
</ItemTemplate>
</asp:Repeater></DIV>
<DIV class=post>
<DIV class=moreinfo>
<DIV id=postcomment>
<form action="/savecomment.aspx" method=post name="PostComment" id="PostComment">
<input type=hidden name="action" value="add">
<input type=hidden name="article_id" value="<%=id%>">
<input id="TitleText" name="TitleText" type="hidden" value="回复：<%=title%>">
<fieldset class="postFSet"><legend><%=ph.g("请发表评论")%></legend>
	　 <font size="2">大名：</font><INPUT type="text" name="AuthorText"  size=40 style="border: 1px solid #AAAAAA;" maxlength=50 value="<%=viewer%>"><br>
	　 <font size="2">主页：</font><INPUT type="text" name="HomepageText"  size=40 style="border: 1px solid #AAAAAA;" value="<%=hp%>" maxlength=255><br>
	　 <font size="2">内容：</font><br/>　 <TEXTAREA name="ContentText" rows="5" cols="50"></TEXTAREA></p>
    　　　<INPUT type="submit" value="提交" class="logbutton"></form></fieldset>
</form>
</DIV></DIV></DIV>

<DIV class=spacer>&nbsp;</DIV></DIV>
<P id=footer>Powered by: <BR><A id=Footer2_Hyperlink2 href="http://www.blogwind.com" name=Hyperlink1>Blogwind</a></P></DIV>

<script type="text/javascript">tagging();</script></body>
</html>