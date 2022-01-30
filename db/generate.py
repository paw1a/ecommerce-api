import os
import random

from faker import Faker
import faker_commerce

import json
import bson

USER_NUM = 20
PRODUCT_NUM = 20
REVIEW_NUM = 10

if not os.path.exists('data'):
    os.makedirs('data')

fake = Faker(['en_US'])
fake.add_provider(faker_commerce.Provider)

# generate users
users = open('data/users.json', 'w')
userList = []
userIds = []
for _ in range(USER_NUM):
    userId = str(bson.ObjectId())
    userIds.append(userId)
    user = fake.json(data_columns={'_id': {"$oid": f'@{userId}'},
                                   'name': 'name',
                                   'email': 'free_email',
                                   'password': 'password'}, num_rows=1)
    parsedUser = json.loads(user)
    userList.append(parsedUser)

data = json.dumps(userList, indent=4)
users.write(data)
users.close()

# generate products
products = open('data/products.json', 'w')
fake.set_arguments('product_desc_arg', {'nb_words': 10})
fake.set_arguments('category_desc_arg', {'nb_words': 5})

productList = []
productIds = []
for _ in range(PRODUCT_NUM):
    productId = str(bson.ObjectId())
    productIds.append(productId)
    categoriesFormatter = [{'name': 'ecommerce_category',
                            'description': 'sentence:category_desc_arg'}] * random.randint(0, 5)

    product = fake.json(data_columns={'_id': {"$oid": f'@{productId}'},
                                      'name': 'ecommerce_name',
                                      'description': 'sentence:product_desc_arg',
                                      'price': 'ecommerce_price',
                                      'categories': categoriesFormatter}, num_rows=1)

    parsedProduct = json.loads(product)
    productList.append(parsedProduct)

data = json.dumps(productList, indent=4)
products.write(data)
products.close()

# generate reviews
fake.set_arguments('rating', {'min_value': 1, 'max_value': 5})
reviews = open('data/reviews.json', 'w')
reviewsList = []
for productId in productIds:
    for _ in range(random.randint(0, REVIEW_NUM)):
        reviewId = str(bson.ObjectId())
        reviewText = fake.text().replace('\n', ' ')
        review = fake.json(data_columns={
            '_id': {"$oid": f'@{reviewId}'},
            'text': f'@{reviewText}',
            'rating': f'pyint:rating',
            'productID': f'@{productId}',
            'userID': f'@{random.choice(userIds)}'
        }, num_rows=1)
        parsedReview = json.loads(review)
        reviewsList.append(parsedReview)

data = json.dumps(reviewsList, indent=4)
reviews.write(data)
reviews.close()

# generate admins
admins = open('data/admins.json', 'w')
data = [
    {
        'name': 'Admin',
        'email': 'paw1a@yandex.ru',
        'password': '123'
    },
    {
        'name': 'Admin 2',
        'email': 'admin@admin.com',
        'password': 'admin'
    }
]
data = json.dumps(data, indent=4)
admins.write(data)
admins.close()
