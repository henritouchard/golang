Hello and welcome to my golang github repository

This was my first experience with golang.
if you work on windowds and discovered "everything" software
(like : http://download.cnet.com/Everything/3000-2379_4-10890746.html)
and want the same thing for your linux, you can have a try.
still have not enough time to finish it. probably i will one day. but i don't
really like UI ...
this is a fast file finder.
the principe is easy. there's an sql database which store every files of your linux pc
(yes i know with golang i can easily do it for windows and mac too, it will come "soon" :)
because i would try webserver i chosed to do the "interface"(yes i know it looks like nothing)
on a web page served on 3040. you type a word in the search bar and it will fast display results
with file path anywhere on your computer.
of course there are things to do:

* upgrade UI.
* add better UX with instant research without page reload.
* implement regexp (easy we could use those of mysql directly)
* maybe do it in a window, it would be interesting.
* optimise files-indexation to store only where there's a new file (it doesn't take a lot of CPU, it's less than 1%).
* add a oneclick go to file near the results.

feel free to do it by yourself!

Have fun (golang seem's awsome).
BTW i remember there is one error : the file indexor need to have root rights to open any folders
i didn't handle it but it should be really easy!

p.s don't forget not to execute random code you don't know! it can be dangerous!
