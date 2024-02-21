from fake_useragent import UserAgent

LENGTH = 10000

ua = UserAgent()
random_ua_list = [f"{ua.random}\n" for _ in range(LENGTH)]
with open("public/random_useragent.txt", "w") as f:
    f.writelines(random_ua_list)