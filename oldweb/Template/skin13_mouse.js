/*
Magic Wand cursor II (By Kurt at kurt.grigg@virgin.net)
Modified and permission granted to Dynamic Drive to feature script in archive
For full source, usage terms, and 100's more DHTML scripts, visit http://dynamicdrive.com
*/

if (document.all){
with (document){
write('<div id="starsDiv" style="position:absolute;top:0px;left:0px">')
write('<div style="position:relative;width:1px;height:1px;background:#60bb60;font-size:1px;visibility:visible"></div>')
write('<div style="position:relative;width:1px;height:1px;background:#60bb60;font-size:1px;visibility:visible"></div>')
write('<div style="position:relative;width:1px;height:1px;background:#60bb60;font-size:1px;visibility:visible"></div>')
write('<div style="position:relative;width:1px;height:1px;background:#60bb60;font-size:2px;visibility:visible"></div>')
write('<div style="position:relative;width:1px;height:1px;background:#60bb60;font-size:2px;visibility:visible"></div>')
write('<div style="position:relative;width:1px;height:1px;background:#60bb60;font-size:2px;visibility:visible"></div>')
write('<div style="position:relative;width:1px;height:1px;background:#60bb60;font-size:3px;visibility:visible"></div>')
write('<div style="position:relative;width:2px;height:2px;background:#60bb60;font-size:3px;visibility:visible"></div>')
write('<div style="position:relative;width:2px;height:2px;background:#60bb60;font-size:3px;visibility:visible"></div>')
write('<div style="position:relative;width:2px;height:2px;background:#60bb60;font-size:4px;visibility:visible"></div>')
write('<div style="position:relative;width:2px;height:2px;background:#60bb60;font-size:4px;visibility:visible"></div>')
write('<div style="position:relative;width:2px;height:2px;background:#60bb60;font-size:5px;visibility:visible"></div>')
write('<div style="position:relative;width:2px;height:2px;background:#60bb60;font-size:5px;visibility:visible"></div>')
write('<div style="position:relative;width:3px;height:3px;background:#60bb60;font-size:6px;visibility:visible"></div>')
write('</div>')
}
}

if (document.layers)
{window.captureEvents(Event.MOUSEMOVE);}
var yBase = 200;
var xBase = 200;
var step = 1;
var currStep = 0;
var Xpos = 1;
var Ypos = 1;

if (document.all)
{
  function MoveHandler(){
  Xpos = document.body.scrollLeft+event.x;
  Ypos = document.body.scrollTop+event.y;
  }
  document.onmousemove = MoveHandler;
}

else if (document.layers)
{
  function xMoveHandler(evnt){
  Xpos = evnt.pageX;
  Ypos = evnt.pageY;
  }
  window.onMouseMove = xMoveHandler;
}

function animateLogo() {
if (document.all)
{
 yBase = window.document.body.offsetHeight/6;
 xBase = window.document.body.offsetWidth/6;
}
else if (document.layers)
{
 yBase = window.innerHeight/8;
 xBase = window.innerWidth/8;
}

if (document.all)
{
 for ( i = 0 ; i < starsDiv.all.length ; i++ )
 {
  starsDiv.all[i].style.top = Ypos + yBase*Math.sin((currStep + i*4)/12)*Math.cos(400+currStep/200);
 starsDiv.all[i].style.left = Xpos + xBase*Math.sin((currStep + i*3)/10)*Math.sin(currStep/200);
 }
}

else if (document.layers)
{
 for ( j = 0 ; j < 14 ; j++ ) //number of NS layers!
 {
  var templayer="a"+j
  document.layers[templayer].top = Ypos + yBase*Math.sin((currStep + j*4)/12)*Math.cos(400+currStep/200);
  document.layers[templayer].left = Xpos + xBase*Math.sin((currStep + j*3)/10)*Math.sin(currStep/200);
 }
}
currStep+= step;
setTimeout("animateLogo()", 10);
}
animateLogo();