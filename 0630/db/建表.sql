-- type Book struct{
-- ID string 
-- Title string 
-- Author string 
-- }
-- type Storage struct{
-- ID int
-- Stock int
-- }
-- type Person struct{
-- ID int 
-- Name string
-- }
-- 

-- DROP TABLE Book CASCADE;
-- DROP table Storage CASCADE;
-- DROP table Person CASCADE;
-- DROP table PersonLikes CASCADE;


CREATE TABLE Book (
  id VARCHAR(20) PRIMARY KEY,
  title VARCHAR(50),
  author VARCHAR(50)
);

CREATE TABLE Storage (
  id INT PRIMARY KEY,
  bookid VARCHAR(20),
  stock INT,
  FOREIGN KEY (bookid) REFERENCES Book(id)
);

CREATE TABLE Person (
  id INT PRIMARY KEY,
  name VARCHAR(10)
);

CREATE TABLE PersonLikes (
  person_id INT,
  book_id VARCHAR(20),
  PRIMARY KEY (person_id, book_id),
  FOREIGN KEY (person_id) REFERENCES Person(id),
  FOREIGN KEY (book_id) REFERENCES Book(id)
);



