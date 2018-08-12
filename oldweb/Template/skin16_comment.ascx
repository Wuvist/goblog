<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<html>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<link rel="stylesheet" href="/Template/skin16.css" type="text/css"> 
</head>

<META content="MSHTML 6.00.2800.1106" name=GENERATOR></HEAD>
<BODY leftMargin=0 background=/images/skins/skin16/qback1.gif topMargin=0 
marginheight="0" marginwidth="0">
<TABLE cellSpacing=0 cellPadding=0 width=750 align=center border=0>
  <TBODY>
  <TR>
      <TD width=304 height="50"><IMG height=43 src="/images/skins/skin16/qleft4.gif" width=304 
      border=0></TD>
    <TD>
      <TABLE cellSpacing=0 cellPadding=0 align=right border=0>
        <TBODY>
        <TR>
          <TD width=26><IMG height=32 src="/images/skins/skin16/qleft5.gif" width=26 
            border=0></TD>
          <TD class=text1 align=middle width=250 
          background=/images/skins/skin16/qback5.gif><%=path%></TD>
          <TD width=25><IMG height=32 src="/images/skins/skin16/qright2.gif" width=25 
            border=0></TD></TR></TBODY></TABLE></TD>
    <TD width=1 bgColor=#615f64><BR></TD></TR></TBODY></TABLE>
<TABLE cellSpacing=0 cellPadding=0 width=750 align=center border=0>
  <TBODY>
  <TR>
    <TD width=1 bgColor=#615f64><BR></TD>
    <TD width=748 background=/images/skins/skin16/qback.gif height=37>
      <TABLE border=0 align="right" cellPadding=0 cellSpacing=0>
        <TBODY>
        <TR>
          <TD width=10><BR></TD>
          <TD><IMG height=22 src="/images/skins/skin16/butleft.gif" width=21 
          border=0></TD>
          <TD class=nav background=/images/skins/skin16/butback.gif><A href="/"><%=ph.g("博客风")%></A></TD>
          <TD><IMG height=22 src="/images/skins/skin16/butright.gif" width=8 
          border=0></TD>
          <TD width=1><BR></TD>
          <TD><IMG height=22 src="/images/skins/skin16/butleft.gif" width=21 
          border=0></TD>
          <TD class=nav background=/images/skins/skin16/butback.gif><a href="/<%=uid%>/"><%=path%>首页</a></TD>
          <TD><IMG height=22 src="/images/skins/skin16/butright.gif" width=8 
          border=0></TD>
          <TD width=1><BR></TD>
          <TD><IMG height=22 src="/images/skins/skin16/butleft.gif" width=21 
          border=0></TD>
          <TD class=nav background=/images/skins/skin16/butback.gif><a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a></TD>
          <TD><IMG height=22 src="/images/skins/skin16/butright.gif" width=8 
          border=0></TD>
        <TD width=1><BR></TD></TR></TBODY></TABLE></TD>
    <TD width=1 bgColor=#615f64><BR></TD></TR></TBODY></TABLE>
<TABLE cellSpacing=0 cellPadding=0 width=750 align=center border=0>
  <TBODY>
  <TR>
    <TD><IMG height=10 src="/images/skins/skin16/qleft1.gif" width=273 border=0></TD>
    <TD width=1 bgColor=#615f64><SPACER height="10" 
