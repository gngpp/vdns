#!/bin/sh

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o linux_amd64_vdns ./main.go
mv ./linux_amd64_vdns ./vdns && upx ./vdns
tar -czvf linux_amd64_vdns.tar.gz ./vdns && rm ./vdns

CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o linux_386_vdns ./main.go
mv ./linux_386_vdns ./vdns && upx ./vdns
tar -czvf linux_386_vdns.tar.gz ./vdns && rm ./vdns

# not upx
CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -ldflags="-s -w" -o freebsd_386_vdns ./main.go
mv ./freebsd_386_vdns ./vdns
tar -czvf freebsd_386_vdns.tar.gz ./vdns && rm ./vdns

# not upx
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o freebsd_amd64_vdns ./main.go
mv ./freebsd_amd64_vdns ./vdns
tar -czvf freebsd_amd64_vdns.tar.gz ./vdns && rm ./vdns

# not upx
CGO_ENABLED=0 GOOS=freebsd GOARCH=arm go build -ldflags="-s -w" -o freebsd_arm64_vdns ./main.go
mv ./freebsd_arm64_vdns ./vdns
tar -czvf freebsd_arm64_vdns.tar.gz ./vdns && rm ./vdns

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o linux_armv7_vdns ./main.go
mv ./linux_armv7_vdns ./vdns && upx ./vdns
tar -czvf linux_armv7_vdns.tar.gz ./vdns && rm ./vdns

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-s -w" -o linux_armv6_vdns ./main.go
mv ./linux_armv6_vdns ./vdns && upx ./vdns
tar -czvf linux_armv6_vdns.tar.gz ./vdns && rm ./vdns

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -o linux_armv5_vdns ./main.go
mv ./linux_armv5_vdns ./vdns && upx ./vdns
tar -czvf linux_armv5_vdns.tar.gz ./vdns && rm ./vdns

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o linux_arm64_vdns ./main.go
mv ./linux_arm64_vdns ./vdns && upx ./vdns
tar -czvf linux_arm64_vdns.tar.gz ./vdns && rm ./vdns

# not upx
CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -ldflags="-s -w" -o linux_mips64_vdns ./main.go
mv ./linux_mips64_vdns ./vdns
tar -czvf linux_mips64_vdns.tar.gz ./vdns && rm ./vdns

# not upx
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -ldflags="-s -w" -o linux_mips64le_vdns ./main.go
mv ./linux_mips64le_vdns ./vdns
tar -czvf linux_mips64le_vdns.tar.gz ./vdns && rm ./vdns

CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -ldflags="-s -w" -o linux_mipsle_vdns ./main.go
mv ./linux_mipsle_vdns ./vdns && upx ./vdns
tar -czvf linux_mipsle_vdns.tar.gz ./vdns && rm ./vdns

CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -ldflags="-s -w" -o linux_mips_vdns ./main.go
mv ./linux_mips_vdns ./vdns && upx ./vdns
tar -czvf linux_mips_vdns.tar.gz ./vdns && rm ./vdns

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o windows_386_vdns.exe ./main.go
mv ./windows_386_vdns.exe ./vdns.exe && upx ./vdns.exe
tar -czvf windows_386_vdns.tar.gz ./vdns.exe && rm ./vdns.exe

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o windows_amd64_vdns.exe ./main.go
mv ./windows_amd64_vdns.exe ./vdns.exe && upx ./vdns.exe
tar -czvf windows_amd64_vdns.tar.gz ./vdns.exe && rm ./vdns.exe

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o darwin_amd64_vdns ./main.go
mv ./darwin_amd64_vdns ./vdns && upx ./vdns
tar -czvf darwin_amd64_vdns.tar.gz ./vdns && rm ./vdns

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o darwin_arm64_vdns ./main.go
mv ./darwin_arm64_vdns ./vdns && upx ./vdns
tar -czvf darwin_arm64_vdns.tar.gz ./vdns && rm ./vdns