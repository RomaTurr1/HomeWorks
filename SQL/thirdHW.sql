-- 1)
SELECT 
    p.product_title,
    p.price,
    c.category_title
FROM products p
JOIN categories c ON p.category_id = c.category_id
ORDER BY c.category_title, p.price DESC;

-- 2)
SELECT 
    o.order_id,
    o.created_at
FROM orders o
WHERE o.user_name = 'Platon'
ORDER BY o.created_at ASC;

-- 3)
SELECT 
    o.order_id,
    o.user_name,
    p.product_title,
    oi.qty,
    oi.qty * p.price AS line_sum
FROM orders o
JOIN order_items oi ON o.order_id = oi.order_id
JOIN products p ON oi.product_id = p.product_id
ORDER BY o.order_id, p.product_title;

-- 4)
SELECT 
    o.order_id,
    o.user_name,
    SUM(oi.qty * p.price) AS order_total
FROM orders o
JOIN order_items oi ON o.order_id = oi.order_id
JOIN products p ON oi.product_id = p.product_id
GROUP BY o.order_id, o.user_name
ORDER BY order_total DESC;

-- 5)
SELECT 
    p.product_title,
    c.category_title
FROM products p
JOIN categories c ON p.category_id = c.category_id
LEFT JOIN order_items oi ON p.product_id = oi.product_id
WHERE oi.product_id IS NULL;