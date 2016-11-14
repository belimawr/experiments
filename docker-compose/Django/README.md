Django
======

A small experiment with Docker-compose + Django + Redis + Postgres

All django apps/models/views are based on [Django tutorial](https://docs.djangoproject.com/en/1.10/intro/tutorial01/)

Running
-------
1. ``docker-compose run --rm web mysite/manage.py migrate``: Apply all migrations
2. ``docker-compose run --rm web mysite/manage.py createsuperuser``: Create a super user
3. ``docker-compose up -d``: Start all dockers as daemons