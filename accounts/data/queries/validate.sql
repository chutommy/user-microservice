SELECT hpassword
FROM account
WHERE id = $1 AND deleted_at == NULL
LIMIT 1;
