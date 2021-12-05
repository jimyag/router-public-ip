from email import message
import json
import asyncio
from config_interface import IConfig
from mail import EmailSender
from ip_getter import IpGetter


class AutoIpReportSystem:
    def __init__(self, _config: IConfig):
        self.__ip_getter = IpGetter()
        self.__email_sender = EmailSender(_config.get_email_config())
        self.__config = _config

    async def __aenter__(self):
        await self.__email_sender.login()
        return self

    async def __aexit__(self, exc_type, exc_val, exc_tb):
        await self.close()

    async def send_ip(self):
        ip_info = await self.__ip_getter.get_ip()
        await self.__email_sender.send_mail(
            receivers=self.__config.get_receivers(),
            sender_name=self.__config.get_sender_name(),
            message_text=str(ip_info)
        )


async def main():
    from config import my_config
    async with AutoIpReportSystem(my_config) as ip_report_system:
        await ip_report_system.send_ip()

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())
    loop.close()
