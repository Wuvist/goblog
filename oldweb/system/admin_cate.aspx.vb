Imports SqlConnHelper
Partial Class admin_cate
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
      BindCateData()
    End If
  End Sub
  Sub BindCateData()
    Dim sql As String
    Dim dt As DataTable

    sql = "select_cate_by_blogger"
    Dim pa As New ParaHelper
    pa.addVarchar("@blogger", Session("blogger"), 50)
    dt = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
    dgCate.DataSource = dt
    dgCate.DataBind()

    ddlCate.DataSource = dt
    ddlCate.DataBind()

    If dgCate.Items.Count = 1 Then
      dgCate.Columns(3).Visible = False
      hint.Text = "博客风的网志系统需要每个用户保留至少一个网志分类"
    Else
      dgCate.Columns(3).Visible = True
      hint.Text = ""
    End If

    sql = "select_doc_by_blogger"
    pa.clearPara()
    pa.addVarchar("@blogger", Session("blogger"), 50)
    dgArticlesCate.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
    dgArticlesCate.DataBind()

  End Sub
  Sub EditCate(ByVal Obj As Object, ByVal e As DataGridCommandEventArgs) Handles dgCate.EditCommand
    dgCate.EditItemIndex = e.Item.ItemIndex
    BindCateData()
  End Sub

  Sub UpdateCate(ByVal Obj As Object, ByVal e As DataGridCommandEventArgs) Handles dgCate.UpdateCommand
    Dim sql As String

    sql = "update_cate_name"
    Dim ctlCate As TextBox = e.Item.Cells(0).Controls(0)
    Dim pa As New ParaHelper
    pa.addInt("@cate", dgCate.DataKeys.Item(e.Item.ItemIndex))
    pa.addVarchar("@name", ctlCate.Text, 50)
    connHelper.exeNonQuery(sql, CommandType.StoredProcedure, pa.parameters)
    dgCate.EditItemIndex = -1
    BindCateData()
  End Sub

  Private Sub AddNewCate_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles AddNewCate.Click
    Dim dsCate As DataSet

    Dim sql As String = "insert_cate"
    Dim pa As New ParaHelper
    pa.addVarchar("@name", NewCate.Text, 50)
    pa.addVarchar("@blogger", Session("blogger"), 50)

    connHelper.exeNonQuery(sql, CommandType.StoredProcedure, pa.parameters)

    NewCate.Text = ""
    BindCateData()
  End Sub

  Sub ddlCate_SelectedIndexChanged(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles ddlCate.SelectedIndexChanged
    Dim dg As DataGrid
    dg = CType(dgArticlesCate, DataGrid)

    Dim GridItem As DataGridItem

    For Each GridItem In dgArticlesCate.Items
      If GridItem.Cells(0).Text = ddlCate.SelectedItem.Value Then
        Dim myCheckbox As CheckBox = CType(GridItem.Cells(3).Controls(1), CheckBox)
        myCheckbox.Checked = True
      Else
        Dim myCheckbox As CheckBox = CType(GridItem.Cells(3).Controls(1), CheckBox)
        myCheckbox.Checked = False
      End If
    Next
  End Sub
  Sub dgArticlesCate_PageIndexChanged(ByVal sender As System.Object, ByVal e As System.Web.UI.WebControls.DataGridPageChangedEventArgs) Handles dgArticlesCate.PageIndexChanged
    dgArticlesCate.CurrentPageIndex = e.NewPageIndex
    Dim sql As String
    sql = "select_doc_by_blogger"
    Dim pa As New ParaHelper
    pa.addVarchar("@blogger", Session("blogger"), 50)
    dgArticlesCate.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
    dgArticlesCate.DataBind()
  End Sub

  Sub UpdateArticleCate_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles UpdateArticleCate.Click
    Dim dg As DataGrid
    dg = CType(dgArticlesCate, DataGrid)
    Dim GridItem As DataGridItem

    Dim sql As String
    For Each GridItem In dgArticlesCate.Items
      Dim myCheckbox As CheckBox = CType(GridItem.Cells(3S).Controls(1), CheckBox)
      If myCheckbox.Checked = "true" Then
        sql = "update_cate"
        Dim pa As New ParaHelper
        pa.addInt("@new_cate", ddlCate.SelectedItem.Value)
        pa.addInt("@Article_id", dgArticlesCate.DataKeys.Item(GridItem.ItemIndex))
        connHelper.exeNonQuery(sql, CommandType.StoredProcedure, pa.parameters)
      End If
    Next

    BindCateData()
  End Sub
  Sub dgArticlesCate_ItemCreated(ByVal Obj As Object, ByVal e As DataGridItemEventArgs) Handles dgArticlesCate.ItemCreated
    Dim GridItem As DataGridItem
    For Each GridItem In dgArticlesCate.Items
      If GridItem.Cells(0).Text = ddlCate.SelectedItem.Value Then
        Dim myCheckbox As CheckBox = CType(GridItem.Cells(3).Controls(1), CheckBox)
        myCheckbox.Checked = True
      End If
    Next
  End Sub
  Sub dgCate_ItemCreated(ByVal Sender As Object, ByVal e As DataGridItemEventArgs) Handles dgCate.ItemCreated
    Select Case e.Item.ItemType
      Case ListItemType.Item, ListItemType.AlternatingItem, ListItemType.EditItem
        Dim myTableCell As TableCell
        myTableCell = e.Item.Cells(3)
        Dim myDeleteButton As LinkButton
        myDeleteButton = myTableCell.Controls(0)
        myDeleteButton.Attributes.Add("onclick", "return confirm('你真的要删除此分类? 如果有文章属于此分类，请把先文章转移到其他分类，再删除此分类');")

    End Select
  End Sub
  Sub DeleteCate(ByVal Obj As Object, ByVal e As DataGridCommandEventArgs) Handles dgCate.DeleteCommand
    If e.Item.Cells(1).Text.Compare(e.Item.Cells(1).Text, "0") <> 0 Then
      Exit Sub
    End If

    Dim sql As String = "DELETE from UserDefineCategory Where [index] = @Index"
    Dim pa As New ParaHelper
    pa.addInt("@Index", dgCate.DataKeys.Item(e.Item.ItemIndex))
    connHelper.exeNonQuery(sql, CommandType.Text, pa.parameters)
    BindCateData()
  End Sub

End Class
