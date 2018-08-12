<%@ Page Language="vb" AutoEventWireup="false" CodeFile="friendlist.aspx.vb" Inherits="blogwind.friendlist"%>
<HTML>
	<HEAD>
		<title>博客风－朋友列表</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<SCRIPT LANGUAGE="JavaScript">
function shows(link) {
	if(link=='')
	{
		alert("不好意思，对方选择隐藏自己的网志。");
	}
	else
	{
		this.location="friendlist.aspx?blogger=" + link;
	}
}

function  opens(link) {
	if(link=='')
	{
		alert("不好意思，对方选择隐藏自己的网志。");
	}
	else
	{
		window.open('/' + link,'_blank');
	}
}
</SCRIPT>
<meta content="博客，博客风，留学" name="keywords">
<meta content="留学生博客网站" name="description">
<style type="text/css">
<!--
body{
MARGIN: 5px 5px 5px 5px;
PADDING: 0px; 
BORDER: 0px;
COLOR: #333;
FONT-SIZE:12px;
LINE-HEIGHT:150%;
FONT-FAMILY: arial,verdana,sans-serif,'宋体';
FONT-SIZE: 12px;
COLOR: #3E62DB;
background:	#f7f7f7;
SCROLLBAR-FACE-COLOR: #85A2E0; 
SCROLLBAR-HIGHLIGHT-COLOR: #ffffff; 
SCROLLBAR-SHADOW-COLOR: #ffffff; 
SCROLLBAR-3DLIGHT-COLOR: #ffffff; 
SCROLLBAR-ARROW-COLOR: #FFFF00; 
SCROLLBAR-TRACK-COLOR: #DAD9EA; 
SCROLLBAR-DARKSHADOW-COLOR: #DAD9EA;
}
A:link { COLOR: #3e62db; TEXT-DECORATION: none }
A:active { COLOR: #3e62db; TEXT-DECORATION: none }
A:visited { COLOR: #3e62db; TEXT-DECORATION: none }
A:hover { COLOR: red; TEXT-DECORATION: underline }
-->
</style>
	</HEAD>
	<body>
		<P><%=nick%>的朋友：
			<asp:Repeater id="friends" runat="server" EnableViewState=False><ItemTemplate><a href="#" oncontextmenu="shows('<%# DataBinder.Eval(Container.DataItem, "id") %>');" onclick="opens('<%# DataBinder.Eval(Container.DataItem, "id") %>');"><%# DataBinder.Eval(Container.DataItem, "nick") %>【<%# DataBinder.Eval(Container.DataItem, "blogname") %>】</a>　　</ItemTemplate></asp:Repeater></P>
		<P>添加<%=nick%>为朋友的人：<asp:Repeater id="flist" runat="server" EnableViewState=False><ItemTemplate><a href="#" oncontextmenu="shows('<%# DataBinder.Eval(Container.DataItem, "id") %>');" onclick="opens('<%# DataBinder.Eval(Container.DataItem, "id") %>');"><%# DataBinder.Eval(Container.DataItem, "nick") %>【<%# DataBinder.Eval(Container.DataItem, "blogname") %>】</a>　　</ItemTemplate></asp:Repeater></P>
		<P> <INPUT style="BORDER-RIGHT: #aaaaaa 1px solid; BORDER-TOP: #aaaaaa 1px solid; FONT-SIZE: 9pt; BORDER-LEFT: #aaaaaa 1px solid; LINE-HEIGHT: 14px; BORDER-BOTTOM: #aaaaaa 1px solid; FONT-FAMILY: 宋体; BACKGROUND-COLOR: #665582;color:white" type="button" value="关闭" onclick="window.parent.hides();"></P>
	</body>
</HTML>
