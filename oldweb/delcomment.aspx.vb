Imports SqlConnHelper
Namespace blogwind
  Partial Class delcomment
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
      If Session("blogger") = "" Then
        Response.Write("2")
      Else

        Dim i As Integer
        Try
          i = CInt(Request.QueryString("cid"))
        Catch ex As Exception
          Exit Sub
        End Try
        Dim pa As New ParaHelper
        pa.addInt("@id", i)
        pa.addNVarchar("@blogger", Session("blogger"), 50)
        i = connHelper.exeScalar("del_comment_safe", CommandType.StoredProcedure, pa.parameters)

        Response.Write(i)
      End If
    End Sub

  End Class
End Namespace