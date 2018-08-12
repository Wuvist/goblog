<%@ Register TagPrefix="Acme" TagName="Message" Src="log.ascx" %>
<%@ Register TagPrefix="Acme" TagName="userpic" Src="icon.ascx" %>
<%@ Page Language="vb" AutoEventWireup="false" CodeFile="default.aspx.vb" Inherits="blogwind._default"%>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<HTML>
	<HEAD>
		<title>博客风 Blogwind</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<link rel="icon" href="/favicon.ico" type="image/x-icon" />
		<link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />
		<meta content="博客,博客风,留学,blog, weblog, 留学生, 网志, 部落格" name="keywords">
		<meta content="留学生博客网站, blog service" name="description">
		<link rel="stylesheet" href="home.css" type="text/css">
		<link rel="commands" href="/js/ub.js" name="博客风心情发布的Ubiquity command" />
		<link rel="stylesheet" href="<%=ini.GetFileWithVersion("css/base.css") %>" type="text/css">
		<script type="text/javascript" src="/baseutil.js"></script>
		<script type="text/javascript" src="/homeajax.js"></script>
		<script type="text/javascript" src="js/util.js"></script>
		<script type="text/javascript" src="/icon.js"></script>
		<script type="text/javascript" src="js/jquery.js"></script>
		<script type="text/javascript" src="js/jquery.highlightFade.js"></script>
		<link rel="stylesheet" href="images/media/jtip/global.css" type="text/css" />
	    <script type="text/javascript"><!--
var blogpage=1;
var cpage=1;
var fblogpage=1;
var fcpage=1;
var fbmpage=1;
var fmpage=1;
var word_cancel="<%=ph.g("取消")%>";
var word_save="<%=ph.g("保存")%>";
var word_url="<%=ph.g("网址")%>";
var word_public="<%=ph.g("公开")%>";
var word_title="<%=ph.g("标题")%>";
var ShowFrineHint ="<%=ph.g("在朋友昵称上点左键可打开朋友网志，点右键则可查看其朋友圈")%>";

 //--></script>
		<script type="text/javascript" src="/menu.js"></script>
	</HEAD>
	<body>
		<div id="header"><table><tr><td></td></tr></table></div>
		<div id="contain">
			<div class="left">
				<asp:Panel id="userview" runat="server" EnableViewState="False" Visible="False">
					<IMG height="27" alt="" src="images<%=ph.g("/")%>title_favourite.gif" width="611"><DIV class="tb"><DIV class="fulltext">
					<table border=0>
						<tr>
							<td valign="top" width="150px">
								<div style='width:150px;padding-top:5px'><Acme:userpic runat="server" ID="userpic1"/>
								<br><INPUT class="logbutton" type="button" value="<%=ph.g("我的好友")%>" onclick="GetFriends(this);"></div>
							</td>
							<td valign="top">
										<table>
											<tr style="line-height:150%">
												<td valign="top">我的心情：<br><input type="button" value="修改" onclick="EditMood(this)" id="BtnEditMood" class="logsub"></td>
												<td valign="top"><div id='MyMood'>读取中...</div>
													<div style="display:none" id="NewMood"><select id="mood" name="mood">
