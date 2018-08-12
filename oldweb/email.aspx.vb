Imports SqlConnHelper

Namespace blogwind
    Partial Class email : Inherits System.Web.UI.Page
        Protected nick As String

        Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
            'Put user code to initialize the page here
            Dim sql As String
            sql = "select nick from blogger where id=@id"
            Dim pa As New ParaHelper
            nick = Request.QueryString("user")
            pa.addNVarchar("@id", nick, 50)
            nick = connHelper.exeString(sql, CommandType.Text, pa.parameters, False) & "(" & nick & ")"
        End Sub

        Private Sub Button1_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button1.Click
            Dim sql As String
            sql = "select email from blogger where id=@id"
            Dim pa As New ParaHelper
            pa.addNVarchar("@id", Request.QueryString("user"), 50)

            Dim msg As New System.Net.Mail.MailMessage


            msg.From = New System.Net.Mail.MailAddress("info@blogwind.com", "Blogwind")
            msg.To.Add(New System.Net.Mail.MailAddress(connHelper.exeString(sql, CommandType.Text, pa.parameters, False)))
            msg.Subject = "您的博客访客给你留言了"
            msg.BodyEncoding = System.Text.Encoding.GetEncoding("GB2312")
            msg.ReplyTo = New System.Net.Mail.MailAddress(e_mail.Text, name.Text)
            msg.Body = "访客：" + name.Text + vbCrLf + "电邮：" + e_mail.Text + vbCrLf + "主题：" + title.Text + vbCrLf + content.Text
            SMTP.send(msg)

            title.Text = ""
            hint.Text = "发送成功"
            content.Text = ""
            e_mail.Text = ""
        End Sub
    End Class
End Namespace