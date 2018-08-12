Imports SqlConnHelper
Namespace Blogwind
  Partial Class tDefault
    Inherits System.Web.UI.UserControl
    Protected path, uid, uinfo As String
    Protected count As Integer
    Protected nick As String
    Protected ph As pageHelper
    'Protected dlCate As Repeater
    'Protected dlArticles As Repeater
    'Protected archive As Repeater
    'Protected links As Repeater


#Region " Web Form Designer Generated Code "

    'This call is required by the Web Form Designer.
    <System.Diagnostics.DebuggerStepThrough()> Private Sub InitializeComponent()

    End Sub
    Protected WithEvents skin_list As System.Web.UI.WebControls.Repeater


    Private Sub Page_Init(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Init
      'CODEGEN: This method call is required by the Web Form Designer
      'Do not modify it using the code editor.
      InitializeComponent()
    End Sub

#End Region

    Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
      'Dim ds As New DataSet
      ph = New pageHelper(Me.Session, Me.Request)
      Dim sql As String
      Dim Blogger_id As Integer

      Dim dr As DataRow
      Dim dt As DataTable

      'ds.Tables.Add("Blogger")
      'ds.Tables.Add("Articles")
      ''ds.Tables.Add("Comment")
      'ds.Tables.Add("UserDefineCategory")
      'ds.Tables.Add("links")
      path = Request.QueryString("blogger")
      If path.ToLower = "wuvist" And Request.RawUrl.IndexOf("cn.blogwind.com") > 0 Then
        Response.Redirect("http://wiki.blogwind.com/doku.php?id=Wuvist")
      End If

      sql = "SELECT * FROM Blogger WHERE id= '" & path & "'"
      dr = connHelper.exeDataRow(sql, CommandType.Text, Nothing)
      Blogger_id = dr.Item(0)
      nick = dr.Item(2)
      uinfo = dr.Item("intro")
      uid = dr.Item(1)
      path = dr.Item("blogname")
      Dim cpage As Integer
      cpage = CInt(Request.QueryString("page"))
      If cpage < 1 Then cpage = 1


      sql = "select_doc_page"
      Dim pa As New ParaHelper

      pa.addInt("@bloggerid", Blogger_id)
      pa.addInt("@page", cpage)

      dlArticles.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
      dlArticles.DataBind()
      'sql = "select Comment.title from Comment WHERE Comment.blogger=" & blogger_id
      'dbComm = New SqlDataAdapter(sql, conn)
      'dbComm.Fill(ds, "Comment")

      sql = "select UserDefineCategory.* from UserDefineCategory WHERE Blogger=" & Blogger_id
      dlCate.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
      dlCate.DataBind()

      sql = "select url,link from links WHERE reveal=1 and blogger= " & Blogger_id & " order by link"
      links.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
      links.DataBind()

      sql = "update_user_counter '" & uid & "'"
      count = connHelper.exeInteger(sql, CommandType.Text, Nothing, False)

      Dim odate As Date
      Dim dc As New DataColumn

      Dim y, m, ynow, mnow As Integer
      sql = "select min(add_date) from Articles where blogger=" & Blogger_id

      Try
        odate = connHelper.exeScalar(sql, CommandType.Text, Nothing)
      Catch ex As Exception
        odate = Now
      End Try

      dt = New DataTable("archive")
      y = odate.Year
      m = odate.Month
      ynow = Now.AddMonths(-1).Year
      mnow = Now.AddMonths(-1).Month

      dc.DataType = System.Type.GetType("System.Int32")
      dc.ColumnName = "year"
      dt.Columns.Add(dc)

      dc = New DataColumn
      dc.DataType = System.Type.GetType("System.Int32")
      dc.ColumnName = "month"
      dt.Columns.Add(dc)

      While (ynow = y And mnow >= m) Or (ynow > y)
        dr = dt.NewRow
        dr("year") = ynow
        dr("month") = mnow
        dt.Rows.Add(dr)
        mnow = mnow - 1
        If mnow < 1 Then
          mnow = 12
          ynow = ynow - 1
        End If
      End While

      archive.DataSource = dt
      archive.DataBind()
    End Sub
  End Class
End Namespace

