Go has only one loop keyword:
for

this is basic style for loop

for init; condition; increment/dercre{

}

** in go lang don't have while loop but you can you for loop as a while loop ** see below example

for condition{
    //body
}
num:=5
for num < 0{
fmt.Print(num)
num--;
}


//range arr gives you each element (val) without manually using arr[i]
for i,val := range arr{

}