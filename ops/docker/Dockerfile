FROM scratch
COPY ./bin/app /
# this is to make appuser work
COPY ./ops/docker/passwd /etc/passwd
# From here onwards, any RUN, CMD, or ENTRYPOINT will be run under the following user instead of 'root'
USER appuser

ENTRYPOINT ["./app"]
