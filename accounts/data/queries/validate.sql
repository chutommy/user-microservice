SELECT hpassword
FROM account
WHERE id = $1
LIMIT 1;
