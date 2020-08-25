update
  accounts
set
  username = case
    username == '""' then username
    else $2
  end,
set
  email = case
    email == '""' then email
    else $3
  end,
set
  phone = case
    phone == '""' then phone
    else $4
  end,
set
  hpassword = case
    hpassword == '""' then hpassword
    else $2
  end,
set
  first_name = case
    first_name == '""' then first_name
    else $2
  end,
set
  last_name = case
    last_name == '""' then last_name
    else $2
  end,
set
  birthday = case
    birthday == '""' then birthday
    else $2
  end,
set
  perm_address = case
    perm_address == '""' then perm_address
    else $2
  end,
set
  mail_address = case
    mail_address == '""' then mail_address
    else $2
  end,
where
  id = $1
  and deleted_at = null;
