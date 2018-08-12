Imports SqlConnHelper
Namespace blogwind
  Partial Class savelink
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

      If Session("blogger") <> "" Then
        Dim blogger, url, link As String

        Dim id As Integer

        url = Request.Form("url")
        link = Request.Form("link")

        Dim pa As New ParaHelper
        pa.addNVarchar("@link", link, 50)
        pa.addNVarchar("@url", url, 100)
        pa.addNVarchar("@blogger", Session("blogger"), 50)
        'pa.Add("@reveal", SqlDbType.Bit)
        If Request.Form("reveal") <> "" Then
          pa.addBit("@reveal", 1)
        Else
          pa.addBit("@reveal", 0)
          'cn.cmd.Parameters.Item(3).Value = 0
        End If

        connHelper.exeNonQuery("insert_link_reveal", CommandType.StoredProcedure, pa.parameters)
        Response.Redirect("/")
      End If
    End Sub

  End Class
End Namespace
