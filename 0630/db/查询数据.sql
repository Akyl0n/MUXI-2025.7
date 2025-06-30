-- 按要求进进行查询
-- 1). 查询喜欢阅读#3的人
SELECT p.name
FROM Person p
JOIN PersonLikes pl ON p.id = pl.person_id
WHERE pl.book_id = 'go-web';

-- 2). 查询没有被喜欢阅读的书的信息(id,作者,标题,库存)
SELECT b.id, b.author, b.title, s.stock
FROM Book b
JOIN Storage s ON b.id = s.bookid
WHERE b.id NOT IN (
    SELECT book_id FROM PersonLikes
);

-- 3). 查询哪些人喜欢哪本书,列出人名和书名
SELECT p.name, b.title
FROM Person p
JOIN PersonLikes pl ON p.id = pl.person_id
JOIN Book b ON pl.book_id = b.id
ORDER BY p.name;


