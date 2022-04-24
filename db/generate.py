import os
import random
import time

import bson
from faker import Faker
import faker_commerce

import json
import stripe

from pymongo import MongoClient

USER_NUM = 10
PRODUCT_NUM = 10
REVIEW_NUM = 6

stripe.api_key = os.getenv('STRIPE_KEY')
CONNECTION_STRING = "mongodb://mongo:27017"

client = MongoClient(CONNECTION_STRING)
db = client['ecommerce']

fake = Faker(['en_US'])
fake.add_provider(faker_commerce.Provider)

# generate users
user_list = []

for _ in range(USER_NUM):
    user = fake.json(data_columns={'name': 'name',
                                   'email': 'free_email',
                                   'password': 'password'}, num_rows=1)
    parsed_user = json.loads(user)
    user_list.append(parsed_user)

user_collection = db['users']
user_collection.drop()
user_collection.insert_many(user_list)

# generate products
fake.set_arguments('product_desc_arg', {'nb_words': 40})
fake.set_arguments('category_desc_arg', {'nb_words': 5})
fake.set_arguments('price', {'min_value': 100, 'max_value': 100000})

product_list = []

for _ in range(PRODUCT_NUM):
    categories_formatter = [{'name': 'ecommerce_category',
                            'description': 'sentence:category_desc_arg'}] * random.randint(0, 5)

    product: str = fake.json(data_columns={'name': 'ecommerce_name',
                                           'description': 'sentence:product_desc_arg',
                                           'price': 'pyint:price',
                                           'categories': categories_formatter}, num_rows=1)

    parsed_product = json.loads(product)
    product_list.append(parsed_product)

product_collection = db['products']
product_collection.drop()
product_collection.insert_many(product_list)

for product in product_list:
    product: dict
    stripe.Product.create(name=product["name"],
                          description=product["description"],
                          id=product['_id'])
    stripe.Price.create(unit_amount_decimal=product["price"] * 100,
                        currency="RUB",
                        product=product['_id'])

# generate reviews
fake.set_arguments('rating', {'min_value': 1, 'max_value': 5})

reviews_list = []

for product in product_list:
    product: dict
    for _ in range(random.randint(REVIEW_NUM // 2, REVIEW_NUM)):
        reviewText = fake.text().replace('\n', ' ')
        review = fake.json(data_columns={
            'text': f'@{reviewText}',
            'rating': f'pyint:rating',
        }, num_rows=1)

        parsed_review: dict = json.loads(review)
        parsed_review['productID'] = product['_id']
        user = random.choice(user_list)
        parsed_review['userID'] = user['_id']
        parsed_review['username'] = user['name']
        date = bson.timestamp.Timestamp(int(time.time()), 0)
        parsed_review['date'] = date

        reviews_list.append(parsed_review)

review_collection = db['reviews']
review_collection.drop()
review_collection.insert_many(reviews_list)

# generate admins
admin_collection = db['admins']
admin_collection.drop()

admin_1 = {'name': 'Admin',
           'email': 'paw1a@yandex.ru',
           'password': '123'}
admin_collection.insert_one(admin_1)

admin_2 = {'name': 'Admin 2',
           'email': 'admin@admin.com',
           'password': 'admin'}
admin_collection.insert_one(admin_2)

# clean up carts and orders collections
carts_collection = db['carts']
carts_collection.drop()

orders_collection = db['orders']
orders_collection.drop()
