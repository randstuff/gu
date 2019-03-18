#!/bin/sh



for l in $(grep -Rh "func Test" * | grep "testing.T" | grep Test | cut -d " " -f 2  | cut -d "(" -f1
) ; do 
	echo go test -v -run  $l
done 





