<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.news" CodeFile="news.aspx.vb" %>
<%@ Register TagPrefix="uc1" TagName="bottom" Src="bottom.ascx" %>
<HTML>
	<HEAD>
		<title>博客风公告板</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<link rel="icon" href="/blogwind.ico" type="image/x-icon">
			<link rel="shortcut icon" href="blogwind.ico" type="image/x-icon">
				<meta content="博客，博客风，留学" name="keywords">
				<meta content="留学生博客网站" name="description">
				<style type="text/css"> .loginput { BORDER-RIGHT: #b6caff 1px solid; BORDER-TOP: #b6caff 1px solid; FONT-SIZE: 12px; BORDER-LEFT: #b6caff 1px solid; WIDTH: 250px; LINE-HEIGHT: 15px; BORDER-BOTTOM: #b6caff 1px solid } .logbutton { BORDER-RIGHT: #aaaaaa 1px solid; BORDER-TOP: #aaaaaa 1px solid; FONT-SIZE: 12px; BORDER-LEFT: #aaaaaa 1px solid; LINE-HEIGHT: 15px; BORDER-BOTTOM: #aaaaaa 1px solid; FONT-FAMILY: 宋体; BACKGROUND-COLOR: #d9e3ff } </style>
	</HEAD>
	<body>
		<form id="Form1" method="post" runat="server">
			<table width="780" border="1" align="center" cellpadding="0" cellspacing="0" bordercolor="#aaaaaa">
				<tr>
					<td><img src="images/titlev2.jpg" width="780" height="124">
						<table width="100%" border="0" cellspacing="0" cellpadding="0">
							<tr>
								<td width="100">&nbsp;</td>
								<td><p align="center"><%=subject%></p>
									<p><%=content%></p>
								</td>
								<td width="100">&nbsp;</td>
							</tr>
						</table>
							<div align="center"><font size="2">
											<uc1:bottom id="Bottom1" runat="server"></uc1:bottom></div></FONT></td>
				</tr>
			</table>
		</form>
	</body>
</HTML>
