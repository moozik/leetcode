#!/bin/bash
pName=$1
mkdir "./problems/$pName"
echo "package "${pName} >> "./problems/$pName/${pName}.go"