# 関数の中で定義された変数は、関数の外では使えない

a = 1

def fun(a):
    aiueo = a + 1
    print(aiueo)
    print(id(aiueo))

    return aiueo

fun(a)
print(a)  # これはエラーにならない
#print(aiueo)  # これはエラーになる