<script>
moodOrder =[51,12,43,35,32,5,0,11,23,26,49,48,14,44,17,2,19,22,18,33,36,24,1,29,37,31,40,8,21,28,6,7,15,34,45,25,3,10,4,16,27,46,41,38,20,50,9,30,47,42,39,13];
for (i=0;i<moods.length ;i++ )
{
	j = moodOrder[i];
	if (j==<%=MyMood%>)
		document.write("<option value='" + j + "' selected>" + moods[j] + "</option>");
	else
		document.write("<option value='" + j + "'>" + moods[j] + "</option>");
}
</script>
</select>
<input type="text" id="msg" name="msg" style="width:310px" maxlength=500><br><input type="button" class="logsub" onclick="SaveMood(this)" value="保存"></div>
<script>
$.get("rpc/moods.aspx",{
  },function(xml){
  if(xml=='')
  {
	  $('#MyMood').html('嗯？貌似您还没有设置过心情，点击下面紫色的“修改”按钮便可以设置了~');
  }
  else
  {
	  var m;
	  m = eval('(' + xml + ")");
	  $('#MyMood').html('【' + moods[m.Mood] + "】" + m.MoodMsg);
  }
});
</script>
												</td>
											</tr>
										</table>
										<DIV id="grey">
										<%=ph.g("我的链接")%>：<a href='<%=Session("blogger")%>/' target="_blank"><font color=red><%=ph.g("我的网志")%></font></a>　　<asp:Repeater id="links" runat="server" EnableViewState="False">
											<ItemTemplate><a href='<%# DataBinder.Eval(Container.DataItem, "URL") %>' target="_blank"><%# DataBinder.Eval(Container.DataItem, "link") %></a>　　</ItemTemplate>
										</asp:Repeater><br>
										<INPUT class="logsub" type="button" value="<%=ph.g("添加链接")%>" onclick="addlink(this);">
										<span id=newlink></span>
										</div>
							</td>
						</tr>
					</table>
							<div id="white">								 
								<div id='MyFriends' style="display:none;width:600px;height:250px;overflow-y:auto;"><img src="images/ajaxing.gif"></div>
								<div id='viewfriend'></div>
							</div>

						</DIV>
					</DIV><div class="footright"><IMG height="5" alt="" src="images/middle_shadow_foot.gif" width="307"></div><!--bloggerlink结束-->

		<br />
		<IMG height="27" alt="" src="images/title_fmood.png" width="611">
		<div id="fm"><asp:repeater id="friendmood" runat="server" enableViewState="False">
					<ItemTemplate>
                        <table width="400px" align="right">
                            <tr>
                                <td width="300px" align=right>                                    
                                        <div class="rbubble_tr  rbubble_<%#blogger.GetBloggerBubbleStyle(Container.DataItem("blogger_id"))%>">
                                            <p class="rbubble_tl rbubble_<%#blogger.GetBloggerBubbleStyle(Container.DataItem("blogger_id"))%>"><%#util.AddHyperLink(Server.HtmlEncode(Container.DataItem("msg")))%></p>
                                            <p class="rbubble_bl rbubble_<%#blogger.GetBloggerBubbleStyle(Container.DataItem("blogger_id"))%>">
                                                <em class="rbubble_br rbubble_<%#blogger.GetBloggerBubbleStyle(Container.DataItem("blogger_id"))%>">&nbsp;</em>
                                            </p>
                                        </div>                                        
                                </td>
                                <td>
                                    <%#Container.DataItem("nick") & "【" & blogwind.BloggerMood.GetMoodStr(Container.DataItem("mood_id")) & "】"%>
                                </td>
                            </tr>
                        </table><br style="clear:both"/>
					</ItemTemplate>
					<AlternatingItemTemplate>
						<table width="380px">
                            <tr>
                                <td>
                                    <%#Container.DataItem("nick") & "【" & blogwind.BloggerMood.GetMoodStr(Container.DataItem("mood_id")) & "】"%>
                                </td>                            
                                <td width="300px">
                                        <div class="lbubble_tr  lbubble_<%#blogger.GetBloggerBubbleStyle(Container.DataItem("blogger_id"))%>">
                                            <p class="lbubble_tl  lbubble_<%#blogger.GetBloggerBubbleStyle(Container.DataItem("blogger_id"))%>"><%#util.AddHyperLink(Server.HtmlEncode(Container.DataItem("msg")))%></p>
                                            <p class="lbubble_bl  lbubble_<%#blogger.GetBloggerBubbleStyle(Container.DataItem("blogger_id"))%>">
                                                <em class="lbubble_br  lbubble_<%#blogger.GetBloggerBubbleStyle(Container.DataItem("blogger_id"))%>">&nbsp;</em>
                                            </p>
                                        </div>   
                                </td>
                            </tr>
                        </table>
					</AlternatingItemTemplate>
				</asp:repeater></div>
				<div id="fmpaging" style="LINE-HEIGHT: 25px; text-align:right">
					<a id="fmpagingup" href="javascript:fmpageup()" style="display:none"><%=ph.g("上一页")%></a>
					<a id="fmpagingdown" href="javascript:fmpagedown()"><%=ph.g("下一页")%></a>&nbsp;&nbsp;&nbsp;&nbsp;</div>
		<div class="footright"><IMG height="5" alt="" src="images/middle_shadow_foot.gif" width="307"></div>
		<br />
		<!--friendmood结束-->	

<div class="leftsub">

