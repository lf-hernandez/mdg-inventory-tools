SELECT
	external_id,
	count(external_id)
FROM
	item
GROUP BY
	external_id
HAVING
	count(external_id) > 1;
