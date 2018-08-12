Imports Microsoft.VisualBasic
Namespace Blogwind.WindMill
    Public Class Account
        Enum AccountTypes
            MSNSpace = 1
            Blogger = 2
            WordPress = 3
            MetaWebLog = 4
            Twitter = 5
            Fanfou = 6
            Baidu = 100
        End Enum

        Public Shared Function GetAccountTypeName(ByVal AccountType As AccountTypes) As String
            Return AccountType.ToString
        End Function

        Public Shared Function GetAccountStatus(ByVal Status As Integer) As String
            If Status = 0 Then
                Return "OK"
            Else
                Return "Error"
            End If
        End Function
    End Class
End Namespace

