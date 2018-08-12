Imports Microsoft.VisualBasic

Namespace Blogwind.Blog
    Public MustInherit Class BWAuthHandler
        Inherits BWHandler
        Protected ph As pageHelper

        Public NotOverridable Overrides Sub ProcessBWRequest(ByVal context As System.Web.HttpContext)
            ph = New pageHelper(context.Session, context.Request)
            If ph.bloggerID < 1 Then
                context.Response.Write("2") '2 = Not login
                Exit Sub
            End If

            ProcessAuthRequest(context)
        End Sub

        Public MustOverride Sub ProcessAuthRequest(ByVal context As System.Web.HttpContext)
    End Class
End Namespace
