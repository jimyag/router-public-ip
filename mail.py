import aiosmtplib
# 负责构造文本
from email.mime.text import MIMEText
# 负责将多个对象集合起来
from email.mime.multipart import MIMEMultipart
from email.header import Header
from typing import *


class EmailReceiver(NamedTuple):
    name: str
    email: str


class EmailConfig(NamedTuple):
    smtp_server: str
    smtp_port: int
    use_tls: bool
    user_name: str
    password: str


class EmailSender:
    def __init__(self,
                 email_config: EmailConfig):
        self.__smtp = aiosmtplib.SMTP(
            hostname=email_config.smtp_server,
            port=email_config.smtp_port,
            use_tls=email_config.use_tls)
        self.__user_name = email_config.user_name
        self.__password = email_config.password

    async def login(self):
        await self.__smtp.__aenter__()
        return await self.__smtp.login(self.__user_name, self.__password)

    async def send_mail(self,
                        receivers: List[EmailReceiver],
                        sender_name: str,
                        subject: str,
                        message_text: str):
        # 构造一封邮件
        mm = MIMEMultipart('related')
        mm['From'] = f'{sender_name}<{self.__user_name}>'
        send_to = ''
        for receiver in receivers:
            send_to += f'{receiver[0]}<{receiver[1]}>,'
        mm['To'] = send_to[:-1]
        mm['Subject'] = Header(subject, 'utf-8')
        mm.attach(MIMEText(message_text, 'html', 'utf-8'))

        # 发送邮件
        await self.__smtp.sendmail(
            self.__user_name,
            list(map(lambda x: x[1], receivers)),
            mm.as_string())

    async def quit(self):
        await self.__smtp.quit()
