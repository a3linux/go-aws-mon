FROM alpine:latest 
MAINTAINER Allen Chen(a3linux@gmail.com)

RUN apk --update add bash openssl 
COPY go-aws-mon /usr/bin/go-aws-mon 

CMD /usr/bin/go-aws-mon --mem-util --mem-used --mem-avail --disk-space-util --disk-inode-util --disk-space-used --disk-space-avail --disk-path=/
