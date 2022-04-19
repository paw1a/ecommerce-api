import React, {useEffect, useState} from 'react';
import axios from 'axios'
import ProductItem from "./ProductItem";
import {Grid} from "@mui/material";
import ProductCard from "./catalog/Card";

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
        <div className='productList'>
            <h1>Product List</h1>
            <Grid container spacing={2}>
                <>
                    {products.map(product => (
                        <Grid item xs={4} key={product.id}>
                            <ProductCard product={product} />
                        </Grid>
                    ))}
                </>
            </Grid>
        </div>
    );
};

export default ProductList;