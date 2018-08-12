Imports SqlConnHelper

Namespace blogwind
    Partial Class pics
        Inherits AuthPage

        Protected pi As New WidgetHelper.PagingInfo
        Protected blogger As String
        Protected dt As DataTable
        Protected TotalCount As Integer

        Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
            Dim sql As String
            Dim uid As Integer = app.GetUid(ph.blogger)
            pi.Count = 20
            FormHelper.DirectBind(Request.QueryString, pi)

            sql = "select * from pic where uid=" & uid & " order by id desc"
            sql = util.GetPagingSQL(sql, pi.Start, pi.Count)
            dt = connHelper.retrieve(sql, CommandType.Text, Nothing)

            sql = "select count(*) from pic where uid=" & uid
            TotalCount = connHelper.exeInteger(sql, CommandType.Text, Nothing, False)
        End Sub
    End Class
End Namespace
