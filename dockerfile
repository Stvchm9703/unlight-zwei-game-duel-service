FROM apline
WORKDIR /app
COPY ./account_server ./ 
COPY ./insecure ./insecure
COPY ./config.test_server.yaml ./config.yaml
EXPOSE 12000

ENTRYPOINT ["account_server"]
CMD [ "start", "-c" ,"config.yaml"]