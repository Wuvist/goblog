Partial Class mills
    Inherits AuthPage
    Protected accs As Blogwind.WindMill.AccountCollection
    Private Shared ReadOnly Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)

    Protected Sub Page_Load(ByVal sender As Object, ByVal e As System.EventArgs) Handles Me.Load
        Dim query As New SubSonic.Select(Blogwind.WindMill.Account.Schema.Provider)
        query.From("accounts")
        query.Where("user_id").IsEqualTo(Me.ph.bloggerID)
        accs = query.ExecuteAsCollection(Of Blogwind.WindMill.AccountCollection)()

        'Dim acc As New blogwind.WindMill.Account
        'acc.ServerHost = "https://storage.msn.com/storageservice/MetaWeblog.rpc"
        'acc.UserName = "wuvist"
        'acc.Pwd = Cryptor.Encrypt("tankeshi")

        'Metaweblog.Post(acc, "test", "i'm a test", "test")

        'Dim c As New Metaweblog
        'c.Url = "https://storage.msn.com/storageservice/MetaWeblog.rpc"
        'Dim blogs As CookComputing.Blogger.BlogInfo() = c.getUsersBlogs("d", "wuvist", "tankeshi")
        'For Each b As CookComputing.Blogger.BlogInfo In blogs
        '    Response.Write(b.blogid)
        '    Response.Write("<br>")
        'Next

    End Sub
End Class

