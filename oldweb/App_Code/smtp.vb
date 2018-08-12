Imports Microsoft.VisualBasic

Public Class SMTP
    Public Shared Sub send(ByVal msg As System.Net.Mail.MailMessage)
        Dim s As New System.Net.Mail.SmtpClient
        s.EnableSsl = ini.EnableSmtpSSL
        s.Send(msg)
    End Sub
End Class
