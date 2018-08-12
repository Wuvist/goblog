Imports Microsoft.VisualBasic

Public Class ini
    Public Enum CommentorType
        Blogger = 1
        BloggerFans = 2
        RegisterUser = 3
        BloggerFriend = 4
        BloggerGoodFriend = 5
        UnrevealUser = 6
        Visitor = 7
    End Enum

    Public Shared ReadOnly Property PhotoBasePath() As String
        Get
            Return ConfigurationManager.AppSettings("PhotoBasePath")
        End Get
    End Property

    Public Shared ReadOnly Property EnableSmtpSSL() As Boolean
        Get
            Return System.Configuration.ConfigurationManager.AppSettings("EnableSmtpSSL")
        End Get
    End Property

    Public Shared ReadOnly Property BfKey() As String
        Get
            Return System.Configuration.ConfigurationManager.AppSettings("BfKey")
        End Get
    End Property

    Public Shared Function GetFileWithVersion(ByVal file As String) As String
        Dim fi As String
        fi = System.IO.File.GetLastWriteTime(System.Web.HttpContext.Current.Server.MapPath(file)).Ticks
        'Dim i As Integer
        'i = file.LastIndexOf(".")

        'Return ini.BaseURL & file.Substring(0, i) & "." & fi & "." & file.Substring(i + 1)

        Return "/" & file & "?v=" & fi
    End Function

    Public Shared ReadOnly Property IsProduction() As Boolean
        Get
            If ConfigurationManager.AppSettings("IsProduction") = "True" Then
                Return True
            Else
                Return False
            End If
        End Get
    End Property
End Class
