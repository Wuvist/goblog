Imports DotNetOpenId.Provider
Partial Class test
    Inherits System.Web.UI.Page

    Protected Sub Page_Load(ByVal sender As Object, ByVal e As System.EventArgs) Handles Me.Load
        'Dim doc As New HtmlAgilityPack.HtmlDocument
        'doc.Load(Server.MapPath("t.tpl"))
        'Dim i As Integer
        'For Each node As HtmlAgilityPack.HtmlNode In doc.DocumentNode.ChildNodes
        '    Response.Write(i)
        '    Response.Write(node.Name)
        '    Response.Write("<br>")
        '    For Each att As HtmlAgilityPack.HtmlAttribute In node.Attributes
        '        Response.Write(att.Name)
        '        Response.Write(att.Value)
        '    Next
        '    Response.Write("<br>")
        '    i += 1
        'Next
    'Response.Write(Blogwind.Photos.GetDayPath)
        'Metaweblog.Validate("https://storage.msn.com/storageservice/MetaWeblog.rpc", "wuvist", "tankeshi")
        Response.Write(Request.UserHostAddress)
    End Sub

    Protected Sub IdentityEndpoint20_NormalizeUri(ByVal sender As Object, ByVal e As IdentityEndpointNormalizationEventArgs)
        'Dim normalized As New UriBuilder(e.UserSuppliedIdentifier)
        'Dim username As String = Request.QueryString("username").TrimEnd("/"c).ToLowerInvariant()
        'username = username.Substring(0, 1).ToUpperInvariant() + username.Substring(1)
        'normalized.Path = "/user/" + username
        'normalized.Scheme = "http"
        '' for a real Provider, this should be HTTPS if supported.
        'e.NormalizedIdentifier = normalized.Uri
    End Sub

End Class
