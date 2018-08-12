<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.changepass" CodeFile="changepass.aspx.vb" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<HTML>
	<HEAD>
		<title>changepass</title>
		<meta name="GENERATOR" content="Microsoft Visual Studio .NET 7.1">
		<meta name="CODE_LANGUAGE" content="Visual Basic .NET 7.1">
		<meta name="vs_defaultClientScript" content="JavaScript">
		<meta name="vs_targetSchema" content="http://schemas.microsoft.com/intellisense/ie5">
	</HEAD>
	<body>
		<form id="Form1" method="post" runat="server">
			<P><FONT face="宋体"><%=request.querystring("id")%>，欢迎您使用博客风密码恢复功能：</FONT></P>
			<P><FONT face="宋体">请输入新密码：</FONT>
				<asp:TextBox id="p1" runat="server" TextMode="Password"></asp:TextBox></P>
			<P><FONT face="宋体">请重复密码&nbsp; ：</FONT>
				<asp:TextBox id="p2" runat="server" TextMode="Password"></asp:TextBox><FONT face="宋体">&nbsp;
					<asp:CompareValidator id="CompareValidator1" runat="server" ErrorMessage="两次密码输入不一致！" ControlToCompare="p1"
						ControlToValidate="p2"></asp:CompareValidator></FONT></P>
			<P><FONT face="宋体">
					<asp:Label id="hint" runat="server" ForeColor="Red"></asp:Label></P>
			</FONT>
			<P>
				<asp:Button id="Button1" runat="server" Text="提交"></asp:Button></P>
		</form>
	</body>
</HTML>
