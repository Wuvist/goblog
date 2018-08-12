Imports Microsoft.VisualBasic

Namespace Blogwind
    Public Class BWPage
        Inherits System.Web.UI.Page
        Protected Helper As New UIHelper

        Protected Overrides Sub Render(ByVal writer As System.Web.UI.HtmlTextWriter)
            Dim sw As New System.IO.StringWriter()
            Dim htmlw As New HtmlTextWriter(sw)
            MyBase.Render(htmlw)
            htmlw.Flush()
            htmlw.Close()
            Dim pageContent As String = sw.ToString()
            Dim etag As String
            etag = """" & util.ToHexString(util.MD5(pageContent)) & """"
            If etag = Request.Headers("Etag") or etag = Request.Headers("If-None-Match")Then
                Response.StatusCode = 304
                Response.End()
            Else
                Response.AddHeader("Etag", etag)
                Response.Write(pageContent)
            End If
        End Sub

        Public Function GetPaging(ByVal start As Integer, ByVal count As Integer, ByVal TotalCount As Integer) As String
            Dim sb As New StringBuilder
            Dim param As NameValueCollection
            param = System.Web.HttpUtility.ParseQueryString(Me.Request.Url.Query)

            sb.Append("<div class=""paging"">")
            Dim TotalPage As Integer
            TotalPage = TotalCount \ count
            If TotalCount Mod count > 0 Then
                TotalPage += 1
            End If

            Dim pagestart As Integer
            For i As Integer = 1 To TotalPage
                pagestart = (i - 1) * count
                param.Item("start") = pagestart
                If start = pagestart Then
                    sb.Append("<span>" & i & "</span> ")
                Else
                    sb.Append("<a href='" & Me.Request.Url.AbsolutePath & "?" & param.ToString & "'>" & i & "</a> ")
                End If
            Next
            sb.Append("</div>")
            Return sb.ToString
        End Function
    End Class
End Namespace