
## testing install
go install -v github.com/cweill/gotests/gotests@v1.6.0

## generate test code by VSCODE
1. 該当のファイルの上で、コマンドバレットを開く「Command + Shift + p」(Mac)
2. 「Go: Generate Unit Tests For Function」を選択
3. テストファイルが自動される

##　Write test case
上記で生成したテストコード内に、テストケースを記述する
※サンプル1(カバレッジ率80%のケース, 通常の関数タイプ)
		{
			name: "return 'Hello DDD!!' if you set 'DDD'",
			args: args{
				"DDD",
			},
			want: "Hello DDD!!",
		},
		{
			name: "空文字を指定したら Hello が返ってくる",
			args: args{
				"",
			},
			want: "Hello",
		},
※サンプル2(カバレッジ率80%のケース, レシーバー関数タイプ)
		{
			name: "return 'Hello DDD!!' if you set 'DDD'",
			fields: fields{
				Say: "DDD",
			},
			want: "Hello DDD!!",
		},
		{
			name: "空文字を指定したら Hello が返ってくる",
			fields: fields{
				Say: "",
			},
			want: "Hello",
		},

## Unit testing
1. unit test
コマンド: go test -v -cover -coverprofile=[カバレッジ出力ファイルPath]] [パッケージPath]
例1(package): go test -v -cover -coverprofile=./coverage/coverage.txt ./hello
例2(all): go test -v -cover -coverprofile=./coverage/coverage.txt ./...

参考:
go test -v -cover -coverprofile=./coverage/coverage.txt ./...  
        unittest                coverage: 0.0% of statements
=== RUN   TestGetHello1
=== RUN   TestGetHello1/return_'Hello1_DDD!!'_if_you_set_'DDD'
=== RUN   TestGetHello1/空文字を指定したら_Hello1_が返ってくる
--- PASS: TestGetHello1 (0.00s)
    --- PASS: TestGetHello1/return_'Hello1_DDD!!'_if_you_set_'DDD' (0.00s)
    --- PASS: TestGetHello1/空文字を指定したら_Hello1_が返ってくる (0.00s)
=== RUN   TestHello_GetHello2
=== RUN   TestHello_GetHello2/return_'Hello2_DDD!!'_if_you_set_'DDD'
=== RUN   TestHello_GetHello2/空文字を指定したら_Hello2_が返ってくる
--- PASS: TestHello_GetHello2 (0.00s)
    --- PASS: TestHello_GetHello2/return_'Hello2_DDD!!'_if_you_set_'DDD' (0.00s)
    --- PASS: TestHello_GetHello2/空文字を指定したら_Hello2_が返ってくる (0.00s)
PASS
coverage: 80.0% of statements
ok      unittest/hello  0.234s  coverage: 80.0% of statements

## Unit test Coverage check
1. カバレッジHTML生成
コマンド: go tool cover -html=[カバレッジ出力ファイルPath] -o [カバレッジ出力HTMLファイルPath]  
例: go tool cover -html=./coverage/coverage.txt -o ./coverage/coverage.html
2. カバレッジHTML確認
コマンド: open [カバレッジ出力HTMLファイルPath]
例: open ./coverage/coverage.html 

## test cache のクリア
コマンド: go clean -testcache