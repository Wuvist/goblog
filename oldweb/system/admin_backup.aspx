<%@ Page Language="VB" AutoEventWireup="false" CodeFile="admin_backup.aspx.vb" Inherits="system_admin_backup" %>
<HTML>
	<HEAD>
		<meta http-equiv="Content-Type" content="text/html; charset=gb2312">
		<LINK href="/blogwind.ico" type="image/x-icon" rel="icon">
			<LINK href="blogwind.ico" type="image/x-icon" rel="shortcut icon">
				<meta content="博客，博客风，留学" name="keywords">
				<meta content="留学生博客网站" name="description">
				<link rel="stylesheet" href="admin.css" type="text/css">
				<style type="text/css">
<!--
/* DATE AND TIME */
p.datetime { line-height:20px; margin:0; padding:0; color:#666; font-size:11px; font-weight:bold; }
.datetime span { font-size:11px; color:#ccc; font-weight:normal; white-space:nowrap; }
.vDateField { margin-left:4px; }
table p.datetime { font-size:10px; margin-left:0; padding-left:0; }
.calendarbox, .clockbox { margin:5px auto; font-size:11px; width:16em; text-align:center; background:white; position:relative; }
.clockbox { width:9em; }
.calendar { margin:0; padding: 0; }
.calendar table { margin:0; padding:0; border-collapse:collapse; background:white; width:99%; }
.calendar caption, .calendarbox h2 { margin: 0; font-size:11px; text-align:center; border-top:none; }
.calendar th { font-size:10px; color:#666; padding:2px 3px; text-align:center; background:#e1e1e1 url(../img/admin/nav-bg.gif) 0 50% repeat-x; border-bottom:1px solid #ddd; }
.calendar td { font-size:11px; text-align: center; padding: 0; border-top:1px solid #eee; border-bottom:none; }
.calendar td.selected a { background: #C9DBED; }
.calendar td.nonday { background:#efefef; }
.calendar td.today a { background:#ffc; }
.calendar td a, .timelist a { display: block; font-weight:bold; padding:4px; text-decoration: none; color:#444; }
.calendar td a:hover, .timelist a:hover { background: #5b80b2; color:white; }
.calendar td a:active, .timelist a:active { background: #036; color:white; }
.calendarnav { font-size:10px; text-align: center; color:#ccc; margin:0; padding:1px 3px; }
.calendarnav a:link, #calendarnav a:visited, #calendarnav a:hover { color: #999; }
.calendar-shortcuts { background:white; font-size:10px; line-height:11px; border-top:1px solid #eee; padding:3px 0 4px; color:#ccc; }
.calendarbox .calendarnav-previous, .calendarbox .calendarnav-next { display:block; position:absolute; font-weight:bold; font-size:12px; background:#C9DBED url(../img/admin/default-bg.gif) bottom left repeat-x; padding:1px 4px 2px 4px; color:white; }
.calendarnav-previous:hover, .calendarnav-next:hover { background:#036; }
.calendarnav-previous { top:0; left:0; }
.calendarnav-next { top:0; right:0; }
.calendar-cancel { margin:0 !important; padding:0; font-size:10px; background:#e1e1e1 url(../img/admin/nav-bg.gif) 0 50% repeat-x;  border-top:1px solid #ddd; }
.calendar-cancel a { padding:2px; color:#999; }
ul.timelist, .timelist li { list-style-type:none; margin:0; padding:0; }
.timelist a { padding:2px; }
a img { border:none; }
a:link, a:visited { color: #5b80b2; text-decoration:none; }
a:hover { color: #036; }
-->
</style>

<script type="text/javascript" src="/images/media/js/jsi18n.js"></script>
<script type="text/javascript" src="/images/media/js/core.js"></script>
<script type="text/javascript" src="/images/media/js/admin/RelatedObjectLookups.js"></script>
<script type="text/javascript" src="/images/media/js/calendar.js"></script>
<script type="text/javascript" src="/images/media/js/admin/DateTimeShortcuts.js"></script>
	</HEAD>
	<BODY bgColor="#ffffff" leftMargin="0" topMargin="0" marginheight="0" Marginwidth="0">
		<table width="100%" border="0" cellPadding="0" cellSpacing="0">
				<tr><td width="200px"><img src='/images/new/banner.jpg' width="200" height="50" border="0"></td>
				<td style="background: url(/images/new/banner_dot.jpg)" valign="middle" class="menu"><a href="/write.aspx" target="_self"><%=ph.g("编写网志")%></a> 
					<a href="/system/admin_userinfo.aspx" target="_self"><%=ph.g("个人资料")%></a>
					<a href="/system/admin_cate.aspx" target="_self"><%=ph.g("分类管理")%></a>
					<a href="/system/admin_article.aspx" target="_self"><%=ph.g("网志管理")%></a>
					<a href="/pics.aspx" target="_self"><%=ph.g("图片管理")%></a>
					<a href='/Template/skin_switch.aspx?blogger=<%=Session("blogger")%>' target="_self"><%=ph.g("选择模板")%></a>
					<a href='/system/admin_link.aspx' target="_self"><%=ph.g("链接管理")%></a>
					<a href="/" target="_top"><%=ph.g("返回首页")%></a>
					<a href='/<%=Session("blogger")%>' target="_top"><%=ph.g("返回网志")%></a>
					<a href="/logout.aspx" target="_top"><%=ph.g("退出登陆")%></a></td></tr>
				</table>
		<h3><font face="Verdana">网志备份 </font>
		</h3>
				<form id="Form1" runat="server">
				<div style="padding:10px">
        <p>
            备份</p>
        <p>
            从：<asp:TextBox ID="sd" runat="server" CssClass="vDateField" MaxLength="10"></asp:TextBox></p>
        <p>
            <font face="Verdana">至：<asp:TextBox ID="ed" runat="server" CssClass="vDateField" MaxLength="10"></asp:TextBox></font></p>
        <p>
            的网志</p>

		
			<asp:label id="hint" runat="server" ForeColor="Red"></asp:label><br>
			<asp:button id="GenBackup" runat="server" Text="生成备份"></asp:button></div>
		</form>
	</BODY>
</HTML>
