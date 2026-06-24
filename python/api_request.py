from urllib.request import urlopen

with urlopen("https://api.github.com") as response:
    print("Status Code:", response.status)
