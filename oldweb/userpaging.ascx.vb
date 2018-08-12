Imports SqlConnHelper
Namespace blogwind
  Partial Class userpaging
    Inherits System.Web.UI.UserControl
    Public tpage, cpage As Integer

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
      Dim sql As String
      sql = "SELECT * FROM Blogger WHERE id= '" & Request.QueryString("blogger") & "'"
      Dim blogger_id As Integer
      Blogger_id = connHelper.exeDataRow(sql, CommandType.Text, Nothing).Item(0)

      sql = "select count(*) from Articles where blogger=" & Blogger_id

      tpage = connHelper.exeInteger(sql, CommandType.Text, Nothing, False)


      If tpage Mod 10 > 0 Then
        tpage = tpage - tpage Mod 10
        tpage = tpage / 10 + 1
      Else
        tpage = tpage / 10
      End If

      cpage = CInt(Request.QueryString("page"))
      If cpage < 1 Then cpage = 1

      If cpage >= tpage Then
        cpage = tpage
      Else
        down.Text = "<a href='/" & Request.QueryString("blogger") & "/default.aspx?page=" & cpage + 1 & "'>下一页</a>"
      End If

      If cpage > 1 Then
        up.Text = " <a href='/" & Request.QueryString("blogger") & "/default.aspx?page=1'>首页</a>"
        up.Text = up.Text & " <a href='/" & Request.QueryString("blogger") & "/default.aspx?page=" & cpage - 1 & "'>上一页</a>"
      End If

      If cpage < tpage Then
        down.Text = down.Text & " <a href='/" & Request.QueryString("blogger") & "/default.aspx?page=" & tpage & "'>尾页</a>"
      End If

      down.Text = down.Text & " (本页为第 " & cpage & " 页 共" & tpage & "页)"
    End Sub
  End Class
End Namespace
