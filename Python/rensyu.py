# Pythonの点数結果プログラム

benkyo = input('点数結果を入力してください：')
try:
    benkyo = int(benkyo)
except ValueError:
    print("点数は数字で入力してください")
    exit()

if benkyo >= 80:
    answer = "大変よくできました！"
elif benkyo >= 60:
    answer = "よくできました！"
elif benkyo >= 40:
    answer = "もう少し頑張りましょう"
elif benkyo >= 1:
    answer = "もっと頑張りましょう"
else:
    answer = "点数は1以上で入力してください"

print(f'{benkyo}点ですね。{answer}')