<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3c.org/TR/1999/REC-html401-19991224/loose.dtd"><html>
<head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link><TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<link rel=stylesheet href="/Template/skin6.css" type=text/css> 
</head>
<BODY leftMargin=0 topMargin=0 marginheight="0" marginwidth="0">

<TABLE cellSpacing=0 cellPadding=0 align=center border=0>
  <TBODY>
  <TR>
    <TD vAlign=bottom align=middle>
      <TABLE cellSpacing=0 cellPadding=0 border=0>
        <TBODY>
        <TR vAlign=bottom align=middle>
          <TD width=98 height=60><b><%=path%></b></TD>
			</TR></TBODY></TABLE></TR></TBODY></TABLE>

<DIV id=box>
<DIV class=blog>
<TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
  <TBODY>
  <TR height=23>
    <TD vAlign=top align=middle height=23>
      
	  <TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR height=23>

          <TD width=81 height=23><IMG title="" height=23 alt="" 
            src="/images/skins/skin6/1.gif" width=81 align=top border=0></TD>
          
		  <TD vAlign=center align=middle background="/images/skins/skin6/2.gif" 
          height=23><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0><SPAN 
            class=title><%=title%></SPAN></TD>
          
		  <TD align=right width=48 height=23><IMG title="" height=23 alt="" 
            src="/images/skins/skin6/3.gif" width=48 align=top 
        border=0></TD></TR></TBODY></TABLE></TD></TR>
  <TR>
    <TD vAlign=center align=left>
      <TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD vAlign=top align=left width=48 
          background="/images/skins/skin6/4.gif"><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0></TD>
          <TD vAlign=top align=left >
		  <br>
		  <%=content%>
<div id="tags"><%=tags%></div>		   
		 <div class="com"><%=nick%> @<%=adddate%><br><a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div>

		 
		 </div>
		   </TD>
         
		  <TD vAlign=top align=right width=48 
          background="/images/skins/skin6/5.gif"><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 
      border=0></TD></TR></TBODY></TABLE></TD></TR>
  <TR>
    <TD vAlign=bottom align=right>
      <TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR height=41>
          <TD vAlign=top align=left width=81 height=41><IMG title="" height=41 
            alt="" src="/images/skins/skin6/6.gif" width=81 border=0></TD>
          <TD vAlign=center align=middle background="/images/skins/skin6/7.gif" 
          height=41><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0></TD>
          <TD vAlign=top align=right width=48 height=41><IMG title="" 
            height=41 alt="" src="/images/skins/skin6/8.gif" width=48 
        border=0></TD></TR></TBODY></TABLE></TD></TR></TBODY></TABLE>

	
 	<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a>

<TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR height=23>

          <TD width=81 height=23><IMG title="" height=23 alt="" 
            src="/images/skins/skin6/1.gif" width=81 align=top border=0></TD>
          
		  <TD vAlign=center align=middle background="/images/skins/skin6/2.gif" 
          height=23><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0></TD>
          
		  <TD align=right width=48 height=23><IMG title="" height=23 alt="" 
            src="/images/skins/skin6/3.gif" width=48 align=top 
        border=0></TD></TR></TBODY></TABLE></TD></TR>
  <TR>
    <TD vAlign=center align=left>
      <TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD vAlign=top align=left width=48 
          background="/images/skins/skin6/4.gif"><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0></TD>
          <TD vAlign=top align=left >
	   <div class="post">
	<br><div id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("author")%> 在 <%#Container.DataItem("add_date")%> 说:</div><br />
	<%# Container.DataItem("content") %>
	</div>
		</div>
		   </TD>
         
		  <TD vAlign=top align=right width=48 
          background="/images/skins/skin6/5.gif"><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 
      border=0></TD></TR></TBODY></TABLE></TD></TR>
  <TR>
    <TD vAlign=bottom align=right>
      <TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR height=41>
          <TD vAlign=top align=left width=81 height=41><IMG title="" height=41 
            alt="" src="/images/skins/skin6/6.gif" width=81 border=0></TD>
          <TD vAlign=center align=middle background="/images/skins/skin6/7.gif" 
          height=41><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0></TD>
          <TD vAlign=top align=right width=48 height=41><IMG title="" 
            height=41 alt="" src="/images/skins/skin6/8.gif" width=48 
        border=0></TD></TR></TBODY></TABLE></TD></TR></TBODY></TABLE></div>
