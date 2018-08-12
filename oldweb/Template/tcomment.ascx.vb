Imports SqlConnHelper
Namespace Blogwind
  Partial Class tcomment
    Inherits System.Web.UI.UserControl
    'Protected WithEvents dlCate As System.Web.UI.WebControls.Repeater
    'Protected archive As Repeater
    'Protected dlComment As Repeater
    'Protected links As Repeater
    Protected path As String
    Protected count, cateid As Integer
    Protected viewer, hp, uid, uinfo, title, content, adddate, catename, blogname, subject, tags As String
    Protected nick As String
    Protected ph As pageHelper

    Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
      ph = New pageHelper(Me.Session, Me.Request)
      Dim sql As String
      Dim Blogger_id As Integer
      viewer = ph.nickname
      hp = ph.url

      ID = CInt(Request("article_id"))

      path = Request.QueryString("blogger")
      If path.ToLower = "wuvist" And Request.RawUrl.IndexOf("cn.blogwind.com") > 0 Then
        Response.Redirect("http://wiki.blogwind.com/doku.php?id=Wuvist")
      End If

      uid = path

      Dim bloggerdr, blogdr As DataRow
      Dim dr As DataRow
      Dim dt As DataTable
      Dim tmp, IpPic As String
      Dim i As Integer

      sql = "select_doc"
      Dim pa As New ParaHelper
      pa.addVarchar("@blogger", path, 50)
      pa.addInt("@id", ID)
      blogdr = connHelper.exeDataRow(sql, CommandType.StoredProcedure, pa.parameters)
      If blogdr Is Nothing Then
        Response.Redirect("/404.html")
        Exit Sub
      End If

      sql = "SELECT * FROM Blogger WHERE id= '" & path & "'"
      bloggerdr = connHelper.exeDataRow(sql, CommandType.Text, Nothing)
      Blogger_id = bloggerdr.Item(0)
      nick = bloggerdr.Item(2)

      sql = "select_comment " & ID
      dt = connHelper.retrieve(sql, CommandType.Text, Nothing)

      Dim ct As ini.CommentorType
      For Each dr In dt.Rows
        If Not TypeOf dr.Item("ip") Is DBNull Then
          tmp = dr.Item("ip")
          IpPic = "<img src='/images/flag/" & GeoIP.LookupCountryCode(tmp) & ".png'> "
          i = tmp.LastIndexOf(".")
          tmp = tmp.Substring(0, i) & ".*"
        Else
          tmp = "*.*.*.*"
          IpPic = ""
        End If

        If TypeOf dr.Item("Reveal") Is DBNull Then
          ct = ini.CommentorType.Visitor
          dr.Item("author") = "<img src='/images/u" & CInt(ct) & ".png' title='" & util.GetCTDes(ct) & "，IP:" & tmp & "'> " & dr.Item("author")
        Else
          If TypeOf dr.Item("bf") Is DBNull And TypeOf dr.Item("fb") Is DBNull Then
            If dr.Item("blogger") = path Then
              ct = ini.CommentorType.Blogger
            Else
              If dr.Item("Reveal") = 0 Then
                ct = ini.CommentorType.UnrevealUser
              Else
                ct = ini.CommentorType.RegisterUser
              End If
            End If
          Else
            If TypeOf dr.Item("bf") Is DBNull Then
              ct = ini.CommentorType.BloggerFans
            Else
              If TypeOf dr.Item("fb") Is DBNull Then
                ct = ini.CommentorType.BloggerFriend
              Else
                If dr.Item("blogger") = path Then
                  ct = ini.CommentorType.Blogger
                Else
                  ct = ini.CommentorType.BloggerGoodFriend
                End If
              End If
            End If
          End If

          If dr.Item("Reveal") = 0 Then
            dr.Item("author") = "<img src='/images/u" & CInt(ct) & ".png' title='" & util.GetCTDes(ct) & "，IP:" & tmp & "'> " & dr.Item("author")
          Else
            dr.Item("author") = "<a href='http://www.blogwind.com/" & dr.Item("blogger") & "'><img border='0' src='/images/u" & CInt(ct) & ".png' title='" & util.GetCTDes(ct) & "，IP:" & tmp & "'></a> " & dr.Item("author")
          End If
        End If
        dr.Item("author") = IpPic & dr.Item("author")
      Next
      dlComment.DataSource = dt
      dlComment.DataBind()

      sql = "select UserDefineCategory.* from UserDefineCategory WHERE Blogger=" & Blogger_id
      dlCate.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
      dlCate.DataBind()

      sql = "select url,link from links WHERE reveal=1 and blogger= " & Blogger_id & " order by link"
      links.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
      links.DataBind()


      sql = "select tag from RR_tag where tid in (select tid from blog_tag where aid=" & ID & ") order by tag"
      dt = connHelper.retrieve(sql, CommandType.Text, Nothing)
      tags = ID & "|"

      For Each dr In dt.Rows
        tags = tags & dr(0) & "|"
      Next
      If tags <> "" Then
        tags = tags.TrimEnd("|")
      End If



      sql = "update_user_counter '" & path & "'"
      count = connHelper.exeInteger(sql, CommandType.Text, Nothing, False)

      uid = bloggerdr.Item(1)
      uinfo = bloggerdr.Item("intro")
      title = "【" & Trim(blogdr.Item("cate_name")) & "】" & blogdr.Item("Title")
      subject = blogdr.Item("Title")
      catename = Trim(blogdr.Item("cate_name"))
      path = bloggerdr.Item("blogname")
      blogname = bloggerdr.Item("blogname")
      content = blogdr.Item("Content")
      adddate = blogdr.Item("add_date")
      cateid = blogdr.Item("cate_index")

      Dim odate As Date
      dt = New DataTable("archive")
      Dim dc As New DataColumn


      Dim y, m, ynow, mnow As Integer
      sql = "select min(add_date) from Articles where blogger=" & Blogger_id

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
