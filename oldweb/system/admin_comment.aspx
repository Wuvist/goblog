<%@ Page Language="vb" AutoEventWireup="false" validaterequest="false" CodeFile="admin_comment.aspx.vb" Inherits="blogwind.admin_comment" %>
<HTML>
	<HEAD>
		<meta http-equiv="Content-Type" content="text/html; charset=gb2312">
		<LINK href="/blogwind.ico" type="image/x-icon" rel="icon">
			<LINK href="blogwind.ico" type="image/x-icon" rel="shortcut icon">
				<meta content="博客，博客风，留学" name="keywords">
				<meta content="留学生博客网站" name="description">
				<LINK href="admin.css" type="text/css" rel="stylesheet">
	</HEAD>
<BODY BGCOLOR="FFFFFF" Marginwidth=0 marginheight=0 leftmargin=0 topmargin=0>
			<img src='/images/system/title3.jpg' width="778" height="141" border="0" usemap="#Map">
			<map name="Map">
				<area shape="RECT" coords="10,120,69,138" href="/write.aspx" target="_self">
				<area shape="RECT" coords="89,121,148,139" href="/system/admin_userinfo.aspx" target="_self">
				<area shape="RECT" coords="164,120,226,136" href="/system/admin_cate.aspx" target="_self">
				<area shape="RECT" coords="241,121,301,136" href="/system/admin_article.aspx" target="_self">
				<area shape="RECT" coords="320,120,379,138" href="/system/admin_comment.aspx" target="_self">
				<area shape="RECT" coords="398,121,457,137" href='/Template/skin_switch.aspx?blogger=<%=Session("blogger")%>' target="_self">
				<area shape="RECT" coords="475,121,537,137" href='/system/admin_link.aspx' target="_self">
				<area shape="RECT" coords="630,121,688,136" href='http://www.blogwind.com/<%=Session("blogger")%>' target="_top">
				<area shape="RECT" coords="707,121,766,135" href="/logout.aspx" target="_top">
				<area shape="rect" coords="552,121,609,136" href="/" target="_top">
		  </map>
		<h3>回复管理页面</h3>
		<form runat="server">
			<asp:DataGrid id="dgComment" runat="server" Font-Name="Verdana" Font-Size="10pt" Font-Color="White"
				AlternatingfItemStyle-BackColor="lightGray"  AutoGenerateColumns="False"
				DataKeyField="index" AllowPaging="True" Font-Names="Verdana">
				<Columns>
					<asp:BoundColumn DataField="add_date" HeaderText="发表日期"></asp:BoundColumn>
					<asp:BoundColumn DataField="title" HeaderText="日志标题">
						<ItemStyle Width="250px"></ItemStyle>
					</asp:BoundColumn>
					<asp:BoundColumn DataField="author" HeaderText="作者"></asp:BoundColumn>
					<asp:ButtonColumn Text="删除" HeaderText="删除" CommandName="Delete"></asp:ButtonColumn>
				</Columns>
				<PagerStyle NextPageText="上一页" PrevPageText="下一页" HorizontalAlign="Right" ForeColor="Black"
					BackColor="#9BB6F9" PageButtonCount="22" Mode="NumericPages"></PagerStyle>
			</asp:DataGrid>
		</form>
	</body>
</HTML>
