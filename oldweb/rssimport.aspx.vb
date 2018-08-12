Imports SqlConnHelper
Imports System.Text
Imports System.Xml
Namespace blogwind

  Partial Class rssimport
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
      Dim ph As New pageHelper(Me.Session, Me.Request)
    End Sub

    Private Sub Button1_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button1.Click
      Dim ds As New XmlDocument
      ds.Load(r.Text)
      Dim List As XmlNodeList

      Dim dr As XmlNode
      Dim atom As Boolean
      atom = False
      Dim root As XmlNode = ds.DocumentElement

      List = ds.GetElementsByTagName("item")
      If List.Count = 0 Then
        List = ds.GetElementsByTagName("entry")
        If List.Count > 0 Then
          atom = True

        End If
      End If

      Dim sb As New StringBuilder
      Dim i As Integer
      i = 0
      sb.Append("请选择需要导入的文章，然后按确认：<br><br><INPUT type=""hidden"" value=""" + r.Text + """ name=""rss"">")
      If atom = True Then
        sb.Append("<INPUT type=""hidden"" value=""true"" name=""atom"">")
      End If

      For Each dr In List
        sb.Append("<INPUT type=""checkbox"" Checked name=""vv" & i & """  value=""i"" >" & dr.Item("title").InnerText)
        sb.Append("<br>")
        i = i + 1
      Next
      sb.Append("")
      view.Text = sb.ToString

      hint.Text = "请务必选择需要导入的网志分类："
      s.Text = "<INPUT class=""logsub"" type=""submit"" value=""确认"">"

      Dim pa As New ParaHelper

      pa.addVarchar("@blogger", Session("blogger"), 50)
      cate.DataSource = connHelper.retrieve("select_cate_by_Blogger", CommandType.StoredProcedure, pa.parameters)
      cate.DataBind()
    End Sub
  End Class
End Namespace