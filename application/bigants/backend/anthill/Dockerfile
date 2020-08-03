FROM gcr.io/google-appengine/python

RUN virtualenv /env -p python3.7
ENV VIRTUAL_ENV /env
ENV PATH /env/bin:$PATH

RUN pip3 install --upgrade pip

ADD requirements.base.txt ./
RUN pip3 install -r requirements.base.txt
ADD requirements.txt ./
RUN pip3 install -r requirements.txt

ADD . ./
CMD gunicorn --bind=0.0.0.0:$PORT --access-logfile - --error-logfile - anthill.wsgi
