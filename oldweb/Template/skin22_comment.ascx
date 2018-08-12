<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<html>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<link rel="stylesheet" type="text/css" href="/Template/skin22.css">
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
</head>
<body class="bodyBox">
<table align="center" border="0" cellpadding="0" cellspacing="0" class="bodyTop"><tr><td></td></tr></table>
<table align="center" border="0" cellpadding="0" cellspacing="0" class="bodyBg"><tr><td class="banner"><table class="title" border="0" cellpadding="0" cellspacing="0"><tr><td width="20"></td><td><%=path%></td><td width="20"></td></tr></table>
<table class="link" border="0" cellpadding="0" cellspacing="0"><tr><td width="20"></td><td><%=uinfo%></td><td width="20"></td></tr></table></td></tr></table>

<table align="center" border="0" cellpadding="0" cellspacing="0" class="bodyBg"><tr><td class="menu"><table border="0" cellpadding="0" cellspacing="0" width="100%"><tr><td class="home"><input style="width:0px;height:0px;border:0px" id="HOME"></td><td id="login" class="text" align="right">
</td></tr></table></td></tr></table>

<table class="sysW770 bodyBg" align="center" border="0" cellpadding="0" cellspacing="0"><tbody><tr><td valign="top"><table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td class="box_1" valign="top"><div class="margin"></div><div class="margin" id="box_1"><div class="photo"><table width="100%" border="0" cellpadding="0" cellspacing="0"><tbody><tr><td class="top"></td></tr><tr><td class="mid" align="center"><table align="center" border="0" cellpadding="0" cellspacing="0"><tr><td align="center">
<Acme:userpic runat="server" ID="userpic1"/>
</td></tr>
<tr><td><div class="infoText"><%=nick%>的BLOG</div></td></tr></table>

</td></tr><tr><td class="bottom"></td></tr></tbody></table><table width="100%" border="0" cellpadding="0" cellspacing="0"><tr class="line"><td></td></tr></table></div><div id="labelFM000009"></div><div class="links"><table id="linksFM000008" width="100%" border="0" cellpadding="0" cellspacing="0"><tbody><tr><td>

<table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td class="nav"><%=ph.g("导航")%></td></tr></table></td></tr><tr  ><td class="mid"><table height="6"><tr><td></td></tr></table>

<table class="wd" border="0" cellpadding="0" cellspacing="0"><tr><td calss="ico" width="20" align="center" valign="top"><font style="font-size:6px;">■</font></td><td class="dashed"><table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td><div class="sysBr180"><a href="/"><%=ph.g("博客风")%></a></div></td></tr></table></td></tr></table>
<table class="wd" border="0" cellpadding="0" cellspacing="0"><tr><td calss="ico" width="20" align="center" valign="top"><font style="font-size:6px;">■</font></td><td class="dashed"><table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td><div class="sysBr180"><a href="/<%=uid%>/"><%=path%>首页</a></div></td></tr></table></td></tr></table>
<table class="wd" border="0" cellpadding="0" cellspacing="0"><tr><td calss="ico" width="20" align="center" valign="top"><font style="font-size:6px;">■</font></td><td class="dashed"><table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td><div class="sysBr180"><a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a></div></td></tr></table></td></tr></table>

