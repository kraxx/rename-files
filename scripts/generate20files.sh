#! /bin/bash
seq -w 1 20 | xargs -n1 -I% sh -c 'dd if=/dev/urandom of=file.% bs=$(shuf -i1-10 -n1) count=1024'