<%@ WebHandler Language="VB" Class="xmlrpc" %>
<%@ Assembly Name="CookComputing.XmlRpcV2" %>

Imports CookComputing.XmlRpc
Imports SqlConnHelper
Imports System.Drawing
Imports System.Drawing.Imaging

<XmlRpcService()> _
Public Class xmlrpc
    Inherits XmlRpcService
    Implements CookComputing.MetaWeblog.IMetaWeblog
    
    Public Function editPost(ByVal postid As String, ByVal username As String, ByVal password As String, ByVal post As CookComputing.MetaWeblog.Post, ByVal publish As Boolean) As Object Implements CookComputing.MetaWeblog.IMetaWeblog.editPost
        Dim aid As Integer
        Dim BloggerId As Integer
        Dim pa As New ParaHelper
        BloggerId = Blogger.Auth(username, password)
        
        aid = CInt(postid)
        pa.addVarchar("@title", post.title, 50)
        pa.addVarchar("@blogger", username, 50)
        pa.addNText("@content", post.description.replace("127.0.0.1:2323", "blog.cctvcnm.com"))
        pa.addInt("@id", aid)
        If post.categories.Length = 0 Then
            pa.addInt("@cate", Categories.GetBloggerDefaultCategoryId(BloggerId))
        Else
            Dim CateId As Integer
            CateId = Categories.GetBloggerCategoryIdByName(BloggerId, post.categories(0))
            If CateId = 0 Then
                CateId = Categories.GetBloggerDefaultCategoryId(BloggerId)
            End If
            pa.addInt("@cate", CateId)
        End If

        connHelper.exeNonQuery("update_doc", CommandType.StoredProcedure, pa.parameters)
        
        Return aid.ToString
    End Function

    Public Function getCategories(ByVal blogid As String, ByVal username As String, ByVal password As String) As CookComputing.MetaWeblog.CategoryInfo() Implements CookComputing.MetaWeblog.IMetaWeblog.getCategories
        Dim BloggerId As Integer
        Dim sql As String
        Dim pa As New ParaHelper
        Dim i As Integer

        BloggerId = Blogger.Auth(username, password)
        
        Dim dt As DataTable
        sql = "select [index],cate from UserDefineCategory WHERE Blogger=" & BloggerId
        dt = connHelper.retrieve(sql, CommandType.Text, Nothing)
        
        Dim cs(dt.Rows.Count - 1) As CookComputing.MetaWeblog.CategoryInfo
        i = 0
        For Each dr As DataRow In dt.Rows
            cs(i) = New CookComputing.MetaWeblog.CategoryInfo
            cs(i).categoryid = dr(0)
            cs(i).title = dr(1)
            cs(i).description = dr(1)
            cs(i).htmlUrl = "http://www.blogwind.com/" & username & "/cate" & dr(0) & ".shtml"
            cs(i).rssUrl = "http://www.blogwind.com/rss.aspx?blogger=" & username & "&cate" & dr(0)
            i += 1
        Next

        Return cs
    End Function

    Public Function getPost(ByVal postid As String, ByVal username As String, ByVal password As String) As CookComputing.MetaWeblog.Post Implements CookComputing.MetaWeblog.IMetaWeblog.getPost
        Dim p As New CookComputing.MetaWeblog.Post
        p.title = ""
        Return p
    End Function

    Public Function getRecentPosts(ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal numberOfPosts As Integer) As CookComputing.MetaWeblog.Post() Implements CookComputing.MetaWeblog.IMetaWeblog.getRecentPosts
        Dim BloggerId As Integer
        Dim sql As String
        Dim pa As New ParaHelper
        Dim i As Integer

        BloggerId = Blogger.Auth(username, password)
        
        Dim dt As DataTable
        sql = "select  top " & numberOfPosts & " articles.[index], " _
       & " CASE Articles.title when '' then 'ÎÞÌâ'  else Articles.title end as title, " _
       & "articles.content,articles.add_date from Articles where blogger=" & BloggerId & " order by [index] DESC"
        
        dt = connHelper.retrieve(sql, CommandType.Text, Nothing)
        
        Dim Blogs(dt.Rows.Count - 1) As CookComputing.MetaWeblog.Post
        i = 0

        For Each dr As DataRow In dt.Rows
            Blogs(i) = New CookComputing.MetaWeblog.Post
            Blogs(i).postid = dr(0)
            Blogs(i).title = dr(1)
            Blogs(i).description = dr(2)
            Blogs(i).dateCreated = dr(3)
            Blogs(i).permalink = "http://www.blogwind.com/" & username & "/" & dr(0) & ".shtml"
 
            i += 1
        Next

        Return Blogs
        
    End Function

    Public Function newMediaObject(ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal file As CookComputing.MetaWeblog.FileData) As CookComputing.MetaWeblog.UrlData Implements CookComputing.MetaWeblog.IMetaWeblog.newMediaObject
        Dim BloggerId As Integer
        BloggerId = Blogger.Auth(username, password)
        
        
        Dim u As New CookComputing.MetaWeblog.UrlData
        u.url = ""
        If Not file.name.ToLower.EndsWith(".jpg") Then
            Return Nothing
        End If
        
        Dim fpath As String = System.Web.HttpContext.Current.Server.MapPath("images/pic")
        Dim fname As String = Now.Ticks.ToString & ".jpg"
        Dim imgpath As String
        
        imgpath = System.IO.Path.Combine(fpath & "/up", System.IO.Path.GetFileName(fname))

        Dim fio As New System.IO.FileInfo(imgpath)
        Dim im As Image
        
        Try
            im = System.Drawing.Image.FromStream(New System.IO.MemoryStream(file.bits))
            im.Save(imgpath)
            Dim pa As New ParaHelper
            pa.addNVarchar("@blogger", username, 50)
            pa.addNVarchar("@path", fname, 50)
            pa.addInt("@file_size", fio.Length / 1000)
            pa.addInt("@width", im.Height)
            pa.addInt("@height", im.Width)
            connHelper.exeNonQuery("insert_blog_pic", CommandType.StoredProcedure, pa.parameters)
        Catch ex As Exception
            System.IO.File.Delete(imgpath)
            Return Nothing
        Finally
            If Not im Is Nothing Then
                im.Dispose()
            End If
        End Try

        Dim host As String
	host = System.Web.HttpContext.Current.Request.Headers("host")
	If host.indexOf("127.0.") > 0 Then
		host = "blog.cctvcnm.com"
	End If

        u.url = "http://" & host & "/images/pic/up/" & fname
        'u.url = "/images/pic/up/" & fname
        Return u
    End Function

    <XmlRpcMethod("metaWeblog.newMediaObject", _
   Description:="Makes a new file to a designated blog using the " _
   & "metaWeblog API. Returns url as a string of a struct.")> _
     Public Function newMediaObject(ByVal blogid As Object, ByVal username As String, ByVal password As String, ByVal file As CookComputing.MetaWeblog.FileData) As CookComputing.MetaWeblog.UrlData
        Return newMediaObject(blogid.ToString, username, password, file)
    End Function

    Public Function newPost(ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal post As CookComputing.MetaWeblog.Post, ByVal publish As Boolean) As String Implements CookComputing.MetaWeblog.IMetaWeblog.newPost
        Dim aid As Integer
        Dim BloggerId As Integer
        Dim pa As New ParaHelper
        BloggerId = Blogger.Auth(username, password)
        pa.addVarchar("@title", post.title, 50)
        pa.addVarchar("@blogger", username, 50)
        pa.addNText("@content", post.description.replace("127.0.0.1:2323", "blog.cctvcnm.com"))
        If post.categories Is Nothing OrElse post.categories.Length = 0 Then
            pa.addInt("@cate", Categories.GetBloggerDefaultCategoryId(BloggerId))
        Else
            pa.addInt("@cate", Categories.GetBloggerCategoryIdByName(BloggerId, post.categories(0)))
        End If

        pa.addNVarchar("@tags", "", 50)
        aid = connHelper.exeInteger("insert_doc", CommandType.StoredProcedure, pa.parameters, False)

        Return aid.ToString
    End Function
    
    <XmlRpcMethod("blogger.getUsersBlogs", _
    Description:="Returns information on all the blogs a given user " _
    & "is a member.")> _
 Public Function getUsersBlogs(ByVal appKey As String, ByVal username As String, ByVal password As String) As CookComputing.Blogger.BlogInfo()
        Dim b(0) As CookComputing.Blogger.BlogInfo
        Dim bb As New CookComputing.Blogger.BlogInfo
        
        Dim BloggerID As Integer
        BloggerID = Blogger.Auth(username, password)
        bb.blogid = "1"
        bb.blogName = Blogger.GetBlogName(BloggerID)
        bb.url = "http://www.blogwind.com/" & username
        b(0) = bb
        Return b

    End Function
End Class