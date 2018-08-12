Imports SqlConnHelper
Namespace blogwind
  Partial Class news
    Inherits System.Web.UI.Page
    Public subject As String
    Public content As String

#Region " Web Form Designer Generated Code "

    'This call is required by the Web Form Designer.
    <System.Diagnostics.DebuggerStepThrough()> Private Sub InitializeComponent()

    End Sub
    Protected WithEvents Hint As System.Web.UI.WebControls.Label


    Private Sub Page_Init(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Init
      'CODEGEN: This method call is required by the Web Form Designer
      'Do not modify it using the code editor.
      InitializeComponent()
    End Sub

#End Region

    Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
      'Put user code to initialize the page here
      Dim ds As DataTable
      Dim i As Integer
      Try
        i = CInt(Request.QueryString("id"))
      Catch ex As Exception
        i = 3
      End Try

      ds = connHelper.retrieve("select * from Announcement where id=" & i, CommandType.Text, Nothing)
      subject = ds.Rows(0).Item("subject")
      content = ds.Rows(0).Item("content")
      content = Replace(content, "  ", "&nbsp;&nbsp;")
      content = Replace(content, vbNewLine, "<br>")
      ds.Dispose()
    End Sub

  End Class
End Namespace