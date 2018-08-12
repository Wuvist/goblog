<%@ Page Language="VB" MasterPageFile="~/master/base.master" AutoEventWireup="false" CodeFile="login.aspx.vb" Inherits="login" title="Untitled Page" %>
<asp:Content ID="Content1" ContentPlaceHolderID="css" Runat="Server">
</asp:Content>
<asp:Content ID="Content2" ContentPlaceHolderID="content" Runat="Server">
	<% If Not String.IsNullOrEmpty(Request("username")) Then%>
	    <div class="error-msg-box">用户名或密码错误</div>
	<%End If %> 
	<form action="/system/check.aspx" method="post">
        <div class="input-form">
            <div>
                  <label>用户名：</label>
		          <input id="id" class="loginput" name="id" type="text" value="<%=request("username") %>">
	        </div>
	        <div>
                  <label>密　码：</label>
		          <input id="pass" class="loginput" name="pass" type="password">
		          <input checked="checked" name="record" type="checkbox">&nbsp;记录我的登陆信息
	        </div>	 
	        <div>
		        <label>&nbsp;</label>
	        <input id="Submit" class="logbutton" value="登录" name="Submit" type="submit">&nbsp;<a href="/apply.aspx">申请</a>&nbsp;
        		
		        <a href="/getpass.aspx">忘记密码？</a></div>
        </div>
    </form>
</asp:Content>
<asp:Content ID="Content3" ContentPlaceHolderID="js" Runat="Server">
</asp:Content>