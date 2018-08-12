<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.selectskin" CodeFile="selectskin.aspx.vb" %>
<%@ Register TagPrefix="uc1" TagName="bottom" Src="bottom.ascx" %>
<HTML>
	<HEAD>
		<title>博客风－注册博客－选取模板</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<LINK href="/blogwind.ico" type="image/x-icon" rel="icon">
			<LINK href="blogwind.ico" type="image/x-icon" rel="shortcut icon">
				<meta content="博客，博客风，留学" name="keywords">
				<meta content="留学生博客网站" name="description">
				<STYLE type="text/css">
		.loginput { BORDER-RIGHT: #b6caff 1px solid; BORDER-TOP: #b6caff 1px solid; FONT-SIZE: 12px; BORDER-LEFT: #b6caff 1px solid; WIDTH: 250px; LINE-HEIGHT: 15px; BORDER-BOTTOM: #b6caff 1px solid }
		.logbutton { BORDER-RIGHT: #aaaaaa 1px solid; BORDER-TOP: #aaaaaa 1px solid; FONT-SIZE: 12px; BORDER-LEFT: #aaaaaa 1px solid; LINE-HEIGHT: 15px; BORDER-BOTTOM: #aaaaaa 1px solid; FONT-FAMILY: 宋体; BACKGROUND-COLOR: #d9e3ff }
		.t { BORDER-RIGHT: #3366cc 1px; BORDER-TOP: #3366cc 1px; BORDER-LEFT: #3366cc 1px; BORDER-BOTTOM: #3366cc 1px; BORDER-COLLAPSE: collapse; BACKGROUND-COLOR: white; align: left }
		</STYLE>
	</HEAD>
	<body>
		<form id="Form1" method="post" runat="server">
			<table borderColor="#aaaaaa" cellSpacing="0" cellPadding="0" width="780" align="center"
				border="1">
				<tr>
					<td><IMG height="123" src="images/titlev2.jpg" width="778">
						<table cellSpacing="0" cellPadding="0" width="100%" border="0">
							<tr>
								<td width="20">&nbsp;</td>
								<td>
									<P><%=blogger%>, 请为您的博客选取一款模板：</P>										
										<table cellspacing="0" cellpadding="0" border="0" width="760">
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
													<TABLE cellspacing="0" cellpadding="4" rules="all" bordercolor="#3366cc" border="1" class="t">
														<TR>
															<TD width="200">模板名称：<%=dt.Rows(k).Item("name")%><br>
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
															<TD colspan="2"><img src='/images/skin<%=dt.Rows(k).Item("index")%>.jpg' width=300 height=190></TD>
														</TR>
														<TR>
															<TD colspan="2" align="center"><a href='selectskin.aspx?blogger=<%=blogger%>&amp;skin=<%=dt.Rows(k).Item("index")%>'>使用</a></TD>
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
													<TABLE cellspacing="0" cellpadding="4" rules="all" bordercolor="#3366cc" border="1" class="t">
														<TR>
															<TD width="150">模板名称：<%=dt.Rows(k).Item("name")%><br>
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
															<TD colspan="2" align="center"><a href='selectskin.aspx?blogger=<%=blogger%>&amp;skin=<%=dt.Rows(k).Item("index")%>'>使用</a></TD>
														</TR>
													</TABLE>
												</td>
											</tr>
											<%end if%>
										</table>
								</td>
							</tr>
						</table>
						<div>
						</div>
						<DIV><FONT size="2">
								<DIV align="center">
									<uc1:bottom id="Bottom1" runat="server"></uc1:bottom></DIV>
						</DIV>
						</FONT>
					</td>
				</tr>
			</table>
		</form>
	</body>
</HTML>
