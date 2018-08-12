Imports Microsoft.VisualBasic

Public Class AuthPage
    Inherits Blogwind.BWPage
    Protected ph As pageHelper

    Protected Sub Pre_Load(ByVal sender As Object, ByVal e As System.EventArgs) Handles Me.PreLoad
        ph = New pageHelper(Me.Session, Me.Request)

        If Session("blogger") = "" Then
            Response.Write("请重新登陆")
            Response.End()
        End If
    End Sub
End Class
