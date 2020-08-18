SELECT id, username, email, phone, first_name, last_name, birth_day, perm_address, mail_address, created_at, updated_at
FROM accounts
WHERE id == $1 AND deleted_at == NULL
LIMIT 1;

