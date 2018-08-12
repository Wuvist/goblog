Imports DotNetOpenId.Extensions.SimpleRegistration
Imports DotNetOpenId.Provider
Partial Class openid_decide
    Inherits Blogwind.BWPage

    Protected ph As pageHelper

    Protected Sub Page_Load(ByVal sender As Object, ByVal e As System.EventArgs) Handles Me.Load
        ph = New pageHelper(Me.Session, Me.Request)

        If ProviderEndpoint.PendingAuthenticationRequest Is Nothing Then
            Response.Redirect("~/")
        End If

        'If ProviderEndpoint.PendingAuthenticationRequest.IsDirectedIdentity Then
        '    ProviderEndpoint.PendingAuthenticationRequest.LocalIdentifier = util.BuildIdentityUrl()
        'End If
        If ProviderEndpoint.PendingAuthenticationRequest.IsReturnUrlDiscoverable Then
            relyingPartyVerificationResultLabel.Text = "passed"
        Else
            relyingPartyVerificationResultLabel.Text = "failed"
        End If

        identityUrlLabel.Text = ProviderEndpoint.PendingAuthenticationRequest.LocalIdentifier.ToString()
        realmLabel.Text = ProviderEndpoint.PendingAuthenticationRequest.Realm.ToString()

        ' check that the logged in user is the same as the user requesting authentication to the consumer. If not, then log them out.
        Dim u As New Uri(ProviderEndpoint.PendingAuthenticationRequest.LocalIdentifier)
        If [String].Equals(ph.blogger, u.Segments(u.Segments.Length - 1), StringComparison.OrdinalIgnoreCase) Then
            ' if simple registration fields were used, then prompt the user for them
            Dim requestedFields As ClaimsRequest = ProviderEndpoint.PendingAuthenticationRequest.GetExtension(Of ClaimsRequest)()
            'If requestedFields IsNot Nothing Then
            '    Me.profileFields.Visible = True
            '    Me.profileFields.SetRequiredFieldsFromRequest(requestedFields)
            '    If Not IsPostBack Then
            '        Dim sregResponse As var = requestedFields.CreateResponse()
            '        sregResponse.Email = Membership.GetUser().Email
            '        Me.profileFields.SetOpenIdProfileFields(sregResponse)
            '    End If
            'End If
        Else
            Response.Redirect(Request.Url.AbsoluteUri)
        End If

    End Sub

    Protected Sub Yes_Click(ByVal sender As Object, ByVal e As System.EventArgs) Handles yes_button.Click
        Dim sregRequest As ClaimsRequest = ProviderEndpoint.PendingAuthenticationRequest.GetExtension(Of ClaimsRequest)()
        Dim sregResponse As ClaimsResponse = Nothing
        'If sregRequest IsNot Nothing Then
        '    sregResponse = profileFields.GetOpenIdProfileFields(sregRequest)
        'End If
        ProviderEndpoint.PendingAuthenticationRequest.IsAuthenticated = True
        If sregResponse IsNot Nothing Then
            ProviderEndpoint.PendingAuthenticationRequest.AddResponseExtension(sregResponse)
        End If

        ProviderEndpoint.PendingAuthenticationRequest.Response.Send()
        ProviderEndpoint.PendingAuthenticationRequest = Nothing

    End Sub

    Protected Sub No_Click(ByVal sender As Object, ByVal e As System.EventArgs) Handles no_button.Click
        ProviderEndpoint.PendingAuthenticationRequest.IsAuthenticated = False
        ProviderEndpoint.PendingAuthenticationRequest.Response.Send()
        ProviderEndpoint.PendingAuthenticationRequest = Nothing
    End Sub
End Class
