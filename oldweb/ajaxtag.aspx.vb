Imports SqlConnHelper
Namespace blogwind
  Partial Class ajaxtag
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
      Dim ph As New pageHelper(Me.Session, Me.Request)
      'Put user code to initialize the page here
      If Session("blogger") <> "" Then
        Dim aid As Integer = CInt(Request("aid"))
        Dim action As String = Request("action")
        Dim tags As String = Request("tags")

        Dim r As Integer
        Dim pa As New ParaHelper
        'cn.addPara("@blogger", SqlDbType.VarChar, 50, Session("blogger"))
        'cn.addPara("@aid", SqlDbType.Int, 4, aid)
        pa.addVarchar("@blogger", Session("blogger"), 50)
        pa.addInt("@aid", aid)

        If action = "add" Then
          tags = tags.Replace("，", ",")
          tags = tags.Replace(";", ",")
          tags = tags.Replace("；", ",")
          tags = tags.Replace("'", "")
          tags = tags.Replace("""", "")
          tags = tags.Replace("|", "")
          '加多一个逗号，方便存储过程切分
          tags = tags & ","
          'cn.addPara("@tags", SqlDbType.NVarChar, 50, tags)
          pa.addNVarchar("@tags", tags, 50)
          r = connHelper.exeScalar("add_tags", CommandType.StoredProcedure, pa.parameters)
        ElseIf action = "delete" Then
          'cn.addPara("@tag", SqlDbType.NVarChar, 50, tags)
          pa.addNVarchar("@tag", tags, 50)
          r = connHelper.exeScalar("del_tag", CommandType.StoredProcedure, pa.parameters)
        End If

        If r = 0 Then
          Response.Write(r)
          Dim dt As DataTable
          Dim dr As DataRow
          dt = connHelper.retrieve("select tag from RR_tag where tid in (select tid from blog_tag where aid=" & aid & ") order by tag", CommandType.Text, Nothing)
          Response.Write(r & aid & "|")
          For Each dr In dt.Rows
            Response.Write(dr(0) & "|")
          Next
        Else
          Response.Write(r)
        End If
      Else
        Response.Write(2)
      End If
    End Sub

  End Class
End Namespace