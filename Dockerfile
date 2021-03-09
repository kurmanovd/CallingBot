FROM golang:buster

RUN echo "Switch to Python 3.7" \
	&& rm -f /usr/bin/python && ln -s /usr/bin/python /usr/bin/python3

RUN echo "installing dependencies" \
	&& apt-get update && apt-get install -y build-essential libcurl4-openssl-dev cmake pkg-config libasound2-dev \
	&& apt-get -y install libssl-dev git

RUN echo "Copying sources" \
	&& mkdir -p /app

COPY . /app

RUN echo "building VoIP" \
	&& cd /app/voip \
	&& cp include/config_site.h  pjproject/pjlib/include/pj/config_site.h \
	&& cd pjproject && ./configure --disable-libwebrtc --disable-opencore-amr \
	&& make dep && make && make install \
	&& cd .. && cmake CMakeLists.txt && make

RUN mkdir /output

RUN echo "building Server" \
	&& cd /app \
	&& go build -o bin/voip-backend cmd/api/main.go
 
RUN echo "Copy entrypoints" \
	&& cp /app/voip/docker/entry.sh / \
	&& cp /app/bin/voip-backend /

CMD ["/voip-backend"]