Imports Microsoft.VisualBasic
Imports System.Reflection

Public Class FormHelper
    Public Shared Function GetColumnName(ByVal FieldName As String) As String
        Dim result As String = ""
        Dim InUpper As Boolean = False
        Dim i As Integer
        For i = 0 To FieldName.Length - 1
            If Char.IsUpper(FieldName.Chars(i)) Then
                If InUpper Then
                    result &= Char.ToLower(FieldName.Chars(i))
                Else
                    If i = 0 Then
                        result = Char.ToLower(FieldName.Chars(i))
                    Else
                        result &= "_" & Char.ToLower(FieldName.Chars(i))
                        InUpper = True
                    End If
                End If
            Else
                result &= FieldName.Chars(i)
                InUpper = False
            End If
        Next

        Return result
    End Function

    Public Shared Sub Bind(ByVal data As System.Collections.Specialized.NameValueCollection, ByVal obj As Object)
        Dim ColumnName As String
        Dim value As String

        For Each i As FieldInfo In obj.GetType.GetFields
            ColumnName = GetColumnName(i.Name)
            value = data.Get(ColumnName)
            If Not String.IsNullOrEmpty(value) Then
                Try
                    i.SetValue(obj, value)
                Catch ex As System.ArgumentException
                    i.SetValue(obj, CInt(value))
                End Try
            End If
        Next
    End Sub

    Public Shared Sub DirectBind(ByVal data As System.Collections.Specialized.NameValueCollection, ByVal obj As Object)
        Dim value As String
        Dim values() As String
        Dim intValues() As Integer

        For Each fi As FieldInfo In obj.GetType.GetFields
            value = data.Get(fi.Name)
            If Not String.IsNullOrEmpty(value) Then
                If fi.FieldType Is GetType(Integer) Then
                    fi.SetValue(obj, CInt(value))
                ElseIf fi.FieldType Is GetType(String()) Then
                    values = data.GetValues(fi.Name)
                    fi.SetValue(obj, values)
                ElseIf fi.FieldType Is GetType(Integer()) Then
                    values = data.GetValues(fi.Name)
                    ReDim intValues(values.Length - 1)
                    For i As Integer = 0 To values.Length - 1
                        intValues(i) = CInt(values(i))
                    Next
                    fi.SetValue(obj, intValues)
                ElseIf fi.FieldType Is GetType(Boolean) Then
                    If value = "on" Then
                        fi.SetValue(obj, True)
                    Else
                        fi.SetValue(obj, value)
                    End If
                Else
                    fi.SetValue(obj, value)
                End If
            End If
        Next
    End Sub

    Public Shared Function RenderAsTable(ByVal obj As Object) As String
        Dim s As New System.Text.StringBuilder
        Dim att As System.Xml.Serialization.XmlAttributeAttribute()

        For Each i As FieldInfo In obj.GetType.GetFields
            att = i.GetCustomAttributes(GetType(System.Xml.Serialization.XmlAttributeAttribute), False)
            If att.Length > 0 Then
                s.Append("<th>" & i.Name & att(0).AttributeName & " :</th><td></td>")
            Else
                s.Append("<th>" & i.Name & " :</th><td></td>")
            End If

        Next
        Return s.ToString
    End Function
End Class