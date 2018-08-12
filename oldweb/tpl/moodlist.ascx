<%@ Control Language="VB" AutoEventWireup="false" CodeFile="moodlist.ascx.vb" Inherits="moodlist" %>
<div class="frame1 txt">
<a href="/moods.aspx">心情</a>：
<%  For Each dr As DataRow In dt.Rows%>
<span style="font-size:<%=GetMoodSize(dr.Item("blogger_count")) %>pt">
<a href="/moods.aspx?mood=<%=System.Web.HttpUtility.UrlEncode(Blogwind.BloggerMood.GetMoodStr(dr.Item("id"))) %>"><%=Blogwind.BloggerMood.GetMoodStr(dr.Item("id"))%></a></span>
<%Next %>	
</div>