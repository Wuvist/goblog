<%@ Page language="vb" debug="true"%>
<%@Import Namespace ="System.Data" %>
<%@Import Namespace ="System.Data.SQLClient" %>
<%
dim id() As String ={"Wuvist"} '用户名
dim pass() As String={"xanadu"}'密码


Dim configurationAppSettings As System.Configuration.AppSettingsReader = New System.Configuration.AppSettingsReader
Dim conn As SqlConnection
Dim sql As String
Dim i As Integer
dim cmd As SqlCommand
Dim md5 As New System.Security.Cryptography.MD5CryptoServiceProvider
conn = New SqlConnection(CType(configurationAppSettings.GetValue("ConnectionString", GetType(System.String)), String))
conn.open()
Dim data() As Byte
sql="update users set pwd=@pass where user_name=@id"
for i=0 to 0
	data = System.Text.Encoding.ASCII.GetBytes(pass(i))
	cmd = new SqlCommand(sql, conn)
	data = System.Text.Encoding.ASCII.GetBytes(pass(i).ToCharArray)
	cmd.Parameters.Add("@id", SqlDbType.NVarChar, 20)
	cmd.Parameters.Item(0).Value = id(i)
	cmd.Parameters.Add("@pass", SqlDbType.binary, 16)
	cmd.Parameters.Item(1).Value = md5.ComputeHash(data)
	cmd.ExecuteNonQuery()
Next
%>
congrates