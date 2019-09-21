FROM clearlinux/golang AS build
ADD . /go/src/pike
RUN go install pike

FROM clearlinux
COPY --from=build /go/bin/pike /usr/local/bin/pike
EXPOSE 8080
CMD exec /usr/local/bin/pike
