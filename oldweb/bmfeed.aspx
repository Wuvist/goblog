<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.bmfeed" CodeFile="bmfeed.aspx.vb" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<HTML>
	<HEAD>
		<title>bmfeed</title>
		<meta name="GENERATOR" content="Microsoft Visual Studio .NET 7.1">
		<meta name="CODE_LANGUAGE" content="Visual Basic .NET 7.1">
		<meta name="vs_defaultClientScript" content="JavaScript">
		<meta name="vs_targetSchema" content="http://schemas.microsoft.com/intellisense/ie5">
	</HEAD>
	<body>
		<form id="Form1" method="post" runat="server"><ul>
			<asp:Repeater id="fed" runat="server" EnableViewState="False">
				<ItemTemplate><li><a title="<%# DataBinder.Eval(Container.DataItem, "des") %>" href="<%# DataBinder.Eval(Container.DataItem, "url") %>" target="_blank"><%# DataBinder.Eval(Container.DataItem, "title") %></a></li></ItemTemplate>
			</asp:Repeater></ul>
		</form>
	</body>
</HTML>