<table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td  class="nav"><div class="sysBr180 text"><%=ph.g("网志分类")%></div></td></tr></table></td></tr><tr  ><td class="mid"><table height="6"><tr><td></td></tr></table>
<asp:Repeater id="dlCate" runat="server">
<ItemTemplate>
<table class="wd" border="0" cellpadding="0" cellspacing="0"><tr><td calss="ico" width="20" align="center" valign="top"><font style="font-size:6px;">■</font></td><td class="dashed"><table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td><div class="sysBr180"><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %>(<%# Container.DataItem("files") %>)</a></div></td></tr></table></td></tr></table>
</ItemTemplate>
</asp:Repeater>
<table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td  class="nav"><div class="sysBr180 text"><%=ph.g("网志存档")%></div></td></tr></table></td></tr><tr  ><td class="mid"><table height="6"><tr><td></td></tr></table>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate><table class="wd" border="0" cellpadding="0" cellspacing="0"><tr><td calss="ico" width="20" align="center" valign="top"><font style="font-size:6px;">■</font></td><td class="dashed"><table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td><div class="sysBr180"><a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a></div></td></tr></table></td></tr></table></ItemTemplate>
</asp:Repeater>
<table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td  class="nav"><div class="sysBr180 text"><%=ph.g("个人链接")%></div></td></tr></table></td></tr><tr  ><td class="mid"><table height="6"><tr><td></td></tr></table>
<asp:Repeater id="links" runat="server">
	<ItemTemplate><table class="wd" border="0" cellpadding="0" cellspacing="0"><tr><td calss="ico" width="20" align="center" valign="top"><font style="font-size:6px;">■</font></td><td class="dashed"><table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td><div class="sysBr180"><a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a></div></td></tr></table></td></tr></table></ItemTemplate>
</asp:Repeater>
<table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td class="nav"><div class="sysBr180 text"><%=ph.g("累计浏览")%>: <%=count%></div></td></tr></table></td></tr><tr  ><td class="mid"><table height="6"><tr><td></td></tr></table>
<table class="wd" border="0" cellpadding="0" cellspacing="0"><tr><td calss="ico" width="20" align="center" valign="top"><font style="font-size:6px;">■</font></td><td class="dashed"><table width="100%" border="0" cellpadding="0" cellspacing="0"><tr><td><div class="sysBr180"><Acme:Message runat="server" ID="Message1"/></div></td></tr></table></td></tr></table>


<table width="100%" border="0" cellpadding="0" cellspacing="0"><tr height="10"><td></td></tr><tr class="bottom"><td></td></tr></table></td></tr></tbody></table><table width="100%" border="0" cellpadding="0" cellspacing="0"><tr class="line"><td></td></tr></table></div><div id="labelFM000013">
</div></div><div class="margin"></div></td><td class="box_2" valign="top">

<div class="margin" id="box_2"><center><div class="article"><table border="0" cellpadding="0" cellspacing="0"><tr><td align="center"><table border="0" cellpadding="0" cellspacing="0"><tr><td class="up"><div class="sysBr500 title"><%=title%></div></td></tr></table></td></tr><tr><td class="aBody"><table align="center" border="0" cellpadding="0" cellspacing="0" class="dashed"><tr><td></td></tr></table><table align="center" border="0" cellpadding="0" cellspacing="0" class="description"><tr><td align="center"><div class="sysBr500 text" align="left">
<%=content%>
<div id="tags"><%=tags%></div>


<DIV></DIV></div></td></tr></table><table align="center" border="0" cellpadding="0" cellspacing="0" class="dashed"><tr><td></td></tr></table><table align="center" border="0" cellpadding="0" cellspacing="0" class="function"><tr><td><div class="com"><%=nick%> 发表于<%=adddate%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div></td></tr></table></td></tr></table><table align="center" border="0" cellpadding="0" cellspacing="0"><tr><td id="articleChild471facb1010000fe" style="display:none"></td></tr></table></div></center><center><table class="comment" border="0" cellpadding="0" cellspacing="0"><tr><td>

<asp:Repeater id="dlComment" runat="server">
<ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onMouseOver="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><hr><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><table align="center" cellpadding="0" cellspacing="0"><tr class="iTitle"><td class="iUp"><div class="sysBr500" id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("author")%> 在 <%#Container.DataItem("add_date")%> 说:</div></td></tr></table><table class="item" align="center" cellpadding="0" cellspacing="0"><tr><td align="center"><table class="itemBody" align="center" border="0" cellpadding="0" cellspacing="0"><tr><td class="iText"><div class="sysBr476"><%# Container.DataItem("content") %></div></td></tr></table></td></tr></table></div></ItemTemplate>
</asp:Repeater>




<table class="item" align="center" cellpadding="0" cellspacing="0"><tr><td class="iBottom"></td></tr></table><table width="100%"><tr height="12"><td></td></tr></table>

	<div class="dashed"></div><br><table align="center" width="100%" cellpadding="0" cellspacing="0"><tr class="line" height="1"><td></td></tr><tr><td class="postBox">
<form action="/savecomment.aspx" method=post name="PostComment" id="PostComment">
<input type=hidden name="action" value="add">
<input type=hidden name="article_id" value="<%=id%>">
<input id="TitleText" name="TitleText" type="hidden" value="回复：<%=title%>">
<fieldset class="postFSet"><legend><%=ph.g("请发表评论")%></legend><table align="center" cellpadding="0" cellspacing="0"><tr><td width="48"></td><td></td></tr>
<tr height="6"><td colspan="2"></td></tr><tr><td>昵　称：</td><td><input id="AuthorText" style="width:360px;" name="AuthorText" type="text" size="76" class="colorCCC" onfocus="this.className='color999';" onblur="this.className='colorCCC';" maxlength=50 value="<%=viewer%>"></td></tr><tr><td>主　页：</td><td><input id="HomepageText" style="width:360px;" name="HomepageText" type="text" size="76" class="colorCCC" onfocus="this.className='color999';" onblur="this.className='colorCCC';" value="<%=hp%>" maxlength=255></td></tr><tr><td valign="top">内　容：</td><td><textarea id="ContentText" style="width:360px;" name="ContentText" cols="59" class="colorCCC" onfocus="this.className='color999';if(this.value=='　') this.value=''" onblur="this.className='colorCCC';">　</textarea></td></tr><tr><td class="post" colspan="2"><a href="javascript:document.forms.PostComment.submit();"><%=ph.g("请发表评论")%></a></td></tr></table></fieldset>
</form></td></tr></table></div></div></td></tr></table></center><br></div></td></tr></table></td></tr></tbody></table><script type="text/javascript">tagging();</script></body>
</html>