<IMG height="27" alt="" src="images<%=ph.g("/")%>title_fblog.gif" width="303">
			<div class="tbh">
				<div class="texthalf"><div id="fblog"><asp:repeater id="friendblog" runat="server" enableViewState="False">
						<ItemTemplate>
							<div id="white"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "index") %>.shtml' target='_blank'><%#Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "nick") & " " & _
							                                                                                                                                                                 ph.g("写了")) & _
							                                                                                                                                                                 "【" & DataBinder.Eval(Container.DataItem, "title") & "】"%></a></div>
						</ItemTemplate>
						<AlternatingItemTemplate>
							<div id="grey"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "index") %>.shtml' target='_blank'><%#Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "nick") & " " & _
							                                                                                                                                                                ph.g("写了")) & _
							                                                                                                                                                                "【" & DataBinder.Eval(Container.DataItem, "title") & "】"%></a></div>
						</AlternatingItemTemplate>
					</asp:repeater></div>
					<div id="fblogpaging" style="LINE-HEIGHT: 25px; text-align:right">
<a id="fblogpagingup" href="javascript:fblogpageup()" style="display:none"><%=ph.g("上一页")%></a>
<a id="fblogpagingdown" href="javascript:fblogpagedown()"><%=ph.g("下一页")%></a>&nbsp;&nbsp;&nbsp;&nbsp;</div>
				</div>
			</div>
			<div class="footright"><IMG height="5" alt="" src="images/middle_shadow_foot.gif" width="290"></div> <!--friendblog结束--></div>
		<div class="rightsub"><IMG height="27" alt="" src="images<%=ph.g("/")%>title_fcomment.gif" width="303">
			<div class="tb">
				<div class="texthalf"><div id="fc"><asp:repeater id="friendcomment" runat="server" enableViewState="False">
						<ItemTemplate>
	                        <div id="white"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "article_id") %>.shtml' target='_blank'><%#app.GetEventMsg(DataBinder.Eval(Container.DataItem, "payload"), DataBinder.Eval(Container.DataItem, "title"), DataBinder.Eval(Container.DataItem, "article_id"), DataBinder.Eval(Container.DataItem, "Comment"))%></a>
	                        </div>
                        </ItemTemplate>
                        <AlternatingItemTemplate>
	                        <div id="grey"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "article_id") %>.shtml' target='_blank'><%#app.GetEventMsg(DataBinder.Eval(Container.DataItem, "payload"), DataBinder.Eval(Container.DataItem, "title"),DataBinder.Eval(Container.DataItem, "article_id"),DataBinder.Eval(Container.DataItem, "Comment"))%></a>
	                        </div>
                        </AlternatingItemTemplate>
					</asp:repeater></div>
					<div id="fcpaging" style="LINE-HEIGHT: 25px; text-align:right">
<a id="fcpagingup" href="javascript:fcpageup()" style="display:none"><%=ph.g("上一页")%></a>
<a id="fcpagingdown" href="javascript:fcpagedown()"><%=ph.g("下一页")%></a>&nbsp;&nbsp;&nbsp;&nbsp;</div>
				</div>
			</div>
			<div class="footright"><IMG height="5" alt="" src="images/middle_shadow_foot.gif" width="290"></div>
		</div> <!--friendcomment结束-->
</asp:Panel>				
<% If Not Request.Url.AbsoluteUri.IndexOf("http://127.0.0.1/") > -1 Then%>
<img src="images/title_moodlist.png" alt="" width="611" height="27">
	<div id='MoodMsg'><asp:repeater id="MoodMsgList" runat="server" enableViewState="False">
<ItemTemplate>
<div id="white"><a href="/<%# Container.DataItem("user_name") %>/"><%# Container.DataItem("nick") %></a><%#util.AddHyperLink(Server.HtmlEncode("【" & blogwind.BloggerMood.GetMoodStr(Container.DataItem("mood_id")) & "】" & Container.DataItem("msg")))%></div></ItemTemplate>
<AlternatingItemTemplate>
<div id="grey"><a href="/<%# Container.DataItem("user_name") %>/"><%# Container.DataItem("nick") %></a><%#util.AddHyperLink(Server.HtmlEncode("【" & blogwind.BloggerMood.GetMoodStr(Container.DataItem("mood_id")) & "】" & Container.DataItem("msg")))%></div>
</AlternatingItemTemplate>
</asp:repeater></div>
<script>
var MyFriendsIntiHTML=$("#MyFriends").html();
</script>
<div style="LINE-HEIGHT: 25px; text-align:right"><a href="moods.aspx">更多..</a></div><br>
<% End If %>
<form id="Form1" method="post" runat="server" style="display: inline;">
				<img src="images<%=ph.g("/")%>title_newblog.gif" alt="" width="611" height="27"><div class="tb"><div class="text"><div id="blogs">
				<asp:Repeater id="new_article" runat="server" enableViewState="False">	
					<ItemTemplate>
						<div id="white">
							<%# DataBinder.Eval(Container.DataItem, "nick") %> <%=ph.g("写了")%> <a href='<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "[index]") %>.shtml' target="_blank">【<%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "title")) %>】</a><font color="DodgerBlue"><%=ph.g("评论")%>：<%# DataBinder.Eval(Container.DataItem, "Comment") %></font>
							</div>
						</ItemTemplate>
					<AlternatingItemTemplate>
						<div id="grey">
							<%# DataBinder.Eval(Container.DataItem, "nick") %> <%=ph.g("写了")%> <a href='<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "[index]") %>.shtml' target="_blank">【<%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "title")) %>】</a><font color="DodgerBlue"><%=ph.g("评论")%>：<%# DataBinder.Eval(Container.DataItem, "Comment") %></font>
							</div>
					</AlternatingItemTemplate>
				</asp:Repeater></div><div id="blogpaging" style="text-align:right;LINE-HEIGHT: 25px; ">
