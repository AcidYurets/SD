drop user test;
create user test password '123';


-- ================== При регистрации нового пользователя ==================
create role "70073bfb-a1a9-49dd-b305-6978bcafc56c" IN ROLE simple_user; -- Создаем отдельную роль для него
grant "70073bfb-a1a9-49dd-b305-6978bcafc56c" to test; -- Разрешаем основному пользователю эту роль
-- ================== При регистрации нового пользователя ==================





drop role "70073bfb-a1a9-49dd-b305-6978bcafc56c"



