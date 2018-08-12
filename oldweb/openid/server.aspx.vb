Imports DotNetOpenId.Provider

Partial Class openid_server
    Inherits System.Web.UI.Page

    Protected Sub Page_Load(ByVal sender As Object, ByVal e As System.EventArgs) Handles Me.Load
        serverEndpointUrl.Text = Request.Url.ToString()
    End Sub

    Protected Sub provider_AuthenticationChallenge(ByVal sender As Object, ByVal e As AuthenticationChallengeEventArgs)
        Response.Redirect("/openid/decide.aspx", True)
    End Sub
End Class
