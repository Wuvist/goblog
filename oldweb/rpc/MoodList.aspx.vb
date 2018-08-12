Imports SqlConnHelper
imports Newtonsoft

Namespace blogwind
    Partial Class rpc_MoodList
        Inherits System.Web.UI.Page
        Private PageSize As Integer = 20

        Protected Sub Page_Load(ByVal sender As Object, ByVal e As System.EventArgs) Handles Me.Load
            Dim sql As String
            Dim start As Integer
            Dim m As BloggerMood.moods
            Dim PageNumber As Integer
            PageNumber = Request.QueryString("page")
            If Not String.IsNullOrEmpty(Request.QueryString("size")) Then
                PageSize = Request.QueryString("size")
            End If

            start = (PageNumber - 1) * PageSize

            If String.IsNullOrEmpty(Request.QueryString("mood")) Then
                sql = "select * from public_blogger_mood_v"
                sql = util.GetPagingSQL(sql, "id desc", start, PageSize)
                MoodMsgList.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
            Else
                m = Request.QueryString("mood")
                sql = "select * from public_blogger_mood_v where mood_id=" & CInt(m)
                sql = util.GetPagingSQL(sql, "id desc", start, PageSize)
                MoodMsgList.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
            End If
            MoodMsgList.DataBind()
        End Sub
    End Class
End Namespace
