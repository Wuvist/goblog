<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tDefault.ascx.vb" Inherits="Blogwind.tDefault" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="uc1" TagName="userpaging" Src="../userpaging.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>

<html><head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<link rel="stylesheet" type="text/css" href="/Template/skin21.css">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<script type="text/javascript" src="/icon.js"></script>
<title><%=path%></title>


</head>

<center><img style="position:absolute; top:0px; left:0px; " 
img src="/images/skins/skin21/couple.jpg" style="filter:alpha(opacity=100)"></center>

<DIV id=table 
style="border:1px none ; padding:5px; z-index: 2; left: 50; position: 
absolute; top: 370; height: 250; width: 190; overflow: auto; FILTER: chroma(color=#FFFFFF)" allowTransparency>
<font color=9FAACD>
<p align="centre">
<asp:Repeater id="dlCate" runat="server">
	<ItemTemplate>
	<a class="b" href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %>(<%# Container.DataItem("files") %>)</a><br>
	</ItemTemplate>
</asp:Repeater>
</p>
</font>
</div>

<DIV id=table 
style="border:1px none ; padding:5px; z-index: 2; left: 20; position: 
absolute; top: 675; height: 250; width: 220; overflow: auto; FILTER: chroma(color=#FFFFFF)" allowTransparency>
<font color=C0CD94>
<p align="centre">
<a href="http://www.blogwind.com">Blogwind Home</a><br>
<a href="/<%=uid%>/"><%=path%> Home</a><br>
<a href="/email.aspx?user=<%=uid%>">Contact <%=nick%></a><br>
<Acme:userpic runat="server" ID="userpic1" deco="<table align='center'><tr><td><div class='shadow1'><div class='shadow2'><div class='shadow3'><div class='alpha'><img src='" enddeco="'></div></div></div></div></td></tr></table>"/>
<br>
<b><MARQUEE scrollAmount=1 width="3%">*</marquee><%=ph.g("个人链接")%><MARQUEE scrollAmount=1 direction=right width="3%"> *</MARQUEE></b><br>
<asp:Repeater id="links" runat="server">
	<ItemTemplate>
	<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
	</ItemTemplate>
</asp:Repeater>
</p>
</font>
</div>

<DIV id=table 
style="border:1px none ; padding:5px; z-index: 2; left: 720;position: 
absolute; top: 370; height: 250; width: 190; overflow: auto; FILTER: chroma(color=#FFFFFF)" allowTransparency>
<font color=9FAACD>
<p align="centre">
<asp:Repeater id="archive" runat="server">
	<ItemTemplate><a class="b" href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("month") %>-<%# Container.DataItem("year") %></a><br>
	</ItemTemplate>
</asp:Repeater>
</p>
</font>
</div>


<DIV id=table 
style="border:1px none ; padding:5px; z-index: 2; left: 735;position: 

absolute; top: 675; height: 245; width: 220; overflow: auto; FILTER: chroma(color=#FFFFFF)" allowTransparency>
<font color=FFB87B>
<br><br>
<p align="centre">
<Acme:Message runat="server" ID="Message1"/>
</p>
</font>
</div>

<DIV id=table 
style="border:1px ; padding:5px; z-index: 2; left:280;position: absolute; top: 
180; height: 750; width: 420 ;overflow: auto; FILTER: chroma(color=#FFFFFF)" allowTransparency>
<font color=7C88A3>
<p> 
<div style="position:absolute; top:0; left:0; width:350;">
<table><tr><td><b><%=path%></b></td><td>&nbsp;</td><td><%=ph.g("累计浏览")%>: <%=count%></td></tr></table>
<div align="right"><p><i><%=uinfo%></i></p></div>
<asp:Repeater id="dlArticles" runat="server">
	<ItemTemplate>
<font size=4>
<a class="s" href="/<%# uid & "/" & Container.DataItem("index")%>.shtml" class="title"><%# Container.DataItem("title") %></a>
</font>
<p>
<div align="left"><%#Container.DataItem("content")%></div>
<div align="right">
-----------------------------------------------------------<br>
<%=nick%> winks at |<%# Container.DataItem("add_date") %>|</font>
<a href="/<%# uid & "/" & Container.DataItem("index")%>.shtml" class="s">Comments(<%#Container.DataItem("comment")%>)</a>
</div>
</p>
</ItemTemplate>
</asp:Repeater>
<uc1:userpaging runat="server" ID="paging"/>

</div>
</font>
</div>

</html>