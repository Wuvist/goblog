Imports System.Xml
Imports SqlConnHelper

Partial Class system_admin_backup
    Inherits System.Web.UI.Page
    Protected ph As pageHelper

    Protected Sub Page_Load(ByVal sender As Object, ByVal e As System.EventArgs) Handles Me.Load
        ph = New pageHelper(Me.Session, Me.Request)
    End Sub

    Protected Sub GenBackup_Click(ByVal sender As Object, ByVal e As System.EventArgs) Handles GenBackup.Click
        Dim nick, blogger, sql, blogname, uinfo As String
        Dim Blogger_id As Integer
        Dim dt As New DataTable
        Blogger_id = Session("bloggerID")


        sql = "SELECT * FROM Bloggers WHERE id= " & Blogger_id
        dt = connHelper.retrieve(sql, CommandType.Text, Nothing)

        If dt.Rows.Count = 0 Then
            dt.Dispose()
            Response.End()
        End If

        blogger = dt.Rows(0).Item("user_name")
        blogname = dt.Rows(0).Item("blogname")
        nick = dt.Rows(0).Item("nick")
        If nick = "" Then nick = blogger
        uinfo = dt.Rows(0).Item("intro")

        sql = "select A.[index], CASE A.title when '' then '无题'  else A.title end as title, A.content, A.add_date, "
        sql = sql & "A.Comment, u.cate, u.[index] as cid from Articles A inner join UserDefineCategory u on A.cate = u.[index]  WHERE"
        sql = sql & " A.blogger=@blogger and A.add_date>=@sd and A.add_date<=@ed order by a.[index]"
        Dim pa As New ParaHelper
        pa.addInt("@blogger", Blogger_id)
        pa.addDateTime("@sd", cdate(sd.Text))
        pa.addDateTime("@ed", CDate(ed.Text))

        dt = connHelper.retrieve(sql, CommandType.Text, pa.parameters)

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


        monode.AppendChild(elenode)

        _xml.AppendChild(monode)
        root = _xml.DocumentElement

        Dim pi As XmlProcessingInstruction
        pi = _xml.CreateProcessingInstruction("xml-stylesheet", "type=""text/xsl"" href=""http://www.blogwind.com/rss2full.xsl"" media=""screen""")

        _xml.InsertBefore(xmldecl, root)
        _xml.InsertBefore(pi, root)
        pi = _xml.CreateProcessingInstruction("xml-stylesheet", "type=""text/css"" href=""http://www.blogwind.com/itemcontent.css"" media=""screen""")
        _xml.InsertBefore(pi, root)

        sql = "/images/backup/" & Blogger_id & "_" & Now.Ticks & ".xml"
        _xml.Save(Server.MapPath(sql))
        hint.Text = "您自 " & sd.Text & " 到 " & ed.Text & " 共" & dt.Rows.Count & " 篇网志已经备份在 <a href='" & sql & "'>这里</a> 请点击右键，选择 目标另存为 下载保存～"
    End Sub
End Class
