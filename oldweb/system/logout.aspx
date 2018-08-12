<%
Session("blogger")=""
Response.Cookies("blogger").Value=""
response.redirect("/")
%>