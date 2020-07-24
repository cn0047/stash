from PyQt5.QtWidgets  import *
import sys


app = QApplication(sys.argv)
d = QMainWindow()
d.setGeometry(0, 0, 200, 100)
d.setWindowTitle("It works!")
d.show()
sys.exit(app.exec_())
