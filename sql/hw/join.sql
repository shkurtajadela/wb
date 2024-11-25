-- Часть 1.1
SELECT 
    c.customer_id,
    c.name,
    o.order_id,
    o.order_date,
    o.shipment_date,
    ROUND(EXTRACT(EPOCH FROM (o.shipment_date - o.order_date)) / (24 * 60 * 60),2) AS waiting_days 
FROM 
    Customers c
JOIN 
    Orders o
ON 
    c.customer_id = o.customer_id
ORDER BY 
    waiting_days DESC
LIMIT 1;

-- Часть 1.2
WITH customer_order_data AS (
    SELECT 
        c.customer_id,
        c.name,
        COUNT(o.order_id) AS total_orders, 
        ROUND(AVG(EXTRACT(EPOCH FROM (o.shipment_date - o.order_date)) / (24 * 60 * 60)), 2) AS avg_waiting_days,
        SUM(o.order_ammount) AS total_order_amount
    FROM 
        Customers c
    JOIN 
        Orders o
    ON 
        c.customer_id = o.customer_id
    GROUP BY 
        c.customer_id, c.name
)
SELECT 
    customer_id,
    name,
    total_orders,
    avg_waiting_days,
    total_order_amount
FROM 
    customer_order_data
WHERE 
    total_orders = (SELECT MAX(total_orders) FROM customer_order_data)
ORDER BY 
    total_order_amount DESC;



-- Часть 1.3
SELECT 
    c.customer_id,
    c.name,
    SUM(CASE WHEN EXTRACT(EPOCH FROM (o.shipment_date - o.order_date)) / (24 * 60 * 60) > 5 THEN 1 ELSE 0 END) AS delayed_deliveries,
    SUM(CASE WHEN o.order_status = 'Cancel' THEN 1 ELSE 0 END) AS canceled_orders, 
    SUM(o.order_ammount) AS total_order_amount 
FROM 
    Customers c
JOIN 
    Orders o
ON 
    c.customer_id = o.customer_id
GROUP BY 
    c.customer_id, c.name
HAVING 
    SUM(CASE WHEN EXTRACT(EPOCH FROM (o.shipment_date - o.order_date)) / (24 * 60 * 60) > 5 THEN 1 ELSE 0 END) > 0
    OR SUM(CASE WHEN o.order_status = 'Cancel' THEN 1 ELSE 0 END) > 0 
ORDER BY 
    total_order_amount DESC;


-- Часть 2.1
SELECT 
    p.product_category,
    SUM(o.order_ammount) AS total_sales
FROM 
    products_3 p
JOIN 
    orders_2 o
ON 
    p.product_id = o.product_id
GROUP BY 
    p.product_category;

-- Часть 2.2
SELECT 
    product_category,
    total_sales
FROM (
    SELECT 
        p.product_category,
        SUM(o.order_ammount) AS total_sales
    FROM 
        products_3 p
    JOIN 
        orders_2 o
    ON 
        p.product_id = o.product_id
    GROUP BY 
        p.product_category
) AS category_totals
ORDER BY 
    total_sales DESC
LIMIT 1;

-- Часть 2.3
SELECT 
    p.product_category,
    p.product_name,
    MAX(o.order_ammount) AS max_sales_per_product
FROM 
    products_3 p
JOIN 
    orders_2 o
ON 
    p.product_id = o.product_id
WHERE 
    o.order_ammount = (
        SELECT 
            MAX(o2.order_ammount)
        FROM 
            orders_2 o2
        JOIN 
            products_3 p2
        ON 
            p2.product_id = o2.product_id
        WHERE 
            p2.product_category = p.product_category
    )
GROUP BY 
    p.product_category, p.product_name
ORDER BY 
    p.product_category;
