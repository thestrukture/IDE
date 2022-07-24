FROM golang:1.17.6
ENV WEBAPP /go/src/IDE
ENV HOME=/home/ide
RUN mkdir -p ${WEBAPP}
COPY . ${WEBAPP}
ENV PORT=8884 
ENV GOPATH=/go
WORKDIR ${WEBAPP}
RUN go install
RUN rm -rf ${WEBAPP}
# untested
VOLUME /home
RUN mkdir -p ${HOME}
EXPOSE 8884
CMD IDE --headless
