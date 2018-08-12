document.ns = navigator.appName == "Netscape"
if (navigator.appName == "Netscape")
{document.write("<layer id=dangdang top=20 width=100 height=500><img src=images/right.jpg border='0'></layer>");
myload()}
else
{document.write("<div id=dangdang style='position: absolute;width:100;top:20;left:800;visibility: visible;z-index: 1'><img src=images/right.jpg border='0'></div>");
myload()
}


