<%@ Page Language="vb" AutoEventWireup="false" Inherits="blogwind.write" validaterequest="false" CodeFile="write.aspx.vb" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<HTML>
	<HEAD>
		<title><%=ph.g("编写网志")%></title>
		<link rel="icon" href="/favicon.ico" type="image/x-icon">
		<link rel="shortcut icon" href="/favicon.ico" type="image/x-icon">
		<link rel="stylesheet" href="write.css" type="text/css">
	</HEAD>
	<BODY onbeforeunload="RunOnBeforeUnload()">
		<textarea name='holdtext' style="DISPLAY:none"></textarea>
		<form id="edit_form" runat="server" name="edit_form" onsubmit="fsub()">
			<table width="100%" border="0" cellPadding="0" cellSpacing="0">
				<tr>
					<td width="200"><img src='/images/new/banner.jpg' width="200" height="50" border="0"></td>
					<td valign="middle" style="background: url(/images/new/banner_dot.jpg)" class="menu"><a href="/write.aspx" target="_self"><%=ph.g("编写网志")%></a> 
					<a href="/system/admin_userinfo.aspx" target="_self"><%=ph.g("个人资料")%></a>
					<a href="/system/admin_cate.aspx" target="_self"><%=ph.g("分类管理")%></a>
					<a href="/system/admin_article.aspx" target="_self"><%=ph.g("网志管理")%></a>
					<a href="/pics.aspx" target="_self"><%=ph.g("图片管理")%></a>
					<a href='/Template/skin_switch.aspx?blogger=<%=Session("blogger")%>' target="_self"><%=ph.g("选择模板")%></a>
					<a href='/system/admin_link.aspx' target="_self"><%=ph.g("链接管理")%></a>
					<a href="/" target="_top"><%=ph.g("返回首页")%></a>
					<a href='/<%=Session("blogger")%>' target="_top"><%=ph.g("返回网志")%></a>
					<a href="/logout.aspx" target="_top"><%=ph.g("退出登陆")%></a></td>
				</tr><tr height="5px"><td colspan=2></td></tr>
			</table>
			<%=ph.g("标题")%>：
			<asp:textbox id="title" runat="server" MaxLength="50" Width="280px" CssClass="title"></asp:textbox><asp:label id="message" Runat="server"></asp:label>
			<asp:Label id="b" runat="server" Visible="False"></asp:Label><BR>
			<%=ph.g("分类")%>：<asp:radiobuttonlist id="cate" runat="server" DataValueField="index" DataTextField="cate" RepeatDirection="Horizontal"
				RepeatLayout="Flow"></asp:radiobuttonlist><BR>
			<%=ph.g("标签")%>：
			<asp:TextBox id="mytags" MaxLength="50" runat="server"></asp:TextBox>(<%=ph.g("多个标签请用逗号分开，不知道什么是标签")%><a href="http://wiki.blogwind.com/doku.php?id=%E6%A0%87%E7%AD%BE" target="_blank">？</a>)
			<TABLE width="990" border="0" cellPadding="0" cellSpacing="0">
				<TR>
					<TD vAlign="top" align="left" width="780">
						<asp:TextBox id="WTB1" runat="server" TextMode="MultiLine"></asp:TextBox></TD>
					<TD vAlign="top" align="left" width="230">
						<table height="350" cellSpacing="0" cellPadding="0" width="230" border="0">
							<tr>
								<td vAlign="top" align="left" width="80">
									<P style="LINE-HEIGHT:25px" class="menu">
										<a href="#" onclick="uppic(this);"><%=ph.g("上传图片")%></a>
										<BR><a href="#" onclick="v('QQ');"><%=ph.g("QQ表情")%></a>
										<BR><a href="#" onclick="v('flag');"><%=ph.g("小旗帜")%></a>
										<BR><a href="#" onclick="v('emot');"><%=ph.g("小包子")%></a>
										<BR><a href="#" onclick="v('face');"><%=ph.g("小图片")%></a>
										<BR><a href="#" onclick="v('neko');"><%=ph.g("可爱猫咪")%></a>
										<BR><br><a href="/system/admin_backup.aspx"><%=ph.g("备份网志")%></a>
									</P>
								</td>
								<td width="150"><div class="sidediv" id="e"><div style="FONT-SIZE:11pt">提示：现在可以<a href="/system/admin_backup.aspx">备份网志</a>哦～</div>
									</div>
								</td>
							</tr>
							<tr>
								<td colspan="2"><div id="uploader" style="DISPLAY:none"><iframe frameborder="0" width="220" height="50" marginheight="0" marginwidth="0" scrolling="no"
											id="uploaderframe"></iframe></div>
								</td>
							</tr>
						</table>
						<asp:Label id="sw" runat="server"></asp:Label>
						<div class="text"><asp:button id="Button1" runat="server" Text="发表" class="b"></asp:button>
							<asp:button id="Button2" runat="server" Text="保存" class="b"></asp:button><br>
							<input type="checkbox" id="saveclip" name="saveclip" checked> <%=ph.g("提交前将网志备份在剪贴版")%>
						</div>
					</TD>
				</TR>
			</TABLE>
		</form>
