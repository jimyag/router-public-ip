from config_interface import *
from typing import *


class MyConfig(IConfig):
    def get_email_config(self) -> EmailConfig:
        return EmailConfig(
            smtp_server='',
            smtp_port=0,
            use_tls=True,
            user_name='',
            password='')

    def get_sender_name(self) -> str:
        return ''

    def get_receivers(self) -> List[EmailReceiver]:
        return [
            EmailReceiver('', '')
        ]


my_config = MyConfig()
