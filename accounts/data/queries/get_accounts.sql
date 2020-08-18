SELECT id, username, email, phone, hpassword, first_name, last_name, birth_day, perm_address, mail_address, created_at, updated_at
FROM accounts
WHERE deleted_at == NULL
ORDER BY $3 $4
OFFSET $2
LIMIT $1;
