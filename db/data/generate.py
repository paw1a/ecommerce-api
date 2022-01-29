import random

from faker import Faker
import faker_commerce

import json
import bson

fake = Faker(['en_US'])
fake.add_provider(faker_commerce.Provider)

# generate users
users = open('users.json', 'w')
userList = []
userIds = []
for _ in range(10):
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
products = open('products.json', 'w')
fake.set_arguments('product_desc_arg', {'nb_words': 10})
fake.set_arguments('category_desc_arg', {'nb_words': 5})

productList = []
productIds = []
for _ in range(10):
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

reviews = []
for productId in productIds:
    for _ in range(random.randint(0, 10)):
        reviewId = str(bson.ObjectId())
        review = fake.json(data_columns={
            '_id': f'@{reviewId}',
            'text': 'text',
            'rating': 'pyint',
            'productID': f'@{productId}',
            'userID': f'@{random.choice(userIds)}'
        })

