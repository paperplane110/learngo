a=1

b=$(echo ${a}+0.1 | bc)
c=$(python -c "print(${a}+0.1)")

echo ${b}
echo ${c}