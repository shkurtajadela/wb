--- Часть 1.1
SELECT 
    city,
    CASE
        WHEN age BETWEEN 0 AND 20 THEN 'young'
        WHEN age BETWEEN 21 AND 49 THEN 'adult'
        WHEN age >= 50 THEN 'old'
    END AS age_category,
    COUNT(*) AS customer_count
FROM users
GROUP BY city, age_category
ORDER BY city, customer_count DESC;


-- Часть 1.2
SELECT 
    category,
    ROUND(CAST(AVG(price) AS NUMERIC), 2) AS avg_price
FROM 
    products
WHERE 
    LOWER(name) LIKE '%hair%' OR LOWER(name) LIKE '%home%'
GROUP BY 
    category;

-- Часть 2.1
SELECT 
    seller_id,
    COUNT(DISTINCT category) AS total_categ,
    ROUND(AVG(rating), 2) AS avg_rating,
    SUM(revenue) AS total_revenue,
    CASE 
        WHEN SUM(revenue) > 50000 AND COUNT(DISTINCT category) > 1 THEN 'rich'
        WHEN SUM(revenue) <= 50000 AND COUNT(DISTINCT category) > 1 THEN 'poor'
    END AS seller_type
FROM 
    sellers
WHERE 
    category != 'Bedding'
GROUP BY 
    seller_id
ORDER BY 
    seller_id;


-- Часть 2.2
WITH poor_sellers AS (
    SELECT 
        seller_id,
        MIN(date_reg) AS min_date_reg, 
        MAX(delivery_days) - MIN(delivery_days) AS max_delivery_difference,
        COUNT(DISTINCT category) AS total_categ,
        SUM(revenue) AS total_revenue
    FROM 
        sellers
    WHERE 
        category != 'Bedding' 
    GROUP BY 
        seller_id
    HAVING 
        COUNT(DISTINCT category) > 1 AND SUM(revenue) <= 50000
)
SELECT 
    seller_id,
    FLOOR((CURRENT_DATE - min_date_reg) / 30) AS month_from_registration, 
    max_delivery_difference
FROM 
    poor_sellers
ORDER BY 
    seller_id;


-- Часть 2.3
WITH seller_categories AS (
    SELECT 
        seller_id,
        ARRAY_AGG(DISTINCT category ORDER BY category) AS categories,
        SUM(revenue) AS total_revenue
    FROM 
        sellers
    WHERE 
        EXTRACT(YEAR FROM date_reg) = 2022
    GROUP BY 
        seller_id
    HAVING 
        COUNT(DISTINCT category) = 2
        AND SUM(revenue) > 75000 
	    AND EXTRACT(YEAR FROM MIN(date_reg)) = 2022
)
SELECT 
    seller_id,
    ARRAY_TO_STRING(categories, ' - ') AS category_pair
FROM 
    seller_categories
ORDER BY 
    seller_id;
