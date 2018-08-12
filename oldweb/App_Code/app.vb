Imports Microsoft.VisualBasic
Imports SqlConnHelper

Public Class app
    Public Shared Function GetIconPath(ByVal username As String, ByVal IconType As Integer) As String
        Select Case IconType
            Case 2
                Return "/images/userpic/" & username & ".png"
            Case 1
                Return "/images/userpic/" & username & ".jpg"
            Case 3
                Return "/images/userpic/" & username & ".gif"
        End Select

        Return "/images/m.gif"
    End Function

    Public Shared Function GetUid(ByVal username As String) As Integer
        Dim pa As New ParaHelper
        pa.addNVarchar("@id", username, 50)
        Return connHelper.exeInteger("select id from users where user_name=@id", CommandType.Text, pa.parameters, False)
    End Function

    Public Shared Function GetUserId(ByVal username As String) As Integer
        Dim pa As New ParaHelper
        pa.addNVarchar("@id", username, 50)
        Return connHelper.exeInteger("select id from bloggers where user_name=@id", CommandType.Text, pa.parameters, False)
    End Function

    Public Shared Function GetEventMsg(ByVal payload As Object, ByVal title As String, ByVal ArticleId As Integer, ByVal CommentCount As String) As String
        Dim msg As String
        Dim MyPayload As String
        Dim IsNullPayload As Boolean
        Try
            MyPayload = payload
        Catch ex As System.InvalidCastException
            Dim dt As DataTable
            dt = connHelper.retrieve("select top 4  author from comment where article=" & ArticleId & " group by author order by max([index]) desc", CommandType.Text, Nothing)
            Dim ar As New ArrayList
            For Each dr As DataRow In dt.Rows
                If String.IsNullOrEmpty(dr.Item("author")) Then
                    ar.Add("Œﬁ√˚ œ")
                Else
                    ar.Add(dr.Item("author"))
                End If
            Next

            MyPayload = util.JoinStr(ar, vbCr)
            Dim pa As New ParaHelper
            pa.addNVarchar("payload", MyPayload, 500)
            connHelper.exeNonQuery("update events set payload=@payload where related_id=" & ArticleId, CommandType.Text, pa.parameters)
            IsNullPayload = True
        End Try


        Dim strs As String() = MyPayload.Split(vbCr)

        If CommentCount < 10 Then
            msg = "[ " & CommentCount & "]" & title & "  - " & util.JoinStr(strs, ", ")
        Else
            msg = "[" & CommentCount & "]" & title & "  - " & util.JoinStr(strs, ", ")
        End If

        Return msg
    End Function
End Class
