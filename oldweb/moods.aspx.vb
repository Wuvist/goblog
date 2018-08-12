Imports SqlConnHelper
Namespace Blogwind
    Partial Class moods
        Inherits BWPage
        Private MaxCount As Integer = 0
        Private MinCount As Integer = Integer.MaxValue
        Private Shared MaxSize As Integer = 20
        Private Shared MinSize As Integer = 11
        Public CurrentMood As String = ""
        Protected MoodList As String

        Protected Sub Page_Load(ByVal sender As Object, ByVal e As System.EventArgs) Handles Me.Load
            MoodList = Blogwind.BWControl.RenderMoControl(Blogwind.BWControl.LoadMoControl("/tpl/moodlist.ascx"))

            Dim m As BloggerMood.moods
            Dim sql As String

            If String.IsNullOrEmpty(Request.QueryString("mood")) Then
                sql = "select * from public_blogger_mood_v"
                sql = util.GetPagingSQL(sql, "id desc", 0, 20)
                MoodMsgList.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
            Else
                m = [Enum].Parse(GetType(BloggerMood.moods), Request.QueryString("mood"))
                CurrentMood = CInt(m)
                sql = "select * from public_blogger_mood_v where mood_id=" & CInt(m)
                sql = util.GetPagingSQL(sql, "id desc", 0, 20)
                MoodMsgList.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
            End If

            MoodMsgList.DataBind()
        End Sub

        Public Function GetMoodSize(ByVal count As Integer) As Integer
            Dim size As Integer
            size = MinSize + (count - MinCount) * (MaxSize - MinSize) / (MaxCount - MinCount)
            Return size
        End Function
    End Class
End Namespace
