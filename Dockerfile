FROM alpine

ADD dist/searchzin /opt/searchzin/searchzin
ADD templates /opt/searchzin/templates

EXPOSE 8080

WORKDIR "/opt/searchzin"
CMD [ "/opt/searchzin/searchzin" ]
