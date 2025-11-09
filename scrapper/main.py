import os
from dotenv import load_dotenv
import requests
from bs4 import BeautifulSoup
from lxml import etree, html
import re

# This gets me 15 out of 74 pins


def main():
    load_dotenv()
    url = os.getenv("URL")

    response = requests.get(url)
    resp_html = html.fromstring(response.content)
    links = []
    links_element_all = list(resp_html.iterlinks())
    for link in links_element_all:
        (element, attribute, link, pos) = link
        # loking for link like /pin/[0-9]*
        pin_link = re.search("^/pin/[0-9]*/$", link)
        if pin_link and link not in links:
            links.append(link)

    for link in links:
        print("www.pinterest.com" + link)


if __name__ == "__main__":
    main()
