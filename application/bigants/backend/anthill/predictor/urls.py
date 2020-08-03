from django.urls import path

from . import views

app_name = 'predictor'
urlpatterns = [
    path('', views.IndexView.as_view(), name='index'),
    path('stocks', views.get_stocks, name='stocks'),
    path('history', views.get_history, name='history'),
    path('predict', views.get_predict, name='predict'),
    path('update', views.update_prices, name='update'),
    path('result', views.result, name='result'),
]
