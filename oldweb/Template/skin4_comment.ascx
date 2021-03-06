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
<link rel="stylesheet" href="/Template/skin4.css" type="text/css"> 
</head>

<body>

<div align="center">
  <center>
  <table width="564" height="584" cellspacing="4" cellpadding="4" bgcolor="#FFFFFF">
    <tr>
      <td width="540" height="98" valign="bottom" colspan="2">
		<img src="/images/skins/skin4.png">
        <p align="center"><font color="#CC00CC" size="5"><b>::<%=path%>::</b></font></p>
      </td>
    </tr>
    <tr>
      <td width="374" height="430" valign="top">
      <div id="blog" style="width: 368; height: 423; background-color: ffffff">
      <p class="title"><b><%=title%></b></p>
		   <div class="blogs"><%=content%></div><br>
		   <div id="tags"><%=tags%></div>
		   <div class="main">>><%=nick%>　@<%=adddate%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div><br><br>
				<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a>
		   <div class="main">		<div id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("author")%> 在 <%#Container.DataItem("add_date")%> 说:</div><br />
		<%# Container.DataItem("content") %><hr>
		</div>
		<br/></div>
	</ItemTemplate>
	</asp:Repeater>
    <p><%=ph.g("请发表评论")%><form action="/savecomment.aspx" method=post>
		<input type=hidden name="action" value="add">
		<input type=hidden name="article_id" value="<%=id%>"></p>
	<p><INPUT type="hidden" name="TitleText" value="回复:<%=title%>"><br>　 <font size="2">名字：</font><INPUT type="text" name="AuthorText"  size=40 style="border: 1px solid #AAAAAA;" value="<%=viewer%>"><br>　 <font size="2">主页：</font><INPUT type="text" name="HomepageText"  size=40 style="border: 1px solid #AAAAAA;" value="<%=hp%>"><br>　 <font size="2">内容：</font><br/>　 <TEXTAREA name="ContentText" rows="8" cols="40"></TEXTAREA></p>
    　　　<INPUT type="submit" value="提交" style="border: 1px solid #AAAAAA;background-color: #cc66ff;">
	</form>
      </div>

      
      
      </td>
      <td width="162" height="430" valign="top">
      <div id="sidebar" style="width: 155; height: 426; background-color: ffffff">
      <b><span class="title"><%=ph.g("导航")%></span></b><br>
		<a href="/"><%=ph.g("博客风")%></a><BR>
		<a href="/<%=uid%>/"><%=path%>首页</a><BR>
		<a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR>
	  <br>
      </font>
      <p>
      <b><span class="title"><%=ph.g("网志介绍")%></span></b><br>
      <span class="main"><Acme:userpic runat="server" ID="userpic1"/><%=uinfo%><br>
      </span>
      <p><b><span class="title"><%=ph.g("网志分类")%></span></b><br>
      <asp:Repeater id="dlCate" runat="server">
						<ItemTemplate>
							<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %>(<%# Container.DataItem("files") %>)</a><br>
						</ItemTemplate>
					</asp:Repeater>
      <br></p>
      <p><b><span class="title"><%=ph.g("网志存档")%></span></b><br>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate>
		<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
	</ItemTemplate>
</asp:Repeater>
      <br></p>
      <p><b><span class="title"><%=ph.g("个人链接")%></span></b><br>
      <asp:Repeater id="links" runat="server">
						<ItemTemplate>
							<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
						</ItemTemplate>
					</asp:Repeater>
	<br></p>
      <p><b><span class="title"><%=ph.g("累计浏览")%>: <%=count%></span></b></p>
<Acme:Message runat="server" ID="Message1"/></div>

      
      
      </td>
    </tr>
    <tr>
      <td width="540" height="42" colspan="2">
        <table width="101%" cellspacing cellpadding height="28">
          <tr>
            <td width="8%" height="28" bgcolor="#FFCCFF"></td>
            <td width="2%" height="28"></td>
            <td width="8%" height="28" bgcolor="#FF99FF"></td>
            <td width="2%" height="28"></td>
            <td width="8%" height="28" bgcolor="#FF66FF"></td>
            <td width="2%" height="28"></td>
            <td width="8%" height="28" bgcolor="#CC00CC"></td>
            <td width="2%" height="28"></td>
            <td width="8%" height="28" bgcolor="#CC66FF"></td>
            <td width="2%" height="28"></td>
            <td width="8%" height="28" bgcolor="#CC99FF"></td>
            <td width="2%" height="28"></td>
            <td width="8%" height="28" bgcolor="#CCCCFF"></td>
          <td width="34%" height="28"></td>
        </tr>
      </table>
    </td>
    </tr>
  </table>
</div>

</body>

<script type="text/javascript">tagging();</script>
</html>
