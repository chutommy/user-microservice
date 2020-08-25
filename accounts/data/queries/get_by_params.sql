select id, username, email, phone, first_name, last_name, birth_day, perm_address, mail_address, created_at, updated_at
from accounts
where deleted_at == null
and ($1 = "" or $1 = id)
and ($2 = "" or $2 = username)
and ($3 = "" or $3 = email)
and ($4 = "" or $4 = phone)
and ($5 = "" or $5 = first_name)
and ($6 = "" or $6 = last_name)
and ($7 = "" or $7 = birth_day)
and ($8 = "" or $8 = perm_address)
and ($9 = "" or $9 = mail_address)
and ($10 = "" or $10 = created_at)
and ($11 = "" or $11 = updated_at);
