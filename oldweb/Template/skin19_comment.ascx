<%@ Control Language="vb" AutoEventWireup="false" CodeFile="tcomment.ascx.vb" Inherits="Blogwind.tcomment" TargetSchema="http://schemas.microsoft.com/intellisense/ie5" %>
<%@ Register TagPrefix="Acme" TagName="Message" Src="../log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="../icon.ascx" %>

<html>
<head>
<link id="RSSLink" title="RSS" type="application/rss+xml" rel="alternate" href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"></link><TITLE><%= subject & " - " & catename & " - " & blogname%></TITLE>
<script type="text/javascript" src="/icon.js"></script>
<script type="text/javascript" src="/util.js"></script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<LINK href="/Template/skin19.css" type=text/css rel=stylesheet>
<META content="MSHTML 6.00.3790.1830" name=GENERATOR></HEAD>
<BODY bgColor=#993300 leftMargin=0 background=/images/skins/skin19/bg.gif topMargin=0 
marginheight="0" marginwidth="0">
<TABLE cellSpacing=0 cellPadding=0 width=740 align=center border=0>
  <TBODY>
  <TR>
    <TD vAlign=top width=202 background=/images/skins/skin19/default_15.gif height=411>
      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD vAlign=top width=202 background=/images/skins/skin19/default_01.gif 
          height=387>
            <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
              <TBODY>
              <TR>
                <TD class=unnamed1 height=250>&nbsp;</TD></TR>
              <TR>
                <TD>
                  <TABLE cellSpacing=0 cellPadding=0 width="65%" align=center 
                  border=0>
                    <TBODY>
                    <TR>
                      <TD class=unnamed1> </TD></TR></TBODY></TABLE></TD></TR></TBODY></TABLE></TD></TR>
        <TR>
          <TD vAlign=top>&nbsp; </TD></TR></TBODY></TABLE>

      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD width=202 background=/images/skins/skin19/default_14.gif height=22>
            <TABLE cellSpacing=0 cellPadding=0 width="70%" align=center 
border=0>
              <TBODY>
              <TR>
                <TD 
        class=unnamed1><STRONG><%=ph.g("网志介绍")%></STRONG></TD></TR></TBODY></TABLE></TD></TR>
        <TR>
          <TD vAlign=top height=22>
            <TABLE width="70%" align=center>
              <TBODY>
              <TR>
                <TD class=unnamed2><Acme:userpic runat="server" ID="userpic1" deco="<div align='center'><img src='" enddeco="'></div>"/><%=uinfo%></TD></TR>
              <TR>
                <TD height=5></TD></TR></TBODY></TABLE></TD></TR></TBODY></TABLE>
      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD width=202 background=/images/skins/skin19/default_14.gif height=22>
            <TABLE cellSpacing=0 cellPadding=0 width="70%" align=center 
border=0>
              <TBODY>
              <TR>
                <TD 
          class=unnamed1><STRONG><%=ph.g("网志分类")%></STRONG></TD></TR></TBODY></TABLE></TD></TR>
        <TR>
          <TD vAlign=top height=22>
            <TABLE width="89%" align=center>
              <TBODY>
              <TR>
                <TD class=unnamed1 width="7%">&nbsp;</TD>
                <TD class=unnamed2 width="93%">				
				<asp:Repeater id="dlCate" runat="server">
					<ItemTemplate>
					<a href = "/<%# uid & "/cate" & Container.DataItem("index") %>.shtml" ><%# Container.DataItem("cate") %></a>(<%# Container.DataItem("files") %>)<br>
					</ItemTemplate>
				</asp:Repeater>
 </TD></TR></TBODY></TABLE></TD></TR></TBODY></TABLE>

      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD width=202 background=/images/skins/skin19/default_14.gif height=22>
            <TABLE cellSpacing=0 cellPadding=0 width="70%" align=center 
border=0>
              <TBODY>
              <TR>
                <TD 
        class=unnamed1><STRONG><%=ph.g("网志存档")%></STRONG></TD></TR></TBODY></TABLE></TD></TR>
        <TR>
          <TD vAlign=top height=22>
            <TABLE width="89%" align=center>
              <TBODY>
              <TR>
                <TD class=unnamed1 width="7%">&nbsp;</TD>
                <TD class=unnamed2 width="93%">
								<asp:Repeater id="archive" runat="server">
					<ItemTemplate><a href = "/<%# uid & "/" & Container.DataItem("year") %>/<%# Container.DataItem("month") %>/" ><%# Container.DataItem("year") %>年<%# Container.DataItem("month") %>月</a><br>
					</ItemTemplate>
				</asp:Repeater>

	</TD></TR></TBODY></TABLE></TD></TR></TBODY></TABLE>
      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD width=202 background=/images/skins/skin19/default_14.gif height=22>
            <TABLE cellSpacing=0 cellPadding=0 width="70%" align=center 
