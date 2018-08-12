<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.bookmark" validaterequest="false" CodeFile="bookmark.aspx.vb" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<HTML>
	<HEAD>
		<title>���ͷ���ժ</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	</HEAD>
	<body>
		<form id="Form1" method="post" runat="server">
			<table width="475" border="0" cellspacing="0" cellpadding="0">
				<tr>
					<td colspan="3"><div align="center">���ͷ���ժ</div>
					</td>
				</tr>
				<tr>
					<td width="20">&nbsp;</td>
					<td width="20">&nbsp;</td>
					<td>&nbsp;</td>
				</tr>
				<tr>
					<td>&nbsp;</td>
					<td>��ַ��
					</td>
					<td><asp:TextBox ID="url" runat="server"></asp:TextBox></td>
				</tr>
				<tr>
					<td>&nbsp;</td>
					<td>���⣺
					</td>
					<td><asp:TextBox ID="title" runat="server"></asp:TextBox></td>
				</tr>
				<tr>
					<td>&nbsp;</td>
					<td align="left" valign="top">ժҪ��
					</td>
					<td><asp:TextBox ID="des" runat="server" TextMode="MultiLine" MaxLength="400"></asp:TextBox></td>
				</tr>
				<tr>
					<td>&nbsp;</td>
					<td>���ࣺ</td>
					<td>&nbsp;</td>
				</tr>
				<tr>
					<td>&nbsp;</td>
					<td>&nbsp;
					</td>
					<td><asp:Button ID="Button1" runat="server" Text="���"></asp:Button></td>
				</tr>
			</table>
		</form>
	</body>
</HTML>
