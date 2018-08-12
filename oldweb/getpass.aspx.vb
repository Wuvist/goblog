Imports System.Web.Mail
Imports SqlConnHelper
Imports System.Text
Namespace blogwind
  Partial Class getpass
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
    End Sub

    Private Sub retrive_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles retrive.Click

      Dim user, email, act, sql As String
      user = ids.Value

      If user <> "" Then
        sql = "select Email from blogger where id=@id"
        Dim pa As New ParaHelper
        pa.addNVarchar("@id", user, 50)
        email = connHelper.exeString(sql, CommandType.Text, pa.parameters, False)
        If email = ems.Value Then
          Dim msg As New MailMessage
          Dim data() As Byte
          msg.From = "info@blogwind.com"
          msg.To = email
          sql = "select pwd from users where user_name=@id"
          pa.clearPara()
          pa.addNVarchar("@id", user, 50)

          data = connHelper.exeScalar(sql, CommandType.Text, pa.parameters)
          act = Encoding.ASCII.GetString(data).GetHashCode.GetHashCode.ToString

          msg.Subject = "您在博客风上的密码"
          msg.Body = user + ",您好！" + vbCrLf + "请点击下列链接设置您的新密码： http://www.blogwind.com/changepass.aspx?id=" + user + "&code=" + act + vbCrLf + vbCrLf + "博客风" + vbCrLf + "http://www.blogwind.com" + vbCrLf + "info@blogwind.com"
          SmtpMail.SmtpServer = "127.0.0.1"
          SmtpMail.Send(msg)
          ids.Value = "密码已经送出"
          ems.Value = "请查询您的信箱。"
        Else
          ems.Value = "请输入您注册帐号时填写的信箱。"
        End If
      End If
    End Sub
  End Class
End Namespace