<script language="javascript" type="text/javascript" src="/images/jscripts/tiny_mce3202/tiny_mce.js"></script>
<script language="javascript" type="text/javascript">
	// Notice: The simple theme does not use all options some of them are limited to the advanced theme
	tinyMCE.init({
	mode : "exact",
	elements : "WTB1",
	theme : "advanced",
	language  :"<%=ph.g("zh")%>",
	plugins : "inlinepopups,preview,paste,Music,media",
	extended_valid_elements : "img[class|src|border=0|alt|title|hspace|vspace|width|height|align|onmouseover|onmouseout|name]",
	theme_advanced_buttons1:"bold,italic,underline,strikethrough,removeformat,separator,justifyleft,justifycenter,justifyright,separator,formatselect,fontselect,fontsizeselect",
	theme_advanced_buttons2:"bullist,numlist,separator,outdent,indent,separator,forecolor,backcolor,separator,undo,redo,separator,link,unlink,separator,hr,charmap,image,media,separator,paste,pastetext,pasteword,separator,cleanup,code,separator,help",
	theme_advanced_buttons3:"",
	theme_advanced_path : false,
	theme_advanced_path_location : "bottom",
	theme_advanced_resizing : true,
	flash_wmode : "transparent",
	flash_quality : "high",
	flash_menu : "false",
	relative_urls : false,
	width : "750",
	height : "440",
	theme_advanced_toolbar_location : "top",
	theme_advanced_toolbar_align : "left"
	});

