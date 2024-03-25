#!/bin/bash

replacement_link="$1"

if [ -z "$replacement_link" ]; then
	echo "Usage: $0 <replacement_link>"
	exit 1
fi

find . -type f -exec grep -l -v 'utl' {} + | xargs sed -i "s~github\.com\/tarqeem\/template$~$replacement_link~"

rm -rf utl translate .git README.org LICENSE init.sh
