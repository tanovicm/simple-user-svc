curl -XPOST -d '{"firstname":"Marijana","lastname":"Tanovic","nickname":"Mara","password":"123","email":"mara@mara","country":"Srbija"}' 'localhost:8000/users'
curl -XGET 'localhost:8000/users/61ffb1934697be8880d8b843'
curl -XDELETE 'localhost:8000/users/61ffb1934697be8880d8b843'
