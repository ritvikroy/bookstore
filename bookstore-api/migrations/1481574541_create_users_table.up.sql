DROP TABLE Users;
CREATE TABLE Users (
 id  SERIAL PRIMARY KEY,
  userName VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  isAdmin Boolean NOT NULL);

Insert into Users (userName,password,isAdmin) values ('user','user',true);
Insert into Users (userName,password,isAdmin) values ('user1','user1',true);
Insert into Users (userName,password,isAdmin) values ('user2','user2',false);
Insert into Users (userName,password,isAdmin) values ('user3','user3',false);



  

