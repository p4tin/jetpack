#!/bin/bash
filename="$1"
while read -r line
do
    name=$line
    url="https://s3-us-west-2.amazonaws.com/taxrates.csv/TAXRATES_ZIP5_"$name"201506.csv"
    local_file="TAXRATES_ZIP5_"$name"201506.csv"
    wget $url
    gawk -F',' 'NR>1{gsub(/; +/,";",$0);printf("{_id:\"%s\",zip:\"%s\",state:\"%s\",tax:\"%s\"}\n",$2,$2,$1,$5)}' $local_file | mongoimport -d "GeoDB" -c "TaxInfo" --upsert
done < "$filename"

