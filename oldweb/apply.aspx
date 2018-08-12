<%@ Register TagPrefix="uc1" TagName="bottom" Src="bottom.ascx" %>
<%@ Page Language="vb" AutoEventWireup="false" CodeFile="apply.aspx.vb" Inherits="blogwind.apply"%>
<HTML>
	<HEAD>
		<title>博客风－注册博客</title>
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
												<td width="90">&nbsp;</td>
												<td>&nbsp;</td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">用户名:</td>
												<td><asp:textbox class="loginput" id="userid" runat="server" maxLength="50"></asp:textbox>&nbsp;<asp:requiredfieldvalidator id="RequiredFieldValidator1" runat="server" ErrorMessage="请输入用户名" ControlToValidate="userid"></asp:requiredfieldvalidator></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">&nbsp;</td>
												<td>（用户名请使用<font color="#ff0000">英文</font>，并且不包含空格。用户名将会决定您的网志网址，如果您的用户名是Wuvist，那么您的博客网址便是：<A href="http://www.blogwind.com/Wuvist">http://www.blogwind.com/Wuvist</A>）</td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">昵称：</td>
												<td><asp:textbox class="loginput" id="nick" runat="server" maxLength="50"></asp:textbox></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td height="23" width="90">密码：</td>
												<td height="23"><asp:textbox class="loginput" id="pw1" runat="server" TextMode="Password"></asp:textbox>
													<asp:RequiredFieldValidator id="RequiredFieldValidator2" runat="server" ControlToValidate="pw1" ErrorMessage="请输入密码"></asp:RequiredFieldValidator></td>
												<td height="23">&nbsp;</td>
											</tr>
											<tr>
												<td width="90">重复密码：</td>
												<td><asp:textbox class="loginput" id="pw2" runat="server" TextMode="Password"></asp:textbox><asp:comparevalidator id="pw" runat="server" ErrorMessage="两次密码输入不一致" ControlToCompare="pw2" ControlToValidate="pw1"></asp:comparevalidator></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">Email：</td>
												<td><asp:textbox class="loginput" id="email" runat="server"></asp:textbox>
													<asp:RegularExpressionValidator id="RegularExpressionValidator1" runat="server" ControlToValidate="email" ErrorMessage="您填写的信箱地址不正确"
														ValidationExpression="\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*"></asp:RegularExpressionValidator></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">&nbsp;</td>
												<td>（请务必填写<font color="#ff0000">正确的email地址</font>，否则您将无法使用博客风的所有功能。）</td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">重复Email：</td>
												<td><asp:textbox class="loginput" id="remail" runat="server"></asp:textbox><asp:comparevalidator id="CompareValidator1" runat="server" ErrorMessage="两次email地址输入不一致" ControlToCompare="email"
														ControlToValidate="remail"></asp:comparevalidator></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">生日：</td>
												<td><asp:textbox class="loginput" id="DOB" runat="server"></asp:textbox></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">&nbsp;</td>
												<td>(格式如：1982-09-17)</td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">网志名称：</td>
												<td><asp:textbox class="loginput" id="blogname" runat="server"></asp:textbox>
													<asp:RequiredFieldValidator id="RequiredFieldValidator3" runat="server" ControlToValidate="blogname" ErrorMessage="请输入网志名称"></asp:RequiredFieldValidator></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td vAlign="top" align="left" width="90">网志简介：<br>
													(最多125字)</td>
												<td><asp:textbox class="loginput" id="intro" runat="server" TextMode="MultiLine" Rows="8"></asp:textbox></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">&nbsp;</td>
												<td><FONT color="#ff0000"><asp:label id="Hint" runat="server" Width="256px"></asp:label></FONT></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td width="90">&nbsp;</td>
												<td><INPUT class="logbutton" id="Submit1" type="submit" value="提交" name="submit" runat="server">
												</td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td vAlign="top" align="left" colSpan="2" rowSpan="6"><font color="#ff0000"></font></td>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td>&nbsp;</td>
											</tr>
											<tr>
												<td>&nbsp;</td>
											</tr>
										</TBODY></table>
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
