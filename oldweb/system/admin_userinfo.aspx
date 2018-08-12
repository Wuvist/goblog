<%@ Page CodeFile="admin_userinfo.aspx.vb" Language="vb" AutoEventWireup="false" ValidateRequest="false" Inherits="blogwind.admin_userinfo" %>
<HTML>
	<HEAD>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<link rel="icon" href="/blogwind.ico" type="image/x-icon" />
		<link rel="shortcut icon" href="/blogwind.ico" type="image/x-icon" />
		<meta content="博客，博客风，留学" name="keywords" />
		<meta content="留学生博客网站" name="description" />
		<title><%=ph.g("个人资料")%></title>
		<link rel="stylesheet" href="admin.css" type="text/css" /> 
		<style type="text/css">
		    th{vertical-align:top}
		</style>
	</HEAD>
	<BODY BGCOLOR="#ffffff" Marginwidth="0" marginheight="0" leftmargin="0" topmargin="0">
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
		<form runat="server">
			<h3>博客个人资料管理</h3>
			<table>
				<tr>
					<th>User ID:
					</th>
					<td><asp:Label id="UserID" runat="server" ReadOnly="true" border="1" /></td>
				</tr>
				<tr>
					<th>Nick Name:
					</th>
					<td><asp:TextBox id="NickName" runat="server" ReadOnly="false" /></td>
                </tr>					
				<tr>
					<th>生日:
					</th>
					<td><asp:TextBox id="DOB" runat="server" ReadOnly="false" /></td>
				</tr>
				<tr>
					<th>Blog Name:
					</th>
					<td><asp:TextBox id="BlogName" runat="server" ReadOnly="false" /></td>
				</tr>
				<tr>
					<th>E-Mail:
					</th>
					<td><asp:TextBox id="Email" runat="server" ReadOnly="false" /></td>
				</tr>
				<tr>
					<th>网络日志介绍:</th>
					<td><asp:TextBox id="Intro" runat="server" ReadOnly="false" textmode="multiline" height="100" width="200" /></td>
				</tr>
				<tr>
					<th>Widget:</th>
					<td><asp:TextBox id="widget" runat="server" ReadOnly="false" textmode="multiline" height="100" width="200" /></td>
				</tr>
				<TR>
					<th>是否公开网志：</th>
					<TD>
						<asp:RadioButtonList id="reveal" runat="server" RepeatDirection="Horizontal" Width="169px" Height="2px">
							<asp:ListItem Value="yes" Selected="True">是</asp:ListItem>
							<asp:ListItem Value="No">否</asp:ListItem>
						</asp:RadioButtonList></TD>
				</TR>
				<tr>
					<td></td>
					<td><asp:Button id="Update" runat="server" Text="Submit 更改" /></td>
				</tr>
				<tr>
					<td></td>
					<td><asp:Label id="MessagePage1" runat="server" /></td>
				</tr>
			</table>
			<br>
			<br>
			<table>
				<tr>
					如果你要更换密码, 请填写下面表格:
				</tr>
				<tr>
					<td>
						原密码:
					</td>
					<td>
						<asp:TextBox id="OldPwd" textmode="password" runat="server" />
					</td>
				</tr>
				<tr>
					<td>
						新密码:
					</td>
					<td>
						<asp:TextBox id="NewPwd" textMode="password" runat="server" />
					</td>
				</tr>
				<tr>
					<td>
						重复输入新密码:
					</td>
					<td>
						<asp:TextBox id="ReNewPwd" textMode="password" runat="server" />
					</td>
				</tr>
				<tr>
					<td>
						<asp:Button id="pwd" runat="server" text="change password" />
					</td>
				</tr>
				<tr>
					<td>
						<asp:Label id="MessagePage2" runat="server" />
					</td>
				</tr>
			</table>
		</form>
	</BODY>
</HTML>
