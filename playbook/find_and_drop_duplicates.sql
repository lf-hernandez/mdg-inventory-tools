/*
 * SQL Play - Find duplicate inventory records
 * This should no longer be needed after clean up effort
 * constraints will be added and guards implemented such that
 * neither API or client will allow the creation of a dupe
 */

-- Find total pn duplicate count
SELECT COUNT(*) AS part_number_duplicates
FROM (
    SELECT
        part_number,
        COUNT(part_number) AS count_pn
    FROM
        item
    GROUP BY
        part_number
    HAVING
        COUNT(part_number) > 1
) AS pn_duplicates;

-- Find total sn duplicate count
SELECT COUNT(*) AS serial_number_duplicates
FROM (
    SELECT
        serial_number,
        COUNT(serial_number) AS count_sn
    FROM
        item
    GROUP BY
        serial_number
    HAVING
        COUNT(serial_number) > 1
) AS sn_duplicates;

-- List the dupes by pn
SELECT part_number, count(part_number)
FROM item
GROUP BY part_number
HAVING COUNT(part_number) > 1;

-- List the dupes by sn
SELECT serial_number, count(serial_number)
FROM item
GROUP BY serial_number
HAVING COUNT(serial_number) > 1;

-- Display all columns in pn dupes to see diff
SELECT *
FROM item
WHERE part_number IN (
    SELECT part_number
    FROM item
    GROUP BY part_number
    HAVING COUNT(part_number) > 1
)
ORDER BY part_number;


-- Display all columns in sn dupes to see diff
SELECT *
FROM item
WHERE serial_number IN (
    SELECT serial_number
    FROM item
    GROUP BY serial_number
    HAVING COUNT(serial_number) > 1
)
ORDER BY serial_number;


-- DANGER: Be careful, point of no return son
DELETE FROM item
WHERE id IN (
    SELECT id
    FROM item
    WHERE part_number IN (
 	    SELECT part_number
 	    FROM item
	    GROUP BY part_number
	    HAVING COUNT(part_number) > 1
    )
    ORDER BY part_number
  );

-- Let's try to prevent this from happening again
ALTER TABLE item
ADD CONSTRAINT unique_part_number UNIQUE (part_number);
