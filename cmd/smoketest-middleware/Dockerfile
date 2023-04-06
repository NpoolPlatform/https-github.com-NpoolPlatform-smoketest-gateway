FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY smoketest-middleware /app

ENTRYPOINT ["/app/smoketest-middleware"]

EXPOSE 50330 50331 50332 50333

CMD ["run"]
