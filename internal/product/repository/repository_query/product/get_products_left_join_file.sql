SELECT
    p.id,
    p.product_name,
    p.product_description,
    p.product_price,
    p.created_at,
    p.updated_at,
    p.stock_qty,
    f.file_name,
    f.file_data,
    f.file_type
FROM
    Product p
    LEFT JOIN File f ON p.file_fk_id = f.id;