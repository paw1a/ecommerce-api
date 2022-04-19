import React, {useEffect, useState} from 'react';
import './Catalog.css'
import Filter from "./Filter";
import {Grid} from "@mui/material";
import ProductItem from "../ProductItem";
import axios from "axios";
import ProductCard from "./Card";

const Catalog = () => {
    let [products, setProducts] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/api/v1/products/')
            .then((resp) => {
                const products = resp.data.data;
                setProducts(products);
            });
    }, [setProducts]);

    console.log(products)

    return (
        <div className='catalog'>
            <Filter/>
            <div className='product-catalog' style={{marginTop: '30px'}}>
                <Grid container columnSpacing={2}>
                    {products.map(product => (
                        <Grid item xs={6} md={4} key={product.id} style={{marginBottom: '20px'}}>
                            <ProductCard product={product}/>
                        </Grid>
                    ))}
                </Grid>
            </div>
        </div>
    );
};

export default Catalog;