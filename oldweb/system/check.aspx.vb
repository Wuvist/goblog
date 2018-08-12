Imports System.Data
Imports System.Web
Imports System.Text
Imports SqlConnHelper

Namespace blogwind
    Partial Class check
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
            Dim id As String
            Dim md5 As New System.Security.Cryptography.MD5CryptoServiceProvider
            Dim pass As Byte()
            id = Request.Form("id")

            If id <> "" Then
                Dim ck As HttpCookie
                Dim nick As String
                Dim pa As New ParaHelper

                pass = util.MD5(Request.Form("pass"))

                pa.addVarchar("@id", id, 50)
                pa.addBinary("@pwd", pass)
                Dim b As Byte()
                Dim dr As DataRow
                dr = connHelper.exeDataRow("select nick, id, pwd,TS,lang,DOB from bloggers where user_name=@id and pwd=@pwd", CommandType.Text, pa.parameters)

                If Not dr Is Nothing Then
                    Session("blogger") = id
                    nick = dr("nick")
                    Session("nick") = nick
                    Session("bloggerID") = CInt(dr("id"))
                    Session("lang") = langs.lang.Parse(GetType(langs.lang), CStr(dr("lang")))
                End If

                If nick <> "" Then
                    Session("blogger") = id
                    If Request.Form("record") <> "" Then
                        ck = New HttpCookie("blogger")
                        ck.Name = "blogger"
                        ck.Item("id") = id
                        ck.Item("pass") = Encoding.ASCII.GetString(pass).GetHashCode.ToString
                        ck.Expires = DateTime.Now.AddDays(365)
                        Response.Cookies.Set(ck)
                    End If
                    Response.Redirect("/")
                Else
                    Response.Redirect("/login.aspx?username=" & id)
                End If
            Else
                Response.Redirect("/login.aspx?username=" & id)
            End If
        End Sub
    End Class
End Namespace