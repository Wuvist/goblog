<%@ Import Namespace ="System" %>
<%@ Import Namespace ="System.Data" %>
<%@ Import Namespace ="System.Data.SQLClient" %>
<% dim ph as New pageHelper(Me.Session, Me.Request)%>
<script language="vb" runat="server">



sub page_load(s as object, e as EventArgs)
	
	if Not isPostBack Then
		BindLinkData()
		NewURL.text="http://"
	end if
end sub

sub BindLinkData()
			Dim dsLink As DataSet
			Dim dbComm As SqlDataAdapter
			Dim conn As SqlConnection
			Dim sql As String

			Dim configurationAppSettings As System.Configuration.AppSettingsReader = New System.Configuration.AppSettingsReader
			conn = New SqlConnection(CType(configurationAppSettings.GetValue("ConnectionString", GetType(System.String)), String))
             
			
			dsLink= New DataSet
			dsLink.Tables.Add("links")
		    
			
			sql = "select_links_by_blogger"
			dbComm = New SqlDataAdapter(sql, conn)
			dbComm.SelectCommand.CommandType = CommandType.StoredProcedure
			dbComm.SelectCommand.Parameters.Add("@blogger", SqlDbType.VarChar, 50)
			dbComm.SelectCommand.Parameters.Item(0).Value = Session("blogger")
			dbComm.Fill(dsLink, "links")
				
			dgLink.DataSource=dsLink.Tables("links")
			dgLink.DataBind()
end sub

Sub dgLink_ItemCreated(Sender As Object, e As DataGridItemEventArgs)
	  
	  Select Case e.Item.ItemType
            Case ListItemType.Item, ListItemType.AlternatingItem, ListItemType.EditItem
		Dim myTableCell As TableCell
		myTableCell = e.Item.Cells(3)
	        Dim myDeleteButton As LinkButton
	        myDeleteButton = myTableCell.Controls(0)
 	        myDeleteButton.Attributes.Add("onclick","return confirm('你真的要删除此链接?');")
        End Select
end sub

sub EditLink(Obj as Object, e As DataGridCommandEventArgs)
	dgLink.EditItemIndex=e.Item.ItemIndex
	BindLinkData()
end sub

sub UpdateLink(Obj as Object, e As DataGridCommandEventArgs)
	
	Dim conn As SqlConnection
	Dim sql As String
	
	Dim configurationAppSettings As System.Configuration.AppSettingsReader = New System.Configuration.AppSettingsReader
	conn = New SqlConnection(CType(configurationAppSettings.GetValue("ConnectionString", GetType(System.String)), String))
	
	sql="update_link"
	dim Cmd as New SQLCommand(sql,conn)
	cmd.CommandType = CommandType.StoredProcedure
	
	Dim ctlLink As TextBox= e.Item.Cells(0).Controls(0)
	Dim ctlURL As TextBox= e.Item.Cells(1).Controls(0)
	
	cmd.Parameters.Add(New SQLParameter("@name", ctlLink.text))
	cmd.Parameters.Add(New SQLParameter("@url", ctlURL.text))
	cmd.Parameters.Add(New SQLParameter("@id", dgLink.DataKeys.Item(e.Item.ItemIndex)))

	conn.Open()
	cmd.ExecuteNonQuery
	dgLink.EditItemIndex=-1
	BindLinkData()

end sub

sub DeleteLink(Obj as Object, e As DataGridCommandEventArgs)
	
	Dim conn As SqlConnection
	Dim configurationAppSettings As System.Configuration.AppSettingsReader = New System.Configuration.AppSettingsReader
	conn = New SqlConnection(CType(configurationAppSettings.GetValue("ConnectionString", GetType(System.String)), String))

	Dim DeleteCmd As String = "DELETE from links Where [id] = @id" 
	Dim Cmd as New SqlCommand(DeleteCmd, Conn)

	Cmd.Parameters.Add(New SqlParameter("@id", dgLink.DataKeys.Item(e.Item.ItemIndex)))
	
	Conn.Open()
	Cmd.ExecuteNonQuery()
	Conn.Close()

	BindLinkData()
end sub

