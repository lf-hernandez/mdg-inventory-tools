SELECT
	part_number,
	count(part_number)
FROM
	item
GROUP BY
	part_number
HAVING
	count(part_number) > 1;