<a id="blogpagingup" href="javascript:blogpageup()" style="display:none"><%=ph.g("上一页")%></a>
<a id="blogpagingdown" href="javascript:blogpagedown()"><%=ph.g("下一页")%></a>&nbsp;&nbsp;&nbsp;&nbsp;</div>
					</div>
				</div><div class="footright"><IMG height="5" alt="" src="images/middle_shadow_foot.gif" width="307" ></div><br>
				
				<img src="images<%=ph.g("/")%>title_reply.gif" alt="" width="611" height="27"><div class="tb"><div class="text">
				<div id="comments">
				<asp:Repeater id="new_comment" runat="server" enableViewState="False">							
							<ItemTemplate>
								<div id="white">
								【<%# DataBinder.Eval(Container.DataItem, "blogname") %>】<a href='<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "Article") %>.shtml#blogc<%# DataBinder.Eval(Container.DataItem, "cid") %>' target="_blank"><%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "title")) %></a>-<%# DataBinder.Eval(Container.DataItem, "Author") %>
								</div>
							</ItemTemplate>
							<AlternatingItemTemplate>
							<div id="grey">
								【<%# DataBinder.Eval(Container.DataItem, "blogname") %>】<a href='<%# DataBinder.Eval(Container.DataItem, "id") %>/<%# DataBinder.Eval(Container.DataItem, "Article") %>.shtml#blogc<%# DataBinder.Eval(Container.DataItem, "cid") %>' target="_blank"><%# Server.HtmlEncode(DataBinder.Eval(Container.DataItem, "title")) %></a>-<%# DataBinder.Eval(Container.DataItem, "Author") %>
								</div>
							</AlternatingItemTemplate>
				</asp:Repeater></div><div id="cpaging" style="text-align:right;LINE-HEIGHT: 25px; ">
<a id="cpagingup" href="javascript:cpageup()" style="display:none"><%=ph.g("上一页")%></a>
<a id="cpagingdown" href="javascript:cpagedown()"><%=ph.g("下一页")%></a>&nbsp;&nbsp;&nbsp;&nbsp;</div>
					</div>
				</div><div class="footright"><IMG height="5" alt="" src="images/middle_shadow_foot.gif" width="307" ></div><br>
				<img src="images<%=ph.g("/")%>title_news.gif" alt="" width="611" height="27"><div class="tb"><div class="text"><asp:Repeater id="Anounce" runat="server" enableViewState="False">
							<ItemTemplate>
								<div id="white"><a href='/Wuvist/<%# DataBinder.Eval(Container.DataItem, "id") %>.shtml' target='_blank'><%# DataBinder.Eval(Container.DataItem, "Subject") %> @ <%# DataBinder.Eval(Container.DataItem, "addtime") %></a></div>
							</ItemTemplate>
							<AlternatingItemTemplate>
								<div id="grey"><a href='/Wuvist/<%# DataBinder.Eval(Container.DataItem, "id") %>.shtml' target='_blank'><%# DataBinder.Eval(Container.DataItem, "Subject") %> @ <%# DataBinder.Eval(Container.DataItem, "addtime") %></a></div>
							</AlternatingItemTemplate>
						</asp:Repeater></div>
				</div><div class="footright"><IMG height="5" alt="" src="images/middle_shadow_foot.gif" width="307" ></div><br><!--网站公告结束-->
			</FORM>
			</div>

			<div class="right">
				<img src="images<%=ph.g("/")%>title_sys.gif" width="161" height="23" alt=""><div class="tbs"><div class="subtext"><br><Acme:Message runat="server" ID="Message1" />
						<div id="stat">
							<%=ph.g("博客总数")%>：<%=count_blogger%><br>
							<%=ph.g("网志总数")%>：<%=count_doc%><br>
							博客风QQ群：1683930<br>
							<br></div>

				<h3><a href="moods.aspx"><img src="images/title_moods.png" alt="" width="161" height="23" border="0"></a></h3>
				<div style="line-height:25px">
				<asp:repeater id="MoodList" runat="server" enableViewState="False">
					<ItemTemplate>
						<span style="font-size:<%# GetMoodSize(Container.DataItem("blogger_count")) %>pt"><a href="moods.aspx?mood=<%# Server.UrlEncode(Blogwind.BloggerMood.GetMoodStr(Container.DataItem("id"))) %>"><%# Blogwind.BloggerMood.GetMoodStr(Container.DataItem("id")) %></a></span>
					</ItemTemplate>
				</asp:repeater></div>
					</div>
				</div><IMG height="2" alt="" src="images/middle_sub_shadow_foot.gif" width="161"><br><br>


				<img src="images<%=ph.g("/")%>title_sub_last.jpg" alt="" width="161" height="23"><div class="tbs"><div class="subtext">
						<div id="stat">
						<asp:Repeater id="Recommend" runat="server" EnableViewState="False">
								<ItemTemplate>
									<div id="white"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/' title='【人气：<%# DataBinder.Eval(Container.DataItem, "hit") %>】
