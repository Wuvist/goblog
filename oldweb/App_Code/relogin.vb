Imports System.Text
Imports SqlConnHelper

Namespace blogwind
    Public Class relogin
        Inherits System.Web.UI.Page
        Public Function login(ByRef ck As HttpCookie) As Boolean
            If Not ck Is Nothing Then
                Dim blogger, password As String

                Dim result As Boolean
                result = False

                blogger = ck.Values("id")
                password = ck.Values("pass")

                If blogger <> "" And password <> "" Then
                    Dim pa As New ParaHelper
                    pa.addNVarchar("@id", blogger, 50)
                    Dim b As Byte()
                    b = connHelper.exeScalar("select pwd from bloggers where user_name=@id", CommandType.Text, pa.parameters)

                    If b Is Nothing Then
                        Return False
                    End If

                    If password = System.Text.Encoding.ASCII.GetString(b).GetHashCode.ToString Then
                        Session("blogger") = blogger
                        result = True
                    End If
                End If
                Return result
            End If
            Return False
        End Function
    End Class
End Namespace