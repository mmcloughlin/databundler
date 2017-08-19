databundler -pkg sevensummits -data sevensummits.csv -schema schema.yaml -output sevensummits.go
gofmt -w sevensummits.go
