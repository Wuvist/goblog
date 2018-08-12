<%@ Page Language="VB" MasterPageFile="~/master/base.master" AutoEventWireup="false" CodeFile="moods.aspx.vb" Inherits="Blogwind.moods" %>
<asp:Content ID="Content4" ContentPlaceHolderID="title" Runat="Server">博客心情 - 博客风</asp:Content>
<asp:Content ID="Content1" ContentPlaceHolderID="css" Runat="Server">
	<style>
	.frame1{border:1px solid #5da3fc;padding:2px;}
	.frame0{border:1px solid #addf53;padding:2px;}
	.frame2{border:1px solid #cc90d9;padding:2px;}
	body{ FONT-SIZE: 10.5pt; FONT-FAMILY: "微软雅黑",simsun,"宋体"; }
	.bg1{background-color:#5da3fc}
	.bg0{background-color:#addf53}
	.bg2{background-color:#cc90d9}
	.inline{display:inline}
	.txt {padding:3px;font-size:10.5pt}
	.left{float:left}
	.right{float:right}
	.innerbar{padding: 3px 0px 0px 2px; color:#333333; font-weight:bold;font-size:9pt;}
	.bartitle{font-size:9pt;padding:4px;color:#fff}
	#white { overflow:hidden;LINE-HEIGHT: 160% }
	#grey { overflow:hidden;BACKGROUND: #e5e5e5; LINE-HEIGHT: 160%;}
	</style>
</asp:Content>

<asp:Content ID="Content2" ContentPlaceHolderID="content" Runat="Server">
<%=MoodList %>
	<br>
	<div id='MoodMsg'>
    <asp:repeater id="MoodMsgList" runat="server" enableViewState="False">
		<ItemTemplate>
		<div id="white"><a href="/<%# Container.DataItem("user_name") %>/"><%# Container.DataItem("nick") %></a><%#util.AddHyperLink(Server.HtmlEncode("【" & blogwind.BloggerMood.GetMoodStr(Container.DataItem("mood_id")) & "】" & Container.DataItem("msg")))%></div></ItemTemplate>
		<AlternatingItemTemplate>
		<div id="grey"><a href="/<%# Container.DataItem("user_name") %>/"><%# Container.DataItem("nick") %></a><%# util.AddHyperLink(Server.HtmlEncode("【" & Blogwind.BloggerMood.GetMoodStr(Container.DataItem("mood_id")) & "】" & Container.DataItem("msg")) )%></div>
		</AlternatingItemTemplate>
	</asp:repeater>
    </div>
	<div id="fblogpaging" class="innerbar bartitle right bg0">
		<a id="pagingup" href="javascript:pageup()" style="display:none">上一页</a>
		<a id="pagingdown" href="javascript:pagedown()">下一页</a>&nbsp;&nbsp;&nbsp;&nbsp;</div>
</asp:Content>

<asp:Content ID="Content3" ContentPlaceHolderID="js" Runat="Server">
<script>
var page=1;
var mood="<%=CurrentMood%>";
function pageup()
{
	page --;
	if (page<2)
		$("#pagingup").hide();
	GetPage();
}

function GetPage()
{
	$("#MoodMsg").html("读取中...");
	$.get("rpc/MoodList.aspx",{
	mood:mood,
	page:page
	  },function(xml){
  		if (!xml)
			xml="无资料";
		$('#MoodMsg').html(xml);
	});
}
function pagedown()
{
	page++;
	$("#pagingup").show();
	GetPage();
}

</script>
</asp:Content>