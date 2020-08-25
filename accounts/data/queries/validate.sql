select hpassword
from account
where id = $1 and deleted_at = null
limit 1;
