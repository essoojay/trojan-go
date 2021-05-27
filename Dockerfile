FROM golang:alpine AS builder
WORKDIR /
ARG REF
RUN apk add git make &&\
    git clone https://github.com/frainzy1477/trojan-go.git
RUN if [[ -z "${REF}" ]]; then \
        echo "No specific commit provided, use the latest one." \
    ;else \
        echo "Use commit ${REF}" &&\
        cd trojan-goo &&\
        git checkout ${REF} \
    ;fi
RUN cd trojan-goo &&\
    make &&\
    wget https://github.com/v2fly/domain-list-community/raw/release/dlc.dat -O build/geosite.dat &&\
    wget https://github.com/v2fly/geoip/raw/release/geoip.dat -O build/geoip.dat

FROM alpine
WORKDIR /
RUN apk add --no-cache tzdata ca-certificates
COPY --from=builder /trojan-goo/build /usr/local/bin/
RUN mv /usr/local/bin/trojan-goo /usr/local/bin/trojan-go

ENTRYPOINT ["/usr/local/bin/trojan-go", "-config"]
CMD ["/etc/trojan-go/*.json"]
