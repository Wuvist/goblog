Imports SqlConnHelper
Namespace blogwind
  Partial Class skin_switch
    Inherits System.Web.UI.Page
    Public skin_id As Integer
    Public dt As DataTable
    Public blogger As String
    Public sql As String
    Protected ph As pageHelper

#Region " Web Form Designer Generated Code "

    'This call is required by the Web Form Designer.
    <System.Diagnostics.DebuggerStepThrough()> Private Sub InitializeComponent()

    End Sub
    Protected WithEvents skin_list As System.Web.UI.WebControls.Repeater


    Private Sub Page_Init(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Init
      'CODEGEN: This method call is required by the Web Form Designer
      'Do not modify it using the code editor.
      InitializeComponent()
    End Sub

#End Region

    Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
      'Put user code to initialize the page here
      ph = New pageHelper(Me.Session, Me.Request)
      blogger = Session("blogger")

      'blogger = Request.ServerVariables("REMOTE_ADDR")

      Dim path As String
      path = Request.QueryString("skin")
      If path <> "" Then
        skin_id = CInt(path)
      Else
        skin_id = -1
      End If


      If skin_id > -1 Then
        sql = "update_skin '" & blogger & "'," & skin_id
        connHelper.exeNonQuery(sql, CommandType.Text, Nothing)
      End If

      sql = "select Blogger.blogskin from Blogger inner join Skin on (Skin.[Index]=Blogger.blogskin) where Blogger.id='" & blogger & "'"

      skin_id = connHelper.exeInteger(sql, CommandType.Text, Nothing, False)

      sql = "select * from Skin where [index]<>" & skin_id & " order by popularity DESC"
      dt = connHelper.retrieve(sql, CommandType.Text, Nothing)
    End Sub
  End Class
End Namespace