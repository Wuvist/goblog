<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>
<HTML>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link>
<TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<LINK href="/Template/skin18.css" type=text/css rel=stylesheet>
<BODY leftMargin=0 background=/images/skins/skin18/bg.gif topMargin=0>
<DIV style="BACKGROUND: url(/images/skins/skin18/bg.gif)" align=center>
<TABLE cellSpacing=0 cellPadding=0 width=776 border=0>
  <TBODY>
  <TR>
    <TD><IMG height=20 alt="" src="/images/skins/skin18/top_r1_c1.jpg" width=105 
      border=0 name=top_r1_c1></TD>
    <TD><IMG height=20 alt="" src="/images/skins/skin18/top_r1_c2.jpg" width=177 
      border=0 name=top_r1_c2></TD>
    <TD><IMG height=20 alt="" src="/images/skins/skin18/top_r1_c3.jpg" width=200 
      border=0 name=top_r1_c3></TD>
    <TD colSpan=2><IMG height=20 alt="" src="/images/skins/skin18/top_r1_c4.jpg" 
      width=294 border=0 name=top_r1_c4></TD></TR>
  <TR>
    <TD><IMG height=123 alt="" src="/images/skins/skin18/top_r2_c1.jpg" width=105 
      border=0 name=top_r2_c1></TD>
    <TD><IMG height=123 alt="" src="/images/skins/skin18/top_r2_c2.jpg" width=177 
      border=0 name=top_r2_c2></TD>
    <TD><IMG height=123 alt="" src="/images/skins/skin18/top_r2_c3.jpg" width=200 
      border=0 name=top_r2_c3></TD>
    <TD align=middle width=253 background=/images/skins/skin18/top_r2_c4.jpg 
      height=123><SPAN style="OVERFLOW: hidden; WIDTH: 253px; HEIGHT: 123px">
      <TABLE height=0 cellSpacing=0 cellPadding=8 width="100%" border=0>
        <TBODY>
        <TR>
          <TD align=right><SPAN class=page-header><a href="/<%=uid%>/"><%=path%></A></SPAN> <BR><SPAN 
            class=page-subheader height="100%" width="100%"><%=uinfo%></SPAN> 
          </TD>
          <TD width=1></TD></TR></TBODY></TABLE></SPAN></TD>
    <TD><IMG height=123 alt="" src="/images/skins/skin18/top_r2_c5.jpg" width=41 
      border=0 name=top_r2_c5></TD></TR>
  <TR>
    <TD><IMG height=14 alt="" src="/images/skins/skin18/top_r3_c1.jpg" width=105 
      border=0 name=top_r3_c1></TD>
    <TD><IMG height=14 alt="" src="/images/skins/skin18/top_r3_c2.jpg" width=177 
      border=0 name=top_r3_c2></TD>
    <TD><IMG height=14 alt="" src="/images/skins/skin18/top_r3_c3.jpg" width=200 
      border=0 name=top_r3_c3></TD>
    <TD><IMG height=14 alt="" src="/images/skins/skin18/top_r3_c4.jpg" width=253 
      border=0 name=top_r3_c4></TD>
    <TD><IMG height=14 alt="" src="/images/skins/skin18/top_r3_c5.jpg" width=41 
      border=0 name=top_r3_c5></TD></TR></TBODY></TABLE>
