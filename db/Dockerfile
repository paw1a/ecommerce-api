FROM python:alpine3.9

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

CMD [ "python", "./generate.py"]