<%@ Page Language="VB" MasterPageFile="~/master/base.master" AutoEventWireup="false" CodeFile="decide.aspx.vb" Inherits="openid_decide" title="Untitled Page" %>

<asp:Content ID="Content1" ContentPlaceHolderID="title" Runat="Server">
</asp:Content>
<asp:Content ID="Content2" ContentPlaceHolderID="css" Runat="Server">
</asp:Content>
<asp:Content ID="Content3" ContentPlaceHolderID="content" Runat="Server">
<form id="form1" runat="server">
<p>
		A site has asked to authenticate that you own the identifier below.&nbsp; You should 
		only do this if you wish to log in to the site given by the Realm.</p>
	<p>
		This site
		<asp:Label ID="relyingPartyVerificationResultLabel" runat="server" 
			Font-Bold="True" Text="failed" /> verification. </p>
	<table>
		<tr>
			<td>
				Identifier: 			</td>
			<td>
				<asp:Label runat="server" ID='identityUrlLabel' />
			</td>
		</tr>
		<tr>
			<td>
				Realm: 			</td>
			<td>
				<asp:Label runat="server" ID='realmLabel' />
			</td>
		</tr>
	</table>
	<p>
		Allow this authentication to proceed?
	</p>
	<asp:Button ID="yes_button" OnClick="Yes_Click" Text="  yes  " runat="Server" />
	<asp:Button ID="no_button" OnClick="No_Click" Text="  no  " runat="Server" />
</form>
</asp:Content>
<asp:Content ID="Content4" ContentPlaceHolderID="js" Runat="Server">
</asp:Content>

