Imports SqlConnHelper

Partial Class admin_article
    Inherits System.Web.UI.Page
    Protected ph As pageHelper

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

        If Not IsPostBack Then
            BindArticlesData()
        End If
    End Sub
    Sub BindArticlesData()
        'Dim dsArticle As DataSet   
        Dim pa As New ParaHelper
        pa.addVarchar("blogger", ph.blogger, 50)

        dgArticles.DataSource = connHelper.retrieve("select_doc_by_blogger", CommandType.StoredProcedure, pa.parameters) 'dsArticle.tables("Articles")
        dgArticles.DataBind()
    End Sub

    Sub EditArticle(ByVal Obj As Object, ByVal e As DataGridCommandEventArgs) Handles dgArticles.EditCommand
        Response.Redirect("/write.aspx?edit=" & dgArticles.DataKeys.Item(e.Item.ItemIndex))
    End Sub

    Sub DeleteArticle(ByVal Obj As Object, ByVal e As DataGridCommandEventArgs) Handles dgArticles.DeleteCommand
        Dim blogId As Integer
        blogId = dgArticles.DataKeys.Item(e.Item.ItemIndex)
        Dim pa As New ParaHelper
        pa.addInt("id", blogId)
        pa.addInt("blogger_id", ph.bloggerID)
        connHelper.exeNonQuery("del_doc", CommandType.StoredProcedure, pa.parameters)
        'connHelper.exeNonQuery("delete from articles where [index]=" & blogId & " and blogger=" & ph.bloggerID, CommandType.Text, Nothing)
        'connHelper.exeNonQuery("delete from comment where article=" & blogId & " and blogger=" & ph.bloggerID, CommandType.Text, Nothing)

        BindArticlesData()
    End Sub

    Sub DataGrid_ItemCreated(ByVal Sender As Object, ByVal e As DataGridItemEventArgs) Handles dgArticles.ItemCreated

        Select Case e.Item.ItemType
            Case ListItemType.Item, ListItemType.AlternatingItem, ListItemType.EditItem
                Dim myTableCell As TableCell
                myTableCell = e.Item.Cells(4)
                Dim myDeleteButton As LinkButton
                myDeleteButton = myTableCell.Controls(0)
                myDeleteButton.Attributes.Add("onclick", "return confirm('你真的要删除此文章?');")
        End Select

    End Sub
    Private Sub dgArticles_PageIndexChanged(ByVal sender As System.Object, ByVal e As System.Web.UI.WebControls.DataGridPageChangedEventArgs) Handles dgArticles.PageIndexChanged
        dgArticles.CurrentPageIndex = e.NewPageIndex
        BindArticlesData()
        'Dim dsArticle As DataSet
        'dsArticle = New DataSet
        'dsArticle.Tables.Add("Articles")
        'sql = "select_doc_by_blogger"
        'dbComm = New SqlDataAdapter(sql, conn)
        'dbComm.SelectCommand.CommandType = CommandType.StoredProcedure
        'dbComm.SelectCommand.Parameters.Add("@blogger", SqlDbType.VarChar, 50)
        'dbComm.SelectCommand.Parameters.Item(0).Value = Session("blogger")
        'dbComm.Fill(dsArticle, "Articles")
        'dgArticles.DataSource = dsArticle.tables("Articles")
        'dgArticles.DataBind()
    End Sub

End Class
