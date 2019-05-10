#! /bin/bash

create_vendor() {
  local name="${1}"
  shift
  local code="${1}"
  shift
  curl -s -X POST -H "Content-Type: application/json" \
    -d "{\"name\":\"${name}\",\"code\":\"${code}\"}" \
    http://localhost:3000/vendor/ |
    sed -e 's,[{}" ],,g' -e 's/,/\n/g' |
    gawk 'BEGIN{FS=":";} /uuid/ { print $2; }'
}

create_flavor() {
  local name="${1}"
  shift
  local vendor_uuid="${1}"
  shift
  curl -s -X POST -H "Content-Type: application/json" \
    -d "{\"name\":\"${name}\",\"vendor_uuid\":\"${vendor_uuid}\"}" \
    http://localhost:3000/flavor/
}

prev_vendor=''
while IFS=',' read -r code vendor flavor; do
  if [[ "${vendor}" != "${prev_vendor}" ]]; then
    prev_vendor="${vendor}"
    i="$(create_vendor "${vendor}" "${code}")"
    echo "${vendor}" "${i}"
  fi
  create_flavor "${flavor}" "${i}"
done <flavors.csv

# eof
