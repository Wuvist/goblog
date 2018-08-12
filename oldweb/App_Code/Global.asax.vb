Imports System.Web
Imports System.Web.SessionState
Imports SqlConnHelper
Namespace Blogwind
    Public Class [Global]
        Inherits System.Web.HttpApplication

#Region " Component Designer Generated Code "

        Public Sub New()
            MyBase.New()

            'This call is required by the Component Designer.
            InitializeComponent()

            'Add any initialization after the InitializeComponent() call

        End Sub

        'Required by the Component Designer
        Private components As System.ComponentModel.IContainer

        'NOTE: The following procedure is required by the Component Designer
        'It can be modified using the Component Designer.
        'Do not modify it using the code editor.
        <System.Diagnostics.DebuggerStepThrough()> Private Sub InitializeComponent()
            components = New System.ComponentModel.Container
        End Sub

#End Region

        Sub Application_Start(ByVal sender As Object, ByVal e As EventArgs)
            ' Fires when the application is started
            log4net.Config.XmlConfigurator.Configure()
            Dim Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)

            Dim configurationAppSettings As System.Configuration.AppSettingsReader = New System.Configuration.AppSettingsReader
            connHelper.init(CType(configurationAppSettings.GetValue("ConnectionString", GetType(System.String)), String))
            Dim p As String
            p = Web.VirtualPathUtility.AppendTrailingSlash(Context.Request.ApplicationPath)
            LanguageHelper.initcate(Server.MapPath(p & "langs"))
            GeoIP.init(Server.MapPath(p & "App_Data/GeoIP.dat"))
			System.Net.ServicePointManager.Expect100Continue = false
            Log.Info("Started")
        End Sub

        Sub Session_Start(ByVal sender As Object, ByVal e As EventArgs)
            ' Fires when the session is started
        End Sub

        Sub Application_BeginRequest(ByVal sender As Object, ByVal e As EventArgs)
            ' Fires at the beginning of each request
        End Sub

        Sub Application_AuthenticateRequest(ByVal sender As Object, ByVal e As EventArgs)
            ' Fires upon attempting to authenticate the use
        End Sub

        Sub Application_Error(ByVal sender As Object, ByVal e As EventArgs)
            ' Code that runs when an unhandled error occurs
            Dim strPageUrl As String = Request.Path
            Dim ErrorInfo As Exception = Server.GetLastError()

            Dim str As String
            str = "Url:" + strPageUrl & vbNewLine
            str = str & "Date:" + Now.ToString("yyyy-MM-dd HH:mm") & vbNewLine
            str = str + "Header: "
            str = str + Request.Headers.ToString & vbNewLine & vbNewLine
            str = str + "Get: "
            str = str + Request.QueryString.ToString & vbNewLine & vbNewLine
            str = str + "Post: "
            str = str + Request.Form.ToString & vbNewLine & vbNewLine
            str = str + "Error: "
            str = str + ErrorInfo.ToString() & vbNewLine

            If ErrorInfo.Message.StartsWith("The file '") And ErrorInfo.Message.EndsWith(" does not exist.") Then
                Response.StatusCode = 404
                Server.Transfer("/404.html", False)
            End If

            If Not ini.IsProduction Then
                Dim Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)
                Log.Error(str)
            Else
                Dim msg As System.Net.Mail.MailMessage = New System.Net.Mail.MailMessage()
                msg.To.Add(New System.Net.Mail.MailAddress("wuvist@gmail.com"))
                msg.From = New System.Net.Mail.MailAddress("info@blogwind.com", "Blogwind")
                msg.Subject = "Blogwind Web Error"
                msg.IsBodyHtml = False
                msg.Body = str
                SMTP.send(msg)
            End If

            If Response.StatusCode = 404 Then
                Response.Clear()
                Server.Transfer("/404.html", False)
            Else
                If ini.IsProduction Then
                    Response.Clear()
                    Response.StatusCode = 500
                    Server.Transfer("/err.html", False)
                End If
            End If
        End Sub

        Sub Session_End(ByVal sender As Object, ByVal e As EventArgs)
            ' Fires when the session ends
        End Sub

        Sub Application_End(ByVal sender As Object, ByVal e As EventArgs)
            ' Fires when the application ends
        End Sub

    End Class
End Namespace