<TABLE cellSpacing=0 cellPadding=0 width=776 background=/images/skins/skin18/bg_1.gif 
border=0>
  <TBODY>
  <TR>
    <TD width=34>&nbsp;</TD>
    <TD vAlign=top align=middle width=532>
      <TABLE cellSpacing=0 cellPadding=0 width=532 border=0>
        <TBODY>
        <TR>
          <TD><IMG height=60 alt="" src="/images/skins/skin18/nav_r1_c1.jpg" width=49 
            border=0 name=nav_r1_c1></TD>
          <TD><IMG height=60 alt="" src="/images/skins/skin18/nav_r1_c2.jpg" width=53 
            border=0 name=nav_r1_c2></TD>
          <TD><IMG height=60 alt="" src="/images/skins/skin18/nav_r1_c3.jpg" width=75 
            border=0 name=nav_r1_c3></TD>
          <TD><IMG height=60 alt="" src="/images/skins/skin18/nav_r1_c4.jpg" width=56 
            border=0 name=nav_r1_c4></TD>
          <TD><IMG height=60 alt="" src="/images/skins/skin18/nav_r1_c5.jpg" width=55 
            border=0 name=nav_r1_c5></TD>
          <TD><IMG height=60 alt="" src="/images/skins/skin18/nav_r1_c6.jpg" width=66 
            border=0 name=nav_r1_c6></TD>
          <TD><IMG height=60 alt="" src="/images/skins/skin18/nav_r1_c7.jpg" width=78 
            border=0 name=nav_r1_c7></TD>
          <TD><IMG height=60 alt="" src="/images/skins/skin18/nav_r1_c8.jpg" width=54 
            border=0 name=nav_r1_c8></TD>
          <TD><IMG height=60 alt="" src="/images/skins/skin18/nav_r1_c9.jpg" width=46 
            border=0 name=nav_r1_c9></TD></TR>
        <TR>
          <TD><IMG height=43 
            alt="" src="/images/skins/skin18/nav_r2_c1.jpg" width=49 border=0 
            name=nav_r2_c1></TD>
          <TD><IMG height=43 alt="" 
            src="/images/skins/skin18/nav_r2_c2.jpg" width=53 border=0 
            name=nav_r2_c2></TD>
          <TD><IMG height=43 
            alt="" src="/images/skins/skin18/nav_r2_c3.jpg" width=75 border=0 
            name=nav_r2_c3></TD>
          <TD><IMG height=43 alt="" 
            src="/images/skins/skin18/nav_r2_c4.jpg" width=56 border=0 
            name=nav_r2_c4></TD>
          <TD><IMG height=43 
            alt="" src="/images/skins/skin18/nav_r2_c5.jpg" width=55 border=0 
            name=nav_r2_c5></TD>
          <TD><IMG height=43 alt="" src="/images/skins/skin18/nav_r2_c6.jpg" width=66 
            border=0 name=nav_r2_c6></TD>
          <TD><IMG height=43 alt="" src="/images/skins/skin18/nav_r2_c7.jpg" width=78 
            border=0 name=nav_r2_c7></TD>
          <TD><IMG height=43 alt="" src="/images/skins/skin18/nav_r2_c8.jpg" width=54 
            border=0 name=nav_r2_c8></TD>
          <TD><IMG height=43 alt="" src="/images/skins/skin18/nav_r2_c9.jpg" width=46 
            border=0 name=nav_r2_c9></TD></TR></TBODY></TABLE>
      <TABLE class=marron1 height=28 cellSpacing=0 cellPadding=0 width="100%" 
      border=0>
        <TBODY>
        <TR align=middle>
          <TD>
            <TABLE cellSpacing=0 cellPadding=4 align=center 
border=0>
              <TBODY>
              <TR>
                <TD><a href="http://www.blogwind.com"><%=ph.g("博客风")%></a></TD><td> | </td>
                <TD><a href="/<%=uid%>/"><%=path%>首页</a></TD><td> | </td>
                <TD><a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a></TD>
</TR></TBODY></TABLE>
<BR></TD></TR>
        <TR>
          <TD><IMG height=7 src="/images/skins/skin18/line1.gif" 
        width=516></TD></TR></TBODY></TABLE>
      <DIV id=ArticleList1_Panel1>

      <TABLE class=marron1 cellSpacing=0 cellPadding=0 width=510 align=center 
      border=0>
        <TBODY>
        <TR>
		  <TD class=textbox-title background=/images/skins/skin18/titie_bg.gif 
          height=36><IMG height=3 src="/images/skins/skin18/pic_lit_001.gif" width=8 
            align=absMiddle> <B><%=title%></B></TD>
          <TD class=textbox-title style="BACKGROUND-POSITION: right 50%" 
          align=right background=/images/skins/skin18/titie_bg.gif>&nbsp; </TD></TR>
        <TR>
          <TD class=textbox-content width="95%" colSpan=2>

				<%=content%><br><br>
		</TD></TR>
        <TR>
          <TD class=textbox-bottom colSpan=2><IMG height=11 
            src="/images/skins/skin18/killtu.gif" width=17> 
			<div id="tags"><%=tags%></div>
			<%=nick%>发表于<%=adddate%> | <a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a>

