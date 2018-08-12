Imports SqlConnHelper
Namespace blogwind
    Partial Class ajaxcomment
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

            Dim psize, pnum As Integer
            Try
                psize = CInt(Request.QueryString("psize"))
            Catch ex As Exception
                psize = 10
            End Try
            Try
                pnum = CInt(Request.QueryString("pnum"))
            Catch ex As Exception
                pnum = 1
            End Try
            Dim sql As String
            Dim pa As New ParaHelper
            'cn.addPara("@psize", SqlDbType.Int, 4, psize)
            'cn.addPara("@pnum", SqlDbType.Int, 4, pnum)
            pa.addInt("@psize", psize)
            pa.addInt("@pnum", pnum)

            sql = "select_ajax_comment_page"
            comments.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
            comments.DataBind()
        End Sub

    End Class
End Namespace