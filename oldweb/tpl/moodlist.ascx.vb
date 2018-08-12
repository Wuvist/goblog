Imports SqlConnHelper
Partial Class moodlist
    Inherits Blogwind.BWControl
    Implements Blogwind.IBindableControl

    Private MaxCount As Integer = 0
    Private MinCount As Integer = Integer.MaxValue
    Private Shared MaxSize As Integer = 20
    Private Shared MinSize As Integer = 11
    Protected dt As DataTable

    Public Sub bind() Implements Blogwind.IBindableControl.bind
        dt = connHelper.retrieve("select * from moods", CommandType.Text, Nothing)
        For Each dr As DataRow In dt.Rows
            If MaxCount < dr.Item("blogger_count") Then
                MaxCount = dr.Item("blogger_count")
            End If
            If MinCount > dr.Item("blogger_count") Then
                MinCount = dr.Item("blogger_count")
            End If
        Next
    End Sub

    Public Function GetMoodSize(ByVal count As Integer) As Integer
        Dim size As Integer
        size = MinSize + (count - MinCount) * (MaxSize - MinSize) / (MaxCount - MinCount)
        Return size
    End Function
End Class