border=0>
              <TBODY>
              <TR>
                <TD 
        class=unnamed1><STRONG><%=ph.g("个人链接")%></STRONG></TD></TR></TBODY></TABLE></TD></TR>
        <TR>
          <TD vAlign=top height=22>
            <TABLE width="89%" align=center>
              <TBODY>
              <TR>
                <TD class=unnamed1 width="7%">&nbsp;</TD>
                <TD class=unnamed2 width="93%">

  				<asp:Repeater id="links" runat="server">
					<ItemTemplate>
					<a href = "<%# Container.DataItem("URL") %>" target="_blank"><%# Container.DataItem("link") %></a><br>
					</ItemTemplate>
				</asp:Repeater>

			  </TD></TR></TBODY></TABLE>
			  
			 </TD></TR></TBODY></TABLE>
      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD width=202 background=/images/skins/skin19/default_14.gif height=22>
            <TABLE cellSpacing=0 cellPadding=0 width="70%" align=center 
border=0>
              <TBODY>
              <TR>
                <TD 
        class=unnamed1><STRONG><%=ph.g("累计浏览")%>：<%=count%></STRONG></TD></TR></TBODY></TABLE></TD></TR>
        <TR>
          <TD vAlign=top height=22>
            <TABLE width="89%" align=center>
              <TBODY>
              <TR>
                <TD class=unnamed1 width="7%">&nbsp;</TD>
                <TD class=unnamed2 width="93%">


			  </TD></TR></TBODY></TABLE>
			  
			 </TD></TR></TBODY></TABLE>
			       <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD vAlign=top height=22>
            <TABLE width="89%" align=center>
              <TBODY>
              <TR>
                <TD class=unnamed1 width="7%">&nbsp;</TD>
                <TD class=unnamed2 width="93%">
<Acme:Message runat="server" ID="Message1"/>

			  </TD></TR></TBODY></TABLE>
			  
			 </TD></TR></TBODY></TABLE>
			 			 
			 </TD>
    <TD vAlign=top width=538 background=/images/skins/skin19/default_12.gif>
      <TABLE cellSpacing=0 cellPadding=0 width="100%" border=0>
        <TBODY>
        <TR>
          <TD><IMG height=42 alt="" src="/images/skins/skin19/default_02.gif" 
          width=538></TD></TR>
        <TR>
          <TD>
            <TABLE cellSpacing=0 cellPadding=0 width="100%" 
            background=/images/skins/skin19/default_04.gif border=0>
              <TBODY>
              <TR>
                <TD width="9%"><IMG height=47 alt="" 
                  src="/images/skins/skin19/default_03.gif" width=46></TD>
                <TD class=unnamed4 vAlign=bottom>
                  <TABLE cellSpacing=0 cellPadding=0 width="75%" border=0>
                    <TBODY>
                    <TR>
                      <TD class=unnamed4>
                        <DIV align=center><%=path%></DIV></TD></TR></TBODY></TABLE></TD>
                <TD width="15%">
                  <DIV align=right><IMG height=47 alt="" 
                  src="/images/skins/skin19/default_06.gif" 
            width=61></DIV></TD></TR></TBODY></TABLE></TD></TR>
        <TR>
          <TD><IMG height=33 alt="" src="/images/skins/skin19/default_07.gif" 
          width=538></TD></TR>
        <TR>
          <TD class=unnamed5 vAlign=bottom background=/images/skins/skin19/default_08.gif 
          height=22>
            <DIV align=center>
            <TABLE cellSpacing=0 cellPadding=0 width="84%" align=center 
