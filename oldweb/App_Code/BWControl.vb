Imports Microsoft.VisualBasic

Namespace Blogwind
    Public Interface IBindableControl
        Sub bind()
    End Interface

    Public Class BWControl
        Inherits System.Web.UI.UserControl

        Public Shared Function LoadMoControl(ByVal VirtualPath As String) As BWControl
            Dim c As New BWControl
            Return c.LoadControl(VirtualPath)
        End Function

        Public Shared Function RenderMoControl(ByVal control As BWControl) As String
            Dim sb As New StringBuilder
            Dim htw As New System.Web.UI.HtmlTextWriter(New System.IO.StringWriter(sb))
            Dim bindable As IBindableControl
            bindable = TryCast(control, IBindableControl)
            If bindable IsNot Nothing Then
                bindable.bind()
            End If

            control.RenderControl(htw)
            htw.Flush()
            Return sb.ToString
        End Function

        Public ViewData As New Hashtable
    End Class
End Namespace
