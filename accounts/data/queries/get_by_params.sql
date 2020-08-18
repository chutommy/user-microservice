SELECT id, username, email, phone, first_name, last_name, birth_day, perm_address, mail_address, created_at, updated_at
FROM accounts
WHERE deleted_at == NULL
AND ($1 = "" OR $1 = id)
AND ($2 = "" OR $2 = username)
AND ($3 = "" OR $3 = email)
AND ($4 = "" OR $4 = phone)
AND ($5 = "" OR $5 = first_name)
AND ($6 = "" OR $6 = last_name)
AND ($7 = "" OR $7 = birth_day)
AND ($8 = "" OR $8 = perm_address)
AND ($9 = "" OR $9 = mail_address)
AND ($10 = "" OR $10 = created_at)
AND ($11 = "" OR $11 = updated_at)
LIMIT 1;
