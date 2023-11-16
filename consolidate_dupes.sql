BEGIN;
WITH summed_quantities AS (
	SELECT
		external_id,
		sum(
			CASE WHEN quantity >= 0 THEN
				quantity
			ELSE
				- quantity
			END) AS total_quantity
	FROM
		item
	GROUP BY
		external_id
	HAVING
		count(*) > 1
),
updated AS (
	UPDATE
		item
	SET
		quantity = sq.total_quantity
	FROM
		summed_quantities sq
	WHERE
		item.external_id = sq.external_id
		AND item.id = (
			SELECT
				id
			FROM
				item i2
			WHERE
				i2.external_id = item.external_id
			LIMIT 1)
	RETURNING
		item.external_id)
DELETE FROM item
WHERE external_id IN (
		SELECT
			external_id
		FROM
			updated)
	AND id NOT IN (
		SELECT
			id
		FROM
			item i2
		WHERE
			i2.external_id = item.external_id
		LIMIT 1);
COMMIT;
