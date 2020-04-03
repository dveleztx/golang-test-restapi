#!/bin/bash
###############################################################################
# Script      : go-get-libs.sh
# Author      : David Velez
# Date        : 03/20/2020
# Description : Retrieve Golang Libraries
###############################################################################

# Variables
golibs=(
	"go get github.com/labstack/echo"
	"go get github.com/jinzhu/gorm"
	"go get github.com/jinzhu/gorm/dialects/postgres"
	"go get github.com/dgrijalva/jwt-go"
)

# Iterate
for libs in "${golibs[@]}"
do
	eval "$libs"
done
