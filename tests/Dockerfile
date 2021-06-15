FROM python:3.9-slim

WORKDIR /tests

RUN apt update -y ; apt install -y curl
RUN curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | POETRY_HOME=/opt/poetry python && \
    cd /usr/local/bin && \
    ln -s /opt/poetry/bin/poetry && \
    poetry config virtualenvs.create false
COPY pyproject.toml ./poetry.lock* /tests/
