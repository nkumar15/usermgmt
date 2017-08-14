# usermgmt
This library has been created solely for learning purpose.


Introduction
-------------
Recently I started learning Go, and I was curious if someone could take a quick look through my repo and suggest me Go idioms that I'm missing out on, parts of code that are not readable, imporoper approach if followed etc. 

I have try to follow below guidelines for learning and practicing

Learning objective and Requirements
------------------
* Mockable database
* Sharing of all global configuration parameters across functions
	 logger
	 db if needed
	 context for request scope
* Provide handler function to library client
* Dependency injection
* App Handler pattern
* Follow REST style
* Logging style and error handling
* Make idiomatic small libaries like user mgmt, login and similar repetetive tasks used in mostly every web project.


Assumptions and limitations
-------------------
* Targeting only sqlite library 
* Made Dependencies on
		upper.io for database access, didn't wanted to spend time on writing sql queries. I understand Go practices suggest sql/sqlx.
		logrus logger for logging needs, wanted to explore logrus.

		
Things finding difficult to do.
-------------------
* Passing logger down to data access layer code (store functions). There are ways like passing through context or explicit argument. 
   Not sure what is suitable approach.
