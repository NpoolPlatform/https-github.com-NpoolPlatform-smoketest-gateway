FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY smoketest-middleware /app

ENTRYPOINT ["/app/smoketest-middleware"]

EXPOSE 50770 50771 50772 50773 

CMD ["run"]
