Imports System.Text
Imports SqlConnHelper
Namespace blogwind
    Partial Class activate
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

            username = Request.QueryString("user")
            p = Request.QueryString("pass")

            sql = "select pwd from users where user_name=@id"
            Dim pa As New ParaHelper

            'cn.cmd.Parameters.Add("@id", SqlDbType.VarChar, 20)
            'cn.cmd.Parameters.Item(0).Value = username
            pa.addNVarchar("@id", username, 50)
            data = connHelper.exeScalar(sql, CommandType.Text, pa.parameters)
            If data Is Nothing Then
                hint.Text = "用户名：" + Request.QueryString("user") + "不存在！"
            Else
                If Encoding.Unicode.GetString(data).GetHashCode.ToString = p Then
                    sql = "update blogger set Activate=1 where id='" + username + "'"
                    connHelper.exeNonQuery(sql, CommandType.Text, Nothing)
                    hint.Text = username + "激活成功！<a href='default.aspx'>点此返回首页</a>"
                Else
                    hint.Text = "激活失败！"
                End If
            End If

        End Sub
    End Class
End Namespace
