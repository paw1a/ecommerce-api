import random

from faker import Faker
from faker.providers import BaseProvider
import faker_commerce

import json
import bson

object_ids = []


class BSONProvider(BaseProvider):
    @staticmethod
    def object_id():
        object_id = str(bson.ObjectId())
        object_ids.append(object_id)
        return object_id


fake = Faker(['en_US'])
fake.add_provider(faker_commerce.Provider)
fake.add_provider(BSONProvider)

users = open('users.json', 'w')
jsonString = fake.json(data_columns={'name': 'name',
                                     'email': 'free_email',
                                     'password': 'password'}, num_rows=20)
parsed = json.loads(jsonString)
pretty = json.dumps(parsed, indent=4)
users.write(pretty)
users.close()

products = open('products.json', 'w')
fake.set_arguments('product_desc_arg', {'nb_words': 10})
fake.set_arguments('category_desc_arg', {'nb_words': 5})
categories = [{'name': 'ecommerce_category',
               'description': 'sentence:category_desc_arg'}] * random.randint(0, 5)

s = fake.json(data_columns={'_id': {"$oid": 'object_id'},
                            'name': 'ecommerce_name',
                            'description': 'sentence:product_desc_arg',
                            'price': 'ecommerce_price',
                            'categories': categories
                            }, num_rows=30)

parsed = json.loads(s)
pretty = json.dumps(parsed, indent=4)
products.write(pretty)
products.close()
