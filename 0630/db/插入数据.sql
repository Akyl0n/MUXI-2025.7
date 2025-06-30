INSERT INTO Book (id, title, author) VALUES
('go-away', 'the way to go', 'Ivo'),
('go-lang', 'Go语言圣经', 'Alan, Brian'),
('go-web', 'Go Web编程', 'Anonymous'),
('con-cur', 'Concurrency in Go', 'Katherine');

INSERT INTO Storage (id, bookid, stock) VALUES
(1, 'go-away', 20),
(2, 'go-lang', 17),
(3, 'go-web', 4),
(4, 'con-cur', 9);

INSERT INTO Person (id, name) VALUES
(1, '小明'),
(2, '张三'),
(3, '翟曙');

-- 小明喜欢：#1 go-away, #3 go-web, #4 con-cur
INSERT INTO PersonLikes (person_id, book_id) VALUES
(1, 'go-away'),
(1, 'go-web'),
(1, 'con-cur');

-- 张三喜欢：#1 go-away
INSERT INTO PersonLikes (person_id, book_id) VALUES
(2, 'go-away');

-- 翟曙喜欢：#3 go-web, #4 con-cur
INSERT INTO PersonLikes (person_id, book_id) VALUES
(3, 'go-web'),
(3, 'con-cur');
