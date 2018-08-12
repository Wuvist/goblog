Imports SqlConnHelper
Namespace blogwind
    Partial Class admin_userinfo
        Inherits System.Web.UI.Page
        Protected ph As pageHelper

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
            ph = New pageHelper(Me.Session, Me.Request)

            If Session("blogger") = "" Then
                Response.Write("请重新登陆。")
                Response.End()
            End If
            If Not IsPostBack Then
                BindUserData()
                MessagePage1.Text = ""
                MessagePage2.Text = ""
            End If
        End Sub
        Sub BindUserData()
            Dim sql As String
            sql = "SELECT * FROM Blogger WHERE id= '" & Session("blogger") & "'"
            Dim dr As DataRow
            dr = connHelper.exeDataRow(sql, CommandType.Text, Nothing)

            UserID.Text = dr("id")
            NickName.Text = dr("nick")
            DOB.Text = dr("DOB")
            BlogName.Text = dr("blogname")
            Email.Text = dr("email")
            Intro.Text = dr("intro")
            If Not dr.IsNull("widget") Then
                widget.Text = dr("widget")
            End If

            If dr("reveal") = True Then
                reveal.SelectedIndex = 0
            Else
                reveal.SelectedIndex = 1
            End If
        End Sub

        Private Sub Update_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Update.Click
            Dim UpdateCmd As String = "UPDATE Blogger SET " & _
                   " nick = @NickName, " & _
                   " DOB = @DOB, " & _
                   " blogname= @BlogName, " & _
                   " email= @Email, " & _
                   " widget= @widget, " & _
                   " intro= @intro, Reveal=@reveal " & _
                                    " WHERE id = '" & Session("blogger") & "'"
            Dim pa As New ParaHelper
            pa.addNVarchar("@NickName", NickName.Text, 50)
            pa.addNVarchar("@DOB", DOB.Text, 50)
            pa.addNVarchar("@BlogName", BlogName.Text, 50)
            pa.addNVarchar("@Email", Email.Text, 50)
            pa.addNVarchar("@intro", Intro.Text, 500)
            pa.addNVarchar("@widget", widget.Text, 500)

            If reveal.SelectedValue = "yes" Then
                pa.addBit("@reveal", 1)
            Else
                pa.addBit("@reveal", 0)
            End If

            connHelper.exeNonQuery(UpdateCmd, CommandType.Text, pa.parameters)

            BindUserData()
            MessagePage1.Text = "资料修改成功！"
        End Sub

        Private Sub pwd_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles pwd.Click
            If String.Compare(NewPwd.Text, ReNewPwd.Text) = -1 Then
                MessagePage2.Text = "两次输入的密码不一致"
                Exit Sub
            Else

                Dim pa As New ParaHelper
                Dim data() As Byte
                Dim id As String
                Dim md5 As New System.Security.Cryptography.MD5CryptoServiceProvider
                id = Session("blogger")
                pa.addVarchar("@id", id, 50)
                data = System.Text.Encoding.ASCII.GetBytes(OldPwd.Text.ToCharArray)
                pa.addBinary("@pwd", md5.ComputeHash(data))

                If connHelper.exeString("blogger_login", CommandType.StoredProcedure, pa.parameters, Nothing) = "" Then
                    MessagePage2.Text = "原密码错误！"
                    Exit Sub
                Else
                    pa.clearPara()
                    pa.addVarchar("@id", id, 50)
                    data = System.Text.Encoding.ASCII.GetBytes(NewPwd.Text.ToCharArray)
                    pa.addBinary("@pass", md5.ComputeHash(data))
                    connHelper.exeNonQuery("update_user_password", CommandType.StoredProcedure, pa.parameters)

                    MessagePage2.Text = "密码已经更改！"
                End If
            End If
        End Sub
    End Class
End Namespace