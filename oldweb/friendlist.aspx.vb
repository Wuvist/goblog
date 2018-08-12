Imports SqlConnHelper
Namespace blogwind
  Partial Class friendlist
    Inherits System.Web.UI.Page
    Public nick As String

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
      'Put user code to initialize the page here3.        
      Dim blogger As String
      blogger = Request.QueryString("blogger")
      If blogger = "" Then
        Response.End()
      End If

      Dim dt As New DataTable("flist")
      Dim ds As New DataTable("friends")

      nick = connHelper.exeString("select nick from blogger where id='" & blogger & "'", CommandType.Text, Nothing, False)
      Dim pa As New ParaHelper

      pa.addNVarchar("@blogger", blogger, 50)
      friends.DataSource = connHelper.retrieve("select_friendlist_by_blogger", CommandType.StoredProcedure, pa.parameters)
      friends.DataBind()

      pa.clearPara()
      pa.addNVarchar("@blogger", blogger, 50)
      flist.DataSource = connHelper.retrieve("select_flist_by_blogger", CommandType.StoredProcedure, pa.parameters)
      flist.DataBind()
    End Sub
  End Class
End Namespace