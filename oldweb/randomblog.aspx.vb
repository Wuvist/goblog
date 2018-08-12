Imports SqlConnHelper
Namespace blogwind
  Partial Class randomblog
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

      Dim blogger As String
      Dim aid As Integer
      Dim r As String
      r = connHelper.exeString("random_blog", CommandType.StoredProcedure, Nothing, False)

      aid = r.IndexOf(" ")
      blogger = r.Substring(aid + 1)
      aid = CInt(r.Substring(0, aid))
      Response.Redirect("/" & blogger & "/" & aid.ToString & ".shtml")
    End Sub
  End Class
End Namespace