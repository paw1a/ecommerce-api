import React, {useEffect, useState} from 'react';
import './Catalog.css'
import Filter from "./Filter";
import {Button, Grid, Pagination} from "@mui/material";
import axios from "axios";
import ProductCard from "./Card";

const Catalog = () => {
    let [products, setProducts] = useState([]);
    useEffect(() => {
        axios.get('http://52.29.184.51:8080/api/v1/products/')
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
                <Button className='show-more'>Показать еще</Button>
                <div className='pagination-container'>
                    <Pagination count={10} shape="rounded" />
                </div>
            </div>
        </div>
    );
};

export default Catalog;