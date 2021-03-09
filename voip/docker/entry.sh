#!/bin/bash

XML_CONF=${XML_CONF-"config.xml"}
RESULT_FILE=${RESULT_FILE-"result.json"}
PORT=${PORT="5060"}

echo "XML_CONF[${XML_CONF}] RESULT FILE[${RESULT_FILE}] PORT[$PORT]"

if [ "$1" = "" ]; then
	CMD="/app/voip/voip_patrol --port ${PORT} --conf /config/${XML_CONF} --output /output/${RESULT_FILE}"
else
	CMD="$*"
fi

echo "Running ${CMD}"
exec ${CMD} > /output/log 2>&1