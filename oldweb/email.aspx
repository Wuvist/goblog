<%@ Register TagPrefix="uc1" TagName="bottom" Src="bottom.ascx"  %>
<%@ Page Language="vb" Inherits="blogwind.email" CodeFile="email.aspx.vb" AutoEventWireup="false" validaterequest="false"%>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<HTML>
	<HEAD>
		<title>博客风－联系</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<LINK href="/blogwind.ico" type="image/x-icon" rel="icon">
			<LINK href="blogwind.ico" type="image/x-icon" rel="shortcut icon">
				<meta content="博客，博客风，留学" name="keywords">
				<meta content="留学生博客网站" name="description">
				<style type="text/css">.loginput { BORDER-RIGHT: #b6caff 1px solid; BORDER-TOP: #b6caff 1px solid; FONT-SIZE: 12px; BORDER-LEFT: #b6caff 1px solid; WIDTH: 250px; LINE-HEIGHT: 15px; BORDER-BOTTOM: #b6caff 1px solid }
	.logbutton { BORDER-RIGHT: #aaaaaa 1px solid; BORDER-TOP: #aaaaaa 1px solid; FONT-SIZE: 12px; BORDER-LEFT: #aaaaaa 1px solid; LINE-HEIGHT: 15px; BORDER-BOTTOM: #aaaaaa 1px solid; FONT-FAMILY: 宋体; BACKGROUND-COLOR: #d9e3ff }
	</style>
	</HEAD>
	<body>
		<form id="Form1" method="post" runat="server">
			<table borderColor="#aaaaaa" cellSpacing="0" cellPadding="0" width="780" align="center"
				border="1">
				<tr>
					<td><IMG height="123" src="images/titlev2.jpg" width="778">
						<table cellSpacing="0" cellPadding="0" width="100%" border="0">
							<tr>
								<td width="100">&nbsp;</td>
								<td>
									<table cellSpacing="0" cellPadding="0" width="100%" border="0">
										<TBODY>
											<tr>
												<td>&nbsp;</td>
												<td>&nbsp;</td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="80">&nbsp;</td>
												<td>您通过电子邮件要给
													<%=nick%>
													留言么？</td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td>&nbsp;</td>
												<td>&nbsp;</td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td>您的姓名：</td>
												<td><asp:textbox id="name" runat="server" Width="352px"></asp:textbox>
													<asp:RequiredFieldValidator id="RequiredFieldValidator1" runat="server" ErrorMessage="您似乎忘了填写您的姓名" ControlToValidate="name"></asp:RequiredFieldValidator></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td vAlign="top">您的电邮：</td>
												<td><asp:textbox id="e_mail" runat="server" Width="352px"></asp:textbox>
													<asp:RequiredFieldValidator id="RequiredFieldValidator2" runat="server" ErrorMessage="您似乎忘了填写您的电邮" ControlToValidate="e_mail"></asp:RequiredFieldValidator><BR>
													<asp:RegularExpressionValidator id="RegularExpressionValidator1" runat="server" ControlToValidate="e_mail" ErrorMessage="请填写正确的电邮地址"
														ValidationExpression="\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*"></asp:RegularExpressionValidator></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td height="23">主题：</td>
												<td height="23"><asp:textbox id="title" runat="server" Width="352px"></asp:textbox>
													<asp:RequiredFieldValidator id="RequiredFieldValidator3" runat="server" ErrorMessage="您似乎忘了填写留言的主题。" ControlToValidate="title"></asp:RequiredFieldValidator></td>
												<td height="23">&nbsp;</td>
											</tr>
											<tr>
												<td vAlign="top" align="left">内容：</td>
												<td><asp:textbox id="content" runat="server" Width="352px" Height="168px" TextMode="MultiLine"></asp:textbox><BR>
													<asp:RequiredFieldValidator id="RequiredFieldValidator4" runat="server" ErrorMessage="您似乎忘了填写留言的内容。" ControlToValidate="content"></asp:RequiredFieldValidator></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td>&nbsp;</td>
												<td>&nbsp;
													<asp:label id="hint" runat="server" ForeColor="Red"></asp:label></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td>&nbsp;</td>
												<td><asp:button id="Button1" runat="server" Text="发送"></asp:button></td>
												<td>&nbsp;</td>
											</tr>
										</TBODY>
									</table>
								</td>
							</tr>
						</table>
						<div>
							<div align="center"><font size="2"><uc1:bottom id="Bottom1" runat="server"></uc1:bottom></font></div>
						</div>
					</td>
				</tr>
			</table>
		</form>
	</body>
</HTML>
