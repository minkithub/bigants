from django.db import models


class Stock(models.Model):
    code = models.CharField(primary_key=True, max_length=20)
    name_ko = models.CharField(max_length=50)

    def __str__(self):
        return f'{self.code} {self.name_ko}'

    class Meta:
        db_table = 'stock'


class Price(models.Model):
    stock_code = models.ForeignKey(Stock, on_delete=models.PROTECT)
    record_date = models.DateField()
    open_value = models.FloatField()
    high_value = models.FloatField()
    low_value = models.FloatField()
    close_value = models.FloatField()
    volume = models.IntegerField()

    def __str__(self):
        return f'{self.stock_code.code} {self.record_date} {self.close_value}'

    class Meta:
        db_table = 'price'
        indexes = [
            models.Index(fields=['stock_code', 'record_date']),
        ]
