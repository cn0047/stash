pip
-

````sh
pip3 -V

PIP_CONFIG_FILE=/tmp/pip.conf

pip3 list
pip3 list -v # show dir for packages

pip3 show requests
pip3 search requests

pip3 install realtimelog                # pylint: disable=invalid-name
pip3 install -U realtimelog             # upgrade
pip3 install --user realtimelog         # install to user dir (not into system dir)
pip3 install --upgrade realtimelog
pip3 install -r requirements.txt
pip3 install pylint requests virtualenv
pip3 install --find-links https://download.pytorch.org/whl/torch_stable.html torchvision==0.7.0+cu101
pip3 install --no-cache --verbose --no-index realtimelog
pip3 install /app/my_prj
pip3 install tensorflow==1.14 --force-reinstall

pip3 uninstall requests

pip3 config -v list

python3 -m pip install requests # install for python3
python3 -m pip freeze > requirements.txt


python3 -m pip install --user --upgrade setuptools wheel twine
rm -rf  build/ dist/ realtimelog.egg-info/
python3 setup.py sdist bdist_wheel
python3 -m twine upload dist/*
python3 -m twine upload --repository testpypi dist/*
# token @see: https://pypi.org/manage/account/
#   Enter your username: __token__
#   Enter your password: $token
````
