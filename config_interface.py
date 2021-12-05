from abc import abstractmethod
from mail import EmailConfig, EmailReceiver
from typing import *


class IConfig:
    @abstractmethod
    def get_email_config(self) -> EmailConfig:
        pass

    @abstractmethod
    def get_receivers(self) -> List[EmailReceiver]:
        pass

    @abstractmethod
    def get_sender_name(self) -> str:
        pass
