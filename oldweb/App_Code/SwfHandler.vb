Imports Microsoft.VisualBasic
Namespace Blogwind.Blog
    Public MustInherit Class SwfHandler
        Implements IHttpHandler, IRequiresSessionState

        Protected ph As pageHelper

        Public ReadOnly Property IsReusable() As Boolean Implements System.Web.IHttpHandler.IsReusable
            Get
                Return False
            End Get
        End Property


        Public Sub ProcessRequest(ByVal context As System.Web.HttpContext) Implements System.Web.IHttpHandler.ProcessRequest
            ph = New pageHelper(context)

            If context.Session("blogger") = "" Then
                context.Response.Write("ÇëÖØÐÂµÇÂ½")
                context.Response.End()
            End If

            ProcessAuthRequest(context)
        End Sub
        Public MustOverride Sub ProcessAuthRequest(ByVal context As System.Web.HttpContext)
    End Class

End Namespace
