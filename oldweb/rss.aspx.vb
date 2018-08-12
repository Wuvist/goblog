Imports System.Xml
Imports SqlConnHelper
Namespace blogwind
  Partial Class rss
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
      Dim nick, blogger, sql, blogname, uinfo As String
      Dim Blogger_id As Integer
      Dim dt As New DataTable
      blogger = Request.QueryString("blogger")

      If blogger.indexOf("'") > -1 Then
	    Response.Write("'")
        Response.End()
      End If

      sql = "SELECT * FROM Blogger WHERE id= '" & blogger & "'"
      dt = connHelper.retrieve(sql, CommandType.Text, Nothing)

      If dt.Rows.Count = 0 Then
        dt.Dispose()
        Response.End()
      End If

      Blogger_id = dt.Rows(0).Item(0)
      blogname = dt.Rows(0).Item("blogname")
      nick = dt.Rows(0).Item("nick")
      If nick = "" Then nick = blogger
      uinfo = dt.Rows(0).Item("intro")

      sql = "select_doc_page"
      Dim pa As New ParaHelper
      pa.addInt("@bloggerid", Blogger_id)
      pa.addInt("@page", 1)

      dt = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)

      Dim _xml As New XmlDocument
      Dim monode, elenode, tmpNode, tmpnode2 As XmlNode
      Dim xmldecl As XmlNode
      Dim att As XmlAttribute

      xmldecl = _xml.CreateXmlDeclaration("1.0", "utf-8", String.Empty)
      Dim root As XmlElement
      monode = _xml.CreateElement("rss")

      'version="2.0" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:trackback="http://madskills.com/public/xml/rss/module/trackback/" xmlns:wfw="http://wellformedweb.org/CommentAPI/" xmlns:slash="http://purl.org/rss/1.0/modules/slash/
      att = _xml.CreateAttribute("version")
      att.Value = "2.0"
      monode.Attributes.Append(att)

      att = _xml.CreateAttribute("xmlns:dc")
      att.Value = "http://purl.org/dc/elements/1.1"
      monode.Attributes.Append(att)

      att = _xml.CreateAttribute("xmlns:slash")
      att.Value = "http://purl.org/rss/1.0/modules/slash/"
      monode.Attributes.Append(att)

      elenode = _xml.CreateNode(XmlNodeType.Element, "channel", "")


      tmpNode = _xml.CreateNode(XmlNodeType.Element, "title", "")
      tmpNode.InnerText = blogname
      elenode.AppendChild(tmpNode)

      tmpNode = _xml.CreateNode(XmlNodeType.Element, "link", "")
      tmpNode.InnerText = "http://www.blogwind.com/" & blogger & "/"
      elenode.AppendChild(tmpNode)

      tmpNode = _xml.CreateNode(XmlNodeType.Element, "description", "")
      tmpNode.InnerText = uinfo
      elenode.AppendChild(tmpNode)

      tmpNode = _xml.CreateNode(XmlNodeType.Element, "webMaster", "")
      tmpNode.InnerText = "info@blogwind.com (Blogwind)"
      elenode.AppendChild(tmpNode)

      tmpNode = _xml.CreateNode(XmlNodeType.Element, "managingEditor", "")
      tmpNode.InnerText = nick
      elenode.AppendChild(tmpNode)

      tmpNode = _xml.CreateNode(XmlNodeType.Element, "generator", "")
      tmpNode.InnerText = "Blogwind RSS Version 0.1"
      elenode.AppendChild(tmpNode)

      Dim url As String

      For Each dr As DataRow In dt.Rows

        url = "http://www.blogwind.com/" & blogger & "/" & dr("index") & ".shtml"

        tmpNode = _xml.CreateNode(XmlNodeType.Element, "item", "")

        'tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "creator", "dc")
        'tmpnode2.InnerText = nick
        'tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "title", "")
        tmpnode2.InnerText = dr("title")
        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "link", "")
        tmpnode2.InnerText = url
        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "author", "")
        tmpnode2.InnerText = nick
        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "pubDate", "")
        tmpnode2.InnerText = String.Format("{0:R}", CDate(dr("add_date")).ToUniversalTime)
        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "guid", "")
        tmpnode2.InnerText = url
        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "comments", "")
        tmpnode2.InnerText = url
        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "slash:comments", "comments")
        tmpnode2.InnerText = dr("Comment")
        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "category", "")
        tmpnode2.InnerText = dr("cate")
        att = _xml.CreateAttribute("domain")
        att.Value = "http://www.blogwind.com/" & blogger & "/cate" & dr("cid") & ".shtml"
        'tmpnode2.Attributes.Append(att)
        tmpNode.AppendChild(tmpnode2)


        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "description", "")
        tmpnode2.InnerText = dr("content")
        tmpNode.AppendChild(tmpnode2)

        elenode.AppendChild(tmpNode)
      Next
If False Then
      sql = "select top 10 * from comment where blogger=" & Blogger_id.ToString & " order by [index] desc"
      dt = connHelper.retrieve(sql, CommandType.Text, Nothing)

      For Each dr As DataRow In dt.Rows
        url = "http://www.blogwind.com/" & blogger & "/" & dr("article") & ".shtml#blogc" & dr("index")

        tmpNode = _xml.CreateNode(XmlNodeType.Element, "item", "")

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "title", "")
        tmpnode2.InnerText = dr("title")
        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "link", "")
        tmpnode2.InnerText = url
        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "author", "")
        If TypeOf dr("author") Is DBNull Then
          tmpnode2.InnerText = "无名氏"
        Else
          tmpnode2.InnerText = dr("author")
        End If

        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "pubDate", "")
        tmpnode2.InnerText = String.Format("{0:R}", CDate(dr("add_date")).ToUniversalTime)
        tmpNode.AppendChild(tmpnode2)

        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "guid", "")
        tmpnode2.InnerText = url
        tmpNode.AppendChild(tmpnode2)


        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "category", "")
        tmpnode2.InnerText = "网志回复"
        att = _xml.CreateAttribute("domain")
        att.Value = "http://www.blogwind.com/" & blogger & "/" & dr("article") & ".shtml"
        'tmpnode2.Attributes.Append(att)
        tmpNode.AppendChild(tmpnode2)


        tmpnode2 = _xml.CreateNode(XmlNodeType.Element, "description", "")
        tmpnode2.InnerText = dr("content")
        tmpNode.AppendChild(tmpnode2)

        elenode.AppendChild(tmpNode)
      Next

      dt.Dispose()
End If
      monode.AppendChild(elenode)

      _xml.AppendChild(monode)
      root = _xml.DocumentElement

      Dim pi As XmlProcessingInstruction
      pi = _xml.CreateProcessingInstruction("xml-stylesheet", "type=""text/xsl"" href=""rss2full.xsl"" media=""screen""")

      _xml.InsertBefore(xmldecl, root)
      _xml.InsertBefore(pi, root)
      pi = _xml.CreateProcessingInstruction("xml-stylesheet", "type=""text/css"" href=""itemcontent.css"" media=""screen""")
      _xml.InsertBefore(pi, root)
      '            Response.ContentEncoding = System.Text.Encoding.GetEncoding("gb2312")
      'Response.ContentType = "application/rss+xml"
      Response.ContentType = "text/xml"
      Response.Write(_xml.OuterXml)
    End Sub
  End Class
End Namespace