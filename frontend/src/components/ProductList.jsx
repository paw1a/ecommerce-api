import React, {useEffect, useState} from 'react';
import axios from 'axios'
import ProductItem from "./ProductItem";

const ProductList = () => {
    let [products, setProducts] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/api/v1/products/')
            .then((resp) => {
            const products = resp.data.data;
            console.log(typeof products);
            setProducts(products);
        });
    }, [setProducts]);

    console.log(products)

    return (
        <div>
            <h1>Product List</h1>
            {products.map(product => (
                <div>
                    <ProductItem product={product}/>
                    <br/>
                </div>
            ))}
        </div>
    );
};

export default ProductList;