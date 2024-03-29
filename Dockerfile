FROM ubuntu:latest
EXPOSE 8000
WORKDIR /app
ENV PORT=8000 
ENV DATABASE_URL=localhost DATABASE_DBPORT=5432 
ENV DATABASE_USER=root DATABASE_PWD=root DATABASE_DBNAME=root
COPY ./main main
RUN chmod +x main
COPY ./templates/ templates/
CMD [ "./main" ]