"""Django settings for FeedVault project.

Generated by 'django-admin startproject' using Django 5.0.1.

For more information on this file, see
https://docs.djangoproject.com/en/5.0/topics/settings/

For the full list of settings and their values, see
https://docs.djangoproject.com/en/5.0/ref/settings/
"""

from __future__ import annotations

import os
from pathlib import Path

from dotenv import find_dotenv, load_dotenv

load_dotenv(dotenv_path=find_dotenv(), verbose=True)

# Build paths inside the project like this: BASE_DIR / 'subdir'.
BASE_DIR: Path = Path(__file__).resolve().parent.parent

# The secret key is used for cryptographic signing, and should be set to a unique, unpredictable value.
SECRET_KEY: str = os.getenv("SECRET_KEY", default="")

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG: bool = os.getenv(key="DEBUG", default="True").lower() == "true"

# A list of all the people who get code error notifications. When DEBUG=False and a view raises an exception
# Django will email these people with the full exception information.
ADMINS: list[tuple[str, str]] = [("Joakim Hellsén", "django@feedvault.se")]

# A list of strings representing the host/domain names that this Django site can serve.
# .feedvault.se will match *.feedvault.se and feedvault.se
ALLOWED_HOSTS: list[str] = [".feedvault.se", ".localhost", "127.0.0.1"]

# The time zone that Django will use to display datetimes in templates and to interpret datetimes entered in forms
TIME_ZONE: str = os.getenv(key="TZ", default="Europe/Stockholm")

# If datetimes will be timezone-aware by default. If True, Django will use timezone-aware datetimes internally.
USE_TZ = True

# Turn off Django's translation system
USE_I18N = False

# Decides which translation is served to all users.
LANGUAGE_CODE = "en-us"

# Default decimal separator used when formatting decimal numbers.
DECIMAL_SEPARATOR = ","

# Use a space as the thousand separator instead of a comma
THOUSAND_SEPARATOR = " "

# Use gmail for sending emails
EMAIL_HOST = "smtp.gmail.com"
EMAIL_PORT = 587
EMAIL_USE_TLS = True
EMAIL_HOST_USER: str = os.getenv(key="EMAIL_HOST_USER", default="webmaster@localhost")
EMAIL_HOST_PASSWORD: str = os.getenv(key="EMAIL_HOST_PASSWORD", default="")
EMAIL_SUBJECT_PREFIX = "[FeedVault] "
EMAIL_USE_LOCALTIME = True
EMAIL_TIMEOUT = 10
DEFAULT_FROM_EMAIL: str = os.getenv(key="EMAIL_HOST_USER", default="webmaster@localhost")
SERVER_EMAIL: str = os.getenv(key="EMAIL_HOST_USER", default="webmaster@localhost")

# Use the X-Forwarded-Host header
USE_X_FORWARDED_HOST = True

# Set the Referrer Policy HTTP header on all responses that do not already have one.
SECURE_REFERRER_POLICY = "strict-origin-when-cross-origin"

# Full Python import path to our main URL configuration.
ROOT_URLCONF = "feedvault.urls"

# URL to use when referring to static files located in STATIC_ROOT.
# https://feedvault.se/static/...
STATIC_URL = "static/"

# Use a 64-bit integer as a primary key for models that don't have a field with primary_key=True.
DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"

# A string representing the full Python import path to our WSGI application object
WSGI_APPLICATION = "feedvault.wsgi.application"

# Internal IPs that are allowed to see debug views
INTERNAL_IPS: list[str] = ["127.0.0.1", "localhost"]

# Applications include some combination of models, views, templates, template tags, static files, URLs, middleware, etc
INSTALLED_APPS: list[str] = [
    # First-party apps
    "feeds.apps.FeedsConfig",
    # Third-party apps
    "whitenoise.runserver_nostatic",  # https://whitenoise.readthedocs.io/en/latest/index.html
    "debug_toolbar",  # https://github.com/jazzband/django-debug-toolbar/
    "django_celery_results",  # https://github.com/celery/django-celery-results
    # Django apps
    "django.contrib.sites",
    "django.contrib.admin",
    "django.contrib.sitemaps",
    "django.contrib.auth",
    "django.contrib.contenttypes",
    "django.contrib.sessions",
    "django.contrib.messages",
    "django.contrib.staticfiles",
]

# Middleware is a framework of hooks into Django's request/response processing.
MIDDLEWARE: list[str] = [
    "django.middleware.security.SecurityMiddleware",
    "whitenoise.middleware.WhiteNoiseMiddleware",
    "debug_toolbar.middleware.DebugToolbarMiddleware",
    "django.contrib.sessions.middleware.SessionMiddleware",
    "django.middleware.common.CommonMiddleware",
    "django.middleware.csrf.CsrfViewMiddleware",
    "django.contrib.auth.middleware.AuthenticationMiddleware",
    "django.contrib.messages.middleware.MessageMiddleware",
    "django.middleware.clickjacking.XFrameOptionsMiddleware",
]

# A list containing the settings for all template engines to be used with Django.
TEMPLATES = [
    {
        "BACKEND": "django.template.backends.django.DjangoTemplates",
        "DIRS": [BASE_DIR / "templates"],
        "OPTIONS": {
            "context_processors": [
                "django.template.context_processors.debug",
                "django.template.context_processors.request",
                "django.contrib.auth.context_processors.auth",
                "django.contrib.messages.context_processors.messages",
            ],
            "loaders": [
                (
                    "django.template.loaders.cached.Loader",
                    [
                        "django.template.loaders.filesystem.Loader",
                        "django.template.loaders.app_directories.Loader",
                    ],
                ),
            ],
        },
    },
]

# A dictionary containing the settings for how we should connect to our PostgreSQL database.
DATABASES: dict[str, dict[str, str]] = {
    "default": {
        "ENGINE": "django.db.backends.postgresql",
        "NAME": "feedvault",
        "USER": os.getenv(key="PGUSER", default=""),
        "PASSWORD": os.getenv(key="PGPASSWORD", default=""),
        "HOST": os.getenv(key="PGHOST", default=""),
        "PORT": os.getenv(key="PGPORT", default="5432"),
    },
}

# The absolute path to the directory where 'python manage.py collectstatic' will copy static files for deployment
STATIC_ROOT: Path = BASE_DIR / "staticfiles"
STATICFILES_DIRS: list[Path] = [BASE_DIR / "static"]

# Use WhiteNoise to serve static files. https://whitenoise.readthedocs.io/en/latest/django.html
STORAGES: dict[str, dict[str, str]] = {
    "staticfiles": {"BACKEND": "whitenoise.storage.CompressedManifestStaticFilesStorage"},
}

# Our site ID
SITE_ID = 1
