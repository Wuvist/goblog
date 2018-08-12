<%@ Page Language="VB" AutoEventWireup="false" CodeFile="FriendMood.aspx.vb" Inherits="blogwind.rpc_FriendMood" %>
<asp:repeater id="friendmood" runat="server" enableViewState="False">
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
</asp:repeater>