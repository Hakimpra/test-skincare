# API-login-go-mysql
login API with bcrypt

# How to Run ?
#1.Create The Database with name go_api

#2. Create Table 
create table with name users

	CREATE TABLE `users` (
	  `id` int(11) NOT NULL AUTO_INCREMENT,
	  `username` varchar(50) DEFAULT NULL,
	  `first_name` varchar(200) NOT NULL,
	  `last_name` varchar(200) NOT NULL,
	  `password` varchar(120) DEFAULT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=latin1;

#run
you should import the library first in terminal

	go get database/sql

	go get golang.org/x/crypto/bcrypt

	go get github.com/go-sql-driver/mysql
	go get github.com/kataras/go-sessions

### And here we go 
	go run main.go