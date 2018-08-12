var hintText;
hintText=new String("读取中。。。<img src='/images/ajaxing.gif' height='16' width='16'>　");
function fcpageup()
{
	fcpage--;
	if(fcpage<2)
	{
		fcpage=1;
		document.getElementById("fcpagingup").style.display = "none";
	}
	document.getElementById("fcpaging").innerHTML = hintText + document.getElementById("fcpaging").innerHTML;
	ajaxRead("myfriends.aspx?fcom=" + fcpage + "","fc","fcpaging")
}
function fcpagedown()
{
	fcpage++;
	document.getElementById("fcpagingup").style.display = "";
	document.getElementById("fcpaging").innerHTML = hintText + document.getElementById("fcpaging").innerHTML;
	ajaxRead("myfriends.aspx?fcom=" + fcpage + "","fc","fcpaging")
}
function fmpageup()
{
	fmpage--;
	if(fmpage<2)
	{
		document.getElementById("fmpagingup").style.display = "none";
	}
	document.getElementById("fmpaging").innerHTML = hintText + document.getElementById("fmpaging").innerHTML;
	ajaxRead("rpc/FriendMood.aspx?fbm=" + fmpage + "","fm","fmpaging")
}
function fmpagedown()
{
	fmpage++;
	document.getElementById("fmpagingup").style.display = "";
	document.getElementById("fmpaging").innerHTML = hintText + document.getElementById("fmpaging").innerHTML;
	ajaxRead("rpc/FriendMood.aspx?fbm=" + fmpage + "","fm","fmpaging")
}
function fbmpageup()
{
	fbmpage--;
	if(fbmpage<2)
	{
		bmcpage=1;
		document.getElementById("fbmpagingup").style.display = "none";
	}
	document.getElementById("fbmpaging").innerHTML = hintText + document.getElementById("fbmpaging").innerHTML;
	ajaxRead("myfriends.aspx?fbm=" + fbmpage + "","fbm","fbmpaging")
}
function fbmpagedown()
{
	fbmpage++;
	document.getElementById("fbmpagingup").style.display = "";
	document.getElementById("fbmpaging").innerHTML = hintText + document.getElementById("fbmpaging").innerHTML;
	ajaxRead("myfriends.aspx?fbm=" + fbmpage + "","fbm","fbmpaging")
}
function fblogpageup()
{
	fblogpage--;
	if(fblogpage<2)
	{
		fblogpage=1;
		document.getElementById("fblogpagingup").style.display = "none";
	}
	document.getElementById("fblogpaging").innerHTML = hintText + document.getElementById("fblogpaging").innerHTML;
	ajaxRead("myfriends.aspx?fblog=" + fblogpage + "","fblog","fblogpaging")
}
function fblogpagedown()
{
	fblogpage++;
	document.getElementById("fblogpagingup").style.display = "";
	document.getElementById("fblogpaging").innerHTML = hintText + document.getElementById("fblogpaging").innerHTML;
	ajaxRead("myfriends.aspx?fblog=" + fblogpage + "","fblog","fblogpaging")
}

function cpagedown()
{
	cpage++;
	document.getElementById("cpaging").innerHTML = hintText + document.getElementById("cpaging").innerHTML;
	ajaxRead("ajaxcomment.aspx?pnum=" + cpage + "&psize=10","comments","cpaging");
	document.getElementById("cpagingup").style.display = "";
}

function cpageup()
{
	cpage--;
	if(cpage<2)
	{
		cpage=1;
		document.getElementById("cpagingup").style.display = "none";
	}
	document.getElementById("cpaging").innerHTML = hintText + document.getElementById("cpaging").innerHTML;
	ajaxRead("ajaxcomment.aspx?pnum=" + cpage + "&psize=10","comments","cpaging")
}

function blogpageup()
{
	blogpage--;
	if(blogpage<2)
	{
		blogpage=1;
		document.getElementById("blogpagingup").style.display = "none";
	}
	document.getElementById("blogpaging").innerHTML = hintText + document.getElementById("blogpaging").innerHTML;
	ajaxRead("ajaxblog.aspx?pnum=" + blogpage + "&psize=10","blogs","blogpaging")
}

function blogpagedown()
{
	blogpage++;
	document.getElementById("blogpaging").innerHTML = hintText + document.getElementById("blogpaging").innerHTML;
	ajaxRead("ajaxblog.aspx?pnum=" + blogpage + "&psize=10","blogs","blogpaging");
	document.getElementById("blogpagingup").style.display = "";
}

function ajaxRead(file,divid,hintid){
 var xmlObj = null;
 if(window.XMLHttpRequest){
 xmlObj = new XMLHttpRequest();
 } else if(window.ActiveXObject){
 xmlObj = new ActiveXObject("Microsoft.XMLHTTP");
 } else {
 return;
 }
 xmlObj.onreadystatechange = function(){
 if(xmlObj.readyState == 4){
 updateObj(divid, xmlObj.responseText,hintid);
 }
 }
 xmlObj.open ('GET', file, true);
 xmlObj.send ('');
 }

function updateObj(obj, data,hintid){
	document.getElementById(obj).innerHTML = data;
	tmp=document.getElementById(hintid).innerHTML;
	document.getElementById(hintid).innerHTML =  tmp.substring(tmp.indexOf("　")+1);
}