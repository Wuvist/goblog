<%@ Reference Page="~/userpic.aspx" %>
<%@ Control Language="vb" AutoEventWireup="false" Inherits="blogwind.icon" CodeFile="icon.ascx.vb" %>
<style type="text/css"> .shadow1 {float:left;background:url(/images/shadow.gif) right bottom no-repeat}
	.shadow2 {background:url(/images/corner_bl.gif) left bottom no-repeat}
	.shadow3 {padding:0 5px 5px 0;background:url(/images/corner_tr.gif) right top no-repeat}
	.alpha {background-color: #fff; border: 1px solid #a9a9a9; padding: 4px;}
	.userpic {DISPLAY: block;CLEAR: both; float:right }
</style>
<%=userpic%>
<%=cmdlist%>
<%=formlist%><script type="text/javascript"><!--
isblogger=<%=isblogger%>;
uid = "<%=friendid%>";
 //--></script> 
 <%=widget%>