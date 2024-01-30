# Generated by Django 5.0.1 on 2024-01-30 16:45

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('feeds', '0002_blocklist'),
    ]

    operations = [
        migrations.AddField(
            model_name='blocklist',
            name='active',
            field=models.BooleanField(default=True, help_text='Is this URL still blocked?'),
        ),
    ]
