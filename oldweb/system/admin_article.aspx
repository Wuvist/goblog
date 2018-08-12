<%@ Page CodeFile="admin_article.aspx.vb" Language="vb" AutoEventWireup="false" Inherits="admin_article" %>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<link rel="icon" href="/blogwind.ico" type="image/x-icon" />
<link rel="shortcut icon" href="blogwind.ico" type="image/x-icon" />
<meta content="博客，博客风，留学" name="keywords">
<meta content="留学生博客网站" name="description">
<link rel="stylesheet" href="admin.css" type="text/css"> 
</head>
<BODY BGCOLOR="FFFFFF" Marginwidth=0 marginheight=0 leftmargin=0 topmargin=0>
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
	<h3><font face="Verdana">日志管理页面</font></h3>
<form runat="server">

	<asp:DataGrid id="dgArticles"
			runat="server"
			Font-Name="Verdana"
			Font-Size="10"
			Font-Color="White"
			AlternatingfItemStyle-BackColor="lightGray" 			
			AutoGenerateColumns="false"
			DataKeyField="index"
			CommandArgument="cate"
			AllowPaging="True"
			>

		<Columns>
			
			<asp:BoundColumn
				HeaderText="发表日期"
				DataField="add_date" />
			
			<asp:BoundColumn
				HeaderText="日志标题"
				DataField="title" 
				ItemStyle-width="250px"/>
			
			<asp:BoundColumn
				HeaderText="留言数目"
				DataField="Comment" />

 			<asp:ButtonColumn
				Text="编辑"
				HeaderText="编辑"
				ItemStyle-width="50px"
				CommandName="Edit"/>

			<asp:ButtonColumn
				HeaderText="删除"
				Text="删除"
				CommandName="Delete"
                />
		</Columns>
		<PagerStyle NextPageText="上一页" PrevPageText="下一页" HorizontalAlign="Right" ForeColor="Black"
		BackColor="#9BB6F9" PageButtonCount="22" Mode="NumericPages"></PagerStyle>
    </asp:DataGrid>

</form>
</body>
</html>
