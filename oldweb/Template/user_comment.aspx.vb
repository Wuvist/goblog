Imports SqlConnHelper
Namespace blogwind
    Partial Class user_comment
        Inherits BWPage

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
            'Dim dt As Date = Now
            Dim ph As New pageHelper(Me.Session, Me.Request)
            Dim s As Integer
            Dim pa As New ParaHelper

            pa.addVarchar("@id", Request.QueryString("blogger"), 50)
            Dim t As Boolean = False
            s = connHelper.exeInteger("select blogskin from blogger where id=@id", CommandType.Text, pa.parameters, t)
            If t Then
                Response.Redirect("/404.html")
            Else
                If s = 0 Then
                    s = CInt(Int((30 * Rnd()) + 1))
                End If
                Dim c As UserControl
                'If ch.ht.Contains("skin" & s & "_comment.ascx") Then
                '  c = ch.ht.Item("skin" & s & "_comment.ascx")
                'Else
                '  c = LoadControl("skin" & s & "_comment.ascx")
                '  ch.ht.Add("skin" & s & "_comment.ascx", c)
                'End If
                c = LoadControl("skin" & s & "_comment.ascx")
                pch.Controls.Add(c)
            End If
            'Response.Write(Now.Subtract(dt).ToString)
        End Sub

    End Class
End Namespace