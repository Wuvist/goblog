<%@ Page Language="VB"  MasterPageFile="~/master/base.master" AutoEventWireup="false"  CodeFile="mills.aspx.vb" Inherits="mills" %>
<asp:Content ID="Content1" ContentPlaceHolderID="css" Runat="Server">
<style>
    #server_host{width:500px}
  </style>  
</asp:Content>
<asp:Content ID="Content2" ContentPlaceHolderID="content" Runat="Server">
<div>
<h3>帐户列表</h3>
<%  If accs.Count = 0 Then%>
尚无帐号
<%Else%>
    <table class="list-table" id="accounts" cellspacing="0px" cellpadding="5px">
        <tr>
            <th>帐号类型</th>
            <th>用户名</th>
            <th>状态</th>
			<th>删除</th>
        </tr>
    <%  For Each acc As Blogwind.WindMill.Account In accs%>  
    <tr id="account<%=acc.Id%>">
    <td><%=blogwind.WindMill.Account.GetAccountTypeName(acc.AccountType)%></td>
    <td><%=acc.UserName%></td>
    <td><%=blogwind.WindMill.Account.GetAccountStatus(acc.Status)%></td>
	<td><a onclick="delete_account(<%=acc.Id%>)"><img src="/images/delete.gif" class="x" /></a></td>
    </tr>
    <%  Next%>
    </table>    
<%End If%>
<h3>添加帐号</h3>
<form action="" method="post">
<div class="input-form">
    <div>
          <label>帐号类型：</label><select name="account_type" id="account_type">
            <option value="3">请选择帐号类型</option> 
            <option value="1">MSN Space</option>
			<option value="100">百度空间</option>
            <option value="2">Blogger</option>
            <option value="3">WordPress</option>
            <option value="4">MetaWebLog</option>
            <option value="5">Twitter</option>
            <option value="6">饭否</option></select>
    </div>
    <div>
          <label>服务器地址：</label><input type="text" id="server_host" name="server_host" value=""  />
    </div>    
    <div>
      <label>用户名：</label><input type="text" id="username" name="username" value="" />
    </div>      
    <div>
      <label>密码：</label><input type="password" id="password" name="password" value="" />
    </div>  
    <div><label>&nbsp;</label>
       <input type="button"  value="添加" id="submitbtn" />
    </div>
 </div>  
       </form>
        </div>    
</asp:Content>
<asp:Content ID="Content3" ContentPlaceHolderID="js" Runat="Server">
    <script type="text/javascript" src="<%=ini.GetFileWithVersion("js/mills.js") %>"></script>       
</asp:Content>

