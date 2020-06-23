from typing import List, Set, Dict, Tuple, Optional


x: int = 1
x: float = 1.0
x: bool = True
x: str = "test"
x: bytes = b"test"

x: List[int] = [1]
x: Set[int] = {6, 7}
x: Dict[str, float] = {'field': 2.0}
x: Tuple[int, ...] = (1, 2, 3)
x: Tuple[int, str, float] = (3, "yes", 7.5)


print(x)
