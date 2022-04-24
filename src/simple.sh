a=%d

# this file cannot be executed directly

while [ $a -gt 0 ]
do
  echo '%s'
  a=`expr $a - 1`
done
