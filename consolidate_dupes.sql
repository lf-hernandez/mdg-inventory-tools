BEGIN;
WITH summed_quantities AS (
	SELECT
		part_number,
		sum(
			CASE WHEN quantity >= 0 THEN
				quantity
			ELSE
				- quantity
			END) AS total_quantity
	FROM
		item
	GROUP BY
		part_number
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
		item.part_number = sq.part_number
		AND item.id = (
			SELECT
				id
			FROM
				item i2
			WHERE
				i2.part_number = item.part_number
			LIMIT 1)
	RETURNING
		item.part_number)
DELETE FROM item
WHERE part_number IN (
		SELECT
			part_number
		FROM
			updated)
	AND id NOT IN (
		SELECT
			id
		FROM
			item i2
		WHERE
			i2.part_number = item.part_number
		LIMIT 1);
COMMIT;
