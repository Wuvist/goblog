Imports Microsoft.VisualBasic

Public Class WidgetHelper
    Public Class PagingInfo
        Public Start As Integer
        Public Count As Integer
    End Class

    Public Shared Function GetIntArrayFromString(ByVal values As String) As Integer()
        Dim data As New System.Collections.Generic.List(Of Integer)
        If Not String.IsNullOrEmpty(values) Then
            For Each value As String In values.Split(",")
                data.Add(CInt(value))
            Next
        End If
        Return data.ToArray
    End Function

    Public Shared Function GetRadios(ByVal Name As String, ByVal enumType As Type, ByVal SelectedValue As Nullable(Of Integer)) As String
        If SelectedValue.HasValue Then
            Return GetRadios(Name, enumType, SelectedValue.Value)
        Else
            Return GetRadios(Name, enumType, -1)
        End If
    End Function


    Public Shared Function GetCheckboxes(ByVal Name As String, ByVal data As DictionaryEntry(), ByVal SelectedValues As String) As String
        Dim MyUIHelper As New UIHelper
        Dim values As Integer() = GetIntArrayFromString(SelectedValues)


        Dim sb As New StringBuilder
        For Each de As DictionaryEntry In data
            sb.Append("<input type=""checkbox"" class=""checkbox"" name='" & Name & "' value='")
            sb.Append(de.Key)
            sb.Append("'")

            If Array.IndexOf(values, CInt(de.Key)) > -1 Then
                sb.Append(" checked")
            End If
            sb.Append(" /> ")
            sb.Append(MyUIHelper.g(de.Value))
            sb.Append("<br />")
        Next
        Return sb.ToString
    End Function

    Public Shared Function GetSelectOptions(ByVal data As DictionaryEntry(), ByVal SelectedValue As String) As String
        Dim MyUIHelper As New UIHelper
        Dim sb As New StringBuilder
        For Each de As DictionaryEntry In data
            sb.Append("<option value='")
            sb.Append(de.Key)
            sb.Append("'")
            If de.Key = SelectedValue Then
                sb.Append(" selected")
            End If
            sb.Append(">")
            sb.Append(MyUIHelper.g(de.Value))
            sb.Append("</option>")
        Next
        Return sb.ToString
    End Function

    Public Shared Function GetSelectOptions(ByVal enumType As Type, ByVal SelectedValue As Nullable(Of Integer)) As String
        If SelectedValue.HasValue Then
            Return GetSelectOptions(enumType, SelectedValue.Value)
        Else
            Return GetSelectOptions(enumType, -1)
        End If
    End Function
End Class
