Imports Microsoft.VisualBasic

Public Class UIHelper
    Private cu As System.Globalization.CultureInfo

    Private _PageDisplayLanguage As langs.lang
    Public Property PageDisplayLanguage() As langs.lang
        Get
            Return _PageDisplayLanguage
        End Get
        Set(ByVal value As langs.lang)
            _PageDisplayLanguage = value
            cu = New System.Globalization.CultureInfo(langs.GetLangString(value))
        End Set
    End Property

    Public Sub New(ByVal context As HttpContext)
        Me._PageDisplayLanguage = LanguageHelper.GetDisplayLanguage(context)
    End Sub

    Public Sub New()
        Me._PageDisplayLanguage = LanguageHelper.GetDisplayLanguage(HttpContext.Current)
    End Sub

    Public Function g(ByVal msg As String) As String
        If cu Is Nothing Then
            cu = New System.Globalization.CultureInfo(langs.GetLangString(Me.PageDisplayLanguage))
        End If

        Return LanguageHelper.catalog.GetString(msg, cu)
    End Function

    Public Function ShowFromSource(ByVal FromSource As Byte)
        'Web = 0
        'Mobile = 1
        If FromSource = 0 Then
            Return Me.g("web")
        End If

        If FromSource = 1 Then
            Return Me.g("mobile phone")
        End If

        Return Me.g("web")
    End Function


    Public Function ShowDateInPastStyle(ByVal DateTime As Date)
        Dim d As TimeSpan
        d = Now.Subtract(DateTime)
        Dim abs As Date
        If DateTime = abs Then
            Return Me.g("Never")
        End If
        If d.Days > 1 Then
            Return String.Format(Me.g("{0} days ago"), d.Days)
        End If

        If d.Days = 1 Then
            Return Me.g("Yesterday")
        End If

        If d.Hours > 1 Then
            Return String.Format(Me.g("{0} hours ago"), d.Hours, d.Minutes)
        End If

        If d.Hours = 1 Then
            Return String.Format(Me.g("{0} hour ago"), d.Hours, d.Minutes)
        End If

        If d.Minutes > 1 Then
            Return String.Format(Me.g("{0} minutes ago"), d.Minutes)
        End If

        If d.Minutes = 1 Then
            Return String.Format(Me.g("1 minute ago"), d.Minutes)
        End If

        Return Me.g("less than 1 minute ago")
    End Function
End Class
