<%@ Page Language="vb" AutoEventWireup="false" CodeFile="fileuploader.aspx.vb" Inherits="blogwind.fileuploader"%>
<!doctype html public "-//W3C//DTD html 4.01 Transitional//EN">
<HTML>
	<HEAD>
		<title>SlickUpload Advanced Upload Sample</title>
		<link rel="stylesheet" href="fileuploader.css" type="text/css">
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<script language="javascript">
		<!--
		function pre(b)
		{
			t=b.value.substring(b.value.lastIndexOf(".")+1);
			if(t!="jpg"&&t!="JPG")
			{
				alert("只允许上传jpg格式文件");
				b.value="";
			}
			else
			window.parent.pre(b.value);
		}

		//-->
		</script>
	</HEAD>
	<body>
		<form id="Form1" runat="server">
			<div id="wrap">
			<input contentEditable="false" class="up" type="file" size="10" name="file1" id="file1" onchange="pre(this);"  runat="server"> <asp:button id="btn_upload" runat="server" Text="上传" class="b"></asp:button>
		</form>
		<script language="javascript">
		<!--
		var hint='<%=hint%>';
		if(hint=='')
		{}
		else
		{
			if(hint.indexOf("f")>0)
				alert(hint);
			else
			{
				var fname, w, h,h;
				fname=hint.substring(1,hint.indexOf('w'));
				w=hint.substring(hint.indexOf('w')+1,hint.indexOf('h'));
				h=hint.substring(hint.indexOf('h')+1)
				window.parent.uploaded(fname,w,h);
			}
		}	
			
		//-->
		</script>
	</body>
</HTML>
