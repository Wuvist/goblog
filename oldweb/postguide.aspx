<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<HTML>
	<HEAD>
		<title>博客风－导航</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<link rel="icon" href="/blogwind.ico" type="image/x-icon">
			<link rel="shortcut icon" href="blogwind.ico" type="image/x-icon">
				<meta content="博客，博客风，网志" name="keywords">
				<link rel="stylesheet" href="write.css" type="text/css">
	</HEAD>
	<BODY BGCOLOR="#ffffff" Marginwidth="0" marginheight="0" leftmargin="0" topmargin="0">
			    <table width="100%" border="0" cellPadding="0" cellSpacing="0">
				<tr><td width="200px"><img src='/images/new/banner.jpg' width="200" height="50" border="0"></td>
				<td style="background: url(/images/new/banner_dot.jpg)" valign="middle" class="menu"><a href="/write.aspx" target="_self">编写网志</a>
				<a href="/system/admin_userinfo.aspx" target="_self">个人资料</a>
				<a href="/system/admin_cate.aspx" target="_self">分类管理</a>
				<a href="/system/admin_article.aspx" target="_self">网志管理</a>
				<a href='/Template/skin_switch.aspx?blogger=<%=Session("blogger")%>' target="_self">选择模板</a>
				<a href='/system/admin_link.aspx' target="_self">链接管理</a>
				<a href="/" target="_top">返回首页</a>
				<a href='/<%=Session("blogger")%>' target="_top">返回网志</a>
				<a href="/logout.aspx" target="_top">退出登陆</a></td></tr>
				</table>
<%
dim action as String
dim aid as Integer
action=request.querystring("action")
aid=request.querystring("aid")

if action="added" then
%><h3>网志发表成功！</h3> 接下来您要：
<ul class="menu">
<li><a href="/<%=Session("blogger") & "/" & aid %>.shtml">查看刚刚发表的网志</a></li>
<li><a href="/write.aspx">继续编写新的网志</a></li>
<li><a href="/">回博客风首页瞅瞅</a></li>
<li><a href="/randomblog.aspx" target="_blank">还是觉得有点无聊，让博客风随机给您一篇网志看看呢？</a></li>
</ul>
<%elseif action="edited" then%>
<h3>网志修改成功！</h3> 接下来您要：
<ul class="menu">
<li><a href="/<%=Session("blogger") & "/" & aid %>.shtml">查看刚刚修改的网志</a></li>
<li><a href="/write.aspx">编写新的网志</a></li>
<li><a href="/">回博客风首页瞅瞅</a></li>
<li><a href="/randomblog.aspx" target="_blank">还是觉得有点无聊，让博客风随机给您一篇网志看看呢？</a></li>
</ul>
<%elseif action="deleted" then%>
<h3>网志删除成功！</h3> 接下来您要：
<ul class="menu">
<li><a href="/write.aspx">编写新的网志</a></li>
<li><a href="/">回博客风首页瞅瞅</a></li>
<li><a href="/randomblog.aspx" target="_blank">还是觉得有点无聊，让博客风随机给您一篇网志看看呢？</a></li>
</ul>
<%end if%>		
	</BODY>
</HTML>
