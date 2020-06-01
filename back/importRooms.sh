rm -f ./db.sqlite3
go run ./cmd/studi-guide-ctl migrate import rooms ./internal/rooms.json
cp -f ./db.sqlite3 ../bin