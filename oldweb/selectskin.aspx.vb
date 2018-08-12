Imports SqlConnHelper
Namespace blogwind
  Partial Class selectskin
    Inherits System.Web.UI.Page
    Public dt As DataTable
    Public blogger As String
    Public sql As String
    Public skin_id As Integer

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
        Response.End()
      End If
      blogger = Session("blogger")


      'blogger = Request.ServerVariables("REMOTE_ADDR")

      Dim path As String

      skin_id = Request.QueryString("skin")

      If skin_id <> 0 Then
        path = Server.MapPath("")
        sql = "update_skin '" & blogger & "'," & skin_id
        connHelper.exeNonQuery(sql, CommandType.Text, Nothing)
        Response.Redirect("confirm.aspx")

        Response.End()
      Else
        sql = "select * from Skin order by popularity DESC"
        dt = connHelper.retrieve(sql, CommandType.Text, Nothing)
      End If
    End Sub

  End Class
End Namespace
