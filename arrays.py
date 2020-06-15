a1 = [int(x) for x in input('Array 1:').split(' ')]
#se for com string então fica  mais simples
# a1 = input('Array 1:').split(' ')
a2 = [int(x) for x in input('Array 2:').split(' ')]

intersecao = []

for item in a1:
    if item in a2:
        intersecao.append(item)
        # custa menos processamento armazenar a interseção e imprimir tudo de uma vez
        # do que imprimir conforme itera os arrays

print("Interseção: ", intersecao)