<%@ Page Language="VB" MasterPageFile="~/master/base.master" AutoEventWireup="false" CodeFile="list.aspx.vb" Inherits="Blogwind.moods_list" title="Untitled Page" %>

<asp:Content ID="Content4" ContentPlaceHolderID="title" Runat="Server"><%=nick %> 的心情</asp:Content>

<asp:Content ID="Content1" ContentPlaceHolderID="css" Runat="Server">
    <link rel="stylesheet" href="<%=ini.GetFileWithVersion("css/mood.css") %>" type="text/css">
</asp:Content>

<asp:Content ID="Content2" ContentPlaceHolderID="content" Runat="Server">
<h3><%=nick %> 的心情</h3>
<%=MoodList %>
<br />
<div id='MoodMsg'>
    <asp:repeater id="MoodMsgList" runat="server" enableViewState="False">
		<ItemTemplate>
		<div id="white"><span class="mood"><%#"【" & Blogwind.BloggerMood.GetMoodStr(Container.DataItem("mood_id")) & "】"%></span><%#util.AddHyperLink(Server.HtmlEncode(Container.DataItem("msg")))%>
		  <span class="time"><%#Me.Helper.ShowDateInPastStyle(Container.DataItem("create_dt"))%></span>
		</div></ItemTemplate>
		<AlternatingItemTemplate>
		<div id="grey"><span class="mood"><%#"【" & Blogwind.BloggerMood.GetMoodStr(Container.DataItem("mood_id")) & "】"%></span><%# util.AddHyperLink(Server.HtmlEncode(Container.DataItem("msg")) )%>
		<span class="time"><%#Me.Helper.ShowDateInPastStyle(Container.DataItem("create_dt"))%></span>
		</div>
		</AlternatingItemTemplate>
	</asp:repeater>
</div>
<div id="fblogpaging" class="innerbar bartitle right bg0">
	<a id="pagingup" href="javascript:pageup()" style="display:none">上一页</a>
	<a id="pagingdown" href="javascript:pagedown()">下一页</a>&nbsp;&nbsp;&nbsp;&nbsp;
</div>
</asp:Content>
<asp:Content ID="Content3" ContentPlaceHolderID="js" Runat="Server">
</asp:Content>

