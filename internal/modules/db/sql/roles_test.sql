-- +++++++++++++++++++ Проверка корректности переключения роли +++++++++++++++++++
SET ROLE superuser;

INSERT INTO users
VALUES ( 'ccd1937e-d48e-496b-8e13-80a30a7c2300'
       , now()
       , now()
       , null
       , '+79197624300'
       , 'test_user'
       , '$2a$14$FFn5L2FT7gc8ZukWKEULMOerJr12NzVNs3V31OCaqvoIuZGMQgqvK'
       , 'simple_user');

CALL before_each_query('ccd1937e-d48e-496b-8e13-80a30a7c2300');

SELECT grantee, privilege_type, table_name
FROM information_schema.role_table_grants
WHERE table_name NOT LIKE 'pg_%' AND grantee NOT IN ('postgres', 'PUBLIC');

SELECT current_user;
SELECT user_uuid
FROM session;


-- Проверка прав доступа к пользователям
SELECT *
FROM users;

UPDATE users
SET login = 'test_user_updated'
WHERE uuid = 'c635e505-f957-4e33-b065-694e9cdf4e5c';

SELECT *
FROM users
WHERE uuid = 'c635e505-f957-4e33-b065-694e9cdf4e5c';

UPDATE users
SET login = 'test_user_updated'
WHERE uuid = 'ccd1937e-d48e-496b-8e13-80a30a7c2300';


-- Проверка прав доступа к тегам
SELECT *
FROM tags;

UPDATE tags
SET name = 'Для себя любимого'
WHERE uuid = '1d719d39-fba4-464f-b1cc-37cde5d44b41';

INSERT INTO tags
VALUES (default, now(), now(), null, 'Тест', 'Тест');































