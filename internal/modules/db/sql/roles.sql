DROP role if exists simple_user;

create role admin;
create role simple_user;
create role premium_user;

grant all on all tables in schema public to admin;

revoke all on all tables in schema public from simple_user;
grant select on access_rights to simple_user;
grant all on events to simple_user;
grant all on invitations to simple_user;
grant all on events_tags to simple_user;
grant select, insert, update on users to simple_user;
grant select on tags to simple_user;

-- Разрешаем RLS
ALTER TABLE users ENABLE ROW LEVEL SECURITY;

create policy select_user_policy
ON users
FOR SELECT
USING (true);

-- Создаем политику для обычного пользователя, чтобы пользователь мог менять только свою запись в таблице users
create policy user_policy
    ON users
    FOR UPDATE
    TO simple_user
    USING (uuid = current_user::uuid);



drop policy user_policy on users;





