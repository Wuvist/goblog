<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.rssimport" CodeFile="rssimport.aspx.vb" %>
<HTML>
	<HEAD>
		<title>rssimport</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<LINK href="/blogwind.ico" type="image/x-icon" rel="icon">
			<LINK href="blogwind.ico" type="image/x-icon" rel="shortcut icon">
				<meta content="博客，博客风，留学" name="keywords">
				<meta content="留学生博客网站" name="description">
				<style type="text/css">BODY { BORDER-RIGHT: 0px; PADDING-RIGHT: 0px; BORDER-TOP: 0px; PADDING-LEFT: 0px; SCROLLBAR-FACE-COLOR: #85a2e0; FONT-SIZE: 12px; PADDING-BOTTOM: 0px; MARGIN: 5px; SCROLLBAR-HIGHLIGHT-COLOR: #ffffff; BORDER-LEFT: 0px; SCROLLBAR-SHADOW-COLOR: #ffffff; COLOR: #3e62db; SCROLLBAR-3DLIGHT-COLOR: #ffffff; LINE-HEIGHT: 150%; SCROLLBAR-ARROW-COLOR: #ffff00; PADDING-TOP: 0px; SCROLLBAR-TRACK-COLOR: #dad9ea; BORDER-BOTTOM: 0px; FONT-FAMILY: arial,verdana,sans-serif,'宋体'; SCROLLBAR-DARKSHADOW-COLOR: #dad9ea; BACKGROUND-COLOR: #ffffff }
	</style>
	</HEAD>
	<body MS_POSITIONING="FlowLayout">
		<P>
			<form action="rsssave.aspx" method="post">
				<P><asp:label id="view" runat="server"></asp:label><FONT face="宋体"><BR>
					</FONT>
					<asp:label id="hint" runat="server"></asp:label><FONT face="宋体"><BR>
					</FONT>
					<asp:repeater id="cate" runat="server" EnableViewState="False">
						<ItemTemplate>
							<%# Container.DataItem("cate") %>
							<input name=cate type=radio value='<%# Container.DataItem("index") %>'>
						</ItemTemplate>
					</asp:repeater></P>
				<P></P>
				<P>
					<asp:Label id="s" runat="server"></asp:Label></P>
			</form>
			<form id="Form1" method="post" runat="server">
				<P>请输入RSS/ATOM数据源：
					<asp:textbox id="r" runat="server"></asp:textbox><asp:button id="Button1" runat="server" Text="提交"></asp:button></P>
			</form>
	</body>
</HTML>
