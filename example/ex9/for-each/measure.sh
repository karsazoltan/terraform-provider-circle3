#!/bin/bash 

echo "Start"

for i in 10 20 40 80 160
do
    list="[\"1\""
    for k in $(seq 2 1 $i)
    do
        list="${list},\"${k}\""
    done
    list="${list}]"
    echo "Apply (${i}) ..."
    time terraform apply -var "list=${list}" -auto-approve -no-color
    sleep 8m
    echo "Destroy (${i}) ..."
    terraform destroy -var "list=${list}" -auto-approve -no-color
done

echo "End"