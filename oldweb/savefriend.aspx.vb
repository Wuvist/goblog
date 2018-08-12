Imports System.IO
Imports System.Web.Mail
Imports SqlConnHelper
Namespace blogwind
    Partial Class savefriend
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

            If Session("blogger") <> "" Then
                Dim blogger, friendid, action As String

                Dim id As Integer

                friendid = Request("friendid")
                action = Request("action")
                Dim pa As New ParaHelper
                If action = "add" Then
                    pa.addVarchar("@friendid", friendid, 50)
                    pa.addVarchar("@blogger", Session("blogger"), 50)

                    If Request.Form("reveal") <> "" Then
                        pa.addBit("reveal", 1)
                    Else
                        pa.addBit("reveal", 0)
                    End If
                    connHelper.exeNonQuery("insert_friend_reveal", CommandType.StoredProcedure, pa.parameters)
                    Dim email As String
                    email = connHelper.exeString("select email from blogger where id='" & friendid & "'", CommandType.Text, Nothing, False)

                    If (Not (email Is Nothing Or email = "")) And Request.Form("reveal") <> "" Then
                        Dim msg As New MailMessage
                        msg.From = "info@blogwind.com"
                        msg.BodyFormat = MailFormat.Html
                        msg.To = email
                        msg.Subject = "博客风提示：" & Session("blogger") & " 添加您为朋友了"
                        msg.Body = "<p>博客风提示：<a href=""http://www.blogwind.com/" & Session("blogger") & "/"" target=""_blank"">" & Session("blogger") & "</a> 添加您为朋友了</p>"
                        msg.Body = msg.Body + "<p>点击 <a href=""http://wiki.blogwind.com/doku.php?id=%E6%9C%8B%E5%8F%8B"" target=""_blank"">这里</a>  查看博客风朋友功能的详细介绍</p>"
                        msg.Body = msg.Body + "<p>博客风<br><a href=""http://www.blogwind.com"">http://www.blogwind.com</a><br>info@blogwind.com</p>"

                        SmtpMail.SmtpServer = "127.0.0.1"
                        SmtpMail.Send(msg)
                    End If
                Else
                    pa.addVarchar("@friendid", friendid, 50)
                    pa.addVarchar("@blogger", Session("blogger"), 50)
                    connHelper.exeNonQuery("del_friend", CommandType.StoredProcedure, pa.parameters)
                End If


                'Dim fpath, content As String
                'fpath = Session("blogger") & ".xml"
                'fpath = Server.MapPath("/images/" + fpath)
                'Dim strwriterobj As StreamWriter
                'Dim dbComm As SqlDataAdapter
                'Dim ds As New DataTable("flist")

                'dbComm = New SqlDataAdapter("select_friendlist_by_blogger", conn)
                'dbComm.SelectCommand.CommandType = CommandType.StoredProcedure
                'dbComm.SelectCommand.Parameters.Add("@blogger", SqlDbType.VarChar, 50)
                'dbComm.SelectCommand.Parameters.Item(0).Value = Session("blogger")
                'dbComm.Fill(ds)
                'Dim dr As DataRow

                'content = "<?xml version=""1.0""?><tree>"
                'For Each dr In ds.Rows
                '    content = content & "    <tree text=""" & dr.Item("nick") & """ src=""" & dr.Item("id") & ".xml"" />" & vbCrLf
                'Next
                'content = content + "</tree>"

                'strwriterobj = File.CreateText(fpath)
                'strwriterobj.WriteLine(content)
                'strwriterobj.Close()


                If Not Request.Form("popout") = "pop" Then
                    Response.Redirect("/")
                End If
            Else
                Response.Write("请重新登陆博客风。")
            End If
        End Sub

    End Class
End Namespace