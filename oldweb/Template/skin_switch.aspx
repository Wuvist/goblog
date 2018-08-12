<!--#include file="../check_id.aspx" -->
<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.skin_switch" CodeFile="skin_switch.aspx.vb" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<HTML>
	<HEAD>
		<title>skin_switch</title>
		<link rel="stylesheet" href="/write.css" type="text/css">
	</HEAD>
<BODY BGCOLOR="FFFFFF" Marginwidth=0 marginheight=0 leftmargin=0 topmargin=0>
		<table width="100%" border="0" cellPadding="0" cellSpacing="0">
		<tr><td width="200px"><img src='/images/new/banner.jpg' width="200" height="50" border="0"></td>
		<td style="background: url(/images/new/banner_dot.jpg)" valign="middle" class="menu"><a href="/write.aspx" target="_self"><%=ph.g("编写网志")%></a> 
					<a href="/system/admin_userinfo.aspx" target="_self"><%=ph.g("个人资料")%></a>
					<a href="/system/admin_cate.aspx" target="_self"><%=ph.g("分类管理")%></a>
					<a href="/system/admin_article.aspx" target="_self"><%=ph.g("网志管理")%></a>
					<a href='/Template/skin_switch.aspx?blogger=<%=Session("blogger")%>' target="_self"><%=ph.g("选择模板")%></a>
					<a href='/system/admin_link.aspx' target="_self"><%=ph.g("链接管理")%></a>
					<a href="/" target="_top"><%=ph.g("返回首页")%></a>
					<a href='/<%=Session("blogger")%>' target="_top"><%=ph.g("返回网志")%></a>
					<a href="/logout.aspx" target="_top"><%=ph.g("退出登陆")%></a></td></tr>
		</table>
		<table width="300">
			<tr>
				<td>当前模板：</td>
			</tr>
			<tr>
				<td><img src='../images/skin<%=skin_id%>.jpg' width=300 height=190><p>不选了，跟我<a href='skin_switch.aspx?blogger=<%=blogger%>&amp;skin=0'>用随机模板</a>吧～</a></p></td>
			</tr>
		</table>
		<br>
			<table cellspacing="0" cellpadding="0" border="0" width="650">
				<%
			dim j as Integer
			dim i as Integer
			dim k as Integer
			dim r as Integer
			dim re as boolean

			r = dt.Rows.Count()
			k = 0 

			if r mod 2 = 1 then
				r = (r-1)/2				
				re = true
			else
				r = r/2
				re = false
			end if			

		for j=1 to r
			response.write("<tr>")
			for i=0 to 1
		%>
					<td>
						<TABLE cellspacing="0" cellpadding="4" rules="all" bordercolor="#3366cc" border="1" style="BORDER-RIGHT:#3366cc 1px; BORDER-TOP:#3366cc 1px; Z-INDEX:101; LEFT:32px; BORDER-LEFT:#3366cc 1px; BORDER-BOTTOM:#3366cc 1px; BORDER-COLLAPSE:collapse; BACKGROUND-COLOR:white">
							<TR>
								<TD style="WIDTH: 200px">模板名称：<%=dt.Rows(k).Item("name")%><br>
									设计者：<%=dt.Rows(k).Item("designer")%></TD>
								<TD>人气：<%=dt.Rows(k).Item("popularity")%></TD>
							</TR>
							<TR>
								<TD colspan="2">
									<P>简介：<%=dt.Rows(k).Item("intro")%></P>
									<P>缩略图：</P>
								</TD>
							</TR>
							<TR>
								<TD colspan="2"><img src='../images/skin<%=dt.Rows(k).Item("index")%>.jpg' width=300 height=190></TD>
							</TR>
							<TR>
								<TD colspan="2" align="center"><a href='skin_switch.aspx?blogger=<%=blogger%>&amp;skin=<%=dt.Rows(k).Item("index")%>'>启用</a></TD>
							</TR>
						</TABLE>
					</td>
					<%
				k = k+1
			Next%>
				</TR>
				<tr>
					<td colspan="2" height="5"></td>
				</tr>
				<%Next%>
				<%if re then%>
				<tr>
					<td colspan="2">
						<TABLE cellspacing="0" cellpadding="4" rules="all" bordercolor="#3366cc" border="1" style="BORDER-RIGHT:#3366cc 1px; BORDER-TOP:#3366cc 1px; Z-INDEX:101; LEFT:32px; BORDER-LEFT:#3366cc 1px; BORDER-BOTTOM:#3366cc 1px; BORDER-COLLAPSE:collapse; BACKGROUND-COLOR:white">
							<TR>
								<TD style="WIDTH: 150px">模板名称：<%=dt.Rows(k).Item("name")%><br>
									设计者：<%=dt.Rows(k).Item("designer")%></TD>
								<TD>人气：<%=dt.Rows(k).Item("popularity")%></TD>
							</TR>
							<TR>
								<TD colspan="2">
									<P>简介：<%=dt.Rows(k).Item("intro")%></P>
									<P>缩略图：</P>
								</TD>
							</TR>
							<TR>
								<TD colspan="2"><img src='../images/skin<%=dt.Rows(k).Item("index")%>.jpg' width=300 height=190></TD>
							</TR>
							<TR>
								<TD colspan="2" align="center"><a href='skin_switch.aspx?blogger=<%=blogger%>&amp;skin=<%=dt.Rows(k).Item("index")%>'>启用</a></TD>
							</TR>
						</TABLE>
					</td>
				</tr>
				<%end if%>
			</table>
	</body>
</HTML>
