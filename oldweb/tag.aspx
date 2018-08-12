<%@ Page Language="VB" AutoEventWireup="false" CodeFile="tag.aspx.vb" Inherits="tag" %>

<%@ Register TagPrefix="uc1" TagName="bottom" Src="bottom.ascx" %>
<HTML>
	<HEAD>
		<title>博客风 标签搜索</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<link rel="icon" href="/blogwind.ico" type="image/x-icon">
		<link rel="shortcut icon" href="blogwind.ico" type="image/x-icon">
		<meta content="博客，博客风，留学" name="keywords">
		<meta content="留学生博客网站" name="description">
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
		.white { overflow:hidden;LINE-HEIGHT: 160% }
		.grey { overflow:hidden;BACKGROUND: #e5e5e5; LINE-HEIGHT: 160%;}
		</style>
	</HEAD>
	<body>
			<table width="780" align="center" cellpadding="0" cellspacing="0" bordercolor="#aaaaaa" style="border: #aaaaaa 1px solid">
				<tr>
					<td><img src="images/titlev2.jpg" width="780" height="124">
						<table width="100%" border="0" cellspacing="0" cellpadding="0">
							<tr>
								<td width="10">&nbsp;</td>
								<td>
<form method="get" action="/tag.aspx">
<div class="innerbar bartitle right bg1">标签</div>
<br style="clear:both" />
<div class="frame1 txt">
	<% If nick<>"" %>
	<input type="hidden" name="blogger" id="blogger" value="<%=blogger%>">
	<% End If %>
	<input type="text" name="tag" id="tag" value="<%=tw%>"> <input type="submit" value="搜索"><br>
	<% If nick<>"" %>
		<div><a href="/tag.aspx?tag=<%=HttpUtility.UrlEncode(tw) %>">搜索整个博客风</a></div>
		<a href="/<%=blogger%>/"><%=Nick%></a> 的网志内，带这个 <%=tw%> 标签的有：
	<% End If %>
	
</div>
</form>

<% If nick="" %>
	<br style="clear:both" />
	<div class="innerbar bartitle left bg0">使用此标签的博客</div>
	<br style="clear:both" />
	<div class="frame0 txt">
	<asp:repeater id="bloggers" runat="server" enableViewState="False">
		<ItemTemplate>
			<a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/'>
			<%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "nick")) %></a>　　
		</ItemTemplate>
	</asp:repeater>
	</div>
<% End If %>


<br style="clear:both" />
<asp:repeater id="docs" runat="server" enableViewState="False">
	<ItemTemplate>
		<div class="grey txt">
		<div>标题: <a href="/<%# DataBinder.Eval(Container.DataItem, "blogger") %>/<%# DataBinder.Eval(Container.DataItem, "index") %>.shtml"><%# DataBinder.Eval(Container.DataItem, "title") %></a></div>
		<div><a href="/<%# DataBinder.Eval(Container.DataItem, "blogger") %>/"><%# DataBinder.Eval(Container.DataItem, "nick") %></a> 发表于 <%# DataBinder.Eval(Container.DataItem, "add_date") %> 有 <%# DataBinder.Eval(Container.DataItem, "Comment") %> 个评论 </div>
		</div>
	</ItemTemplate>
	<AlternatingItemTemplate>
		<div class="white txt">
		<div>标题: <a href="/<%# DataBinder.Eval(Container.DataItem, "blogger") %>/<%# DataBinder.Eval(Container.DataItem, "index") %>.shtml"><%# DataBinder.Eval(Container.DataItem, "title") %></a></div>
		<div><a href="/<%# DataBinder.Eval(Container.DataItem, "blogger") %>/"><%# DataBinder.Eval(Container.DataItem, "nick") %></a> 发表于 <%# DataBinder.Eval(Container.DataItem, "add_date") %> 有 <%# DataBinder.Eval(Container.DataItem, "Comment") %> 个评论</div>
		</div>
	</AlternatingItemTemplate>
</asp:repeater>


<div class="innerbar bartitle right bg0"><%=paging%></div>
								</td>
								<td width="10">&nbsp;</td>
							</tr>
						</table>
							<div align="center"><font size="2">
											<uc1:bottom id="Bottom1" runat="server"></uc1:bottom></div></FONT></td>
				</tr>
			</table>
	</body>
</HTML>

