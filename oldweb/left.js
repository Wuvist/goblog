var imgheight
document.ns = navigator.appName == "Netscape"
window.screen.width>800 ? imgheight=550:imgheight=550
 function myload()
{
if (navigator.appName == "Netscape")
{document.dang.pageY=pageYOffset+window.innerHeight-imgheight;
document.dang.pageX=+window.innerWidth-105
document.dangdang.pageY=pageYOffset+window.innerHeight-imgheight;
document.dangdang.pageX=+window.innerWidth-105
mymove();
}
else
{
dang.style.top=document.body.scrollTop+document.body.offsetHeight-imgheight;
dang.style.left=10
dangdang.style.top=document.body.scrollTop+document.body.offsetHeight-imgheight;
dangdang.style.left=document.body.offsetWidth-130
mymove();
}
}
 function mymove()
 {
 if(document.ns)
 {
 document.dang.top=pageYOffset+window.innerHeight-imgheight
 dang.style.left=10
 document.dangdang.top=pageYOffset+window.innerHeight-imgheight
 document.dangdang.left=pageXOffset+window.innerWidth-130
 setTimeout("mymove();",5)
 }
 else
 {
 dang.style.top=document.body.scrollTop+document.body.offsetHeight-imgheight;
 dang.style.left=10
 dangdang.style.top=document.body.scrollTop+document.body.offsetHeight-imgheight;
 dangdang.style.left=document.body.scrollLeft+document.body.offsetWidth-130
 setTimeout("mymove();",5)
 }
 }

function MM_reloadPage(init) {  //reloads the window if Nav4 resized
  if (init==true) with (navigator) {if ((appName=="Netscape")&&(parseInt(appVersion)==4)) {
    document.MM_pgW=innerWidth; document.MM_pgH=innerHeight; onresize=MM_reloadPage; }}
  else if (innerWidth!=document.MM_pgW || innerHeight!=document.MM_pgH) location.reload();
}
MM_reloadPage(true)




if (navigator.appName == "Netscape")
{document.write("<layer id=dang top=20 width=100 height=500><img src=images/left.jpg border='0'></layer>");
myload()}
else
{document.write("<div id=dang style='position: absolute;width:100;top:20;left:10;visibility: visible;z-index: 1'><img src=images/left.jpg border='0'></div>");
myload()
}