border=0>
              <TBODY>
              <TR>
                <TD class=unnamed5>
				
				<center>
				<table><tr><TD><a href="http://www.blogwind.com"><%=ph.g("博客风")%></a></TD><td> | </td>
                <TD><a href="/<%=uid%>/"><%=path%>首页</a></TD><td> | </td>
                <TD><a href="/email.aspx?user=<%=uid%>"><%=ph.g("联系")%></a></TD></tr></table></center>
			</TD></TR></TBODY></TABLE></DIV></TD></TR>
        <TR>
          <TD><IMG height=16 alt="" src="/images/skins/skin19/default_09.gif" 
          width=538></TD></TR>
        <TR>
          <TD vAlign=top><IMG height=15 alt="" 
            src="/images/skins/skin19/default_11.gif" width=538></TD></TR></TBODY></TABLE>

	  <TABLE class=textbox style="MARGIN-RIGHT: 10px" cellSpacing=0 
      cellPadding=4 width="95%" align=center border=0>
        <TBODY>
        <TR>
          <TD class=unnamed3><IMG height=14 src="/images/skins/skin19/a.gif" width=15 
            align=absMiddle><%=title%></a> </TD>
        <TR>
          <TD class=unnamed7 style="PADDING-LEFT: 8px" colSpan=2>				
			<%=content%>
			</TD></TR>
        <TR>
          <TD class=unnamed7 style="PADDING-LEFT: 8px" colSpan=2>				
		  <div id="tags"><%=tags%></div>
			<div align="right"><%=nick%>发表于<%=adddate%> | <a href='/<%=uid%>/cate<%=cateid%>.shtml'><%=ph.g("查看本分类的所有网志")%>:<%= catename%></a></div>
	  <BR><BR></TD></TR></TBODY></TABLE>

	  <BR>
	  <TABLE class=textbox style="MARGIN-RIGHT: 10px" cellSpacing=0 
      cellPadding=4 width="95%" align=center border=0>
        <TBODY>
        <TR>
          <TD class=unnamed3><IMG height=14 src="/images/skins/skin19/a.gif" width=15 
            align=absMiddle><%=title%></a> </TD>
        <TR>
          <TD class=unnamed7 style="PADDING-LEFT: 8px" colSpan=2>				
				<asp:Repeater id="dlComment" runat="server">
	   <ItemTemplate><div id="_c<%# Container.DataItem("cid") %>" onMouseOver="s(<%# Container.DataItem("cid") %>)" onmouseout="h(<%# Container.DataItem("cid") %>)"><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a><a href="#" name="blogc<%# Container.DataItem("cid") %>"></a>
			<div class="comments">
				<p><%# Container.DataItem("title") %></p>
				<p><%# Container.DataItem("content") %></p>
				<div style="text-align:right" id="_r<%# Container.DataItem("cid") %>"><%#Container.DataItem("author")%> @ <%#Container.DataItem("add_date")%></div>
				<hr>
			</div></div>
		</ItemTemplate>
	</asp:Repeater>
			</TD></TR>
</TBODY></TABLE>


		  <TABLE class=textbox style="MARGIN-RIGHT: 10px" cellSpacing=0 
      cellPadding=4 width="95%" align=center border=0>
        <TBODY>
        <TR>
          <TD class=unnamed7>
				<form action="/savecomment.aspx" method=post>
				<input type=hidden name="action" value="add">
				<input type=hidden name="article_id" value="<%=id%>">
				<table width="100%">
				<tr><td colspan="2" class=unnamed7><%=ph.g("请发表评论")%>:</td></tr>
				<tr><td class=unnamed7>标题:</td><td><input type="text" name="TitleText" size=50 value="回复：<%=title%>" maxlength=50></td></tr>
				<tr><td class=unnamed7>名字:</td><td><input type="text" name="AuthorText"  size=50 maxlength=50 value="<%=viewer%>"></td></tr>
				<tr><td class=unnamed7>主页:</td><td><input type="text" name="HomepageText" size=50  value="<%=hp%>" maxlength=255></td></tr>
				<tr><td valign="top" class=unnamed7>内容:</td><td><textarea class=button cols="50" rows="5" name="ContentText"></textarea></td></tr>
				<tr><td></td><td><input class=msgsubmit type="submit" value="发表"></td></tr>
				</table>
				</form>
	  
	  </TD></TR></TBODY></TABLE>
</TD></TR></TBODY></TABLE>
<TABLE cellSpacing=0 cellPadding=0 width=740 align=center border=0>
  <TBODY>
  <TR>
    <TD><IMG height=98 alt="" src="/images/skins/skin19/default_17.gif" width=202><IMG 
      height=98 alt="" src="/images/skins/skin19/default_18.gif" 
width=538></TD></TR></TBODY></TABLE>
</BODY><script type="text/javascript">tagging();</script>
</html>

