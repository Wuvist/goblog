<%@ Page Language="VB" AutoEventWireup="false" CodeFile="MoodList.aspx.vb" Inherits="blogwind.rpc_MoodList" %>
<asp:repeater id="MoodMsgList" runat="server" enableViewState="False">
<ItemTemplate>
<div id="white"><a href="/<%# Container.DataItem("user_name") %>/"><%# Container.DataItem("nick") %></a><%# util.AddHyperLink(Server.HtmlEncode("【" & blogwind.BloggerMood.GetMoodStr(Container.DataItem("mood_id")) & "】" & Container.DataItem("msg")))%></div></ItemTemplate>
<AlternatingItemTemplate>
<div id="grey"><a href="/<%# Container.DataItem("user_name") %>/"><%# Container.DataItem("nick") %></a><%# util.AddHyperLink(Server.HtmlEncode("【" & blogwind.BloggerMood.GetMoodStr(Container.DataItem("mood_id")) & "】" & Container.DataItem("msg")))%></div>
</AlternatingItemTemplate>
</asp:repeater>