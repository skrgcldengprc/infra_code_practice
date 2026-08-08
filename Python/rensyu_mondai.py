# できる限り答えを見ずに進める。
# CHATGPTから出されたお題に対し、最初からヒントがある場合は例外
# 分からないことに対して、わからないから教えてと聞くのではなく調べる癖をつける。

# CHATGPTの練習問題1 サーバーの状態チェック

# status = "OK"

# if status == "OK":
#     print("サーバーは正常です")
# elif status == "WARNING":
#     print("サーバーに注意が必要です")
# elif status == "CRITICAL":
#     print("サーバーに異常があります")
# else:
#     print("不明なステータスです")

# CHATGPTの練習問題2 CPU使用率を判定する

# cpu_usage = 75

# if cpu_usage >= 90:
#     print("CRITICAL: CPU使用率が高すぎます")
# elif cpu_usage >=70 and cpu_usage <= 90:
#     print("WARNING: CPU使用率に注意してください")
# else:
#     print("OK: CPU使用率は正常です")

# CHATGPTの練習問題3 複数サーバーの状態を確認する

# servers = ["OK","WARNING","OK","CRITICAL","OK"]

# for i in servers:

#     if i == "OK":
#         print("サーバーは正常です")
#     elif i == "WARNING":
#         print("サーバーに注意が必要です")
#     elif i == "CRITICAL":
#         print("サーバーに異常があります")
#     else:
#         print("サーバーの状態は不明です")

# CHATGPTの練習問題4 異常なサーバーを数える

# count = 0
# servers = ["OK","WARNING","OK","CRITICAL","OK"]

# for i in servers:

#     if i == "OK":
#         print("サーバーは正常です")
#     elif i == "WARNING":
#         print("サーバーに注意が必要です")
#     elif i == "CRITICAL":
#         print("サーバーに異常があります")
#         count += 1
#     else:
#         print("サーバーの状態は不明です")  

# print(f"CRITICALなサーバーは{count}台です")

# CHATGPTの練習問題5 正常なサーバーだけ別のリストに入れる

# servers = ["OK","WARNING","OK","CRITICAL","OK"]
# ok_servers = []

# for i in servers:

#     if i == "OK":
#         print("サーバーは正常です")
#         ok_servers.append(i)
#     elif i == "WARNING":
#         print("サーバーに注意が必要です")
#     elif i == "CRITICAL":
#         print("サーバーに異常があります")
#     else:
#         print("サーバーの状態は不明です")
# print(f"OKのサーバーは{len(ok_servers)}台です")

# CHATGPTの練習問題6 サーバー名と状態を一緒に扱う

# servers = {
#     "web01": "OK",
#     "web02": "CRITICAL",
#     "web03": "OK",
#     "web04": "WARNING"
# }

# for name, status in servers.items():
#     if status == "CRITICAL":
#         print(f"{name}は異常です")

# CHATGPTの練習問題7 障害サーバーの一覧を作る

# servers = {
#     "web01": "OK",
#     "web02": "CRITICAL",
#     "web03": "OK",
#     "web04": "WARNING",
#     "web05": "CRITICAL"
# }

# critical_servers = []
# for name, status in servers.items():
#     if status == "CRITICAL":
#         critical_servers.append(name)
# print(f"CRITICALなサーバーは{critical_servers}です")

# CHATGPTの練習問題8 サーバーの状態を集計する

# servers = {
#     "web01": "OK",
#     "web02": "CRITICAL",
#     "web03": "OK",
#     "web04": "WARNING",
#     "web05": "CRITICAL",
#     "web06": "OK"
# }

# ok_count = 0
# warning_count = 0
# critical_count = 0

# for name, status in servers.items():
#     if status == "OK":
#         ok_count +=1
#     elif status == "CRITICAL":
#         critical_count += 1
#     elif status == "WARNING":
#         warning_count +=1
#     else:
#         print(f"不明なステータス{name}です")

# print(f"正常{ok_count}台")
# print(f"警告{warning_count}台")
# print(f"異常{critical_count}台")

# CHATGPTの練習問題9 関数にしてみよう

# servers = {
#     "web01": "OK",
#     "web02": "CRITICAL",
#     "web03": "OK",
#     "web04": "CRITICAL"
# }

# def check_servers(servers):
#     critical_server = 0
#     for name, status in servers.items():
#         if status == "CRITICAL":
#             critical_server += 1

#     return critical_server

# result = check_servers(servers)

# print(f"CRITICALなサーバは{result}台です")

# CHATGPTの練習問題10 CRITICALなサーバー一覧を返す関数

# servers = {
#     "web01": "OK",
#     "web02": "CRITICAL",
#     "web03": "OK",
#     "web04": "CRITICAL",
#     "web05": "WARNING"
# }

# def check_servers(servers):
#     critical_servers = []
#     for name , status in servers.items():
#         if status == "CRITICAL":
#             critical_servers.append(name)

#     return critical_servers

# result = check_servers(servers)
# print(f"CRITICALなサーバー:{result}")

# CHATGPTの練習問題11 異常サーバーをまとめて取得する

# servers = {
#     "web01": "OK",
#     "web02": "CRITICAL",
#     "web03": "OK",
#     "web04": "WARNING",
#     "web05": "CRITICAL",
#     "web06": "OK"
# }

# def check_servers(servers):
#     critical_server = []
#     for name, status in servers.items():
#         if status == "CRITICAL" or status == "WARNING":
#             critical_server.append(name)
    
#     return critical_server

# result = (check_servers(servers))

# print(f"異常・注意が必要なサーバー：{result}")

# CHATGPTの練習問題12 状態を入力して検索する

# servers = {
#     "web01": "OK",
#     "web02": "CRITICAL",
#     "web03": "OK",
#     "web04": "WARNING",
#     "web05": "CRITICAL",
#     "web06": "OK"
# }

# def check_servers(servers):
#     ok_servers = []
#     warning_servers = []
#     critical_servers = []

#     for name, status in servers.items():
#         if status == "OK":
#             ok_servers.append(name)
#         elif status == "WARNING":
#             warning_servers.append(name)
#         elif status == "CRITICAL":
#             critical_servers.append(name)

#     return {
#         "OK": ok_servers,
#         "WARNING": warning_servers,
#         "CRITICAL": critical_servers
#     }

# result = check_servers(servers)

# input_result = input("確認したい状態を入力してください：")

# if input_result == "OK":
#     print(result["OK"])
# elif input_result == "WARNING":
#     print(result["WARNING"])
# elif input_result == "CRITICAL":
#     print(result["CRITICAL"])

# CHATGPTの練習問題13 サーバー名から状態を取得する

#自分の答え

# servers = {
#     "web01": "OK",
#     "web02": "CRITICAL",
#     "web03": "OK",
#     "web04": "WARNING",
#     "web05": "CRITICAL"
# }
# for name, status in servers.items():

#   server_name = input("確認したいサーバー名を入力してください：")
#   if server_name in name:
#       print(f"{name}の状態は{status}です")
#       break
#   else:
#       print(f"指定されたサーバー{server_name}は存在しません")
#       continue

#CHATGPTの答え

servers = {
    "web01": "OK",
    "web02": "CRITICAL",
    "web03": "OK",
    "web04": "WARNING",
    "web05": "CRITICAL"
}

server_name = input("確認したいサーバー名を入力してください：")

if server_name in servers:
    print(f"{server_name}の状態は{servers[server_name]}です")
else:
    print(f"指定されたサーバー{server_name}は存在しません")