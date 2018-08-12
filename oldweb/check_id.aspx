<%
If Session("blogger") = "" Then
	Response.Write("请登陆")
	Response.End()
End If
'If Session("blogger") <> Request.QueryString("blogger") Then
'	Response.Write("登陆错误")
'	Response.End()
'End If
%>