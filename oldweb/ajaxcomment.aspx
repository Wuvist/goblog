<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.ajaxcomment" CodeFile="ajaxcomment.aspx.vb" %>
<asp:Repeater id="comments" runat="server" enableViewState="False">							
	<ItemTemplate>
		<div id="white">
		【<%# DataBinder.Eval(Container.DataItem, "blogname") %>】<a href='<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "Article") %>.shtml#blogc<%# DataBinder.Eval(Container.DataItem, "cid") %>' target="_blank"><%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "title")) %></a>-<%# DataBinder.Eval(Container.DataItem, "Author") %>
		</div>
	</ItemTemplate>
	<AlternatingItemTemplate>
	<div id="grey">
		【<%# DataBinder.Eval(Container.DataItem, "blogname") %>】<a href='<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "Article") %>.shtml#blogc<%# DataBinder.Eval(Container.DataItem, "cid") %>' target="_blank"><%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "title")) %></a>-<%# DataBinder.Eval(Container.DataItem, "Author") %>
		</div>
	</AlternatingItemTemplate>
</asp:Repeater>
