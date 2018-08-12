<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tDefault.ascx.vb" Inherits="Blogwind.tDefault" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html>
<head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<script type="text/javascript" src="/icon.js"></script>
<title><%=path%></title>
<link rel="stylesheet" href="/Template/skin3.css" type="text/css"> 
</head>

<body bgcolor="#FDFFFC">
<div id="Layer1" style="position:absolute; width:100%; height:100%; z-index:1; top: 0; background:#FDFFFC;">
<br>
<div align="center">
  <center>
  <table border="0" width="750" height="550" cellspacing="0" cellpadding="0" class=borderrr>
    <tr>
      <td width="750" colspan="2" height="85" valign="bottom" align="right">
        <h1><%=path%>　</h1>
      </td>
    </tr>
    <tr>
      <td width="750" colspan="2" height="40"></td>
    </tr>
    <tr>
      <td width="210" height="430" valign="top">
      <div class="blog" style="width: 210; height: 430">
	  <p class=head><%=ph.g("导航")%></p>
      <p><a href="/<%=uid%>/"><%=path%>首页</a><BR>
  <a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR>
  <A href="#"><IMG alt="RSS频道" src="../images/xml.gif" border="0" width="36" height="14"></A></p>
      <p class=head><%=ph.g("网志介绍")%></p>
	  <Acme:userpic runat="server" ID="userpic1"/>
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
	<p class=head><%=ph.g("累计浏览")%>：<%=count%></p><Acme:Message runat="server" ID="Message1"/></div>
      </td>
      <td width="540" height="430" valign="top">
      <div class="blog" style="width: 500; height: 430">
	  <asp:Repeater id="dlArticles" runat="server">
		<ItemTemplate>
				<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml"><div class="date"><p><font size="3"><b><%# Container.DataItem("title") %></p></b></font> </div></a>
			   <div align="justify"><%#Container.DataItem("content")%> </div><br>
			   <div align="right"><%=nick%> @<%# Container.DataItem("add_date") %>  
			<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml">评论:<%#Container.DataItem("comment")%></a></div>
		</ItemTemplate>
	</asp:Repeater><uc1:userpaging runat="server" ID="paging"/>
      </div>
      </td>
    </tr>
  </table>
</div>
</div>
</body>
</html>