Imports Microsoft.VisualBasic
Imports BlowfishNET

Public Class Cryptor
    Private Shared _Cryptor As New BlowfishSimple(ini.BfKey)
    Public Shared Function Decrypt(ByVal msg As String) As String
        Return _Cryptor.Decrypt(msg)
    End Function

    Public Shared Function Encrypt(ByVal msg As String) As String
        Return _Cryptor.Encrypt(msg)
    End Function
End Class
