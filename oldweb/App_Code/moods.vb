Imports SqlConnHelper
Imports Microsoft.VisualBasic

Namespace blogwind
    Public Class BloggerMood
        Public Mood As moods
        Public MoodMsg As String
        Public Id As Integer
        Public BloggerId As Integer
        Public MoodDate As DateTime

        Public Shared Function GetMoodStr(ByVal mood As Integer) As String
            Dim m As moods
            m = mood
            Return m.ToString
        End Function

        Public Enum moods
            饿
            离
            哭
            食
            爽
            读
            拼
            泣
            闷
            羞
            暑
            烦
            哀
            拽
            囧
            穷
            睡
            空
            乐
            困
            闲
            恼
            懒
            愤
            冷
            湿
            贺
            思
            怒
            聊
            厌
            忙
            吵
            泪
            热
            病
            累
            乱
            吓
            赞
            美
            喜
            怨
            憋
            渴
            伤
            痛
            幽
            静
            恨
            凶
            爱
        End Enum

        Public Shared Function GetFriendMoods(ByVal bloggerId As Integer, ByVal PageNumber As Integer, ByVal PageSize As Integer) As DataTable
            Dim pa As New ParaHelper
            pa.addInt("bloggerid", bloggerId)
            pa.addInt("page", PageNumber)
            pa.addInt("pagesize", PageSize)
            Return connHelper.retrieve("select_friendmood", CommandType.StoredProcedure, pa.parameters)
        End Function

        Public Shared Function MapDr(ByVal dr As DataRow) As BloggerMood
            If dr Is Nothing Then
                Return Nothing
            End If
            Dim m As New BloggerMood
            m.Mood = dr.Item("mood_id")
            m.MoodMsg = dr.Item("msg")
            m.MoodDate = dr.Item("create_dt")
            m.Id = dr.Item("id")
            Return m
        End Function

        Public Shared Function GetBloggerCureentMood(ByVal bloggerId As Integer) As BloggerMood
            Dim m As New BloggerMood
            Dim dr As DataRow

            Dim pa As New ParaHelper
            pa.addInt("bloggerid", bloggerId)
            dr = connHelper.exeDataRow("select top 1 * from blogger_moods where blogger_id=@bloggerid order by id desc", CommandType.Text, pa.parameters)
            Return BloggerMood.MapDr(dr)
        End Function

        Public Sub Save()
            Dim pa As New ParaHelper
            pa.addInt("bloggerid", Me.BloggerId)
            pa.addInt("mood", Me.Mood)
            pa.addNVarchar("msg", Me.MoodMsg, 500)
            Dim sql As String
            sql = "insert into blogger_moods (blogger_id,mood_id,msg)values(@bloggerid, @mood, @msg)"
            connHelper.exeNonQuery(sql, CommandType.Text, pa.parameters)
            sql = "update moods set blogger_count=blogger_count+1 where id=" & CInt(Me.Mood).ToString
            connHelper.exeNonQuery(sql, CommandType.Text, Nothing)
        End Sub

    End Class
End Namespace