type="block"></TD></TR></TBODY></TABLE>
<TABLE cellSpacing=0 cellPadding=0 width=750 align=center border=0>
  <TBODY>
  <TR>
    <TD width=1 bgColor=#615f64></TD>
    <TD vAlign=top width=204><IMG height=337 src="/images/skins/skin16/qleft3.jpg" 
      width=204 border=0><div class="menu">
	  <p class=head>*<%=ph.g("网志介绍")%>*</p>
	  <Acme:userpic runat="server" ID="userpic1" deco="<img src='" enddeco="'>"/>
      <p><%=uinfo%></p><br/><br/><br/><br/>
      <p class=head>*<%=ph.g("个人链接")%>*</p><asp:Repeater id="links" runat="server">
						<ItemTemplate>
							<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
						</ItemTemplate>
					</asp:Repeater></p>
      <br/><br/><br/><br/><p class=head>*<%=ph.g("网志分类")%>*</p>
	  <asp:Repeater id="dlCate" runat="server">
						<ItemTemplate><a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %>(<%# Container.DataItem("files") %>)</a><br>
						</ItemTemplate>
					</asp:Repeater>
      <br/><br/><br/><br/><p class=head>*<%=ph.g("网志存档")%>*</p>
<asp:Repeater id="archive" runat="server">
		<ItemTemplate>
			<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
		</ItemTemplate>
	</asp:Repeater>
	<br/><br/><br/><br/><p class=head>*<%=ph.g("累计浏览")%>：<%=count%>*<br>
	<Acme:Message runat="server" ID="Message1"/></p></div>
	  </TD>
    <TD vAlign=top align=left width=544>
      <TABLE height=300 cellSpacing=0 cellPadding=0 width=520 border=0>
        <TBODY>
        <TR>
          <TD width=1 background=/images/skins/skin16/qback3.gif rowSpan=5><SPACER 
            height="1" type="block"></TD>
          <TD width=508 background=/images/skins/skin16/qline1.jpg><SPACER height="1" 
            type="block"></TD>
          <TD width=1 background=/images/skins/skin16/qback3.gif rowSpan=5><SPACER 
            height="1" type="block"></TD></TR>
        <TR>
          <TD bgColor=#dbdbdb height=13><BR></TD></TR>
        <TR>
          <TD vAlign=top bgColor=#ffffff><BR>
            <TABLE cellSpacing=0 cellPadding=0 align=center border=0>
              <TBODY>
              <TR>
                <TD width=35><IMG height=24 src="/images/skins/skin16/qcorner.gif" 
                  width=35 border=0></TD>
                <TD width=290 
background=/images/skins/skin16/qline2.gif><span class="head"><%=title%></span></TD>
                <TD width=140 
            background=/images/skins/skin16/qline3.gif></TD></TR></TBODY></TABLE><BR>
            <TABLE cellSpacing=0 cellPadding=0 width="90%" align=center 
border=0>
              <TBODY>
              <TR>
                <TD class=text><%=content%><br><br><div id="tags"><%=tags%></div><div align=right> <%=nick%> @ <%=adddate%> <br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a> </div><br/><br/>
					<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onMouseOver="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><hr>
		<b><%# Container.DataItem("title") %></b><br>
		<%# Container.DataItem("content") %><br>
		<div align=right id="_r<%# Container.DataItem("cid") %>">By:<%#Container.DataItem("author")%> @ <%#Container.DataItem("add_date")%></div>
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
    　　　<INPUT type="submit" value="提交" class="logbutton"></form>
                </TD></TR></TBODY></TABLE></TD></TR>
        <TR>
          <TD bgColor=#dbdbdb height=13><BR></TD></TR>
        <TR>
          <TD width=508 background=/images/skins/skin16/qline1.jpg><SPACER height="1" 
            type="block"></TD></TR></TBODY></TABLE><B><BR></B></TD>
    <TD width=1 bgColor=#615f64></TD></TR></TBODY></TABLE>
<TABLE cellSpacing=0 cellPadding=0 width=750 align=center border=0>
  <TBODY>
  <TR>
    <TD width=1 bgColor=#615f64 height=88><BR></TD>
    <TD align=middle width=748 background=/images/skins/skin16/qback2.gif>......................................Copyright (C) 2004, syl. all right reserved...........................................</TD>
    <TD width=1 bgColor=#615f64><BR></TD></TR></TBODY></TABLE></BODY><script type="text/javascript">tagging();</script>
</html>