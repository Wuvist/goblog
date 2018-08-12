Imports SqlConnHelper

Partial Class setlang
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

        Dim l As langs.lang
        l = CInt(Request.QueryString("lang"))
        ph.lang = l
        Session("lang") = l
        connHelper.exeNonQuery("update blogger set lang='" & l.ToString & "' where [index]=" & Session("bloggerID"), CommandType.Text, Nothing)
        Response.Redirect(Request.UrlReferrer.AbsoluteUri)
    End Sub

End Class
