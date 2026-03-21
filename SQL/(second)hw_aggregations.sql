INSERT INTO categories (category_id, name)
VALUES (100, 'New Category');


INSERT INTO products (product_id, name, price, category_id) VALUES
(201, 'Product A', 12000, 100),
(202, 'Product B', 15000, 100),
(203, 'Product C', 5000, 1),
(204, 'Product D', 20000, 1),
(205, 'Product E', 3000, 2);


INSERT INTO orders (order_id, customer, total) VALUES
(301, 'Ramazan', 2500),
(302, 'Ramazan', 90000),
(303, 'Aidar', 15000),
(304, 'Aidar', 20000),
(305, 'Dana', 80000),
(306, 'Ali', 5000);


-- 1)
SELECT 
    category_id,
    COUNT(*) AS total_products
FROM products
GROUP BY category_id
ORDER BY total_products DESC;


-- 2)
SELECT 
    category_id,
    COUNT(*) AS total_products
FROM products
GROUP BY category_id
HAVING COUNT(*) >= 3;


-- 3)
SELECT 
    category_id,
    MIN(price) AS min_price,
    MAX(price) AS max_price,
    ROUND(AVG(price), 2) AS avg_price
FROM products
GROUP BY category_id;


-- 4)
SELECT 
    category_id,
    COUNT(*) AS expensive_count
FROM products
WHERE price > 10000
GROUP BY category_id
HAVING COUNT(*) >= 2;


-- 5)
SELECT 
    customer,
    COUNT(*) AS orders_count
FROM orders
GROUP BY customer
ORDER BY orders_count DESC;


-- 6)
SELECT 
    customer,
    COUNT(*) AS orders_count
FROM orders
GROUP BY customer
HAVING COUNT(*) >= 2;


-- 7)
SELECT 
    customer,
    SUM(total) AS spent_total,
    ROUND(AVG(total), 2) AS avg_check
FROM orders
GROUP BY customer
ORDER BY spent_total DESC;