sub AddNewLink_OnClick(obj as Object, e As EventArgs)
	
	dim dsLink as DataSet
	dim MyCommand as SqlCommand
	
	Dim conn As SqlConnection
	Dim configurationAppSettings As System.Configuration.AppSettingsReader = New System.Configuration.AppSettingsReader
	conn = New SqlConnection(CType(configurationAppSettings.GetValue("ConnectionString", GetType(System.String)), String))
	
	Dim InsertCmd as string ="insert_link"
	MyCommand= new SqlCommand(InsertCmd, conn)
	MyCommand.CommandType = CommandType.StoredProcedure
	
	MyCommand.Parameters.Add(New SqlParameter("@link",SqlDbType.varchar,50))
	MyCommand.Parameters.Add(New SqlParameter("@url",SqlDbType.varchar,100))
	MyCommand.Parameters.Add(New SqlParameter("@blogger",SqlDbType.varchar,50))

	MyCommand.Parameters("@link").Value= NewLink.text
	MyCommand.Parameters("@url").Value= NewURL.text
	MyCommand.Parameters("@blogger").Value= Session("blogger")
	
	conn.Open()
	MyCommand.ExecuteNonQuery
	conn.Close()
	
	NewLink.text=""
	NewURL.text="http://"
	BindLinkData()        
end sub



</script>

<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=gb2312">
<link rel="icon" href="/blogwind.ico" type="image/x-icon" />
<link rel="shortcut icon" href="blogwind.ico" type="image/x-icon" />
<meta content="博客，博客风，留学" name="keywords">
<meta content="留学生博客网站" name="description">
<link rel="stylesheet" href="admin.css" type="text/css"> 
</head>
<BODY BGCOLOR="FFFFFF" TEXT="White" LINK="Silver" VLINK="Silver" Marginwidth=0 marginheight=0 leftmargin=0 topmargin=0>
			<table width="100%" border="0" cellPadding="0" cellSpacing="0">
				<tr><td width="200px"><img src='/images/new/banner.jpg' width="200" height="50" border="0"></td>
				<td style="background: url(/images/new/banner_dot.jpg)" valign="middle" class="menu"><a href="/write.aspx" target="_self"><%=ph.g("编写网志")%></a> 
					<a href="/system/admin_userinfo.aspx" target="_self"><%=ph.g("个人资料")%></a>
					<a href="/system/admin_cate.aspx" target="_self"><%=ph.g("分类管理")%></a>
					<a href="/system/admin_article.aspx" target="_self"><%=ph.g("网志管理")%></a>
					<a href="/pics.aspx" target="_self"><%=ph.g("图片管理")%></a>
					<a href='/Template/skin_switch.aspx?blogger=<%=Session("blogger")%>' target="_self"><%=ph.g("选择模板")%></a>
					<a href='/system/admin_link.aspx' target="_self"><%=ph.g("链接管理")%></a>
					<a href="/" target="_top"><%=ph.g("返回首页")%></a>
					<a href='/<%=Session("blogger")%>' target="_top"><%=ph.g("返回网志")%></a>
					<a href="/logout.aspx" target="_top"><%=ph.g("退出登陆")%></a></td></tr>
				</table>
	<h3><font face="Verdana">友情链接管理页面 </font></h3>

<form runat="server">

	<asp:DataGrid id="dgLink"
			runat="server"
			Font-Name="Verdana"
			Font-Size="10"
			Font-Color="White"
			AlternatingfItemStyle-BackColor="lightGray" 
			OnEditCommand="EditLink"
			OnUpdateCommand="UpdateLink"
			AutoGenerateColumns="false"
			DataKeyField="id"
			OnDeleteCommand="DeleteLink"
			OnItemCreated="dgLink_ItemCreated"
			>
		<Columns>
		<asp:BoundColumn
				HeaderText="链接名称"
				DataField="link" 
				ItemStyle-width="200px"/>

		<asp:BoundColumn
				HeaderText="网址"
				DataField="URL" 
				ItemStyle-width="200px"/>

		<asp:EditCommandColumn
				EditText="编辑"
				UpdateText="更新"
				HeaderText="编辑"/>

		<asp:ButtonColumn
				HeaderText="删除"
				Text="删除"
				CommandName="Delete"
                />
		</Columns>
        </asp:DataGrid>
		<br/>
	    <table><tr><td>
		添加新链接：
		</td></tr>
		<tr><td>
		友情链接名称：</td><td><asp:TextBox id="NewLink" runat="server"/>
		</td>
		<tr><td>
		链接网址: </td><td><asp:TextBox id="NewURL" runat="server"/>
		</td></tr>
        <tr><td>
		<asp:Button id="AddNewLink" runat="server"
			Text="添加"
			OnClick="AddNewLink_OnClick"/>
		</td></tr></table>
        
		
</form>
</body>
</html>