Imports SqlConnHelper
Namespace blogwind
  Partial Class userpic
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
      Dim blogger, msg As String
      Dim rd As Boolean
      rd = True

      blogger = Session("blogger")
      If blogger <> "" Then

        Dim action As String
        action = Request.Form("action")
        Select Case action
          Case "close"
            connHelper.exeNonQuery("update blogger set userpic = userpic %10 +20 where id='" & blogger & "'", CommandType.Text, Nothing)
          Case "show"
            connHelper.exeNonQuery("update blogger set userpic = userpic %10 +10 where id='" & blogger & "'", CommandType.Text, Nothing)
          Case Else
            If action.StartsWith("copy:") Then
              Dim friendid As String
              friendid = action.Substring(5)
              If System.IO.File.Exists(Server.MapPath("/images/userpic/" & friendid & ".jpg")) Then
                connHelper.exeNonQuery("update blogger set userpic = 11 where id='" & blogger & "'", CommandType.Text, Nothing)
                System.IO.File.Copy(Server.MapPath("/images/userpic/" & friendid & ".jpg"), Server.MapPath("/images/userpic/" & blogger & ".jpg"), True)
                msg = "头像复制成功！"
                rd = False
              Else
                rd = False
                connHelper.exeNonQuery("update blogger set userpic = 10 where id='" & blogger & "'", CommandType.Text, Nothing)
                msg = "汗？为什么有人要复制系统内置的头像呢？你现在的头像已经是默认头像了～"
              End If
            End If
        End Select

        If rd Then
          Response.Redirect(Request.UrlReferrer.AbsoluteUri)
        Else
          Response.Write("<script>alert('" & msg & "');history.go(-1);</script>")
        End If

      End If
    End Sub

  End Class
End Namespace