</ItemTemplate>
</asp:Repeater>
<TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR height=23>

          <TD width=81 height=23><IMG title="" height=23 alt="" 
            src="/images/skins/skin6/1.gif" width=81 align=top border=0></TD>
          
		  <TD vAlign=center align=middle background="/images/skins/skin6/2.gif" 
          height=23><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0><SPAN 
            class=title>请<%=ph.g("请发表评论")%>：</SPAN></TD>
          
		  <TD align=right width=48 height=23><IMG title="" height=23 alt="" 
            src="/images/skins/skin6/3.gif" width=48 align=top 
        border=0></TD></TR></TBODY></TABLE></TD></TR>
  <TR>
    <TD vAlign=center align=left>
      <TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD vAlign=top align=left width=48 
          background="/images/skins/skin6/4.gif"><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0></TD>
          <TD vAlign=top align=left >
		  <br>
		 
		 <p><form action="/savecomment.aspx" method=post>
		<input type=hidden name="action" value="add">
		<input type=hidden name="article_id" value="<%=id%>">
	<INPUT type="hidden" name="TitleText" value="回复：<%=title%>">
	<font size="2">名字：</font><INPUT type="text" name="AuthorText"  size=40 style="border: 1px solid #AAAAAA;"   maxlength=50 value="<%=viewer%>"><br>
	<font size="2">主页：</font><INPUT type="text" name="HomepageText"  size=40 style="border: 1px solid     #AAAAAA;" value="<%=hp%>" maxlength=255><br>
	<font size="2">内容：</font><br/><TEXTAREA name="ContentText" rows="5" cols="50"></TEXTAREA></p>
    <INPUT type="submit" value="提交" style="border: 1px solid #AAAAAA;background-color: #CCD4E0;"></form><br/><br/>
		 
		 </div>
		   </TD>
         
		  <TD vAlign=top align=right width=48 
          background="/images/skins/skin6/5.gif"><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 
      border=0></TD></TR></TBODY></TABLE></TD></TR>
  <TR>
    <TD vAlign=bottom align=right>
      <TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR height=41>
          <TD vAlign=top align=left width=81 height=41><IMG title="" height=41 
            alt="" src="/images/skins/skin6/6.gif" width=81 border=0></TD>
          <TD vAlign=center align=middle background="/images/skins/skin6/7.gif" 
          height=41><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0></TD>
          <TD vAlign=top align=right width=48 height=41><IMG title="" 
            height=41 alt="" src="/images/skins/skin6/8.gif" width=48 
        border=0></TD></TR></TBODY></TABLE></TD></TR></TBODY></TABLE>

<CENTER>Layout inspired from <A href="http://www.apple.com/" target=_blank>MacOS 
X</A>.</CENTER></DIV></DIV>
<DIV id=macmenu>
<TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
  <TBODY>
  <TR height=23>
    <TD vAlign=top align=middle height=23>
      <TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR height=23>
          <TD width=81 height=23><IMG alt="" 
            src="/images/skins/skin6/macmenu.gif" align=top 
      border=0 width="230" height="30"></TD></TR></TBODY></TABLE></TD></TR>
  <TR>
    <TD vAlign=center align=left>
      <TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD vAlign=top align=left width=48 
          background="/images/skins/skin6/4.gif"><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0></TD>
          <TD vAlign=top align=left>
            <DIV class=macside>
                    <DIV class=blocktitle> <%=ph.g("导航")%></DIV>
                    <BR><BR>
                    <IMG src="/images/skins/skin6/li.gif" width="7" height="8">&nbsp;<a href="/"><%=ph.g("博客风")%></a> <br>
					<IMG src="/images/skins/skin6/li.gif" width="7" height="8">&nbsp;<a href="/<%=uid%>/"><%=path%>首页</a> <br>
					<IMG src="/images/skins/skin6/li.gif" width="7" height="8">&nbsp;<a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a><BR>
					</div> 
					<BR>
					<BR>
					
<br>
					 </DIV>
            <DIV class=macside>
            <DIV class=blocktitle><%=ph.g("网志介绍")%></DIV>
			<Acme:userpic runat="server" ID="userpic1"/>
			<%=uinfo%>
			<BR></DIV>

                  <DIV class=macside> 
                    <DIV class=blocktitle><%=ph.g("网志分类")%></DIV>
					<asp:Repeater id="dlCate" runat="server">
					<ItemTemplate>
					<IMG src="/images/skins/skin6/li.gif" width="7" height="8">&nbsp;<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
					</ItemTemplate>
					</asp:Repeater>
                     
            <BR></DIV>
                  <DIV class=macside> 
                    <DIV class=blocktitle><%=ph.g("网志存档")%></DIV>
<asp:Repeater id="archive" runat="server">
	<ItemTemplate>
		<IMG src="/images/skins/skin6/li.gif" width="7" height="8">&nbsp;<a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
	</ItemTemplate>
</asp:Repeater>            <BR></DIV>

            <DIV class=macside>
            <DIV class=blocktitle><%=ph.g("个人链接")%></DIV>
			<asp:Repeater id="links" runat="server">
				<ItemTemplate>
				<IMG 
            src="/images/skins/skin6/bullet_double.gif" width="8" height="9">&nbsp;<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
				</ItemTemplate>
				</asp:Repeater>
				<p class="nav"><%=ph.g("累计浏览")%>: <%=count%></p>
			<BR><Acme:Message runat="server" ID="Message1"/></DIV>
            </TD>
          <TD vAlign=top align=right width=48 
          background="/images/skins/skin6/5.gif"><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 
      border=0></TD></TR></TBODY></TABLE></TD></TR>
  <TR>
    <TD vAlign=bottom align=right>
      <TABLE height="100%" cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR height=41>
          <TD vAlign=top align=left width=81 height=41><IMG title="" height=41 
            alt="" src="/images/skins/skin6/6.gif" width=81 border=0></TD>
          <TD vAlign=center align=middle background="/images/skins/skin6/7.gif" 
          height=41><IMG title="" height=1 alt="" 
            src="/images/skins/skin6/pixel.gif" width=1 border=0></TD>
          <TD vAlign=top align=right width=48 height=41><IMG title="" 
            height=41 alt="" src="/images/skins/skin6/8.gif" width=48 
        border=0></TD></TR></TBODY></TABLE></TD></TR></TBODY></TABLE><BR>
</BODY><script type="text/javascript">tagging();</script>
</html>
