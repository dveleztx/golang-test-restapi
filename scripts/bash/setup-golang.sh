#!/bin/bash
###############################################################################
# Script      : setup-golang.sh
# Author      : David Velez
# Date        : 03/20/2020
# Description : Setup Golang directory structure and ENV variables into 
#               .bashrc. This has been tested on Debian/Ubuntu distros only.
###############################################################################

# Variables
USERBASHRC=$HOME/.bashrc
FILEPATH=$HOME/go/src

# Install Golang
sudo apt update && sudo apt upgrade -y
sudo apt install -y golang

# Setup Go Directory
if [ -f $FILEPATH ]; then
	echo "Filepath $FILEPATH exists."
else
	mkdir -p $HOME/go/src
fi

# Write ENV variables to bashrc
if grep -q "export GOROOT=/usr/lib/go" $USERBASHRC; then
	echo "GOROOT variable already set."
else
	echo "GOROOT variable is unset..."
	echo "export GOROOT=/usr/lib/go" >> $USERBASHRC
	echo "Exported GOROOT variable in bashrc."
fi

if grep -q "export GOPATH=\$HOME/go" $USERBASHRC; then
	echo "GOPATH variable already set."
else
	echo "GOPATH variable is unset..."
	echo "export GOPATH=\$HOME/go" >> $USERBASHRC
	echo "Exported GOPATH variable in bashrc."
fi

if grep -q "export PATH=\$GOROOT/bin:\$GOPATH/bin:\$PATH" $USERBASHRC; then
	echo ".bashrc is already up to date with latest PATH"
else
	echo "Updating PATH with GOROOT and GOPATH"
	echo "export PATH=\$GOROOT/bin:\$GOPATH/bin:\$PATH" >> $USERBASHRC
	echo "Updated!"
fi
