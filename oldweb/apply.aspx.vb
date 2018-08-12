Imports System.Text
Imports SqlConnHelper

Namespace blogwind
    Partial Class apply
        Inherits System.Web.UI.Page

#Region " Web Form Designer Generated Code "

        'This call is required by the Web Form Designer.
        <System.Diagnostics.DebuggerStepThrough()> Private Sub InitializeComponent()

        End Sub
        Protected WithEvents Panel1 As System.Web.UI.WebControls.Panel


        Private Sub Page_Init(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Init
            'CODEGEN: This method call is required by the Web Form Designer
            'Do not modify it using the code editor.
            InitializeComponent()
        End Sub

#End Region

        Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
            'Put user code to initialize the page here
        End Sub

        Private Sub Submit1_ServerClick(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Submit1.ServerClick
            Dim err As Boolean
            Dim id, sql As String
            id = userid.Text

            err = False

            If check(id) <> "valid" Then
                Hint.Text = check(id)
                err = True
            Else

                sql = "select count(id) from blogger where id='" & id & "'"

                If connHelper.exeInteger(sql, CommandType.Text, Nothing, False) > 0 Or (id.ToLower = "rewu") Or (id.ToLower = "system") Or (id.ToLower = "images") Or (id.ToLower = "template") Then
                    Hint.Text = "该用户名已经被注册！"
                    err = True
                End If
                sql = "select count(picker) from RR_picker where picker='" & id & "'"

                If connHelper.exeInteger(sql, CommandType.Text, Nothing, False) > 0 Then
                    Hint.Text = "该用户名已经被注册！"
                    err = True
                End If
            End If
                    Hint.Text = "该用户名已经被注册！"
                    err = True
            If err = False Then
                Dim md5 As New System.Security.Cryptography.MD5CryptoServiceProvider
                Dim data() As Byte
                Dim pass As Byte()
                Dim act As String
                Dim un, password As String
                un = id
                password = pw1.Text


                sql = "insert_blogger"
                Dim pa As New ParaHelper
                pa.addVarchar("@id", un, 50)
                data = System.Text.Encoding.ASCII.GetBytes(pw1.Text.ToCharArray)
                pass = md5.ComputeHash(data)
                pa.addBinary("@password", pass)

                act = "http://www.blogwind.com/activate.aspx?user=" + un + "&pass=" + Encoding.Unicode.GetString(pass).ToString.GetHashCode.ToString

                pa.addVarchar("@nick", nick.Text, 50)

                Dim d As New Date
                Try
                    d = CDate(DOB.Text)
                Catch
                    d = Now
                End Try
                pa.addDateTime("@DOB", d)

                pa.addVarchar("@blogname", blogname.Text, 50)
                pa.addInt("@blogskin", 6)
                pa.addNVarchar("@email", email.Text, 50)
                pa.addVarchar("@intro", intro.Text, 500)

                connHelper.exeNonQuery(sql, CommandType.StoredProcedure, pa.parameters)

                Dim msg As New System.Net.Mail.MailMessage
                msg.From = New System.Net.Mail.MailAddress("info@blogwind.com", "Blogwind")
                msg.To.Add(New System.Net.Mail.MailAddress(Request.Form("email")))
                msg.Subject = "欢迎您在博客风注册帐号！"
                msg.IsBodyHtml = False
                msg.BodyEncoding = System.Text.Encoding.GetEncoding("GB2312")
                msg.Body = "请您先访问以下连接激活您的帐号：" + vbCrLf
                msg.Body = msg.Body + act + vbCrLf
                msg.Body = msg.Body + "您的用户名是：" + un + vbCrLf + "密码是：" + password + "（大小写敏感）"
                msg.Body = msg.Body + vbCrLf + "Please visit the following link to activate your account first:" + vbCrLf
                msg.Body = msg.Body + act + vbCrLf
                msg.Body = msg.Body + "Your login ID is:" + un + vbCrLf + "Password is:" + password + "(Case sensitive)"

                msg.Body = msg.Body + vbCrLf + "博客风" + vbCrLf + "http://www.blogwind.com" + vbCrLf + "info@blogwind.com"

                Try
                    SMTP.send(msg)
                Catch ex As Exception

                End Try

                Session("blogger") = un
                Dim ck As HttpCookie
                ck = New HttpCookie("blogger")
                ck.Name = "blogger"
                ck.Item("id") = un
                ck.Item("pass") = Encoding.ASCII.GetString(pass).GetHashCode.ToString
                ck.Expires = DateTime.Now.AddDays(365)
                Response.Cookies.Set(ck)

                Response.Redirect("selectskin.aspx")
            End If
        End Sub

        Function check(ByVal a As String) As String
            Dim i As Integer
            For i = 0 To a.Length - 1
                If Not (Char.IsDigit(userid.Text.Chars(i)) Or Char.IsLower(userid.Text.Chars(i)) Or Char.IsUpper(userid.Text.Chars(i))) Then
                    Return "用户名只能由数字或者英文字母组成，并且不可以包括空格。"
                End If
            Next
            Return "valid"
        End Function
    End Class
End Namespace