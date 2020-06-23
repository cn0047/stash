from enum import Enum
from typing import Tuple
from dataclasses import dataclass


class Height(Enum):
  SHORT = 0
  TALL = 1


@dataclass
class BaseConfig:
  num: int = 10
  pi: float = 3.1415
  is_awesome: bool = True
  height: Height = Height.SHORT
  description: str = 'text'
  run_command: Tuple[str] = ('sh', '-c', 'echo', '200')



@dataclass
class WebAppConf(BaseConfig):
  name: str = 'WebApp'


c = WebAppConf()
print(c)
