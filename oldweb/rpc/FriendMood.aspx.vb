Imports SqlConnHelper

Namespace blogwind
    Partial Class rpc_FriendMood
        Inherits System.Web.UI.Page

        Protected Sub Page_Load(ByVal sender As Object, ByVal e As System.EventArgs) Handles Me.Load
            Dim ph As New pageHelper(Me.Session, Me.Request)
            Dim sql As String
            Dim p As Integer
            p = Request.QueryString("fbm")

            sql = "select_friendmood " & ph.bloggerID & "," & p & ",10"
            friendmood.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
            friendmood.DataBind()
        End Sub
    End Class

End Namespace
