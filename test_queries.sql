SELECT * FROM weather;

SELECT city, temp_lo, temp_hi, prcp, date FROM weather;

SELECT * FROM weather
    ORDER BY city;

SELECT * FROM weather
    ORDER BY city, temp_lo;

SELECT * FROM weather
    WHERE city = 'San Francisco' AND prcp > 0.0;

CREATE TABLE cities (
    name            text,
    population      real,
    altitude        int     -- (in ft)
);

CREATE TABLE capitals (
    state           char(2)
) INHERITS (cities);

SELECT pg_size_pretty(pg_relation_size('big_table'));

INSERT INTO numbers (num) VALUES ( generate_series(1,1000));

SELECT W1.city, W1.temp_lo AS low, W1.temp_hi AS high,
    W2.city, W2.temp_lo AS low, W2.temp_hi AS high
    FROM weather W1, weather W2
    WHERE W1.temp_lo < W2.temp_lo
    AND W1.temp_hi > W2.temp_hi;

SELECT
 A.pka,
 A.c1,
 B.pkb,
 B.c2
FROM
 A
INNER JOIN B ON A .pka = B.fka;