<br><br>
        </TD></TR>
		<tr><td>
				<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onmouseover="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a>
			<div class="comments">
				<p><%# Container.DataItem("title") %></p>
				<p><%# Container.DataItem("content") %></p>
				<div style="text-align:right" id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("author")%> @ <%#Container.DataItem("add_date")%></div>
				<hr>
			</div></div>
		</ItemTemplate>
	</asp:Repeater>
			<div class="addcomment">
				<form action="/savecomment.aspx" method=post>
				<input type=hidden name="action" value="add">
				<input type=hidden name="article_id" value="<%=id%>">
				<table width="100%">
				<tr><td colspan="2"><%=ph.g("请发表评论")%>:</td></tr>
				<tr><td>标题:</td><td><input type="text" name="TitleText" size=50 value="回复：<%=title%>" maxlength=50></td></tr>
				<tr><td>名字:</td><td><input type="text" name="AuthorText"  size=50 maxlength=50 value="<%=viewer%>"></td></tr>
				<tr><td>主页:</td><td><input type="text" name="HomepageText" size=50  value="<%=hp%>" maxlength=255></td></tr>
				<tr><td valign="top">内容:</td><td><textarea cols="50" rows="5" name="ContentText"></textarea></td></tr>
				<tr><td></td><td><input type="submit" value="发表"></td></tr>
				</table>
				</form>
			</div>

		</td></tr>
		</TBODY></TABLE><BR>


      <TABLE height=5 cellSpacing=0 cellPadding=0 width=525 align=center 
      border=0>
        <TBODY>
        <TR>
          <TD><IMG height=7 src="/images/skins/skin18/line1.gif" 
        width=516></TD></TR></TBODY></TABLE><BR></DIV></TD>
    <TD vAlign=top align=middle width=180>
      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD background=/images/skins/skin18/dot.gif colSpan=2 height=5></TD></TR>
        <TR>
          <TD class=marron1 width="43%" height=20><IMG height=12 
            src="/images/skins/skin18/icon.gif" width=18><STRONG><%=ph.g("网志介绍")%></STRONG></TD>
          <TD class=marron1 width="57%"></TD></TR>
        <TR>
          <TD background=/images/skins/skin18/dot.gif colSpan=2 height=5></TD></TR>
        <TR>
          <TD colSpan=2>
            <TABLE class=marron1 cellSpacing=0 cellPadding=0 width="90%" 
            align=center border=0>
              <TBODY>
              <TR>
                <TD class=normalfont>
                  <TABLE class=panel cellSpacing=1 cellPadding=4 width="100%" 
                  align=center border=0>
                    <TBODY>
                    <TR>
                      <TD class=comment-text><Acme:userpic runat="server" ID="userpic1" deco="<div align='center'><img src='" enddeco="'></div>"/><%=uinfo%></TD></TR>
		</TBODY></TABLE><BR></TD></TR></TBODY></TABLE></TD></TR>
</TBODY></TABLE>
    
      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD background=/images/skins/skin18/dot.gif height=5></TD></TR>
        <TR>
          <TD class=marron1 height=20><IMG height=12 
            src="/images/skins/skin18/icon.gif" width=18><STRONG><%=ph.g("网志分类")%></STRONG></TD></TR>
        <TR>
          <TD background=/images/skins/skin18/dot.gif height=5></TD></TR>
        <TR align=middle>
          <TD>
            <TABLE class=marron1 cellSpacing=0 cellPadding=3 width="100%" 
            border=0>
              <TBODY>
              <TR>
                <TD>
				<asp:Repeater id="dlCate" runat="server">
					<ItemTemplate>
					<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
					</ItemTemplate>
				</asp:Repeater>

				</TD></TR>
