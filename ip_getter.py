import aiohttp
from typing import *
import json


class IpInfo(NamedTuple):
    ip: str
    full_ip: str
    country: str
    country_code: str
    province: str
    city: str
    distinct: str
    isp: str
    operator: str
    lon: str
    lat: str
    net_str: str

    @staticmethod
    def from_json(json_str: str):
        dic = json.loads(json_str)
        # return IpInfo(
        #     ip=dic['ip'],
        #     full_ip=dic['full_ip'],
        #     country=dic['country'],
        #     country_code=dic['country_code'],
        #     province=dic['province'],
        #     city=dic['city'],
        #     distinct=dic['distinct'],
        #     isp=dic['isp'],
        #     operator=dic['operator'],
        #     lon=dic['lon'],
        #     lat=dic['lat'],
        #     net_str=dic['net_str'])
        return IpInfo(**dic)


class IpGetter:
    def __init__(self):
        self.__session = aiohttp.ClientSession()
        self.__api_url = 'https://forge.speedtest.cn/api/location/info'
        self.__timeout = aiohttp.ClientTimeout(
            total=330,          # 全部请求最终完成时间
            connect=2,          # 从本机连接池里取出一个将要进行的请求的时间
            sock_connect=15,    # 单个请求连接到服务器的时间
            sock_read=10        # 单个请求从服务器返回的时间
        )
        self.__header = {
            'User-Agent': 'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:94.0) Gecko/20100101 Firefox/94.0',
        }

    async def get_ip(self):
        response = await self.__session.get(
            url=self.__api_url,
            headers=self.__header,
            timeout=self.__timeout)
        response_body_text = await response.text()
        return IpInfo.from_json(response_body_text)
