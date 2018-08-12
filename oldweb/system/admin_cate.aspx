<%@ Import Namespace ="System.Data.SQLClient" %>
<%@ Import Namespace ="System.Data" %>
<%@ Import Namespace ="System" %>
<%@ Page CodeFile="admin_cate.aspx.vb" Language="vb" AutoEventWireup="false" Inherits="admin_cate" %>
<HTML>
	<HEAD>
		<meta http-equiv="Content-Type" content="text/html; charset=gb2312">
		<LINK href="/blogwind.ico" type="image/x-icon" rel="icon">
			<LINK href="blogwind.ico" type="image/x-icon" rel="shortcut icon">
				<meta content="博客，博客风，留学" name="keywords">
				<meta content="留学生博客网站" name="description">
				<link rel="stylesheet" href="admin.css" type="text/css">
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
		<h3><font face="Verdana">文章分类管理页面 </font>
		</h3>
		<form runat="server">
			<asp:datagrid id="dgCate" runat="server" OnItemCreated="dgCate_ItemCreated" OnDeleteCommand="DeleteCate"
				DataKeyField="index" AutoGenerateColumns="false" OnUpdateCommand="UpdateCate" OnEditCommand="EditCate"
				AlternatingfItemStyle-BackColor="lightGray" Font-Color="White" Font-Size="10" Font-Name="Verdana">
				<Columns>
					<asp:BoundColumn HeaderText="日志分类" DataField="cate" ItemStyle-width="200px" />
					<asp:BoundColumn HeaderText="文章数目" DataField="files" ItemStyle-width="200px" />
					<asp:EditCommandColumn EditText="编辑" UpdateText="更新" HeaderText="编辑" />
					<asp:ButtonColumn HeaderText="删除" Text="删除" CommandName="Delete" />
				</Columns>
			</asp:datagrid><asp:label id="hint" runat="server" ForeColor="Red"></asp:label><br>
			<table>
				<tr>
					<td>添加新分类：
					</td>
					<td><asp:textbox id="NewCate" runat="server"></asp:textbox></td>
					<td><asp:button id="AddNewCate" runat="server" Text="添加"></asp:button></td>
				</tr>
			</table>
			<table width="600" border="1">
				<tr>
					<td align="right">将选择的文章分类到：
						<asp:dropdownlist id="ddlCate" runat="server" onselectedindexchanged="ddlCate_SelectedIndexChanged"
							AutoPostBack="true" DataTextField="cate" DataValueField="index"></asp:dropdownlist><asp:button id="UpdateArticleCate" runat="server" text="更新"></asp:button>
				<tr>
					<td><asp:datagrid id="dgArticlesCate" runat="server" DataKeyField="index" AutoGenerateColumns="false"
							AlternatingfItemStyle-BackColor="lightGray" Font-Color="White" Font-Size="10" Font-Name="Verdana"
							AllowPaging="True">
							<Columns>
								<asp:BoundColumn DataField="cate" visible="false" />
								<asp:BoundColumn HeaderText="发表日期" DataField="add_date" />
								<asp:BoundColumn HeaderText="日志标题" DataField="title" ItemStyle-width="300px" />
								<asp:TemplateColumn HeaderText="属于所选分类" ItemStyle-width="150px">
									<ItemTemplate>
										<asp:CheckBox ID="myCheckbox" Runat="server" />
									</ItemTemplate>
								</asp:TemplateColumn>
							</Columns>
							<PagerStyle NextPageText="上一页" PrevPageText="下一页" HorizontalAlign="Right" ForeColor="Black"
								BackColor="#9BB6F9" PageButtonCount="22" Mode="NumericPages"></PagerStyle>
						</asp:datagrid></td>
				</tr>
			</table>
		</form>
	</BODY>
</HTML>
