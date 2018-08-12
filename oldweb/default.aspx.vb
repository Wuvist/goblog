Imports SqlConnHelper
Imports System.Text

Namespace blogwind
    Partial Class _default
        Inherits BWPage
        Protected sql, fcom As String

        Protected count_com, count, count_doc, count_blogger, MyMood As Integer
        Protected ph As pageHelper
        Private MaxCount As Integer = 0
        Private MinCount As Integer = Integer.MaxValue
        Private Shared MaxSize As Integer = 20
        Private Shared MinSize As Integer = 11

#Region " Web Form Designer Generated Code "

        'This call is required by the Web Form Designer.
        <System.Diagnostics.DebuggerStepThrough()> Private Sub InitializeComponent()

        End Sub


        Private Sub Page_Init(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Init
            'CODEGEN: This method call is required by the Web Form Designer
            'Do not modify it using the code editor.
            InitializeComponent()
        End Sub

#End Region

        Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
            'Put user code to initialize the page here

            ph = New pageHelper(Me.Session, Me.Request)
            If ph.relogin Then
                connHelper.exeNonQuery("update blogger set IP='" & Request.ServerVariables("REMOTE_ADDR") & "', Last_log=getdate() where [index]=" & ph.bloggerID, CommandType.Text, Nothing)
            End If

            If ph.login Then
                userview.Visible = True
                Dim bid As Integer
                bid = ph.bloggerID

                sql = "select_links_by_id " & bid
                links.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
                links.DataBind()

                sql = "select_friendblog " & bid & ",1"
                friendblog.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)

                sql = "select_friendmood " & bid & ",1,10"
                friendmood.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)

                sql = "select_friendevents " & bid & ",1"
                friendcomment.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)

                friendblog.DataBind()
                friendcomment.DataBind()
                friendmood.DataBind()

                sql = "select top 1 mood_id from blogger_moods where blogger_id=" & ph.bloggerID & " order by id desc"
                MyMood = connHelper.exeInteger(sql, CommandType.Text, Nothing, False)
            End If

            Recommend.DataSource = ch.getbloggerbyHits
            Recommend.DataBind()

            Newest.DataSource = ch.getNewblogger
            Newest.DataBind()

            sql = "select_new_Doc"
            new_article.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, Nothing)
            new_article.DataBind()

            sql = "select_new_Comment"
            new_comment.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, Nothing)
            new_comment.DataBind()

            sql = "select_new_Announce"
            Anounce.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, Nothing)
            Anounce.DataBind()

            sql = "select * from public_blogger_mood_v"
            sql = util.GetPagingSQL(sql, "id desc", 0, 20)
            MoodMsgList.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
            MoodMsgList.DataBind()

            Dim dt As DataTable
            dt = connHelper.retrieve("select * from moods order by mood", CommandType.Text, Nothing)
            MoodList.DataSource = dt

            For Each dr As DataRow In dt.Rows
                If MaxCount < dr.Item("blogger_count") Then
                    MaxCount = dr.Item("blogger_count")
                End If
                If MinCount > dr.Item("blogger_count") Then
                    MinCount = dr.Item("blogger_count")
                End If
            Next
            MoodList.DataBind()

            count_doc = connHelper.exeInteger("get_stat_doc", CommandType.StoredProcedure, Nothing, False)
            count_blogger = connHelper.exeInteger("get_stat_blogger", CommandType.StoredProcedure, Nothing, False)
        End Sub

        Public Function GetMoodSize(ByVal count As Integer) As Integer
            Dim size As Integer
            size = MinSize + (count - MinCount) * (MaxSize - MinSize) / (MaxCount - MinCount)
            Return size
        End Function
    End Class
End Namespace