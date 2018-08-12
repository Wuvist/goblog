<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.myfriends" CodeFile="myfriends.aspx.vb" %>
<asp:repeater id="friendblog" runat="server" enableViewState="False">
	<ItemTemplate><div id="white"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "index") %>.shtml' target='_blank'><%#Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "nick")) & "发表了【" & DataBinder.Eval(Container.DataItem, "title") & "】"%></a></div></ItemTemplate>
	<AlternatingItemTemplate><div id="grey"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "index") %>.shtml' target='_blank'><%#Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "nick")) & "发表了【" & DataBinder.Eval(Container.DataItem, "title") & "】"%></a></div></AlternatingItemTemplate>
</asp:repeater>
<asp:repeater id="friendcomment" runat="server" enableViewState="False">
<ItemTemplate>
<div id="white"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "article_id") %>.shtml' target='_blank'><%#app.GetEventMsg(DataBinder.Eval(Container.DataItem, "payload"), DataBinder.Eval(Container.DataItem, "title"), DataBinder.Eval(Container.DataItem, "article_id"), DataBinder.Eval(Container.DataItem, "Comment"))%></a>
</div>
</ItemTemplate>
<AlternatingItemTemplate>
<div id="grey"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "article_id") %>.shtml' target='_blank'><%#app.GetEventMsg(DataBinder.Eval(Container.DataItem, "payload"), DataBinder.Eval(Container.DataItem, "title"), DataBinder.Eval(Container.DataItem, "article_id"), DataBinder.Eval(Container.DataItem, "Comment"))%></a>
</div>
</AlternatingItemTemplate>
</asp:repeater>
<asp:Panel id="userview" runat="server" EnableViewState="False" Visible="False">
<script>
function firefox()
{
if (navigator.appName == "Microsoft Internet Explorer")
{
alert("对不起,请在Firefox浏览器下安装此插件!");
}
else
{
void(InstallTrigger.install({'人人网摘firefox插件':'http://www.renrenweb.com/download/renrenweb.xpi'}));
}
}
</script>
<table border="0" cellpadding="0" cellspacing="0" style="border-collapse: collapse" bordercolor="#111111" width="100%">
<tr>
<td height="12"></td>
</tr>
<tr>
<td width="100%" style="font-size: 16px;color=violet"><b>博客风新添加网摘（人人网摘）功能！</b>  -- <a href="http://www.renrenweb.com/help.aspx" target="_blank" style="font-size: 16px;color=red">网摘介绍</a></td>
</tr>
<tr>
<td bgcolor="#6487DC" height="2"></td>
</tr>
<tr>
<td width="100%">
<p style="line-height: 100%; font-size: 14px; margin: 15"><b>● 网摘IE插件</b><br>
&nbsp;&nbsp;&nbsp; 1. IE右键注册表插件（IE用户推荐使用）：点击下载 <a href="http://www.renrenweb.com/download/renren.reg"><font color=blue>renren.reg</font></a><br><br>
&nbsp;&nbsp;&nbsp; 如果无法直接下载reg文件则请下载：点击下载 <a href="http://www.renrenweb.com/download/renren.zip"><font color=blue>reren.zip</font></a><br><br>
&nbsp;&nbsp;&nbsp; 2. IE插件完整安装包（测试版）：点击下载 <a href="http://www.renrenweb.com/download/renrenwebsetup.msi"><font color=blue>renrenwebsetup.msi</font></a><br><br>
<b>● 网摘Firefox插件（Firefox/Mozilla/Netscape用户适用）：</b><br>&nbsp;&nbsp;&nbsp; 点击安装 <a href="javascript:firefox()"title="Firefox下安装成功后,页面上点右键选收藏此页到人人网摘,收藏网页"><font color=blue>renrenweb.xpi</font></a></p></td>
</tr>
</table>
</asp:Panel>
<asp:repeater id="friendbookmark" runat="server" enableViewState="False">
	<ItemTemplate><div id="white"><a href='http://www.renrenweb.com/item/<%# DataBinder.Eval(Container.DataItem, "lid") %>.shtml' target='_blank' title="<%# DataBinder.Eval(Container.DataItem, "nick") %> 摘录了【<%# DataBinder.Eval(Container.DataItem, "title") %>】

	<%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "des")) %>"><%# left(DataBinder.Eval(Container.DataItem, "nick") & " 摘录了【" & DataBinder.Eval(Container.DataItem, "title") & "】",48)%></a> <a href='<%# DataBinder.Eval(Container.DataItem, "url") %>' target='_blank' >#</a></div></ItemTemplate>
	<AlternatingItemTemplate><div id="grey"><a href='http://www.renrenweb.com/item/<%# DataBinder.Eval(Container.DataItem, "lid") %>.shtml' target='_blank' title="<%# DataBinder.Eval(Container.DataItem, "nick") %> 摘录了【<%# DataBinder.Eval(Container.DataItem, "title") %>】

	<%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "des")) %>"><%# left(DataBinder.Eval(Container.DataItem, "nick") & " 摘录了【" & DataBinder.Eval(Container.DataItem, "title") & "】",48)%></a> <a href='<%# DataBinder.Eval(Container.DataItem, "url") %>' target='_blank' >#</a></div></AlternatingItemTemplate>
</asp:repeater>