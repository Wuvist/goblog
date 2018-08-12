moods=new Array();
moods[0]='饿';moods[1]='离';moods[2]='哭';moods[3]='食';moods[4]='爽';moods[5]='读';moods[6]='拼';moods[7]='泣';moods[8]='闷';moods[9]='羞';moods[10]='暑';moods[11]='烦';moods[12]='哀';moods[13]='拽';moods[14]='囧';moods[15]='穷';moods[16]='睡';moods[17]='空';moods[18]='乐';moods[19]='困';moods[20]='闲';moods[21]='恼';moods[22]='懒';moods[23]='愤';moods[24]='冷';moods[25]='湿';moods[26]='贺';moods[27]='思';moods[28]='怒';moods[29]='聊';moods[30]='厌';moods[31]='忙';moods[32]='吵';moods[33]='泪';moods[34]='热';moods[35]='病';moods[36]='累';moods[37]='乱';moods[38]='吓';moods[39]='赞';moods[40]='美';moods[41]='喜';moods[42]='怨';moods[43]='憋'; moods[44]='渴';moods[45]='伤';moods[46]='痛';moods[47]='幽';moods[48]='静'; moods[49]='恨';moods[50]='凶';moods[51]='爱';
function EditMood(b)
{
	if(b.value == "修改")
	{
		b.value = '取消';
		$("#MyMood").fadeOut("Slow",function (){$('#NewMood').fadeIn()});
	}
	else
	{
		b.value = '修改';
		$("#NewMood").fadeOut("Slow",function (){$('#MyMood').fadeIn()});
	}
}
function SaveMood(b)
{	
	var m, msg;
	b.disabled= "true";
	b.value = "保存中..";
	m = $("#mood").val();
	msg = $("#msg").val();
	$.post("rpc/moods.aspx",{
	mood:m,
	msg:msg
	  },function(xml){
		$('#MyMood').html('【' + moods[m] + "】" + msg);
		$("#BtnEditMood").val('修改');
		b.disabled= "";
		b.value = "保存";
		$("#NewMood").fadeOut("Slow",function (){$('#MyMood').fadeIn()});
		$.get("rpc/MoodList.aspx",{	size:10,page:1,t:(new Date()).getTime()},function(xml){
			$('#MoodMsg').html(xml);		});
		ajaxRead("rpc/FriendMood.aspx?fbm=1","fm","fmpaging");

	});
}

function delete_friend(id)
{
	$.post("rpc/Friends.ashx",{
		action:"delete",
		friendid: id
	  },function(xml){
	    $('#friend_' +id).highlightFade({color:'red',speed:'slow',complete:function() { $('#friend_' +id).remove(); }});
	});
}

function show_friend(id,nick, icon)
{
var str;
str ='<table width="115" style="display:inline;" id="friend_' + id + '" onmouseover="toggle_friend_X(\'' + id + '\',0, event)" onmouseout="toggle_friend_X(\'' + id + '\',1, event)"><tr><td width="70" height="80"><img src="' + icon + '" width="64"></td>';
str += '<td width="30" valign="bottom"><div id="friend_X' + id + '" style="display:none"><a style="cursor:pointer" onclick=\'delete_friend("' + id + '")\' title="点了就删，没得确认哦~">删除</a></div></td></tr><tr><td colspan="2" valign="top" id="friend_N_' + id + '"><a href="#" oncontextmenu="shows(\'' + id + '\');" onclick="opens(\''+ id +'\');">' + nick + '</a></td></table>';
return str;
}

function GetFriends(b)
{
	toggle("MyFriends");
	if(document.getElementById("MyFriends").style.display=='none')
		b.value="我的好友";
	else
		b.value="关闭列表";

	if (document.getElementById("MyFriends").style.display=='block' && $("#MyFriends").html()==MyFriendsIntiHTML)
	{
		$.get("rpc/Friends.ashx",{
		  },function(xml){
			  fs = eval(xml);
			  var html = ShowFrineHint + "<br>";
			  for (var f in fs)
			  {
				html += show_friend(fs[f].id, fs[f].nick, fs[f].icon);
			  }
			  $("#MyFriends").html(html);
		});
	}
}
var FriendMoods=new Array();

function toggle_friend_X(id, action, ev)
{
	var cor;
	ev = ev || window.event;
	cor = mouseCoords(ev);
	
	var paramswidth =250;
	title ='朋友心情';
	
	if (action==0)
	{
		document.getElementById('friend_X' +id).style.display="inline";
		$('#friend_N_' +id).css({ background : "#e5e5e5" });

		var de = document.documentElement;
		var w = self.innerWidth || (de&&de.clientWidth) || document.body.clientWidth;
		var clickElementy = cor.y; //set y position
		$("body").append("<div id='JT' style='width:"+paramswidth*1+"px'><div id='JT_arrow_left'></div><div id='JT_close_left'>"+title+"</div><div id='JT_copy'><div class='JT_loader'><div></div></div>");//right side
		

		var clickElementx = cor.x + 50; //set x position
		
		$('#JT').css({left: clickElementx+"px", top: clickElementy+"px"});
		$('#JT').show();
		if(FriendMoods[id])
		{
			 $('#JT_copy').html(FriendMoods[id]);
		}
		else
		{
			$.get("rpc/moods.aspx",{
				bloggerid: id
			  },function(xml){
			  if(xml=='')
			  {
				  xml = '还没有设置过心情~';
			  }
			  else
			  {
				  var m;
				  m = eval('(' + xml + ")");
				  xml = '【' + moods[m.Mood] + "】" + m.MoodMsg;
			  }
			   FriendMoods[id] = xml;

			   $('#JT_copy').html(xml);
			});
		}
	}
	else
	{
		document.getElementById('friend_X' +id).style.display="none";
		$('#friend_N_' +id).css({ background: "white" });
		$('#JT').remove();
	}
}

function mouseCoords(ev){
	if(ev.pageX || ev.pageY){
	return {x:ev.pageX, y:ev.pageY};
	}
	return {
	x:ev.clientX + document.body.scrollLeft - document.body.clientLeft,
	y:ev.clientY + document.body.scrollTop - document.body.clientTop
	};
} 
