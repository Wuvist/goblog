Imports SqlConnHelper
Namespace blogwind
  Partial Class log
    Inherits System.Web.UI.UserControl
    Public showfriend As String
    Public nick As String
    Public welcomemsg As String
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
      nick = ph.nick

      Dim blogger, friendid As String
      blogger = ph.blogger
      friendid = Request.QueryString("blogger")

      If blogger <> "" AndAlso friendid <> "" AndAlso friendid.ToLower <> blogger.ToLower Then
        Dim pa As New ParaHelper
        pa.addNVarchar("@blogger", blogger, 50)
        pa.addNVarchar("@friend", friendid, 50)
        If connHelper.exeInteger("check_friend", CommandType.StoredProcedure, pa.parameters, False) = 0 Then
          showfriend = "<form action='/savefriend.aspx' method='post' target='_blank'><input type='hidden' name='friendid' value='" & friendid & "'><input type='hidden' name='action' value='add'><input type='submit' value='加为朋友' class='logbutton' title='添加" & friendid & "为您的朋友'><input type='hidden' name='popout' value='pop'></form>"
        Else
          showfriend = ""
        End If
      End If
      If ph.login Then
        welcomemsg = connHelper.exeString("select top 1 content from Announcement order by id desc", CommandType.Text, Nothing, Nothing)
      End If
    End Sub

  End Class
End Namespace