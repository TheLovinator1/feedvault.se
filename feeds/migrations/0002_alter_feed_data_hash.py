# Generated by Django 5.0.6 on 2024-05-20 01:19

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('feeds', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='feed',
            name='data_hash',
            field=models.BinaryField(help_text='The hash of the feed data.', null=True),
        ),
    ]