Imports SqlConnHelper
Namespace blogwind
  Partial Class updateskin
    Inherits System.Web.UI.Page
    Dim sql As String

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
      If Not IsPostBack Then
        bind()
      End If
      'Put user code to initialize the page here
    End Sub

    Private Sub Button1_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button1.Click
      'Dim skin As Integer
      'Dim path As String

      'skin = CInt(skins.SelectedItem.Value)
      'sql = "select id from Blogger where blogskin=@id"
      'Dim pa As New ParaHelper
      'pa.addInt("@id", skin)
      'cn.comm.Fill(ds, "bloggers")

      'path = Server.MapPath("..")
      'For Each row As DataRow In ds.Tables("bloggers").Rows
      '  sql = path & "\" & row.Item(0)
      '  Try
      '    Kill(sql & "\" & "default.aspx")
      '    Kill(sql & "\" & "UserCate.aspx")
      '    Kill(sql & "\" & "comment.aspx")
      '  Catch ex As Exception

      '  End Try


      '  'FileSystem.FileCopy(path & "\Template\skin" + CStr(skin) + "_default.aspx", sql & "\" & "default.aspx")
      '  'FileSystem.FileCopy(path & "\Template\skin" + CStr(skin) + "_UserCate.aspx", sql & "\" & "UserCate.aspx")
      '  'FileSystem.FileCopy(path & "\Template\skin" + CStr(skin) + "_comment.aspx", sql & "\" & "comment.aspx")
      'Next
      'cn.close()
      'bind()
      'hint.Text = "模板 " & skin & "更新成功！"
    End Sub
    Private Sub files_Update(ByVal sender As System.Object, ByVal e As System.Web.UI.WebControls.DataGridCommandEventArgs) Handles skinlist.UpdateCommand
      Dim index As Integer
      Dim t, y, c As TextBox

      t = CType(e.Item.Cells(0).Controls(0), TextBox)
      y = CType(e.Item.Cells(1).Controls(0), TextBox)
      c = CType(e.Item.Cells(2).Controls(0), TextBox)

      index = CInt(skinlist.DataKeys(skinlist.EditItemIndex))

      sql = "update Skin set name=@name, designer=@designer, intro=@intro where [index]=@index"
      Dim pa As New ParaHelper

      'cn.cmd.Parameters.Add("@name", SqlDbType.Char, 10)
      'cn.cmd.Parameters.Item(0).Value = Trim(t.Text)
      'cn.cmd.Parameters.Add("@designer", SqlDbType.Char, 20)
      'cn.cmd.Parameters.Item(1).Value = Trim(y.Text)
      'cn.cmd.Parameters.Add("@intro", SqlDbType.NVarChar)
      'cn.cmd.Parameters.Item(2).Value = Trim(c.Text)
      'cn.cmd.Parameters.Add("@index", SqlDbType.Int, 4)
      'cn.cmd.Parameters.Item(3).Value = index
      'cn.cmd.ExecuteNonQuery()
      'skinlist.EditItemIndex = -1
      'bind()
    End Sub
    Private Sub files_Edit(ByVal sender As System.Object, ByVal e As System.Web.UI.WebControls.DataGridCommandEventArgs) Handles skinlist.EditCommand
      skinlist.EditItemIndex = e.Item.ItemIndex
      bind()
    End Sub
    Public Sub MyDataGrid_ItemDataBound(ByVal senderas As Object, ByVal e As System.Web.UI.WebControls.DataGridItemEventArgs) Handles skinlist.ItemDataBound
      If e.Item.ItemType = ListItemType.EditItem Then
        If e.Item.Controls(0).Controls(0).GetType().ToString = "System.Web.UI.WebControls.TextBox" Then
          Dim tb As TextBox
          tb = e.Item.Controls(0).Controls(0)
          tb.Width = Unit.Parse("90px")
          tb = e.Item.Controls(1).Controls(0)
          tb.Width = Unit.Parse("140px")
          tb = e.Item.Controls(2).Controls(0)
          tb.Width = Unit.Parse("440px")
        End If
      End If
    End Sub

    Private Sub files_Cancel(ByVal sender As System.Object, ByVal e As System.Web.UI.WebControls.DataGridCommandEventArgs) Handles skinlist.CancelCommand
      skinlist.EditItemIndex = -1
      bind()
    End Sub
    Private Sub bind()

      sql = "select * from Skin"
      Dim dt As DataTable
      dt = connHelper.retrieve(sql, CommandType.Text, Nothing)
      skinlist.DataSource = dt
      skinlist.DataBind()
      skins.DataSource = dt
      skins.DataBind()
    End Sub

    Private Sub Button2_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button2.Click
      sql = "insert into Skin (name, designer,intro)values(@name,@designer,@intro)"

      'cn.newcmd(sql)
      'cn.cmd.CommandType = CommandType.Text
      'cn.cmd.Parameters.Add("@name", SqlDbType.Char, 10)
      'cn.cmd.Parameters.Item(0).Value = Trim(nname.Text)
      'cn.cmd.Parameters.Add("@designer", SqlDbType.Char, 20)
      'cn.cmd.Parameters.Item(1).Value = Trim(ndesigner.Text)
      'cn.cmd.Parameters.Add("@intro", SqlDbType.NVarChar)
      'cn.cmd.Parameters.Item(2).Value = Trim(nintro.Text)
      'connHelper.exeNonQuery(sql, CommandType.Text, Nothing)
      bind()

    End Sub
  End Class
End Namespace