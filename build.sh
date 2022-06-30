#!/bin/sh

MAIN_PATH=./main.go
TARGET_PATH=./vdns
WINDOWS_TARGET_PATH=./vdns.exe

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o linux_amd64_vdns $MAIN_PATH
mv ./linux_amd64_vdns $TARGET_PATH && upx $TARGET_PATH
tar -czvf linux_amd64_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o linux_386_vdns $MAIN_PATH
mv ./linux_386_vdns $TARGET_PATH && upx $TARGET_PATH
tar -czvf linux_386_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

# not upx
CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -ldflags="-s -w" -o freebsd_386_vdns $MAIN_PATH
mv ./freebsd_386_vdns $TARGET_PATH
tar -czvf freebsd_386_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

# not upx
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o freebsd_amd64_vdns $MAIN_PATH
mv ./freebsd_amd64_vdns $TARGET_PATH
tar -czvf freebsd_amd64_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

# not upx
CGO_ENABLED=0 GOOS=freebsd GOARCH=arm go build -ldflags="-s -w" -o freebsd_arm64_vdns $MAIN_PATH
mv ./freebsd_arm64_vdns $TARGET_PATH
tar -czvf freebsd_arm64_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o linux_armv7_vdns $MAIN_PATH
mv ./linux_armv7_vdns $TARGET_PATH && upx $TARGET_PATH
tar -czvf linux_armv7_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-s -w" -o linux_armv6_vdns $MAIN_PATH
mv ./linux_armv6_vdns $TARGET_PATH && upx $TARGET_PATH
tar -czvf linux_armv6_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -o linux_armv5_vdns $MAIN_PATH
mv ./linux_armv5_vdns $TARGET_PATH && upx $TARGET_PATH
tar -czvf linux_armv5_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o linux_arm64_vdns $MAIN_PATH
mv ./linux_arm64_vdns $TARGET_PATH && upx $TARGET_PATH
tar -czvf linux_arm64_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

# not upx
CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -ldflags="-s -w" -o linux_mips64_vdns $MAIN_PATH
mv ./linux_mips64_vdns $TARGET_PATH
tar -czvf linux_mips64_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

# not upx
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -ldflags="-s -w" -o linux_mips64le_vdns $MAIN_PATH
mv ./linux_mips64le_vdns $TARGET_PATH
tar -czvf linux_mips64le_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -ldflags="-s -w" -o linux_mipsle_vdns $MAIN_PATH
mv ./linux_mipsle_vdns $TARGET_PATH && upx $TARGET_PATH
tar -czvf linux_mipsle_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -ldflags="-s -w" -o linux_mips_vdns $MAIN_PATH
mv ./linux_mips_vdns $TARGET_PATH && upx $TARGET_PATH
tar -czvf linux_mips_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o windows_386_vdns.exe $MAIN_PATH
mv ./windows_386_vdns.exe $WINDOWS_TARGET_PATH && upx $WINDOWS_TARGET_PATH
tar -czvf windows_386_vdns.tar.gz $WINDOWS_TARGET_PATH && rm $WINDOWS_TARGET_PATH

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o windows_amd64_vdns.exe $MAIN_PATH
mv ./windows_amd64_vdns.exe $WINDOWS_TARGET_PATH && upx $WINDOWS_TARGET_PATH
tar -czvf windows_amd64_vdns.tar.gz $WINDOWS_TARGET_PATH && rm $WINDOWS_TARGET_PATH

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o darwin_amd64_vdns $MAIN_PATH
mv ./darwin_amd64_vdns $TARGET_PATH && upx $TARGET_PATH
tar -czvf darwin_amd64_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o darwin_arm64_vdns $MAIN_PATH
mv ./darwin_arm64_vdns $TARGET_PATH && upx $TARGET_PATH
tar -czvf darwin_arm64_vdns.tar.gz $TARGET_PATH && rm $TARGET_PATH