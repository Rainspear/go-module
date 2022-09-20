FIRST START A FOLDER TO BE O MODULE
go init module exaple.com/hello

SECOND GET DEPENDENCY - RIGHT NOW go get xxx WILL NOT DOWN THE EXECUTABLE FILE, IT WILL GET DEPENDENCY FOR go.mod
go get xyz.com/x/quote
go get rsc.io/quote

AFTER THAT go.mod WILL AUTOMATICALLY ADD DEPENDENCY AND INDIRECT DEPENDENCIES OF IT

TO UPGRADE A EXTERNAL PACKAGE
go get go get rsc.io/quote
go get go get rsc.io/quote@1.5.4

SOME ADDITIONAL COMMANDS
go list -m all // to list all dependecies in module (root folder of project)
go list -m -versions rsc.io/quote // to list all version of a dependency

REFERENCES 
https://go-dev.translate.goog/blog/using-go-modules?_x_tr_sl=en&_x_tr_tl=vi&_x_tr_hl=vi&_x_tr_pto=op,sc