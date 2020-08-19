DELETE FROM accounts
WHERE id = $1 AND deleted_at = NULL;
