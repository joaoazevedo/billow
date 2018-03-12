FROM busybox

COPY ./bin/* /usr/sbin/

WORKDIR /usr/apps/web/

CMD ["/usr/sbin/fileserver", "-d", "/usr/apps/web/"]
