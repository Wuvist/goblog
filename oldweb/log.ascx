<%@ Control Language="vb" AutoEventWireup="false" Inherits="blogwind.log" CodeFile="log.ascx.vb" %>
<% if request.querystring("blogger")<>"" then%>
<A href="http://www.blogwind.com/rss.aspx?blogger=<%=request.querystring("blogger")%>"><IMG alt="RSS 2.0" src="/images/rss2.png" border="0" width="80" height="15"></A><br><br><%end if%>
<%if Session("blogger")="" then%>
<form name="Form2" action="/system/check.aspx" method="post" id="Form2"><div class=log><%=ph.g("用户名")%>：<input id="id" class="loginput" type="text" name="id"><br><%=ph.g("密　码")%>：<input id="pass" class="loginput" type="password" name="pass"><br>
		<INPUT type="checkbox" CHECKED name="record">&nbsp;<%=ph.g("记录我的登陆信息")%><br>
		&nbsp;
		<br><div align=center><input id="Submit" class="logbutton" type="submit" value="<%=ph.g("登录")%>" name="Submit">&nbsp;<INPUT id="Submit1" type='button' class='logbutton' value="<%=ph.g("申请")%>" onclick="javascript:window.open('/apply.aspx','_top')">&nbsp;<INPUT id="Submit1" type='button' class='logbutton' value="<%=ph.g("忘记密码")%>" onclick="javascript:window.open('/getpass.aspx','_self')"></div>
</form></div>
<%else%>
<%=nick%>，<% =welcomemsg
          %>
<br><br><a href='/randomblog.aspx' target="_blank"><%=ph.g("阅读随机网志")%></a>
<br>
<%=showfriend%><br>
<a style="CURSOR: hand" onclick="toggle('admin')"><%=ph.g("网志管理")%></a>
<a href='http://wiki.blogwind.com/doku.php' target="_blank"><%=ph.g("网站帮助")%></a><br />
<ul id="admin" style="display:none">
<li><a href="/system/admin_userinfo.aspx"><%=ph.g("个人资料")%></a>
<li><a href="/system/admin_cate.aspx"><%=ph.g("分类管理")%></a>
<li><a href="/system/admin_article.aspx"><%=ph.g("网志管理")%></a>
<li><a href="/pics.aspx"><%=ph.g("图片管理")%></a>
<li><a href='/Template/skin_switch.aspx?blogger=<%=Session("blogger")%>'><%=ph.g("选择模板")%></a>
<li><a href='/system/admin_link.aspx'><%=ph.g("链接管理")%></a>
<li><a href="/rssimport.aspx"><%=ph.g("RSS导入")%></a>
<li><a href="/system/admin_backup.aspx"><%=ph.g("网志备份")%></a>
<li><a style="CURSOR: hand" onclick="toggle('setlangs')"><%=ph.g("界面语言")%></a>
<ol id="setlangs" style="display:none">
<li><a href="/setlang.aspx?lang=0"><%=ph.g("中文")%></a>
<li><a href="/setlang.aspx?lang=1"><%=ph.g("English")%></a>
<li><a href="/setlang.aspx?lang=2"><%=ph.g("Deutsch")%></a>
<li><a href="/setlang.aspx?lang=3"><%=ph.g("日本語")%></a>
</ol>
</ul>
<A href="/write.aspx"><%=ph.g("编写网志")%></A>
<A href="/system/logout.aspx"><%=ph.g("退出登陆")%></A><br />
<%end if%>