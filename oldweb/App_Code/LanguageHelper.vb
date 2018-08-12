Imports GNU.Gettext

Public Class LanguageHelper
    Public Shared catalog As GettextResourceManager

    Public Shared Sub initcate(ByVal dictpath As String)
        catalog = New GettextResourceManager("blogwind", dictpath)
    End Sub

    Public Shared Function GetDisplayLanguage(ByVal context As HttpContext) As langs.lang
        Dim ck As HttpCookie
        ck = context.Request.Cookies("L")
        If ck IsNot Nothing Then
            Return langs.GetLang(ck.Value, context.Request.UserLanguages)
        End If
    End Function
End Class
