a=%d

while [ $a -gt 0 ]
do
  echo '%s'
  a=`expr $a - 1`
done
