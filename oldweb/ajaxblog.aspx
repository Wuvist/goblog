<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.ajaxblog" CodeFile="ajaxblog.aspx.vb" %>
<asp:Repeater id="blogs" runat="server" enableViewState="False">
<ItemTemplate>
	<div id="white">
		<%# DataBinder.Eval(Container.DataItem, "nick") %> 发表了<a href='<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "[index]") %>.shtml' target="_blank">【<%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "title")) %>】</a><font color="DodgerBlue">评论：<%# DataBinder.Eval(Container.DataItem, "Comment") %></font>
		</div>
	</ItemTemplate>
<AlternatingItemTemplate>
	<div id="grey">
		<%# DataBinder.Eval(Container.DataItem, "nick") %> 发表了<a href='<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "[index]") %>.shtml' target="_blank">【<%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "title")) %>】</a><font color="DodgerBlue">评论：<%# DataBinder.Eval(Container.DataItem, "Comment") %></font>
		</div>
</AlternatingItemTemplate>
</asp:Repeater>
