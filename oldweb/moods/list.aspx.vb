Imports SqlConnHelper

Namespace Blogwind
    Partial Class moods_list
        Inherits Blogwind.BWPage
        Protected MoodList As String
        Protected blogger, nick As String

        Protected Sub Page_Load(ByVal sender As Object, ByVal e As System.EventArgs) Handles Me.Load
            blogger = Request("blogger")
            MoodList = Blogwind.BWControl.RenderMoControl(Blogwind.BWControl.LoadMoControl("/tpl/moodlist.ascx"))

            Dim pa As New ParaHelper
            pa.addVarchar("@id", blogger, 50)
            Dim t As Boolean = False
            Dim dr As DataRow
            dr = connHelper.exeDataRow("select [index], nick from blogger where id=@id", CommandType.Text, pa.parameters)
            If dr Is Nothing Then
                Response.StatusCode = 404
                Server.Transfer("/404.html")
            End If

            If dr.IsNull(1) Then
                nick = blogger
            Else
                nick = dr.Item(1)
            End If

            Dim sql As String

            sql = "select * from blogger_moods where blogger_id=" & dr.Item(0)
            sql = util.GetPagingSQL(sql, "id desc", 0, 20)
            MoodMsgList.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
            MoodMsgList.DataBind()
        End Sub
    End Class
End Namespace