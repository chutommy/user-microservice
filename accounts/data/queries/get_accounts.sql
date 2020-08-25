select id, username, email, phone, first_name, last_name, birth_day, perm_address, mail_address, created_at, updated_at
from accounts
where deleted_at = null
order by $3 $4
offset $2
limit $1;
