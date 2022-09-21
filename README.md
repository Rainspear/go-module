FIRST START A FOLDER TO BE O MODULE

**go init module exaple.com/hello**


SECOND GET DEPENDENCY - RIGHT NOW go get xxx WILL NOT DOWN THE EXECUTABLE FILE, IT WILL GET DEPENDENCY FOR go.mod

**go get xyz.com/x/quote** 

**go get rsc.io/quote**


AFTER THAT go.mod WILL AUTOMATICALLY ADD DEPENDENCY AND INDIRECT DEPENDENCIES OF IT


TO UPGRADE A EXTERNAL PACKAGE

**go get go get rsc.io/quote**

**go get go get rsc.io/quote@1.5.4**


SOME ADDITIONAL COMMANDS

**go list -m all** // to list all dependecies in module (root folder of project)

**go list -m -versions rsc.io/quote** // to list all version of a dependency


REFERENCES 

https://go.dev/blog/migrating-to-go-modules

https://stackoverflow.com/questions/61845013/package-xxx-is-not-in-goroot-when-building-a-go-project

https://stackoverflow.com/questions/35480623/how-to-import-local-packages-in-go