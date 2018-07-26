def library_sort(l):
    # Initialization
    d = len(l)
    k = [None]*(d<<1)
    m = d.bit_length() # floor(log2(n) + 1)
    for i in range(d): k[2*i+1] = l[i]

    # main loop
    a,b = 1,2
    for i in range(m):
        # Because multiplication by 2 occurs at the beginning of the loop,
        # the first element will not be sorted at first pass, wich is wanted
        # (because a single element does not need to be sorted)
        a <<= 1
        b <<= 1
        for j in range(a,min(b,d+1)):
            p = 2*j-1
            s = k[p]
            # Binary search
            x, y = 0, p
            while y-x > 1:
                c = (x+y)>>1
                if k[c] != None:
                    if k[c] <  s: x = c
                    else: y = c
                else:
                    e,f = c-1,c+1
                    while k[e] == None: e -= 1
                    while k[f] == None: f += 1
                    if k[e] > s: y = e
                    elif k[f] < s: x = f
                    else:
                        x, y = e, f
                        break
            if y-x > 1: k[ (x+y)>>1 ] = s
            else:
                if k[x] != None:
                    if k[x] > s: y = x # case may occur for [2,1,0]
                    while s != None:
                        k[y], s = s, k[y]
                        y += 1
                else: k[x] = s
            k[p] = None
        if b > d: break
        if i < m-1:
            s = p
            while s >= 0:
                if k[s] != None:
                    # In the following line, the order is very important
                    # because s and p may be equal in which case we want
                    # k[s] (which is k[p]) to remain unchanged
                    k[s], k[p] = None, k[s]
                    p -= 2
                s -= 1
    return [ x for x in k if x != None ]

arr = [10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2]

mm = library_sort(arr)

print arr

print mm
