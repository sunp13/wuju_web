FROM centos:7

ENV TZ=Asia/shanghai

COPY ./conf /app/conf
COPY ./logs /app/logs
COPY ./static /app/static
COPY ./views /app/views
COPY ./wuju_web /app
WORKDIR /app

EXPOSE 8080

CMD [ "./wuju_web" ]


