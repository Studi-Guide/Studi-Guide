rm -f ./db.sqlite3
go run ./cmd/studi-guide-ctl migrate import campus ./internal/campus.json
go run ./cmd/studi-guide-ctl migrate import rooms ./internal/rooms.json
go run ./cmd/studi-guide-ctl migrate import rssfeed ./internal/rssfeeds.json
go run ./cmd/studi-guide-ctl migrate import location ./internal/location.json
