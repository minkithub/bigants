from django.contrib import admin
from .models import Stock, Price



def hello_admin_site_action(modeladmin, request, queryset):
  print("hi")


admin.site.register(Stock)
admin.site.register(Price)

admin.site.add_action(hello_admin_site_action)
