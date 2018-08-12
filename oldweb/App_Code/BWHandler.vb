Imports Microsoft.VisualBasic
Namespace Blogwind.Blog
    Public MustInherit Class BWHandler
        Implements IHttpHandler, IRequiresSessionState
        Private Shared ReadOnly Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)

        Private _Flash As Blogwind.Flash
        Private Session As HttpSessionState
        Private MyUIHelper As UIHelper
        Public ReadOnly Property Flash() As Blogwind.Flash
            Get
                If _Flash Is Nothing AndAlso Session IsNot Nothing Then
                    _Flash = New Flash(DirectCast(Session(Blogwind.Flash.FlashKey), Blogwind.Flash))
                End If
                Return _Flash
            End Get
        End Property

        Private Sub init(ByVal context As HttpContext)
            Session = context.Session
            MyUIHelper = New UIHelper(context)
        End Sub

        Public Sub CleanUp()
            If _Flash Is Nothing Then
                Session.Remove(Flash.FlashKey)
                Exit Sub
            End If

            Me.Flash.Sweep()
            If Me.Flash.HasItemsToKeep Then
                Session(Flash.FlashKey) = Me.Flash
            Else
                Session.Remove(Flash.FlashKey)
            End If
        End Sub

        Public ReadOnly Property IsReusable() As Boolean Implements System.Web.IHttpHandler.IsReusable
            Get

            End Get
        End Property

        Private Sub ProcessRequest(ByVal context As System.Web.HttpContext) Implements System.Web.IHttpHandler.ProcessRequest
            Me.init(context)
            Try
                ProcessBWRequest(context)
            Catch ex As System.Threading.ThreadAbortException
                'ProcessMoRequest used a context.Response.Redirect
            Catch ex As Exception
                Log.Error("MoHandler Error: " & ex.ToString)
                Throw
            Finally
                AfterProcessRequest(context)
            End Try
        End Sub

        Private Sub AfterProcessRequest(ByVal context As System.Web.HttpContext)
            ' Remove items from flash before leaving the page
            Me.Flash.Sweep()

            If Me.Flash.HasItemsToKeep Then
                context.Session(Flash.FlashKey) = Me.Flash
            Else
                context.Session.Remove(Flash.FlashKey)
            End If
        End Sub

        Public MustOverride Sub ProcessBWRequest(ByVal context As System.Web.HttpContext)

    End Class
End Namespace
