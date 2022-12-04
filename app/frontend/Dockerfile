####################
# Dependencies
####################
FROM node:18-alpine as dep

WORKDIR /app

COPY ./package.json ./yarn.lock ./

RUN yarn install

####################
# Development Image
####################
FROM dep as development
EXPOSE 3000

WORKDIR /app

CMD yarn dev
