Imports SqlConnHelper
Namespace blogwind
  Partial Class myfriends
    Inherits System.Web.UI.Page
    Protected sql As String


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
      If Session("blogger") <> "" Then
        Dim sql As String
        Dim pa As New ParaHelper
        Dim fbpage, fcpage, flpage As Integer
        Dim bid As Integer
        sql = "select [index] from blogger where id='" & Session("blogger") & "'"

        bid = connHelper.exeInteger(sql, CommandType.Text, Nothing, False)

        fbpage = CInt(Request.QueryString("fblog"))
        flpage = CInt(Request.QueryString("fbm"))
        fcpage = CInt(Request.QueryString("fcom"))
        If fcpage < 1 Then
          friendcomment.Visible = False
        Else
                    sql = "select_friendevents"
          pa.addInt("@bloggerid", bid)
          pa.addInt("@page", fcpage)
          friendcomment.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
          friendcomment.DataBind()
        End If

        If fbpage < 1 Then
          friendblog.Visible = False
        Else
          sql = "select_friendblog"
          pa.clearPara()
          pa.addInt("@bloggerid", bid)
          pa.addInt("@page", fbpage)

          friendblog.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
          friendblog.DataBind()
        End If
        If flpage < 1 Then
          friendbookmark.Visible = False
        Else
          sql = "select_friendbookmark"
          pa.clearPara()
          pa.addInt("@bloggerid", bid)
          pa.addInt("@page", flpage)

          friendbookmark.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
          friendbookmark.DataBind()
          If friendbookmark.Items.Count = 0 Then
            userview.Visible = True
          End If
        End If
      End If
    End Sub
  End Class
End Namespace