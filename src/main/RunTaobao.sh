#!/usr/bin/env bash

today=$(date +%F)


./GoTaoBao

cd 搜索结果/${today//-//}

echo $(pwd)


tar -czvf taobao_data_${today}.tar.gz  ./*

/usr/bin/ftp -n <<EOF
open 123.126.65.199 5021
user chliu ZKaCH974yl5NQ3Yq
binary
prompt off
put taobao_data_${today}.tar.gz
bye
EOF