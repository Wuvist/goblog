Imports SqlConnHelper
Imports System.Text
Namespace blogwind
  Partial Class changepass
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
      Dim username, p, sql As String
      Dim data() As Byte

      username = Request.QueryString("id")
      p = Request.QueryString("code")

      sql = "select pwd from users where user_name=@id"

      'cn.cmd.Parameters.Add("@id", SqlDbType.VarChar, 20)
      'cn.cmd.Parameters.Item(0).Value = username
      Dim pa As New ParaHelper
      pa.addNVarchar("@id", username, 50)
      data = connHelper.exeScalar(sql, CommandType.Text, pa.parameters)

      If data.Length = 0 Then
        Response.Write("用户名：" + Request.QueryString("user") + "不存在！")
        Response.End()
      End If

      If Encoding.ASCII.GetString(data).GetHashCode.GetHashCode.ToString <> p Then
        Response.Write("无法重试密码，请联系博客风管理员：info@blogwind.com")
        Response.End()
      End If

    End Sub

    Private Sub Button1_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button1.Click
      Dim data() As Byte
      Dim md5 As New System.Security.Cryptography.MD5CryptoServiceProvider

      Dim sql As String
      sql = "UPDATE users SET pwd= @pass WHERE user_name =@id"
      Dim pa As New ParaHelper
      pa.addNVarchar("id", Request.QueryString("id"), 50)
      'cn.cmd.Parameters.Add("@pass", SqlDbType.Binary, 16)
      data = System.Text.Encoding.ASCII.GetBytes(p1.Text.ToCharArray)
      pa.addBinary("pass", md5.ComputeHash(data))
      connHelper.exeNonQuery(sql, CommandType.Text, pa.parameters)
      hint.Text = "密码修改成功！请<a href='http://www.blogwind.com' target='_top'>返回博客风首页</a>登陆。"
    End Sub
  End Class
End Namespace