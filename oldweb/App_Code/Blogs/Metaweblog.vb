Imports Microsoft.VisualBasic

Public Class Metaweblog
    Inherits CookComputing.XmlRpc.XmlRpcClientProtocol
    Implements CookComputing.MetaWeblog.IMetaWeblog, CookComputing.Blogger.IBlogger

    Public Class Urls
        Public Shared MSNSpace As String = "https://storage.msn.com/storageservice/MetaWeblog.rpc"
    End Class

    Private Shared ReadOnly Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)


    Public Function editPost(ByVal postid As String, ByVal username As String, ByVal password As String, ByVal post As CookComputing.MetaWeblog.Post, ByVal publish As Boolean) As Object Implements CookComputing.MetaWeblog.IMetaWeblog.editPost

    End Function

    <CookComputing.XmlRpc.XmlRpcMethod("metaWeblog.getCategories")> _
    Public Function getCategories(ByVal blogid As String, ByVal username As String, ByVal password As String) As CookComputing.MetaWeblog.CategoryInfo() Implements CookComputing.MetaWeblog.IMetaWeblog.getCategories
        Return Me.Invoke("getCategories", New Object() {blogid, username, password})
    End Function

    Public Function getPost(ByVal postid As String, ByVal username As String, ByVal password As String) As CookComputing.MetaWeblog.Post Implements CookComputing.MetaWeblog.IMetaWeblog.getPost

    End Function

    Public Function getRecentPosts(ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal numberOfPosts As Integer) As CookComputing.MetaWeblog.Post() Implements CookComputing.MetaWeblog.IMetaWeblog.getRecentPosts

    End Function

    Public Function newMediaObject(ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal file As CookComputing.MetaWeblog.FileData) As CookComputing.MetaWeblog.UrlData Implements CookComputing.MetaWeblog.IMetaWeblog.newMediaObject

    End Function

    <CookComputing.XmlRpc.XmlRpcMethod("metaWeblog.newPost")> _
    Public Function newPost(ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal post As CookComputing.MetaWeblog.Post, ByVal publish As Boolean) As String Implements CookComputing.MetaWeblog.IMetaWeblog.newPost
        Return Me.Invoke("newPost", New Object() {blogid, username, password, post, publish})
    End Function


    Public Shared Function Validate(ByVal ServerHost As String, ByVal username As String, ByVal pwd As String) As CookComputing.Blogger.BlogInfo()
        Dim c As New Metaweblog
        c.Url = ServerHost
        Try
            Return c.getUsersBlogs("", username, pwd)
        Catch ex As CookComputing.XmlRpc.XmlRpcFaultException
            Return Nothing
        End Try
    End Function

    Public Shared Sub Post(ByVal Acc As Blogwind.WindMill.Account, ByVal Title As String, ByVal Content As String, ByVal Category As String)
        Dim c As New Metaweblog
        Dim blog As New CookComputing.MetaWeblog.Post
        blog.title = Title
        blog.description = Content
        blog.categories = New String() {Category}
        blog.dateCreated = Now
        c.Url = Acc.ServerHost
        c.newPost(Acc.Blogid, Acc.UserName, Cryptor.Decrypt(Acc.Pwd), blog, True)
    End Sub

    Public Function deletePost(ByVal appKey As String, ByVal postid As String, ByVal username As String, ByVal password As String, ByVal publish As Boolean) As Boolean Implements CookComputing.Blogger.IBlogger.deletePost

    End Function

    Public Function editPost1(ByVal appKey As String, ByVal postid As String, ByVal username As String, ByVal password As String, ByVal content As String, ByVal publish As Boolean) As Object Implements CookComputing.Blogger.IBlogger.editPost

    End Function

    Public Function getCategories1(ByVal blogid As String, ByVal username As String, ByVal password As String) As CookComputing.Blogger.Category() Implements CookComputing.Blogger.IBlogger.getCategories

    End Function

    Public Function getPost1(ByVal appKey As String, ByVal postid As String, ByVal username As String, ByVal password As String) As CookComputing.Blogger.Post Implements CookComputing.Blogger.IBlogger.getPost

    End Function

    Public Function getRecentPosts1(ByVal appKey As String, ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal numberOfPosts As Integer) As CookComputing.Blogger.Post() Implements CookComputing.Blogger.IBlogger.getRecentPosts

    End Function

    Public Function getTemplate(ByVal appKey As String, ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal templateType As String) As String Implements CookComputing.Blogger.IBlogger.getTemplate

    End Function

    Public Function getUserInfo(ByVal appKey As String, ByVal username As String, ByVal password As String) As CookComputing.Blogger.UserInfo Implements CookComputing.Blogger.IBlogger.getUserInfo

    End Function

    <CookComputing.XmlRpc.XmlRpcMethod("blogger.getUsersBlogs")> _
    Public Function getUsersBlogs(ByVal appKey As String, ByVal username As String, ByVal password As String) As CookComputing.Blogger.BlogInfo() Implements CookComputing.Blogger.IBlogger.getUsersBlogs
        Return Me.Invoke("getUsersBlogs", New Object() {appKey, username, password})
    End Function

    Public Function newPost1(ByVal appKey As String, ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal content As String, ByVal publish As Boolean) As String Implements CookComputing.Blogger.IBlogger.newPost

    End Function

    Public Function setTemplate(ByVal appKey As String, ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal template As String, ByVal templateType As String) As Boolean Implements CookComputing.Blogger.IBlogger.setTemplate

    End Function
End Class
