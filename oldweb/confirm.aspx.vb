Imports SqlConnHelper
Imports System.Text
Imports System.Web.Mail
Namespace blogwind
  Partial Class confirm
    Inherits System.Web.UI.Page
    Public nick, blogger As String


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
        Response.End()
      End If


      Dim dt As DataRow
      blogger = Session("blogger")

      dt = connHelper.exeDataRow("select id, nick, email from blogger where id='" & blogger & "'", CommandType.Text, Nothing)
      nick = dt.Item("nick")

      If Not IsPostBack Then
        email.Text = dt.Item("email")
      End If

    End Sub

    Private Sub Button1_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button1.Click
      Dim msg As New MailMessage
      Dim act, sql As String
      Dim md5 As New System.Security.Cryptography.MD5CryptoServiceProvider

      Dim dt As DataTable
      Dim data() As Byte


      sql = "select pwd from users where user_name=@id"

      'cn.cmd.Parameters.Add("@id", SqlDbType.VarChar, 20)
      'cn.cmd.Parameters.Item(0).Value = blogger
      Dim pa As New ParaHelper
      pa.addNVarchar("@id", blogger, 50)
      data = connHelper.exeScalar(sql, CommandType.Text, pa.parameters)

      connHelper.exeNonQuery("update blogger set email='" & email.Text & "' where id='" & blogger & "'", CommandType.Text, Nothing)

      msg.From = "info@blogwind.com"
      msg.To = email.Text
      msg.Subject = "欢迎您在博客风注册帐号！"
      msg.Body = "请您先访问以下连接激活您的帐号：" + vbCrLf
      act = "http://www.blogwind.com/activate.aspx?user=" + blogger + "&pass=" + Encoding.Unicode.GetString(data).ToString.GetHashCode.ToString
      msg.Body = msg.Body + act + vbCrLf

      msg.Body = msg.Body + vbCrLf + "Please visit the following link to activate your account first:" + vbCrLf
      msg.Body = msg.Body + act + vbCrLf

      msg.Body = msg.Body + vbCrLf + "博客风" + vbCrLf + "http://www.blogwind.com" + vbCrLf + "wuvist@gmail.com"

      SmtpMail.SmtpServer = "220.249.204.201"

      hint.Text = "新的确认信件已经发送至<font color=red>" & email.Text & "</font>，"

      SmtpMail.Send(msg)
    End Sub
  End Class
End Namespace