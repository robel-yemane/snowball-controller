FROM debian
COPY ./snowball-controller /snowball-controller
ENTRYPOINT /snowball-controller

