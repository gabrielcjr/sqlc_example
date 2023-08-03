-- name: ListCategories :many
SELECT * FROM categories;
-- name: LisrPtoductsFromCategoryId :many
SELECT * FROM produtos WHERE category_id = ?;