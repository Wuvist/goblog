Imports SqlConnHelper
Imports System.Xml
Namespace blogwind
  Partial Class rsssave
    Inherits System.Web.UI.Page

#Region " Web Form Designer Generated Code "

    'This call is required by the Web Form Designer.
    <System.Diagnostics.DebuggerStepThrough()> Private Sub InitializeComponent()

    End Sub


    Private Sub Page_Init(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Init
      'CODEGEN: This method call is required by the Web Form Designer
      'Do not modify it using the code editor.
      InitializeComponent()
    End Sub

#End Region

    Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
      'Put user code to initialize the page here
      Dim ph As New pageHelper(Me.Session, Me.Request)
      Dim blogger As String
      blogger = Session("blogger")
      If blogger = "" Then
        Response.End()
      End If

      Dim ds As New XmlDocument
      ds.Load(Request.Form("rss"))
      Dim List As XmlNodeList

      Dim dr As XmlNode
      Dim root As XmlNode = ds.DocumentElement
      Dim atom As Boolean

      If Request.Form("atom") = "true" Then
        atom = True
        List = ds.GetElementsByTagName("entry")
      Else
        atom = False
        List = ds.GetElementsByTagName("item")
      End If



      Dim i, cate As Integer
      i = 0

      cate = CInt(Request.Form("cate"))

      Dim pa As ParaHelper

      'cn.cmd.Parameters.Add("@cate", SqlDbType.Int)
      'cn.cmd.Parameters.Add("@tags", SqlDbType.NVarChar, 50)
      'cn.cmd.Parameters.Item(1).Value = blogger
      For Each dr In List
        pa = New ParaHelper

        If Request.Form("vv" & i.ToString) = "i" Then
          pa.addVarchar("@title", dr.Item("title").InnerText, 50)
          pa.addVarchar("@blogger", blogger, 50)
          If atom = False Then
            pa.addNText("@content", dr.Item("description").InnerText)
          Else
            pa.addNText("@content", dr.Item("content").InnerText)
          End If
          pa.addInt("@cate", cate)
          pa.addNVarchar("@tags", "", 50)
          connHelper.exeNonQuery("insert_doc", CommandType.StoredProcedure, pa.parameters)
        End If
        i = i + 1
      Next
    End Sub
  End Class
End Namespace