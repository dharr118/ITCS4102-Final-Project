"""Simple HTTP API request example in Python."""

from urllib.request import urlopen


def main() -> None:
    with urlopen("https://api.github.com", timeout=10) as response:
        print("Status code:", response.status)
        print("Content type:", response.headers.get("Content-Type"))


if __name__ == "__main__":
    main()
