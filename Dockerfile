FROM ubuntu:20.04
WORKDIR vdns
COPY . .
RUN apt update && apt install software-properties-common -y && add-apt-repository ppa:longsleep/golang-backports -y
RUN apt install golang-go -y && apt install gcc-mingw-w64-i686 gcc-multilib -y
RUN sh ./build.sh