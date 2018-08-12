Imports Microsoft.VisualBasic

Public Class Twitter
    Private Shared ReadOnly Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)

    Public Class Urls
        Public Shared Twitter As String = "http://twitter.com/statuses/update.xml"
        Public Shared Fanfou As String = "http://api.fanfou.com/statuses/update.xml"
        Public Shared FanfouAuth As String = "http://api.fanfou.com/users/show.xml"
        Public Shared TwitterAuth As String = "http://twitter.com/users/show/wuvist.xml"
    End Class

    Public Shared Function Validate(ByVal Username As String, ByVal Password As String, ByVal Url As String) As Boolean
        Dim c As New System.Net.WebClient
        c.Credentials = New System.Net.NetworkCredential(Username, Password)
        Dim data As New System.Collections.Specialized.NameValueCollection
        Try
            c.UploadValues(Url, data)
        Catch ex As Exception
            Return False
        End Try
        Return True
    End Function

    Public Shared Sub Post(ByVal Username As String, ByVal Password As String, ByVal Url As String, ByVal msg As String)
        Dim c As New System.Net.WebClient
        c.Credentials = New System.Net.NetworkCredential(Username, Password)
        Dim data As New System.Collections.Specialized.NameValueCollection
        data.Add("status", msg)
        Dim str As String = System.Text.Encoding.UTF8.GetString(c.UploadValues(Url, data))
        Log.Debug(str)
    End Sub

End Class
