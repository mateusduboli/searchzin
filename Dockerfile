FROM alpine

ADD dist/searchzin /opt/searchzin

EXPOSE 8080

CMD [ "/opt/searchzin" ]
