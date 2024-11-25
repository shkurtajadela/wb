--- Часть 1
-- без оконных функций
SELECT 
    s.first_name,
    s.last_name,
    s.salary,
    s.industry,
    (SELECT first_name || ' ' || last_name
     FROM Salary s_sub
     WHERE s_sub.industry = s.industry
     AND s_sub.salary = (SELECT MAX(salary) FROM Salary WHERE industry = s.industry)
    ) AS name_highest_sal
FROM 
    Salary s;

-- с использованием оконных функций
SELECT 
    first_name,
    last_name,
    salary,
    industry,
    FIRST_VALUE(first_name || ' ' || last_name) OVER (PARTITION BY industry ORDER BY salary DESC) AS name_highest_sal
FROM 
    Salary;

-- без оконных функций
SELECT 
    s.first_name,
    s.last_name,
    s.salary,
    s.industry,
    (SELECT first_name || ' ' || last_name
     FROM Salary s_sub
     WHERE s_sub.industry = s.industry
     AND s_sub.salary = (SELECT MIN(salary) FROM Salary WHERE industry = s.industry)
    ) AS name_lowest_sal
FROM 
    Salary s;
	
-- с использованием оконных функций
SELECT 
    first_name,
    last_name,
    salary,
    industry,
    FIRST_VALUE(first_name || ' ' || last_name) OVER (PARTITION BY industry ORDER BY salary ASC) AS name_lowest_sal
FROM 
    Salary;



--- Часть 2.1
SELECT DISTINCT 
    s."SHOPNUMBER",
    sh."CITY",
    sh."ADDRESS",
    SUM(s."QTY") OVER (PARTITION BY s."SHOPNUMBER") AS "SUM_QTY", 
    SUM(s."QTY" * g."PRICE") OVER (PARTITION BY s."SHOPNUMBER") AS "SUM_QTY_PRICE" 
FROM 
    SALES s
JOIN 
    SHOPS sh
ON 
    s."SHOPNUMBER" = sh."SHOPNUMBER"
JOIN 
    GOODS g
ON 
    s."ID_GOOD" = g."ID_GOOD"
WHERE 
    s."DATE" = '2016-01-02'; 


--- Часть 2.2
SELECT 
    s."DATE" AS date_,
    sh."CITY" AS city,
    SUM(s."QTY" * g."PRICE") AS sum_sales, 
    SUM(s."QTY" * g."PRICE") / SUM(SUM(s."QTY" * g."PRICE")) OVER (PARTITION BY s."DATE") AS sum_sales_rel
FROM 
    sales s
JOIN 
    goods g
ON 
    s."ID_GOOD" = g."ID_GOOD"
JOIN 
    shops sh
ON 
    s."SHOPNUMBER" = sh."SHOPNUMBER"
WHERE 
    g."CATEGORY" = 'ЧИСТОТА'
GROUP BY 
    s."DATE", sh."CITY"
ORDER BY 
    s."DATE", sh."CITY";

--- Часть 2.3
WITH ranked_sales AS (
    SELECT 
        s."DATE" AS date_,
        s."SHOPNUMBER" as shopnumber,
        s."ID_GOOD" as id_good,
        s."QTY",
        ROW_NUMBER() OVER (
            PARTITION BY s."DATE", s."SHOPNUMBER" 
            ORDER BY s."QTY" DESC
        ) AS rank
    FROM 
        sales s
)
SELECT 
    date_,
    shopnumber,
    id_good
FROM 
    ranked_sales
WHERE 
    rank <= 3
ORDER BY 
    date_, shopnumber, rank;

--- Часть 2.4
WITH sales_with_lag AS (
    SELECT 
        s."DATE" AS date_,
        s."SHOPNUMBER" as shopnumber,
        g."CATEGORY" as category,
        SUM(s."QTY" * g."PRICE") AS total_sales,
        LAG(SUM(s."QTY" * g."PRICE")) OVER (
            PARTITION BY s."SHOPNUMBER", g."CATEGORY"
            ORDER BY s."DATE"
        ) AS prev_sales 
    FROM 
        sales s
    JOIN 
        goods g
    ON 
        s."ID_GOOD" = g."ID_GOOD"
    JOIN 
        shops sh
    ON 
        s."SHOPNUMBER" = sh."SHOPNUMBER"
    WHERE 
        sh."CITY" = 'СПб' 
    GROUP BY 
        s."DATE", s."SHOPNUMBER", g."CATEGORY"
)
SELECT 
    date_,
    shopnumber,
    category,
    prev_sales
