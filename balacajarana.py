n = int(input("N: "))

for i in range(1,n+1):
    print("Número %d:" %i, end=" ")
    if i%3 == 0:
        print("bala ", end="")
    if i % 7 == 0:
        print("cajarana", end="")
    print("")#quebra linha