<%# DataBinder.Eval(Container.DataItem, "intro") %>'><%# DataBinder.Eval(Container.DataItem, "nick") %>：<%# DataBinder.Eval(Container.DataItem, "blogname") %></a></div>
								</ItemTemplate>
								<AlternatingItemTemplate>
									<div id="grey"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/' title='【人气：<%# DataBinder.Eval(Container.DataItem, "hit") %>】
<%# DataBinder.Eval(Container.DataItem, "intro") %>'><%# DataBinder.Eval(Container.DataItem, "nick") %>：<%# DataBinder.Eval(Container.DataItem, "blogname") %></a></div>
								</AlternatingItemTemplate>
							</asp:Repeater>
						</div>
					</div>
				</div><IMG height="2" alt="" src="images/middle_sub_shadow_foot.gif" width="161"><br><br>

					<img src="images<%=ph.g("/")%>title_newblogger.jpg" alt="" width="161" height="23"><div class="tbs"><div class="subtext">
						<div id="stat">
										<asp:Repeater id="Newest" runat="server" EnableViewState="False">
								<ItemTemplate>
									<div id="white"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/' title='<%# DataBinder.Eval(Container.DataItem, "intro") %>'><%# DataBinder.Eval(Container.DataItem, "nick") %>：<%# DataBinder.Eval(Container.DataItem, "blogname") %></a></div>
								</ItemTemplate>
								<AlternatingItemTemplate>
									<div id="grey"><a href='/<%# DataBinder.Eval(Container.DataItem, "id") %>/' title='<%# DataBinder.Eval(Container.DataItem, "intro") %>'><%# DataBinder.Eval(Container.DataItem, "nick") %>：<%# DataBinder.Eval(Container.DataItem, "blogname") %></a></div>
								</AlternatingItemTemplate>
							</asp:Repeater>
							</div>
					</div>
				</div><IMG height="2" alt="" src="images/middle_sub_shadow_foot.gif" width="161"><br><br>
			</div>
		</div>
		<div id="footer"><p><font color="#FF0000"><b>友情链接</b></font>: 
          <a href="http://www.sgchinese.com/">狮城论坛</a>
		<a href="http://www.xinchaoliu.org/">新潮留</a>
		<a href="http://www.huasing.org/">华新</a>
		<a href="http://www.yanruyu.com/">颜如玉</a>
		<a href="http://www.singeat.com/">新加坡美食网</a>
		<a href="http://www.House-SG.com/">House-SG</a>
		<a href="http://www.jujuya.com/">聚聚呀</a>
		<a href="http://www.careeralbum.com/">CareerAlbum</a>
		</p>
			Copyright &copy;2003 - 2008 <a href="http://www.blogwind.com">BlogWind.com</a> <a href='http://www.miibeian.gov.cn/'>粤ICP备05018623号</a> <a href="http://wiki.blogwind.com/doku.php?id=aboutus">关于博客风</a></div>
<script>
if(document.location.href.indexOf("cn.blogwind.com")>0)
{
	$("a").each(function(){
		if (this.href.indexOf("Wuvist")>0)
		{
			this.href = "http://wiki.blogwind.com/doku.php?id=Wuvist";
		}
	});
}
var t= '<%=Request.Url.AbsoluteUri%>';
</script>
	</body>
</HTML>
