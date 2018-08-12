Imports SqlConnHelper
Partial Class tag
    Inherits System.Web.UI.Page
    Public tw As String
    Public nick, blogger, paging As String
    Private pagesize As Integer = 20

    Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
        Dim u, sql, tmp As String
        Dim tid, bloggerid As Integer
        Dim i, titem, start As Integer
        Dim dr As DataRow
        Dim pa As New ParaHelper
        tmp = Request.RawUrl
        i = InStr(tmp, "tag=")
        If i > 0 Then
            tw = Mid(Request.RawUrl, i + 4, Request.RawUrl.Length - i + 4)
            tw = HttpUtility.UrlDecode(tw, Encoding.UTF8)

            pa.addNVarchar("@tag", tw, 50)
            dr = connHelper.exeDataRow("select tid, blogs from RR_tag where tag=@tag", CommandType.Text, pa.parameters)
            If Not dr Is Nothing Then
                tid = dr.Item("tid")
                titem = dr.Item("blogs")
            End If
        End If


        If String.IsNullOrEmpty(tw) Or tid = 0 Then

        Else
            u = Request.QueryString("blogger")
            start = Request.QueryString("start")
            If start < 1 Then
                start = 1
            End If

            If String.IsNullOrEmpty(u) Then
                sql = "select id, nick from blogger where reveal=1 and [index] in ("
                sql &= "select distinct blogger from Articles where [index]  in (select aid from blog_tag where tid=" & tid & "))"
                bloggers.DataSource = connHelper.retrieve(sql, CommandType.Text, Nothing)
                bloggers.DataBind()
                sql = "select_doclist_by_tid"
                pa.clearPara()

                paging = ""
                If start > 1 Then
                    paging = "<a href='tag.aspx?start=" & (start - pagesize) & "&tag=" & Server.UrlPathEncode(tw) & "'>上一页</a>　　"
                End If
                If start + pagesize - 1 < titem Then
                    paging &= "<a href='tag.aspx?start=" & (start + pagesize) & "&tag=" & Server.UrlPathEncode(tw) & "'>下一页</a>"
                End If
            Else
                pa.clearPara()
                pa.addNVarchar("id", u, 50)
                dr = connHelper.exeDataRow("select [index],nick from blogger where id=@id", CommandType.Text, pa.parameters)
                If dr Is Nothing Then
                    pa.addInt("blogger", 0)
                Else
                    blogger = u
                    bloggerid = dr.Item("index")
                    nick = dr.Item("nick")
                    pa.clearPara()
                    pa.addInt("bloggerid", bloggerid)
                End If

                sql = "select_doclist_by_tid_bloggerid"
            End If

            pa.addInt("tid", tid)
            pa.addInt("start", start)
            pa.addInt("end", start + pagesize - 1)
            docs.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
            docs.DataBind()
        End If
    End Sub

End Class
