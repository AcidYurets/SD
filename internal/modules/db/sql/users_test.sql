select current_user;


-- ================== Перед каждым запросом ==================
SET role "70073bfb-a1a9-49dd-b305-6978bcafc56c";
-- ================== Перед каждым запросом ==================



-- Проверка прав доступа к тегам
SELECT *
FROM tags;

UPDATE tags
SET name = 'Для себя'
WHERE name = 'Для себя';

INSERT INTO tags
VALUES (default, now(), now(), null, 'Тест', 'Тест');

-- Проверка прав доступа к пользователям
SELECT *
FROM users;

UPDATE users
SET login = 'Тест2'
WHERE login = 'yurets';
