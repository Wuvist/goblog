Imports SqlConnHelper
Imports System.Web
Imports GNU.Gettext

Public Class pageHelper
    Private cu As System.Globalization.CultureInfo
    Public blogger As String
    Public bloggerID As Integer
    Public viewer As String = ""
    Public nick As String = ""
    Public url As String = ""
    Public lang As langs.lang = langs.lang.zh_cn
    Public login As Boolean = False
    Public relogin As Boolean = False


    Public ReadOnly Property nickname() As String
        Get
            If Not String.IsNullOrEmpty(viewer) Then
                Return viewer
            ElseIf Not String.IsNullOrEmpty(nick) Then
                Return nick
            ElseIf Not String.IsNullOrEmpty(blogger) Then
                Return blogger
            Else
                Return ""
            End If
        End Get
    End Property

    Public Function g(ByVal word As String) As String
        If cu Is Nothing Then
            cu = New System.Globalization.CultureInfo(langs.getLangString(Me.lang))
        End If
        Return LanguageHelper.catalog.GetString(word, cu)
    End Function

    Public Sub New(ByVal context As HttpContext)
        Dim hash As String
        Dim result As Boolean = False

        blogger = context.Request("blogger")
        hash = context.Request("hash")

        Dim pa As New ParaHelper
        pa.addNVarchar("@id", blogger, 50)
        Dim b As Byte()
        Dim dr As DataRow
        dr = connHelper.exeDataRow("select nick, id, pwd,TS,lang,DOB from bloggers where user_name=@id", CommandType.Text, pa.parameters)

        If Not dr Is Nothing Then
            b = dr("pwd")
            If hash = util.ToHexString(util.MD5(util.ToHexString(b) & blogger)) Then
                context.Session("blogger") = blogger
                nick = dr("nick")
                context.Session("nick") = nick
                bloggerID = CInt(dr("id"))
                context.Session("bloggerID") = bloggerID
                lang = langs.lang.Parse(GetType(langs.lang), CStr(dr("lang")))
                context.Session("lang") = lang
                login = True
                relogin = True
            End If
        End If
    End Sub

    Public Sub New(ByRef session As SessionState.HttpSessionState, ByRef request As HttpRequest)
        Dim ck As HttpCookie
        ck = request.Cookies("viewer")

        If Not ck Is Nothing Then
            viewer = System.Text.Encoding.UTF8.GetString(Convert.FromBase64String(ck.Values("nick")))
            url = ck.Values("hp")
        End If

        If session("blogger") = "" Then
            Try
                ck = request.Cookies("blogger")
            Catch ex As Exception
                ck = Nothing
            End Try

            If Not ck Is Nothing Then
                Dim pwd As String

                Dim result As Boolean
                result = False

                blogger = ck.Values("id")
                pwd = ck.Values("pass")

                If (Not String.IsNullOrEmpty(blogger)) And (Not String.IsNullOrEmpty(pwd)) Then
                    Dim pa As New ParaHelper
                    pa.addNVarchar("@id", blogger, 50)
                    Dim b As Byte()
                    Dim dr As DataRow
                    dr = connHelper.exeDataRow("select nick, id, pwd,TS,lang,DOB from bloggers where user_name=@id", CommandType.Text, pa.parameters)

                    If Not dr Is Nothing Then
                        b = dr("pwd")
                        If pwd = System.Text.Encoding.ASCII.GetString(b).GetHashCode.ToString Then
                            session("blogger") = blogger
                            nick = dr("nick")
                            session("nick") = nick
                            bloggerID = CInt(dr("id"))
                            session("bloggerID") = bloggerID
                            lang = langs.lang.Parse(GetType(langs.lang), CStr(dr("lang")))
                            session("lang") = lang
                            login = True
                            relogin = True
                        End If
                    End If
                End If
            End If
        Else
            blogger = session("blogger")
            login = True
            If session("bloggerID") Is Nothing Then
                Dim dr As DataRow
                Dim pa As New ParaHelper
                pa.addNVarchar("@id", blogger, 50)
                dr = connHelper.exeDataRow("select nick, id, pwd,TS,lang,DOB from bloggers where user_name=@id", CommandType.Text, pa.parameters)
                If dr Is Nothing Then
                    login = False
                    Exit Sub
                Else
                    nick = dr("nick")
                    session("nick") = nick
                    bloggerID = CInt(dr("id"))
                    session("bloggerID") = bloggerID
                    lang = langs.lang.Parse(GetType(langs.lang), CStr(dr("lang")))
                    session("lang") = lang
                End If
            End If
            bloggerID = session("bloggerID")
            lang = session("lang")
            nick = session("nick")
        End If

        If session("blogger") = "" Then
            lang = langs.getLang("", request.UserLanguages)
        End If
    End Sub
End Class
