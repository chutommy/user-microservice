delete from
  accounts
where
  id = $1
  and deleted_at = null;
