Imports sqlconnhelper

Public Class Blogger
    Public BloggerID As Integer
    Public UserName As String

    Public Shared Function Auth(ByVal username As String, ByVal password As String) As Integer
        Dim pa As New ParaHelper
        Dim BloggerID As Integer


        pa.addVarchar("user_name", username, 50)
        pa.addBinary("@pwd", util.MD5(password))

        BloggerID = connHelper.exeInteger("select id from bloggers where user_name=@user_name and pwd=@pwd", CommandType.Text, pa.parameters, False)
        If BloggerID = 0 Then
            Throw New ApplicationException("Auth Fail")
        Else
            Return BloggerID
        End If
    End Function

    Public Shared Function GetBlogName(ByVal BloggerId As Integer) As String
        Dim sql As String
        sql = "SELECT blogname FROM Bloggers WHERE id=" & BloggerId
        Return connHelper.exeString(sql, CommandType.Text, Nothing, Nothing)
    End Function

    Public Shared Function GetBloggerBubbleStyle(ByVal BloggerId As Integer) As String
        Dim color As String
        color = System.Web.HttpContext.Current.Items("BloggerBubbleStyle" & BloggerId)

        If String.IsNullOrEmpty(color) Then
            Dim sql As String
            sql = "SELECT blogs FROM Bloggers WHERE id=" & BloggerId
            color = connHelper.exeInteger(sql, CommandType.Text, Nothing, Nothing) Mod 7 + 1
            System.Web.HttpContext.Current.Items("BloggerBubbleStyle" & BloggerId) = color
        End If
        Return color
    End Function
End Class
