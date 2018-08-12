Imports SqlConnHelper
Namespace blogwind
  Partial Class admin_comment
    Inherits System.Web.UI.Page

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
      Dim ph As New pageHelper(Me.Session, Me.Request)

      If Session("blogger") = "" Then
        Response.Write("请重新登陆。")
        Response.End()
      End If
      If Not IsPostBack Then
        BindCommentData()
      End If
    End Sub

    Private Sub BindCommentData()
      Dim sql As String
      sql = "select_comment_by_blogger"
      Dim pa As New ParaHelper
      pa.addVarchar("@blogger", Session("blogger"), 50)

      dgComment.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
      dgComment.DataBind()
    End Sub

    Private Sub dgComment_ItemCreated(ByVal Sender As Object, ByVal e As DataGridItemEventArgs) Handles dgComment.ItemCreated
      Select Case e.Item.ItemType
        Case ListItemType.Item, ListItemType.AlternatingItem, ListItemType.EditItem
          Dim myTableCell As TableCell
          myTableCell = e.Item.Cells(3)
          Dim myDeleteButton As LinkButton
          myDeleteButton = myTableCell.Controls(0)
          myDeleteButton.Attributes.Add("onclick", "return confirm('你真的要删除此留言?');")
      End Select
    End Sub
    Private Sub DeleteComment(ByVal Obj As System.Object, ByVal e As System.Web.UI.WebControls.DataGridCommandEventArgs) Handles dgComment.DeleteCommand
      Dim pa As New ParaHelper
      pa.addInt("@id", dgComment.DataKeys.Item(e.Item.ItemIndex))
      pa.addNVarchar("@blogger", Session("blogger"), 50)
      connHelper.exeScalar("del_comment_safe", CommandType.StoredProcedure, pa.parameters)

      BindCommentData()
    End Sub

    Private Sub dgComment_PageIndexChanged(ByVal sender As System.Object, ByVal e As System.Web.UI.WebControls.DataGridPageChangedEventArgs) Handles dgComment.PageIndexChanged
      dgComment.CurrentPageIndex = e.NewPageIndex
      BindCommentData()
    End Sub
  End Class
End Namespace