</TBODY></TABLE><BR></TD></TR>
        <TR>
          <TD height=5></TD></TR></TBODY></TABLE>
      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD background=/images/skins/skin18/dot.gif colSpan=2 height=5></TD></TR>
        <TR>
          <TD class=marron1 width="43%" height=20><IMG height=12 
            src="/images/skins/skin18/icon.gif" width=18><STRONG><%=ph.g("网志存档")%></STRONG></TD>
          <TD class=marron1 width="57%"></TD></TR>
        <TR>
          <TD background=/images/skins/skin18/dot.gif colSpan=2 height=5></TD></TR>
        <TR align=middle>
          <TD colSpan=2>
            <TABLE class=marron1 cellSpacing=0 cellPadding=3 width="100%" 
            border=0>
              <TBODY>
              <TR>
                <TD>
				<asp:Repeater id="archive" runat="server">
					<ItemTemplate><a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
					</ItemTemplate>
				</asp:Repeater>

				</TD></TR></TBODY></TABLE><BR></TD></TR>
        <TR>
          <TD colSpan=2 height=5></TD></TR></TBODY></TABLE>
      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD background=/images/skins/skin18/dot.gif colSpan=2 height=5></TD></TR>
        <TR>
          <TD class=marron1 width="43%" height=20><IMG height=12 
            src="/images/skins/skin18/icon.gif" width=18><STRONG><%=ph.g("个人链接")%></STRONG></TD>
          <TD class=marron1 width="57%"></TD></TR>
        <TR>
          <TD background=/images/skins/skin18/dot.gif colSpan=2 height=5></TD></TR>
        <TR align=middle>
          <TD colSpan=2>
            <TABLE class=marron1 cellSpacing=0 cellPadding=3 width="100%" 
            border=0>
              <TBODY>
              <TR>
                <TD>
				<asp:Repeater id="links" runat="server">
					<ItemTemplate>
					<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
					</ItemTemplate>
				</asp:Repeater>

				</TD></TR></TBODY></TABLE></TD></TR>
        <TR>
          <TD colSpan=2 height=5></TD></TR></TBODY></TABLE>
		    <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD background=/images/skins/skin18/dot.gif colSpan=2 height=5></TD></TR>
        <TR>
          <TD class=marron1 height=20><IMG height=12 
            src="/images/skins/skin18/icon.gif" width=18><STRONG>浏览人次：<%=count%></STRONG></TD>
          </TR>
        <TR>
          <TD background=/images/skins/skin18/dot.gif colSpan=2 height=5></TD></TR>
        <TR align=middle>
          <TD colSpan=2>
            <TABLE class=marron1 cellSpacing=0 cellPadding=3 width="100%" 
            border=0>
              <TBODY>
              <TR>
                <TD></TD></TR></TBODY></TABLE><BR></TD></TR>
        <TR>
          <TD colSpan=2 height=5></TD></TR></TBODY></TABLE>
		    <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR align=middle>
          <TD colSpan=2>
            <TABLE class=marron1 cellSpacing=0 cellPadding=3 width="100%" 
            border=0>
              <TBODY>
              <TR>
                <TD><Acme:Message runat="server" ID="Message1"/></TD></TR></TBODY></TABLE><BR></TD></TR>
        <TR>
          <TD colSpan=2 height=5></TD></TR></TBODY></TABLE>
    
		  </TD>
    <TD width=30>&nbsp;</TD></TR></TBODY></TABLE>
<TABLE cellSpacing=0 cellPadding=0 width=776 border=0>
  <TBODY>
  <TR>
    <TD align=middle><IMG height=63 src="/images/skins/skin18/page_foot.gif" 
      width=758></TD></TR></TBODY></TABLE>
</DIV></BODY><script type="text/javascript">tagging();</script>
</html>
