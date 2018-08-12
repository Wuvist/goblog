<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.updateskin" CodeFile="updateskin.aspx.vb" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<HTML>
	<HEAD>
		<title>updateskin</title>
		<meta name="GENERATOR" content="Microsoft Visual Studio .NET 7.1">
		<meta name="CODE_LANGUAGE" content="Visual Basic .NET 7.1">
		<meta name="vs_defaultClientScript" content="JavaScript">
		<meta name="vs_targetSchema" content="http://schemas.microsoft.com/intellisense/ie5">
	</HEAD>
	<body>
		<form id="Form1" method="post" runat="server">
			<asp:DataGrid id="skinlist" runat="server" AutoGenerateColumns="False" Width="950px" BorderColor="#999999"
				BorderStyle="Solid" BorderWidth="1px" BackColor="White" CellPadding="3" GridLines="Vertical"
				ForeColor="Black" DataKeyField="index">
				<SelectedItemStyle Font-Bold="True" ForeColor="White" BackColor="#000099"></SelectedItemStyle>
				<AlternatingItemStyle BackColor="#CCCCCC"></AlternatingItemStyle>
				<HeaderStyle Font-Bold="True" ForeColor="White" BackColor="Black"></HeaderStyle>
				<FooterStyle BackColor="#CCCCCC"></FooterStyle>
				<Columns>
					<asp:BoundColumn DataField="name" HeaderText="模板名称">
						<HeaderStyle Width="100px"></HeaderStyle>
					</asp:BoundColumn>
					<asp:BoundColumn DataField="designer" HeaderText="设计者">
						<HeaderStyle Width="150px"></HeaderStyle>
					</asp:BoundColumn>
					<asp:BoundColumn DataField="intro" HeaderText="简介">
						<HeaderStyle Width="450px"></HeaderStyle>
					</asp:BoundColumn>
					<asp:BoundColumn DataField="popularity" ReadOnly="True" HeaderText="使用人数">
						<HeaderStyle Width="120px"></HeaderStyle>
					</asp:BoundColumn>
					<asp:EditCommandColumn ButtonType="LinkButton" UpdateText="Update" HeaderText="修改" CancelText="Cancel"
						EditText="Edit">
						<HeaderStyle Width="110px"></HeaderStyle>
					</asp:EditCommandColumn>
				</Columns>
				<PagerStyle HorizontalAlign="Center" ForeColor="Black" BackColor="#999999"></PagerStyle>
			</asp:DataGrid>
			<P><FONT face="宋体">模板名称：
					<asp:TextBox id="nname" runat="server" MaxLength="5"></asp:TextBox><BR>
					设计者：
					<asp:TextBox id="ndesigner" runat="server" MaxLength="10"></asp:TextBox><BR>
					简介：
					<asp:TextBox id="nintro" runat="server" Width="344px" MaxLength="100"></asp:TextBox><BR>
					<asp:Button id="Button2" runat="server" Text="添加"></asp:Button></FONT></P>
			<P><FONT face="宋体">
					<asp:Label id="hint" runat="server"></asp:Label></P>
			</FONT>
			<P>
				<asp:RadioButtonList id="skins" runat="server" DataValueField="index" DataTextField="name"></asp:RadioButtonList>
				<asp:Button id="Button1" runat="server" Text="Button"></asp:Button></P>
		</form>
	</body>
</HTML>
