Imports SqlConnHelper
Namespace Blogwind
  Partial Class tcate
    Inherits System.Web.UI.UserControl
    'Protected dlCate As Repeater
    'Protected dlArticles As Repeater
    'Protected archive As Repeater
    'Protected links As Repeater
    Protected blogger_id As Integer
    Protected path As String
    Protected nick, uid, uinfo, catename, blogname As String
    Protected count As Integer
    Protected ph As pageHelper

    Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load

      Dim sql As String
      Dim bloggerdr As DataRow
      ph = New pageHelper(Me.Session, Me.Request)
      path = Request.QueryString("blogger")
      uid = path
      If path.ToLower = "wuvist" And Request.RawUrl.IndexOf("cn.blogwind.com") > 0 Then
        Response.Redirect("http://wiki.blogwind.com/doku.php?id=Wuvist")
      End If


      sql = "SELECT * FROM Blogger WHERE id='" & path & "'"

      bloggerdr = connHelper.exeDataRow(sql, CommandType.Text, Nothing)
      blogger_id = bloggerdr.Item(0)
      nick = bloggerdr.Item(2)
      path = bloggerdr.Item("blogname")
      blogname = bloggerdr.Item("blogname")
      uinfo = bloggerdr.Item("intro")
      uid = bloggerdr.Item(1)

      Dim y, m, cateid As Integer

      cateid = CInt(Request.QueryString("cate_id"))
      Dim pa As New ParaHelper
      If cateid = 0 Then
        Try
          y = CInt(Request.QueryString("y"))
          m = CInt(Request.QueryString("m"))
        Catch ex As System.InvalidCastException
          Response.StatusCode = 404
          Response.Redirect("/404.html")
          Exit Sub
        End Try

        sql = "select_doclist_by_bloggerid_date"

        pa.addInt("@bloggerid", blogger_id)
        pa.addInt("@y", y)
        pa.addInt("@m", m)
      Else
        Dim t As Boolean = False
        catename = connHelper.exeString("select cate from UserDefineCategory where [index]=" & cateid & " and blogger=" & blogger_id, CommandType.Text, Nothing, t)
        If t Then
          Response.StatusCode = 404
          Response.Redirect("/404.html")
          Exit Sub
        End If
        sql = "select_doclist_by_bloggerid_cate"
        pa.addInt("@bloggerid", blogger_id)
        pa.addInt("@cate", cateid)
      End If

      dlArticles.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
      dlArticles.DataBind()

      sql = "select UserDefineCategory.* from UserDefineCategory WHERE Blogger=" & blogger_id
      dlCate.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
      dlCate.DataBind()

      sql = "select url,link from links WHERE reveal=1 and blogger= " & blogger_id & " order by link"
      links.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
      links.DataBind()

      sql = "update_user_counter '" & uid & "'"
      count = connHelper.exeInteger(sql, CommandType.Text, Nothing, False)

      Dim odate As Date
      Dim dt As New DataTable("archive")
      Dim dc = New DataColumn
      Dim dr As DataRow

      Dim ynow, mnow As Integer
      sql = "select min(add_date) from Articles where blogger=" & blogger_id

      Try
        odate = connHelper.exeScalar(sql, CommandType.Text, Nothing)
      Catch ex As Exception
        odate = Now
      End Try

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