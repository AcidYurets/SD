-- =====================================================================================================
-- =============================================== Итого ===============================================

-- =================== При создании БД ===================
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create role admin;
create role simple_user;
create role premium_user;
-- Суперпользователь для выполнения служебных операций (например, для создания пользователей)
create role superuser WITH SUPERUSER;
-- =================== При создании БД (конец блока) ===================

-- =================== При запуске сервера с включенной опцией SQL миграции ===================
SET ROLE superuser;

DROP OWNED BY admin;
DROP OWNED BY simple_user;
DROP OWNED BY premium_user;
drop role admin;
drop role simple_user;
drop role premium_user;

create role admin;
create role simple_user;
create role premium_user;

-- ===== Админ =====
grant all on all tables in schema public to admin;

-- ===== Простой пользователь =====
grant select, update (login, phone, password_hash) on users to simple_user;
grant select on access_rights to simple_user;
grant all on events to simple_user;
grant all on invitations to simple_user;
grant all on events_tags to simple_user;
grant select on tags to simple_user;

-- ===== Премиум пользователь =====
grant select, update (login, phone, password_hash) on users to premium_user;
grant select on access_rights to premium_user;
grant all on events to premium_user;
grant all on invitations to premium_user;
grant all on events_tags to premium_user;
grant all on tags to premium_user;


-- ===== RLS для таблицы 'users' =====
-- Разрешаем RLS
ALTER TABLE users
    ENABLE ROW LEVEL SECURITY;

-- Представление session для получения uuid текущего пользователя
DROP VIEW IF EXISTS session;
CREATE VIEW session AS
SELECT current_setting('app.user_uuid')::uuid AS user_uuid;

-- У всех есть доступ к сессии
grant all on session to simple_user, premium_user, admin;

-- Получать любых пользователей могут все
DROP POLICY IF EXISTS users_select ON users;
CREATE POLICY users_select
    ON users
    FOR SELECT
    USING (true);

-- Простой и премиум пользователь могут изменять только своих пользователей
DROP POLICY IF EXISTS users_update ON users;
CREATE POLICY users_update
    ON users
    FOR UPDATE
    TO simple_user, premium_user
    USING (uuid = (select user_uuid
                   from session));

-- Админ может изменять любого пользователя
DROP POLICY IF EXISTS users_update_admin ON users;
CREATE POLICY users_update_admin
    ON users
    FOR UPDATE
    TO admin
    USING (true);

-- Админ может удалять любого пользователя
DROP POLICY IF EXISTS users_delete_admin ON users;
CREATE POLICY users_delete_admin
    ON users
    FOR DELETE
    TO admin
    USING (true);

-- ===== Процедура, выполняемая перед каждым запросом =====
DROP PROCEDURE IF EXISTS before_each_query(user_uuid uuid);
CREATE OR REPLACE PROCEDURE before_each_query(IN user_uuid uuid)
    LANGUAGE plpgsql
AS
$$
DECLARE
    role_name regrole;
BEGIN
    role_name = (SELECT role FROM users WHERE users.uuid = user_uuid)::regrole;
    execute format('SET role %I', role_name);
    execute format('SET app.user_uuid = %L', user_uuid);
END;
$$;
-- =================== При запуске сервера с включенной опцией SQL миграции (конец блока) ===================

-- =================== При создании пользователя ===================
-- Создание выполняется служебным суперпользователем
SET ROLE superuser;
-- =================== При создании пользователя (конец блока) ===================

-- =================== При каждом запросе ===================
CALL before_each_query(?);
-- =================== Перед каждым запросом (конец блока) ===================
