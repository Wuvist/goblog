<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.userpic_change" CodeFile="userpic_change.aspx.vb" %>
<HTML>
	<HEAD>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	</HEAD>
	<body>
		<form id="Form1" method="post" runat="server">
			<input id="Up_file" style="FONT-SIZE: 12px; WIDTH: 120px; FONT-FAMILY: '宋体'" type="file"
				size="8" name="Up_file" runat="server">
			<BR>
			<asp:RequiredFieldValidator id="RequiredFieldValidator1" runat="server" ErrorMessage="请选择头像上传" ControlToValidate="Up_file"></asp:RequiredFieldValidator><BR>
			<asp:Button id="Button1" runat="server" Text="保存" style="BORDER-RIGHT: #aaaaaa 1px solid; BORDER-TOP: #aaaaaa 1px solid; FONT-SIZE: 9pt; BORDER-LEFT: #aaaaaa 1px solid; LINE-HEIGHT: 15px; BORDER-BOTTOM: #aaaaaa 1px solid; FONT-FAMILY: 宋体; BACKGROUND-COLOR: #F6E39E"></asp:Button><BR>
			<asp:Label id="hint" runat="server" ForeColor="Red" Font-Size="10pt"></asp:Label></form>
	</body>
</HTML>
