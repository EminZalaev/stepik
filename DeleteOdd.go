func (x uint) uint{
    if x==0{
     return 100
    }else{
        s:=strconv.FormatUint(uint64(x), 10)
        f:= []rune(s)
        l:=make([]rune,len(f))
        var j int
        for i:=0;i<len(f);i++{
            if f[i]=='2'||f[i]=='4'||f[i]=='6'||f[i]=='8'{
             l[j]=f[i]
                j++
            }
        }
        p:=make([]rune,len(l))
        var o int
        for i:=0;i<len(l);i++{
            if l[i]==0{
             break
            }else{
             p[i]=l[i]
                o++
            }
        }
        u:=make([]rune,o)
        for i:=0;i<len(u);i++{
            u[i]=p[i]

        }
        q:=string(u)
        number, _ := strconv.ParseUint(string(q), 10,64)
        number1:=uint(number)
        return number1
    }
}
