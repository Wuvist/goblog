<%@ Reference Page="~/email.aspx" %>
<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.confirm" CodeFile="confirm.aspx.vb" %>
<%@ Register TagPrefix="uc1" TagName="bottom" Src="bottom.ascx" %>
<HTML>
	<HEAD>
		<title>博客风－注册博客成功</title>
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
									<P><%=nick%>，您的博客已经注册成功！</P>
									<P>请查取您的电子信箱，使用博客风系统发送的信件中的连接以确认帐号。在确认帐号之前，您将无法使用博客风的所有功能。</P>
									<P>
										<asp:Label id="hint" runat="server"></asp:Label>倘若您在十分钟后仍无法收到系统信件，请确认您的信箱是否为：<BR>
										<asp:TextBox id="email" runat="server"></asp:TextBox><BR>
										倘若信箱有误，请在直接修改上方的信箱地址，并
										<asp:Button id="Button1" runat="server" Text="点此按钮重新发送确认信件"></asp:Button>。</P>
									<P>您可以使用：<BR>
										&nbsp;&nbsp;&nbsp; <A href="http://www.blogwind.com/<%=blogger%>">http://www.blogwind.com/<%=blogger%></A></P>
									<P>访问您的博客。</P>
									<P>请点此进入<a href="/write.aspx">管理后台</a>，发表您在博客风上的第一篇blog吧！</P>
									<P>如果您不熟悉博客风的系统使用，可以点击这里查看<a href="http://wiki.blogwind.com/">博客风的帮助系统</a>。</P>
								</td>
							</tr>
						</table>
						<FONT size="2">
							<DIV align="center">
								<uc1:bottom id="Bottom1" runat="server"></uc1:bottom>
							</DIV>
						</FONT>
					</td>
				</tr>
			</table>
		</form>
	</body>
</HTML>