g_blnCheckUnload = true;
var piclist= "";
var word_upload_img="<%=ph.g("上传图片")%>";
var word_list_upload_img="<%=ph.g("本次上传图片")%>";
var word_upload_img_hint="<%=ph.g("请使用下方文件框选择jpg格式文件上传")%>";
var word_close_write_hint="<%=ph.g("提示：如果您选择'确定'，您将离开本页面，并且您现在所编辑的网志将不会被保存哦！")%>";
var current_emot;
function fsub()
{
	g_blnCheckUnload=false;
	disableB();
}
function pre(img)
{
	if(piclist=="")
		e.innerHTML="<div id=\"preview\" style=\"border:1;\"></div>";
	else
		e.innerHTML="<div id=\"preview\" style=\"border:1;\"></div>" + word_list_upload_img + ":<br>" + piclist;
	document.getElementById('preview').display="inline";
	document.getElementById('preview').innerHTML="您选择了如下图片准备上传：<br><img style=\"filter:progid:DXImageTransform.Microsoft.AlphaImageLoader(src='file:///"  +  img +"',sizingMethod='scale');\"src=\"file://localhost/"+ img + "\" width=\"100px\">";

}
function insertIMG(img,w,h,size)
{
	var editor1 = document.getElementById('CE_WTB1_ID');
	var tmp;
	if (size==0)
	{
		tmp='<img src="/images/img/full/' + img +  '" border=0>';
	}
	else if (size==100)
	{
		if(w>100|h>100)
		{
			tmp='<a href="/images/img/full/' + img + '" target="_blank"><img src="/images/img/100/' + img +  '" border=0></a>';
		}
		else
		{
			tmp='<img src="/images/img/full/' + img +  '" border=0>';
		}
	}
	else if (size==320)
	{
		tmp='<a href="/images/img/full/' + img + '" target="_blank"><IMG  src="/images/img/320/' + img +  '" border=0></a>';
	}
	else if (size==640)
	{
		tmp='<a href="/images/img/full/' + img + '" target="_blank"><IMG  src="/images/img/640/' + img +  '" border=0></a>';
	}
	tinyMCE.execInstanceCommand("WTB1","mceInsertContent",false,tmp);
}
function uploaded(img,w,h)
{
	var fname,pid,blist;
	pid=img.substring(1,img.indexOf('.'));
	if(w>100|h>100)
		fname="/images/img/100/" + img;
	else
		fname="/images/img/full/" + img;
	document.getElementById('preview').display="none";
	blist="";
	if(w>100|h>100)
		blist+="<a href='#' title='插入小型缩略图' onclick='insertIMG(\"" + img + "\"," + w + "," + h +",100)'>" + "小" + "</a> ";
	if(w>320|h>240)
		blist+="<a href='#' title='插入中型缩略图' onclick='insertIMG(\"" + img + "\"," + w + "," + h +",320)'>" + "中" + "</a> ";
	if(w>640|h>480)
		blist+="<a href='#' title='插入大型缩略图' onclick='insertIMG(\"" + img + "\"," + w + "," + h +",640)'>" + "大" + "</a> ";
	blist+="<a href='#' title='插入原始图片' onclick='insertIMG(\"" + img + "\"," + w + "," + h +",0)'>" + "原" + "</a>";

	piclist = "<div onmouseover=\"show_button(" + pid + ")\" onmouseout=\"hide_button(" + pid + ")\"><a href='#' title='插入小型缩略图' onclick='insertIMG(\"" + img + "\"," + w + "," + h +",100)'><img src='"+ fname + "' border=\"0\"></a><div id=\"pic" + pid + "\">" + blist + "</div></div>" + piclist;
	e.innerHTML="<div id=\"preview\" style=\"border:1;\"></div>" + word_list_upload_img + ":<br />" + piclist;
}
function show_button(pid)
{
	document.getElementById('pic' + pid).style.visibility="visible";
}
function hide_button(pid)
{
	document.getElementById('pic' + pid).style.visibility="hidden";
}
function uppic(b)
{
	if (b.innerHTML==word_upload_img)
	{
		document.getElementById("uploaderframe").src="/fileuploader.aspx";
		document.getElementById("uploader").style.display="inline";
		//b.innerHTML="本次上传图片";
		e.innerHTML="<div id=\"preview\" style=\"border:1;\">" + word_upload_img_hint + "</div>";
	}
	else
	{
		e.innerHTML="<div id=\"preview\" style=\"border:1;\"></div>" + word_list_upload_img + ":<br>" + piclist;
	}
	current_emot="";
}

function w(number) { 
	img = '<IMG  src="/images/emot/'+number+'.gif" border=0>';
	tinyMCE.execInstanceCommand("WTB1","mceInsertContent",false,img);
}

function RunOnBeforeUnload() 
{
	wuvist_t=tinyMCE.getContent("WTB1");
	var agree;
	if (g_blnCheckUnload) 
	{
		if (wuvist_t!='<DIV>&nbsp;</DIV>' && wuvist_t!='<DIV></DIV>' && wuvist_t!='')
		{
			window.event.returnValue=(word_close_write_hint);
		}
	} 
	else
	{
		if(document.getElementById("saveclip").checked==true)
			copyclip();
	}
}

function copyclip()
{
	try {
		holdtext.innerText =tinyMCE.getContent("WTB1");
		var Copied = holdtext.createTextRange();
		Copied.execCommand("Copy");
	} catch (e) {
		// Ignore any errors
	}	
}
function disableB()
{
		document.getElementById('Button1').style.visibility='hidden';
		document.getElementById('Button2').style.visibility='hidden';
}
</script>
<script type="text/javascript" src="/emot.js"></script>		
	</BODY>
</HTML>
