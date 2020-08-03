import django
import pandas as pd
import sys
import os
import datetime

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'anthill.settings')
django.setup()

from predictor.models import Stock, Price

def csv_to_model(path):
    stock_code = int(os.path.basename(path).split('.')[0])
    df = pd.read_csv(path)

    for i, row in df.iterrows():
        if row.isnull().values.any():
            continue

        record_date, open_value, high_value, low_value, close_value, adj_close_value, volume = row
        try:
            record = Price(
                stock_code=Stock.objects.get(pk=stock_code),
                record_date=record_date,
                open_value=open_value,
                high_value=high_value,
                low_value=low_value,
                close_value=close_value,
                adj_close_value=adj_close_value,
                volume=volume
            )
            record.save()
        except:
            print('fail', i)


def main():
    if sys.argv[1] == 'clearall':
        for obj in Price.objects.all():
            obj.delete()
    elif len(sys.argv) != 2:
        print('usage: csv_to_model.py [csv_path]')
    else:
        csv_to_model(sys.argv[1])


if __name__ == '__main__':
    main()