FROM 
    sales_with_lag
WHERE 
    prev_sales IS NOT NULL 
ORDER BY 
    date_, shopnumber, category;


--- Часть 3
CREATE TABLE query (
    searchid SERIAL PRIMARY KEY,
    year INT,
    month INT,
    day INT,
    userid INT,
    ts BIGINT, 
    devicetype VARCHAR(50),
    deviceid VARCHAR(50),
    query VARCHAR(255)
);

INSERT INTO query (year, month, day, userid, ts, devicetype, deviceid, query)
VALUES
(2023, 11, 18, 1, 1697659200, 'android', 'dev123', 'к'),
(2023, 11, 18, 1, 1697659260, 'android', 'dev123', 'ку'),
(2023, 11, 18, 1, 1697659320, 'android', 'dev123', 'куп'),
(2023, 11, 18, 1, 1697659500, 'android', 'dev123', 'купить кур'),
(2023, 11, 18, 1, 1697659560, 'android', 'dev123', 'купить куртку'),
(2023, 11, 18, 2, 1697659200, 'ios', 'dev002', 'с'),
(2023, 11, 18, 2, 1697659260, 'ios', 'dev002', 'сап'),
(2023, 11, 18, 2, 1697659300, 'ios', 'dev002', 'сапо'),
(2023, 11, 18, 2, 1697659400, 'ios', 'dev002', 'сапоги'),
(2023, 11, 18, 3, 1697659200, 'android', 'dev003', 'к'),
(2023, 11, 18, 3, 1697659260, 'android', 'dev003', 'кро'),
(2023, 11, 18, 3, 1697659300, 'android', 'dev003', 'крос'),
(2023, 11, 18, 3, 1697659400, 'android', 'dev003', 'кросс'),
(2023, 11, 18, 3, 1697659600, 'android', 'dev003', 'кроссовки'),
(2023, 11, 18, 4, 1697659200, 'desktop', 'dev004', 'п'),
(2023, 11, 18, 4, 1697659260, 'desktop', 'dev004', 'пи'),
(2023, 11, 18, 4, 1697659300, 'desktop', 'dev004', 'пидж'),
(2023, 11, 18, 4, 1697659400, 'desktop', 'dev004', 'пиджа'),
(2023, 11, 18, 4, 1697659500, 'desktop', 'dev004', 'пиджак');

WITH ranked_queries AS (
    SELECT 
        *,
        LEAD(ts) OVER (PARTITION BY userid, deviceid ORDER BY ts) AS next_ts, 
        LEAD(query) OVER (PARTITION BY userid, deviceid ORDER BY ts) AS next_query,
        LENGTH(query) AS current_length,
        LEAD(LENGTH(query)) OVER (PARTITION BY userid, deviceid ORDER BY ts) AS next_length
    FROM 
        query
),
is_final_calculated AS (
    SELECT 
        *,
        CASE
            --  после данного запроса больше ничего не искал
            WHEN next_ts IS NULL THEN 1
            -- до следующего запроса прошло более 3х минут
            WHEN (next_ts - ts) > 180 THEN 1
	        -- след запрос короче, И до него прошло прошло более минуты
            WHEN (next_ts - ts) > 60 AND next_length < current_length THEN 2
            -- иначе
            ELSE 0
        END AS is_final
    FROM 
        ranked_queries
)
SELECT 
    year,
    month,
    day,
    userid,
    ts,
    devicetype,
    deviceid,
    query,
    next_query,
    is_final
FROM 
    is_final_calculated
WHERE 
    day = 18 AND devicetype = 'android' AND is_final IN (1, 2) 
ORDER BY 
    ts;
