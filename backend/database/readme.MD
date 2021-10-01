db.go
- initialize connection to localhost database
- get current database connection

dump.sql
- use this file to init local database table and values


curl 'http://localhost:8000/products'
curl 'http://localhost:8000/product/1'
curl -X POST 'http://localhost:8000/product/insert' -d '{"name": "DVD Game Dota 3", "shop_name" : "Toko Baru Buka", "product_price" : 367000, "image_url" : "product123.png"}'
curl -X PUT 'http://localhost:8000/product/update' -d '{"id" : 1, "name": "DVD Game Resident Evil 4 Nemesis", "shop_name" : "Toko Baru Buka", "product_price" : 367000, "image_url" : "product123.png"}'