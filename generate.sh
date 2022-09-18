#!/bin/bash

rm -R types/

for file in schemas/*
do
    ocpp_message_name=${file/.json/""}
    generated_file_name=$ocpp_message_name."go"
    gojsonschema -p $ocpp_message_name "$file" >> $generated_file_name
done

mkdir temp
mv schemas/*.go temp

for file in temp/*.go
do
    dir_name=${file/.go/""}
    mkdir "$dir_name" 
    mv "$file" "$dir_name"
done

mkdir -p types
mv temp/* types
rm -R temp
