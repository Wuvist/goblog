Imports SqlConnHelper
Namespace blogwind
    Partial Class bmfeed
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
            Dim sql As String

            sql = "RR_select_bookmark_by_blogger"
            'cn.newda(sql)
            'cn.comm.SelectCommand.CommandType = CommandType.StoredProcedure
            'cn.comm.SelectCommand.Parameters.Add("@blogger", SqlDbType.VarChar, 50)
            'cn.comm.SelectCommand.Parameters.Item(0).Value = Request.QueryString("blogger")
            Dim pa As New ParaHelper
            pa.addVarchar("@blogger", Request.QueryString("blogger"), 50)
            fed.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
            fed.DataBind()
        End Sub
    End Class
End Namespace