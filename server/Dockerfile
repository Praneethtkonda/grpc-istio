# pull official base image
FROM team-sigma-docker-local.artifactory.eng.vmware.com/signing-service/0.1/python:3.9.5-slim-buster

# set work directory
WORKDIR /app

# set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# install dependencies
RUN pip install --upgrade pip
COPY ./requirements.txt .
RUN pip install -r requirements.txt

# copy project
COPY . .

# Start the server
EXPOSE 50051

CMD ["python", "server.py"]