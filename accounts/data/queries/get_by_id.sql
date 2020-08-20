select id, username, email, phone, first_name, last_name, birth_day, perm_address, mail_address, created_at, updated_at
from accounts
where id = $1 and deleted_at = null
limit 1;

