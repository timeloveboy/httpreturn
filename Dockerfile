FROM go
MAINTAINER timeloveboy(734991033@qq.com)
ADD httpreturn /usr/local/bin/
EXPOSE 9000
CMD httpreturn hello