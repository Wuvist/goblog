<%
    Session.Abandon()
Response.Cookies("blogger").Value=""
response.redirect("/")
%>