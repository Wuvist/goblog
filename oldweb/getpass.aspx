<%@ Reference Page="~/email.aspx" %>
<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.getpass" CodeFile="getpass.aspx.vb" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<HTML>
	<HEAD>
		<title>取回密码</title>
		<meta content="Microsoft Visual Studio .NET 7.1" name="GENERATOR">
		<meta content="Visual Basic .NET 7.1" name="CODE_LANGUAGE">
		<meta content="JavaScript" name="vs_defaultClientScript">
		<meta content="http://schemas.microsoft.com/intellisense/ie5" name="vs_targetSchema">
		<LINK href="frame.css" type="text/css" rel="stylesheet">
	</HEAD>
	<BODY>
		&gt;
		<script type="text/javascript">
		//<![CDATA[
		var newsoldonloadHndlr=window.onload, newspopupHgt, newsactualHgt, newstmrId=-1, newsresetTimer;
        var newstitHgt, newscntDelta, newstmrHide=-1, newshideAfter=10000, newshideAlpha, newshasFilters=true;
        var newsnWin, newsshowBy=null, newsdxTimer=-1, newspopupBottom;

        var newsnText,newsnMsg,newsnTitle,newsbChangeTexts=false;
		function newsespopup_Close()
        {
          if (newstmrId==-1)
          {
            el=document.getElementById('news');
            el.style.filter='';
            el.style.display='none';
            if (newstmrHide!=-1) clearInterval(newstmrHide); newstmrHide=-1;
            
          }
		  parent.helper.style.visibility='hidden';
        } //]]>
		</script>
		<form id="Form1" method="post" runat="server">
			<div id="news" style="BORDER-RIGHT: #455690 1px solid; BORDER-TOP: #b9c9ef 1px solid; Z-INDEX: 9999; BACKGROUND: #e0e9f8; LEFT: 0px; BORDER-LEFT: #b9c9ef 1px solid; WIDTH: 200px; BORDER-BOTTOM: #455690 1px solid; POSITION: absolute; TOP: 0px; HEIGHT: 120px">
				<div id="news_header" style="FILTER: progid:DXImageTransform.Microsoft.Gradient(GradientType=0,
                 StartColorStr='#FFE0E9F8', EndColorStr='#FFFFFFFF'); LEFT: 2px; FONT: 12px arial,sans-serif; WIDTH: 194px; CURSOR: default; COLOR: #1f336b; POSITION: absolute; TOP: 2px; HEIGHT: 14px; TEXT-DECORATION: none"><span id="newstitleEl">博客风：取回密码</span><span onmousedown="event.cancelBubble=true;" onmouseover="style.color='#455690';" style="RIGHT: 3px; FONT: bold 12px arial,sans-serif; CURSOR: pointer; COLOR: #728eb8; POSITION: absolute; TOP: 0px"
						onclick="newsespopup_Close()" onmouseout="style.color='#728EB8';">X</span></div>
				<div onmousedown="event.cancelBubble=true;" id="news_content" style="BORDER-RIGHT: #b9c9ef 1px solid; PADDING-RIGHT: 2px; BORDER-TOP: #728eb8 1px solid; PADDING-LEFT: 2px; BACKGROUND: #e0e9f8; FILTER: progid:DXImageTransform.Microsoft.Gradient(GradientType=0,
                 StartColorStr='#FFE0E9F8', EndColorStr='#FFFFFFFF'); LEFT: 2px; PADDING-BOTTOM: 2px; OVERFLOW: hidden; BORDER-LEFT: #728eb8 1px solid; WIDTH: 194px; PADDING-TOP: 2px; BORDER-BOTTOM: #b9c9ef 1px solid; POSITION: absolute; TOP: 20px; HEIGHT: 96px; TEXT-ALIGN: center">
					<div id="newsaCnt" style="FONT: 12px arial,sans-serif; COLOR: #1f336b; TEXT-DECORATION: none"><form target="_blank" method="post" action="system/check.aspx" name="log">
							请输入您的帐号信息：<BR>
							用户名： <INPUT class="input" id="ids" type="text" size="15" name="ids" runat="server"><br>
							电&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;邮：&nbsp;&nbsp;<INPUT class="input" id="ems" type="text" size="15" name="ems" runat="server">&nbsp;
							<table height="6" border="0">
								<tr>
									<td><FONT face="宋体"></FONT></td>
								</tr>
							</table>
							<INPUT class="button" id="Submit1" onclick="javascript:window.open('/','_self')" type="button"
								value="返回" name="返回">&nbsp;&nbsp;&nbsp;
							<asp:button id="retrive" runat="server" Text="取回密码"></asp:button></form>
						</FONT></TD></TR></TABLE></div>
				</div>
			</div>
		</form>
	</BODY>
</HTML>
