###
python3.10 -m venv env

###
source env/bin/activate

deactivate

###
pip freeze > requirements.txt

pip install -r requirements.txt
