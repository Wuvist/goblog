<%@ Page Language="VB" AutoEventWireup="false" CodeFile="pics.aspx.vb" Inherits="blogwind.pics" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<html>
	<head>
		<title><%=ph.g("图片管理")%></title>
		<link rel="icon" href="/favicon.ico" type="image/x-icon" />
		<link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />
		<link rel="stylesheet" href="/css/base.css" type="text/css" />
		<link rel="stylesheet" href="write.css" type="text/css" />
		<style type="text/css">
		    .dynamic:after{content:".";display:block;height:0;clear:both;visibility:hidden;}
		    #imgs{width:99%}
		    #imgs div.img{width:120px;height:110px;margin:10px;float:left;position:relative}
		    .paging{margin:5px; text-align:right;margin-right:10px}
		    .paging a, .paging a:visited, .paging a:hover{color:Blue}		    
            .paging a, .paging span{padding:3px;border:solid 1px grey}
            .paging span{color:Red}
            .x{position:absolute;top:0px;right:5px;width:13px;height:14px;background:url(/images/x_to_hide.gif);cursor:pointer}
		</style>
	</head>
	<body>
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
	<h3><font face="Verdana"><%=ph.g("图片管理")%></font></h3>
	
	    <div id="imgs" class="dynamic">
		<%For Each dr As DataRow In dt.rows%>
		    <div class="img" id="img<%=dr.Item("id") %>">
				<img alt="<%=dr.Item("des") %>" src="<%=blogwind.Photos.GetPath(dr, blogwind.Photos.thumbsize.i100_100)%>" />
				<div class="x" id="<%=dr.Item("id") %>"></div>
			</div>
		<%Next%>
        </div>
        <%=Me.GetPaging(pi.Start, pi.Count, Me.TotalCount)%>
        <script language="javascript" type="text/javascript" src="/js/jquery.lastest.js"></script>
        <script type="text/javascript">
        $(function(){
            $("#imgs div.img").hover(function(){
                $(this).find(".x").show();
            }, function(){
                $(this).find(".x").hide();
            });
            
            $(".x").click(function(){
                id = this.id;
                $.post("/rpc/photos.ashx",{action:"delete", id:id},
                    function(xml) {
                        if(xml=="0")
                            $("#img" + id).remove();
                });
            });
        })
        
        </script>
	</body>
</html>
